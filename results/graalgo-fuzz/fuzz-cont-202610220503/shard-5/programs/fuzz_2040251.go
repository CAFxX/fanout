package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int16; F1 int32; F2 bool }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 int; F1 int64; F2 int8 }
type GenS10 struct { F0 uint16 }

func main() {
    var v1 uint = 65535
    v2 := make(map[string]bool)
    var v3 uint8 = 42
    v4 := make(map[int]string)
    var v5 uint16 = 2
    v1 = (v1 * (v1 - (v1 | 0)))
    v6 := []int{-760, 32768, -9223372036854775808}
    switch 32767 {
    case 1:
        v5 = 0
    case 480:
        var v7 int = 82
        _ = v7
    case -728:
        if true {
            v5 = (v5 - 979)
            v3 -= v3
        } else {
            for fi8 := 0; fi8 < 2; fi8++ {
                for fi9 := 0; fi9 < 17; fi9++ {
                    if (1.5 == (0.1 * 0.0)) {
                        if false {
                            fmt.Println(v1, v1)
                        } else {
                            v10 := GenS1{F0: v1}
                            fmt.Println(fi8, v1, fi8)
                            _ = v10
                        }
                    }
                    v11 := GenS1{F0: v1}
                    _ = fi9
                    _ = v11
                }
                switch fi8 {
                case 128:
                    var v12 int16 = 165
                    _ = v12
                case 9223372036854775807:
                    if len(v6) > 0 {
                        v6[0] = -275
                    }
                case 1:
                    if false {
                        var v13 uint16 = (v5 | 0)
                        v5 = (v5 % (((v13 >> ((0) & 63))) | 1))
                        fmt.Println(v3)
                        _ = v13
                    } else {
                        for fi14 := 0; fi14 < 11; fi14++ {
                            if len(v6) > 0 {
                                v6[0] = (^(fi14))
                            }
                            switch "\n\t" {
                            case "":
                                v2[""] = (v3 != (v3 | v3))
                            case "a":
                                var v15 uint8 = (v3 ^ (v3 * v3))
                                _ = v15
                            case "\n\t":
                                v1 = uint(v5)
                            }
                            _ = fi14
                        }
                        v1 = v1
                    }
                }
                _ = fi8
            }
            if false {
                switch ("" + ("hello" + "")) {
                case "hello":
                    var v16 int64 = -1
                    _ = v16
                case "":
                    if true {
                        v3 += v3
                    } else {
                        v17 := GenS2{F0: -1, F1: 4294967295}
                        v6 = append(v6, 127)
                        _ = v17
                    }
                default:
                    v1 -= (^(v1))
                }
                fmt.Println(v3, v1)
                for fi18 := 0; fi18 < 4; fi18++ {
                    var v19 uint32 = 65535
                    if false {
                        if (!(!false)) {
                            for fi20 := 0; fi20 < 12; fi20++ {
                                switch "hello" {
                                case "a":
                                    var v21 int32 = 1
                                    _ = v21
                                case "":
                                    _ = len(v4)
                                case "hello":
                                    fmt.Println(v19, fi18)
                                default:
                                    var v22 uint32 = v19
                                    _ = v22
                                }
                                var v23 bool = (false && true)
                                _ = fi20
                                _ = v23
                            }
                            v1 += 661
                        }
                    } else {
                        for fi24 := 0; fi24 < 5; fi24++ {
                            var v25 uint = v1
                            v25 -= 596
                            var v26 int = (^(fi18))
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        switch ((("foo bar" + "foo bar") + ("foo bar" + ("hello" + ""))) + "a") {
                        case "foo bar":
                            v27 := GenS2{F0: 0, F1: (v19 & (^(v19)))}
                            _ = v27
                        default:
                            for fi28 := 0; fi28 < 6; fi28++ {
                                v29 := GenS2{F0: 42, F1: v19}
                                _ = fi28
                                _ = v29
                            }
                        }
                        v19 += v19
                    }
                    _ = fi18
                    _ = v19
                }
            } else {
                if ((v3 >> (((v1 << ((18446744073709551615) & 63))) & 63)) == (v3 * v3)) {
                    switch "foo bar" {
                    case "hello":
                        if false {
                            v30 := GenS3{F0: int8(v1), F1: 42}
                            _ = v30
                        }
                    default:
                        v3 -= v3
                    }
                    v3 = v3
                    _ = len(v4)
                }
                for fi31 := 0; fi31 < 2; fi31++ {
                    var v32 int = 42
                    var v33 uint8 = 113
                    for fi34 := 0; fi34 < 15; fi34++ {
                        v33 = v33
                        v35 := GenS1{F0: v1}
                        _ = fi34
                        _ = v35
                    }
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                var v36 int32 = -621
                _ = v36
            }
            switch 9223372036854775807 {
            case -1:
                var v37 int8 = -128
                _ = v37
            }
        }
    }
    v38 := GenS4{F0: 0, F1: -142, F2: true}
    if (51 != 117) {
        v39 := GenS3{F0: 42, F1: 1}
        _ = v39
    } else {
        switch -2147483648 {
        case -1:
            fmt.Println(v1, v1)
        case 502:
            fmt.Println(v1)
        default:
            var v40 int32 = -242
            _ = v40
        }
        if (((0.1 + -2.25) / 0.1) != 1.5) {
            delete(v4, 1)
            v1 += uint(v5)
            v41 := v4[128]
            _ = v41
        }
        v1 = 2
    }
    v1 = (v1 - v1)
    if false {
        v1 = 367
        switch "\n\t" {
        case "foo bar":
            v5 = v5
        case "a":
            fmt.Println(v3)
        }
    }
    fmt.Println(v1, v1, v5)
    if (!true) {
        for fi42 := 0; fi42 < 15; fi42++ {
            _ = len(v2)
            if false {
                var v43 int16 = -920
                fmt.Println(v1, v1)
                var v44 bool = true
                _ = v43
                _ = v44
            }
            _ = fi42
        }
        for fi45 := 0; fi45 < 15; fi45++ {
            if true {
                for fi46 := 0; fi46 < 14; fi46++ {
                    if (-32768 > 1000) {
                        var v47 uint16 = v5
                        v3 -= 42
                        if false {
                            var v48 uint16 = v5
                            switch "foo bar" {
                            case "a":
                                var v49 uint8 = v3
                                _ = v49
                            }
                            switch "hello" {
                            case "\n\t":
                                v1 -= v1
                            default:
                                var v50 uint8 = 255
                                _ = v50
                            }
                            _ = v48
                        } else {
                            var v51 int64 = -1
                            var v52 int = 1
                            v53 := GenS3{F0: int8(v47), F1: 112}
                            v5 += v5
                            v54 := GenS3{F0: 100, F1: 100}
                            var v55 int64 = int64(v1)
                            var v56 int8 = 0
                            _ = v51
                            _ = v52
                            _ = v53
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                        v5 += uint16(v1)
                        _ = v47
                    }
                    v3 = (v3 / (((v3 << (((v1 * 255)) & 63))) | 1))
                    v5 += v5
                    _ = fi46
                }
                for fi57 := 0; fi57 < 18; fi57++ {
                    v1 += 42
                    if false {
                        var v58 uint8 = (v3 >> (((^(v1))) & 63))
                        _ = len(v2)
                        if (!(!true)) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if (((-382 >= fi45) && (!(0 != 1))) && ((v3 & (v58 ^ v58)) != (v58 << (((^(v1))) & 63)))) {
                                if (((v3 / ((v3) | 1)) <= (^(v3))) && (3.14159 < -2.25)) {
                                    switch "foo bar" {
                                    case "a":
                                        fmt.Println(v3, fi45, fi45)
                                        var v59 int16 = 853
                                        delete(v2, "foo bar")
                                        _ = v59
                                    case "\n\t":
                                        v60 := GenS1{F0: v1}
                                        _ = v60
                                    case "foo bar":
                                        if true {
                                            v58 = v3
                                            var v61 uint = v1
                                            _ = v61
                                        } else {
                                            switch "" {
                                            case "":
                                                _ = len(v6)
                                            case "a":
                                                v6 = append(v6, (fi57 - (fi57 & 32767)))
                                            case "\n\t":
                                                fmt.Println(v3, v1)
                                            }
                                            _ = len(v2)
                                            var v62 int8 = -128
                                            if true {
                                                v3 = v3
                                            }
                                            _ = v62
                                        }
                                    default:
                                        fmt.Println(v58, fi57, fi45)
                                    }
                                    if true {
                                        var v63 uint8 = 2
                                        var v64 float64 = (1e10 + -2.25)
                                        _ = v63
                                        _ = v64
                                    } else {
                                        var v65 uint16 = (^(v5))
                                        v66 := GenS5{F0: 1, F1: v3}
                                        var v67 int8 = 100
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    }
                                    fmt.Println(v3, fi45)
                                    if false {
                                        v5 = 60
                                        if (true && (true || false)) {
                                            fmt.Println(v1, fi57, fi45)
                                            v68 := v4[(fi57 >> ((v1) & 63))]
                                            if true {
                                                if false {
                                                    fmt.Println(fi57)
                                                    v2[v68] = true
                                                } else {
                                                    switch fi57 {
                                                    case 128:
                                                        switch fi45 {
                                                        case 9223372036854775807:
                                                            v3 = v3
                                                        }
                                                    case 449:
                                                        _ = len(v6)
                                                    case 32767:
                                                        if (-491 == 42) {
                                                            fmt.Println(v3, v5, v5)
                                                            var v69 float32 = 1e10
                                                            _ = v69
                                                        }
                                                    }
                                                    if (float32((-2.25 * (3.14159 / 1e10))) >= 1.5) {
                                                        delete(v2, v68)
                                                        v68 = v68
                                                    }
                                                    v5 += 0
                                                }
                                                var v70 int8 = 42
                                                fmt.Println(v3, v1)
                                                v71 := GenS6{F0: 0}
                                                if false {
                                                    var v72 int8 = 114
                                                    var v73 uint = (^(v1))
                                                    var v74 uint64 = 18446744073709551615
                                                    _ = v72
                                                    _ = v73
                                                    _ = v74
                                                }
                                                v5 += (^(v5))
                                                _ = v70
                                                _ = v71
                                            }
                                            _ = v68
                                        }
                                        if true {
                                            _ = len(v4)
                                            var v75 bool = false
                                            v58 -= (v3 | 42)
                                            _ = v75
                                        } else {
                                            v6 = append(v6, -9223372036854775808)
                                        }
                                    } else {
                                        if false {
                                            var v76 int32 = 701
                                            var v77 float64 = (1.5 / (((float64(3.14159) + 1e10) / 3.14159) + 0.0))
                                            _ = v76
                                            _ = v77
                                        }
                                    }
                                } else {
                                    v3 = v58
                                    v78 := GenS7{F0: -32768, F1: v58}
                                    delete(v4, fi45)
                                    _ = v78
                                }
                                switch fi57 {
                                case 42:
                                    v3 = (v3 / ((255) | 1))
                                    var v79 string = ((("hello" + "a") + "hello") + ("a" + "foo bar"))
                                    v80 := GenS4{F0: 215, F1: 42, F2: (uint32(v3) != 2147483647)}
                                    v58 += (^(v58))
                                    _ = v79
                                    _ = v80
                                case -637:
                                    v81 := GenS1{F0: (v1 * 42)}
                                    _ = v81
                                default:
                                    v82 := v2["\n\t"]
                                    v83 := GenS8{F0: -419, F1: uint8(v1)}
                                    if v82 {
                                        v5 = v5
                                        v84 := GenS9{F0: fi57, F1: -9223372036854775808, F2: 18}
                                        v5 = (^(v5))
                                        fmt.Println(v58, v5)
                                        v1 = v1
                                        fmt.Println(v1, v3, v82)
                                        _ = v84
                                    } else {
                                        v6 = append(v6, (fi45 >> (((v1 / (((^(v1))) | 1))) & 63)))
                                        if v82 {
                                            switch -2147483648 {
                                            case 332:
                                                var v85 uint8 = v3
                                                _ = v85
                                            default:
                                                v86 := GenS7{F0: 42, F1: (v3 | v58)}
                                                _ = v86
                                            }
                                            var v87 bool = (v82 && false)
                                            v6 = append(v6, 437)
                                            switch 610 {
                                            case 42:
                                                fmt.Println(v87, v87)
                                            case -9223372036854775808:
                                                v88 := GenS5{F0: 0, F1: 0}
                                                _ = v88
                                            }
                                            _ = v87
                                        }
                                    }
                                    _ = v82
                                    _ = v83
                                }
                            }
                            switch "hello" {
                            case "a":
                                var v89 uint64 = 989
                                _ = v89
                            case "hello":
                                switch "a" {
                                case "hello":
                                    v1 = uint(v5)
                                case "foo bar":
                                    v90 := GenS6{F0: 4294967295}
                                    _ = v90
                                case "a":
                                    if true {
                                        var v91 float64 = ((((((0.1 + 1e10) - 0.1) + (0.1 * 1e10)) / -2.25) / (0.1 / 1.5)) / 1.0)
                                        if (42 > 42) {
                                            var v92 uint = (v1 + (v1 - v1))
                                            _ = v92
                                        }
                                        fmt.Println(v58, v91)
                                        _ = v91
                                    }
                                default:
                                    v93 := GenS10{F0: v5}
                                    _ = v93
                                }
                            case "foo bar":
                                var v94 bool = (4294967295 > 1)
                                _ = v94
                            default:
                                var v95 uint64 = 274
                                _ = v95
                            }
                        } else {
                            v5 -= v5
                        }
                        fmt.Println(fi57)
                        _ = v58
                    } else {
                        switch (fi45 | 216) {
                        case 42:
                            v96 := GenS9{F0: -1, F1: 42, F2: -107}
                            _ = v96
                        default:
                            var v97 float32 = 0.1
                            if (v1 >= 18446744073709551615) {
                                v98 := GenS6{F0: 4294967295}
                                var v99 uint16 = v5
                                _ = v98
                                _ = v99
                            } else {
                                fmt.Println(fi45, fi57)
                                var v100 int8 = 42
                                v1 += (v1 + v1)
                                _ = v100
                            }
                            _ = v97
                        }
                        v5 += (v5 + (v5 ^ (^(v5))))
                    }
                    v101 := GenS8{F0: -984, F1: (v3 * (v3 >> ((v1) & 63)))}
                    _ = fi57
                    _ = v101
                }
            }
            v3 = (v3 / (((v3 >> (((v1 + v1)) & 63))) | 1))
            fmt.Println(v1, v3, v3)
            _ = fi45
        }
    }
    switch "hello" {
    case "\n\t":
        for fi102 := 0; fi102 < 16; fi102++ {
            fmt.Println(v3, v1, fi102)
            var v103 int16 = 1000
            fmt.Println(v3, v3)
            _ = fi102
            _ = v103
        }
    case "":
        for fi104 := 0; fi104 < 4; fi104++ {
            v105 := GenS3{F0: -1, F1: -96}
            _ = fi104
            _ = v105
        }
    case "hello":
        var v106 int = 0
        _ = v106
    default:
        v5 = v5
    }
    switch 463 {
    case 1:
        var v107 int = 32768
        _ = v107
    case -950:
        v1 += v1
    default:
        v1 += 4294967295
    }
    v3 -= 157
    v108 := GenS2{F0: 2147483647, F1: 65535}
    var v109 float64 = (1e10 * (3.14159 + (((1e10 + 0.0) - 3.14159) + float64(3.14159))))
    fmt.Println(v3, v5)
    var v110 int16 = 231
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v108
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
