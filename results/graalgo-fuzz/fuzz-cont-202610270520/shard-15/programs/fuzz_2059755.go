package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8; F1 uint32; F2 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint8; F1 int16; F2 uint16 }
func genFn0(p0 string, p1 uint32) uint32 {
    switch ((p0 + (p0 + p0)) + ("a" + (p0 + p0))) {
    case "foo bar":
        switch ((("" + (p0 + (p0 + p0))) + ((p0 + p0) + p0)) + ("" + "\n\t")) {
        case "a":
            var v6 int16 = 0
            _ = v6
        }
    case "a":
        for fi7 := 0; fi7 < 13; fi7++ {
            if (!true) {
                var v8 float64 = -2.25
                var v9 uint64 = 4294967295
                v10 := []int{}
                _ = v10
                _ = v8
                _ = v9
            } else {
                var v11 string = ""
                var v12 int64 = -1
                if false {
                    fmt.Println(p1, v11, v11)
                } else {
                    v13 := GenS1{F0: -106}
                    _ = v13
                }
                _ = v11
                _ = v12
            }
            var v14 float32 = (0.0 + -2.25)
            v15 := []int{32768, 2147483647, (-(fi7))}
            _ = fi7
            _ = v14
            _ = v15
        }
    }
    var v16 string = (p0 + p0)
    _ = p0
    _ = p1
    _ = v16
    return (p1 + 65535)
}
func genFn1() int32 {
    for fi17 := 0; fi17 < 7; fi17++ {
        if (1 <= 846) {
            fmt.Println(fi17)
            v18 := GenS2{F0: 0, F1: -32768, F2: 180}
            _ = v18
        } else {
            var v19 int64 = 0
            v20 := GenS1{F0: -92}
            if (!true) {
                for fi21 := 0; fi21 < 8; fi21++ {
                    fmt.Println(v19, fi17)
                    var v22 float32 = (3.14159 / 3.14159)
                    v23 := GenS2{F0: 1, F1: -622, F2: 253}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    if false {
        var v24 float64 = ((-2.25 / 3.14159) - 1.5)
        _ = v24
    }
    var v25 float32 = -2.25
    for fi26 := 0; fi26 < 8; fi26++ {
        switch fi26 {
        case 32768:
            var v27 int = 1
            _ = v27
        case 32767:
            if true {
                v28 := []int{(^(fi26))}
                for fi29 := 0; fi29 < 9; fi29++ {
                    v25 = v25
                    v30 := GenS3{F0: 0, F1: 694}
                    _ = fi29
                    _ = v30
                }
                var v31 int32 = -2147483648
                _ = v28
                _ = v31
            } else {
                var v32 bool = ((true && (-844 > 42)) || (42 == 42))
                v33 := []int{(fi26 + -197), (fi26 % ((-1) | 1)), -1, 9223372036854775807}
                var v34 int32 = -2147483648
                _ = v32
                _ = v33
                _ = v34
            }
        }
        fmt.Println(fi26, fi26)
        v25 -= v25
        _ = fi26
    }
    switch ("a" + "hello") {
    case "":
        var v35 bool = false
        _ = v35
    case "a":
        v25 -= v25
    }
    _ = v25
    return 164
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 int8 = 100
    var v5 uint32 = (v1 ^ v1)
    if true {
        delete(v2, v5)
        var v36 uint16 = 42
        _ = v36
    }
    var v37 int64 = 42
    var v38 int32 = 820
    if (-1 >= 42) {
        for fi39 := 0; fi39 < 17; fi39++ {
            var v40 int64 = -1
            fmt.Println(fi39, v40, v40)
            var v41 float32 = (0.1 - -2.25)
            _ = fi39
            _ = v40
            _ = v41
        }
        for fi42 := 0; fi42 < 10; fi42++ {
            v43 := GenS2{F0: genFn0(("hello" + ""), v5), F1: 1, F2: 255}
            _ = fi42
            _ = v43
        }
        var v44 bool = (!(!((2 != 4294967295) && true)))
        _ = v44
    } else {
        v1 -= (^(v5))
    }
    switch "foo bar" {
    case "hello":
        switch -2147483648 {
        case 2147483647:
            var v45 int32 = v38
            _ = v45
        }
    case "a":
        if (0 == 65535) {
            v38 += (v38 - (^(v38)))
            for fi46 := 0; fi46 < 8; fi46++ {
                var v47 uint = 15
                switch "hello" {
                case "\n\t":
                    var v48 int = fi46
                    _ = v48
                default:
                    if false {
                        var v49 int16 = 677
                        for fi50 := 0; fi50 < 11; fi50++ {
                            v49 -= v49
                            v51 := v2[genFn0("", (^(v1)))]
                            _ = fi50
                            _ = v51
                        }
                        v52 := []int{-9223372036854775808, fi46}
                        _ = v49
                        _ = v52
                    }
                }
                var v53 uint64 = 549
                _ = fi46
                _ = v47
                _ = v53
            }
            v54 := GenS2{F0: v5, F1: 32767, F2: 127}
            _ = v54
        }
    }
    v1 += 227
    fmt.Println(v4)
    var v55 int16 = 412
    for fi56 := 0; fi56 < 18; fi56++ {
        v38 = v38
        v57 := GenS2{F0: 0, F1: 42, F2: 128}
        var v58 int16 = v55
        _ = fi56
        _ = v57
        _ = v58
    }
    for fi59 := 0; fi59 < 8; fi59++ {
        for fi60 := 0; fi60 < 14; fi60++ {
            if false {
                v61 := GenS4{F0: ((("a" + "foo bar") + ("a" + "")) + "\n\t"), F1: false}
                _ = v61
            }
            if (!false) {
                v62 := []int{-1, (fi59 | fi60)}
                if (!false) {
                    switch 9223372036854775807 {
                    case 987:
                        var v63 uint32 = v1
                        _ = v63
                    case -708:
                        _ = len(v62)
                    case 1:
                        v64 := GenS5{F0: "hello", F1: false}
                        _ = v64
                    }
                    var v65 int16 = 1
                    _ = v65
                } else {
                    v4 += (v4 / ((0) | 1))
                    v4 -= v4
                    v5 -= (v1 & 0)
                }
                _ = v62
            } else {
                if false {
                    v37 = (v37 ^ 9223372036854775807)
                    switch "" {
                    case "a":
                        fmt.Println(fi60, v38, v55)
                    case "foo bar":
                        v5 -= v1
                        v4 -= v4
                        v66 := []int{128, 32768}
                        if (uint16(v1) < 140) {
                            v67 := GenS6{F0: 1, F1: int8(v1)}
                            var v68 float32 = 3.14159
                            _ = v67
                            _ = v68
                        } else {
                            var v69 uint8 = 127
                            _ = v69
                        }
                        _ = v66
                    }
                    v70 := []int{fi60, fi60}
                    _ = v70
                } else {
                    var v71 uint64 = 0
                    fmt.Println(v55, v1)
                    var v72 uint64 = v71
                    _ = v71
                    _ = v72
                }
                switch "foo bar" {
                case "foo bar":
                    v55 = (v55 % (((v55 - (v55 % ((-195) | 1)))) | 1))
                case "a":
                    if false {
                        v1 += (v1 + (v1 + 2147483648))
                        v2[(v1 + (v5 / (((v5 >> ((306) & 63))) | 1)))] = fi59
                        var v73 float32 = 0.0
                        _ = v73
                    } else {
                        if false {
                            fmt.Println(v37)
                            v74 := GenS6{F0: 128, F1: 0}
                            var v75 float64 = -2.25
                            _ = v74
                            _ = v75
                        } else {
                            fmt.Println(v5, fi59, v4)
                            var v76 float64 = 0.0
                            var v77 uint8 = 2
                            _ = v76
                            _ = v77
                        }
                        v55 += -38
                        var v78 uint32 = (v5 / ((v1) | 1))
                        _ = v78
                    }
                default:
                    fmt.Println(v4, v37, v5)
                }
                fmt.Println(v38)
            }
            fmt.Println(v37, v5)
            _ = fi60
        }
        v1 = v5
        _ = fi59
    }
    v4 = v4
    if (1e10 <= 0.1) {
        v79 := GenS1{F0: (v4 >> ((42) & 63))}
        v80 := []int{-9223372036854775808, -1, -9223372036854775808, -755}
        fmt.Println(v37, v55, v55)
        _ = v79
        _ = v80
    } else {
        for fi81 := 0; fi81 < 5; fi81++ {
            v55 -= -414
            _ = fi81
        }
        fmt.Println(v37, v37)
    }
    var v82 uint16 = 65535
    v83 := []int{32768, 1, -368}
    v4 += (-(v4))
    fmt.Println(v55, v37)
    if true {
        switch "\n\t" {
        case "hello":
            var v84 uint64 = 864
            _ = v84
        case "\n\t":
            for fi85 := 0; fi85 < 18; fi85++ {
                v55 -= v55
                var v86 int = (fi85 ^ fi85)
                if true {
                    if (v4 != (v4 / ((73) | 1))) {
                        fmt.Println(fi85, fi85)
                        if true {
                            v87 := GenS1{F0: (v4 / ((-1) | 1))}
                            v38 += genFn1()
                            v4 = (v4 - v4)
                            _ = v87
                        }
                        v88 := GenS3{F0: 255, F1: 42}
                        _ = v88
                    } else {
                        switch fi85 {
                        case 32767:
                            for fi89 := 0; fi89 < 15; fi89++ {
                                switch fi89 {
                                case 4294967295:
                                    switch fi89 {
                                    case 127:
                                        var v90 float64 = 0.0
                                        _ = v90
                                    }
                                case -9223372036854775808:
                                    if true {
                                        if true {
                                            v82 = (v82 << ((4294967295) & 63))
                                        }
                                        fmt.Println(fi85, v4, v37)
                                        v91 := v3[fi85]
                                        v92 := GenS3{F0: v82, F1: 821}
                                        _ = v91
                                        _ = v92
                                    } else {
                                        if true {
                                            var v93 int16 = v55
                                            _ = v93
                                        }
                                    }
                                }
                                if (42 < 407) {
                                    v94 := GenS3{F0: (v82 >> ((65535) & 63)), F1: 235}
                                    var v95 uint = 913
                                    switch "\n\t" {
                                    case "":
                                        v96 := GenS6{F0: 255, F1: v4}
                                        _ = v96
                                    default:
                                        switch "a" {
                                        case "foo bar":
                                            fmt.Println(v37)
                                            fmt.Println(v55)
                                        default:
                                            switch "" {
                                            case "":
                                                v83 = append(v83, int(v38))
                                            case "\n\t":
                                                if (2 > 18446744073709551615) {
                                                    var v97 uint8 = 255
                                                    v5 = v1
                                                    var v98 uint8 = v97
                                                    _ = v97
                                                    _ = v98
                                                }
                                            }
                                        }
                                    }
                                    _ = v94
                                    _ = v95
                                } else {
                                    switch ("foo bar" + "foo bar") {
                                    case "hello":
                                        v5 = genFn0((((("" + "a") + (("a" + "a") + "\n\t")) + "") + "foo bar"), uint32(v82))
                                    }
                                    _ = len(v2)
                                    v83 = append(v83, (-(v86)))
                                }
                                _ = fi89
                            }
                        case 128:
                            fmt.Println(v4, fi85, v55)
                        case 257:
                            v99 := GenS7{F0: v4}
                            _ = v99
                        default:
                            if (true && true) {
                                var v100 uint64 = 2
                                delete(v3, v86)
                                _ = v100
                            }
                        }
                    }
                    v38 = (v38 ^ genFn1())
                }
                _ = fi85
                _ = v86
            }
        }
    } else {
        if len(v83) >= 2 {
            v83 = v83[1:len(v83)]
        }
        if false {
            for fi101 := 0; fi101 < 17; fi101++ {
                var v102 int64 = (v37 / ((42) | 1))
                if ((v4 << ((2) & 63)) < 89) {
                    switch "" {
                    case "a":
                        v103 := v2[v5]
                        _ = v103
                    case "hello":
                        if (!(0 >= 4294967295)) {
                            var v104 int16 = v55
                            fmt.Println(v102)
                            _ = v104
                        } else {
                            var v105 int64 = (v102 & int64(v55))
                            v1 += genFn0("", (v1 % (((v1 % (((^(v1))) | 1))) | 1)))
                            _ = v105
                        }
                    }
                    var v106 int8 = (v4 % ((0) | 1))
                    _ = v106
                } else {
                    var v107 uint16 = v82
                    _ = v107
                }
                _ = len(v83)
                _ = fi101
                _ = v102
            }
            for fi108 := 0; fi108 < 12; fi108++ {
                var v109 string = ("hello" + (("\n\t" + "\n\t") + "foo bar"))
                if false {
                    delete(v2, 1)
                }
                _ = fi108
                _ = v109
            }
        } else {
            var v110 uint = 929
            if ((-(v4)) != -1) {
                switch -81 {
                case 42:
                    switch "" {
                    case "\n\t":
                        delete(v2, genFn0("foo bar", v1))
                    case "hello":
                        if false {
                            _ = len(v83)
                        }
                    }
                case 2147483647:
                    switch -883 {
                    case 32768:
                        v111 := GenS8{F0: -16, F1: 840, F2: true}
                        _ = v111
                    default:
                        for fi112 := 0; fi112 < 12; fi112++ {
                            switch "a" {
                            case "hello":
                                fmt.Println(v4)
                            case "\n\t":
                                var v113 int32 = (v38 * v38)
                                _ = v113
                            }
                            v37 = v37
                            _ = fi112
                        }
                    }
                }
            } else {
                var v114 int64 = 1
                v114 += (v37 + -9223372036854775808)
                v115 := GenS2{F0: (v1 * (v1 ^ v5)), F1: v55, F2: 42}
                _ = v114
                _ = v115
            }
            _ = len(v83)
            _ = v110
        }
        v116 := GenS1{F0: (v4 / ((v4) | 1))}
        _ = v116
    }
    switch -2147483648 {
    case -969:
        for fi117 := 0; fi117 < 16; fi117++ {
            fmt.Println(v82)
            if false {
                v118 := GenS9{F0: "foo bar"}
                _ = v118
            } else {
                v119 := GenS10{F0: 214, F1: -1, F2: (v82 + v82)}
                _ = v119
            }
            v120 := GenS3{F0: (v82 + (^(v82))), F1: 991}
            _ = fi117
            _ = v120
        }
    case 128:
        v55 = (v55 >> ((500) & 63))
    default:
        v121 := GenS2{F0: (v1 & genFn0("\n\t", (v1 * 2147483647))), F1: (^(v55)), F2: 142}
        _ = v121
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v82:", v82)
    fmt.Println("slicelen v83:", len(v83), v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v55
    _ = v82
    _ = v83
}
