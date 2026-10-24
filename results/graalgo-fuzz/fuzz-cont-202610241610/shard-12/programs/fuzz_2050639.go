package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int8; F2 string }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 string; F1 uint8; F2 string }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16; F1 int32; F2 string }
type GenS8 struct { F0 int64; F1 string }
type GenS9 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS10 struct { F0 uint; F1 int8 }
func genFn0(p0 uint8, p1 int64) int32 {
    if true {
        if false {
            fmt.Println(p1, p1)
            p0 -= p0
            p1 -= p1
        } else {
            p0 += (p0 + 42)
        }
    }
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 0}
        _ = fi5
        _ = v6
    }
    if true {
        for fi7 := 0; fi7 < 12; fi7++ {
            fmt.Println(p1, p0, p1)
            _ = fi7
        }
        var v8 float32 = 0.1
        _ = v8
    } else {
        if (-32768 < 42) {
            p1 += -1
        }
    }
    var v9 string = "a"
    var v10 uint32 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 42
}
func genFn1(p0 uint16, p1 uint8, p2 uint8) uint64 {
    switch "hello" {
    case "foo bar":
        if (!(((-2147483648 == genFn0(128, -1)) && false) && (p0 >= (p0 | p0)))) {
            for fi11 := 0; fi11 < 8; fi11++ {
                if (true || (0 >= 2147483647)) {
                    v12 := GenS1{F0: 32767}
                    _ = v12
                } else {
                    p1 += p2
                }
                _ = fi11
            }
            if ((!false) && true) {
                v13 := []int{231, 77}
                if (((1e10 - (1e10 / (0.1 * (1e10 + ((0.0 / float64((1.5 - 3.14159))) * 1e10))))) + (-2.25 / ((-2.25 / 1.0) * (float64(1e10) / 1.5)))) > 1e10) {
                    v14 := GenS2{F0: 42, F1: 127, F2: "foo bar"}
                    _ = v14
                } else {
                    var v15 uint16 = p0
                    _ = v15
                }
                _ = v13
            }
        } else {
            var v16 uint = 255
            var v17 int16 = -1
            v18 := []int{-9223372036854775808, 170}
            _ = v16
            _ = v17
            _ = v18
        }
    case "a":
        switch 127 {
        case 448:
            fmt.Println(p2, p2, p2)
        default:
            switch 2147483647 {
            case -2147483648:
                p2 += 191
            case -693:
                var v19 int32 = 2147483647
                _ = v19
            case -9223372036854775808:
                p2 = (p1 * (p2 ^ p1))
            default:
                if true {
                    v20 := []int{1, 916, -2147483648}
                    var v21 int32 = -2147483648
                    _ = v20
                    _ = v21
                }
            }
        }
    }
    var v22 uint64 = 172
    var v23 int32 = -2147483648
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 uint8 = 0
        fmt.Println(v22, v22)
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return 42
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 255
    var v3 uint16 = 255
    var v4 int8 = 100
    var v26 uint64 = 149
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 uint64 = (v26 - (v26 * (^(v26))))
        v29 := GenS1{F0: 211}
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        for fi31 := 0; fi31 < 6; fi31++ {
            fmt.Println(v1, fi31)
            var v32 uint32 = 781
            v2 -= (v2 << ((uint(v3)) & 63))
            _ = fi31
            _ = v32
        }
        fmt.Println(v2, v3, v2)
        _ = fi30
    }
    v4 -= 0
    v3 = 42
    switch "foo bar" {
    case "foo bar":
        v33 := GenS2{F0: -885, F1: v4, F2: (("a" + ("hello" + "\n\t")) + "hello")}
        _ = v33
    default:
        var v34 uint32 = 755
        _ = v34
    }
    if false {
        v35 := []int{-1, 9223372036854775807, 32767, -9223372036854775808}
        v36 := GenS1{F0: 1000}
        _ = v35
        _ = v36
    } else {
        switch "foo bar" {
        case "":
            v37 := []int{4294967295}
            _ = v37
        case "a":
            v38 := []int{int(v2), 128, 128}
            _ = v38
        default:
            v39 := []int{511, 1, -556, 0}
            _ = v39
        }
        var v40 bool = (!((v4 | 0) == v4))
        _ = v40
    }
    if false {
        fmt.Println(v3, v4)
        v3 = v3
    } else {
        v1 -= 0
        v41 := []int{}
        fmt.Println(v4, v3)
        _ = v41
    }
    switch "foo bar" {
    case "\n\t":
        if false {
            v42 := GenS2{F0: -9223372036854775808, F1: (v4 ^ v4), F2: ("\n\t" + "hello")}
            _ = v42
        } else {
            v3 = 353
            v43 := GenS3{F0: (v1 + v1), F1: -1}
            v2 += 255
            _ = v43
        }
    default:
        v44 := []int{}
        _ = v44
    }
    if (true || false) {
        v45 := GenS2{F0: 2, F1: (v4 % ((int8(v3)) | 1)), F2: ("a" + "a")}
        var v46 uint = 42
        var v47 uint32 = 2147483648
        _ = v45
        _ = v46
        _ = v47
    } else {
        for fi48 := 0; fi48 < 16; fi48++ {
            var v49 float64 = (0.1 / ((-2.25 / 1.0) * ((3.14159 + -2.25) / (-2.25 + 0.0))))
            var v50 int8 = v4
            for fi51 := 0; fi51 < 9; fi51++ {
                var v52 float32 = (0.0 + 0.0)
                _ = fi51
                _ = v52
            }
            _ = fi48
            _ = v49
            _ = v50
        }
        var v53 uint64 = genFn1((v3 / ((v3) | 1)), 1, (v2 - 42))
        var v54 uint = 1
        _ = v53
        _ = v54
    }
    if false {
        if ((v2 ^ 255) == v2) {
            v55 := GenS1{F0: 42}
            switch ("hello" + ("foo bar" + "")) {
            case "\n\t":
                v26 += 182
            case "":
                var v56 uint8 = (v2 | (v2 | (v2 ^ 0)))
                _ = v56
            case "a":
                var v57 uint32 = 120
                _ = v57
            }
            _ = v55
        }
        var v58 int8 = (v4 & 42)
        v26 = genFn1(255, 42, v2)
        _ = v58
    } else {
        v3 = v3
        v4 += (v4 + 100)
        for fi59 := 0; fi59 < 20; fi59++ {
            switch "foo bar" {
            case "a":
                v60 := []int{(fi59 & fi59), (fi59 >> ((0) & 63)), fi59, -650}
                _ = v60
            case "foo bar":
                for fi61 := 0; fi61 < 19; fi61++ {
                    v26 -= (v26 >> ((65535) & 63))
                    _ = fi61
                }
            case "":
                var v62 uint64 = 1
                _ = v62
            }
            var v63 uint = 1
            _ = fi59
            _ = v63
        }
    }
    switch -941 {
    case -9223372036854775808:
        fmt.Println(v3, v3, v2)
    case 692:
        for fi64 := 0; fi64 < 12; fi64++ {
            v65 := GenS4{F0: ("" + "\n\t"), F1: 128, F2: ("hello" + "")}
            v1 -= (v1 ^ (^(v1)))
            if (0.0 <= -2.25) {
                if false {
                    switch (fi64 + (fi64 & fi64)) {
                    case -601:
                        var v66 int32 = -1
                        _ = v66
                    case 900:
                        v67 := []int{(fi64 * (fi64 * (-(fi64)))), 4294967295, (fi64 + (fi64 - (fi64 | int(v1)))), int(v1)}
                        _ = v67
                    case -9223372036854775808:
                        switch "hello" {
                        case "foo bar":
                            v68 := GenS5{F0: 52, F1: fi64}
                            _ = v68
                        }
                    default:
                        for fi69 := 0; fi69 < 12; fi69++ {
                            var v70 uint16 = (v3 % ((v3) | 1))
                            fmt.Println(v2, fi64, v70)
                            _ = fi69
                            _ = v70
                        }
                    }
                    var v71 uint32 = 55
                    _ = v71
                } else {
                    if true {
                        for fi72 := 0; fi72 < 9; fi72++ {
                            switch "a" {
                            case "foo bar":
                                v73 := []int{-1, (fi72 | fi64)}
                                _ = v73
                            default:
                                if (140 >= 397) {
                                    v4 = (v4 + 1)
                                    if true {
                                        v1 = (v1 + genFn0((v2 & 127), -912))
                                        var v74 int = 32767
                                        _ = v74
                                    } else {
                                        v75 := GenS3{F0: 0, F1: (v4 >> ((2) & 63))}
                                        fmt.Println(v3, v1)
                                        _ = v75
                                    }
                                    if true {
                                        var v76 int64 = -9223372036854775808
                                        _ = v76
                                    } else {
                                        v77 := GenS6{F0: 42}
                                        var v78 uint = 255
                                        _ = v77
                                        _ = v78
                                    }
                                }
                            }
                            var v79 int16 = -32768
                            switch fi72 {
                            case 0:
                                fmt.Println(v26)
                            case 32768:
                                if ((2147483648 == 4294967295) || (!false)) {
                                    v80 := []int{fi64, 2147483647}
                                    if ("a" == "foo bar") {
                                        switch (("\n\t" + "hello") + "") {
                                        case "hello":
                                            var v81 bool = false
                                            _ = v81
                                        }
                                        var v82 int32 = 0
                                        if len(v80) >= 2 {
                                            v80 = v80[1:len(v80)]
                                        }
                                        if len(v80) > 0 {
                                            v80[0] = -9223372036854775808
                                        }
                                        _ = v82
                                    } else {
                                        v83 := GenS6{F0: 4294967295}
                                        v84 := GenS5{F0: v4, F1: (fi64 / ((fi64) | 1))}
                                        var v85 uint64 = genFn1((v3 - v3), v2, (^(v2)))
                                        fmt.Println(v85, v3)
                                        v4 = 127
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    }
                                    _ = v80
                                }
                            case -1:
                                if true {
                                    v86 := GenS1{F0: -1}
                                    var v87 string = ""
                                    if ((-(fi64)) >= fi64) {
                                        if ((^(v26)) > uint64(v3)) {
                                            v88 := []int{fi72}
                                            var v89 int8 = (v4 / ((127) | 1))
                                            _ = v88
                                            _ = v89
                                        } else {
                                            v90 := []int{(fi64 | fi64), (fi72 / ((fi72) | 1)), fi72}
                                            _ = v90
                                        }
                                        if (!false) {
                                            var v91 bool = (!((v3 - (v3 * v3)) < (v3 % (((v3 | (v3 << ((uint(v2)) & 63)))) | 1))))
                                            fmt.Println(v3, v87)
                                            v92 := []int{(fi64 & (fi64 & (fi72 << ((18446744073709551615) & 63)))), fi72, fi64}
                                            _ = v91
                                            _ = v92
                                        } else {
                                            v4 = (^(v4))
                                            fmt.Println(v2, fi64, v87)
                                            if false {
                                                v87 = v87
                                                v93 := GenS1{F0: (v79 % ((v79) | 1))}
                                                _ = v93
                                            }
                                        }
                                    } else {
                                        var v94 uint16 = (^(v3))
                                        if (!false) {
                                            v79 = v79
                                        }
                                        if true {
                                            v95 := GenS7{F0: int16(v2), F1: v1, F2: v87}
                                            if ((false && (0.1 > 0.0)) && (true && false)) {
                                                v79 = 707
                                            } else {
                                                v96 := []int{(^(fi72)), fi72, fi72}
                                                _ = v96
                                            }
                                            v97 := []int{(fi64 / (((fi72 << ((4294967295) & 63))) | 1)), (fi72 - 32767)}
                                            var v98 int8 = int8(v26)
                                            _ = v95
                                            _ = v97
                                            _ = v98
                                        }
                                        _ = v94
                                    }
                                    v1 += 1
                                    switch (fi72 >> ((371) & 63)) {
                                    case 1:
                                        v26 += v26
                                    case 0:
                                        switch v87 {
                                        case "\n\t":
                                            v99 := []int{(fi64 & fi64), fi64, (fi64 & fi72), 127}
                                            _ = v99
                                        case "":
                                            v100 := []int{(fi64 << ((1) & 63)), (fi64 % (((^(fi64))) | 1))}
                                            _ = v100
                                        case "foo bar":
                                            var v101 uint64 = 2
                                            _ = v101
                                        default:
                                            var v102 uint = 1
                                            _ = v102
                                        }
                                    case 128:
                                        if false {
                                            v87 = "foo bar"
                                        } else {
                                            v103 := GenS7{F0: 1, F1: genFn0((v2 + v2), 101), F2: v87}
                                            var v104 uint = 4294967295
                                            v3 -= (v3 * (v3 % ((v3) | 1)))
                                            fmt.Println(v1, fi72, fi72)
                                            var v105 uint = (v104 >> ((v104) & 63))
                                            v106 := GenS8{F0: 9223372036854775807, F1: v87}
                                            if (true && false) {
                                                var v107 uint32 = 878
                                                var v108 int32 = int32(fi64)
                                                var v109 uint16 = (v3 - v3)
                                                _ = v107
                                                _ = v108
                                                _ = v109
                                            }
                                            _ = v103
                                            _ = v104
                                            _ = v105
                                            _ = v106
                                        }
                                    }
                                    v79 -= int16(v2)
                                    v4 += 0
                                    _ = v86
                                    _ = v87
                                } else {
                                    v3 = (v3 - 32767)
                                    v110 := []int{fi72, fi64}
                                    v4 -= 61
                                    _ = v110
                                }
                                v111 := GenS9{F0: (v2 % ((v2) | 1)), F1: (v3 ^ (^(v3))), F2: 724}
                                var v112 float64 = 1e10
                                _ = v111
                                _ = v112
                            }
                            _ = fi72
                            _ = v79
                        }
                    }
                    v113 := GenS10{F0: 0, F1: (v4 ^ (^(v4)))}
                    if true {
                        var v114 int8 = (^(v4))
                        var v115 uint = 255
                        _ = v114
                        _ = v115
                    }
                    _ = v113
                }
                v3 = (v3 % ((v3) | 1))
                fmt.Println(v1, v4, v2)
            }
            _ = fi64
            _ = v65
        }
    }
    v116 := []int{-9223372036854775808, 379, 0, 127}
    var v117 uint = uint(v26)
    fmt.Println(v117)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v116:", len(v116), v116)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v116
    _ = v117
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
