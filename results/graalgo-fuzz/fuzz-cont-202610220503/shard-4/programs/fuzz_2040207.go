package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8; F1 int32 }
func genFn0() int64 {
    if (149 > 274) {
        v7 := []int{}
        _ = v7
    } else {
        switch -1 {
        case 9223372036854775807:
            if ((false || false) || false) {
                var v8 uint16 = 42
                _ = v8
            }
        }
        var v9 bool = (127 != 222)
        var v10 uint = 18446744073709551615
        _ = v10
        _ = v9
    }
    var v11 int32 = 0
    var v12 uint = 255
    if true {
        var v13 int64 = -9223372036854775808
        _ = v13
    } else {
        if false {
            for fi14 := 0; fi14 < 3; fi14++ {
                if false {
                    v15 := []int{127, 127, (fi14 >> ((v12) & 63))}
                    _ = v15
                }
                var v16 float32 = (1.5 / 3.14159)
                _ = fi14
                _ = v16
            }
            switch "hello" {
            case "foo bar":
                var v17 float32 = (0.0 - ((1e10 + 0.1) + -2.25))
                _ = v17
            case "hello":
                v18 := GenS1{F0: 793, F1: 1}
                _ = v18
            default:
                var v19 int8 = int8(v11)
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 14; fi20++ {
                v21 := GenS1{F0: 951, F1: 42}
                for fi22 := 0; fi22 < 15; fi22++ {
                    if (fi22 == fi20) {
                        v12 = (v12 & 4294967295)
                    } else {
                        var v23 float32 = (1.5 - 0.0)
                        _ = v23
                    }
                    if true {
                        switch "" {
                        case "foo bar":
                            var v24 int64 = -632
                            _ = v24
                        default:
                            var v25 int8 = 42
                            _ = v25
                        }
                        v12 += v12
                    }
                    v26 := GenS2{F0: 1}
                    var v27 bool = false
                    _ = fi22
                    _ = v26
                    _ = v27
                }
                v12 += 0
                _ = fi20
                _ = v21
            }
        }
        var v28 string = "foo bar"
        _ = v28
    }
    switch 4294967295 {
    case 42:
        v29 := []int{32767, -9223372036854775808}
        _ = v29
    case -9223372036854775808:
        v30 := GenS3{F0: v11, F1: (((!true) || true) && (0 >= 0))}
        _ = v30
    case -1:
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v12, v12)
        default:
            v31 := GenS3{F0: (v11 + v11), F1: ((!false) && (733 == 4294967295))}
            _ = v31
        }
    }
    _ = v11
    _ = v12
    return -1
}
func genFn1(p0 uint32, p1 uint8, p2 int) int {
    v32 := []int{249, (p2 / ((int(p0)) | 1)), p2, p2}
    p0 = p0
    for fi33 := 0; fi33 < 19; fi33++ {
        var v34 uint32 = 294
        _ = len(v32)
        _ = fi33
        _ = v34
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return p2
}

func main() {
    var v1 int16 = 1
    var v2 uint16 = 65535
    var v3 uint16 = v2
    var v4 uint = 233
    v5 := []int{696}
    var v6 uint16 = v2
    var v35 bool = false
    switch "a" {
    case "hello":
        switch 1 {
        case 4294967295:
            v2 += v2
        case -1:
            fmt.Println(v35, v2)
        }
    default:
        fmt.Println(v6, v4)
    }
    if v35 {
        v4 += (^(v4))
    } else {
        if len(v5) > 0 {
            v5[0] = 29
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v36 := GenS3{F0: 42, F1: true}
        _ = v36
    }
    if v35 {
        switch "\n\t" {
        case "a":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "":
            for fi37 := 0; fi37 < 18; fi37++ {
                if ((v4 * v4) <= v4) {
                    for fi38 := 0; fi38 < 3; fi38++ {
                        var v39 bool = false
                        v40 := GenS4{F0: (v3 % ((1) | 1)), F1: 65535, F2: 2147483647}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                }
                _ = fi37
            }
        case "foo bar":
            var v41 uint = (^(v4))
            _ = v41
        }
        switch "a" {
        case "foo bar":
            if (42 < 65535) {
                var v42 int8 = 74
                for fi43 := 0; fi43 < 14; fi43++ {
                    v5 = append(v5, genFn1(1, 0, int(v1)))
                    _ = fi43
                }
                _ = v42
            }
        case "a":
            v44 := GenS5{F0: 1, F1: 1, F2: 1}
            _ = v44
        }
    } else {
        v35 = v35
        if len(v5) > 0 {
            v5[0] = -2147483648
        }
        var v45 float32 = (0.1 - (((1e10 * (1.5 + float32((-2.25 + (-2.25 / -2.25))))) * 1e10) - 0.0))
        _ = v45
    }
    v3 = (v3 / ((1) | 1))
    v2 += uint16(v4)
    v3 -= (v6 ^ (v2 ^ v2))
    switch "a" {
    case "a":
        switch 128 {
        case 32767:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case 1:
            if true {
                fmt.Println(v1, v35, v35)
                v1 = (v1 | (v1 - 1))
            } else {
                if ((^(v1)) == v1) {
                    v46 := GenS2{F0: 0}
                    var v47 int = 32768
                    switch v47 {
                    case 0:
                        fmt.Println(v4, v6)
                    case 128:
                        v48 := GenS5{F0: 0, F1: 0, F2: 1}
                        _ = v48
                    }
                    _ = v46
                    _ = v47
                } else {
                    var v49 int64 = 42
                    _ = v49
                }
            }
        default:
            v1 += (v1 - v1)
        }
    case "hello":
        if ((v1 % ((v1) | 1)) <= (^(v1))) {
            for fi50 := 0; fi50 < 14; fi50++ {
                v6 += 1
                _ = fi50
            }
            fmt.Println(v1, v4, v4)
            for fi51 := 0; fi51 < 14; fi51++ {
                v3 += 42
                for fi52 := 0; fi52 < 15; fi52++ {
                    if len(v5) > 0 {
                        v5[0] = fi52
                    }
                    if (v35 || v35) {
                        var v53 int = -2147483648
                        switch "foo bar" {
                        case "":
                            v5 = append(v5, 4294967295)
                        }
                        _ = v53
                    }
                    if true {
                        if (v35 && v35) {
                            switch genFn1(42, 0, genFn1(1, 255, fi51)) {
                            case 128:
                                v1 = (v1 << ((503) & 63))
                            case -9223372036854775808:
                                v6 = 255
                            default:
                                v1 += (v1 ^ 635)
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        switch (fi52 >> (((^(v4))) & 63)) {
                        case -1:
                            v54 := GenS5{F0: 18446744073709551615, F1: 0, F2: 152}
                            _ = v54
                        default:
                            switch "a" {
                            case "hello":
                                if v35 {
                                    v55 := GenS1{F0: v1, F1: uint8(v4)}
                                    var v56 int16 = v1
                                    if (("a" + "\n\t") >= "") {
                                        fmt.Println(v35)
                                        if (v35 || v35) {
                                            var v57 int64 = -1
                                            _ = len(v5)
                                            _ = v57
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                    } else {
                                        v58 := GenS6{F0: 9223372036854775807}
                                        v56 = 645
                                        v59 := GenS5{F0: 18446744073709551615, F1: genFn0(), F2: 141}
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = v55
                                    _ = v56
                                } else {
                                    var v60 float32 = 1.5
                                    v60 += (v60 * ((v60 - 3.14159) + float32(1.5)))
                                    if (((0.0 * -2.25) - 0.1) > (float64(v60) * ((1e10 / 3.14159) - -2.25))) {
                                        switch (fi52 | fi52) {
                                        case -279:
                                            v35 = ((((((1e10 + 1e10) * 0.0) + 0.1) / 1.5) + float64(0.1)) == (1e10 * -2.25))
                                        case 32767:
                                            var v61 uint = 2
                                            _ = v61
                                        }
                                        var v62 float64 = 0.1
                                        fmt.Println(v1, fi52, fi52)
                                        _ = v62
                                    } else {
                                        fmt.Println(v6, v4, v2)
                                        if (!v35) {
                                            if (v35 || (!(472 >= 2147483647))) {
                                                fmt.Println(v6, v4)
                                            } else {
                                                v4 = 822
                                                v60 -= v60
                                                var v63 int = (-(fi52))
                                                _ = v63
                                            }
                                            var v64 uint = 2
                                            _ = v64
                                        } else {
                                            switch ("a" + "a") {
                                            case "foo bar":
                                                v3 += (v6 * (^(v2)))
                                                v35 = (42 <= -9223372036854775808)
                                            }
                                            v6 -= v6
                                            var v65 int32 = 2147483647
                                            _ = v65
                                        }
                                        fmt.Println(v6, v60, fi51)
                                    }
                                    _ = v60
                                }
                            case "a":
                                switch ((("" + (("foo bar" + "foo bar") + ("\n\t" + "\n\t"))) + ("foo bar" + "foo bar")) + "a") {
                                case "\n\t":
                                    v5 = append(v5, (fi51 << (((v4 ^ 868)) & 63)))
                                default:
                                    var v66 uint16 = v2
                                    _ = v66
                                }
                            case "\n\t":
                                v1 -= (v1 | v1)
                            default:
                                var v67 int = fi52
                                _ = v67
                            }
                            var v68 int8 = 127
                            _ = v68
                        }
                        v1 += (v1 & v1)
                    }
                    _ = fi52
                }
                var v69 int = genFn1(0, 0, fi51)
                _ = fi51
                _ = v69
            }
        }
    default:
        v6 = 32767
    }
    switch "foo bar" {
    case "":
        var v70 uint8 = 2
        _ = v70
    case "foo bar":
        _ = len(v5)
    }
    for fi71 := 0; fi71 < 12; fi71++ {
        v3 -= 1
        _ = fi71
    }
    if (-1 >= 9223372036854775807) {
        for fi72 := 0; fi72 < 13; fi72++ {
            var v73 uint64 = 182
            v35 = v35
            v74 := GenS6{F0: 9223372036854775807}
            _ = fi72
            _ = v73
            _ = v74
        }
    }
    for fi75 := 0; fi75 < 11; fi75++ {
        var v76 uint32 = 0
        _ = fi75
        _ = v76
    }
    for fi77 := 0; fi77 < 14; fi77++ {
        if len(v5) > 0 {
            v5[0] = fi77
        }
        switch 42 {
        case 2147483647:
            v35 = (127 < 127)
        case 32768:
            v3 = v6
        }
        _ = fi77
    }
    var v78 int32 = -637
    fmt.Println(v3, v3)
    var v79 float64 = 0.1
    v5 = append(v5, genFn1(80, 0, 4294967295))
    if ((v35 || false) || (!(42 != 167))) {
        v80 := GenS2{F0: 0}
        if ((32767 < v1) && false) {
            v81 := GenS6{F0: -1}
            var v82 uint16 = 2
            for fi83 := 0; fi83 < 12; fi83++ {
                for fi84 := 0; fi84 < 6; fi84++ {
                    switch "a" {
                    case "a":
                        var v85 int = (fi84 % ((fi83) | 1))
                        _ = v85
                    default:
                        if ((((-1 > 91) && (!v35)) || v35) || ((^(v4)) != (v4 - 18446744073709551615))) {
                            switch "\n\t" {
                            case "\n\t":
                                switch "foo bar" {
                                case "hello":
                                    if len(v5) > 0 {
                                        v5[0] = (fi83 % ((127) | 1))
                                    }
                                default:
                                    v3 += 255
                                }
                            case "":
                                switch "a" {
                                case "\n\t":
                                    v4 -= uint(v6)
                                default:
                                    var v86 int64 = genFn0()
                                    _ = v86
                                }
                            default:
                                var v87 int16 = -32768
                                _ = v87
                            }
                            var v88 bool = (v79 >= ((v79 + v79) * (1.5 - v79)))
                            var v89 float32 = 3.14159
                            if (genFn1(4294967295, 94, 479) < genFn1(65535, 255, (fi83 & fi83))) {
                                v90 := GenS1{F0: (v1 & v1), F1: 2}
                                if v88 {
                                    fmt.Println(v82)
                                    v91 := GenS4{F0: (v6 << ((2) & 63)), F1: v3, F2: (-(v78))}
                                    v92 := GenS3{F0: (v78 >> (((v4 << ((1) & 63))) & 63)), F1: v35}
                                    _ = v91
                                    _ = v92
                                }
                                _ = v90
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = len(v5)
                            _ = v88
                            _ = v89
                        }
                        v93 := GenS4{F0: 0, F1: v3, F2: -1}
                        var v94 float32 = 0.0
                        _ = v93
                        _ = v94
                    }
                    v6 = 233
                    _ = fi84
                }
                fmt.Println(fi83)
                _ = fi83
            }
            _ = v81
            _ = v82
        }
        _ = v80
    } else {
        if (v35 && v35) {
            v95 := GenS4{F0: v2, F1: v6, F2: (-(v78))}
            _ = v95
        } else {
            v5 = append(v5, 2147483647)
        }
        v1 += v1
        switch 127 {
        case 4294967295:
            v96 := GenS7{F0: -128, F1: v78}
            _ = v96
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v79
}
