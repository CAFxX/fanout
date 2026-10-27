package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint; F1 string; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint }
type GenS4 struct { F0 uint64; F1 int; F2 int }
type GenS5 struct { F0 uint64; F1 int; F2 uint16 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 int64; F1 int32; F2 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int; F1 bool }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 bool; F1 int16 }
func genFn0() uint16 {
    switch 128 {
    case -2147483648:
        var v5 int = 32768
        _ = v5
    case 876:
        var v6 uint = 105
        _ = v6
    case 0:
        switch "hello" {
        case "\n\t":
            switch -292 {
            case -806:
                for fi7 := 0; fi7 < 15; fi7++ {
                    for fi8 := 0; fi8 < 7; fi8++ {
                        var v9 uint64 = 27
                        _ = fi8
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 10; fi10++ {
                        if (0 >= -490) {
                            v11 := []int{(-(fi7)), fi10}
                            _ = len(v11)
                            switch "foo bar" {
                            case "foo bar":
                                var v12 float32 = -2.25
                                _ = v12
                            case "":
                                v13 := GenS1{F0: ("foo bar" + "hello"), F1: -50, F2: 1}
                                _ = v13
                            }
                            _ = v11
                        } else {
                            var v14 int8 = -128
                            _ = v14
                        }
                        var v15 uint32 = 2147483648
                        _ = fi10
                        _ = v15
                    }
                    _ = fi7
                }
            }
        }
    }
    var v16 uint = 927
    if true {
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 int16 = 834
            v19 := []int{}
            var v20 int32 = -1
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        v16 = v16
        v21 := GenS1{F0: "", F1: 32767, F2: 4294967295}
        _ = v21
    }
    v22 := GenS1{F0: "hello", F1: 42, F2: 2147483647}
    fmt.Println(v16)
    _ = v16
    _ = v22
    return 42
}

func main() {
    var v1 uint64 = 0
    v2 := []uint32{1, 42, 1, 602}
    v3 := make(map[uint32]string)
    var v4 int8 = -1
    var v23 int = -1
    var v24 float64 = 0.0
    v23 += v23
    fmt.Println(v1, v24, v24)
    v1 += v1
    for fi25 := 0; fi25 < 1; fi25++ {
        switch 59 {
        case 127:
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 int = (^(fi25))
                v28 := GenS1{F0: "", F1: 670, F2: 519}
                if (4294967295 >= 0) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if true {
                        v24 -= (-2.25 * float64(float32(((v24 / 1e10) / v24))))
                        v29 := v3[0]
                        switch ((v29 + v29) + "hello") {
                        case "\n\t":
                            if (!(((!(v4 < v4)) && false) && (!true))) {
                                var v30 int64 = 0
                                v1 = 0
                                _ = v30
                            }
                        case "foo bar":
                            fmt.Println(v24)
                        case "":
                            switch (v29 + v29) {
                            case "hello":
                                var v31 uint16 = 416
                                _ = v31
                            case "foo bar":
                                v32 := GenS2{F0: 42, F1: ("hello" + "hello"), F2: (v1 >> ((18446744073709551615) & 63))}
                                _ = v32
                            default:
                                var v33 bool = ((-793 >= -1) && (("foo bar" > v29) || (-2147483648 != -2147483648)))
                                _ = v33
                            }
                        }
                        _ = v29
                    } else {
                        v27 += (^(v23))
                    }
                    v34 := GenS1{F0: (("hello" + "foo bar") + "foo bar"), F1: 1000, F2: 4294967295}
                    _ = v34
                } else {
                    if false {
                        var v35 int64 = 42
                        var v36 bool = true
                        _ = v35
                        _ = v36
                    } else {
                        v37 := GenS3{F0: 478, F1: -815, F2: 42}
                        v38 := GenS4{F0: v1, F1: 2147483647, F2: (^(v23))}
                        var v39 int8 = v4
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                }
                if false {
                    v3[2147483648] = (((("hello" + ("" + ("foo bar" + (("foo bar" + "") + "hello")))) + "hello") + "foo bar") + ((("hello" + ("\n\t" + "\n\t")) + "a") + ""))
                    v40 := GenS5{F0: v1, F1: 32767, F2: 65535}
                    _ = v40
                } else {
                    if (!true) {
                        v24 = (v24 - 0.1)
                        v24 -= v24
                        var v41 uint32 = 2147483648
                        _ = v41
                    }
                    v2 = append(v2, 611)
                    var v42 bool = false
                    _ = v42
                }
                _ = fi26
                _ = v27
                _ = v28
            }
        case 9223372036854775807:
            fmt.Println(v24, v24, v23)
        case -2147483648:
            v43 := GenS6{F0: uint(v1), F1: true}
            _ = v43
        }
        var v44 float32 = (0.0 / 0.1)
        _ = fi25
        _ = v44
    }
    v24 += (float64((0.1 / 1.0)) * ((1.5 * (v24 / v24)) * v24))
    if (!false) {
        for fi45 := 0; fi45 < 16; fi45++ {
            var v46 uint8 = 203
            fmt.Println(fi45, fi45)
            _ = fi45
            _ = v46
        }
    } else {
        if (65535 != 472) {
            v47 := GenS7{F0: (-(v4)), F1: (v1 ^ 4294967295)}
            _ = v47
        }
        delete(v3, 2147483648)
        v3[2147483647] = ""
    }
    fmt.Println(v4)
    v48 := GenS7{F0: 0, F1: (v1 - 4294967295)}
    switch "hello" {
    case "a":
        fmt.Println(v24, v24)
    case "\n\t":
        var v49 uint16 = 65535
        _ = v49
    default:
        if ((1 >= 0) && (42 >= v4)) {
            _ = len(v3)
            v23 += (v23 - 4294967295)
            v50 := GenS8{F0: -1, F1: -2147483648, F2: false}
            _ = v50
        } else {
            if (v23 < (^(v23))) {
                for fi51 := 0; fi51 < 19; fi51++ {
                    v1 = 42
                    if (v4 != int8(v23)) {
                        var v52 int16 = -1
                        for fi53 := 0; fi53 < 19; fi53++ {
                            var v54 uint8 = 128
                            _ = fi53
                            _ = v54
                        }
                        v23 = (v23 ^ (fi51 - (fi51 ^ (fi51 ^ 549))))
                        _ = v52
                    } else {
                        if true {
                            fmt.Println(fi51)
                            var v55 float32 = 3.14159
                            var v56 float64 = v24
                            _ = v55
                            _ = v56
                        } else {
                            var v57 int = fi51
                            for fi58 := 0; fi58 < 18; fi58++ {
                                v24 += (v24 / v24)
                                v23 += fi58
                                v59 := GenS9{F0: (("hello" + "a") + "foo bar")}
                                if false {
                                    v4 += (v4 / ((-128) | 1))
                                    switch "a" {
                                    case "":
                                        v57 = 32768
                                    case "hello":
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(v24)
                                        case "\n\t":
                                            var v60 uint8 = 1
                                            _ = v60
                                        case "hello":
                                            v61 := GenS7{F0: -1, F1: 427}
                                            _ = v61
                                        default:
                                            switch "" {
                                            case "a":
                                                _ = len(v3)
                                            case "hello":
                                                if (1e10 <= -2.25) {
                                                    fmt.Println(fi51, v23, v24)
                                                }
                                            case "\n\t":
                                                v62 := GenS9{F0: ("" + "")}
                                                _ = v62
                                            }
                                        }
                                    }
                                }
                                v63 := GenS3{F0: 0, F1: 0, F2: uint(v1)}
                                _ = fi58
                                _ = v59
                                _ = v63
                            }
                            for fi64 := 0; fi64 < 14; fi64++ {
                                v65 := GenS10{F0: (v1 / ((42) | 1))}
                                _ = fi64
                                _ = v65
                            }
                            _ = v57
                        }
                    }
                    _ = fi51
                }
                v4 -= v4
            } else {
                for fi66 := 0; fi66 < 20; fi66++ {
                    var v67 int64 = 42
                    if true {
                        var v68 uint16 = 65535
                        _ = v68
                    } else {
                        for fi69 := 0; fi69 < 18; fi69++ {
                            if false {
                                fmt.Println(fi66, fi66, v4)
                            } else {
                                var v70 string = ""
                                _ = v70
                            }
                            if true {
                                fmt.Println(v67)
                                if true {
                                    v71 := v3[434]
                                    v24 -= (-2.25 - (float64(1e10) * v24))
                                    _ = v71
                                }
                                var v72 int8 = (v4 | 0)
                                _ = v72
                            }
                            _ = fi69
                        }
                        var v73 float64 = (v24 * (v24 * 1.5))
                        v73 = float64(-2.25)
                        _ = v73
                    }
                    if ((-1 <= v67) && false) {
                        v23 = fi66
                        for fi74 := 0; fi74 < 12; fi74++ {
                            switch "" {
                            case "hello":
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = 42
                                    }
                                } else {
                                    fmt.Println(v4, v67, v24)
                                    if len(v2) > 0 {
                                        v2[0] = 42
                                    }
                                    v4 -= int8(v1)
                                    if len(v2) > 0 {
                                        v2[0] = 65535
                                    }
                                }
                            case "\n\t":
                                switch -1 {
                                case 77:
                                    v67 += 9223372036854775807
                                case -710:
                                    fmt.Println(fi74)
                                case 128:
                                    switch (-(fi66)) {
                                    case -9223372036854775808:
                                        v3[362] = (("\n\t" + "a") + "a")
                                        if true {
                                            if (true || (true || false)) {
                                                if (false && (!((float32(float64((-2.25 / float32((float64(1.5) + v24))))) / 1.5) < 3.14159))) {
                                                    if len(v2) > 0 {
                                                        v2[0] = 2147483647
                                                    }
                                                    fmt.Println(v67, v24)
                                                } else {
                                                    delete(v3, 0)
                                                    v23 = fi66
                                                    v1 += v1
                                                }
                                            } else {
                                                _ = len(v3)
                                                var v75 int32 = 0
                                                if (v75 >= (^(v75))) {
                                                    var v76 float64 = -2.25
                                                    _ = v76
                                                }
                                                _ = v75
                                            }
                                        }
                                        v4 = (-(v4))
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v77 int64 = 9223372036854775807
                                        _ = v77
                                    }
                                default:
                                    if (v1 != (v1 | 0)) {
                                        if false {
                                            v24 -= v24
                                            v78 := GenS11{F0: 0}
                                            _ = v78
                                        }
                                    }
                                }
                            default:
                                v67 = (v67 % (((v67 & v67)) | 1))
                            }
                            _ = fi74
                        }
                        v79 := GenS3{F0: 0, F1: int16(v1), F2: 255}
                        _ = v79
                    } else {
                        v1 = (v1 + v1)
                    }
                    _ = fi66
                    _ = v67
                }
            }
            switch ("\n\t" + ("hello" + "foo bar")) {
            case "\n\t":
                if false {
                    if false {
                        if (3.14159 >= (0.0 - (-2.25 / 1.5))) {
                            var v80 int64 = 191
                            _ = v80
                        } else {
                            var v81 int8 = -62
                            if true {
                                v81 = -1
                                if false {
                                    fmt.Println(v81, v1, v4)
                                    var v82 int = (v23 % ((128) | 1))
                                    _ = v82
                                } else {
                                    v4 = v4
                                }
                            } else {
                                if false {
                                    for fi83 := 0; fi83 < 1; fi83++ {
                                        switch ("hello" + ("" + "")) {
                                        case "a":
                                            for fi84 := 0; fi84 < 9; fi84++ {
                                                fmt.Println(fi84, v1, fi84)
                                                v24 += (0.1 * 0.0)
                                                if (!false) {
                                                    v4 -= 100
                                                    switch "a" {
                                                    case "\n\t":
                                                        if true {
                                                            switch (fi84 & v23) {
                                                            case 128:
                                                                v85 := GenS9{F0: "hello"}
                                                                _ = v85
                                                            default:
                                                                var v86 uint8 = 42
                                                                _ = v86
                                                            }
                                                            v2 = append(v2, 1)
                                                        }
                                                    }
                                                    v1 -= 0
                                                    v4 -= 0
                                                    if true {
                                                        if false {
                                                            switch (^(v23)) {
                                                            case 4294967295:
                                                                var v87 float32 = 3.14159
                                                                _ = v87
                                                            case 2147483647:
                                                                var v88 int16 = 0
                                                                _ = v88
                                                            }
                                                            switch "hello" {
                                                            case "":
                                                                if false {
                                                                    var v89 uint16 = 1
                                                                    _ = v89
                                                                } else {
                                                                    if false {
                                                                        if (false || false) {
                                                                            v23 = fi84
                                                                            if false {
                                                                                v81 -= -112
                                                                                v90 := GenS12{F0: 32768, F1: true}
                                                                                v1 = 667
                                                                                _ = v90
                                                                            } else {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = 42
                                                                                }
                                                                            }
                                                                            switch ("foo bar" + ("foo bar" + ("a" + ("" + "")))) {
                                                                            case "foo bar":
                                                                                fmt.Println(v23)
                                                                            case "":
                                                                                fmt.Println(v1)
                                                                            case "a":
                                                                                v23 = -682
                                                                            }
                                                                        } else {
                                                                            if (false && (v24 > v24)) {
                                                                                var v91 uint = 0
                                                                                _ = v91
                                                                            }
                                                                        }
                                                                        if (635 != -1) {
                                                                            v81 = -1
                                                                            v24 -= v24
                                                                        } else {
                                                                            if false {
                                                                                v92 := GenS9{F0: "\n\t"}
                                                                                var v93 int8 = v81
                                                                                var v94 int32 = -217
                                                                                _ = v92
                                                                                _ = v93
                                                                                _ = v94
                                                                            } else {
                                                                                v81 += v4
                                                                            }
                                                                            switch "a" {
                                                                            case "hello":
                                                                                var v95 uint32 = 2147483647
                                                                                _ = v95
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            case "hello":
                                                                if false {
                                                                    _ = len(v2)
                                                                    if (9223372036854775807 > -9223372036854775808) {
                                                                        v24 -= (v24 + (float64((((1e10 * -2.25) * 3.14159) + 3.14159)) - 3.14159))
                                                                    } else {
                                                                        switch 42 {
                                                                        case 32768:
                                                                            v4 = (v81 / (((v81 & v81)) | 1))
                                                                        default:
                                                                            v4 = -51
                                                                            fmt.Println(v24)
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                switch (v23 >> ((609) & 63)) {
                                                                                case 1:
                                                                                    fmt.Println(v81)
                                                                                case -1:
                                                                                    var v96 float64 = (v24 * v24)
                                                                                    _ = v96
                                                                                case 32768:
                                                                                    var v97 uint16 = genFn0()
                                                                                    _ = v97
                                                                                }
                                                                            }
                                                                        }
                                                                        v98 := GenS6{F0: 4294967295, F1: (!((^(v1)) < v1))}
                                                                        var v99 int8 = v81
                                                                        _ = v98
                                                                        _ = v99
                                                                    }
                                                                    v1 -= (^(v1))
                                                                }
                                                                v4 = (^(v81))
                                                                var v100 float32 = 0.0
                                                                _ = v100
                                                            case "\n\t":
                                                                if false {
                                                                    if false {
                                                                        fmt.Println(v1, v24, fi84)
                                                                        var v101 int64 = -777
                                                                        v23 = fi83
                                                                        _ = v101
                                                                    }
                                                                } else {
                                                                    v4 = (-(v4))
                                                                    if false {
                                                                        switch ("" + ("foo bar" + ("\n\t" + "a"))) {
                                                                        case "foo bar":
                                                                            var v102 uint8 = 255
                                                                            _ = v102
                                                                        case "":
                                                                            var v103 bool = false
                                                                            _ = v103
                                                                        case "hello":
                                                                            fmt.Println(v4)
                                                                        default:
                                                                            fmt.Println(v24, v1)
                                                                        }
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        var v104 bool = false
                                                                        _ = v104
                                                                    }
                                                                    v1 = (v1 ^ (^(v1)))
                                                                }
                                                            default:
                                                                fmt.Println(v81, v4, v81)
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    if (-1 == -468) {
                                                        v1 += v1
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    v1 -= v1
                                                }
                                                _ = fi84
                                            }
                                        case "\n\t":
                                            var v105 uint = 4294967295
                                            _ = v105
                                        }
                                        _ = fi83
                                    }
                                }
                                var v106 int64 = 42
                                var v107 int32 = 1
                                _ = v106
                                _ = v107
                            }
                            if false {
                                v24 = 1e10
                                fmt.Println(v1, v23)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                if false {
                                    fmt.Println(v81)
                                } else {
                                    fmt.Println(v81, v24)
                                }
                                if false {
                                    for fi108 := 0; fi108 < 20; fi108++ {
                                        for fi109 := 0; fi109 < 9; fi109++ {
                                            fmt.Println(v1, v24, fi109)
                                            switch (fi109 - v23) {
                                            case -607:
                                                switch (fi109 & fi108) {
                                                case 4294967295:
                                                    v24 = v24
                                                default:
                                                    v81 += (v4 | v81)
                                                }
                                            case 642:
                                                fmt.Println(fi108, v81, v23)
                                            case 128:
                                                var v110 int8 = -41
                                                _ = v110
                                            default:
                                                v81 -= v4
                                            }
                                            var v111 uint8 = 7
                                            fmt.Println(v81)
                                            _ = len(v3)
                                            _ = fi109
                                            _ = v111
                                        }
                                        _ = fi108
                                    }
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = 42
                                    }
                                    v2 = append(v2, 42)
                                    fmt.Println(v4, v4)
                                }
                                v112 := GenS13{F0: 0}
                                _ = v112
                            }
                            _ = v81
                        }
                        v2 = append(v2, 1)
                    } else {
                        for fi113 := 0; fi113 < 12; fi113++ {
                            var v114 uint32 = 4294967295
                            v115 := GenS14{F0: false, F1: 1}
                            v1 = 4294967295
                            _ = fi113
                            _ = v114
                            _ = v115
                        }
                        v116 := v3[uint32(v1)]
                        if true {
                            fmt.Println(v23, v116)
                        }
                        _ = v116
                    }
                }
            default:
                fmt.Println(v4)
            }
            var v117 int8 = (v4 | (-(v4)))
            _ = v117
        }
    }
    var v118 string = ("hello" + "")
    var v119 int32 = -339
    var v120 int16 = 0
    var v121 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v118
    _ = v119
    _ = v120
    _ = v121
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v48
}
