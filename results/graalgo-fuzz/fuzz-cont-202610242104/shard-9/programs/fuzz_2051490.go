package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint; F1 int; F2 bool }
type GenS6 struct { F0 string; F1 uint64; F2 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 uint8 }
type GenS10 struct { F0 bool; F1 uint8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint64; F1 uint64; F2 int }
type GenS13 struct { F0 int32; F1 bool; F2 uint64 }
func genFn0(p0 int) uint8 {
    if (42 >= 2) {
        v4 := []int{p0}
        if (true && false) {
            fmt.Println(p0)
            for fi5 := 0; fi5 < 11; fi5++ {
                var v6 float32 = (1e10 * (0.1 / ((3.14159 / 1.5) - 3.14159)))
                _ = fi5
                _ = v6
            }
            v7 := GenS1{F0: 255, F1: 42}
            _ = v7
        }
        _ = v4
    } else {
        fmt.Println(p0)
        fmt.Println(p0)
        for fi8 := 0; fi8 < 6; fi8++ {
            p0 += -267
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 float32 = 1e10
                if true {
                    v11 := GenS1{F0: 459, F1: -1}
                    v12 := GenS1{F0: 1, F1: 0}
                    _ = v11
                    _ = v12
                }
                var v13 int64 = 9223372036854775807
                _ = fi9
                _ = v10
                _ = v13
            }
            _ = fi8
        }
    }
    switch "\n\t" {
    case "":
        var v14 int32 = 1
        _ = v14
    case "hello":
        var v15 string = "a"
        _ = v15
    default:
        p0 = (p0 | p0)
    }
    _ = p0
    return 42
}
func genFn1(p0 uint, p1 bool) int {
    for fi16 := 0; fi16 < 18; fi16++ {
        p1 = p1
        _ = fi16
    }
    v17 := GenS1{F0: 901, F1: 32767}
    v18 := []int{-17, 2147483647, 128, -138}
    _ = len(v18)
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 504
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 1
    if true {
        v3 = v3
    } else {
        v1 = append(v1, ("" + (("a" + "\n\t") + ("a" + "\n\t"))))
        v19 := GenS1{F0: 306, F1: 1}
        _ = v19
    }
    var v20 uint16 = 488
    if ((true && true) && (2 < 983)) {
        v3 -= v3
        if true {
            for fi21 := 0; fi21 < 2; fi21++ {
                for fi22 := 0; fi22 < 20; fi22++ {
                    switch (fi22 >> ((2) & 63)) {
                    case -332:
                        v23 := GenS2{F0: -1, F1: v3, F2: 128}
                        _ = v23
                    }
                    v3 += (v3 >> ((943) & 63))
                    _ = fi22
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    var v25 int32 = 1
                    v20 += (v20 - v20)
                    _ = fi24
                    _ = v25
                }
                fmt.Println(v20, fi21)
                _ = fi21
            }
            v20 += (^(v20))
        }
    } else {
        v3 = 2147483647
    }
    var v26 string = "foo bar"
    var v27 uint8 = genFn0(-619)
    if (32767 > -2147483648) {
        for fi28 := 0; fi28 < 8; fi28++ {
            for fi29 := 0; fi29 < 18; fi29++ {
                if (!true) {
                    v3 += (v3 - v3)
                    v30 := GenS2{F0: -2147483648, F1: v3, F2: 42}
                    if true {
                        switch "\n\t" {
                        case "hello":
                            v31 := GenS2{F0: 2147483647, F1: 1, F2: 73}
                            v3 = v3
                            v32 := GenS2{F0: -2147483648, F1: (v3 & 0), F2: (v27 ^ (v27 & (v27 / (((v27 << ((573) & 63))) | 1))))}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v26 = v26
                            _ = v31
                            _ = v32
                        case "a":
                            if false {
                                v26 = ("a" + v26)
                                v3 += v3
                            } else {
                                v33 := v2[(v3 & 2)]
                                v20 = (v20 & 65535)
                                v27 -= (v27 + 2)
                                v34 := v2[v3]
                                switch ("a" + ("\n\t" + v26)) {
                                case "a":
                                    _ = len(v1)
                                    if (!true) {
                                        v35 := GenS3{F0: v27}
                                        v36 := GenS4{F0: (fi29 | -791), F1: 4294967295}
                                        _ = v35
                                        _ = v36
                                    } else {
                                        v37 := GenS4{F0: (-(fi28)), F1: 606}
                                        _ = v37
                                    }
                                    v26 = "hello"
                                    if (-310 < 956) {
                                        v33 = 4294967295
                                    }
                                case "":
                                    _ = len(v1)
                                    if (42 != int64(v20)) {
                                        var v38 float32 = 3.14159
                                        delete(v2, v34)
                                        _ = v38
                                    } else {
                                        v20 = v20
                                        var v39 uint64 = 4294967295
                                        var v40 bool = true
                                        _ = v39
                                        _ = v40
                                    }
                                case "foo bar":
                                    v33 -= 65535
                                }
                                switch v26 {
                                case "":
                                    var v41 float64 = 0.1
                                    _ = v41
                                case "hello":
                                    if false {
                                        fmt.Println(v27)
                                        if false {
                                            var v42 uint32 = 197
                                            _ = v42
                                        } else {
                                            switch (fi29 << ((942) & 63)) {
                                            case 4294967295:
                                                v27 = genFn0(fi28)
                                            case 9223372036854775807:
                                                switch (v26 + "") {
                                                case "\n\t":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case "a":
                                                    v43 := v2[(^(v33))]
                                                    _ = v43
                                                case "foo bar":
                                                    switch "hello" {
                                                    case "hello":
                                                        switch v26 {
                                                        case "\n\t":
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        default:
                                                            var v44 uint8 = (^(v27))
                                                            _ = v44
                                                        }
                                                    default:
                                                        v45 := GenS2{F0: 172, F1: (^(v3)), F2: (v27 & v27)}
                                                        _ = v45
                                                    }
                                                }
                                            case 101:
                                                var v46 uint = 356
                                                _ = v46
                                            default:
                                                v47 := GenS4{F0: -608, F1: 18446744073709551615}
                                                var v48 float64 = 1e10
                                                v49 := GenS5{F0: 287, F1: (fi28 >> ((1) & 63)), F2: (false || (v26 != v26))}
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                        v3 += v33
                                    } else {
                                        v3 -= 218
                                        if (true || true) {
                                            var v50 float32 = 1.5
                                            fmt.Println(fi29)
                                            v51 := GenS6{F0: v26, F1: 2, F2: 24}
                                            switch (fi28 * fi28) {
                                            case 987:
                                                if false {
                                                    v50 += v50
                                                    v2[v34] = (v3 + uint32(v27))
                                                    if ((v27 + v27) < genFn0((fi28 + fi29))) {
                                                        fmt.Println(v33, v20, v3)
                                                    } else {
                                                        if ((v50 >= ((v50 * ((1e10 / -2.25) - v50)) * (v50 * 1e10))) && (786 < 0)) {
                                                            v52 := GenS7{F0: v33}
                                                            var v53 uint8 = 143
                                                            v53 = genFn0((fi29 | (-(fi29))))
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        switch fi29 {
                                                        case -449:
                                                            if false {
                                                                var v54 bool = (!((v33 < v33) && true))
                                                                if v54 {
                                                                    v34 += v34
                                                                }
                                                                fmt.Println(fi29)
                                                                var v55 int64 = 760
                                                                v56 := v2[v3]
                                                                _ = v54
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                        case 9223372036854775807:
                                                            switch (v26 + "\n\t") {
                                                            case "foo bar":
                                                                v33 -= v33
                                                            }
                                                        default:
                                                            var v57 uint32 = 42
                                                            _ = v57
                                                        }
                                                    }
                                                    v58 := GenS8{F0: (fi28 ^ fi29)}
                                                    _ = v58
                                                } else {
                                                    v26 = ""
                                                    var v59 uint16 = (v20 % ((42) | 1))
                                                    fmt.Println(v3)
                                                    _ = v59
                                                }
                                            default:
                                                switch (fi29 * (fi28 % ((fi29) | 1))) {
                                                case 617:
                                                    v27 += (v27 * uint8(v34))
                                                case -907:
                                                    if (3.14159 == (0.0 - (1e10 * (1e10 + 3.14159)))) {
                                                        v1 = append(v1, v26)
                                                    } else {
                                                        var v60 float64 = 3.14159
                                                        if (!false) {
                                                            switch ("" + "") {
                                                            case "":
                                                                var v61 float32 = v50
                                                                _ = v61
                                                            case "hello":
                                                                var v62 int64 = -405
                                                                _ = v62
                                                            default:
                                                                if len(v1) > 0 {
                                                                    v1[0] = v26
                                                                }
                                                            }
                                                            v63 := GenS5{F0: 42, F1: 9223372036854775807, F2: ((v26 + (v26 + (v26 + "a"))) != v26)}
                                                            if len(v1) > 0 {
                                                                v1[0] = v26
                                                            }
                                                            _ = v63
                                                        }
                                                        var v64 uint32 = v3
                                                        _ = v60
                                                        _ = v64
                                                    }
                                                case 42:
                                                    v20 = v20
                                                default:
                                                    if ((745 > 32) && (1e10 >= 0.1)) {
                                                        v65 := GenS2{F0: -204, F1: v33, F2: (^(v27))}
                                                        fmt.Println(v33)
                                                        v50 = v50
                                                        _ = v65
                                                    } else {
                                                        var v66 bool = true
                                                        _ = v66
                                                    }
                                                }
                                            }
                                            _ = v50
                                            _ = v51
                                        } else {
                                            if true {
                                                if false {
                                                    var v67 uint8 = v27
                                                    v68 := GenS9{F0: v20, F1: 127}
                                                    _ = v67
                                                    _ = v68
                                                } else {
                                                    fmt.Println(fi29)
                                                }
                                                var v69 int16 = 1000
                                                var v70 uint16 = v20
                                                _ = v69
                                                _ = v70
                                            } else {
                                                v71 := GenS6{F0: "a", F1: 870, F2: 42}
                                                if true {
                                                    v3 = uint32(v27)
                                                    v72 := GenS6{F0: "\n\t", F1: 18446744073709551615, F2: 1}
                                                    if ((false && false) && (true && false)) {
                                                        v20 = v20
                                                        if true {
                                                            switch 2147483647 {
                                                            case -117:
                                                                if (((^(v20)) != 445) && true) {
                                                                    var v73 string = ""
                                                                    _ = len(v2)
                                                                    if len(v1) > 0 {
                                                                        v1[0] = v73
                                                                    }
                                                                    _ = v73
                                                                } else {
                                                                    var v74 uint64 = 18446744073709551615
                                                                    var v75 uint8 = v27
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                                v26 = ""
                                                                v33 -= 2
                                                                v2[(v33 - v3)] = v33
                                                            case 512:
                                                                var v76 int8 = 42
                                                                _ = v76
                                                            case 734:
                                                                _ = len(v2)
                                                            }
                                                            if true {
                                                                _ = len(v1)
                                                                var v77 uint8 = v27
                                                                _ = v77
                                                            } else {
                                                                v78 := GenS10{F0: (v3 <= (^(v3))), F1: genFn0(fi28)}
                                                                fmt.Println(v27, v33)
                                                                _ = v78
                                                            }
                                                            if false {
                                                                switch fi28 {
                                                                case 42:
                                                                    var v79 uint16 = (v20 ^ (v20 % ((v20) | 1)))
                                                                    _ = v79
                                                                default:
                                                                    var v80 uint64 = 4294967295
                                                                    var v81 uint8 = v27
                                                                    var v82 int = fi28
                                                                    fmt.Println(v81, fi28)
                                                                    v80 = v80
                                                                    _ = v80
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                            } else {
                                                                v3 += (v34 | (^(v33)))
                                                                var v83 int32 = -2147483648
                                                                fmt.Println(v20)
                                                                _ = v83
                                                            }
                                                        }
                                                    } else {
                                                        if false {
                                                            v34 = 0
                                                        }
                                                        var v84 uint16 = (^(v20))
                                                        if true {
                                                            v85 := GenS9{F0: 1, F1: (^(v27))}
                                                            v3 = 2147483648
                                                            _ = v85
                                                        }
                                                        _ = v84
                                                    }
                                                    if true {
                                                        v34 -= 348
                                                        fmt.Println(v20)
                                                        var v86 int16 = 310
                                                        _ = v86
                                                    }
                                                    _ = v72
                                                } else {
                                                    v87 := v2[776]
                                                    var v88 bool = (false || false)
                                                    _ = v87
                                                    _ = v88
                                                }
                                                switch "a" {
                                                case "hello":
                                                    var v89 int16 = 42
                                                    _ = v89
                                                case "foo bar":
                                                    var v90 uint16 = (v20 / ((2) | 1))
                                                    _ = v90
                                                default:
                                                    v91 := GenS11{F0: v27}
                                                    if (1 < 42) {
                                                        var v92 int8 = 127
                                                        _ = v92
                                                    }
                                                    _ = v91
                                                }
                                                _ = v71
                                            }
                                        }
                                    }
                                default:
                                    v93 := GenS5{F0: 299, F1: (fi28 / ((-2147483648) | 1)), F2: (!true)}
                                    _ = v93
                                }
                                v20 -= (v20 * (v20 / ((32767) | 1)))
                                _ = v33
                                _ = v34
                            }
                        }
                        if true {
                            delete(v2, v3)
                            if true {
                                fmt.Println(fi29, v3)
                                switch ((v26 + v26) + ("\n\t" + v26)) {
                                case "":
                                    fmt.Println(v3, v3)
                                case "\n\t":
                                    v27 -= v27
                                }
                                v94 := GenS6{F0: v26, F1: 2, F2: 18}
                                _ = v94
                            } else {
                                if true {
                                    var v95 uint32 = v3
                                    _ = v95
                                }
                                v27 = v27
                                switch "a" {
                                case "\n\t":
                                    if (0.0 <= (float64((float32(1.5) + 0.0)) + ((((1e10 + (1.5 * 1.5)) - 0.1) / 1.5) + (float64(1e10) + 0.1)))) {
                                        fmt.Println(v27, fi29, v3)
                                        v3 = v3
                                        v26 = ("hello" + v26)
                                    } else {
                                        v26 = v26
                                        if false {
                                            var v96 int32 = 2147483647
                                            _ = v96
                                        } else {
                                            v3 = v3
                                        }
                                    }
                                case "":
                                    _ = len(v1)
                                case "foo bar":
                                    v26 = (v26 + "foo bar")
                                default:
                                    if ((-2.25 * ((float32(1e10) * (1.5 * 0.1)) - 1.5)) == 1e10) {
                                        switch "a" {
                                        case "\n\t":
                                            v2[(v3 - (^(v3)))] = 853
                                        default:
                                            fmt.Println(fi29, v20)
                                        }
                                    }
                                }
                            }
                        }
                    }
                    _ = v30
                }
                if false {
                    var v97 float64 = 0.1
                    _ = v97
                }
                _ = fi29
            }
            v27 += (v27 ^ 0)
            _ = fi28
        }
        v98 := GenS12{F0: 42, F1: 0, F2: 32768}
        if (false || (-1 != -50)) {
            for fi99 := 0; fi99 < 12; fi99++ {
                v3 -= 1
                v100 := GenS9{F0: v20, F1: (v27 & 127)}
                _ = fi99
                _ = v100
            }
        }
        _ = v98
    }
    for fi101 := 0; fi101 < 7; fi101++ {
        v102 := GenS6{F0: ("\n\t" + v26), F1: uint64(v20), F2: 42}
        var v103 uint = 4294967295
        for fi104 := 0; fi104 < 12; fi104++ {
            _ = len(v1)
            _ = fi104
        }
        _ = fi101
        _ = v102
        _ = v103
    }
    var v105 int32 = 954
    v106 := GenS11{F0: genFn0(42)}
    for fi107 := 0; fi107 < 2; fi107++ {
        if len(v1) > 0 {
            v1[0] = v26
        }
        var v108 uint64 = 18446744073709551615
        v105 = (v105 ^ (-(v105)))
        _ = fi107
        _ = v108
    }
    for fi109 := 0; fi109 < 15; fi109++ {
        if false {
            var v110 int16 = 32767
            for fi111 := 0; fi111 < 1; fi111++ {
                v26 = v26
                _ = fi111
            }
            _ = v110
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 += 682
        _ = fi109
    }
    v112 := GenS13{F0: 1, F1: false, F2: 419}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v105:", v105)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v105
    _ = v106
    _ = v112
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
}
