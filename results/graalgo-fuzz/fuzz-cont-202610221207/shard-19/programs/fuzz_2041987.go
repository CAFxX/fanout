package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int16; F2 int64 }
type GenS4 struct { F0 bool; F1 int64; F2 uint8 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint64; F1 int64; F2 string }
type GenS8 struct { F0 int64; F1 bool; F2 bool }
type GenS9 struct { F0 uint32; F1 uint64; F2 uint32 }
func genFn0(p0 int8, p1 int64, p2 uint32) int16 {
    p1 = (p1 + 42)
    if true {
        switch "\n\t" {
        case "hello":
            v6 := []int{-1, -2147483648, 32768}
            _ = v6
        }
        if ((1 <= 42) || false) {
            p0 += (p0 / ((-111) | 1))
        }
        v7 := GenS1{F0: 2, F1: 47}
        _ = v7
    } else {
        var v8 uint32 = p2
        fmt.Println(v8, p0)
        _ = v8
    }
    if (!(true || (p0 != (p0 - p0)))) {
        v9 := []int{0, -765, 9223372036854775807, -124}
        _ = v9
    }
    switch -9223372036854775808 {
    case 443:
        if (-680 > 0) {
            var v10 int = 835
            var v11 float64 = (1e10 * float64(1.5))
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 8; fi12++ {
                fmt.Println(fi12, p1, p2)
                p1 -= p1
                v13 := []int{}
                _ = fi12
                _ = v13
            }
        }
    case 127:
        p0 += 118
    case 32767:
        v14 := GenS2{F0: (p1 >> ((642) & 63)), F1: p2}
        _ = v14
    }
    p1 -= (^(p1))
    _ = p0
    _ = p1
    _ = p2
    return 32767
}
func genFn1(p0 int8, p1 uint64) bool {
    v15 := []int{}
    var v16 int8 = p0
    v17 := GenS1{F0: 628, F1: 2}
    var v18 int32 = 2147483647
    if true {
        var v19 int16 = -32768
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    return true
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 576
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    var v5 uint64 = 0
    v2 = 42
    v2 = 0
    switch "foo bar" {
    case "foo bar":
        for fi20 := 0; fi20 < 8; fi20++ {
            v5 = (v5 & v5)
            if (1e10 >= (float64(1.5) * -2.25)) {
                v21 := GenS2{F0: 156, F1: 164}
                v22 := GenS3{F0: 42, F1: 0, F2: -1}
                _ = v21
                _ = v22
            } else {
                var v23 uint16 = 65535
                _ = v23
            }
            if (true && ((false || (0 == 834)) && false)) {
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 int32 = v2
                    v26 := GenS1{F0: 888, F1: 42}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                if ("" < "\n\t") {
                    var v27 int64 = 9223372036854775807
                    _ = v27
                } else {
                    if (v5 <= v5) {
                        fmt.Println(v5, v2, fi20)
                    } else {
                        v28 := []int{fi20, (fi20 << ((255) & 63)), fi20, 128}
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 8; fi29++ {
                        delete(v4, (fi20 | (fi29 >> ((49) & 63))))
                        var v30 int32 = (v2 & v2)
                        _ = fi29
                        _ = v30
                    }
                }
            }
            _ = fi20
        }
    case "\n\t":
        if false {
            var v31 int16 = 92
            v2 = (^(v2))
            for fi32 := 0; fi32 < 13; fi32++ {
                v5 += v5
                fmt.Println(v5, v5)
                _ = fi32
            }
            _ = v31
        } else {
            v33 := GenS1{F0: v5, F1: 24}
            _ = v33
        }
    }
    v34 := []int{128}
    var v35 uint32 = 0
    switch "foo bar" {
    case "a":
        var v36 int8 = 42
        _ = v36
    default:
        var v37 uint = 744
        _ = v37
    }
    switch ("a" + "\n\t") {
    case "foo bar":
        v38 := GenS2{F0: -1, F1: (^(v35))}
        _ = v38
    case "a":
        if true {
            var v39 uint8 = 1
            var v40 uint32 = v35
            if (v5 < (v5 >> ((916) & 63))) {
                _ = len(v34)
                var v41 uint64 = 750
                _ = v41
            } else {
                v39 += (^(v39))
                v42 := GenS4{F0: false, F1: -114, F2: (v39 >> ((65535) & 63))}
                var v43 int32 = v2
                _ = v42
                _ = v43
            }
            _ = v39
            _ = v40
        } else {
            var v44 float64 = float64(3.14159)
            _ = v44
        }
    }
    if true {
        fmt.Println(v35, v35)
        v34 = append(v34, 147)
    }
    switch "foo bar" {
    case "foo bar":
        v2 -= v2
    case "a":
        v35 -= 2147483648
    case "":
        v5 -= 42
    default:
        v45 := GenS2{F0: -9223372036854775808, F1: v35}
        _ = v45
    }
    if true {
        v46 := v3[(("hello" + ("foo bar" + ((("" + "a") + "\n\t") + ("a" + (("foo bar" + "foo bar") + ""))))) + "hello")]
        switch "foo bar" {
        case "foo bar":
            if false {
                var v47 uint16 = 562
                if true {
                    v47 = v47
                    v47 -= v47
                } else {
                    var v48 uint32 = (v35 ^ (^(v35)))
                    if false {
                        var v49 string = ("" + "hello")
                        _ = v49
                    }
                    var v50 float32 = 0.1
                    _ = v48
                    _ = v50
                }
                v35 -= (v35 ^ 1)
                _ = v47
            } else {
                if (((v35 - v35) < v35) || false) {
                    v51 := GenS5{F0: (v5 | v5), F1: 1}
                    delete(v1, v35)
                    _ = v51
                } else {
                    fmt.Println(v5)
                }
                if true {
                    switch "hello" {
                    case "a":
                        for fi52 := 0; fi52 < 14; fi52++ {
                            for fi53 := 0; fi53 < 1; fi53++ {
                                if (1e10 <= 1e10) {
                                    v54 := GenS2{F0: 492, F1: 779}
                                    _ = v54
                                }
                                delete(v3, "a")
                                var v55 int32 = v2
                                _ = fi53
                                _ = v55
                            }
                            v46 += (fi52 / ((fi52) | 1))
                            v56 := GenS2{F0: 9223372036854775807, F1: v35}
                            _ = fi52
                            _ = v56
                        }
                    default:
                        v34 = append(v34, (v46 / ((v46) | 1)))
                    }
                }
            }
        case "\n\t":
            if false {
                var v57 float64 = 1e10
                v5 = 4
                _ = v57
            }
        case "hello":
            var v58 float64 = 0.1
            _ = v58
        default:
            v46 += 0
        }
        switch 32767 {
        case 0:
            _ = len(v34)
        case 32768:
            for fi59 := 0; fi59 < 9; fi59++ {
                v60 := GenS2{F0: 715, F1: (v35 % ((v35) | 1))}
                _ = fi59
                _ = v60
            }
        }
        _ = v46
    } else {
        v61 := GenS2{F0: -1, F1: uint32(v5)}
        for fi62 := 0; fi62 < 6; fi62++ {
            for fi63 := 0; fi63 < 14; fi63++ {
                if (float32(1e10) < (1.5 + -2.25)) {
                    var v64 float32 = 1e10
                    _ = v64
                } else {
                    var v65 uint16 = 0
                    v35 -= 2147483647
                    _ = v65
                }
                if true {
                    v2 = (v2 << ((42) & 63))
                }
                _ = fi63
            }
            v2 += (v2 * v2)
            fmt.Println(v5, fi62, v5)
            _ = fi62
        }
        _ = v61
    }
    switch 9223372036854775807 {
    case -95:
        switch 0 {
        case 32768:
            if false {
                v66 := v1[(^(v35))]
                fmt.Println(v5)
                v35 += (v35 << ((65535) & 63))
                _ = v66
            } else {
                v5 += v5
            }
        }
    default:
        v35 -= v35
    }
    for fi67 := 0; fi67 < 14; fi67++ {
        var v68 int64 = -1
        v5 = (v5 + 291)
        _ = fi67
        _ = v68
    }
    v35 = (v35 + v35)
    v69 := GenS1{F0: 42, F1: 249}
    fmt.Println(v2, v35)
    v2 -= int32(v5)
    if false {
        if false {
            var v70 uint32 = v35
            v71 := GenS2{F0: -32, F1: (^(v35))}
            v72 := GenS6{F0: 1000, F1: 455}
            _ = v70
            _ = v71
            _ = v72
        }
    }
    switch -2147483648 {
    case 127:
        if ((v5 + (v5 * (^(v5)))) < (v5 / ((v5) | 1))) {
            switch ((("foo bar" + (("a" + "\n\t") + ("a" + "hello"))) + ("" + "hello")) + ("a" + ("hello" + "foo bar"))) {
            case "foo bar":
                if true {
                    v2 = -426
                    for fi73 := 0; fi73 < 7; fi73++ {
                        var v74 int64 = 1
                        for fi75 := 0; fi75 < 18; fi75++ {
                            var v76 int8 = 92
                            _ = fi75
                            _ = v76
                        }
                        switch (fi73 / (((fi73 * 32768)) | 1)) {
                        case 2147483647:
                            if len(v34) > 0 {
                                v34[0] = -629
                            }
                        case 1:
                            v4[fi73] = false
                        case 42:
                            var v77 float64 = -2.25
                            _ = v77
                        default:
                            v5 += v5
                        }
                        _ = fi73
                        _ = v74
                    }
                    v2 += (v2 % ((1) | 1))
                }
            case "hello":
                for fi78 := 0; fi78 < 12; fi78++ {
                    var v79 int16 = int16(v35)
                    v80 := v1[0]
                    var v81 int16 = (v79 | v79)
                    _ = fi78
                    _ = v79
                    _ = v80
                    _ = v81
                }
            default:
                if (true || true) {
                    if false {
                        v2 = v2
                        var v82 bool = true
                        _ = v82
                    } else {
                        for fi83 := 0; fi83 < 13; fi83++ {
                            fmt.Println(v2, v5, v2)
                            _ = fi83
                        }
                        for fi84 := 0; fi84 < 13; fi84++ {
                            v34 = append(v34, fi84)
                            v35 = (v35 / (((^(v35))) | 1))
                            if (!true) {
                                var v85 int32 = v2
                                var v86 uint32 = (^(v35))
                                v87 := GenS5{F0: (v5 * v5), F1: 1}
                                _ = v85
                                _ = v86
                                _ = v87
                            }
                            _ = fi84
                        }
                    }
                    v88 := GenS7{F0: v5, F1: 42, F2: ("\n\t" + "hello")}
                    _ = v88
                } else {
                    if (!true) {
                        v89 := v1[679]
                        if (32767 < 1000) {
                            if true {
                                v2 = v2
                                var v90 int16 = -32768
                                if len(v34) > 0 {
                                    v34[0] = (^(v89))
                                }
                                _ = v90
                            }
                            _ = len(v1)
                            if len(v34) >= 2 {
                                v34 = v34[1:len(v34)]
                            }
                        } else {
                            v91 := GenS8{F0: -543, F1: (v5 != 1), F2: true}
                            _ = v91
                        }
                        fmt.Println(v2, v2, v35)
                        _ = v89
                    } else {
                        v35 = (v35 / (((v35 | (v35 - 0))) | 1))
                        var v92 uint16 = 42
                        v93 := GenS5{F0: (v5 >> ((18446744073709551615) & 63)), F1: 68}
                        _ = v92
                        _ = v93
                    }
                    v94 := GenS4{F0: (!false), F1: 1, F2: 229}
                    var v95 uint16 = 2
                    _ = v94
                    _ = v95
                }
            }
            var v96 int = -60
            fmt.Println(v2, v35)
            _ = v96
        }
    default:
        for fi97 := 0; fi97 < 4; fi97++ {
            switch ("" + ((((("hello" + "") + "\n\t") + "a") + "hello") + "\n\t")) {
            case "":
                for fi98 := 0; fi98 < 12; fi98++ {
                    switch 399 {
                    case -960:
                        var v99 string = ("" + "a")
                        _ = v99
                    case -117:
                        v1[(v35 & (v35 ^ 65535))] = fi98
                    case 1:
                        v35 = (v35 & (v35 & 418))
                    default:
                        var v100 int = 32768
                        _ = v100
                    }
                    switch fi97 {
                    case 32767:
                        fmt.Println(fi98)
                    case 1:
                        fmt.Println(fi98)
                    }
                    v5 -= (v5 << ((0) & 63))
                    _ = fi98
                }
            case "\n\t":
                for fi101 := 0; fi101 < 4; fi101++ {
                    fmt.Println(fi97, v2)
                    v102 := GenS9{F0: v35, F1: v5, F2: (^(v35))}
                    if len(v34) >= 2 {
                        v34 = v34[1:len(v34)]
                    }
                    _ = fi101
                    _ = v102
                }
            }
            delete(v1, (v35 / (((v35 | v35)) | 1)))
            v2 += (-(v2))
            _ = fi97
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v69
}
