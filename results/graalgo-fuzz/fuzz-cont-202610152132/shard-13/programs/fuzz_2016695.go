package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 int; F2 int64 }
type GenS7 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint; F1 uint32; F2 int8 }
func genFn0(p0 uint, p1 int16, p2 uint) int8 {
    p1 = p1
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := []int{(-(fi4)), (^(fi4)), -1, fi4}
        var v6 string = "hello"
        var v7 string = "foo bar"
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return -128
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int16 = 1
        var v10 uint64 = 960
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int32 = -1
    v11 = (-(v11))
    var v12 float32 = 1.5
    v13 := GenS1{F0: -2147483648}
    v14 := []int{1, 241}
    for fi15 := 0; fi15 < 9; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS1{F0: fi16}
            switch ("hello" + "hello") {
            case "hello":
                v18 := GenS1{F0: fi16}
                _ = v18
            case "\n\t":
                _ = len(v14)
            default:
                fmt.Println(v11)
            }
            v12 -= 1.5
            _ = fi16
            _ = v17
        }
        if true {
            var v19 float32 = float32((-2.25 * (0.1 - 1.5)))
            switch "a" {
            case "":
                switch ((("foo bar" + ("hello" + "a")) + "a") + "") {
                case "hello":
                    v20 := GenS2{F0: 1000, F1: 42}
                    _ = v20
                case "":
                    var v21 uint32 = 4294967295
                    _ = v21
                default:
                    for fi22 := 0; fi22 < 14; fi22++ {
                        if false {
                            if true {
                                v19 += (v12 - (0.0 / 1e10))
                                var v23 uint32 = 2147483647
                                var v24 uint16 = 0
                                var v25 uint = 2
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS3{F0: 0, F1: 65535}
                            _ = len(v3)
                            _ = v26
                        } else {
                            if len(v14) > 0 {
                                v14[0] = fi22
                            }
                            if ((true && (fi15 < fi22)) && false) {
                                if true {
                                    v11 = v11
                                    switch "hello" {
                                    case "":
                                        switch "\n\t" {
                                        case "a":
                                            var v27 uint64 = 18446744073709551615
                                            _ = v27
                                        case "":
                                            v12 = v12
                                        default:
                                            var v28 string = ""
                                            _ = v28
                                        }
                                    case "a":
                                        var v29 int64 = 42
                                        v29 -= (v29 % ((v29) | 1))
                                        v19 = 3.14159
                                        _ = v29
                                    default:
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    }
                                    fmt.Println(v11, fi22, fi22)
                                    switch "a" {
                                    case "\n\t":
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    case "a":
                                        fmt.Println(v12)
                                    default:
                                        if (!false) {
                                            var v30 int64 = -318
                                            if true {
                                                fmt.Println(fi22)
                                                v12 += -2.25
                                            } else {
                                                switch "hello" {
                                                case "a":
                                                    v31 := GenS3{F0: 9223372036854775807, F1: 490}
                                                    _ = v31
                                                case "\n\t":
                                                    if false {
                                                        v30 -= (v30 * -1)
                                                        var v32 bool = (502 > 305)
                                                        v11 = (v11 % ((v11) | 1))
                                                        _ = v32
                                                    }
                                                default:
                                                    if false {
                                                        if false {
                                                            _ = len(v14)
                                                            _ = len(v14)
                                                        } else {
                                                            v33 := GenS4{F0: 1, F1: 0}
                                                            v34 := GenS2{F0: -32768, F1: 0}
                                                            var v35 float32 = float32(-2.25)
                                                            _ = v33
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                        v19 = v12
                                                        delete(v2, ("\n\t" + "hello"))
                                                    }
                                                }
                                            }
                                            _ = v30
                                        }
                                    }
                                    switch ("" + "") {
                                    case "foo bar":
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    case "":
                                        var v36 uint = 1
                                        _ = v36
                                    }
                                    v12 = 0.0
                                }
                            } else {
                                switch fi22 {
                                case 1:
                                    v19 -= float32((3.14159 + float64(v19)))
                                    var v37 string = "foo bar"
                                    _ = v37
                                case -65:
                                    switch (-(fi22)) {
                                    case 1:
                                        if ("a" > (("a" + (("hello" + "foo bar") + "\n\t")) + "foo bar")) {
                                            v12 -= 3.14159
                                        } else {
                                            if (!false) {
                                                var v38 uint8 = 3
                                                delete(v3, 559)
                                                fmt.Println(v19, fi22, v12)
                                                _ = v38
                                            } else {
                                                var v39 float64 = -2.25
                                                _ = v39
                                            }
                                            var v40 bool = ((v11 | (^(v11))) != v11)
                                            if (!true) {
                                                switch 42 {
                                                case 973:
                                                    if (580 > 255) {
                                                        switch ("\n\t" + "\n\t") {
                                                        case "a":
                                                            var v41 string = "hello"
                                                            _ = v41
                                                        default:
                                                            var v42 uint16 = 0
                                                            _ = v42
                                                        }
                                                        var v43 int16 = -863
                                                        _ = v43
                                                    } else {
                                                        v12 = 3.14159
                                                    }
                                                case 100:
                                                    var v44 uint8 = 255
                                                    _ = v44
                                                case 9223372036854775807:
                                                    v45 := v1[2147483647]
                                                    _ = v45
                                                default:
                                                    var v46 uint = 42
                                                    _ = v46
                                                }
                                                var v47 bool = (!v40)
                                                _ = v47
                                            } else {
                                                if true {
                                                    var v48 float64 = 0.0
                                                    v49 := GenS2{F0: 1000, F1: 649}
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    v11 -= v11
                                                    var v50 int32 = (^(v11))
                                                    fmt.Println(fi22, fi15, v12)
                                                    _ = v50
                                                }
                                            }
                                            _ = v40
                                        }
                                        v51 := GenS1{F0: fi22}
                                        var v52 uint32 = 42
                                        if (true && true) {
                                            v11 -= (v11 | (v11 << ((18446744073709551615) & 63)))
                                            v2[((("a" + "a") + "\n\t") + "hello")] = (-(fi15))
                                            switch (fi15 & fi22) {
                                            case 136:
                                                v53 := GenS3{F0: -1, F1: 18446744073709551615}
                                                _ = v53
                                            default:
                                                if (!(int64(v11) != -9223372036854775808)) {
                                                    var v54 float32 = v12
                                                    _ = v54
                                                }
                                            }
                                        } else {
                                            v55 := GenS4{F0: 32767, F1: fi22}
                                            var v56 bool = true
                                            fmt.Println(v12)
                                            var v57 string = "foo bar"
                                            _ = v55
                                            _ = v56
                                            _ = v57
                                        }
                                        var v58 uint32 = (v52 - v52)
                                        _ = v51
                                        _ = v52
                                        _ = v58
                                    case 136:
                                        _ = len(v2)
                                    case 370:
                                        v12 = v19
                                    default:
                                        v59 := v3[fi22]
                                        _ = v59
                                    }
                                case 32768:
                                    switch (fi22 + fi22) {
                                    case 42:
                                        var v60 int32 = -1
                                        _ = v60
                                    case 1:
                                        if (966 < 2) {
                                            fmt.Println(v11, fi15)
                                            fmt.Println(v19, v19)
                                        }
                                        v61 := GenS5{F0: (fi22 << ((65535) & 63))}
                                        fmt.Println(fi15)
                                        _ = v61
                                    default:
                                        if (407 >= 2147483647) {
                                            switch ("hello" + ("foo bar" + ("hello" + "foo bar"))) {
                                            case "":
                                                var v62 string = "hello"
                                                _ = v62
                                            case "hello":
                                                v19 = 1e10
                                            case "\n\t":
                                                if (true || false) {
                                                    var v63 uint64 = 1
                                                    _ = v63
                                                }
                                            }
                                        } else {
                                            var v64 uint = 18446744073709551615
                                            var v65 int8 = 1
                                            _ = v64
                                            _ = v65
                                        }
                                    }
                                default:
                                    if true {
                                        v19 += 0.1
                                        v66 := GenS6{F0: 2, F1: fi22, F2: 9223372036854775807}
                                        fmt.Println(fi22, fi15, v12)
                                        _ = v66
                                    }
                                }
                                if (!false) {
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                } else {
                                    if false {
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    }
                                }
                                v67 := GenS5{F0: (fi15 - (-(fi15)))}
                                _ = v67
                            }
                            _ = len(v14)
                            v14 = append(v14, fi22)
                        }
                        switch ("a" + "a") {
                        case "foo bar":
                            v12 += (((v19 / float32((1e10 * 0.1))) / (((v19 / v12) - float32(float64(v12))) - 0.1)) + v19)
                        }
                        _ = fi22
                    }
                }
            case "foo bar":
                v12 = (float32(3.14159) * float32((0.0 / -2.25)))
            }
            _ = v19
        } else {
            for fi68 := 0; fi68 < 15; fi68++ {
                if false {
                    fmt.Println(v11, v12, fi15)
                    if len(v14) > 0 {
                        v14[0] = 32767
                    }
                    fmt.Println(v11)
                } else {
                    fmt.Println(v11, v11)
                }
                switch "hello" {
                case "":
                    if (-2.25 >= (-2.25 - v12)) {
                        var v69 int64 = -9223372036854775808
                        v12 = v12
                        _ = len(v3)
                        _ = v69
                    } else {
                        v11 += v11
                    }
                    fmt.Println(v11, fi68)
                case "\n\t":
                    var v70 uint64 = 58
                    switch 32768 {
                    case 32768:
                        v71 := GenS4{F0: -1, F1: fi68}
                        _ = v71
                    case -500:
                        var v72 float32 = 1e10
                        _ = v72
                    case 1:
                        var v73 uint8 = 127
                        _ = v73
                    }
                    var v74 int64 = -9223372036854775808
                    _ = v70
                    _ = v74
                }
                var v75 float32 = v12
                _ = fi68
                _ = v75
            }
        }
        for fi76 := 0; fi76 < 14; fi76++ {
            var v77 string = "a"
            _ = fi76
            _ = v77
        }
        _ = fi15
    }
    var v78 int8 = 124
    v78 = 0
    for fi79 := 0; fi79 < 10; fi79++ {
        if false {
            v11 = v11
            switch "a" {
            case "foo bar":
                v12 += ((v12 - float32(1e10)) + (float32(((3.14159 - 1e10) + (-2.25 + (1.5 - (-2.25 + 3.14159))))) * v12))
            default:
                v12 += v12
            }
        } else {
            fmt.Println(v12, v12, fi79)
        }
        v80 := GenS7{F0: 1, F1: 0, F2: 789}
        _ = fi79
        _ = v80
    }
    var v81 string = ("foo bar" + ("" + "a"))
    for fi82 := 0; fi82 < 11; fi82++ {
        v83 := GenS1{F0: 880}
        _ = fi82
        _ = v83
    }
    if true {
        switch 534 {
        case 42:
            _ = len(v14)
        case 728:
            v78 += -128
        case 32768:
            v81 = v81
        default:
            fmt.Println(v11)
        }
        var v84 uint32 = 700
        if true {
            var v85 int64 = -9223372036854775808
            fmt.Println(v84, v81, v78)
            var v86 uint16 = 65535
            _ = v85
            _ = v86
        }
        _ = v84
    } else {
        _ = len(v3)
        for fi87 := 0; fi87 < 1; fi87++ {
            if ((v78 + (v78 << ((65535) & 63))) < v78) {
                v11 = -101
                v88 := GenS1{F0: fi87}
                v11 = v11
                _ = v88
            }
            v14 = append(v14, fi87)
            _ = fi87
        }
        v12 += ((v12 / (v12 / v12)) * 0.0)
    }
    _ = len(v2)
    switch "foo bar" {
    case "hello":
        for fi89 := 0; fi89 < 6; fi89++ {
            v90 := GenS4{F0: 0, F1: (fi89 & fi89)}
            if false {
                fmt.Println(v81, v81, fi89)
                v12 = v12
                var v91 float32 = v12
                _ = v91
            }
            _ = fi89
            _ = v90
        }
    case "foo bar":
        switch 42 {
        case 139:
            v78 = v78
        case -113:
            v3[128] = 367
        default:
            if (!(!(!(true || (-9223372036854775808 == int64(v11)))))) {
                v81 = v81
            }
        }
    case "\n\t":
        v81 = v81
    default:
        if (42 <= 1) {
            fmt.Println(v78, v81, v81)
        }
    }
    switch 0 {
    case 32767:
        for fi92 := 0; fi92 < 9; fi92++ {
            v81 = v81
            _ = fi92
        }
    case 32768:
        var v93 int32 = (v11 | (v11 >> ((170) & 63)))
        _ = v93
    case 0:
        delete(v1, 2)
    default:
        v11 = (v11 * (v11 & (v11 + 42)))
    }
    if (float32((0.0 - ((0.0 / (((((1e10 + 0.0) * 1e10) + (0.0 / 1.0)) - 1e10) - 0.1)) - (float64((((v12 + 1e10) - 0.1) / (0.0 / v12))) - 0.0)))) > (v12 / (v12 - 1e10))) {
        v81 = v81
    } else {
        if (4294967295 < 729) {
            if (v11 != v11) {
                if (2 < 232) {
                    var v94 int64 = 0
                    var v95 int16 = 1
                    var v96 uint16 = 0
                    _ = v94
                    _ = v95
                    _ = v96
                }
                if true {
                    var v97 int32 = v11
                    if true {
                        for fi98 := 0; fi98 < 3; fi98++ {
                            _ = len(v14)
                            var v99 float64 = -2.25
                            v12 += 0.1
                            _ = fi98
                            _ = v99
                        }
                        for fi100 := 0; fi100 < 17; fi100++ {
                            switch v81 {
                            case "\n\t":
                                v101 := GenS5{F0: -631}
                                _ = v101
                            case "hello":
                                v102 := GenS1{F0: (fi100 % ((fi100) | 1))}
                                _ = v102
                            case "a":
                                _ = len(v3)
                            }
                            var v103 uint64 = 165
                            _ = fi100
                            _ = v103
                        }
                    } else {
                        if (int16(v78) >= -1) {
                            v104 := v3[127]
                            _ = v104
                        }
                    }
                    _ = v97
                }
                if false {
                    var v105 int = 32767
                    switch v105 {
                    case 127:
                        v81 = ""
                    default:
                        v105 = 127
                    }
                    _ = v105
                } else {
                    v81 = "a"
                }
            }
            for fi106 := 0; fi106 < 12; fi106++ {
                delete(v2, v81)
                var v107 int64 = -394
                if true {
                    for fi108 := 0; fi108 < 4; fi108++ {
                        if true {
                            v11 = v11
                        }
                        _ = fi108
                    }
                } else {
                    if (2147483648 != 220) {
                        v78 -= (v78 - v78)
                        v109 := GenS8{F0: v81}
                        _ = v109
                    } else {
                        v110 := GenS6{F0: 0, F1: (^(fi106)), F2: v107}
                        _ = v110
                    }
                    v111 := GenS4{F0: -586, F1: -163}
                    v2["foo bar"] = fi106
                    _ = v111
                }
                _ = fi106
                _ = v107
            }
        }
        v112 := GenS9{F0: 1, F1: 1, F2: (v78 << ((816) & 63))}
        _ = v112
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v78
    _ = v81
}
