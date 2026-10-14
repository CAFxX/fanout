package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16; F1 int16; F2 uint16 }
type GenS9 struct { F0 int64 }
func genFn0() int32 {
    var v4 int32 = 0
    v4 += v4
    v4 += -2147483648
    switch 0 {
    case 42:
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 uint8 = 1
            v7 := GenS1{F0: (v4 / ((v4) | 1)), F1: 4294967295}
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 uint64 = 42
                v10 := GenS1{F0: -2147483648, F1: (^(fi5))}
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi5
            _ = v6
            _ = v7
        }
    case 32767:
        fmt.Println(v4)
    default:
        var v11 int64 = 1
        _ = v11
    }
    var v12 uint64 = 4294967295
    _ = v12
    _ = v4
    return 42
}
func genFn1(p0 int16, p1 uint8, p2 uint) int8 {
    var v13 int = 0
    v14 := []int{v13, (v13 + (v13 & v13)), 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 46
}

func main() {
    var v1 int32 = 42
    var v2 uint16 = 1
    var v3 float64 = -2.25
    switch 42 {
    case 9223372036854775807:
        var v15 int32 = (v1 ^ v1)
        _ = v15
    case 1:
        fmt.Println(v2, v1, v2)
    case -995:
        var v16 float32 = float32(v3)
        _ = v16
    }
    v17 := GenS1{F0: (^(v1)), F1: 2147483647}
    if true {
        if (!(!(0 == 854))) {
            var v18 float64 = ((1.5 + 0.1) - -2.25)
            _ = v18
        }
        if true {
            v19 := GenS2{F0: 2147483648, F1: 127}
            v20 := []int{0}
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 13; fi21++ {
                v22 := GenS1{F0: 42, F1: (fi21 - (-(fi21)))}
                _ = fi21
                _ = v22
            }
        }
        v3 = v3
    } else {
        v1 -= -1
        switch "" {
        case "\n\t":
            fmt.Println(v2, v2)
        default:
            v23 := []int{1, 128}
            _ = v23
        }
    }
    v24 := []int{-2147483648, -9223372036854775808, -15}
    var v25 bool = (false && false)
    v26 := GenS1{F0: v1, F1: -2147483648}
    v25 = v25
    var v27 int32 = v1
    switch 589 {
    case 9223372036854775807:
        v25 = (42 <= 81)
    case 127:
        v28 := GenS1{F0: v27, F1: 611}
        _ = v28
    default:
        if (-32768 <= 1000) {
            switch -661 {
            case 0:
                switch ("foo bar" + "") {
                case "a":
                    switch "hello" {
                    case "\n\t":
                        v1 = v1
                    case "foo bar":
                        switch 1 {
                        case -106:
                            v2 = v2
                        case -9223372036854775808:
                            var v29 uint64 = 860
                            _ = v29
                        default:
                            v27 = genFn0()
                        }
                    case "a":
                        if len(v24) > 0 {
                            v24[0] = int(v2)
                        }
                    default:
                        var v30 int16 = 0
                        _ = v30
                    }
                case "hello":
                    if (-9223372036854775808 > -75) {
                        if (!((-128 == 0) || v25)) {
                            var v31 int64 = 0
                            _ = v31
                        } else {
                            if (!v25) {
                                fmt.Println(v25, v27, v1)
                            }
                            var v32 float32 = (1e10 * 3.14159)
                            for fi33 := 0; fi33 < 15; fi33++ {
                                if v25 {
                                    var v34 int8 = 42
                                    _ = v34
                                }
                                if len(v24) > 0 {
                                    v24[0] = int(v2)
                                }
                                v35 := GenS2{F0: 2, F1: fi33}
                                _ = fi33
                                _ = v35
                            }
                            _ = v32
                        }
                        var v36 int64 = -438
                        _ = v36
                    } else {
                        for fi37 := 0; fi37 < 12; fi37++ {
                            fmt.Println(v27, v25)
                            v38 := GenS3{F0: (305 != 0)}
                            _ = fi37
                            _ = v38
                        }
                        v1 = (v27 ^ genFn0())
                        if v25 {
                            var v39 int32 = v27
                            var v40 int64 = 1
                            _ = v39
                            _ = v40
                        } else {
                            v24 = append(v24, 127)
                            v1 += (^(v27))
                        }
                    }
                case "\n\t":
                    v41 := GenS4{F0: v27}
                    _ = v41
                default:
                    v42 := GenS1{F0: 2147483647, F1: -883}
                    _ = v42
                }
            case 9223372036854775807:
                v1 = 2147483647
            default:
                _ = len(v24)
            }
            v43 := GenS2{F0: 2147483647, F1: -9223372036854775808}
            switch -9223372036854775808 {
            case 42:
                var v44 int = 32767
                _ = v44
            case 9223372036854775807:
                var v45 uint32 = 823
                _ = v45
            case 2147483647:
                v27 -= v1
            default:
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
            }
            _ = v43
        }
    }
    v46 := GenS2{F0: uint32(v2), F1: 127}
    switch 9223372036854775807 {
    case 895:
        v27 += 0
    case 0:
        var v47 int = -789
        _ = v47
    case 2147483647:
        var v48 uint = 0
        _ = v48
    default:
        if ((v27 - int32(v2)) != (^(v1))) {
            for fi49 := 0; fi49 < 19; fi49++ {
                v24 = append(v24, (fi49 >> ((2) & 63)))
                fmt.Println(v1, v2)
                var v50 uint32 = 953
                _ = fi49
                _ = v50
            }
            if (!v25) {
                switch 42 {
                case 32768:
                    v51 := GenS2{F0: 70, F1: -1}
                    _ = v51
                default:
                    v2 -= v2
                }
                if len(v24) > 0 {
                    v24[0] = -9223372036854775808
                }
                switch "hello" {
                case "a":
                    switch "" {
                    case "a":
                        v1 = (v1 & (v1 - v27))
                    case "":
                        v1 = (v1 >> ((42) & 63))
                    case "\n\t":
                        v52 := GenS4{F0: (v27 % (((v27 & (v1 & (v27 ^ v1)))) | 1))}
                        _ = v52
                    }
                case "":
                    var v53 int64 = 9223372036854775807
                    _ = v53
                case "hello":
                    v1 += (^(v1))
                default:
                    var v54 uint16 = (v2 % (((v2 * (v2 ^ v2))) | 1))
                    _ = v54
                }
            }
        } else {
            fmt.Println(v3)
            fmt.Println(v2)
            fmt.Println(v3)
        }
    }
    switch 127 {
    case 0:
        v55 := GenS2{F0: 2, F1: -1}
        _ = v55
    case -1:
        var v56 int32 = v1
        _ = v56
    case 452:
        if (357 != 1) {
            if (-32768 == -32768) {
                var v57 int32 = v27
                switch "hello" {
                case "":
                    if v25 {
                        switch "" {
                        case "a":
                            switch "" {
                            case "":
                                v58 := GenS5{F0: (v27 ^ (v1 / ((genFn0()) | 1))), F1: 1}
                                _ = v58
                            }
                        case "":
                            v1 -= 42
                        default:
                            v59 := GenS6{F0: 2, F1: (v27 * v57)}
                            _ = v59
                        }
                    } else {
                        var v60 int64 = 0
                        _ = v60
                    }
                case "a":
                    var v61 int32 = 42
                    _ = v61
                case "hello":
                    for fi62 := 0; fi62 < 7; fi62++ {
                        var v63 float32 = 1.5
                        _ = fi62
                        _ = v63
                    }
                default:
                    if v25 {
                        var v64 float32 = float32(v3)
                        v27 = genFn0()
                        switch -101 {
                        case 4294967295:
                            var v65 string = "a"
                            _ = v65
                        default:
                            v66 := GenS2{F0: 65535, F1: 127}
                            _ = v66
                        }
                        _ = v64
                    } else {
                        if len(v24) > 0 {
                            v24[0] = 311
                        }
                    }
                }
                _ = v57
            } else {
                if v25 {
                    v25 = v25
                } else {
                    fmt.Println(v1, v1)
                    v24 = append(v24, 9223372036854775807)
                }
                for fi67 := 0; fi67 < 13; fi67++ {
                    v1 += -403
                    v68 := GenS1{F0: (v1 / ((v27) | 1)), F1: fi67}
                    var v69 string = (((((("a" + "\n\t") + "a") + "") + "a") + "") + "a")
                    _ = fi67
                    _ = v68
                    _ = v69
                }
            }
        }
    default:
        if v25 {
            for fi70 := 0; fi70 < 6; fi70++ {
                switch (^(fi70)) {
                case 2147483647:
                    switch ("foo bar" + ("" + ("foo bar" + "\n\t"))) {
                    case "a":
                        if ((true && false) || ((0 < 42) && false)) {
                            v2 -= (v2 % ((v2) | 1))
                            v71 := GenS6{F0: 111, F1: 2147483647}
                            _ = v71
                        } else {
                            v72 := GenS7{F0: (v1 / ((genFn0()) | 1))}
                            if len(v24) > 0 {
                                v24[0] = (fi70 & (fi70 ^ (fi70 | -1)))
                            }
                            _ = v72
                        }
                    case "hello":
                        if ((0.0 + (float64(0.0) * (-2.25 * (float64(1e10) + (((float64((float32(-2.25) + float32(v3))) / ((-2.25 + v3) / (float64(float32(v3)) + v3))) - v3) / v3))))) <= (v3 - -2.25)) {
                            v73 := GenS1{F0: (v1 ^ v27), F1: (fi70 & -689)}
                            var v74 float32 = (1e10 * 1e10)
                            v24 = append(v24, -2147483648)
                            _ = v73
                            _ = v74
                        }
                    default:
                        var v75 int16 = -557
                        _ = v75
                    }
                default:
                    v24 = append(v24, 195)
                }
                v1 += (v27 & (v1 % ((genFn0()) | 1)))
                _ = fi70
            }
            switch "hello" {
            case "\n\t":
                v27 = genFn0()
            case "hello":
                if (116 >= 910) {
                    _ = len(v24)
                    fmt.Println(v2, v3)
                }
            case "a":
                v3 += v3
            default:
                switch -247 {
                case 42:
                    v3 -= (0.1 / (-2.25 * v3))
                case -1:
                    v76 := GenS8{F0: 1, F1: -515, F2: 65535}
                    _ = v76
                case -924:
                    var v77 uint32 = 2
                    _ = v77
                }
            }
            v78 := GenS9{F0: -363}
            _ = v78
        } else {
            v79 := GenS7{F0: 2147483647}
            v25 = v25
            var v80 uint64 = 1
            _ = v79
            _ = v80
        }
    }
    var v81 int16 = -547
    v24 = append(v24, 252)
    var v82 float64 = (((-2.25 / v3) - 1e10) + ((0.1 * v3) - 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v46
    _ = v81
    _ = v82
}
