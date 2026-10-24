package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int; F1 string; F2 uint8 }
type GenS11 struct { F0 int; F1 int; F2 int64 }

func main() {
    var v1 int8 = 127
    v2 := []uint32{}
    var v3 uint64 = 0
    for fi4 := 0; fi4 < 9; fi4++ {
        if true {
            if true {
                var v5 int16 = -917
                _ = v5
            }
            if (false && ((1 >= -1) || ((!false) || (972 != 320)))) {
                if (0 == 42) {
                    for fi6 := 0; fi6 < 15; fi6++ {
                        v7 := GenS1{F0: fi4}
                        _ = fi6
                        _ = v7
                    }
                }
                var v8 bool = (1 == 65535)
                _ = v8
            } else {
                switch int(v3) {
                case 0:
                    var v9 uint8 = 243
                    _ = v9
                case 1:
                    v10 := GenS2{F0: 1, F1: fi4}
                    _ = v10
                default:
                    if len(v2) > 0 {
                        v2[0] = 2147483647
                    }
                }
                var v11 float32 = 0.1
                v11 += v11
                _ = v11
            }
            v3 += (^(v3))
        } else {
            _ = len(v2)
            fmt.Println(v1, fi4)
        }
        switch ("" + "a") {
        case "":
            var v12 int64 = 1
            _ = v12
        case "hello":
            if true {
                for fi13 := 0; fi13 < 10; fi13++ {
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                    v1 = v1
                    var v14 uint16 = 65535
                    _ = fi13
                    _ = v14
                }
            } else {
                v3 += v3
                var v15 float64 = (float64(1e10) + (-2.25 / 1.0))
                v2 = append(v2, 567)
                _ = v15
            }
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = 1
            }
        default:
            fmt.Println(v1, fi4, fi4)
        }
        _ = fi4
    }
    v16 := GenS2{F0: -414, F1: -2147483648}
    fmt.Println(v1, v1)
    var v17 uint8 = 2
    for fi18 := 0; fi18 < 8; fi18++ {
        switch "foo bar" {
        case "\n\t":
            if false {
                v3 = (v3 << ((453) & 63))
                switch -545 {
                case 512:
                    v19 := GenS3{F0: (v17 * (v17 ^ v17)), F1: 18446744073709551615}
                    _ = v19
                case -36:
                    fmt.Println(v1, v1)
                case 0:
                    v1 = -1
                }
            } else {
                var v20 uint = 255
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v21 uint16 = 255
                _ = v20
                _ = v21
            }
        case "hello":
            if (((((float64((1e10 + 0.0)) - 0.0) / -2.25) - ((0.0 - (1e10 + float64(0.0))) + 1.5)) - ((0.1 - float64(1.5)) / 1.0)) != (3.14159 + 1.5)) {
                fmt.Println(v17)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch fi18 {
                case 128:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 429:
                    switch int(v17) {
                    case -1:
                        switch "foo bar" {
                        case "hello":
                            switch "foo bar" {
                            case "foo bar":
                                switch "foo bar" {
                                case "a":
                                    v22 := GenS4{F0: 1, F1: -1}
                                    _ = v22
                                case "hello":
                                    v23 := GenS2{F0: 924, F1: fi18}
                                    _ = v23
                                default:
                                    v17 = v17
                                }
                            case "hello":
                                switch "foo bar" {
                                case "a":
                                    for fi24 := 0; fi24 < 3; fi24++ {
                                        v1 -= v1
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = fi24
                                    }
                                case "foo bar":
                                    if len(v2) > 0 {
                                        v2[0] = 2147483648
                                    }
                                case "\n\t":
                                    var v25 bool = (!false)
                                    _ = v25
                                }
                            case "\n\t":
                                v26 := GenS5{F0: (v3 << ((uint(v3)) & 63))}
                                _ = v26
                            default:
                                var v27 int64 = 0
                                _ = v27
                            }
                        case "foo bar":
                            var v28 uint16 = 1
                            _ = v28
                        case "":
                            if false {
                                v3 = 1
                                for fi29 := 0; fi29 < 9; fi29++ {
                                    var v30 float32 = 0.1
                                    fmt.Println(v1, fi18)
                                    fmt.Println(v1, v17)
                                    v31 := GenS6{F0: (!(v17 > v17))}
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                            } else {
                                if (1e10 != ((3.14159 * 0.1) * 1e10)) {
                                    v2 = append(v2, 2)
                                }
                            }
                        default:
                            fmt.Println(fi18, v17)
                        }
                    case 2147483647:
                        fmt.Println(fi18, fi18)
                    default:
                        if true {
                            v3 = (v3 % (((^(v3))) | 1))
                        } else {
                            if false {
                                if ("hello" > "foo bar") {
                                    var v32 uint16 = 32767
                                    _ = v32
                                }
                                if true {
                                    for fi33 := 0; fi33 < 13; fi33++ {
                                        var v34 uint = 1
                                        _ = fi33
                                        _ = v34
                                    }
                                    var v35 uint16 = 727
                                    v1 -= (v1 & v1)
                                    _ = v35
                                } else {
                                    var v36 uint32 = 656
                                    v17 = 205
                                    if (true || false) {
                                        v37 := GenS7{F0: -1, F1: v1}
                                        _ = v37
                                    }
                                    _ = v36
                                }
                                if len(v2) > 0 {
                                    v2[0] = 2147483647
                                }
                            } else {
                                fmt.Println(v3, v1)
                                var v38 uint8 = 0
                                v38 += (v17 & (v17 % ((v17) | 1)))
                                _ = v38
                            }
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(fi18)
                            }
                        }
                    }
                }
            } else {
                fmt.Println(v17, v1, fi18)
                v39 := GenS8{F0: ""}
                v40 := GenS3{F0: uint8(v3), F1: 183}
                _ = v39
                _ = v40
            }
        }
        _ = fi18
    }
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch 9223372036854775807 {
        case 1:
            v1 = v1
        case -416:
            fmt.Println(v1, v3, v17)
        }
        if false {
            var v41 int = 4294967295
            _ = v41
        } else {
            v42 := GenS1{F0: 32768}
            v43 := GenS2{F0: -288, F1: -977}
            if true {
                var v44 int32 = 42
                fmt.Println(v3, v17, v17)
                switch "foo bar" {
                case "\n\t":
                    v45 := GenS9{F0: 65535}
                    _ = v45
                }
                _ = v44
            }
            _ = v42
            _ = v43
        }
    }
    v17 = v17
    for fi46 := 0; fi46 < 6; fi46++ {
        var v47 uint32 = 0
        switch "\n\t" {
        case "\n\t":
            v2 = append(v2, (v47 >> ((349) & 63)))
        case "a":
            switch (fi46 * 42) {
            case 896:
                switch "a" {
                case "":
                    if (0 <= v3) {
                        var v48 float32 = (0.0 + (3.14159 * 1e10))
                        _ = len(v2)
                        _ = v48
                    }
                case "\n\t":
                    for fi49 := 0; fi49 < 17; fi49++ {
                        v50 := GenS5{F0: uint64(v47)}
                        _ = fi49
                        _ = v50
                    }
                default:
                    if (((fi46 % (((fi46 & 127)) | 1)) >= (fi46 >> ((368) & 63))) || ((false || (fi46 != (fi46 << ((18446744073709551615) & 63)))) && (!true))) {
                        fmt.Println(fi46, v17, v3)
                        var v51 uint16 = 942
                        v52 := GenS6{F0: ((true && (0 >= v3)) || false)}
                        _ = v51
                        _ = v52
                    } else {
                        if (42 < 9223372036854775807) {
                            v17 = (^(v17))
                            fmt.Println(v47)
                            if (1 == 9223372036854775807) {
                                var v53 int32 = 2147483647
                                _ = v53
                            } else {
                                if (false || false) {
                                    if true {
                                        if (false || true) {
                                            fmt.Println(v3)
                                            if (true || false) {
                                                v1 = int8(fi46)
                                                for fi54 := 0; fi54 < 7; fi54++ {
                                                    _ = len(v2)
                                                    var v55 uint = 2
                                                    _ = fi54
                                                    _ = v55
                                                }
                                                switch fi46 {
                                                case 9223372036854775807:
                                                    if true {
                                                        for fi56 := 0; fi56 < 20; fi56++ {
                                                            v2 = append(v2, v47)
                                                            fmt.Println(v3)
                                                            _ = fi56
                                                        }
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v1, fi46)
                                                if true {
                                                    var v57 int = (fi46 << ((983) & 63))
                                                    fmt.Println(v47, v3)
                                                    _ = v57
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = 0
                                                    }
                                                }
                                            }
                                        }
                                        if true {
                                            switch "a" {
                                            case "":
                                                v47 += v47
                                            case "foo bar":
                                                var v58 uint8 = (^(v17))
                                                _ = v58
                                            case "hello":
                                                switch "\n\t" {
                                                case "\n\t":
                                                    _ = len(v2)
                                                default:
                                                    for fi59 := 0; fi59 < 2; fi59++ {
                                                        var v60 uint16 = 255
                                                        v61 := GenS10{F0: fi46, F1: "foo bar", F2: uint8(v3)}
                                                        v62 := GenS9{F0: 42}
                                                        if true {
                                                            v1 += v1
                                                            if (true && (42 >= 1000)) {
                                                                v2 = append(v2, v47)
                                                            }
                                                        }
                                                        _ = fi59
                                                        _ = v60
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                }
                                            default:
                                                for fi63 := 0; fi63 < 19; fi63++ {
                                                    var v64 uint64 = v3
                                                    if (!true) {
                                                        if len(v2) > 0 {
                                                            v2[0] = uint32(v3)
                                                        }
                                                        v64 = 18446744073709551615
                                                        v64 -= v3
                                                        if len(v2) > 0 {
                                                            v2[0] = (v47 / ((116) | 1))
                                                        }
                                                    } else {
                                                        switch fi63 {
                                                        case 0:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        default:
                                                            switch "a" {
                                                            case "\n\t":
                                                                switch fi46 {
                                                                case 2147483647:
                                                                    v1 -= (v1 ^ 0)
                                                                case 4294967295:
                                                                    if true {
                                                                        v3 -= 1
                                                                    } else {
                                                                        v1 = v1
                                                                        var v65 int16 = 575
                                                                        var v66 float32 = (1e10 * 1.5)
                                                                        _ = v65
                                                                        _ = v66
                                                                    }
                                                                case -2147483648:
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v47 = (v47 - (v47 << ((4294967295) & 63)))
                                                                    fmt.Println(v17, v1, v3)
                                                                default:
                                                                    v47 -= 449
                                                                }
                                                            case "foo bar":
                                                                var v67 uint = 811
                                                                _ = v67
                                                            case "":
                                                                _ = len(v2)
                                                            default:
                                                                fmt.Println(v3, v47)
                                                            }
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    _ = len(v2)
                                                    v1 += (v1 / (((v1 - 100)) | 1))
                                                    _ = fi63
                                                    _ = v64
                                                }
                                            }
                                            var v68 uint = 65535
                                            var v69 int8 = v1
                                            _ = v68
                                            _ = v69
                                        }
                                    }
                                    _ = len(v2)
                                    v3 += v3
                                }
                                if len(v2) > 0 {
                                    v2[0] = v47
                                }
                                var v70 int32 = 2147483647
                                _ = v70
                            }
                        }
                        var v71 uint8 = 0
                        if false {
                            switch "hello" {
                            case "\n\t":
                                for fi72 := 0; fi72 < 2; fi72++ {
                                    fmt.Println(fi46, v17)
                                    _ = fi72
                                }
                            case "a":
                                if (2 >= (^(v3))) {
                                    fmt.Println(v17, v3, v71)
                                    v73 := GenS5{F0: (v3 | v3)}
                                    var v74 int8 = v1
                                    _ = v73
                                    _ = v74
                                }
                            default:
                                v17 = v17
                            }
                            v75 := GenS7{F0: 1, F1: (^(v1))}
                            for fi76 := 0; fi76 < 11; fi76++ {
                                var v77 int32 = 42
                                var v78 int = fi46
                                _ = fi76
                                _ = v77
                                _ = v78
                            }
                            _ = v75
                        } else {
                            var v79 uint = 255
                            if len(v2) > 0 {
                                v2[0] = 0
                            }
                            v79 = (^(v79))
                            _ = v79
                        }
                        _ = v71
                    }
                }
            default:
                for fi80 := 0; fi80 < 9; fi80++ {
                    var v81 uint8 = v17
                    if false {
                        if (!(18446744073709551615 == 2)) {
                            var v82 uint16 = 190
                            if (("foo bar" + "foo bar") > (("" + (("foo bar" + "foo bar") + "hello")) + ("foo bar" + "a"))) {
                                fmt.Println(v81)
                                fmt.Println(v17, v3)
                                if (!false) {
                                    v82 -= 573
                                    v2 = append(v2, 2147483648)
                                    var v83 uint = 42
                                    switch ("hello" + "foo bar") {
                                    case "":
                                        var v84 uint16 = (^(v82))
                                        _ = v84
                                    case "hello":
                                        if len(v2) > 0 {
                                            v2[0] = v47
                                        }
                                    default:
                                        if (false && (v82 >= (v82 | (v82 * v82)))) {
                                            fmt.Println(v3)
                                            fmt.Println(v1)
                                        } else {
                                            var v85 uint16 = v82
                                            switch (-(fi80)) {
                                            case 802:
                                                if false {
                                                    v86 := GenS11{F0: int(v1), F1: (fi80 ^ 213), F2: -9223372036854775808}
                                                    v85 = 2
                                                    _ = v86
                                                } else {
                                                    fmt.Println(v3, fi46, v85)
                                                    if true {
                                                        if len(v2) > 0 {
                                                            v2[0] = 2147483647
                                                        }
                                                        var v87 uint = (v83 & (v83 % ((v83) | 1)))
                                                        v2 = append(v2, 42)
                                                        _ = v87
                                                    } else {
                                                        fmt.Println(v82)
                                                        if (182 <= v17) {
                                                            fmt.Println(v47)
                                                        } else {
                                                            v3 = 827
                                                            if (v3 >= v3) {
                                                                var v88 uint = (v83 - v83)
                                                                var v89 float32 = float32(-2.25)
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v90 := GenS5{F0: v3}
                                                                    _ = v90
                                                                case "":
                                                                    var v91 string = "\n\t"
                                                                    _ = v91
                                                                default:
                                                                    v92 := GenS9{F0: 65535}
                                                                    _ = v92
                                                                }
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                            var v93 int = (-(fi80))
                                                            _ = v93
                                                        }
                                                    }
                                                }
                                                v94 := GenS2{F0: -9223372036854775808, F1: (fi80 * (fi80 - (fi80 + (fi80 / ((fi46) | 1)))))}
                                                v1 -= v1
                                                v85 = uint16(v47)
                                                _ = v94
                                            }
                                            switch int(v83) {
                                            case -2147483648:
                                                fmt.Println(v83, v81)
                                            }
                                            _ = v85
                                        }
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = (^(v47))
                                    }
                                    _ = v83
                                } else {
                                    v47 -= 608
                                }
                            } else {
                                switch "hello" {
                                case "\n\t":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "a":
                                    v17 = (v17 >> ((65535) & 63))
                                case "foo bar":
                                    if len(v2) > 0 {
                                        v2[0] = 306
                                    }
                                }
                                v17 = v17
                                var v95 int8 = 127
                                _ = v95
                            }
                            fmt.Println(v47)
                            fmt.Println(v82)
                            _ = v82
                        }
                        v3 -= (v3 ^ (v3 | v3))
                        var v96 bool = ((-2.25 <= 0.1) || (false && true))
                        if true {
                            var v97 bool = (42 <= v3)
                            _ = v97
                        } else {
                            v47 -= v47
                            var v98 int = (fi80 + 2147483647)
                            _ = v98
                        }
                        _ = v96
                    } else {
                        v1 = (v1 << ((uint(v17)) & 63))
                        fmt.Println(v17, fi46, v17)
                    }
                    _ = fi80
                    _ = v81
                }
            }
        case "":
            v1 = (v1 - v1)
        default:
            var v99 int16 = 0
            _ = v99
        }
        _ = fi46
        _ = v47
    }
    for fi100 := 0; fi100 < 19; fi100++ {
        v2 = append(v2, 65535)
        _ = fi100
    }
    fmt.Println(v17, v1, v1)
    _ = len(v2)
    for fi101 := 0; fi101 < 2; fi101++ {
        var v102 int16 = 1
        fmt.Println(v17)
        v2 = append(v2, 533)
        _ = fi101
        _ = v102
    }
    if (542 <= -1) {
        v17 = v17
        v2 = append(v2, 65535)
    }
    switch "hello" {
    case "a":
        if ((v3 >> ((328) & 63)) <= (v3 / (((v3 % ((42) | 1))) | 1))) {
            for fi103 := 0; fi103 < 11; fi103++ {
                v3 = 42
                v1 -= -62
                _ = fi103
            }
        }
    case "":
        var v104 bool = (1.5 >= float64(1e10))
        _ = v104
    default:
        for fi105 := 0; fi105 < 14; fi105++ {
            v17 -= v17
            _ = fi105
        }
    }
    v17 = 128
    if (284 < 65535) {
        v106 := GenS7{F0: 1, F1: 1}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v106
    } else {
        v17 = 42
    }
    v107 := GenS7{F0: -1, F1: (v1 << ((4294967295) & 63))}
    for fi108 := 0; fi108 < 4; fi108++ {
        v17 -= 2
        for fi109 := 0; fi109 < 19; fi109++ {
            v1 = (v1 >> ((65535) & 63))
            _ = fi109
        }
        _ = fi108
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v107
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
