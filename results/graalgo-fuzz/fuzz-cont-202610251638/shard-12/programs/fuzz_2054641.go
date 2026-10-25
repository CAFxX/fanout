package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint16; F2 int }
type GenS6 struct { F0 int; F1 uint; F2 int }
type GenS7 struct { F0 int64; F1 uint; F2 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16; F1 uint8 }
func genFn0(p0 bool, p1 int8) int8 {
    p1 += 0
    var v6 int8 = (p1 - (^(p1)))
    switch 2147483647 {
    case 128:
        if p0 {
            v7 := GenS1{F0: p0, F1: 42}
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 uint16 = 2
                v10 := []int{}
                var v11 uint16 = v9
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            v6 = (p1 - 127)
            _ = v7
        }
    case 0:
        switch ("\n\t" + "") {
        case "a":
            var v12 int = 4294967295
            _ = v12
        }
    case -1:
        switch -2147483648 {
        case 917:
            v13 := []int{127, -2147483648, -124, 32768}
            _ = v13
        default:
            fmt.Println(p1)
        }
    }
    if (p0 || p0) {
        var v14 uint16 = 866
        v15 := []int{-1, 32768}
        _ = v14
        _ = v15
    } else {
        switch "hello" {
        case "\n\t":
            if false {
                var v16 float64 = 1e10
                _ = v16
            } else {
                var v17 bool = true
                _ = v17
            }
        default:
            v6 = (-(p1))
        }
        if p0 {
            for fi18 := 0; fi18 < 4; fi18++ {
                v19 := GenS1{F0: (182 <= 0), F1: 42}
                fmt.Println(p1, fi18)
                _ = fi18
                _ = v19
            }
            v20 := []int{}
            switch 32768 {
            case 42:
                fmt.Println(v6, p1, v6)
            case -9223372036854775808:
                for fi21 := 0; fi21 < 19; fi21++ {
                    fmt.Println(p1, fi21, fi21)
                    for fi22 := 0; fi22 < 19; fi22++ {
                        p1 -= v6
                        _ = fi22
                    }
                    _ = fi21
                }
            case 4294967295:
                p0 = (255 > 128)
            default:
                v23 := GenS1{F0: (4294967295 != 81), F1: 0}
                _ = v23
            }
            _ = v20
        }
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            p0 = true
            if (1 != 42) {
                v26 := []int{(fi25 % ((936) | 1)), (fi25 / (((^(fi24))) | 1)), fi24, fi25}
                v27 := GenS1{F0: p0, F1: 127}
                _ = v26
                _ = v27
            } else {
                var v28 uint16 = 1
                v29 := GenS2{F0: 42, F1: 448}
                var v30 int16 = 1000
                var v31 int64 = 9223372036854775807
                var v32 float32 = 1.5
                _ = v28
                _ = v29
                _ = v30
                _ = v31
                _ = v32
            }
            if p0 {
                p0 = p0
            }
            _ = fi25
        }
        var v33 int8 = -128
        p0 = ((!(!(p0 || (!(p0 && true))))) && true)
        _ = fi24
        _ = v33
    }
    _ = p0
    _ = p1
    _ = v6
    return (p1 % (((v6 >> ((2) & 63))) | 1))
}
func genFn1() uint16 {
    v34 := GenS2{F0: 136, F1: 42}
    var v35 uint32 = 65535
    v35 += (v35 ^ v35)
    _ = v34
    _ = v35
    return 42
}

func main() {
    v1 := []uint32{65535, 1, 990}
    var v2 bool = (true || false)
    var v3 int32 = -221
    var v4 uint8 = 255
    var v5 uint8 = 0
    v5 += v4
    var v36 uint64 = 18446744073709551615
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v2, v5)
    default:
        if true {
            if (2147483647 >= 663) {
                v37 := GenS1{F0: false, F1: v5}
                for fi38 := 0; fi38 < 14; fi38++ {
                    for fi39 := 0; fi39 < 18; fi39++ {
                        v5 = v4
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v1 = append(v1, 0)
                        _ = fi39
                    }
                    var v40 uint16 = 42
                    switch (-(fi38)) {
                    case 42:
                        if (!(!v2)) {
                            var v41 float64 = 1e10
                            _ = v41
                        } else {
                            v42 := GenS3{F0: 0, F1: (0 != 1)}
                            v43 := GenS4{F0: "\n\t", F1: (v40 & (v40 | (v40 << ((1) & 63))))}
                            for fi44 := 0; fi44 < 7; fi44++ {
                                if (v36 > (v36 << ((255) & 63))) {
                                    if len(v1) > 0 {
                                        v1[0] = 2147483648
                                    }
                                    var v45 uint8 = v4
                                    _ = v45
                                } else {
                                    if false {
                                        var v46 uint32 = 686
                                        _ = v46
                                    } else {
                                        _ = len(v1)
                                    }
                                    fmt.Println(v2, fi38, v2)
                                }
                                var v47 uint8 = 70
                                switch "" {
                                case "":
                                    v5 = v5
                                case "a":
                                    var v48 uint32 = 614
                                    _ = v48
                                case "\n\t":
                                    var v49 int64 = 451
                                    _ = v49
                                default:
                                    if v2 {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v50 := GenS5{F0: 127, F1: (^(v40)), F2: 542}
                                        _ = v50
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            if v2 {
                                                _ = len(v1)
                                                v36 = (v36 | v36)
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if (-124 >= -128) {
                                                    v47 += (v4 - 128)
                                                    v3 -= int32(v47)
                                                } else {
                                                    var v51 int32 = int32(v40)
                                                    _ = v51
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = fi44
                                _ = v47
                            }
                            _ = v42
                            _ = v43
                        }
                    case 9223372036854775807:
                        if len(v1) > 0 {
                            v1[0] = 65535
                        }
                    }
                    _ = fi38
                    _ = v40
                }
                if (!(9223372036854775807 > 9223372036854775807)) {
                    v4 -= (v5 & v5)
                    if false {
                        switch 9223372036854775807 {
                        case 778:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v52 := GenS6{F0: -30, F1: 812, F2: 32768}
                        _ = v52
                    }
                    var v53 int8 = 0
                    _ = v53
                }
                _ = v37
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v54 string = "hello"
                _ = v54
            }
            if v2 {
                v3 += v3
            }
        }
    }
    var v55 int32 = (v3 + 42)
    switch 127 {
    case 32767:
        v2 = (!((v36 > v36) || v2))
    case -355:
        if true {
            v5 += 128
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
        } else {
            if v2 {
                var v56 int32 = v55
                _ = v56
            }
            if (9223372036854775807 <= -1) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v5 = v4
                v3 -= (-(v55))
            }
        }
    }
    v57 := GenS2{F0: 1, F1: 1000}
    v55 += (v3 / ((int32(v4)) | 1))
    v58 := GenS4{F0: (("hello" + ("hello" + "\n\t")) + "\n\t"), F1: 32767}
    for fi59 := 0; fi59 < 1; fi59++ {
        var v60 float32 = 3.14159
        if ((v55 % (((-(v55))) | 1)) > 16) {
            for fi61 := 0; fi61 < 5; fi61++ {
                if v2 {
                    _ = len(v1)
                    fmt.Println(v36, v36)
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                }
                if v2 {
                    fmt.Println(v5)
                    if v2 {
                        switch fi61 {
                        case -1:
                            fmt.Println(v5)
                        case 4294967295:
                            v2 = (1.5 > ((3.14159 + -2.25) / (float64(v60) / 1e10)))
                        case 0:
                            fmt.Println(fi59)
                        default:
                            fmt.Println(v4, v60, fi59)
                        }
                    }
                } else {
                    if ((0.0 + (1.5 - (v60 - (v60 / (v60 * (v60 / v60)))))) != float32(1e10)) {
                        if v2 {
                            if (42 == -716) {
                                v4 = v5
                                if ((fi59 / ((fi59) | 1)) == -2147483648) {
                                    v55 += -2147483648
                                    v60 = ((v60 / (-2.25 + v60)) * float32((1.5 / 3.14159)))
                                }
                            } else {
                                v1 = append(v1, 4294967295)
                                if (!("" <= ("a" + (("\n\t" + "hello") + ("foo bar" + ("hello" + ("" + "a"))))))) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    _ = len(v1)
                                }
                                v2 = false
                                var v62 int8 = -128
                                _ = v62
                            }
                        } else {
                            var v63 uint = 918
                            var v64 int = (fi59 & (fi59 ^ (fi59 % ((fi59) | 1))))
                            var v65 int16 = 42
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                        if v2 {
                            v1 = append(v1, 2147483647)
                        } else {
                            var v66 float32 = ((v60 - v60) - (v60 - v60))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if (3.14159 >= 0.0) {
                                v4 += v5
                                v67 := GenS2{F0: 65535, F1: 0}
                                v2 = v2
                                _ = v67
                            } else {
                                fmt.Println(v2, v4, v55)
                                fmt.Println(v55, fi61, v60)
                            }
                            _ = v66
                        }
                        if false {
                            v68 := GenS7{F0: 722, F1: 1, F2: fi59}
                            var v69 float64 = ((1.5 + 1e10) + 1.5)
                            v3 = v55
                            _ = v68
                            _ = v69
                        } else {
                            var v70 string = "hello"
                            _ = v70
                        }
                    } else {
                        var v71 int8 = genFn0(true, 127)
                        _ = v71
                    }
                    v36 += 42
                }
                var v72 int64 = -1
                v5 = (v5 & (v4 >> ((2) & 63)))
                v5 += (v4 >> ((255) & 63))
                _ = fi61
                _ = v72
            }
            for fi73 := 0; fi73 < 12; fi73++ {
                switch "\n\t" {
                case "\n\t":
                    if false {
                        var v74 uint32 = 451
                        var v75 int = 127
                        var v76 bool = ((v2 || true) && v2)
                        _ = v74
                        _ = v75
                        _ = v76
                    } else {
                        v55 -= (v55 >> ((255) & 63))
                        var v77 float64 = 1e10
                        if len(v1) > 0 {
                            v1[0] = 2147483648
                        }
                        v78 := GenS8{F0: 9223372036854775807}
                        switch (fi73 - -590) {
                        case -1:
                            fmt.Println(v4, v36)
                        }
                        switch (fi73 * fi73) {
                        case -9223372036854775808:
                            fmt.Println(v60, v36, fi59)
                        }
                        var v79 uint = 49
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if false {
                            var v80 int8 = genFn0((301 > 0), 58)
                            var v81 int64 = 1
                            if false {
                                fmt.Println(v60, v5)
                                v1 = append(v1, 1)
                            } else {
                                if (false && true) {
                                    v81 = int64(v3)
                                    v55 = v55
                                    var v82 uint8 = (v5 % (((v4 & 127)) | 1))
                                    _ = v82
                                } else {
                                    var v83 int64 = (v81 ^ v81)
                                    var v84 uint = (v79 / (((v79 ^ v79)) | 1))
                                    _ = v83
                                    _ = v84
                                }
                                v85 := GenS9{F0: 778, F1: 95}
                                _ = v85
                            }
                            _ = v80
                            _ = v81
                        }
                        if true {
                            if v2 {
                                if true {
                                    fmt.Println(fi59, v2, v3)
                                }
                                switch ("a" + "\n\t") {
                                case "hello":
                                    switch "a" {
                                    case "hello":
                                        var v86 uint32 = 2147483647
                                        _ = v86
                                    case "\n\t":
                                        v3 = v55
                                    case "foo bar":
                                        _ = len(v1)
                                    }
                                }
                                var v87 uint64 = v36
                                _ = v87
                            }
                            fmt.Println(v55, v55)
                            v77 -= 3.14159
                            if (0 != genFn0(true, genFn0((v55 <= (v3 % ((v3) | 1))), -128))) {
                                fmt.Println(fi59, v3, v4)
                                var v88 int8 = 0
                                var v89 int64 = -9223372036854775808
                                _ = v88
                                _ = v89
                            } else {
                                if ("\n\t" != "a") {
                                    _ = len(v1)
                                }
                            }
                        } else {
                            fmt.Println(v60)
                            switch (fi59 << ((v79) & 63)) {
                            case 127:
                                _ = len(v1)
                                v77 = (v77 * (((v77 / v77) + v77) * v77))
                                v77 += 1e10
                                if (!v2) {
                                    if (v2 && v2) {
                                        fmt.Println(v77, v77, fi73)
                                        fmt.Println(v55, v2)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    v1 = append(v1, 4294967295)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                var v90 uint64 = v36
                                _ = v90
                            case 1:
                                v5 = v5
                            default:
                                var v91 float64 = -2.25
                                _ = v91
                            }
                            v92 := GenS1{F0: false, F1: 2}
                            _ = v92
                        }
                        if v2 {
                            v55 += 1
                            v93 := GenS9{F0: 42, F1: v5}
                            var v94 int32 = (v55 - (v3 ^ (-(v3))))
                            _ = v93
                            _ = v94
                        }
                        _ = v77
                        _ = v78
                        _ = v79
                    }
                case "hello":
                    var v95 int16 = 1000
                    _ = v95
                case "foo bar":
                    if (v2 && ((v5 <= 157) && v2)) {
                        v96 := GenS5{F0: 107, F1: 1, F2: fi59}
                        v55 -= v3
                        v55 = 42
                        _ = v96
                    }
                }
                v1 = append(v1, uint32(v4))
                _ = fi73
            }
            if v2 {
                fmt.Println(v5)
            }
        }
        v3 -= (v3 << ((18446744073709551615) & 63))
        _ = fi59
        _ = v60
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v55
    _ = v57
    _ = v58
}
