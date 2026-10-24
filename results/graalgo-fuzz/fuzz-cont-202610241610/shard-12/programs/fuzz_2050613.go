package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 bool; F1 uint8 }
type GenS9 struct { F0 uint; F1 uint16 }
type GenS10 struct { F0 string; F1 uint; F2 bool }
type GenS11 struct { F0 uint8; F1 uint64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint64; F1 int8 }
func genFn0(p0 bool) uint64 {
    var v5 uint = 99
    var v6 string = "hello"
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            if ((0.0 - float64(0.1)) > 3.14159) {
                var v9 uint = (v5 ^ (v5 / (((v5 & 42)) | 1)))
                if (v6 < v6) {
                    var v10 uint32 = 716
                    v11 := []int{(fi7 >> (((v9 << (((v9 | v5)) & 63))) & 63)), 503, 0, fi7}
                    if p0 {
                        v6 = "\n\t"
                        var v12 uint8 = 180
                        v11 = append(v11, fi7)
                        _ = v12
                    } else {
                        p0 = p0
                        if true {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            v9 = v9
                            if (!true) {
                                if (!p0) {
                                    fmt.Println(v5, v10, v10)
                                    if len(v11) > 0 {
                                        v11[0] = fi8
                                    }
                                }
                                fmt.Println(v5, p0, v9)
                                if len(v11) > 0 {
                                    v11[0] = fi8
                                }
                            }
                        }
                        v11 = append(v11, fi7)
                    }
                    var v13 int32 = 2147483647
                    var v14 uint32 = v10
                    _ = v10
                    _ = v11
                    _ = v13
                    _ = v14
                }
                v15 := []int{fi8}
                if p0 {
                    if p0 {
                        p0 = true
                        if (-32768 == 0) {
                            if p0 {
                                v9 += (v9 >> (((v5 >> ((v9) & 63))) & 63))
                            }
                        }
                    }
                } else {
                    var v16 bool = (true || p0)
                    fmt.Println(v9)
                    if (((!(v16 && v16)) || false) || true) {
                        if len(v15) > 0 {
                            v15[0] = -294
                        }
                        var v17 int8 = 100
                        v15 = append(v15, (fi8 - -9223372036854775808))
                        _ = v17
                    } else {
                        v18 := GenS1{F0: 32767, F1: 7, F2: 255}
                        _ = v18
                    }
                    _ = v16
                }
                _ = len(v15)
                if p0 {
                    v19 := GenS1{F0: -32768, F1: 42, F2: uint8(v9)}
                    if p0 {
                        var v20 float64 = -2.25
                        _ = v20
                    }
                    v15 = append(v15, (^(fi8)))
                    var v21 bool = ((fi8 & fi8) != fi7)
                    var v22 bool = (3.14159 == 0.0)
                    _ = v19
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS1{F0: -592, F1: 321, F2: 0}
                    if p0 {
                        var v24 uint32 = 1
                        if (uint32(v5) > v24) {
                            if p0 {
                                var v25 uint32 = (v24 ^ v24)
                                var v26 uint32 = v24
                                v27 := GenS2{F0: 42, F1: (v24 | v24), F2: v26}
                                _ = v25
                                _ = v26
                                _ = v27
                            } else {
                                switch 140 {
                                case -94:
                                    v9 = v5
                                case 128:
                                    v24 += v24
                                case -1:
                                    fmt.Println(fi7)
                                }
                            }
                        } else {
                            switch fi7 {
                            case 42:
                                v28 := GenS2{F0: 1, F1: (v24 & (^(v24))), F2: (v24 / (((v24 - v24)) | 1))}
                                _ = v28
                            case 0:
                                v6 = (v6 + v6)
                            default:
                                switch fi8 {
                                case -272:
                                    v9 = v5
                                case 0:
                                    fmt.Println(p0, v24, fi8)
                                default:
                                    switch (^(fi7)) {
                                    case 127:
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                    case -2147483648:
                                        if true {
                                            if (!p0) {
                                                if len(v15) > 0 {
                                                    v15[0] = fi7
                                                }
                                                if p0 {
                                                    switch v6 {
                                                    case "\n\t":
                                                        var v29 float32 = (1e10 / float32(float64((1e10 + 0.0))))
                                                        _ = v29
                                                    default:
                                                        v30 := GenS2{F0: 296, F1: v24, F2: 1}
                                                        _ = v30
                                                    }
                                                    v31 := GenS3{F0: -1, F1: 65, F2: 0}
                                                    _ = v31
                                                }
                                                v6 = "foo bar"
                                                switch ("foo bar" + v6) {
                                                case "a":
                                                    fmt.Println(v5, fi8)
                                                case "":
                                                    v5 -= (v5 + v5)
                                                case "foo bar":
                                                    v15 = append(v15, (fi7 % ((fi8) | 1)))
                                                }
                                            }
                                            var v32 int16 = -277
                                            v33 := GenS4{F0: 42}
                                            _ = v32
                                            _ = v33
                                        } else {
                                            v34 := GenS5{F0: 32767, F1: "hello"}
                                            if p0 {
                                                if (false && (!(p0 || true))) {
                                                    if len(v15) > 0 {
                                                        v15[0] = (-(fi8))
                                                    }
                                                    switch (fi8 << ((4294967295) & 63)) {
                                                    case 473:
                                                        var v35 uint8 = 2
                                                        _ = v35
                                                    case -232:
                                                        v9 += (v9 | v5)
                                                        var v36 uint = (v9 ^ v9)
                                                        var v37 float64 = 0.1
                                                        var v38 int32 = 1
                                                        var v39 float32 = 1e10
                                                        _ = v36
                                                        _ = v37
                                                        _ = v38
                                                        _ = v39
                                                    case 32768:
                                                        if p0 {
                                                            if len(v15) >= 2 {
                                                                v15 = v15[1:len(v15)]
                                                            }
                                                            if ((v24 * v24) < 2147483647) {
                                                                _ = len(v15)
                                                                var v40 string = (v6 + v6)
                                                                _ = v40
                                                            }
                                                        } else {
                                                            if (p0 && false) {
                                                                v41 := GenS6{F0: (fi7 + 127)}
                                                                var v42 float64 = (0.1 - ((0.1 + (((float64(float32(1.5)) - (float64(0.0) - -2.25)) + (-2.25 * ((3.14159 + 1.5) + (-2.25 * (3.14159 / (-2.25 * 1.5)))))) - 1e10)) - 3.14159))
                                                                _ = v41
                                                                _ = v42
                                                            } else {
                                                                switch (-(fi8)) {
                                                                case 128:
                                                                    if p0 {
                                                                        var v43 int = fi7
                                                                        if p0 {
                                                                            var v44 uint64 = 18446744073709551615
                                                                            v45 := GenS5{F0: int16(v44), F1: v6}
                                                                            v46 := GenS3{F0: 1000, F1: uint8(v44), F2: 128}
                                                                            _ = v44
                                                                            _ = v45
                                                                            _ = v46
                                                                        }
                                                                        _ = v43
                                                                    } else {
                                                                        var v47 uint64 = 0
                                                                        v6 = v6
                                                                        switch fi7 {
                                                                        case 4294967295:
                                                                            v24 = 1
                                                                        case 32768:
                                                                            v48 := GenS6{F0: (fi7 % ((fi7) | 1))}
                                                                            _ = v48
                                                                        case -1:
                                                                            var v49 int32 = 0
                                                                            _ = v49
                                                                        }
                                                                        if true {
                                                                            v24 = v24
                                                                            p0 = true
                                                                            fmt.Println(fi8, v47, v47)
                                                                        } else {
                                                                            switch v6 {
                                                                            case "a":
                                                                                var v50 int = (fi8 & (fi7 + fi7))
                                                                                _ = v50
                                                                            case "":
                                                                                if (873 == 1) {
                                                                                    if len(v15) >= 2 {
                                                                                        v15 = v15[1:len(v15)]
                                                                                    }
                                                                                }
                                                                            case "\n\t":
                                                                                var v51 uint32 = v24
                                                                                _ = v51
                                                                            default:
                                                                                v52 := GenS1{F0: 728, F1: v47, F2: 2}
                                                                                _ = v52
                                                                            }
                                                                            if len(v15) >= 2 {
                                                                                v15 = v15[1:len(v15)]
                                                                            }
                                                                        }
                                                                        _ = v47
                                                                    }
                                                                case 127:
                                                                    var v53 float64 = (0.0 / 3.14159)
                                                                    _ = v53
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v6 = ""
                                                    v24 += v24
                                                } else {
                                                    fmt.Println(p0)
                                                    var v54 float64 = ((-2.25 / 3.14159) * (1e10 * 1e10))
                                                    var v55 int64 = 1
                                                    _ = v54
                                                    _ = v55
                                                }
                                            }
                                            _ = v34
                                        }
                                    case 611:
                                        var v56 string = "hello"
                                        _ = v56
                                    }
                                }
                            }
                            v9 += v5
                            switch (fi7 << ((v5) & 63)) {
                            case 32767:
                                p0 = p0
                            case 32768:
                                if (636 >= 764) {
                                    _ = len(v15)
                                    var v57 float32 = float32(-2.25)
                                    if p0 {
                                        v58 := GenS7{F0: (v5 & 907), F1: 127}
                                        _ = len(v15)
                                        switch ("foo bar" + v6) {
                                        case "a":
                                            var v59 float64 = 0.0
                                            _ = v59
                                        case "hello":
                                            var v60 uint8 = 0
                                            _ = v60
                                        case "":
                                            p0 = (0 < 522)
                                        default:
                                            switch 128 {
                                            case -1:
                                                v61 := GenS5{F0: 0, F1: ""}
                                                _ = v61
                                            case 2147483647:
                                                if len(v15) >= 2 {
                                                    v15 = v15[1:len(v15)]
                                                }
                                            case -541:
                                                if ((1 < -525) || p0) {
                                                    _ = len(v15)
                                                } else {
                                                    if p0 {
                                                        _ = len(v15)
                                                    }
                                                    switch 9223372036854775807 {
                                                    case 2147483647:
                                                        fmt.Println(v57, fi7)
                                                        v62 := GenS4{F0: 100}
                                                        _ = v62
                                                    }
                                                }
                                            default:
                                                if true {
                                                    v63 := GenS8{F0: ("\n\t" < ((("\n\t" + ("foo bar" + "hello")) + "") + v6)), F1: 6}
                                                    var v64 int32 = 42
                                                    v65 := GenS9{F0: (v9 * 18446744073709551615), F1: 32767}
                                                    _ = v63
                                                    _ = v64
                                                    _ = v65
                                                }
                                            }
                                        }
                                        if false {
                                            v66 := GenS1{F0: -773, F1: 2, F2: 127}
                                            _ = v66
                                        } else {
                                            v5 += v5
                                            var v67 int64 = -1
                                            _ = v67
                                        }
                                        v24 = v24
                                        _ = v58
                                    }
                                    _ = v57
                                }
                            default:
                                if (1e10 == 0.1) {
                                    switch v6 {
                                    case "\n\t":
                                        switch ("foo bar" + v6) {
                                        case "a":
                                            switch v6 {
                                            case "\n\t":
                                                var v68 uint = (v9 / ((v5) | 1))
                                                _ = v68
                                            case "hello":
                                                if (2147483647 < (^(v24))) {
                                                    var v69 int8 = 42
                                                    _ = len(v15)
                                                    var v70 int64 = -261
                                                    _ = v69
                                                    _ = v70
                                                } else {
                                                    v71 := GenS10{F0: v6, F1: uint(v24), F2: p0}
                                                    _ = v71
                                                }
                                            case "":
                                                v24 += (^(v24))
                                            }
                                        }
                                    case "foo bar":
                                        var v72 float32 = 1.5
                                        _ = v72
                                    default:
                                        v73 := GenS11{F0: uint8(v9), F1: 4294967295}
                                        _ = v73
                                    }
                                }
                            }
                        }
                        _ = v24
                    } else {
                        v6 = "hello"
                    }
                    var v74 uint32 = 598
                    if len(v15) > 0 {
                        v15[0] = (^(fi8))
                    }
                    _ = v23
                    _ = v74
                }
                p0 = false
                p0 = p0
                _ = v15
                _ = v9
            }
            _ = fi8
        }
        v75 := []int{}
        v76 := GenS12{F0: -688}
        _ = fi7
        _ = v75
        _ = v76
    }
    _ = p0
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1() uint16 {
    var v77 int16 = 1
    switch "foo bar" {
    case "a":
        var v78 float32 = float32((((-2.25 - ((float64((-2.25 + -2.25)) * (float64(1.5) + 1e10)) - 1.5)) * 1e10) / 1.5))
        _ = v78
    case "foo bar":
        for fi79 := 0; fi79 < 3; fi79++ {
            if (1 <= 42) {
                var v80 bool = ((255 < 65535) || false)
                var v81 uint = 2
                _ = v80
                _ = v81
            } else {
                v82 := []int{-35}
                for fi83 := 0; fi83 < 18; fi83++ {
                    v82 = append(v82, (-(fi79)))
                    v84 := GenS6{F0: fi83}
                    _ = fi83
                    _ = v84
                }
                _ = len(v82)
                _ = v82
            }
            for fi85 := 0; fi85 < 15; fi85++ {
                v86 := GenS7{F0: 2, F1: -128}
                switch fi79 {
                case 2147483647:
                    v87 := []int{fi79}
                    _ = v87
                default:
                    v77 = v77
                }
                _ = fi85
                _ = v86
            }
            v77 += (^(v77))
            _ = fi79
        }
    }
    v77 += 42
    _ = v77
    return 2
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]bool)
    var v3 int8 = 35
    v4 := make(map[uint32]bool)
    fmt.Println(v1, v3)
    v88 := v4[65535]
    var v89 int8 = v3
    v90 := []int{-444}
    switch 128 {
    case 0:
        var v91 float64 = (0.0 / (3.14159 / 1e10))
        _ = v91
    }
    if v88 {
        switch "hello" {
        case "a":
            v92 := GenS13{F0: 4294967295, F1: v89}
            _ = v92
        case "\n\t":
            fmt.Println(v3, v3, v88)
        case "hello":
            for fi93 := 0; fi93 < 13; fi93++ {
                var v94 float32 = -2.25
                switch fi93 {
                case -980:
                    v1 = (v1 % (((v1 % ((-9223372036854775808) | 1))) | 1))
                }
                _ = fi93
                _ = v94
            }
        }
        v88 = v88
    }
    switch 127 {
    case 2147483647:
        for fi95 := 0; fi95 < 19; fi95++ {
            for fi96 := 0; fi96 < 16; fi96++ {
                v97 := GenS1{F0: 42, F1: 1, F2: 127}
                v1 -= v1
                fmt.Println(v1, v1)
                fmt.Println(fi96)
                _ = fi96
                _ = v97
            }
            _ = fi95
        }
    case 621:
        delete(v2, 127)
    default:
        fmt.Println(v88, v3)
    }
    fmt.Println(v89, v3, v88)
    v98 := GenS10{F0: "foo bar", F1: 65535, F2: (!(!(1 > 42)))}
    var v99 int8 = (v89 + (v3 / (((-(v89))) | 1)))
    for fi100 := 0; fi100 < 8; fi100++ {
        var v101 string = ("" + (("" + "hello") + (("" + "\n\t") + "a")))
        _ = fi100
        _ = v101
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("slicelen v90:", len(v90), v90)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v88
    _ = v89
    _ = v90
    _ = v98
    _ = v99
}
