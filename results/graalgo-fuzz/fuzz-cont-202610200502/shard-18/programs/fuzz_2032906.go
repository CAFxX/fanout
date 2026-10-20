package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64; F1 string }
func genFn0(p0 uint16, p1 uint64) int8 {
    p1 = (p1 << ((454) & 63))
    for fi7 := 0; fi7 < 16; fi7++ {
        p1 += (p1 % ((1) | 1))
        v8 := []int{(^(fi7)), (fi7 * (fi7 / ((fi7) | 1))), fi7}
        for fi9 := 0; fi9 < 3; fi9++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    var v10 int16 = 564
    fmt.Println(p1, v10, p0)
    _ = p0
    _ = p1
    _ = v10
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = (((("a" + "a") + "\n\t") + "\n\t") + "hello")
    var v3 uint32 = 2147483647
    var v4 uint64 = 688
    var v5 uint8 = 1
    var v6 float32 = 3.14159
    fmt.Println(v4, v5)
    v6 -= v6
    v4 += 42
    switch ("hello" + "") {
    case "\n\t":
        var v11 uint16 = 42
        _ = v11
    default:
        switch v2 {
        case "":
            var v12 int64 = 0
            _ = v12
        case "a":
            var v13 uint16 = 858
            _ = v13
        default:
            for fi14 := 0; fi14 < 8; fi14++ {
                v5 = v5
                _ = fi14
            }
        }
    }
    v15 := GenS1{F0: (!(((v6 / (-2.25 - v6)) - (v6 / (v6 * (v6 - -2.25)))) < v6))}
    switch v2 {
    case "a":
        switch v2 {
        case "a":
            switch 42 {
            case -2147483648:
                var v16 string = "\n\t"
                _ = v16
            default:
                for fi17 := 0; fi17 < 3; fi17++ {
                    fmt.Println(v6, v6)
                    for fi18 := 0; fi18 < 18; fi18++ {
                        if (true && true) {
                            switch "\n\t" {
                            case "foo bar":
                                v19 := []int{490, (fi17 >> ((958) & 63)), 128, (fi18 + (fi17 + fi18))}
                                _ = v19
                            default:
                                var v20 int16 = -1
                                _ = v20
                            }
                            var v21 uint = 340
                            _ = v21
                        }
                        v22 := GenS1{F0: (-103 == genFn0(42, 2))}
                        fmt.Println(fi17)
                        _ = fi18
                        _ = v22
                    }
                    _ = fi17
                }
            }
        case "foo bar":
            for fi23 := 0; fi23 < 17; fi23++ {
                if false {
                    v5 += v5
                    var v24 uint32 = v3
                    switch fi23 {
                    case -675:
                        v25 := GenS1{F0: false}
                        _ = v25
                    default:
                        v26 := GenS1{F0: false}
                        _ = v26
                    }
                    _ = v24
                } else {
                    var v27 uint64 = v4
                    v6 -= (v6 + v6)
                    _ = v27
                }
                _ = fi23
            }
        default:
            if (false || (109 >= int8(v4))) {
                switch -9223372036854775808 {
                case -2147483648:
                    v6 += 1.5
                case 4294967295:
                    var v28 string = ((v2 + v2) + (v2 + v2))
                    _ = v28
                case 42:
                    switch (v2 + v2) {
                    case "hello":
                        v29 := GenS1{F0: false}
                        _ = v29
                    }
                }
                v30 := []int{-244}
                _ = v30
            }
        }
    case "\n\t":
        fmt.Println(v4, v2)
    case "foo bar":
        v4 -= v4
    default:
        if false {
            switch 9223372036854775807 {
            case 128:
                var v31 int = 229
                _ = v31
            case 1:
                if (66 > 32768) {
                    fmt.Println(v2)
                    v3 += v3
                    fmt.Println(v3, v3, v6)
                } else {
                    v3 = 1
                }
            }
            var v32 float64 = float64(-2.25)
            v4 = 429
            _ = v32
        }
    }
    var v33 bool = ((true || (true && (0 != 4294967295))) || (2 == uint(v5)))
    v33 = v33
    _ = len(v1)
    if v33 {
        if (9223372036854775807 != 32768) {
            switch 1 {
            case 802:
                if (true || ((-1 >= -2147483648) || false)) {
                    switch -1 {
                    case -157:
                        delete(v1, v3)
                    default:
                        if true {
                            var v34 string = "a"
                            _ = v34
                        } else {
                            fmt.Println(v3)
                            v4 += 2
                        }
                    }
                    switch -9223372036854775808 {
                    case 1:
                        for fi35 := 0; fi35 < 20; fi35++ {
                            var v36 float64 = float64(v6)
                            for fi37 := 0; fi37 < 17; fi37++ {
                                v1[(v3 & v3)] = (v3 << ((42) & 63))
                                v36 += 0.1
                                _ = fi37
                            }
                            if (-1 >= -1) {
                                if ("" <= ((v2 + v2) + "\n\t")) {
                                    fmt.Println(v4, v36)
                                    if (255 < 1) {
                                        delete(v1, (v3 | (v3 + (^(v3)))))
                                    } else {
                                        var v38 float32 = 0.0
                                        _ = v38
                                    }
                                } else {
                                    v36 += v36
                                    v4 += 0
                                }
                                for fi39 := 0; fi39 < 20; fi39++ {
                                    var v40 int16 = -32768
                                    _ = fi39
                                    _ = v40
                                }
                            } else {
                                var v41 bool = (v5 < (v5 - 128))
                                var v42 uint8 = 0
                                if (42 > 2) {
                                    var v43 int16 = 0
                                    v5 = v42
                                    var v44 int64 = 114
                                    _ = v43
                                    _ = v44
                                }
                                _ = v41
                                _ = v42
                            }
                            _ = fi35
                            _ = v36
                        }
                    case 32767:
                        v2 = v2
                    default:
                        if (v6 == (v6 + v6)) {
                            var v45 int = -9223372036854775808
                            _ = v45
                        } else {
                            if (-77 >= 1) {
                                v6 = v6
                            } else {
                                var v46 uint32 = 311
                                _ = v46
                            }
                        }
                    }
                } else {
                    v47 := []int{1}
                    if v33 {
                        v48 := GenS1{F0: v33}
                        switch v2 {
                        case "\n\t":
                            for fi49 := 0; fi49 < 6; fi49++ {
                                switch 4294967295 {
                                case 0:
                                    var v50 float32 = 1e10
                                    _ = v50
                                case 32768:
                                    v33 = (1 > 18446744073709551615)
                                }
                                var v51 int = fi49
                                _ = fi49
                                _ = v51
                            }
                        case "hello":
                            v52 := GenS2{F0: v2, F1: 2147483647}
                            _ = v52
                        case "":
                            v53 := GenS3{F0: 1}
                            _ = v53
                        default:
                            if v33 {
                                var v54 int16 = -842
                                v6 -= v6
                                if (!v33) {
                                    switch v2 {
                                    case "":
                                        if ((v3 ^ v3) != v3) {
                                            if (!(2 < uint16(v5))) {
                                                for fi55 := 0; fi55 < 17; fi55++ {
                                                    for fi56 := 0; fi56 < 16; fi56++ {
                                                        var v57 uint8 = 42
                                                        _ = fi56
                                                        _ = v57
                                                    }
                                                    if (v5 != (v5 >> ((4294967295) & 63))) {
                                                        switch (fi55 >> ((313) & 63)) {
                                                        case 42:
                                                            fmt.Println(v4)
                                                        }
                                                        var v58 int64 = 1
                                                        _ = v58
                                                    }
                                                    _ = fi55
                                                }
                                                v59 := GenS4{F0: -2147483648, F1: v4}
                                                var v60 int32 = -494
                                                _ = v59
                                                _ = v60
                                            } else {
                                                if v33 {
                                                    v6 -= v6
                                                    switch "hello" {
                                                    case "a":
                                                        v5 -= 228
                                                    case "":
                                                        fmt.Println(v6, v33)
                                                    default:
                                                        if (-1 >= -1) {
                                                            for fi61 := 0; fi61 < 5; fi61++ {
                                                                v62 := v1[v3]
                                                                _ = fi61
                                                                _ = v62
                                                            }
                                                            if false {
                                                                var v63 uint32 = (v3 >> ((1) & 63))
                                                                var v64 int = 459
                                                                if false {
                                                                    v64 = v64
                                                                    var v65 string = "hello"
                                                                    for fi66 := 0; fi66 < 2; fi66++ {
                                                                        if (!(-9223372036854775808 <= 0)) {
                                                                            var v67 float32 = v6
                                                                            v4 -= v4
                                                                            _ = v67
                                                                        }
                                                                        for fi68 := 0; fi68 < 4; fi68++ {
                                                                            switch (^(fi68)) {
                                                                            case -2147483648:
                                                                                if (v65 <= ("foo bar" + ("" + ("a" + "")))) {
                                                                                    v3 -= v3
                                                                                }
                                                                            case 9223372036854775807:
                                                                                fmt.Println(fi68)
                                                                            case 127:
                                                                                var v69 uint64 = (v4 * 2)
                                                                                _ = v69
                                                                            }
                                                                            v70 := GenS5{F0: -9223372036854775808}
                                                                            v64 -= fi68
                                                                            _ = fi68
                                                                            _ = v70
                                                                        }
                                                                        _ = fi66
                                                                    }
                                                                    _ = v65
                                                                }
                                                                _ = v63
                                                                _ = v64
                                                            } else {
                                                                v6 = (-2.25 + 1e10)
                                                            }
                                                        }
                                                    }
                                                    var v71 uint = 521
                                                    _ = v71
                                                } else {
                                                    if true {
                                                        v54 = v54
                                                    }
                                                }
                                                var v72 uint8 = v5
                                                var v73 int16 = v54
                                                _ = v72
                                                _ = v73
                                            }
                                            var v74 int32 = -921
                                            _ = v74
                                        }
                                    default:
                                        for fi75 := 0; fi75 < 5; fi75++ {
                                            v76 := GenS3{F0: 0}
                                            _ = fi75
                                            _ = v76
                                        }
                                    }
                                    if len(v47) >= 2 {
                                        v47 = v47[1:len(v47)]
                                    }
                                } else {
                                    var v77 uint64 = 0
                                    if (!v33) {
                                        for fi78 := 0; fi78 < 6; fi78++ {
                                            var v79 float32 = (float32(0.1) / v6)
                                            v47 = append(v47, fi78)
                                            var v80 uint = 42
                                            _ = fi78
                                            _ = v79
                                            _ = v80
                                        }
                                        if true {
                                            v1[v3] = v3
                                            for fi81 := 0; fi81 < 12; fi81++ {
                                                switch fi81 {
                                                case 4294967295:
                                                    _ = len(v1)
                                                case 127:
                                                    var v82 int16 = (v54 | (^(v54)))
                                                    _ = v82
                                                case 1:
                                                    if (false || (!v33)) {
                                                        v4 = (v4 | v4)
                                                    } else {
                                                        v6 -= v6
                                                        v4 = 42
                                                    }
                                                default:
                                                    var v83 uint32 = 271
                                                    _ = v83
                                                }
                                                var v84 int64 = -1
                                                _ = fi81
                                                _ = v84
                                            }
                                        }
                                    }
                                    _ = v77
                                }
                                _ = v54
                            }
                        }
                        _ = v48
                    }
                    _ = v47
                }
            }
            v85 := v1[v3]
            v86 := GenS1{F0: v33}
            _ = v85
            _ = v86
        }
        for fi87 := 0; fi87 < 13; fi87++ {
            var v88 int16 = 293
            var v89 uint16 = 678
            var v90 uint32 = (v3 << ((42) & 63))
            _ = fi87
            _ = v88
            _ = v89
            _ = v90
        }
    } else {
        v5 = 255
        v4 += (v4 | 18446744073709551615)
    }
    v91 := GenS4{F0: 1, F1: v4}
    if v33 {
        if v33 {
            v92 := []int{32768}
            var v93 int16 = 557
            _ = v92
            _ = v93
        }
        if true {
            fmt.Println(v33, v33, v6)
            v94 := GenS6{F0: 42}
            v5 -= 1
            _ = v94
        }
        for fi95 := 0; fi95 < 15; fi95++ {
            fmt.Println(v4, v2, v5)
            fmt.Println(v4, v33)
            v96 := GenS7{F0: 193, F1: v2}
            _ = fi95
            _ = v96
        }
    }
    var v97 uint32 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v91
    _ = v97
}
