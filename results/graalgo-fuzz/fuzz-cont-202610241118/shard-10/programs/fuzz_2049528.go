package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int64; F2 uint64 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint; F1 int64; F2 int8 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int; F1 int; F2 uint }
type GenS11 struct { F0 bool; F1 uint16 }
type GenS12 struct { F0 string }
func genFn0() uint64 {
    v5 := GenS1{F0: 423}
    if true {
        var v6 float64 = 1e10
        _ = v6
    } else {
        if (-133 >= -32768) {
            for fi7 := 0; fi7 < 4; fi7++ {
                for fi8 := 0; fi8 < 11; fi8++ {
                    var v9 uint32 = 2147483648
                    var v10 float32 = -2.25
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi7
            }
        }
    }
    v11 := []int{-1, -2147483648, 32768, -2147483648}
    _ = v11
    _ = v5
    return 2
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 string = ("" + "a")
    var v4 uint = 0
    var v12 uint16 = 157
    v4 -= uint(v12)
    var v13 uint16 = v12
    var v14 float32 = (0.0 * 1.5)
    v12 = v12
    v13 = 1
    for fi15 := 0; fi15 < 6; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 uint8 = 255
            _ = fi16
            _ = v17
        }
        var v18 int8 = 42
        v14 = 0.1
        _ = fi15
        _ = v18
    }
    v19 := []int{128, -768, 32768}
    if true {
        v14 = 0.1
    } else {
        var v20 int = 1
        fmt.Println(v20)
        _ = v20
    }
    if ((false || true) || true) {
        var v21 int8 = 1
        if false {
            v3 = v3
        } else {
            var v22 int = 4294967295
            for fi23 := 0; fi23 < 12; fi23++ {
                if true {
                    v1[2147483647] = v22
                } else {
                    v13 = v13
                    var v24 int64 = -1
                    v13 += v13
                    _ = v24
                }
                for fi25 := 0; fi25 < 7; fi25++ {
                    if (!(2 != 41)) {
                        if true {
                            fmt.Println(v3)
                            var v26 uint16 = v13
                            switch (v22 ^ fi25) {
                            case -466:
                                v27 := GenS1{F0: 42}
                                _ = v27
                            case 814:
                                var v28 uint32 = 4294967295
                                _ = v28
                            case 1:
                                var v29 uint = (v4 ^ 18446744073709551615)
                                _ = v29
                            }
                            if true {
                                if len(v19) > 0 {
                                    v19[0] = (v22 ^ (v22 - fi25))
                                }
                            } else {
                                v14 = float32(0.0)
                                fmt.Println(v3, v13)
                                var v30 float32 = v14
                                var v31 int32 = 0
                                var v32 int8 = 100
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            switch fi25 {
                            case 128:
                                var v33 uint64 = 2
                                _ = v33
                            }
                            _ = v26
                        } else {
                            switch "a" {
                            case "":
                                v12 += v13
                                var v34 float32 = (v14 + 3.14159)
                                switch fi23 {
                                case 1:
                                    if true {
                                        if true {
                                            v12 -= 32767
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                            v13 -= (v12 >> ((2) & 63))
                                        } else {
                                            v21 = v21
                                        }
                                        v35 := GenS2{F0: 0, F1: 42, F2: 2}
                                        _ = v35
                                    } else {
                                        var v36 int = fi25
                                        fmt.Println(fi23)
                                        _ = v36
                                    }
                                    var v37 uint = 1
                                    var v38 uint8 = 255
                                    if (!(!true)) {
                                        switch v3 {
                                        case "\n\t":
                                            var v39 int32 = int32(v37)
                                            _ = v39
                                        case "":
                                            v40 := GenS3{F0: (v22 % ((32767) | 1)), F1: 42}
                                            _ = v40
                                        case "a":
                                            switch v3 {
                                            case "foo bar":
                                                var v41 int = -369
                                                _ = v41
                                            case "":
                                                v12 -= v13
                                            default:
                                                if (!(true || (-1 == 42))) {
                                                    fmt.Println(v37)
                                                    if true {
                                                        var v42 int32 = 0
                                                        _ = v42
                                                    } else {
                                                        v13 = 255
                                                        switch fi23 {
                                                        case 0:
                                                            v4 = (v4 | v37)
                                                        case 32768:
                                                            if false {
                                                                v43 := GenS4{F0: (v37 % ((v4) | 1)), F1: -1, F2: (^(v21))}
                                                                _ = v43
                                                            } else {
                                                                v19 = append(v19, (fi25 % ((fi25) | 1)))
                                                                switch 4294967295 {
                                                                case 2147483647:
                                                                    if len(v19) > 0 {
                                                                        v19[0] = (v22 * fi23)
                                                                    }
                                                                case 127:
                                                                    v21 = v21
                                                                default:
                                                                    if true {
                                                                        var v44 int16 = int16(v13)
                                                                        _ = v44
                                                                    } else {
                                                                        if (42 == v21) {
                                                                            v14 = ((((v14 / v14) - v34) - v34) / (((v14 * v34) * v14) + 1e10))
                                                                        } else {
                                                                            v19 = append(v19, fi25)
                                                                        }
                                                                        switch (fi23 ^ fi23) {
                                                                        case 362:
                                                                            var v45 int32 = 33
                                                                            _ = v45
                                                                        case 952:
                                                                            switch fi23 {
                                                                            case -300:
                                                                                switch "\n\t" {
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        fmt.Println(v14, v13)
                                                                                        switch fi23 {
                                                                                        case 2147483647:
                                                                                            if len(v19) > 0 {
                                                                                                v19[0] = (^(fi23))
                                                                                            }
                                                                                        case 4294967295:
                                                                                            fmt.Println(fi25, v12, v34)
                                                                                        }
                                                                                        v4 += (v4 - (^(v4)))
                                                                                        v46 := GenS2{F0: 42, F1: -1, F2: 18446744073709551615}
                                                                                        _ = v46
                                                                                    } else {
                                                                                        v13 = (v13 << ((65535) & 63))
                                                                                    }
                                                                                case "hello":
                                                                                    if (-1 > -1) {
                                                                                        switch "\n\t" {
                                                                                        case "\n\t":
                                                                                            if false {
                                                                                                v47 := GenS2{F0: 0, F1: int64(v38), F2: 42}
                                                                                                _ = v47
                                                                                            } else {
                                                                                                v14 = (v34 / v14)
                                                                                                fmt.Println(v34, v14)
                                                                                                if (true && true) {
                                                                                                    if false {
                                                                                                        _ = len(v1)
                                                                                                    }
                                                                                                    var v48 uint = 4294967295
                                                                                                    v49 := GenS1{F0: 65535}
                                                                                                    _ = v48
                                                                                                    _ = v49
                                                                                                } else {
                                                                                                    if ((false && (((-32768 <= 42) || ((v38 % (((v38 % ((v38) | 1))) | 1)) != uint8(v4))) || true)) || (0 < 2147483647)) {
                                                                                                        if (v21 >= 100) {
                                                                                                            v50 := GenS5{F0: 0, F1: v12}
                                                                                                            fmt.Println(v37)
                                                                                                            if false {
                                                                                                                var v51 int32 = 2147483647
                                                                                                                _ = v51
                                                                                                            } else {
                                                                                                                var v52 uint32 = 940
                                                                                                                _ = v52
                                                                                                            }
                                                                                                            _ = v50
                                                                                                        } else {
                                                                                                            fmt.Println(v21, v12)
                                                                                                        }
                                                                                                        var v53 int32 = 1
                                                                                                        v54 := GenS6{F0: (v38 * v38), F1: fi25}
                                                                                                        fmt.Println(v4, v22)
                                                                                                        _ = v53
                                                                                                        _ = v54
                                                                                                    } else {
                                                                                                        v55 := GenS7{F0: (v38 / (((v38 % ((255) | 1))) | 1)), F1: 625, F2: 1}
                                                                                                        var v56 float64 = 1.5
                                                                                                        switch (v3 + v3) {
                                                                                                        case "hello":
                                                                                                            _ = len(v2)
                                                                                                        }
                                                                                                        var v57 int8 = (^(v21))
                                                                                                        fmt.Println(v3, v12)
                                                                                                        _ = v55
                                                                                                        _ = v56
                                                                                                        _ = v57
                                                                                                    }
                                                                                                    switch (v3 + (v3 + v3)) {
                                                                                                    case "":
                                                                                                        fmt.Println(v34, v12)
                                                                                                    case "hello":
                                                                                                        fmt.Println(v38, fi25, v21)
                                                                                                    case "\n\t":
                                                                                                        if false {
                                                                                                            switch fi25 {
                                                                                                            case 9223372036854775807:
                                                                                                                switch v3 {
                                                                                                                case "\n\t":
                                                                                                                    v58 := GenS1{F0: v37}
                                                                                                                    _ = v58
                                                                                                                case "":
                                                                                                                    var v59 uint32 = 411
                                                                                                                    _ = v59
                                                                                                                case "a":
                                                                                                                    v14 = 1.5
                                                                                                                }
                                                                                                            case 1:
                                                                                                                v34 = v34
                                                                                                            }
                                                                                                            var v60 int64 = 42
                                                                                                            _ = v60
                                                                                                        } else {
                                                                                                            switch v3 {
                                                                                                            case "foo bar":
                                                                                                                var v61 float32 = 0.0
                                                                                                                _ = v61
                                                                                                            case "hello":
                                                                                                                var v62 string = v3
                                                                                                                _ = v62
                                                                                                            case "a":
                                                                                                                v37 += 2
                                                                                                            default:
                                                                                                                v21 -= v21
                                                                                                            }
                                                                                                            v4 = v37
                                                                                                            v12 = 65535
                                                                                                        }
                                                                                                        _ = len(v1)
                                                                                                    default:
                                                                                                        if true {
                                                                                                            var v63 uint8 = (v38 | v38)
                                                                                                            v34 = v34
                                                                                                            _ = v63
                                                                                                        }
                                                                                                    }
                                                                                                    if false {
                                                                                                        v3 = v3
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        case "hello":
                                                                                            var v64 float64 = 0.0
                                                                                            _ = v64
                                                                                        case "foo bar":
                                                                                            fmt.Println(v13)
                                                                                        }
                                                                                        var v65 uint8 = (^(v38))
                                                                                        v14 = 3.14159
                                                                                        v66 := GenS8{F0: v21}
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                    } else {
                                                                                        var v67 float64 = 1e10
                                                                                        var v68 uint8 = 1
                                                                                        _ = v67
                                                                                        _ = v68
                                                                                    }
                                                                                case "a":
                                                                                    v14 -= v34
                                                                                default:
                                                                                    v19 = append(v19, 32768)
                                                                                }
                                                                            case 2147483647:
                                                                                var v69 float32 = v14
                                                                                _ = v69
                                                                            }
                                                                        case 1:
                                                                            fmt.Println(v3, fi23, v34)
                                                                        default:
                                                                            var v70 uint = v37
                                                                            _ = v70
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        case -1:
                                                            if (true || (v13 > v12)) {
                                                                v3 = v3
                                                            }
                                                            if false {
                                                                v71 := GenS1{F0: v4}
                                                                _ = v71
                                                            } else {
                                                                var v72 int = fi25
                                                                if len(v19) > 0 {
                                                                    v19[0] = fi25
                                                                }
                                                                if true {
                                                                    var v73 int = (fi23 % (((v72 % ((-9223372036854775808) | 1))) | 1))
                                                                    v34 += 0.1
                                                                    _ = len(v1)
                                                                    _ = v73
                                                                }
                                                                _ = v72
                                                            }
                                                        default:
                                                            fmt.Println(v12)
                                                            var v74 uint8 = 15
                                                            v75 := GenS2{F0: -833, F1: 9223372036854775807, F2: 876}
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                    }
                                                    var v76 int32 = 42
                                                    _ = v76
                                                } else {
                                                    v12 += 42
                                                }
                                            }
                                        }
                                    } else {
                                        if (687 > -1) {
                                            if (9223372036854775807 == 892) {
                                                v14 -= 1e10
                                            } else {
                                                fmt.Println(v13, v3)
                                            }
                                        } else {
                                            if ((!(false && (true || (v4 < 0)))) && ((1 <= 862) || false)) {
                                                v77 := GenS4{F0: (v37 << ((v4) & 63)), F1: -1, F2: v21}
                                                _ = v77
                                            } else {
                                                switch v22 {
                                                case 434:
                                                    v78 := GenS1{F0: v37}
                                                    _ = v78
                                                case 42:
                                                    if len(v19) >= 2 {
                                                        v19 = v19[1:len(v19)]
                                                    }
                                                default:
                                                    switch "" {
                                                    case "\n\t":
                                                        v21 += v21
                                                    }
                                                }
                                            }
                                            fmt.Println(v38, fi23, v4)
                                        }
                                    }
                                    v3 = "\n\t"
                                    _ = v37
                                    _ = v38
                                case 127:
                                    var v79 int64 = -9223372036854775808
                                    _ = v79
                                case 546:
                                    switch v3 {
                                    case "":
                                        v3 = (v3 + v3)
                                    }
                                default:
                                    v80 := GenS2{F0: 1000, F1: 1, F2: 18446744073709551615}
                                    _ = v80
                                }
                                _ = v34
                            case "hello":
                                if (!false) {
                                    if true {
                                        if (1000 >= 839) {
                                            var v81 uint16 = 42
                                            _ = len(v1)
                                            if (-32768 != 1000) {
                                                if false {
                                                    switch v3 {
                                                    case "foo bar":
                                                        fmt.Println(v14)
                                                    case "\n\t":
                                                        fmt.Println(v81)
                                                    case "a":
                                                        if false {
                                                            v21 = int8(fi23)
                                                            delete(v1, -9223372036854775808)
                                                            if (18446744073709551615 > 457) {
                                                                v14 += v14
                                                                var v82 uint16 = (v13 % (((v13 >> ((v4) & 63))) | 1))
                                                                v83 := GenS2{F0: 42, F1: -668, F2: 496}
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                        } else {
                                                            v4 -= (v4 / ((42) | 1))
                                                            _ = len(v2)
                                                        }
                                                    }
                                                    var v84 float32 = v14
                                                    _ = v84
                                                } else {
                                                    var v85 int16 = 0
                                                    if len(v19) >= 2 {
                                                        v19 = v19[1:len(v19)]
                                                    }
                                                    if len(v19) > 0 {
                                                        v19[0] = (v22 << ((v4) & 63))
                                                    }
                                                    _ = v85
                                                }
                                                v86 := GenS5{F0: 0, F1: (v81 - v13)}
                                                v87 := GenS9{F0: 2147483647}
                                                _ = v86
                                                _ = v87
                                            } else {
                                                v4 -= (v4 ^ (v4 >> ((472) & 63)))
                                            }
                                            _ = v81
                                        }
                                        var v88 int16 = 1000
                                        if (!(500 == 18446744073709551615)) {
                                            v88 -= -392
                                        } else {
                                            var v89 int32 = -498
                                            _ = v89
                                        }
                                        _ = v88
                                    }
                                    v19 = append(v19, fi25)
                                    fmt.Println(v22, v21)
                                    fmt.Println(fi23)
                                    v14 = ((3.14159 * 0.1) / v14)
                                } else {
                                    if true {
                                        delete(v1, int(v21))
                                    }
                                    v90 := GenS7{F0: 0, F1: 892, F2: 0}
                                    v4 = (^(v4))
                                    _ = v90
                                }
                            case "\n\t":
                                if (float32(0.1) > v14) {
                                    fmt.Println(v21, v13, v13)
                                    v22 += (v22 << ((v4) & 63))
                                    v91 := GenS9{F0: 974}
                                    _ = v91
                                }
                            }
                            fmt.Println(fi23, v3, v13)
                        }
                        fmt.Println(v4, v21)
                        switch "\n\t" {
                        case "\n\t":
                            v22 -= 32767
                        default:
                            if false {
                                v13 = (v13 ^ v13)
                                switch v3 {
                                case "foo bar":
                                    v19 = append(v19, -9223372036854775808)
                                default:
                                    fmt.Println(fi25)
                                    var v92 float32 = v14
                                    v93 := GenS10{F0: fi23, F1: v22, F2: 2}
                                    _ = v92
                                    _ = v93
                                }
                                var v94 int32 = 839
                                _ = v94
                            } else {
                                var v95 uint8 = 128
                                if (fi23 > v22) {
                                    v95 -= (v95 - v95)
                                    v96 := v1[v22]
                                    switch v3 {
                                    case "a":
                                        var v97 int = 128
                                        _ = v97
                                    case "\n\t":
                                        if true {
                                            var v98 uint16 = 42
                                            v3 = v3
                                            _ = v98
                                        } else {
                                            v99 := GenS10{F0: (v22 / ((v96) | 1)), F1: fi23, F2: (v4 ^ 820)}
                                            switch v22 {
                                            case -46:
                                                var v100 int8 = (v21 | -128)
                                                switch v3 {
                                                case "a":
                                                    var v101 uint64 = genFn0()
                                                    _ = v101
                                                case "foo bar":
                                                    switch v22 {
                                                    case 9223372036854775807:
                                                        fmt.Println(v95, v12, v21)
                                                    case 64:
                                                        v102 := GenS2{F0: 32767, F1: -9223372036854775808, F2: 0}
                                                        _ = v102
                                                    case 32768:
                                                        v19 = append(v19, 32767)
                                                    default:
                                                        switch -719 {
                                                        case -242:
                                                            switch v96 {
                                                            case 127:
                                                                v103 := GenS9{F0: 805}
                                                                _ = v103
                                                            case 4294967295:
                                                                var v104 uint64 = 18446744073709551615
                                                                _ = v104
                                                            case -387:
                                                                if len(v19) > 0 {
                                                                    v19[0] = v22
                                                                }
                                                            default:
                                                                v2[v3] = ("foo bar" + (("" + "foo bar") + v3))
                                                            }
                                                        default:
                                                            var v105 bool = (762 > 42)
                                                            _ = v105
                                                        }
                                                        if (v12 != v12) {
                                                            switch v3 {
                                                            case "\n\t":
                                                                v106 := v1[fi23]
                                                                _ = v106
                                                            case "a":
                                                                var v107 int16 = -589
                                                                _ = v107
                                                            case "":
                                                                switch v3 {
                                                                case "":
                                                                    var v108 uint = 176
                                                                    _ = v108
                                                                case "hello":
                                                                    var v109 int16 = 147
                                                                    _ = v109
                                                                case "foo bar":
                                                                    v12 = v13
                                                                }
                                                            default:
                                                                var v110 int32 = -1
                                                                _ = v110
                                                            }
                                                            v111 := GenS11{F0: (!true), F1: (v13 - v13)}
                                                            var v112 int64 = -9223372036854775808
                                                            _ = v111
                                                            _ = v112
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(v12, fi25, fi23)
                                                }
                                                _ = len(v19)
                                                _ = v100
                                            case 42:
                                                var v113 float64 = -2.25
                                                _ = v113
                                            case 127:
                                                var v114 uint16 = v13
                                                _ = v114
                                            default:
                                                if (719 <= uint64(v95)) {
                                                    v4 += v4
                                                }
                                            }
                                            _ = v99
                                        }
                                    case "hello":
                                        v115 := GenS8{F0: -8}
                                        _ = v115
                                    }
                                    _ = v96
                                }
                                v13 = uint16(v95)
                                _ = v95
                            }
                        }
                    }
                    if true {
                        if (-2.25 > (v14 + v14)) {
                            if (true && (218 < genFn0())) {
                                v116 := GenS12{F0: v3}
                                var v117 int32 = -1
                                _ = v116
                                _ = v117
                            } else {
                                fmt.Println(v14, v3)
                            }
                            v3 = "a"
                        } else {
                            v118 := GenS4{F0: 850, F1: -9223372036854775808, F2: (v21 + v21)}
                            var v119 uint64 = 0
                            _ = v118
                            _ = v119
                        }
                    } else {
                        if true {
                            if len(v19) > 0 {
                                v19[0] = (^(v22))
                            }
                            v21 += (v21 / ((127) | 1))
                            switch fi23 {
                            case 678:
                                if (127 < 42) {
                                    v120 := GenS11{F0: false, F1: v13}
                                    if false {
                                        _ = len(v2)
                                    }
                                    var v121 int16 = 1
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v122 uint64 = 0
                                        _ = v122
                                    case "hello":
                                        v19 = append(v19, fi25)
                                    }
                                    _ = v120
                                    _ = v121
                                } else {
                                    v12 = (v13 * v12)
                                    var v123 float64 = ((0.0 + 1.5) * float64(((((v14 + float32(1e10)) / v14) / 3.14159) - 3.14159)))
                                    _ = v123
                                }
                            case 9223372036854775807:
                                v2[v3] = v3
                            }
                        } else {
                            if (!(false || false)) {
                                var v124 uint = (^(v4))
                                var v125 string = v3
                                var v126 uint32 = 1
                                _ = v124
                                _ = v125
                                _ = v126
                            } else {
                                switch ((v3 + v3) + "hello") {
                                case "":
                                    var v127 string = v3
                                    _ = v127
                                case "hello":
                                    switch (fi25 + -924) {
                                    case 32:
                                        switch (v3 + v3) {
                                        case "hello":
                                            v19 = append(v19, fi23)
                                            var v128 int8 = (v21 << ((v4) & 63))
                                            _ = len(v19)
                                            _ = v128
                                        default:
                                            v129 := GenS10{F0: (fi23 % ((v22) | 1)), F1: (v22 | -2147483648), F2: uint(v12)}
                                            _ = v129
                                        }
                                    case -178:
                                        var v130 int16 = -619
                                        _ = v130
                                    case -2147483648:
                                        _ = len(v19)
                                    }
                                case "a":
                                    switch -658 {
                                    case 4294967295:
                                        v131 := GenS9{F0: 4294967295}
                                        _ = v131
                                    }
                                default:
                                    if false {
                                        var v132 string = (((((v3 + (("\n\t" + "") + v3)) + "") + ((v3 + v3) + "")) + v3) + v3)
                                        v14 = 1.5
                                        _ = v132
                                    }
                                }
                                v12 -= v13
                                var v133 uint8 = 0
                                v14 -= float32(3.14159)
                                _ = v133
                            }
                            var v134 int = 32768
                            _ = v134
                        }
                        var v135 uint32 = 325
                        _ = v135
                    }
                    v3 = ""
                    _ = fi25
                }
                _ = fi23
            }
            _ = v22
        }
        _ = v21
    }
    fmt.Println(v12)
    v14 = float32(-2.25)
    var v136 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v136:", v136)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v136
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
