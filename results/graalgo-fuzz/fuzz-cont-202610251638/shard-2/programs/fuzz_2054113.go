package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint32; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 bool; F2 int }
type GenS5 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint; F1 uint64; F2 int64 }
type GenS7 struct { F0 string; F1 int8 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 uint32; F1 string; F2 string }
type GenS10 struct { F0 uint32; F1 bool }
type GenS11 struct { F0 uint64; F1 int; F2 int8 }
type GenS12 struct { F0 uint; F1 uint32 }
func genFn0(p0 int, p1 bool, p2 uint64) int16 {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 float64 = 0.1
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint32 = 407
            switch (p0 + fi7) {
            case 1:
                p0 -= 2147483647
            case -1:
                var v11 int32 = -2147483648
                _ = v11
            case -9223372036854775808:
                if p1 {
                    var v12 float64 = v8
                    p0 -= 999
                    _ = v12
                }
            default:
                p2 = p2
            }
            v10 = 4294967295
            if p1 {
                if p1 {
                    var v13 uint = 792
                    _ = v13
                } else {
                    if p1 {
                        var v14 uint64 = (p2 | p2)
                        v15 := GenS1{F0: (v10 >> ((857) & 63))}
                        var v16 uint32 = (v10 + (v10 / (((^(v10))) | 1)))
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    if (p1 && true) {
                        var v17 string = (("" + "a") + ("hello" + "foo bar"))
                        switch (v17 + (v17 + "\n\t")) {
                        case "a":
                            v18 := []int{32768, p0, fi7, 9223372036854775807}
                            _ = v18
                        case "hello":
                            v19 := GenS2{F0: 295, F1: 2, F2: true}
                            _ = v19
                        default:
                            v20 := GenS2{F0: -1, F1: v10, F2: (390 == 255)}
                            _ = v20
                        }
                        _ = v17
                    } else {
                        v8 += -2.25
                        var v21 uint64 = p2
                        _ = v21
                    }
                    if p1 {
                        switch "hello" {
                        case "":
                            if p1 {
                                if p1 {
                                    v10 -= v10
                                    var v22 bool = (!p1)
                                    v23 := []int{}
                                    _ = v22
                                    _ = v23
                                } else {
                                    var v24 float32 = (((-2.25 * -2.25) + 0.0) / -2.25)
                                    _ = v24
                                }
                                p1 = p1
                                var v25 uint64 = 42
                                _ = v25
                            } else {
                                p2 = 2
                            }
                        case "foo bar":
                            var v26 float64 = (0.1 / v8)
                            _ = v26
                        case "\n\t":
                            v27 := GenS3{F0: p0}
                            var v28 int32 = 2147483647
                            _ = v27
                            _ = v28
                        }
                    }
                }
            } else {
                v8 = float64(((0.0 * 0.0) - 3.14159))
                if p1 {
                    var v29 int32 = -1
                    var v30 uint16 = 1
                    _ = v29
                    _ = v30
                } else {
                    switch ("hello" + ("" + "a")) {
                    case "foo bar":
                        var v31 int = (fi9 | p0)
                        _ = v31
                    case "a":
                        var v32 int = fi7
                        _ = v32
                    }
                    switch "foo bar" {
                    case "\n\t":
                        v33 := GenS4{F0: 42, F1: (42 <= 9223372036854775807), F2: (-(p0))}
                        _ = v33
                    case "":
                        v34 := []int{(fi9 + (^(fi9)))}
                        _ = v34
                    default:
                        v10 = (^(v10))
                    }
                    var v35 uint64 = p2
                    _ = v35
                }
                if (((!p1) || p1) || p1) {
                    v36 := []int{(-(fi9)), fi9}
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    fmt.Println(p1, fi9)
                    _ = v36
                }
                if (0 >= 92) {
                    var v37 int = (p0 + p0)
                    _ = v37
                }
                fmt.Println(p1)
            }
            _ = fi9
            _ = v10
        }
        if (3.14159 < ((0.0 / ((v8 / (-2.25 - v8)) + 1.5)) + 1e10)) {
            fmt.Println(v8, p1)
            var v38 float32 = 0.1
            for fi39 := 0; fi39 < 8; fi39++ {
                p1 = p1
                if (!(!(!p1))) {
                    var v40 uint64 = p2
                    v8 -= float64(float32(v8))
                    _ = v40
                } else {
                    if true {
                        switch (^(fi39)) {
                        case 876:
                            var v41 string = ("a" + "a")
                            _ = v41
                        case 44:
                            fmt.Println(fi39, p1)
                        }
                        if p1 {
                            var v42 int64 = -1
                            v38 += (v38 + -2.25)
                            v43 := GenS1{F0: 0}
                            _ = v42
                            _ = v43
                        } else {
                            p2 -= (p2 & (p2 * (p2 * (p2 * p2))))
                            v44 := GenS1{F0: 4294967295}
                            v45 := []int{2147483647}
                            _ = v44
                            _ = v45
                        }
                        var v46 int16 = 42
                        var v47 uint64 = 292
                        _ = v46
                        _ = v47
                    } else {
                        var v48 float64 = v8
                        v49 := []int{(fi39 * fi7), fi7, fi7}
                        if len(v49) > 0 {
                            v49[0] = p0
                        }
                        v50 := GenS5{F0: p2, F1: 2147483648, F2: 1}
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                    v8 -= 3.14159
                    v51 := GenS6{F0: 4294967295, F1: (p2 >> ((2) & 63)), F2: -9223372036854775808}
                    _ = v51
                }
                _ = fi39
            }
            _ = v38
        } else {
            v52 := []int{fi7}
            for fi53 := 0; fi53 < 19; fi53++ {
                v54 := GenS7{F0: (("a" + "foo bar") + ""), F1: -2}
                var v55 uint = 1
                var v56 uint = v55
                var v57 int = fi53
                _ = fi53
                _ = v54
                _ = v55
                _ = v56
                _ = v57
            }
            var v58 int32 = -182
            _ = v52
            _ = v58
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p2, p0)
    var v59 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v59
    return 1
}
func genFn1(p0 int64, p1 int64) uint8 {
    if ((42 >= 4294967295) && true) {
        for fi60 := 0; fi60 < 9; fi60++ {
            v61 := []int{-173}
            _ = fi60
            _ = v61
        }
        switch -766 {
        case -2147483648:
            switch "a" {
            case "":
                if true {
                    v62 := []int{32768, -9223372036854775808}
                    _ = v62
                }
            case "\n\t":
                p0 += 547
            case "hello":
                for fi63 := 0; fi63 < 10; fi63++ {
                    v64 := GenS8{F0: "\n\t", F1: (p1 - (p0 - 9223372036854775807))}
                    for fi65 := 0; fi65 < 20; fi65++ {
                        var v66 uint8 = 167
                        v67 := GenS1{F0: 396}
                        v68 := GenS9{F0: 0, F1: "hello", F2: "hello"}
                        _ = fi65
                        _ = v66
                        _ = v67
                        _ = v68
                    }
                    _ = fi63
                    _ = v64
                }
            }
        case -119:
            switch ("hello" + ("a" + "hello")) {
            case "\n\t":
                fmt.Println(p0, p1)
            case "hello":
                fmt.Println(p0, p1)
            case "a":
                switch "foo bar" {
                case "a":
                    var v69 string = "a"
                    _ = v69
                case "\n\t":
                    v70 := GenS10{F0: 4294967295, F1: false}
                    _ = v70
                default:
                    var v71 uint16 = 1
                    _ = v71
                }
            }
        case 4294967295:
            v72 := GenS3{F0: 1}
            _ = v72
        }
        var v73 float64 = 1e10
        _ = v73
    }
    var v74 int8 = 0
    if (211 <= 65535) {
        if true {
            p1 = (p1 % ((1) | 1))
        }
        switch -241 {
        case 32767:
            if true {
                var v75 uint8 = 255
                _ = v75
            }
        }
        var v76 int32 = -42
        _ = v76
    }
    _ = p0
    _ = p1
    _ = v74
    return 0
}

func main() {
    v1 := []uint32{2, 42}
    v2 := make(map[int]int)
    v3 := []string{"foo bar", "foo bar", "", "hello"}
    var v4 int64 = 9223372036854775807
    var v5 bool = (true || (1e10 < 0.0))
    var v6 uint = 0
    if true {
        fmt.Println(v6, v6)
    }
    for fi77 := 0; fi77 < 3; fi77++ {
        if ("a" < "hello") {
            for fi78 := 0; fi78 < 7; fi78++ {
                if len(v1) > 0 {
                    v1[0] = 1
                }
                v2[(fi77 + -2147483648)] = fi78
                _ = fi78
            }
            v4 -= v4
        }
        _ = fi77
    }
    var v79 uint8 = 1
    fmt.Println(v5, v6)
    v6 += v6
    fmt.Println(v4)
    var v80 int = 32767
    for fi81 := 0; fi81 < 19; fi81++ {
        fmt.Println(v79, v4)
        switch "a" {
        case "a":
            fmt.Println(v4, v5)
        case "\n\t":
            v4 -= int64(fi81)
        default:
            var v82 float32 = 1e10
            _ = v82
        }
        _ = fi81
    }
    for fi83 := 0; fi83 < 7; fi83++ {
        for fi84 := 0; fi84 < 11; fi84++ {
            if (!((0 <= 32767) && false)) {
                switch 331 {
                case 110:
                    fmt.Println(v6)
                case 0:
                    v79 = v79
                default:
                    if (0.0 <= 1.5) {
                        v85 := GenS11{F0: 701, F1: fi83, F2: 0}
                        _ = v85
                    }
                }
            } else {
                if ((v79 - genFn1(988, 42)) < (v79 - 128)) {
                    var v86 uint = 78
                    _ = v86
                } else {
                    switch "\n\t" {
                    case "hello":
                        v80 -= v80
                    }
                }
            }
            _ = fi84
        }
        v5 = v5
        v79 = (v79 ^ v79)
        _ = fi83
    }
    v87 := GenS12{F0: 42, F1: 2147483647}
    var v88 string = "a"
    switch 128 {
    case 42:
        _ = len(v2)
    case 9223372036854775807:
        var v89 string = v88
        _ = v89
    case 393:
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
    }
    var v90 int16 = 0
    delete(v2, v80)
    for fi91 := 0; fi91 < 2; fi91++ {
        if v5 {
            _ = len(v1)
        }
        _ = fi91
    }
    var v92 uint32 = 604
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v88:", v88)
    fmt.Println("v90:", v90)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
    _ = v80
    _ = v87
    _ = v88
    _ = v90
    _ = v92
}
