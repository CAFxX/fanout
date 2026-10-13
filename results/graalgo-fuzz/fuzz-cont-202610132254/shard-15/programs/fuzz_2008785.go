package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint32; F2 uint }
type GenS6 struct { F0 int8; F1 bool; F2 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 int8; F2 int }
type GenS9 struct { F0 bool; F1 bool; F2 int8 }
type GenS10 struct { F0 int8; F1 int8 }
type GenS11 struct { F0 uint64; F1 int64 }
type GenS12 struct { F0 uint32; F1 uint }
type GenS13 struct { F0 int; F1 int32 }

func main() {
    var v1 uint16 = 0
    var v2 int64 = 1
    var v3 uint = 42
    var v4 uint = (v3 + 4294967295)
    v5 := make(map[uint32]string)
    var v6 int8 = 39
    v7 := GenS1{F0: (!(true && false)), F1: 0}
    switch 32768 {
    case -2147483648:
        if true {
            fmt.Println(v3)
            if ((v2 ^ v2) > (v2 * 759)) {
                fmt.Println(v4)
            }
            for fi8 := 0; fi8 < 12; fi8++ {
                switch fi8 {
                case -186:
                    v6 = (v6 ^ v6)
                }
                _ = fi8
            }
        } else {
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 uint64 = 18446744073709551615
                fmt.Println(fi9)
                _ = fi9
                _ = v10
            }
            if (true || true) {
                var v11 uint8 = 204
                _ = v11
            }
        }
    default:
        v2 += v2
    }
    fmt.Println(v2, v4, v3)
    v12 := GenS1{F0: true, F1: 465}
    v2 = (v2 * v2)
    v3 += v4
    v4 -= uint(v1)
    _ = len(v5)
    v4 = v3
    switch 127 {
    case 317:
        v13 := []int{127, 32768, 32767}
        _ = v13
    case -2147483648:
        if (("foo bar" > "a") || (true || (true && (false && (3.14159 != (1.5 / ((1e10 - 1.5) / 1.0))))))) {
            if (false && (!false)) {
                v3 -= v3
                for fi14 := 0; fi14 < 5; fi14++ {
                    switch "hello" {
                    case "foo bar":
                        var v15 uint8 = 149
                        _ = v15
                    case "a":
                        switch "foo bar" {
                        case "\n\t":
                            var v16 int = (fi14 * fi14)
                            _ = v16
                        case "":
                            var v17 uint64 = 0
                            _ = v17
                        default:
                            switch ("foo bar" + "a") {
                            case "foo bar":
                                v1 = 0
                            default:
                                if false {
                                    switch fi14 {
                                    case 127:
                                        for fi18 := 0; fi18 < 19; fi18++ {
                                            v19 := GenS1{F0: ("\n\t" == "hello"), F1: 945}
                                            _ = fi18
                                            _ = v19
                                        }
                                    case -530:
                                        var v20 bool = true
                                        _ = v20
                                    case -1:
                                        fmt.Println(v3, v2)
                                    default:
                                        if (false || true) {
                                            fmt.Println(v1, v3)
                                            v4 = (^(v3))
                                            if (1e10 <= (1e10 - 3.14159)) {
                                                if (!false) {
                                                    if false {
                                                        for fi21 := 0; fi21 < 4; fi21++ {
                                                            if true {
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    var v22 int64 = -63
                                                                    _ = v22
                                                                default:
                                                                    v23 := GenS2{F0: 324}
                                                                    _ = v23
                                                                }
                                                                var v24 uint = (v3 ^ v3)
                                                                if false {
                                                                    switch (fi14 >> ((18446744073709551615) & 63)) {
                                                                    case 2147483647:
                                                                        v25 := GenS1{F0: (2 <= 2), F1: 2147483647}
                                                                        _ = v25
                                                                    case 32768:
                                                                        v26 := []int{(-(fi14)), fi21, (fi21 / (((fi14 / ((1) | 1))) | 1))}
                                                                        switch fi21 {
                                                                        case -336:
                                                                            v1 -= (v1 + uint16(v24))
                                                                        default:
                                                                            if (-12 != v6) {
                                                                                var v27 int8 = v6
                                                                                _ = v27
                                                                            } else {
                                                                                v6 += v6
                                                                                if true {
                                                                                    _ = len(v26)
                                                                                    var v28 uint = 255
                                                                                    v29 := GenS3{F0: (v6 & v6)}
                                                                                    switch ("hello" + "foo bar") {
                                                                                    case "foo bar":
                                                                                        if (true || true) {
                                                                                            v30 := GenS3{F0: v6}
                                                                                            v31 := GenS4{F0: (true || (!false))}
                                                                                            _ = v30
                                                                                            _ = v31
                                                                                        } else {
                                                                                            switch 1 {
                                                                                            case -1:
                                                                                                switch "foo bar" {
                                                                                                case "foo bar":
                                                                                                    v4 = 65535
                                                                                                case "hello":
                                                                                                    var v32 string = "a"
                                                                                                    _ = v32
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v6, v2)
                                                                                                }
                                                                                            case -855:
                                                                                                _ = len(v5)
                                                                                            case 127:
                                                                                                v28 -= (v4 + (v28 / ((2) | 1)))
                                                                                            }
                                                                                            v33 := GenS5{F0: 65, F1: 1, F2: (v24 >> (((v3 + 475)) & 63))}
                                                                                            v5[1] = "a"
                                                                                            if true {
                                                                                                delete(v5, 65535)
                                                                                            } else {
                                                                                                if true {
                                                                                                    switch fi14 {
                                                                                                    case 188:
                                                                                                        if true {
                                                                                                            if len(v26) >= 2 {
                                                                                                                v26 = v26[1:len(v26)]
                                                                                                            }
                                                                                                            var v34 int32 = 42
                                                                                                            v35 := GenS6{F0: v6, F1: (uint(v1) >= 1), F2: 92}
                                                                                                            var v36 uint64 = uint64(v3)
                                                                                                            _ = v34
                                                                                                            _ = v35
                                                                                                            _ = v36
                                                                                                        }
                                                                                                    case -2147483648:
                                                                                                        v1 -= uint16(v4)
                                                                                                    case 127:
                                                                                                        if (1e10 >= float64((0.1 / 3.14159))) {
                                                                                                            v3 = v28
                                                                                                            v26 = append(v26, (fi14 / ((730) | 1)))
                                                                                                        }
                                                                                                    default:
                                                                                                        v3 += v3
                                                                                                    }
                                                                                                } else {
                                                                                                    var v37 uint8 = 128
                                                                                                    v1 = v1
                                                                                                    switch "hello" {
                                                                                                    case "\n\t":
                                                                                                        if false {
                                                                                                            var v38 int16 = -1
                                                                                                            v39 := GenS7{F0: (^(v37))}
                                                                                                            v1 = v1
                                                                                                            _ = v38
                                                                                                            _ = v39
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        if len(v26) >= 2 {
                                                                                                            v26 = v26[1:len(v26)]
                                                                                                        }
                                                                                                    }
                                                                                                    if false {
                                                                                                        fmt.Println(v24)
                                                                                                    } else {
                                                                                                        v40 := GenS1{F0: (v2 >= (v2 + v2)), F1: -1}
                                                                                                        _ = v40
                                                                                                    }
                                                                                                    switch fi14 {
                                                                                                    case 127:
                                                                                                        if ((v37 >> (((v3 ^ 82)) & 63)) != v37) {
                                                                                                            fmt.Println(v24, fi21, fi21)
                                                                                                            delete(v5, 225)
                                                                                                            v3 = v24
                                                                                                            if false {
                                                                                                                switch "foo bar" {
                                                                                                                case "\n\t":
                                                                                                                    fmt.Println(v1)
                                                                                                                case "a":
                                                                                                                    if len(v26) >= 2 {
                                                                                                                        v26 = v26[1:len(v26)]
                                                                                                                    }
                                                                                                                case "":
                                                                                                                    var v41 int64 = (v2 % (((v2 << ((42) & 63))) | 1))
                                                                                                                    _ = v41
                                                                                                                default:
                                                                                                                    var v42 int64 = v2
                                                                                                                    _ = v42
                                                                                                                }
                                                                                                                v26 = append(v26, fi21)
                                                                                                                fmt.Println(v2, v28)
                                                                                                            }
                                                                                                        }
                                                                                                    case -647:
                                                                                                        if true {
                                                                                                            switch fi14 {
                                                                                                            case 2147483647:
                                                                                                                v37 += (^(v37))
                                                                                                            case 747:
                                                                                                                v28 = 65535
                                                                                                            default:
                                                                                                                fmt.Println(v6, v37)
                                                                                                            }
                                                                                                            v26 = append(v26, (fi21 ^ 388))
                                                                                                            if true {
                                                                                                                v6 += 109
                                                                                                            } else {
                                                                                                                switch "" {
                                                                                                                case "foo bar":
                                                                                                                    v43 := GenS6{F0: -1, F1: (false && true), F2: v37}
                                                                                                                    _ = v43
                                                                                                                case "\n\t":
                                                                                                                    switch ("a" + "foo bar") {
                                                                                                                    case "hello":
                                                                                                                        v44 := GenS3{F0: v6}
                                                                                                                        _ = v44
                                                                                                                    case "foo bar":
                                                                                                                        v45 := GenS6{F0: (v6 * (v6 % (((v6 - v6)) | 1))), F1: (!true), F2: v37}
                                                                                                                        _ = v45
                                                                                                                    case "\n\t":
                                                                                                                        if (-84 == int16(fi21)) {
                                                                                                                            if false {
                                                                                                                                switch fi14 {
                                                                                                                                case 127:
                                                                                                                                    _ = len(v5)
                                                                                                                                default:
                                                                                                                                    v46 := GenS8{F0: (("\n\t" + ("a" + "a")) + "a"), F1: 1, F2: (fi14 - fi14)}
                                                                                                                                    _ = v46
                                                                                                                                }
                                                                                                                                _ = len(v26)
                                                                                                                            } else {
                                                                                                                                v37 -= uint8(v3)
                                                                                                                                v47 := GenS6{F0: (v6 ^ (v6 % ((v6) | 1))), F1: true, F2: v37}
                                                                                                                                _ = v47
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            var v48 int = 185
                                                                                                                            fmt.Println(v2, v6, v24)
                                                                                                                            _ = len(v26)
                                                                                                                            v49 := GenS8{F0: "a", F1: (-(v6)), F2: -424}
                                                                                                                            _ = v48
                                                                                                                            _ = v49
                                                                                                                        }
                                                                                                                    }
                                                                                                                default:
                                                                                                                    var v50 int = 576
                                                                                                                    _ = v50
                                                                                                                }
                                                                                                                v51 := GenS9{F0: ((^(v4)) < v24), F1: false, F2: (v6 + v6)}
                                                                                                                v24 -= 835
                                                                                                                _ = v51
                                                                                                            }
                                                                                                            v5[65535] = (("" + "a") + "hello")
                                                                                                            var v52 float64 = -2.25
                                                                                                            _ = v52
                                                                                                        } else {
                                                                                                            var v53 uint = v4
                                                                                                            if true {
                                                                                                                fmt.Println(fi21)
                                                                                                            } else {
                                                                                                                switch (fi14 / (((fi21 << ((v4) & 63))) | 1)) {
                                                                                                                case -1:
                                                                                                                    switch int(v1) {
                                                                                                                    case -686:
                                                                                                                        v53 = (v28 ^ v4)
                                                                                                                    default:
                                                                                                                        var v54 uint32 = 65535
                                                                                                                        _ = v54
                                                                                                                    }
                                                                                                                case 4294967295:
                                                                                                                    v55 := GenS3{F0: 0}
                                                                                                                    _ = v55
                                                                                                                }
                                                                                                                if true {
                                                                                                                    var v56 int8 = (v6 | int8(v3))
                                                                                                                    v57 := GenS1{F0: (1000 == 0), F1: 42}
                                                                                                                    _ = v56
                                                                                                                    _ = v57
                                                                                                                } else {
                                                                                                                    v1 = (v1 ^ (v1 * 2))
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v53
                                                                                                        }
                                                                                                    case -591:
                                                                                                        switch "foo bar" {
                                                                                                        case "foo bar":
                                                                                                            var v58 int8 = v6
                                                                                                            _ = v58
                                                                                                        case "\n\t":
                                                                                                            v5[4294967295] = ((("" + "\n\t") + (("\n\t" + "\n\t") + (("a" + (("\n\t" + "\n\t") + "foo bar")) + ""))) + ("\n\t" + ""))
                                                                                                        default:
                                                                                                            v28 = v4
                                                                                                        }
                                                                                                    default:
                                                                                                        v1 -= v1
                                                                                                    }
                                                                                                    _ = v37
                                                                                                }
                                                                                                var v59 uint8 = 0
                                                                                                switch "" {
                                                                                                case "a":
                                                                                                    v3 += (v4 & v4)
                                                                                                case "":
                                                                                                    var v60 int8 = -42
                                                                                                    _ = v60
                                                                                                }
                                                                                                _ = v59
                                                                                            }
                                                                                            switch fi21 {
                                                                                            case -56:
                                                                                                if len(v26) > 0 {
                                                                                                    v26[0] = (fi21 | (fi14 % (((^(fi21))) | 1)))
                                                                                                }
                                                                                            case -1:
                                                                                                var v61 uint8 = 2
                                                                                                _ = v61
                                                                                            case -2147483648:
                                                                                                var v62 uint32 = 1
                                                                                                _ = v62
                                                                                            default:
                                                                                                if (false || true) {
                                                                                                    v24 += v3
                                                                                                }
                                                                                            }
                                                                                            _ = v33
                                                                                        }
                                                                                        v63 := GenS10{F0: (v6 & v6), F1: v6}
                                                                                        _ = v63
                                                                                    }
                                                                                    var v64 uint = 269
                                                                                    _ = v28
                                                                                    _ = v29
                                                                                    _ = v64
                                                                                }
                                                                                v6 += 122
                                                                            }
                                                                        }
                                                                        v65 := GenS9{F0: (!(238 <= 128)), F1: false, F2: v6}
                                                                        _ = v26
                                                                        _ = v65
                                                                    }
                                                                    var v66 int8 = v6
                                                                    var v67 int8 = 100
                                                                    var v68 int32 = 938
                                                                    _ = v66
                                                                    _ = v67
                                                                    _ = v68
                                                                }
                                                                _ = v24
                                                            }
                                                            v69 := GenS8{F0: ("a" + "foo bar"), F1: v6, F2: fi14}
                                                            _ = fi21
                                                            _ = v69
                                                        }
                                                        for fi70 := 0; fi70 < 19; fi70++ {
                                                            fmt.Println(fi70, v6, v1)
                                                            _ = fi70
                                                        }
                                                    } else {
                                                        var v71 bool = false
                                                        v72 := []int{fi14}
                                                        var v73 bool = (18446744073709551615 != 537)
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    v1 = 32767
                                                } else {
                                                    if true {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            for fi74 := 0; fi74 < 19; fi74++ {
                                                                var v75 int32 = -2147483648
                                                                var v76 string = ("\n\t" + "a")
                                                                var v77 float32 = -2.25
                                                                _ = fi74
                                                                _ = v75
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                        case "a":
                                                            for fi78 := 0; fi78 < 4; fi78++ {
                                                                var v79 float32 = 3.14159
                                                                var v80 uint16 = (v1 % (((v1 * v1)) | 1))
                                                                v81 := []int{fi14, (fi14 ^ (fi78 << ((18446744073709551615) & 63)))}
                                                                _ = fi78
                                                                _ = v79
                                                                _ = v80
                                                                _ = v81
                                                            }
                                                        case "":
                                                            var v82 string = ("hello" + "\n\t")
                                                            _ = v82
                                                        default:
                                                            v83 := []int{fi14, 0, fi14, fi14}
                                                            _ = v83
                                                        }
                                                    }
                                                    var v84 float32 = (3.14159 * 0.1)
                                                    v85 := v5[1]
                                                    _ = v84
                                                    _ = v85
                                                }
                                            } else {
                                                v2 = (v2 + (v2 - v2))
                                                if (0.1 > 1.5) {
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v86 uint16 = v1
                                                        _ = v86
                                                    case "hello":
                                                        if (65535 > 680) {
                                                            if true {
                                                                v2 = v2
                                                                var v87 bool = ((v1 > 62) && (1e10 != (1e10 * 0.0)))
                                                                _ = v87
                                                            }
                                                            if true {
                                                                v88 := GenS11{F0: 92, F1: (v2 - (v2 ^ (v2 / (((v2 - v2)) | 1))))}
                                                                _ = v88
                                                            }
                                                        }
                                                    }
                                                }
                                                v89 := []int{32768, fi14}
                                                _ = v89
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    case "":
                        for fi90 := 0; fi90 < 12; fi90++ {
                            v3 = (v3 + v4)
                            _ = fi90
                        }
                    default:
                        v91 := GenS12{F0: 255, F1: (v4 + v4)}
                        _ = v91
                    }
                    _ = fi14
                }
            }
            switch 0 {
            case 0:
                v6 += v6
            }
            var v92 int16 = -1
            _ = v92
        } else {
            v93 := v5[2]
            _ = v93
        }
    case 32768:
        v94 := []int{}
        _ = v94
    default:
        v2 += -1
    }
    for fi95 := 0; fi95 < 6; fi95++ {
        var v96 bool = true
        v97 := v5[2]
        if (false || v96) {
            if (!(true && v96)) {
                for fi98 := 0; fi98 < 6; fi98++ {
                    var v99 uint8 = 1
                    var v100 uint = 856
                    v101 := []int{-543, fi98, 331, -2147483648}
                    _ = fi98
                    _ = v100
                    _ = v101
                    _ = v99
                }
                v102 := GenS10{F0: v6, F1: (v6 >> (((v3 >> ((v3) & 63))) & 63))}
                for fi103 := 0; fi103 < 6; fi103++ {
                    v104 := GenS10{F0: (v6 / ((v6) | 1)), F1: v6}
                    _ = fi103
                    _ = v104
                }
                _ = v102
            } else {
                var v105 uint64 = 42
                _ = v105
            }
        }
        _ = fi95
        _ = v96
        _ = v97
    }
    switch -2147483648 {
    case 2147483647:
        fmt.Println(v2)
    }
    for fi106 := 0; fi106 < 12; fi106++ {
        switch ("\n\t" + ("a" + ("\n\t" + "foo bar"))) {
        case "":
            v4 = v4
        case "foo bar":
            v6 = (v6 >> ((4294967295) & 63))
        }
        _ = fi106
    }
    fmt.Println(v3, v4, v3)
    if (78 < 4294967295) {
        var v107 uint32 = 996
        _ = v107
    } else {
        if true {
            v5[42] = "a"
            v1 = v1
        } else {
            v4 = 515
        }
        if (127 >= 0) {
            fmt.Println(v4, v6)
        } else {
            for fi108 := 0; fi108 < 1; fi108++ {
                for fi109 := 0; fi109 < 6; fi109++ {
                    v110 := GenS13{F0: -2147483648, F1: 83}
                    v3 = (v3 >> (((^(v3))) & 63))
                    _ = fi109
                    _ = v110
                }
                _ = fi108
            }
        }
        var v111 uint32 = 2147483648
        _ = v111
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
