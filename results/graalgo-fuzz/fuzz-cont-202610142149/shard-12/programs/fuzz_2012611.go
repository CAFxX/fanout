package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int8; F1 uint16; F2 string }
type GenS8 struct { F0 uint; F1 bool; F2 int32 }
type GenS9 struct { F0 int32; F1 int64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint16; F1 uint32; F2 uint }
type GenS12 struct { F0 int8 }
func genFn0(p0 int16, p1 uint, p2 uint32) uint64 {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{-607, 32768}
        if (!((!(2 > (p2 >> ((p1) & 63)))) || true)) {
            if (true && true) {
                var v8 int32 = -98
                for fi9 := 0; fi9 < 15; fi9++ {
                    v8 -= -538
                    _ = fi9
                }
                var v10 string = "hello"
                _ = v10
                _ = v8
            }
        } else {
            fmt.Println(p0, p0, p2)
            var v11 float32 = 1.5
            fmt.Println(fi6, fi6, v11)
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(p1, p1)
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 915
}
func genFn1() uint32 {
    var v12 int8 = 12
    v12 += (v12 * v12)
    fmt.Println(v12)
    v13 := GenS1{F0: 65535}
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 42
    var v3 int64 = 42
    var v4 int8 = 112
    v5 := make(map[int]string)
    var v14 int64 = (v3 + (v3 ^ v3))
    switch 4294967295 {
    case 163:
        v3 = 858
    default:
        var v15 uint64 = 867
        _ = v15
    }
    v3 = v3
    fmt.Println(v4, v3, v14)
    var v16 float32 = float32(float64(3.14159))
    v4 = -128
    v3 = (v3 << ((0) & 63))
    switch int(v3) {
    case 1:
        fmt.Println(v2, v16)
    case 40:
        if (-371 < 108) {
            var v17 int = 32767
            for fi18 := 0; fi18 < 1; fi18++ {
                if (695 == (v14 << ((255) & 63))) {
                    if (false && (!true)) {
                        if false {
                            if true {
                                var v19 uint16 = (v2 >> ((1) & 63))
                                for fi20 := 0; fi20 < 9; fi20++ {
                                    fmt.Println(fi18, v2)
                                    if true {
                                        fmt.Println(v16, fi18)
                                        switch "hello" {
                                        case "\n\t":
                                            var v21 uint16 = v2
                                            _ = v21
                                        case "hello":
                                            var v22 float32 = 1e10
                                            _ = v22
                                        case "":
                                            if true {
                                                _ = len(v1)
                                            } else {
                                                var v23 uint64 = 1
                                                v24 := v5[(fi20 | -2147483648)]
                                                if true {
                                                    v25 := GenS1{F0: v2}
                                                    switch 127 {
                                                    case 4294967295:
                                                        var v26 int8 = v4
                                                        _ = v26
                                                    case -676:
                                                        var v27 float32 = v16
                                                        _ = v27
                                                    default:
                                                        fmt.Println(fi18)
                                                    }
                                                    v19 -= (^(v2))
                                                    _ = v25
                                                }
                                                _ = v23
                                                _ = v24
                                            }
                                        default:
                                            v28 := v1[9223372036854775807]
                                            _ = v28
                                        }
                                    } else {
                                        switch "a" {
                                        case "":
                                            switch fi18 {
                                            case 1:
                                                fmt.Println(v3, v16, v17)
                                            case 659:
                                                var v29 uint = 2
                                                _ = v29
                                            }
                                        case "a":
                                            delete(v5, (fi20 % ((4294967295) | 1)))
                                        }
                                    }
                                    if false {
                                        switch "foo bar" {
                                        case "\n\t":
                                            v30 := GenS2{F0: (v4 / ((v4) | 1))}
                                            var v31 bool = (v16 > (1e10 * (-2.25 * 0.1)))
                                            _ = v30
                                            _ = v31
                                        case "foo bar":
                                            var v32 int8 = (v4 >> ((4294967295) & 63))
                                            _ = v32
                                        default:
                                            switch (("hello" + ("a" + "")) + "foo bar") {
                                            case "hello":
                                                var v33 int8 = 0
                                                _ = v33
                                            default:
                                                if true {
                                                    if ((true && true) || (2 <= v19)) {
                                                        v34 := []int{v17}
                                                        _ = v34
                                                    } else {
                                                        if true {
                                                            v3 = (v3 & v14)
                                                        }
                                                        var v35 uint32 = 389
                                                        switch v17 {
                                                        case 277:
                                                            var v36 float32 = 0.0
                                                            _ = v36
                                                        default:
                                                            fmt.Println(v4)
                                                        }
                                                        _ = v35
                                                    }
                                                    switch 881 {
                                                    case 0:
                                                        v3 = (v14 / (((v3 >> ((255) & 63))) | 1))
                                                    case -1:
                                                        if (v14 >= (v3 % ((v3) | 1))) {
                                                            if (true && (18446744073709551615 >= 449)) {
                                                                if false {
                                                                    v4 = v4
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(v19, v14)
                                                    }
                                                    if (225 < 188) {
                                                        var v37 int = v17
                                                        v37 = v37
                                                        _ = v37
                                                    }
                                                    var v38 uint16 = (v19 >> ((584) & 63))
                                                    _ = v38
                                                }
                                            }
                                            var v39 int16 = 42
                                            var v40 int = (fi20 + v17)
                                            _ = v39
                                            _ = v40
                                        }
                                        v41 := GenS1{F0: (v2 << ((917) & 63))}
                                        _ = v41
                                    } else {
                                        var v42 uint8 = 255
                                        v43 := []int{}
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = fi20
                                }
                                v1[9223372036854775807] = "hello"
                                _ = v19
                            } else {
                                if false {
                                    for fi44 := 0; fi44 < 19; fi44++ {
                                        if (!(!(true || (775 != 3)))) {
                                            var v45 uint = 813
                                            var v46 uint64 = 0
                                            _ = v45
                                            _ = v46
                                        } else {
                                            switch fi44 {
                                            case -2147483648:
                                                v47 := GenS3{F0: 2147483647, F1: v17}
                                                _ = v47
                                            case -1:
                                                if (1e10 >= (v16 * v16)) {
                                                    if (("foo bar" < ("hello" + ("hello" + ("foo bar" + ("foo bar" + "a"))))) || false) {
                                                        v48 := GenS4{F0: genFn0(-1, 42, 543), F1: (v4 << ((307) & 63))}
                                                        _ = v48
                                                    }
                                                } else {
                                                    if false {
                                                        if (!true) {
                                                            if (2147483647 < 0) {
                                                                v49 := GenS5{F0: v14}
                                                                if (!(true || false)) {
                                                                    var v50 string = ("hello" + "foo bar")
                                                                    _ = v50
                                                                }
                                                                v51 := []int{}
                                                                _ = v49
                                                                _ = v51
                                                            } else {
                                                                if (("a" + "\n\t") >= "hello") {
                                                                    fmt.Println(v3, fi18)
                                                                    if true {
                                                                        fmt.Println(v17, v4, fi18)
                                                                        var v52 int8 = -128
                                                                        var v53 int32 = -1
                                                                        _ = v52
                                                                        _ = v53
                                                                    } else {
                                                                        v1[(fi18 | fi18)] = "a"
                                                                    }
                                                                } else {
                                                                    if (1.5 >= (-2.25 + (1.5 * (0.1 * -2.25)))) {
                                                                        v54 := []int{128}
                                                                        var v55 uint32 = 2147483648
                                                                        _ = v54
                                                                        _ = v55
                                                                    }
                                                                }
                                                                if ((fi44 | fi44) != 1) {
                                                                    v2 -= v2
                                                                }
                                                            }
                                                        }
                                                        var v56 uint16 = v2
                                                        _ = v56
                                                    } else {
                                                        v4 = (v4 % (((v4 % ((1) | 1))) | 1))
                                                    }
                                                    v57 := []int{(fi44 >> ((386) & 63))}
                                                    _ = v57
                                                }
                                            default:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v58 bool = true
                                                    _ = v58
                                                }
                                            }
                                        }
                                        v59 := v5[(^(fi18))]
                                        if (true && false) {
                                            v4 += v4
                                            if false {
                                                var v60 uint64 = 0
                                                v61 := GenS6{F0: 2, F1: uint32(v60)}
                                                if (!(true || false)) {
                                                    var v62 int32 = -1
                                                    v63 := GenS7{F0: v4, F1: (v2 / (((v2 - (v2 ^ 551))) | 1)), F2: "hello"}
                                                    v16 = (float32(1.5) / v16)
                                                    _ = v62
                                                    _ = v63
                                                }
                                                _ = v60
                                                _ = v61
                                            } else {
                                                var v64 int8 = 1
                                                _ = v64
                                            }
                                        } else {
                                            v65 := GenS6{F0: 18446744073709551615, F1: 4294967295}
                                            v66 := v1[fi44]
                                            var v67 uint8 = 73
                                            fmt.Println(v59, v16, v2)
                                            _ = v65
                                            _ = v66
                                            _ = v67
                                        }
                                        _ = fi44
                                        _ = v59
                                    }
                                    fmt.Println(v17, v14, v2)
                                } else {
                                    var v68 bool = true
                                    _ = v68
                                }
                                v5[int(v4)] = "\n\t"
                                for fi69 := 0; fi69 < 16; fi69++ {
                                    switch (("hello" + "\n\t") + "") {
                                    case "hello":
                                        var v70 int16 = 42
                                        _ = v70
                                    case "foo bar":
                                        v71 := GenS3{F0: 1, F1: (fi69 & (v17 % ((fi69) | 1)))}
                                        _ = v71
                                    case "a":
                                        if false {
                                            fmt.Println(v3, fi18)
                                            if (!((-2.25 / -2.25) >= 3.14159)) {
                                                fmt.Println(fi18, fi18, v2)
                                            } else {
                                                var v72 int64 = (v14 ^ v14)
                                                v1[32768] = "a"
                                                v73 := []int{(-(fi18)), (v17 + (v17 * 9223372036854775807)), (-(fi18)), v17}
                                                var v74 int32 = -269
                                                _ = v72
                                                _ = v73
                                                _ = v74
                                            }
                                            v75 := GenS8{F0: uint(v2), F1: false, F2: 42}
                                            if (42 != 2147483648) {
                                                fmt.Println(v2)
                                            }
                                            v3 -= (v14 * (^(v14)))
                                            _ = v75
                                        } else {
                                            v2 -= 310
                                            fmt.Println(v4, fi18)
                                            fmt.Println(v3, v16, v4)
                                            delete(v5, v17)
                                        }
                                    default:
                                        v4 -= (v4 | v4)
                                    }
                                    switch "" {
                                    case "a":
                                        v76 := GenS6{F0: 1, F1: 4294967295}
                                        _ = v76
                                    }
                                    var v77 uint32 = 65535
                                    _ = fi69
                                    _ = v77
                                }
                            }
                            v17 += (v17 | 992)
                        } else {
                            fmt.Println(v17)
                            switch 42 {
                            case 1:
                                if true {
                                    for fi78 := 0; fi78 < 15; fi78++ {
                                        v14 = (v3 + 0)
                                        v4 = v4
                                        v3 = v14
                                        var v79 int16 = 1
                                        _ = fi78
                                        _ = v79
                                    }
                                    v80 := []int{(v17 ^ v17), (v17 * (fi18 - v17))}
                                    var v81 float32 = float32(0.0)
                                    _ = v80
                                    _ = v81
                                } else {
                                    v5[4294967295] = ("hello" + ("foo bar" + "hello"))
                                }
                            case 128:
                                for fi82 := 0; fi82 < 4; fi82++ {
                                    var v83 bool = (!false)
                                    _ = fi82
                                    _ = v83
                                }
                            }
                        }
                    }
                    v84 := GenS6{F0: 42, F1: 4294967295}
                    _ = v84
                }
                switch v17 {
                case 32767:
                    v2 = v2
                case 32768:
                    switch fi18 {
                    case 42:
                        if (2147483647 >= 337) {
                            v5[550] = "foo bar"
                            v17 = v17
                            v14 -= 1
                        } else {
                            if ((!false) && false) {
                                v14 = v14
                                delete(v5, fi18)
                                if (false || (true && true)) {
                                    for fi85 := 0; fi85 < 5; fi85++ {
                                        if (0 >= 2) {
                                            if true {
                                                v16 = v16
                                            }
                                            v86 := v5[v17]
                                            _ = v86
                                        } else {
                                            v87 := v1[v17]
                                            if false {
                                                v5[fi18] = v87
                                            } else {
                                                switch "" {
                                                case "hello":
                                                    switch v87 {
                                                    case "":
                                                        if false {
                                                            if true {
                                                                v88 := GenS9{F0: 42, F1: 1}
                                                                v89 := GenS10{F0: v3}
                                                                v90 := []int{fi85, (-(v17))}
                                                                _ = v88
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                        }
                                                    }
                                                }
                                                switch (^(fi18)) {
                                                case -1:
                                                    var v91 float32 = 0.0
                                                    _ = v91
                                                case 102:
                                                    var v92 uint32 = 1
                                                    _ = v92
                                                default:
                                                    v2 = (v2 * v2)
                                                    v4 -= int8(v2)
                                                    v2 = (^(v2))
                                                }
                                                var v93 uint32 = 192
                                                _ = v93
                                            }
                                            _ = v87
                                        }
                                        switch v17 {
                                        case 2147483647:
                                            v16 = 1e10
                                        case -288:
                                            if ("\n\t" > "") {
                                                v94 := GenS7{F0: (v4 % ((-128) | 1)), F1: 2, F2: (("" + ("foo bar" + ("foo bar" + "\n\t"))) + ("\n\t" + ("hello" + ("" + "foo bar"))))}
                                                _ = v94
                                            }
                                        default:
                                            v4 = v4
                                        }
                                        var v95 string = "hello"
                                        _ = fi85
                                        _ = v95
                                    }
                                    v1[fi18] = ""
                                    var v96 uint16 = v2
                                    _ = v96
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        for fi97 := 0; fi97 < 7; fi97++ {
                                            delete(v1, 128)
                                            _ = fi97
                                        }
                                    case "foo bar":
                                        var v98 bool = (!false)
                                        _ = v98
                                    }
                                    v4 = v4
                                }
                            }
                            v99 := GenS2{F0: v4}
                            _ = v99
                        }
                    case 678:
                        fmt.Println(v17)
                    }
                case 499:
                    v3 = (v14 & v3)
                default:
                    var v100 int64 = (v14 & (v3 % ((v14) | 1)))
                    _ = v100
                }
                _ = fi18
            }
            var v101 int64 = 9223372036854775807
            _ = v101
            _ = v17
        }
    }
    v102 := GenS11{F0: 1, F1: 0, F2: 336}
    var v103 uint32 = genFn1()
    for fi104 := 0; fi104 < 6; fi104++ {
        var v105 string = (("\n\t" + "\n\t") + "a")
        if ((v14 >> ((846) & 63)) <= v14) {
            v14 = (v3 ^ v14)
            for fi106 := 0; fi106 < 20; fi106++ {
                fmt.Println(v16, v105)
                _ = fi106
            }
        }
        _ = fi104
        _ = v105
    }
    fmt.Println(v4)
    fmt.Println(v14, v14)
    var v107 int = -671
    for fi108 := 0; fi108 < 4; fi108++ {
        v109 := []int{(fi108 + (v107 - (v107 + -750))), v107, fi108}
        v109 = append(v109, (v107 | int(v4)))
        delete(v5, 42)
        _ = fi108
        _ = v109
    }
    for fi110 := 0; fi110 < 1; fi110++ {
        if ((!true) || (789 > genFn0(42, 660, 2147483647))) {
            v111 := []int{}
            var v112 int32 = -1
            _ = v111
            _ = v112
        } else {
            switch ("foo bar" + "a") {
            case "foo bar":
                switch int(v3) {
                case -360:
                    v2 = v2
                case -453:
                    var v113 uint8 = 128
                    _ = v113
                case 0:
                    var v114 int16 = 42
                    _ = v114
                default:
                    if (444 > 1000) {
                        var v115 uint64 = 1
                        v107 -= v107
                        var v116 float64 = (-2.25 / 0.1)
                        _ = v115
                        _ = v116
                    }
                }
            case "a":
                v14 += -9223372036854775808
            case "hello":
                switch fi110 {
                case 682:
                    v117 := GenS12{F0: v4}
                    _ = v117
                case -397:
                    for fi118 := 0; fi118 < 14; fi118++ {
                        var v119 bool = (128 == 2)
                        var v120 int32 = -2147483648
                        v119 = (127 >= -574)
                        _ = fi118
                        _ = v119
                        _ = v120
                    }
                }
            default:
                var v121 int32 = 1
                _ = v121
            }
        }
        fmt.Println(v14, v103)
        fmt.Println(v2, v107, v4)
        _ = fi110
    }
    var v122 float32 = (-2.25 - ((v16 - (v16 * v16)) - 3.14159))
    var v123 string = ("" + "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("v107:", v107)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v102
    _ = v103
    _ = v107
    _ = v122
    _ = v123
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
