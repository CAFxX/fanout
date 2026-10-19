package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint16; F1 bool; F2 int64 }
type GenS6 struct { F0 uint8; F1 int; F2 int8 }
type GenS7 struct { F0 int32; F1 int8; F2 int8 }
type GenS8 struct { F0 uint16; F1 int16 }
type GenS9 struct { F0 uint; F1 int32 }
type GenS10 struct { F0 uint8; F1 int; F2 bool }
type GenS11 struct { F0 bool; F1 bool; F2 uint64 }
func genFn0(p0 bool, p1 int64, p2 uint8) int {
    p1 -= (^(p1))
    for fi6 := 0; fi6 < 9; fi6++ {
        switch fi6 {
        case -272:
            v7 := GenS1{F0: "hello", F1: 1, F2: (p1 >> ((516) & 63))}
            _ = v7
        }
        _ = fi6
    }
    v8 := GenS1{F0: "\n\t", F1: 42, F2: (p1 >> ((42) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -705
}

func main() {
    var v1 uint8 = 42
    v2 := []string{(("" + ("a" + "")) + "foo bar")}
    var v3 float32 = 0.0
    var v4 string = ("" + ("a" + "hello"))
    v5 := []uint32{42, 295}
    var v9 int8 = -128
    var v10 uint16 = 32767
    switch v4 {
    case "\n\t":
        v4 = ("" + (v4 + v4))
    }
    var v11 uint64 = 2
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 float64 = 1.5
        fmt.Println(v4, fi12, v3)
        v1 += v1
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 float64 = 0.1
        for fi16 := 0; fi16 < 10; fi16++ {
            v1 = (v1 & (v1 * (v1 >> ((370) & 63))))
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    v1 -= v1
    switch genFn0(true, -9223372036854775808, v1) {
    case 450:
        var v17 bool = (true && false)
        _ = v17
    case -308:
        v4 = v4
    }
    if false {
        v3 += 3.14159
        if false {
            fmt.Println(v1, v9, v1)
            v11 -= (v11 / (((v11 * (v11 + 42))) | 1))
            v3 = (float32(-2.25) / v3)
        } else {
            _ = len(v5)
        }
    } else {
        switch 9223372036854775807 {
        case 0:
            for fi18 := 0; fi18 < 9; fi18++ {
                for fi19 := 0; fi19 < 20; fi19++ {
                    v20 := GenS2{F0: 207, F1: 2147483647, F2: 1}
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 8; fi21++ {
                    v3 -= float32((float64(v3) * (1e10 * (0.0 + -2.25))))
                    _ = fi21
                }
                _ = fi18
            }
        case -698:
            for fi22 := 0; fi22 < 10; fi22++ {
                fmt.Println(v9, v11)
                _ = fi22
            }
        case 32767:
            var v23 bool = (!(!true))
            _ = v23
        default:
            v2 = append(v2, ((v4 + ("a" + v4)) + v4))
        }
    }
    if (!(!((^(v11)) != (v11 % ((0) | 1))))) {
        if (genFn0(false, -938, v1) <= 2147483647) {
            if ((((v4 + ((v4 + ("\n\t" + v4)) + ("\n\t" + v4))) + v4) > ("a" + "")) || true) {
                v2 = append(v2, v4)
                fmt.Println(v9, v10)
            }
            if false {
                if false {
                    switch 297 {
                    case -315:
                        if (((float32((3.14159 + 3.14159)) / 1.0) / 1.0) == 3.14159) {
                            if ((v4 > "hello") || false) {
                                fmt.Println(v9, v1, v9)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch -1 {
                                case -9223372036854775808:
                                    for fi24 := 0; fi24 < 17; fi24++ {
                                        var v25 bool = (false && false)
                                        var v26 uint16 = (v10 % ((v10) | 1))
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                case 4294967295:
                                    v3 = 1.5
                                case 9223372036854775807:
                                    v2 = append(v2, v4)
                                }
                            }
                            v27 := GenS2{F0: 2, F1: 257, F2: 458}
                            if len(v2) > 0 {
                                v2[0] = v4
                            }
                            _ = v27
                        } else {
                            v28 := GenS1{F0: (v4 + "hello"), F1: 76, F2: -1}
                            _ = v28
                        }
                    case 405:
                        var v29 int = int(v10)
                        _ = v29
                    case 32767:
                        switch genFn0(false, 0, (^(v1))) {
                        case -1:
                            var v30 float32 = (3.14159 / v3)
                            _ = v30
                        case 4294967295:
                            for fi31 := 0; fi31 < 19; fi31++ {
                                fmt.Println(fi31)
                                _ = fi31
                            }
                        default:
                            fmt.Println(v1, v3)
                        }
                    default:
                        for fi32 := 0; fi32 < 2; fi32++ {
                            var v33 int32 = -1
                            var v34 uint16 = (v10 | v10)
                            if len(v2) > 0 {
                                v2[0] = (v4 + (("a" + ("a" + v4)) + v4))
                            }
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    }
                }
                for fi35 := 0; fi35 < 17; fi35++ {
                    var v36 int64 = 232
                    _ = fi35
                    _ = v36
                }
                fmt.Println(v11)
            }
        } else {
            v3 = ((0.1 - (1.5 / float32((3.14159 + (-2.25 / -2.25))))) * v3)
        }
        if false {
            switch ((((v4 + "") + v4) + v4) + v4) {
            case "a":
                for fi37 := 0; fi37 < 1; fi37++ {
                    switch v4 {
                    case "\n\t":
                        if false {
                            v1 -= v1
                            var v38 int = genFn0((fi37 == 1), 489, 189)
                            _ = len(v5)
                            _ = v38
                        } else {
                            if len(v2) > 0 {
                                v2[0] = v4
                            }
                            for fi39 := 0; fi39 < 10; fi39++ {
                                var v40 uint8 = (v1 >> ((255) & 63))
                                v2 = append(v2, (v4 + ""))
                                fmt.Println(v11, v40)
                                _ = fi39
                                _ = v40
                            }
                            if (!true) {
                                var v41 int64 = 0
                                var v42 int8 = 1
                                _ = v41
                                _ = v42
                            } else {
                                switch genFn0(false, -9223372036854775808, 128) {
                                case 127:
                                    if true {
                                        if false {
                                            for fi43 := 0; fi43 < 15; fi43++ {
                                                var v44 float32 = v3
                                                if false {
                                                    v11 -= v11
                                                    v44 = 0.1
                                                }
                                                v45 := GenS1{F0: v4, F1: v11, F2: -1}
                                                _ = fi43
                                                _ = v44
                                                _ = v45
                                            }
                                            if len(v5) > 0 {
                                                v5[0] = 4294967295
                                            }
                                            for fi46 := 0; fi46 < 6; fi46++ {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = fi46
                                            }
                                        } else {
                                            if (214 == v1) {
                                                for fi47 := 0; fi47 < 17; fi47++ {
                                                    fmt.Println(fi47, v10)
                                                    _ = fi47
                                                }
                                                for fi48 := 0; fi48 < 10; fi48++ {
                                                    fmt.Println(v9, v1)
                                                    var v49 bool = false
                                                    v50 := GenS2{F0: 632, F1: 42, F2: 2147483647}
                                                    _ = fi48
                                                    _ = v49
                                                    _ = v50
                                                }
                                            } else {
                                                if false {
                                                    var v51 string = v4
                                                    for fi52 := 0; fi52 < 9; fi52++ {
                                                        v9 = int8(fi37)
                                                        _ = fi52
                                                    }
                                                    v51 = "a"
                                                    _ = v51
                                                } else {
                                                    var v53 uint64 = 42
                                                    v11 -= (^(v11))
                                                    _ = v53
                                                }
                                                v54 := GenS3{F0: v10}
                                                if (!(42 != 661)) {
                                                    var v55 uint32 = 2147483647
                                                    _ = v55
                                                }
                                                _ = v54
                                            }
                                            v3 = v3
                                            _ = len(v2)
                                        }
                                        for fi56 := 0; fi56 < 16; fi56++ {
                                            if len(v5) > 0 {
                                                v5[0] = 253
                                            }
                                            _ = fi56
                                        }
                                    }
                                }
                                if len(v5) > 0 {
                                    v5[0] = 2147483647
                                }
                                var v57 string = v4
                                _ = v57
                            }
                        }
                    case "hello":
                        var v58 string = "a"
                        _ = v58
                    default:
                        var v59 int8 = v9
                        _ = v59
                    }
                    v60 := GenS3{F0: v10}
                    _ = fi37
                    _ = v60
                }
            case "\n\t":
                v61 := GenS3{F0: 42}
                _ = v61
            default:
                v10 = v10
            }
        } else {
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            switch v4 {
            case "foo bar":
                v4 = v4
            case "":
                for fi62 := 0; fi62 < 17; fi62++ {
                    v4 = v4
                    v10 += (v10 << ((18446744073709551615) & 63))
                    _ = fi62
                }
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    } else {
        v63 := GenS4{F0: (v9 ^ -1), F1: (^(v1)), F2: v10}
        v10 = v10
        _ = v63
    }
    for fi64 := 0; fi64 < 3; fi64++ {
        var v65 uint = 42
        v2 = append(v2, "foo bar")
        if true {
            for fi66 := 0; fi66 < 16; fi66++ {
                v67 := GenS5{F0: v10, F1: true, F2: 0}
                _ = fi66
                _ = v67
            }
            for fi68 := 0; fi68 < 20; fi68++ {
                if len(v2) > 0 {
                    v2[0] = v4
                }
                var v69 string = ("foo bar" + ("\n\t" + ((("\n\t" + (v4 + v4)) + v4) + (("a" + "foo bar") + v4))))
                _ = fi68
                _ = v69
            }
            v70 := GenS6{F0: v1, F1: (fi64 - fi64), F2: (v9 >> ((v65) & 63))}
            _ = v70
        } else {
            v71 := GenS7{F0: -2147483648, F1: v9, F2: (-(v9))}
            for fi72 := 0; fi72 < 4; fi72++ {
                var v73 int16 = int16(v1)
                _ = fi72
                _ = v73
            }
            switch genFn0(((false || true) && (!false)), -696, v1) {
            case 32767:
                switch fi64 {
                case -1:
                    fmt.Println(v10)
                default:
                    v74 := GenS5{F0: (v10 & (v10 | 65535)), F1: ((^(v10)) <= (v10 * (^(v10)))), F2: 0}
                    _ = v74
                }
            case 9223372036854775807:
                for fi75 := 0; fi75 < 17; fi75++ {
                    switch "hello" {
                    case "hello":
                        fmt.Println(v65, v3)
                    default:
                        switch v4 {
                        case "foo bar":
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = v4
                                }
                                if false {
                                    v76 := GenS8{F0: (v10 + 703), F1: 42}
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v76
                                }
                            }
                        case "":
                            var v77 int32 = 307
                            _ = v77
                        case "\n\t":
                            var v78 uint = (^(v65))
                            _ = v78
                        default:
                            var v79 bool = false
                            _ = v79
                        }
                        v3 = 1.5
                        if false {
                            var v80 uint32 = 0
                            _ = v80
                        } else {
                            v11 = 4294967295
                            v1 -= v1
                        }
                    }
                    var v81 int = 42
                    _ = fi75
                    _ = v81
                }
            case 407:
                if false {
                    if true {
                        switch 339 {
                        case 1:
                            var v82 string = v4
                            _ = v82
                        case 22:
                            var v83 int64 = 0
                            _ = v83
                        case -9223372036854775808:
                            for fi84 := 0; fi84 < 20; fi84++ {
                                v4 = (v4 + v4)
                                _ = fi84
                            }
                        default:
                            v10 = (v10 >> ((v65) & 63))
                        }
                    }
                }
            }
            _ = v71
        }
        _ = fi64
        _ = v65
    }
    fmt.Println(v10)
    switch 127 {
    case 4294967295:
        v85 := GenS9{F0: 2, F1: -1}
        _ = v85
    case 32767:
        v86 := GenS9{F0: 2, F1: 2147483647}
        _ = v86
    }
    switch "foo bar" {
    case "hello":
        for fi87 := 0; fi87 < 10; fi87++ {
            for fi88 := 0; fi88 < 1; fi88++ {
                if false {
                    var v89 int64 = -9223372036854775808
                    v5 = append(v5, 4294967295)
                    _ = v89
                }
                _ = fi88
            }
            _ = len(v2)
            var v90 int = (-(fi87))
            _ = fi87
            _ = v90
        }
    default:
        switch "" {
        case "hello":
            v91 := GenS4{F0: v9, F1: (v1 | 249), F2: v10}
            _ = v91
        case "\n\t":
            fmt.Println(v3, v3, v9)
        default:
            var v92 int64 = 9223372036854775807
            _ = v92
        }
    }
    switch -1 {
    case -1:
        v93 := GenS10{F0: 128, F1: genFn0((false || false), 1, v1), F2: false}
        _ = v93
    case -571:
        switch v4 {
        case "":
            if false {
                v94 := GenS1{F0: "\n\t", F1: v11, F2: -1}
                _ = v94
            } else {
                v95 := GenS6{F0: uint8(v10), F1: 128, F2: -128}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v95
            }
        default:
            var v96 uint16 = 1
            _ = v96
        }
    case 1:
        v5 = append(v5, uint32(v1))
    default:
        v97 := GenS11{F0: true, F1: ((v11 * v11) != v11), F2: v11}
        _ = v97
    }
    switch 42 {
    case 667:
        for fi98 := 0; fi98 < 13; fi98++ {
            fmt.Println(v9)
            for fi99 := 0; fi99 < 16; fi99++ {
                v3 -= (-2.25 / v3)
                _ = len(v5)
                _ = fi99
            }
            _ = fi98
        }
    case 9223372036854775807:
        fmt.Println(v9)
    case 2147483647:
        v2 = append(v2, "hello")
    }
    switch genFn0(false, 9223372036854775807, v1) {
    case 0:
        var v100 int8 = 0
        _ = v100
    case -9223372036854775808:
        var v101 int32 = -2147483648
        _ = v101
    case -83:
        var v102 float64 = -2.25
        _ = v102
    default:
        var v103 float64 = (1e10 + 1.5)
        _ = v103
    }
    switch v4 {
    case "":
        if len(v2) > 0 {
            v2[0] = v4
        }
    case "hello":
        for fi104 := 0; fi104 < 5; fi104++ {
            var v105 int8 = v9
            fmt.Println(v105, v11)
            _ = fi104
            _ = v105
        }
    default:
        _ = len(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
