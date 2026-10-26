package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS4 struct { F0 int64; F1 uint; F2 int32 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 uint16; F1 uint16 }
type GenS8 struct { F0 int; F1 uint64; F2 bool }
type GenS9 struct { F0 bool; F1 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint8; F1 int64; F2 int }
type GenS12 struct { F0 int; F1 int8; F2 uint }
type GenS13 struct { F0 string; F1 uint16 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int32; F1 int8 }
type GenS17 struct { F0 int8 }
func genFn0(p0 uint, p1 string) uint64 {
    if (42 <= 1000) {
        v4 := GenS1{F0: 9223372036854775807, F1: 0, F2: 9223372036854775807}
        v5 := GenS2{F0: -215}
        _ = v4
        _ = v5
    } else {
        v6 := []int{-481, 32768}
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := []int{fi8}
            p1 = p1
            _ = fi8
            _ = v9
        }
        v10 := GenS3{F0: p0, F1: 4294967295, F2: 127}
        _ = fi7
        _ = v10
    }
    _ = p0
    _ = p1
    return 532
}

func main() {
    v1 := []uint32{2, 244, 437, 42}
    v2 := make(map[string]uint32)
    var v3 uint16 = 1000
    v3 = (v3 % (((v3 << ((2) & 63))) | 1))
    v11 := GenS3{F0: 18446744073709551615, F1: 576, F2: 255}
    switch ("\n\t" + "a") {
    case "":
        var v12 bool = false
        _ = v12
    case "\n\t":
        v13 := GenS3{F0: 4294967295, F1: genFn0(820, "a"), F2: 128}
        _ = v13
    default:
        switch "" {
        case "\n\t":
            if (2 < 42) {
                switch "hello" {
                case "foo bar":
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v15 := GenS4{F0: 225, F1: 65535, F2: -2147483648}
                        v3 = 2
                        _ = len(v1)
                        _ = fi14
                        _ = v15
                    }
                case "\n\t":
                    _ = len(v1)
                default:
                    _ = len(v1)
                }
                if (42 < 0) {
                    switch "" {
                    case "hello":
                        fmt.Println(v3)
                    case "":
                        _ = len(v1)
                    default:
                        if true {
                            var v16 int16 = 42
                            _ = v16
                        }
                    }
                    var v17 bool = true
                    for fi18 := 0; fi18 < 14; fi18++ {
                        v3 = 32767
                        fmt.Println(fi18)
                        _ = fi18
                    }
                    _ = v17
                } else {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        var v20 uint8 = 127
                        var v21 bool = (782 <= 0)
                        v21 = false
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                var v22 uint32 = 738
                _ = v22
            } else {
                v3 += v3
            }
        case "hello":
            var v23 bool = (127 == -128)
            _ = v23
        case "":
            v1 = append(v1, 2147483648)
        default:
            fmt.Println(v3)
        }
    }
    v24 := GenS1{F0: 32767, F1: 851, F2: -1}
    var v25 uint = 0
    var v26 int64 = -1
    fmt.Println(v26, v25, v26)
    if (32767 > 1000) {
        for fi27 := 0; fi27 < 2; fi27++ {
            var v28 uint = v25
            _ = len(v2)
            fmt.Println(v26, fi27, v28)
            _ = fi27
            _ = v28
        }
    }
    v29 := GenS4{F0: (v26 / ((1) | 1)), F1: 18446744073709551615, F2: -738}
    for fi30 := 0; fi30 < 17; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            fmt.Println(v3, v26)
            v32 := GenS1{F0: fi30, F1: -772, F2: (v26 & 1)}
            v33 := GenS1{F0: (fi30 >> (((v25 | (v25 % ((v25) | 1)))) & 63)), F1: 2147483647, F2: 42}
            _ = fi31
            _ = v32
            _ = v33
        }
        var v34 string = ("foo bar" + ("\n\t" + "hello"))
        _ = fi30
        _ = v34
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        v36 := GenS1{F0: fi35, F1: 42, F2: (v26 ^ (v26 << (((v25 & 675)) & 63)))}
        for fi37 := 0; fi37 < 8; fi37++ {
            if true {
                v3 -= (v3 & 822)
            }
            v3 += v3
            switch 798 {
            case 164:
                fmt.Println(v26)
            case -423:
                switch "" {
                case "\n\t":
                    fmt.Println(v25)
                }
            }
            if len(v1) > 0 {
                v1[0] = 65535
            }
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    switch 918 {
    case 9223372036854775807:
        for fi38 := 0; fi38 < 14; fi38++ {
            if ((!true) && (int8(v26) <= -1)) {
                switch (fi38 / (((fi38 % ((fi38) | 1))) | 1)) {
                case 9223372036854775807:
                    if ("hello" <= "foo bar") {
                        for fi39 := 0; fi39 < 16; fi39++ {
                            var v40 int32 = 42
                            fmt.Println(v25, fi38)
                            _ = fi39
                            _ = v40
                        }
                        if len(v1) > 0 {
                            v1[0] = 2147483647
                        }
                    }
                case 127:
                    if true {
                        if true {
                            _ = len(v1)
                            if (0 < -1) {
                                v3 = v3
                                v25 -= v25
                                switch -137 {
                                case 9223372036854775807:
                                    if false {
                                        var v41 uint8 = 255
                                        if false {
                                            for fi42 := 0; fi42 < 9; fi42++ {
                                                var v43 float32 = float32(1.5)
                                                if true {
                                                    v43 = v43
                                                    var v44 string = "hello"
                                                    var v45 uint64 = 214
                                                    if ((v3 - (v3 * (^(v3)))) != v3) {
                                                        if (!((genFn0(uint(v45), "") >= uint64(v41)) || ((42 < 483) && false))) {
                                                            switch (fi42 - fi38) {
                                                            case 334:
                                                                var v46 uint16 = (v3 - v3)
                                                                _ = v46
                                                            case -335:
                                                                if ((!false) || false) {
                                                                    _ = len(v2)
                                                                    fmt.Println(v26, v44)
                                                                } else {
                                                                    if false {
                                                                        v41 = 127
                                                                        if false {
                                                                            fmt.Println(v26)
                                                                            if false {
                                                                                fmt.Println(v45, fi42, v3)
                                                                            } else {
                                                                                fmt.Println(v43, v25)
                                                                                v47 := GenS5{F0: 42, F1: 1}
                                                                                v43 -= v43
                                                                                _ = v47
                                                                            }
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        } else {
                                                                            v45 = (v45 - v45)
                                                                            var v48 uint8 = v41
                                                                            _ = v48
                                                                        }
                                                                    } else {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 1
                                                                        }
                                                                        _ = len(v1)
                                                                        var v49 uint32 = 4294967295
                                                                        _ = v49
                                                                    }
                                                                    v50 := GenS5{F0: 0, F1: 127}
                                                                    v43 -= v43
                                                                    _ = v50
                                                                }
                                                            case 577:
                                                                if (!(true || false)) {
                                                                    var v51 int64 = -1
                                                                    var v52 bool = (((v25 % ((4294967295) | 1)) >= uint(v45)) || true)
                                                                    if (42 != -2147483648) {
                                                                        switch 4294967295 {
                                                                        case 42:
                                                                            switch 127 {
                                                                            case 128:
                                                                                if v52 {
                                                                                    var v53 string = v44
                                                                                    var v54 int64 = (v51 % ((v26) | 1))
                                                                                    var v55 int8 = -86
                                                                                    v43 += v43
                                                                                    var v56 uint16 = (v3 / ((v3) | 1))
                                                                                    _ = v53
                                                                                    _ = v54
                                                                                    _ = v55
                                                                                    _ = v56
                                                                                } else {
                                                                                    v1 = append(v1, 256)
                                                                                }
                                                                            case 9223372036854775807:
                                                                                _ = len(v2)
                                                                            case 137:
                                                                                var v57 int64 = (^(v51))
                                                                                _ = v57
                                                                            default:
                                                                                switch (^(fi38)) {
                                                                                case 0:
                                                                                    fmt.Println(v52)
                                                                                    v58 := GenS6{F0: 551, F1: -2147483648}
                                                                                    _ = v58
                                                                                case 1:
                                                                                    var v59 int32 = 0
                                                                                    v59 = v59
                                                                                    v51 = (v51 % ((v51) | 1))
                                                                                    _ = v59
                                                                                case 128:
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                default:
                                                                                    var v60 uint64 = (v45 << ((214) & 63))
                                                                                    _ = v60
                                                                                }
                                                                            }
                                                                        case 2147483647:
                                                                            var v61 int16 = 32767
                                                                            _ = v61
                                                                        default:
                                                                            var v62 uint64 = (v45 % ((genFn0((v25 / ((255) | 1)), (v44 + v44))) | 1))
                                                                            _ = v62
                                                                        }
                                                                        v43 = (1e10 - (3.14159 / 1e10))
                                                                    } else {
                                                                        v63 := GenS7{F0: 748, F1: v3}
                                                                        if v52 {
                                                                            v43 -= v43
                                                                        } else {
                                                                            v64 := GenS4{F0: v26, F1: uint(v3), F2: -2147483648}
                                                                            _ = v64
                                                                        }
                                                                        v26 = (v51 << ((77) & 63))
                                                                        var v65 uint8 = 189
                                                                        _ = v63
                                                                        _ = v65
                                                                    }
                                                                    _ = v51
                                                                    _ = v52
                                                                } else {
                                                                    fmt.Println(v43, v44)
                                                                    if true {
                                                                        if ((0.0 != 1.5) || true) {
                                                                            fmt.Println(v45, fi42, v3)
                                                                            v44 = v44
                                                                        } else {
                                                                            v41 = (v41 - (v41 << ((v25) & 63)))
                                                                        }
                                                                    } else {
                                                                        v45 += (v45 * (v45 << ((65535) & 63)))
                                                                        if (!(!true)) {
                                                                            if false {
                                                                                v25 -= 65535
                                                                                fmt.Println(fi42, v41)
                                                                                fmt.Println(v26, v44)
                                                                                switch v44 {
                                                                                case "a":
                                                                                    v1 = append(v1, 677)
                                                                                case "":
                                                                                    switch fi38 {
                                                                                    case 159:
                                                                                        var v66 uint64 = genFn0(v25, (v44 + "hello"))
                                                                                        _ = v66
                                                                                    }
                                                                                case "hello":
                                                                                    v67 := v2[(v44 + v44)]
                                                                                    _ = v67
                                                                                }
                                                                                var v68 float64 = (-2.25 / float64(v43))
                                                                                _ = v68
                                                                            } else {
                                                                                v26 = v26
                                                                                var v69 int = fi42
                                                                                switch fi38 {
                                                                                case -1:
                                                                                    v41 -= v41
                                                                                case 128:
                                                                                    switch (fi38 / (((-(v69))) | 1)) {
                                                                                    case -821:
                                                                                        switch (v44 + v44) {
                                                                                        case "a":
                                                                                            fmt.Println(v45)
                                                                                        default:
                                                                                            var v70 int16 = 1
                                                                                            _ = v70
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    var v71 int64 = v26
                                                                                    v71 = -9223372036854775808
                                                                                    _ = v71
                                                                                } else {
                                                                                    var v72 float64 = 0.0
                                                                                    _ = v72
                                                                                }
                                                                                if (!false) {
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        var v73 int = fi42
                                                                                        _ = v73
                                                                                    default:
                                                                                        var v74 int64 = (v26 & -9223372036854775808)
                                                                                        var v75 bool = (0 < 1)
                                                                                        _ = v74
                                                                                        _ = v75
                                                                                    }
                                                                                    var v76 float64 = 0.1
                                                                                    v2["hello"] = 4294967295
                                                                                    switch v69 {
                                                                                    case 2147483647:
                                                                                        v43 = 0.0
                                                                                    case 787:
                                                                                        v45 = (v45 | (^(v45)))
                                                                                    case -817:
                                                                                        var v77 int16 = 1
                                                                                        _ = v77
                                                                                    default:
                                                                                        if (false && false) {
                                                                                            switch (fi38 % ((1) | 1)) {
                                                                                            case 9223372036854775807:
                                                                                                v78 := GenS3{F0: 18446744073709551615, F1: (v45 - v45), F2: (v41 | (v41 << ((v25) & 63)))}
                                                                                                if true {
                                                                                                    v25 += 729
                                                                                                    var v79 string = (v44 + v44)
                                                                                                    v76 = (float64(v43) + -2.25)
                                                                                                    _ = v79
                                                                                                } else {
                                                                                                    v43 = float32((3.14159 + v76))
                                                                                                    delete(v2, ("" + ""))
                                                                                                    v76 = 0.0
                                                                                                }
                                                                                                if ((!(genFn0(v25, v44) < (v45 / ((v45) | 1)))) || false) {
                                                                                                    var v80 int8 = -1
                                                                                                    v1 = append(v1, 65535)
                                                                                                    _ = v80
                                                                                                }
                                                                                                _ = v78
                                                                                            case 2147483647:
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            case -379:
                                                                                                var v81 string = (("foo bar" + "foo bar") + "")
                                                                                                _ = v81
                                                                                            default:
                                                                                                if (-1 < -1) {
                                                                                                    if (!true) {
                                                                                                        _ = len(v1)
                                                                                                        v26 = (v26 << (((v25 ^ v25)) & 63))
                                                                                                        var v82 int64 = (-(v26))
                                                                                                        _ = v82
                                                                                                    }
                                                                                                    var v83 int8 = 82
                                                                                                    _ = v83
                                                                                                } else {
                                                                                                    v44 = v44
                                                                                                    v84 := GenS1{F0: (^(fi38)), F1: 42, F2: v26}
                                                                                                    fmt.Println(fi42, v26)
                                                                                                    if false {
                                                                                                        v85 := GenS2{F0: 32767}
                                                                                                        v43 -= v43
                                                                                                        _ = v85
                                                                                                    }
                                                                                                    var v86 float64 = v76
                                                                                                    fmt.Println(fi42)
                                                                                                    _ = v84
                                                                                                    _ = v86
                                                                                                }
                                                                                            }
                                                                                            var v87 float64 = (3.14159 - v76)
                                                                                            _ = v87
                                                                                        }
                                                                                    }
                                                                                    _ = v76
                                                                                } else {
                                                                                    switch "" {
                                                                                    case "hello":
                                                                                        if ((v3 * v3) <= v3) {
                                                                                            v3 = v3
                                                                                        }
                                                                                    case "a":
                                                                                        v88 := GenS1{F0: int(v25), F1: 0, F2: -1}
                                                                                        _ = v88
                                                                                    case "\n\t":
                                                                                        var v89 uint = v25
                                                                                        _ = v89
                                                                                    }
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 2
                                                                                    }
                                                                                    var v90 int16 = 1000
                                                                                    _ = v90
                                                                                }
                                                                                _ = v69
                                                                            }
                                                                            var v91 bool = (1000 != -32768)
                                                                            _ = v91
                                                                        } else {
                                                                            switch 2147483647 {
                                                                            case 128:
                                                                                switch 1 {
                                                                                case -595:
                                                                                    fmt.Println(fi38, v43)
                                                                                default:
                                                                                    v92 := GenS2{F0: 0}
                                                                                    _ = v92
                                                                                }
                                                                                var v93 int = (fi42 >> ((v25) & 63))
                                                                                _ = v93
                                                                            case 127:
                                                                                fmt.Println(v43)
                                                                            case 0:
                                                                                v94 := GenS6{F0: v25, F1: 585}
                                                                                _ = v94
                                                                            }
                                                                            var v95 uint16 = (v3 * (v3 + v3))
                                                                            if true {
                                                                                var v96 uint32 = 407
                                                                                v44 = "\n\t"
                                                                                if (!(-333 <= -2147483648)) {
                                                                                    switch (fi38 & fi38) {
                                                                                    case 2147483647:
                                                                                        switch fi38 {
                                                                                        case 1:
                                                                                            switch v44 {
                                                                                            case "a":
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = 1
                                                                                                }
                                                                                            default:
                                                                                                v97 := GenS8{F0: (fi42 & (fi38 & (^(fi38)))), F1: (v45 >> (((v25 >> (((v25 % (((v25 * (v25 * 1))) | 1))) & 63))) & 63)), F2: (v3 <= (v3 / (((^(v3))) | 1)))}
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = v96
                                                                                                }
                                                                                                _ = len(v1)
                                                                                                _ = v97
                                                                                            }
                                                                                        case 128:
                                                                                            if ((false && (-178 == -717)) || true) {
                                                                                                var v98 uint64 = 0
                                                                                                var v99 string = v44
                                                                                                var v100 uint = (v25 % ((1) | 1))
                                                                                                v101 := GenS2{F0: 1000}
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                                _ = v98
                                                                                                _ = v99
                                                                                            }
                                                                                        default:
                                                                                            switch fi42 {
                                                                                            case 42:
                                                                                                var v102 uint16 = v95
                                                                                                _ = v102
                                                                                            case 32768:
                                                                                                v44 = v44
                                                                                            case -2147483648:
                                                                                                switch (fi38 / ((128) | 1)) {
                                                                                                case 128:
                                                                                                    v44 = ""
                                                                                                    var v103 uint8 = (v41 + (v41 & (v41 * (v41 + v41))))
                                                                                                    v104 := GenS2{F0: -697}
                                                                                                    _ = v103
                                                                                                    _ = v104
                                                                                                }
                                                                                            default:
                                                                                                v105 := GenS9{F0: false, F1: (true || (1 > (v26 >> (((v25 | v25)) & 63))))}
                                                                                                _ = v105
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    delete(v2, v44)
                                                                                } else {
                                                                                    var v106 uint8 = v41
                                                                                    if false {
                                                                                        v41 = 255
                                                                                        v3 += (^(v3))
                                                                                    }
                                                                                    fmt.Println(v26, v26, v45)
                                                                                    _ = v106
                                                                                }
                                                                                _ = v96
                                                                            } else {
                                                                                var v107 int = (fi38 - int(v25))
                                                                                fmt.Println(v3, v44)
                                                                                var v108 uint32 = 2147483647
                                                                                _ = v107
                                                                                _ = v108
                                                                            }
                                                                            _ = v95
                                                                        }
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            v109 := GenS10{F0: v45}
                                                                            _ = v109
                                                                        case "":
                                                                            switch 127 {
                                                                            case 32768:
                                                                                var v110 int8 = 1
                                                                                _ = v110
                                                                            case 2147483647:
                                                                                if (true || ((v25 - v25) > (v25 << ((v25) & 63)))) {
                                                                                    fmt.Println(v25, v43, fi38)
                                                                                    fmt.Println(v26, v3, v26)
                                                                                } else {
                                                                                    var v111 int32 = 42
                                                                                    var v112 uint64 = (v45 % ((v45) | 1))
                                                                                    if (-128 == 1) {
                                                                                        if ((v111 > (v111 - v111)) || true) {
                                                                                            var v113 uint32 = 1
                                                                                            fmt.Println(v41)
                                                                                            _ = v113
                                                                                        } else {
                                                                                            v25 += 1
                                                                                        }
                                                                                        v114 := GenS2{F0: -1}
                                                                                        _ = v114
                                                                                    }
                                                                                    if true {
                                                                                        switch fi42 {
                                                                                        case 626:
                                                                                            var v115 float32 = v43
                                                                                            _ = v115
                                                                                        case -2147483648:
                                                                                            var v116 int = (fi38 & (fi42 - fi38))
                                                                                            _ = v116
                                                                                        }
                                                                                    } else {
                                                                                        if (v25 <= v25) {
                                                                                            v45 = 1
                                                                                            v45 = (^(v45))
                                                                                        } else {
                                                                                            v1 = append(v1, 242)
                                                                                            if ((v3 << (((v25 >> (((v25 & v25)) & 63))) & 63)) <= (v3 & (v3 / (((^(v3))) | 1)))) {
                                                                                                var v117 int = -1
                                                                                                _ = v117
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v118 := GenS11{F0: v41, F1: 9223372036854775807, F2: 128}
                                                                                    if false {
                                                                                        fmt.Println(fi38)
                                                                                    }
                                                                                    _ = v111
                                                                                    _ = v112
                                                                                    _ = v118
                                                                                }
                                                                                var v119 uint64 = (v45 * genFn0(v25, "foo bar"))
                                                                                v43 = (v43 / v43)
                                                                                _ = v119
                                                                            default:
                                                                                v26 -= (v26 - v26)
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                var v120 uint64 = genFn0(v25, v44)
                                                                _ = v120
                                                            }
                                                            switch v44 {
                                                            case "\n\t":
                                                                v121 := GenS12{F0: 855, F1: -92, F2: v25}
                                                                _ = v121
                                                            case "":
                                                                fmt.Println(v44, v25, v25)
                                                            case "a":
                                                                fmt.Println(v25, v26, v26)
                                                            }
                                                        }
                                                    }
                                                    var v122 float32 = 1e10
                                                    var v123 uint = v25
                                                    if false {
                                                        var v124 float64 = -2.25
                                                        fmt.Println(v44)
                                                        _ = v124
                                                    } else {
                                                        if (369 < 910) {
                                                            var v125 string = (("a" + "\n\t") + "\n\t")
                                                            _ = v125
                                                        } else {
                                                            switch v44 {
                                                            case "":
                                                                if true {
                                                                    v43 = 1e10
                                                                }
                                                            case "a":
                                                                var v126 uint = (v123 % (((v25 >> (((v123 / (((v123 * (v25 + (v123 / ((uint(v45)) | 1))))) | 1))) & 63))) | 1))
                                                                _ = v126
                                                            default:
                                                                v127 := GenS13{F0: v44, F1: v3}
                                                                v128 := GenS12{F0: fi42, F1: -22, F2: v123}
                                                                _ = v127
                                                                _ = v128
                                                            }
                                                            var v129 uint32 = 0
                                                            fmt.Println(fi42, v3)
                                                            _ = v129
                                                        }
                                                        fmt.Println(v43, v41)
                                                        var v130 uint8 = 167
                                                        _ = v130
                                                    }
                                                    if (!true) {
                                                        v45 += (v45 ^ 22)
                                                        v41 = 183
                                                        switch (fi38 % ((fi38) | 1)) {
                                                        case 320:
                                                            v45 -= v45
                                                        case -962:
                                                            var v131 uint = (v123 & 2)
                                                            _ = v131
                                                        case 128:
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        default:
                                                            v26 += (-(v26))
                                                        }
                                                    }
                                                    _ = v122
                                                    _ = v123
                                                    _ = v44
                                                    _ = v45
                                                }
                                                v1 = append(v1, 42)
                                                v132 := GenS13{F0: ("\n\t" + ("hello" + "")), F1: (v3 | (v3 >> (((v25 << ((860) & 63))) & 63)))}
                                                var v133 int32 = 2147483647
                                                _ = fi42
                                                _ = v132
                                                _ = v133
                                                _ = v43
                                            }
                                            var v134 float32 = (0.0 * (0.0 - 3.14159))
                                            v26 = -404
                                            _ = v134
                                        }
                                        _ = v41
                                    }
                                default:
                                    var v135 string = "\n\t"
                                    _ = v135
                                }
                            } else {
                                v1 = append(v1, 897)
                            }
                        }
                        var v136 float64 = ((1.5 - -2.25) - 3.14159)
                        _ = v136
                    } else {
                        v137 := GenS14{F0: 32767}
                        _ = v137
                    }
                case 2147483647:
                    for fi138 := 0; fi138 < 18; fi138++ {
                        v26 -= 1
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(fi138, v26, v26)
                        if ((!true) && (false && (!false))) {
                            v139 := GenS15{F0: 245}
                            var v140 uint64 = 931
                            v141 := GenS3{F0: uint(v140), F1: (v140 - (v140 | 2)), F2: 2}
                            _ = v139
                            _ = v140
                            _ = v141
                        }
                        fmt.Println(v26)
                        v142 := GenS16{F0: 757, F1: 42}
                        delete(v2, "hello")
                        var v143 uint8 = 42
                        if false {
                            v143 = (v143 >> ((v25) & 63))
                        }
                        _ = fi138
                        _ = v142
                        _ = v143
                    }
                default:
                    switch 42 {
                    case -2147483648:
                        for fi144 := 0; fi144 < 4; fi144++ {
                            if len(v1) > 0 {
                                v1[0] = 939
                            }
                            _ = fi144
                        }
                    default:
                        var v145 uint = (v25 % (((v25 * (v25 | (v25 - (v25 & (v25 >> ((230) & 63))))))) | 1))
                        _ = v145
                    }
                }
                _ = len(v1)
                fmt.Println(v26, fi38)
            }
            switch fi38 {
            case 127:
                v1 = append(v1, 4294967295)
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi38
        }
    case -1:
        var v146 int64 = (v26 + v26)
        _ = v146
    case 710:
        var v147 int16 = 1000
        _ = v147
    default:
        if false {
            switch -656 {
            case 0:
                for fi148 := 0; fi148 < 2; fi148++ {
                    v26 -= (v26 * v26)
                    fmt.Println(v3)
                    _ = fi148
                }
            case 809:
                var v149 uint8 = 112
                _ = v149
            case 9223372036854775807:
                fmt.Println(v26, v26, v25)
            default:
                if false {
                    v150 := GenS5{F0: 2147483647, F1: -125}
                    v151 := GenS17{F0: 113}
                    _ = v150
                    _ = v151
                }
            }
        }
    }
    _ = len(v1)
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
}
