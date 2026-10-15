package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 int; F1 uint16; F2 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int16; F1 int; F2 int16 }
func genFn0(p0 int, p1 string, p2 int) uint16 {
    switch -2147483648 {
    case 1:
        v6 := GenS1{F0: 1, F1: 2147483647}
        _ = v6
    case 42:
        var v7 uint16 = 65535
        _ = v7
    default:
        fmt.Println(p2)
    }
    var v8 int64 = -9223372036854775808
    var v9 uint32 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int, p1 uint, p2 int16) uint8 {
    p1 = 2
    p0 = 42
    fmt.Println(p2, p1)
    v10 := GenS2{F0: 0}
    v11 := []int{4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 2
}

func main() {
    var v1 bool = false
    var v2 float64 = 1e10
    var v3 float32 = ((3.14159 / 0.1) + (float32(v2) / 1.5))
    var v4 uint32 = 436
    var v5 float32 = (((1e10 * v3) / 0.1) + (float32(((v2 * float64(v3)) / 0.1)) / v3))
    if v1 {
        var v12 string = "\n\t"
        switch 0 {
        case 9223372036854775807:
            var v13 int64 = 9223372036854775807
            _ = v13
        case 4294967295:
            v2 += ((v2 * v2) * ((v2 / v2) * float64(v3)))
        case -1:
            var v14 int = 32767
            _ = v14
        default:
            v15 := GenS1{F0: 42, F1: 42}
            _ = v15
        }
        fmt.Println(v12, v12)
        _ = v12
    } else {
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := []int{(fi16 / (((fi16 << ((164) & 63))) | 1)), (fi16 | -9223372036854775808), (fi16 % ((fi16) | 1)), (-(fi16))}
            _ = fi16
            _ = v17
        }
        if (315 > 18446744073709551615) {
            var v18 uint64 = 0
            _ = v18
        } else {
            v3 -= float32(v2)
            if (-32768 < int16(v4)) {
                var v19 uint8 = 255
                for fi20 := 0; fi20 < 16; fi20++ {
                    switch "hello" {
                    case "a":
                        var v21 uint16 = 42
                        _ = v21
                    case "hello":
                        var v22 int64 = 42
                        _ = v22
                    case "":
                        v23 := GenS1{F0: -2147483648, F1: int32(fi20)}
                        _ = v23
                    default:
                        if v1 {
                            v3 = v5
                            fmt.Println(v1, v4)
                        } else {
                            fmt.Println(v3)
                            for fi24 := 0; fi24 < 10; fi24++ {
                                v5 += ((-2.25 - 1.5) - float32((v2 / 0.1)))
                                _ = fi24
                            }
                        }
                    }
                    if (!v1) {
                        var v25 float64 = ((0.1 / v2) / float64((3.14159 / 3.14159)))
                        v3 += (float32(float64(-2.25)) * ((1e10 - float32((float64(float32(v2)) / v2))) * v5))
                        v26 := []int{(fi20 * fi20), (-(fi20)), fi20, fi20}
                        _ = v25
                        _ = v26
                    }
                    v3 -= v3
                    _ = fi20
                }
                _ = v19
            } else {
                switch (("a" + ("a" + "")) + "") {
                case "\n\t":
                    v5 += 1.5
                default:
                    v4 = v4
                }
                var v27 uint = 255
                switch 128 {
                case -2147483648:
                    if (1 != 127) {
                        v28 := GenS2{F0: 42}
                        if (true && (!(v1 && ((!true) || true)))) {
                            if v1 {
                                for fi29 := 0; fi29 < 2; fi29++ {
                                    v27 -= (^(v27))
                                    for fi30 := 0; fi30 < 5; fi30++ {
                                        v27 = (^(v27))
                                        v31 := GenS1{F0: -2147483648, F1: -927}
                                        v32 := GenS2{F0: 2}
                                        _ = fi30
                                        _ = v31
                                        _ = v32
                                    }
                                    _ = fi29
                                }
                                var v33 uint8 = 2
                                _ = v33
                            } else {
                                v34 := []int{}
                                var v35 uint16 = 710
                                if (v27 >= 255) {
                                    if true {
                                        var v36 uint8 = 128
                                        _ = v36
                                    }
                                    var v37 int8 = -128
                                    for fi38 := 0; fi38 < 9; fi38++ {
                                        v4 = v4
                                        for fi39 := 0; fi39 < 3; fi39++ {
                                            v34 = append(v34, (fi38 >> (((v27 + v27)) & 63)))
                                            switch "hello" {
                                            case "foo bar":
                                                v40 := GenS1{F0: 394, F1: 452}
                                                _ = v40
                                            case "a":
                                                fmt.Println(v37, v5, v27)
                                            case "":
                                                v41 := GenS3{F0: v37, F1: ("hello" + "foo bar")}
                                                _ = v41
                                            default:
                                                v34 = append(v34, int(v4))
                                            }
                                            fmt.Println(v5, v3)
                                            _ = fi39
                                        }
                                        _ = fi38
                                    }
                                    _ = v37
                                }
                                _ = v34
                                _ = v35
                            }
                            v42 := []int{127, 4294967295, 42, 2147483647}
                            var v43 int64 = 9223372036854775807
                            _ = v42
                            _ = v43
                        } else {
                            v27 = 862
                        }
                        if (!v1) {
                            if ((0 != 255) && (!true)) {
                                switch "\n\t" {
                                case "hello":
                                    var v44 string = ""
                                    _ = v44
                                case "foo bar":
                                    for fi45 := 0; fi45 < 9; fi45++ {
                                        switch (("\n\t" + "\n\t") + "hello") {
                                        case "":
                                            fmt.Println(v1)
                                        case "foo bar":
                                            for fi46 := 0; fi46 < 7; fi46++ {
                                                v2 -= ((0.0 / float64(float32(-2.25))) + v2)
                                                v2 += v2
                                                fmt.Println(fi46, fi45, fi46)
                                                _ = fi46
                                            }
                                        }
                                        switch "\n\t" {
                                        case "hello":
                                            for fi47 := 0; fi47 < 5; fi47++ {
                                                v48 := []int{}
                                                _ = fi47
                                                _ = v48
                                            }
                                        case "foo bar":
                                            v49 := GenS1{F0: 42, F1: -2147483648}
                                            _ = v49
                                        default:
                                            var v50 uint16 = 222
                                            _ = v50
                                        }
                                        _ = fi45
                                    }
                                case "a":
                                    v27 = 867
                                default:
                                    for fi51 := 0; fi51 < 15; fi51++ {
                                        var v52 uint64 = 0
                                        _ = fi51
                                        _ = v52
                                    }
                                }
                                if v1 {
                                    if (!(0.0 >= v3)) {
                                        var v53 uint8 = genFn1(423, v27, 1000)
                                        if true {
                                            switch (("\n\t" + ("a" + ((("hello" + ("foo bar" + "a")) + ("a" + "\n\t")) + "foo bar"))) + (("\n\t" + (("a" + "") + (("hello" + "") + "foo bar"))) + "\n\t")) {
                                            case "hello":
                                                if v1 {
                                                    fmt.Println(v2, v1)
                                                }
                                            default:
                                                switch "hello" {
                                                case "a":
                                                    if (1 >= uint16(v53)) {
                                                        for fi54 := 0; fi54 < 8; fi54++ {
                                                            var v55 uint8 = genFn1((fi54 / (((-(fi54))) | 1)), v27, 0)
                                                            v56 := []int{fi54, (fi54 / ((fi54) | 1)), 128, fi54}
                                                            _ = fi54
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    } else {
                                                        v2 = v2
                                                        fmt.Println(v2)
                                                    }
                                                case "":
                                                    if v1 {
                                                        v57 := GenS4{F0: 2, F1: 125, F2: 2}
                                                        _ = v57
                                                    }
                                                }
                                            }
                                        }
                                        _ = v53
                                    }
                                    v3 = float32(v2)
                                } else {
                                    switch -2147483648 {
                                    case 9223372036854775807:
                                        v27 = v27
                                    case 912:
                                        v1 = v1
                                    case 32768:
                                        v4 = v4
                                    }
                                }
                                v58 := []int{219}
                                _ = v58
                            }
                            for fi59 := 0; fi59 < 9; fi59++ {
                                var v60 int16 = 425
                                _ = fi59
                                _ = v60
                            }
                            for fi61 := 0; fi61 < 3; fi61++ {
                                var v62 uint32 = v4
                                _ = fi61
                                _ = v62
                            }
                        } else {
                            if (true && v1) {
                                v63 := GenS2{F0: 255}
                                _ = v63
                            }
                            var v64 float64 = v2
                            _ = v64
                        }
                        _ = v28
                    } else {
                        fmt.Println(v4, v5, v4)
                        v65 := GenS5{F0: ("" + "a")}
                        _ = v65
                    }
                case 2147483647:
                    v1 = v1
                }
                _ = v27
            }
        }
        var v66 int = -9223372036854775808
        _ = v66
    }
    var v67 int64 = 42
    var v68 uint32 = v4
    if true {
        v69 := []int{743, 32767}
        _ = v69
    } else {
        if false {
            var v70 uint16 = 255
            _ = v70
        }
    }
    v71 := GenS6{F0: (v68 - 2147483647), F1: 2147483647}
    for fi72 := 0; fi72 < 15; fi72++ {
        if (203 > 42) {
            var v73 string = ("" + (("foo bar" + ("a" + (("hello" + ("a" + ((("\n\t" + "\n\t") + ("foo bar" + "a")) + "\n\t"))) + ("foo bar" + ("foo bar" + "a"))))) + "a"))
            fmt.Println(v5, v68)
            _ = v73
        } else {
            v2 = v2
            v74 := []int{(fi72 >> ((0) & 63)), fi72, -439, (fi72 * fi72)}
            var v75 uint16 = 1
            _ = v74
            _ = v75
        }
        fmt.Println(v3, v2)
        var v76 uint64 = 4294967295
        _ = fi72
        _ = v76
    }
    v77 := []int{127, -560}
    var v78 int = 127
    if v1 {
        v67 = 0
        if v1 {
            fmt.Println(v68, v68)
            var v79 uint16 = 1
            var v80 float32 = v5
            _ = v79
            _ = v80
        }
    }
    if (!true) {
        for fi81 := 0; fi81 < 10; fi81++ {
            var v82 uint8 = 2
            _ = fi81
            _ = v82
        }
        var v83 float64 = v2
        _ = v83
    } else {
        var v84 uint16 = 255
        if false {
            if (376 >= 160) {
                var v85 float32 = (((float32(1.5) / ((0.1 - v3) / (0.0 + (0.1 / v5)))) * 3.14159) - 0.1)
                v86 := GenS7{F0: (float64(v3) < v2), F1: 4294967295}
                _ = v85
                _ = v86
            }
            for fi87 := 0; fi87 < 3; fi87++ {
                v88 := GenS6{F0: v4, F1: 42}
                if v1 {
                    v5 -= v3
                    var v89 uint = 400
                    _ = v89
                } else {
                    for fi90 := 0; fi90 < 11; fi90++ {
                        v5 = ((v3 + (-2.25 - (v5 * v3))) / v5)
                        _ = len(v77)
                        _ = fi90
                    }
                    v91 := GenS1{F0: 2147483647, F1: -2147483648}
                    _ = v91
                }
                var v92 uint = 994
                _ = fi87
                _ = v88
                _ = v92
            }
            var v93 float64 = 0.0
            _ = v93
        }
        _ = v84
    }
    v94 := GenS2{F0: 42}
    var v95 uint64 = 338
    if v1 {
        var v96 float32 = ((v3 / v5) + (0.1 / float32((((v2 / (float64(0.0) - 3.14159)) - v2) / (v2 * 3.14159)))))
        v97 := GenS1{F0: -1, F1: 0}
        v78 = v78
        _ = v96
        _ = v97
    } else {
        switch "hello" {
        case "a":
            switch "\n\t" {
            case "a":
                switch (("" + ("hello" + "a")) + "a") {
                case "hello":
                    v2 = (1e10 * ((v2 - v2) - float64((0.1 / v5))))
                }
            default:
                var v98 uint16 = 65535
                _ = v98
            }
        case "":
            switch v78 {
            case 875:
                for fi99 := 0; fi99 < 10; fi99++ {
                    for fi100 := 0; fi100 < 9; fi100++ {
                        fmt.Println(fi100, v78)
                        fmt.Println(v5, v1)
                        _ = fi100
                    }
                    _ = fi99
                }
            case 42:
                fmt.Println(v2)
            default:
                if (v67 != (v67 ^ (v67 % ((v67) | 1)))) {
                    v77 = append(v77, v78)
                    fmt.Println(v95)
                }
            }
        default:
            for fi101 := 0; fi101 < 13; fi101++ {
                if (((uint32(v95) <= (v68 | 1)) || ((v67 > 42) || (v1 || (!(1 < 1))))) || v1) {
                    if (!((1 <= -579) || (!((int8(v95) >= 1) && false)))) {
                        for fi102 := 0; fi102 < 8; fi102++ {
                            v103 := GenS8{F0: int(v95), F1: 1, F2: v4}
                            v5 = v3
                            v68 -= v4
                            _ = fi102
                            _ = v103
                        }
                        v104 := GenS9{F0: v67}
                        _ = v104
                    }
                    fmt.Println(v3, v78)
                } else {
                    for fi105 := 0; fi105 < 12; fi105++ {
                        fmt.Println(v5, v78, v95)
                        _ = len(v77)
                        v4 += v68
                        _ = fi105
                    }
                }
                if v1 {
                    fmt.Println(v3, v68, v95)
                    v1 = (-32768 < -123)
                    v1 = (genFn0(128, ("foo bar" + "\n\t"), fi101) < 255)
                }
                _ = fi101
            }
        }
    }
    if len(v77) >= 2 {
        v77 = v77[1:len(v77)]
    }
    fmt.Println(v4)
    for fi106 := 0; fi106 < 2; fi106++ {
        if ((("foo bar" + "foo bar") == "\n\t") || false) {
            _ = len(v77)
            v3 = 1e10
        } else {
            for fi107 := 0; fi107 < 6; fi107++ {
                var v108 int32 = 1
                v77 = append(v77, (^(fi106)))
                _ = fi107
                _ = v108
            }
            var v109 int64 = v67
            if ((v1 && v1) || true) {
                v110 := GenS4{F0: (v68 * 42), F1: 0, F2: v4}
                v77 = append(v77, 454)
                _ = v110
            }
            _ = v109
        }
        _ = fi106
    }
    for fi111 := 0; fi111 < 4; fi111++ {
        switch "a" {
        case "":
            var v112 int64 = v67
            _ = v112
        case "a":
            if true {
                v67 -= (v67 % ((v67) | 1))
                _ = len(v77)
            } else {
                if false {
                    var v113 uint16 = 285
                    v77 = append(v77, v78)
                    var v114 uint32 = v68
                    _ = v113
                    _ = v114
                } else {
                    var v115 int = (fi111 | 42)
                    if len(v77) > 0 {
                        v77[0] = v115
                    }
                    var v116 int8 = 121
                    _ = v115
                    _ = v116
                }
                for fi117 := 0; fi117 < 18; fi117++ {
                    if true {
                        if ((124 == -76) && ((-2.25 / 1.5) != 0.1)) {
                            v118 := GenS10{F0: 809, F1: 666, F2: -32768}
                            _ = v118
                        }
                        v77 = append(v77, -1)
                        var v119 string = ("\n\t" + (("a" + "hello") + "foo bar"))
                        _ = v119
                    }
                    _ = fi117
                }
                v3 += v3
            }
        }
        _ = fi111
    }
    var v120 uint32 = (v4 | 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v78:", v78)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v120
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v68
    _ = v71
    _ = v77
    _ = v78
    _ = v94
    _ = v95
}
