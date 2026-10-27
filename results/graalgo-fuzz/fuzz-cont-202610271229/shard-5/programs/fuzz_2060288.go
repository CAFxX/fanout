package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int64; F1 string; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint; F1 uint; F2 int64 }
type GenS10 struct { F0 uint; F1 int64 }
func genFn0(p0 int32, p1 int32) uint64 {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: (p1 - (^(p0))), F1: 0}
        if true {
            switch 2147483647 {
            case 128:
                v8 := GenS1{F0: p1, F1: 420}
                _ = v8
            case 0:
                fmt.Println(p0, p0, p0)
            default:
                if (127 != fi6) {
                    if (0 <= 255) {
                        fmt.Println(fi6)
                    }
                    switch fi6 {
                    case 32768:
                        if (1 != 986) {
                            for fi9 := 0; fi9 < 18; fi9++ {
                                p1 += (p0 * (-(p0)))
                                var v10 float32 = 0.0
                                _ = fi9
                                _ = v10
                            }
                            v11 := GenS2{F0: "\n\t", F1: true, F2: int8(fi6)}
                            _ = v11
                        } else {
                            p0 -= p1
                        }
                    }
                }
            }
            switch "foo bar" {
            case "foo bar":
                v12 := []int{(^(fi6))}
                _ = v12
            }
        } else {
            p0 += (p0 << ((2) & 63))
            p1 = p1
        }
        switch "hello" {
        case "hello":
            if false {
                switch (fi6 / ((fi6) | 1)) {
                case -54:
                    switch fi6 {
                    case 32767:
                        if false {
                            p1 -= (p1 ^ (-(p0)))
                        }
                    case 1:
                        if (2147483648 >= 2147483647) {
                            var v13 uint32 = 192
                            _ = v13
                        }
                    case 32768:
                        var v14 uint32 = 1
                        _ = v14
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            switch "\n\t" {
                            case "\n\t":
                                p1 -= p0
                            default:
                                fmt.Println(p1, p1, p1)
                            }
                        case "":
                            v15 := []int{(fi6 - fi6), (fi6 << ((0) & 63))}
                            _ = v15
                        case "a":
                            for fi16 := 0; fi16 < 3; fi16++ {
                                fmt.Println(p1, fi16)
                                p0 = -2147483648
                                _ = fi16
                            }
                        }
                    }
                case 42:
                    if true {
                        var v17 uint = 1
                        p0 += p0
                        p0 += (p0 >> (((v17 & v17)) & 63))
                        _ = v17
                    } else {
                        v18 := GenS2{F0: "", F1: false, F2: 100}
                        _ = v18
                    }
                case 127:
                    var v19 uint16 = 1
                    _ = v19
                default:
                    if true {
                        p0 = (^(p0))
                        switch (fi6 << ((116) & 63)) {
                        case 4294967295:
                            var v20 int32 = p0
                            _ = v20
                        case 128:
                            p0 = (p0 + (p0 >> ((2) & 63)))
                        default:
                            switch (fi6 * (-(fi6))) {
                            case 32767:
                                var v21 uint8 = 2
                                _ = v21
                            default:
                                var v22 uint = 255
                                _ = v22
                            }
                        }
                    } else {
                        p1 += 0
                        for fi23 := 0; fi23 < 17; fi23++ {
                            v24 := GenS3{F0: 0, F1: 1}
                            p1 = (p0 & p0)
                            _ = fi23
                            _ = v24
                        }
                    }
                }
                if false {
                    p0 = (p1 / ((p1) | 1))
                    v25 := GenS2{F0: ("" + ("foo bar" + "\n\t")), F1: true, F2: -46}
                    if (-327 == -913) {
                        p1 -= p0
                        if false {
                            fmt.Println(p1, p1, fi6)
                            v26 := GenS4{F0: true, F1: 153}
                            fmt.Println(p1)
                            _ = v26
                        }
                    }
                    _ = v25
                }
            } else {
                v27 := []int{(fi6 << ((42) & 63)), (fi6 & fi6), fi6}
                var v28 int32 = 2147483647
                p0 += p0
                _ = v27
                _ = v28
            }
        case "foo bar":
            if (513 <= 42) {
                if false {
                    p0 -= -2147483648
                    p1 = (-(p0))
                    p0 = (p0 % (((p1 + p1)) | 1))
                }
            } else {
                p1 -= 1
                v29 := GenS1{F0: p1, F1: 18446744073709551615}
                var v30 uint64 = 1
                _ = v29
                _ = v30
            }
        }
        _ = fi6
        _ = v7
    }
    p0 += (p0 | p1)
    fmt.Println(p0, p1)
    for fi31 := 0; fi31 < 20; fi31++ {
        p0 += p1
        if true {
            var v32 uint8 = 223
            var v33 uint64 = 813
            var v34 int64 = -9223372036854775808
            _ = v32
            _ = v33
            _ = v34
        } else {
            fmt.Println(fi31, p0, p1)
        }
        _ = fi31
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 int8 = -67
    v2 := make(map[string]bool)
    var v3 int = 498
    var v4 float64 = 0.1
    var v5 int8 = v1
    switch (^(v3)) {
    case 32768:
        fmt.Println(v5)
    case 42:
        var v35 uint16 = 0
        _ = v35
    }
    v36 := GenS3{F0: 846, F1: 2}
    var v37 uint8 = 97
    if false {
        for fi38 := 0; fi38 < 3; fi38++ {
            if false {
                for fi39 := 0; fi39 < 2; fi39++ {
                    var v40 int64 = 144
                    _ = fi39
                    _ = v40
                }
            }
            _ = fi38
        }
        v41 := GenS4{F0: (!(!false)), F1: 2}
        _ = v41
    } else {
        var v42 int16 = 689
        fmt.Println(v1)
        _ = v42
    }
    var v43 uint8 = 1
    if (!(2147483647 != -2147483648)) {
        for fi44 := 0; fi44 < 15; fi44++ {
            v45 := GenS5{F0: -1, F1: (("" + "") + "hello"), F2: 619}
            for fi46 := 0; fi46 < 2; fi46++ {
                v47 := GenS1{F0: 2147483647, F1: 65535}
                switch "a" {
                case "\n\t":
                    switch fi44 {
                    case 128:
                        if ((v5 / (((v1 - 1)) | 1)) != (v1 & (v1 % (((v1 / ((v5) | 1))) | 1)))) {
                            _ = len(v2)
                            if true {
                                switch (fi44 << ((0) & 63)) {
                                case 0:
                                    var v48 bool = true
                                    _ = v48
                                case -2147483648:
                                    switch (v3 & (-(v3))) {
                                    case 50:
                                        if false {
                                            var v49 uint64 = uint64(v43)
                                            _ = v49
                                        }
                                    case 32767:
                                        v37 -= v37
                                    }
                                }
                                fmt.Println(v43, v5, v5)
                            }
                            v5 = v1
                        }
                    case 32768:
                        switch "\n\t" {
                        case "a":
                            fmt.Println(fi44, v1)
                        case "foo bar":
                            if false {
                                var v50 uint16 = 249
                                switch v3 {
                                case -9223372036854775808:
                                    v51 := []int{-238}
                                    _ = v51
                                case -2147483648:
                                    v52 := GenS4{F0: (false && (fi44 == (fi44 / (((fi46 & (fi44 - (v3 * fi44)))) | 1)))), F1: v50}
                                    _ = v52
                                }
                                var v53 uint8 = v43
                                _ = v50
                                _ = v53
                            } else {
                                v54 := GenS5{F0: 0, F1: ("hello" + "a"), F2: 635}
                                fmt.Println(v1, fi44, fi46)
                                fmt.Println(fi46)
                                v55 := []int{(fi46 * 1), v3, v3, (fi46 >> ((2) & 63))}
                                _ = v54
                                _ = v55
                            }
                        default:
                            v56 := GenS6{F0: -9223372036854775808}
                            fmt.Println(fi46, fi44, v43)
                            _ = v56
                        }
                    case 2147483647:
                        var v57 int8 = (v1 >> ((99) & 63))
                        if false {
                            var v58 uint = 42
                            if true {
                                if (false || (-9223372036854775808 != 42)) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v59 := []int{}
                                        _ = v59
                                    case "a":
                                        _ = len(v2)
                                    }
                                }
                                var v60 string = "foo bar"
                                var v61 int16 = int16(fi44)
                                _ = v60
                                _ = v61
                            } else {
                                if (3.14159 > (3.14159 + 0.1)) {
                                    v57 -= int8(v43)
                                }
                                v1 = 108
                                if (!(v5 >= (-(v1)))) {
                                    var v62 float64 = v4
                                    var v63 int16 = 1
                                    _ = v62
                                    _ = v63
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(v43, v43, v57)
                                    case "hello":
                                        var v64 uint32 = 2
                                        _ = v64
                                    case "":
                                        fmt.Println(v3)
                                    default:
                                        switch -2147483648 {
                                        case 697:
                                            var v65 int64 = 42
                                            _ = v65
                                        case 9223372036854775807:
                                            _ = len(v2)
                                        default:
                                            v66 := GenS7{F0: (^(v58))}
                                            _ = v66
                                        }
                                    }
                                    if true {
                                        var v67 string = ""
                                        var v68 int8 = (-(v1))
                                        v4 = ((1.5 - v4) / (float64(-2.25) - 0.1))
                                        _ = v67
                                        _ = v68
                                    }
                                }
                            }
                            if (false && (!(((!false) && true) && true))) {
                                var v69 float32 = (0.1 * (-2.25 + (0.1 + 3.14159)))
                                if (!true) {
                                    v70 := []int{fi46, -166}
                                    _ = v70
                                } else {
                                    v71 := GenS4{F0: (4294967295 > 938), F1: 1}
                                    v72 := GenS8{F0: 42}
                                    _ = v71
                                    _ = v72
                                }
                                v73 := v2["\n\t"]
                                var v74 string = ""
                                _ = v69
                                _ = v73
                                _ = v74
                            } else {
                                v1 += v5
                            }
                            _ = v58
                        } else {
                            v75 := []int{fi44, fi46, (fi46 | (fi46 - (fi46 + 4294967295)))}
                            _ = v75
                        }
                        v76 := GenS1{F0: -2147483648, F1: uint(v43)}
                        _ = v57
                        _ = v76
                    default:
                        if (4294967295 >= 0) {
                            if (!(1 > 42)) {
                                v3 = (v3 % ((-443) | 1))
                                if false {
                                    v77 := v2["a"]
                                    _ = v77
                                } else {
                                    v1 -= (^(v1))
                                }
                            }
                            fmt.Println(v43)
                            v78 := v2[""]
                            _ = v78
                        } else {
                            v2[("" + ("foo bar" + "hello"))] = (v4 == 3.14159)
                        }
                    }
                default:
                    switch (^(v3)) {
                    case 131:
                        if false {
                            v79 := []int{453, fi44, int(v1), (fi44 + fi46)}
                            var v80 uint32 = 375
                            var v81 uint8 = (v37 % (((v37 / ((1) | 1))) | 1))
                            switch "a" {
                            case "\n\t":
                                v43 = v37
                            case "hello":
                                switch v3 {
                                case 32768:
                                    switch -2147483648 {
                                    case -9223372036854775808:
                                        switch 42 {
                                        case 128:
                                            v4 -= ((v4 * float64(-2.25)) * 1.5)
                                        case 597:
                                            _ = len(v79)
                                        default:
                                            v82 := GenS7{F0: 18446744073709551615}
                                            _ = v82
                                        }
                                    case -349:
                                        fmt.Println(v1, v81, v1)
                                    case -88:
                                        v43 = v81
                                    default:
                                        _ = len(v2)
                                    }
                                case 48:
                                    var v83 float64 = 0.0
                                    _ = v83
                                case -376:
                                    var v84 float64 = v4
                                    _ = v84
                                default:
                                    v85 := GenS9{F0: 4294967295, F1: 591, F2: 992}
                                    _ = v85
                                }
                            }
                            if (!false) {
                                v86 := GenS8{F0: 639}
                                var v87 uint32 = v80
                                v5 += v5
                                _ = v86
                                _ = v87
                            }
                            _ = v79
                            _ = v80
                            _ = v81
                        }
                    case -1:
                        fmt.Println(fi44, v37, v5)
                    }
                }
                v88 := GenS3{F0: 2, F1: 715}
                v89 := GenS3{F0: 2, F1: 875}
                v90 := []int{fi44, (v3 | (v3 | (v3 & (v3 >> ((4294967295) & 63))))), fi44}
                _ = fi46
                _ = v47
                _ = v88
                _ = v89
                _ = v90
            }
            v91 := GenS10{F0: 2, F1: 9223372036854775807}
            _ = fi44
            _ = v45
            _ = v91
        }
        var v92 float64 = v4
        _ = v92
    }
    fmt.Println(v43, v4)
    for fi93 := 0; fi93 < 12; fi93++ {
        v43 -= (^(v43))
        fmt.Println(v5, v43, v5)
        _ = fi93
    }
    v1 += (v1 | v5)
    var v94 int32 = -199
    v94 -= -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v43
    _ = v5
    _ = v94
}
