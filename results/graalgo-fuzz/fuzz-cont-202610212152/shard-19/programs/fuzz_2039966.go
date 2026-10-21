package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16 }
func genFn0(p0 int64) bool {
    p0 += p0
    if false {
        if false {
            if true {
                p0 = p0
                v6 := GenS1{F0: 32767, F1: 1, F2: 18446744073709551615}
                _ = v6
            }
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 uint = 4294967295
                for fi9 := 0; fi9 < 18; fi9++ {
                    fmt.Println(p0, v8, fi7)
                    _ = fi9
                }
                _ = fi7
                _ = v8
            }
        } else {
            v10 := GenS2{F0: 345, F1: 0}
            switch "\n\t" {
            case "\n\t":
                p0 -= p0
            default:
                for fi11 := 0; fi11 < 4; fi11++ {
                    fmt.Println(p0, fi11)
                    switch (("\n\t" + "") + "a") {
                    case "":
                        var v12 uint = 405
                        _ = v12
                    case "foo bar":
                        if (p0 > p0) {
                            var v13 uint = 1
                            _ = v13
                        } else {
                            for fi14 := 0; fi14 < 14; fi14++ {
                                switch "foo bar" {
                                case "":
                                    p0 -= 1
                                case "a":
                                    v15 := []int{2147483647, fi14, (fi14 + fi11)}
                                    _ = v15
                                case "foo bar":
                                    v16 := GenS1{F0: 0, F1: 1, F2: 1}
                                    _ = v16
                                default:
                                    v17 := []int{4294967295, fi11, 32767}
                                    _ = v17
                                }
                                var v18 uint8 = 226
                                if true {
                                    fmt.Println(v18)
                                }
                                var v19 int16 = 32767
                                _ = fi14
                                _ = v18
                                _ = v19
                            }
                        }
                    }
                    for fi20 := 0; fi20 < 5; fi20++ {
                        if false {
                            fmt.Println(p0)
                            var v21 int8 = 65
                            _ = v21
                        }
                        if (!true) {
                            v22 := []int{32768, -647, (fi20 % ((fi11) | 1)), fi20}
                            p0 -= p0
                            switch (-(fi20)) {
                            case 127:
                                p0 += p0
                            case 39:
                                fmt.Println(p0)
                                p0 = 130
                            default:
                                var v23 uint32 = 65535
                                _ = v23
                            }
                            _ = v22
                        }
                        _ = fi20
                    }
                    _ = fi11
                }
            }
            _ = v10
        }
        p0 = (p0 ^ 9223372036854775807)
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        for fi25 := 0; fi25 < 9; fi25++ {
            v26 := GenS1{F0: 65535, F1: 101, F2: 855}
            v27 := []int{fi24, fi25}
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi24
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 uint64 = 1
        v30 := []int{}
        for fi31 := 0; fi31 < 1; fi31++ {
            switch fi28 {
            case -529:
                fmt.Println(fi31, fi31)
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
                if (!true) {
                    if true {
                        var v32 int32 = -1
                        var v33 int64 = p0
                        _ = v32
                        _ = v33
                    } else {
                        switch fi31 {
                        case 32768:
                            if false {
                                p0 -= (p0 - p0)
                            }
                        case -9223372036854775808:
                            var v34 float32 = (0.1 - 0.0)
                            _ = v34
                        case -288:
                            switch fi28 {
                            case 42:
                                p0 += (p0 & (p0 | (^(p0))))
                            }
                        default:
                            v35 := GenS2{F0: 252, F1: 255}
                            _ = v35
                        }
                        var v36 int16 = 32767
                        _ = v36
                    }
                    var v37 int8 = 0
                    _ = v37
                } else {
                    v29 += (v29 & v29)
                    fmt.Println(v29, fi31)
                }
            case 4294967295:
                if (-128 == 100) {
                    p0 = -1
                    fmt.Println(fi31)
                    fmt.Println(fi31, fi28, p0)
                } else {
                    if true {
                        v38 := GenS3{F0: 94, F1: 364}
                        v39 := GenS3{F0: 42, F1: 0}
                        var v40 uint = uint(v29)
                        _ = v38
                        _ = v39
                        _ = v40
                    } else {
                        v41 := GenS3{F0: 42, F1: 2147483648}
                        v29 += (^(v29))
                        fmt.Println(fi28, fi31, p0)
                        fmt.Println(v29, fi28)
                        v42 := GenS1{F0: 342, F1: -2147483648, F2: 0}
                        v43 := GenS3{F0: 42, F1: 42}
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    var v44 uint = 42
                    _ = v44
                }
            case 2147483647:
                p0 = (p0 & (p0 & p0))
            default:
                _ = len(v30)
            }
            var v45 float32 = -2.25
            _ = fi31
            _ = v45
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    _ = p0
    return (3.14159 != ((1.5 - (1e10 / (-2.25 / 3.14159))) + (0.1 / float64((1.5 * 0.1)))))
}

func main() {
    v1 := []string{}
    var v2 uint32 = 1
    v3 := []string{"hello", "a", ""}
    v4 := make(map[int]int)
    var v5 float32 = 0.1
    switch 173 {
    case 42:
        for fi46 := 0; fi46 < 4; fi46++ {
            v2 = 2
            _ = fi46
        }
    case 2147483647:
        switch "a" {
        case "hello":
            var v47 int64 = -9223372036854775808
            _ = v47
        case "":
            if false {
                for fi48 := 0; fi48 < 16; fi48++ {
                    v2 -= (v2 & v2)
                    fmt.Println(v2)
                    v49 := GenS4{F0: 4294967295, F1: 4294967295}
                    _ = fi48
                    _ = v49
                }
                if true {
                    fmt.Println(v5, v2)
                    fmt.Println(v2)
                } else {
                    v50 := GenS2{F0: 2147483647, F1: 11}
                    _ = v50
                }
            }
        case "\n\t":
            if false {
                fmt.Println(v5)
                _ = len(v4)
                _ = len(v4)
            } else {
                if false {
                    v5 = v5
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v2, v2)
                    case "":
                        fmt.Println(v2)
                    case "hello":
                        v51 := v4[42]
                        _ = v51
                    default:
                        v52 := GenS4{F0: 2, F1: 434}
                        _ = v52
                    }
                } else {
                    var v53 uint64 = 514
                    _ = v53
                }
                var v54 int32 = 0
                _ = len(v4)
                _ = v54
            }
        }
    case 32767:
        var v55 uint16 = 42
        _ = v55
    }
    if ("" == "\n\t") {
        var v56 uint64 = 652
        _ = v56
    }
    for fi57 := 0; fi57 < 5; fi57++ {
        if (!true) {
            var v58 int64 = 42
            switch "a" {
            case "\n\t":
                _ = len(v4)
            case "a":
                for fi59 := 0; fi59 < 9; fi59++ {
                    v1 = append(v1, "\n\t")
                    delete(v4, (-(fi57)))
                    _ = fi59
                }
            case "":
                switch "\n\t" {
                case "a":
                    switch ("hello" + ("foo bar" + "a")) {
                    case "\n\t":
                        var v60 uint32 = v2
                        _ = v60
                    case "a":
                        var v61 string = ("hello" + "")
                        _ = v61
                    case "":
                        if false {
                            var v62 uint8 = 2
                            _ = v62
                        } else {
                            v5 += v5
                        }
                    }
                default:
                    if true {
                        var v63 bool = true
                        v4[(fi57 + (-(fi57)))] = (fi57 << ((236) & 63))
                        v5 += (v5 / 0.1)
                        _ = v63
                    } else {
                        v5 -= v5
                        v5 = (3.14159 + v5)
                    }
                }
            default:
                if ("\n\t" >= ("foo bar" + "hello")) {
                    var v64 int16 = 0
                    _ = v64
                } else {
                    if true {
                        v58 = -206
                        v5 -= v5
                        for fi65 := 0; fi65 < 7; fi65++ {
                            v5 -= v5
                            _ = fi65
                        }
                    } else {
                        for fi66 := 0; fi66 < 17; fi66++ {
                            v2 = v2
                            var v67 uint = 728
                            fmt.Println(v58)
                            _ = fi66
                            _ = v67
                        }
                        v2 += (^(v2))
                    }
                    v5 -= (0.0 / ((v5 - (3.14159 - 1.5)) - 0.1))
                    var v68 int32 = 2147483647
                    _ = v68
                }
            }
            if false {
                var v69 int16 = 648
                v58 = int64(v69)
                var v70 string = "hello"
                _ = v69
                _ = v70
            } else {
                for fi71 := 0; fi71 < 16; fi71++ {
                    v72 := GenS4{F0: 0, F1: 0}
                    _ = fi71
                    _ = v72
                }
                var v73 string = ("\n\t" + ("\n\t" + "foo bar"))
                fmt.Println(v5, fi57, v2)
                _ = v73
            }
            _ = v58
        }
        for fi74 := 0; fi74 < 14; fi74++ {
            fmt.Println(v2)
            if true {
                var v75 string = ("a" + "a")
                _ = v75
            }
            if false {
                _ = len(v1)
            }
            _ = fi74
        }
        _ = fi57
    }
    if false {
        for fi76 := 0; fi76 < 17; fi76++ {
            switch (fi76 % ((fi76) | 1)) {
            case 32767:
                v77 := GenS3{F0: -14, F1: (^(v2))}
                _ = v77
            case 9223372036854775807:
                var v78 uint = 797
                _ = v78
            case -838:
                _ = len(v1)
            }
            v2 += (v2 % ((2147483648) | 1))
            _ = fi76
        }
    }
    switch 42 {
    case 32767:
        for fi79 := 0; fi79 < 15; fi79++ {
            switch (("" + "hello") + "\n\t") {
            case "a":
                v2 = (v2 ^ 65535)
            case "foo bar":
                if true {
                    v5 -= -2.25
                    v5 = -2.25
                    fmt.Println(v2, v5)
                }
            default:
                var v80 uint32 = v2
                _ = v80
            }
            v5 = ((v5 - (v5 + 1e10)) + v5)
            _ = fi79
        }
    }
    if (true && false) {
        v2 += 725
    }
    v1 = append(v1, "\n\t")
    switch "hello" {
    case "hello":
        if ((float32(float64(0.0)) * (1e10 * 0.1)) < -2.25) {
            if (((v5 / v5) * 1.5) == 1.5) {
                var v81 int8 = -106
                v5 -= 0.0
                _ = v81
            } else {
                v82 := GenS2{F0: -2147483648, F1: 42}
                if false {
                    var v83 bool = false
                    _ = v83
                }
                if false {
                    v5 = v5
                } else {
                    var v84 float64 = ((1e10 + 1e10) * (0.1 + 1e10))
                    var v85 float32 = float32((((v84 / ((v84 * v84) - v84)) - ((v84 - ((-2.25 + 1.5) - 0.0)) * v84)) + ((0.1 / 3.14159) + 1e10)))
                    _ = len(v3)
                    _ = v84
                    _ = v85
                }
                _ = v82
            }
            v86 := GenS5{F0: 100}
            var v87 int8 = 0
            _ = v86
            _ = v87
        } else {
            v2 -= (v2 * v2)
            v88 := GenS6{F0: -32768}
            _ = v88
        }
    case "a":
        var v89 float32 = ((float32(float64(v5)) / v5) / (float32(3.14159) + ((1e10 + 0.0) * 0.1)))
        _ = v89
    case "":
        var v90 int64 = 9223372036854775807
        _ = v90
    default:
        var v91 uint16 = 269
        _ = v91
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
