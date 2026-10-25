package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int; F2 uint64 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 int; F2 int64 }
type GenS6 struct { F0 uint64; F1 int64; F2 uint }
type GenS7 struct { F0 uint32; F1 int8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16; F1 uint8 }
type GenS10 struct { F0 uint; F1 uint32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int32; F1 int16; F2 int }
type GenS13 struct { F0 uint16; F1 int8 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int8; F1 int8 }
type GenS16 struct { F0 int; F1 uint }
func genFn0(p0 int64, p1 int64) string {
    if ((1 != 127) || true) {
        v7 := GenS1{F0: 18446744073709551615}
        p0 -= (-(p1))
        _ = v7
    }
    var v8 int64 = p1
    if false {
        var v9 int = 32768
        _ = v9
    } else {
        switch -9223372036854775808 {
        case 654:
            var v10 int16 = -1
            _ = v10
        case 596:
            var v11 float32 = 1.5
            _ = v11
        case 127:
            var v12 uint = 657
            _ = v12
        default:
            if true {
                v13 := GenS1{F0: 4294967295}
                _ = v13
            } else {
                v14 := GenS2{F0: 104, F1: 0, F2: 0}
                _ = v14
            }
        }
        v15 := GenS2{F0: 106, F1: 9223372036854775807, F2: 18446744073709551615}
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 bool = true
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v8
    return ((("hello" + "a") + "\n\t") + "hello")
}
func genFn1(p0 uint16, p1 int32) uint {
    v18 := GenS2{F0: 127, F1: -278, F2: uint64(p0)}
    var v19 int16 = -32768
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return uint(p0)
}

func main() {
    var v1 float32 = 0.0
    var v2 uint16 = 0
    var v3 int16 = -837
    var v4 uint16 = (^(v2))
    var v5 uint32 = 2
    var v6 int = -1
    if false {
        v20 := []int{32768, (v6 | (v6 % ((v6) | 1))), (v6 | 0), 128}
        _ = v20
    } else {
        switch (v6 | v6) {
        case 936:
            v21 := GenS2{F0: 127, F1: 32767, F2: 0}
            _ = v21
        case -1:
            v22 := []int{(^(v6))}
            _ = v22
        case 352:
            v23 := GenS1{F0: 0}
            _ = v23
        }
        if true {
            v6 = v6
            var v24 float32 = (1e10 + float32(((float64((-2.25 + (0.1 / v1))) / 1.5) + 1.5)))
            _ = v24
        } else {
            v25 := GenS2{F0: 76, F1: v6, F2: 228}
            var v26 float64 = float64(0.1)
            var v27 float64 = v26
            _ = v25
            _ = v26
            _ = v27
        }
    }
    if true {
        v28 := GenS3{F0: -1, F1: 0}
        v29 := GenS4{F0: 65535}
        v30 := GenS5{F0: uint(v5), F1: v6, F2: 42}
        _ = v28
        _ = v29
        _ = v30
    } else {
        for fi31 := 0; fi31 < 20; fi31++ {
            v32 := GenS4{F0: 65535}
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 5; fi33++ {
            v34 := GenS3{F0: -90, F1: 1}
            var v35 bool = (float32((1e10 - (0.1 * (3.14159 + 0.1)))) <= float32(3.14159))
            var v36 float64 = (1.5 / (1e10 * float64(1.5)))
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
    }
    v37 := GenS1{F0: 4294967295}
    v3 = v3
    for fi38 := 0; fi38 < 3; fi38++ {
        var v39 float32 = v1
        _ = fi38
        _ = v39
    }
    var v40 uint = 255
    if ((true && true) || false) {
        for fi41 := 0; fi41 < 2; fi41++ {
            v4 = 0
            if false {
                v40 += 105
                for fi42 := 0; fi42 < 3; fi42++ {
                    v3 -= v3
                    var v43 int16 = (v3 << (((v40 / ((v40) | 1))) & 63))
                    var v44 int64 = -27
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                var v45 uint16 = 1
                _ = v45
            }
            _ = fi41
        }
        for fi46 := 0; fi46 < 19; fi46++ {
            v2 = (^(v4))
            var v47 uint8 = 1
            for fi48 := 0; fi48 < 14; fi48++ {
                fmt.Println(fi48)
                _ = fi48
            }
            _ = fi46
            _ = v47
        }
    }
    if (("\n\t" >= "foo bar") || (-1 < -721)) {
        if false {
            fmt.Println(v1, v5, v40)
        } else {
            v4 -= 32767
            v49 := []int{}
            var v50 int8 = 58
            _ = v49
            _ = v50
        }
    } else {
        if true {
            var v51 uint = v40
            fmt.Println(v5, v4, v5)
            for fi52 := 0; fi52 < 4; fi52++ {
                var v53 float64 = ((0.1 - 1e10) / 1.5)
                _ = fi52
                _ = v53
            }
            _ = v51
        } else {
            v54 := []int{0, 32768, v6}
            switch (("foo bar" + ("hello" + "\n\t")) + "") {
            case "a":
                var v55 int8 = 127
                _ = v55
            case "hello":
                if len(v54) > 0 {
                    v54[0] = 127
                }
            case "":
                for fi56 := 0; fi56 < 3; fi56++ {
                    if (!(fi56 >= int(v3))) {
                        for fi57 := 0; fi57 < 3; fi57++ {
                            if (false && false) {
                                var v58 int = fi57
                                _ = v58
                            } else {
                                v54 = append(v54, (fi56 << (((v40 - (v40 / (((v40 - genFn1(v4, 1))) | 1)))) & 63)))
                                if (!(42 == 0)) {
                                    fmt.Println(v2)
                                    _ = len(v54)
                                }
                            }
                            switch ("foo bar" + "a") {
                            case "":
                                if (false || (42 == 0)) {
                                    var v59 uint = 716
                                    _ = v59
                                } else {
                                    if (1e10 < -2.25) {
                                        if false {
                                            var v60 int64 = 142
                                            fmt.Println(v60)
                                            _ = v60
                                        } else {
                                            v3 = int16(fi57)
                                            if false {
                                                v61 := GenS2{F0: 127, F1: (fi56 + int(v4)), F2: 2}
                                                switch (fi57 - fi56) {
                                                case 287:
                                                    switch (fi57 & fi57) {
                                                    case 42:
                                                        fmt.Println(v6, v5, fi56)
                                                    case -772:
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v3 = (v3 >> (((v40 / (((v40 ^ (v40 ^ genFn1((v4 >> (((v40 >> (((v40 ^ v40)) & 63))) & 63)), 0)))) | 1))) & 63))
                                                        case "":
                                                            v4 = (v4 % ((v2) | 1))
                                                        default:
                                                            var v62 float32 = v1
                                                            _ = v62
                                                        }
                                                    case 4294967295:
                                                        var v63 float64 = ((((0.1 + 0.1) / 1.5) * 0.0) / -2.25)
                                                        _ = v63
                                                    default:
                                                        if len(v54) > 0 {
                                                            v54[0] = 0
                                                        }
                                                    }
                                                case -9223372036854775808:
                                                    fmt.Println(v5, v3, v1)
                                                case 21:
                                                    if len(v54) > 0 {
                                                        v54[0] = fi56
                                                    }
                                                default:
                                                    var v64 string = ("\n\t" + "")
                                                    _ = v64
                                                }
                                                _ = v61
                                            } else {
                                                v65 := GenS6{F0: 2, F1: 256, F2: (v40 * (^(v40)))}
                                                v2 += v4
                                                var v66 uint = (v40 / ((genFn1((v2 ^ 0), -2147483648)) | 1))
                                                _ = v65
                                                _ = v66
                                            }
                                            if true {
                                                v40 += 114
                                                if false {
                                                    if (!((-(v3)) > v3)) {
                                                        var v67 uint = v40
                                                        switch ("hello" + ("\n\t" + "hello")) {
                                                        case "\n\t":
                                                            var v68 uint16 = (v2 >> (((^(v40))) & 63))
                                                            _ = v68
                                                        case "hello":
                                                            v40 -= (^(v67))
                                                        case "":
                                                            if len(v54) > 0 {
                                                                v54[0] = (fi57 + fi57)
                                                            }
                                                        default:
                                                            v5 = (v5 * v5)
                                                        }
                                                        v54 = append(v54, (fi57 - (fi57 - fi57)))
                                                        _ = v67
                                                    } else {
                                                        v3 = 0
                                                        if true {
                                                            if (true || false) {
                                                                if false {
                                                                    if true {
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            v6 = v6
                                                                            var v69 uint16 = 0
                                                                            var v70 int32 = -1
                                                                            _ = v69
                                                                            _ = v70
                                                                        case "a":
                                                                            var v71 uint32 = v5
                                                                            _ = v71
                                                                        case "":
                                                                            if (v5 < 65535) {
                                                                                v72 := GenS1{F0: 963}
                                                                                switch 42 {
                                                                                case 2147483647:
                                                                                    fmt.Println(v1, v6)
                                                                                default:
                                                                                    switch (("a" + "\n\t") + (("foo bar" + "") + "foo bar")) {
                                                                                    case "hello":
                                                                                        v73 := GenS7{F0: 0, F1: int8(v2)}
                                                                                        _ = v73
                                                                                    case "":
                                                                                        if true {
                                                                                            var v74 uint = 42
                                                                                            if len(v54) > 0 {
                                                                                                v54[0] = fi57
                                                                                            }
                                                                                            if (0 > 0) {
                                                                                                fmt.Println(v4)
                                                                                            } else {
                                                                                                if false {
                                                                                                    var v75 int16 = v3
                                                                                                    fmt.Println(v6)
                                                                                                    fmt.Println(v3)
                                                                                                    _ = v75
                                                                                                } else {
                                                                                                    switch fi57 {
                                                                                                    case -1:
                                                                                                        var v76 float32 = v1
                                                                                                        v6 -= int(v3)
                                                                                                        if (9223372036854775807 >= -480) {
                                                                                                            var v77 uint = (v74 & (v40 ^ uint(v5)))
                                                                                                            _ = v77
                                                                                                        }
                                                                                                        _ = v76
                                                                                                    case 2147483647:
                                                                                                        switch 128 {
                                                                                                        case -742:
                                                                                                            v1 = 3.14159
                                                                                                        case 1:
                                                                                                            fmt.Println(v4, fi56, fi56)
                                                                                                        default:
                                                                                                            switch (("\n\t" + "\n\t") + ("\n\t" + "")) {
                                                                                                            case "foo bar":
                                                                                                                v3 = (v3 << ((v74) & 63))
                                                                                                            case "\n\t":
                                                                                                                v6 = 532
                                                                                                            case "":
                                                                                                                v78 := GenS2{F0: 0, F1: (v6 ^ (v6 * v6)), F2: 0}
                                                                                                                v79 := GenS2{F0: uint8(v74), F1: v6, F2: 42}
                                                                                                                _ = v78
                                                                                                                _ = v79
                                                                                                            default:
                                                                                                                var v80 uint64 = 42
                                                                                                                _ = v80
                                                                                                            }
                                                                                                        }
                                                                                                    case 0:
                                                                                                        switch (v6 / ((127) | 1)) {
                                                                                                        case 2147483647:
                                                                                                            var v81 int32 = 42
                                                                                                            _ = v81
                                                                                                        }
                                                                                                    default:
                                                                                                        var v82 int64 = 42
                                                                                                        _ = v82
                                                                                                    }
                                                                                                }
                                                                                                var v83 int16 = (v3 << ((v74) & 63))
                                                                                                v1 = v1
                                                                                                var v84 int16 = v83
                                                                                                var v85 int8 = 1
                                                                                                _ = v83
                                                                                                _ = v84
                                                                                                _ = v85
                                                                                            }
                                                                                            _ = v74
                                                                                        } else {
                                                                                            fmt.Println(fi57)
                                                                                            if (1 <= 42) {
                                                                                                switch 457 {
                                                                                                case 128:
                                                                                                    var v86 string = "hello"
                                                                                                    _ = v86
                                                                                                case -1:
                                                                                                    if ((-(v3)) >= -32768) {
                                                                                                        switch fi56 {
                                                                                                        case -855:
                                                                                                            fmt.Println(v5, v40, v40)
                                                                                                        case 4294967295:
                                                                                                            switch "hello" {
                                                                                                            case "":
                                                                                                                var v87 int64 = 9223372036854775807
                                                                                                                _ = v87
                                                                                                            }
                                                                                                        case 127:
                                                                                                            var v88 uint64 = 2
                                                                                                            var v89 bool = ((fi57 * (fi57 & int(v88))) <= -21)
                                                                                                            fmt.Println(v40)
                                                                                                            _ = v88
                                                                                                            _ = v89
                                                                                                        }
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            v40 = 0
                                                                                                        }
                                                                                                        v3 -= int16(v6)
                                                                                                        if len(v54) >= 2 {
                                                                                                            v54 = v54[1:len(v54)]
                                                                                                        }
                                                                                                        v90 := GenS2{F0: 127, F1: fi56, F2: 18446744073709551615}
                                                                                                        _ = v90
                                                                                                    }
                                                                                                case 4294967295:
                                                                                                    switch 32768 {
                                                                                                    case 32767:
                                                                                                        if ((-2.25 + float32((1e10 / float64(v1)))) < 1e10) {
                                                                                                            var v91 string = ("" + "foo bar")
                                                                                                            v54 = append(v54, (fi56 / (((fi56 | v6)) | 1)))
                                                                                                            switch ("foo bar" + "foo bar") {
                                                                                                            case "\n\t":
                                                                                                                var v92 uint32 = v5
                                                                                                                _ = v92
                                                                                                            case "hello":
                                                                                                                v40 = v40
                                                                                                            case "foo bar":
                                                                                                                switch -509 {
                                                                                                                case 42:
                                                                                                                    switch v6 {
                                                                                                                    case 9223372036854775807:
                                                                                                                        v5 = v5
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                var v93 uint32 = v5
                                                                                                                _ = v93
                                                                                                            }
                                                                                                            fmt.Println(v6)
                                                                                                            _ = v91
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                if (-119 == -10) {
                                                                                                                    v94 := GenS8{F0: 555}
                                                                                                                    v4 = (v2 % ((873) | 1))
                                                                                                                    if false {
                                                                                                                        v3 -= v3
                                                                                                                        switch (fi56 * -9223372036854775808) {
                                                                                                                        case 754:
                                                                                                                            v95 := GenS5{F0: (v40 - v40), F1: 730, F2: 563}
                                                                                                                            _ = v95
                                                                                                                        default:
                                                                                                                            v1 = 1e10
                                                                                                                        }
                                                                                                                        v96 := GenS3{F0: 42, F1: 718}
                                                                                                                        switch "" {
                                                                                                                        case "a":
                                                                                                                            if len(v54) >= 2 {
                                                                                                                                v54 = v54[1:len(v54)]
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            if len(v54) > 0 {
                                                                                                                                v54[0] = (fi56 << ((genFn1((^(v2)), int32(v40))) & 63))
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = v96
                                                                                                                    } else {
                                                                                                                        var v97 int64 = 81
                                                                                                                        v3 += v3
                                                                                                                        switch "\n\t" {
                                                                                                                        case "foo bar":
                                                                                                                            if false {
                                                                                                                                v6 = -130
                                                                                                                            } else {
                                                                                                                                var v98 int8 = 0
                                                                                                                                _ = v98
                                                                                                                            }
                                                                                                                        case "a":
                                                                                                                            v99 := GenS5{F0: v40, F1: (v6 + (fi56 * v6)), F2: 0}
                                                                                                                            _ = v99
                                                                                                                        case "":
                                                                                                                            switch (fi57 + v6) {
                                                                                                                            case 1:
                                                                                                                                v3 = (v3 + v3)
                                                                                                                                v2 += v4
                                                                                                                            case 0:
                                                                                                                                v100 := GenS9{F0: 255, F1: 1}
                                                                                                                                _ = v100
                                                                                                                            default:
                                                                                                                                switch (fi56 | (fi56 % ((fi56) | 1))) {
                                                                                                                                case 2147483647:
                                                                                                                                    var v101 uint8 = 176
                                                                                                                                    _ = v101
                                                                                                                                case 277:
                                                                                                                                    v40 += 18446744073709551615
                                                                                                                                default:
                                                                                                                                    if len(v54) >= 2 {
                                                                                                                                        v54 = v54[1:len(v54)]
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        var v102 uint64 = 203
                                                                                                                        v97 -= v97
                                                                                                                        var v103 uint16 = (v2 / ((315) | 1))
                                                                                                                        _ = v102
                                                                                                                        _ = v103
                                                                                                                        _ = v97
                                                                                                                    }
                                                                                                                    _ = v94
                                                                                                                } else {
                                                                                                                    fmt.Println(v40)
                                                                                                                    _ = len(v54)
                                                                                                                    v4 -= 816
                                                                                                                    fmt.Println(v2, v1, v4)
                                                                                                                }
                                                                                                                var v104 string = "foo bar"
                                                                                                                switch v104 {
                                                                                                                case "foo bar":
                                                                                                                    v6 = 0
                                                                                                                }
                                                                                                                var v105 int = (fi56 - fi57)
                                                                                                                _ = v104
                                                                                                                _ = v105
                                                                                                            }
                                                                                                            var v106 string = ""
                                                                                                            v6 = (v6 / (((^(fi56))) | 1))
                                                                                                            if true {
                                                                                                                var v107 uint16 = 54
                                                                                                                var v108 uint = 4294967295
                                                                                                                if len(v54) >= 2 {
                                                                                                                    v54 = v54[1:len(v54)]
                                                                                                                }
                                                                                                                _ = v107
                                                                                                                _ = v108
                                                                                                            } else {
                                                                                                                var v109 int8 = 0
                                                                                                                _ = v109
                                                                                                            }
                                                                                                            _ = v106
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = len(v54)
                                                                                            } else {
                                                                                                fmt.Println(v40, v3, v4)
                                                                                                var v110 int = (^(v6))
                                                                                                var v111 float32 = v1
                                                                                                _ = v110
                                                                                                _ = v111
                                                                                            }
                                                                                            v112 := GenS10{F0: 255, F1: (v5 ^ v5)}
                                                                                            if false {
                                                                                                if false {
                                                                                                    v4 = 1
                                                                                                    v113 := GenS11{F0: uint32(v40)}
                                                                                                    v54 = append(v54, 128)
                                                                                                    v3 = v3
                                                                                                    if (!true) {
                                                                                                        v54 = append(v54, -9223372036854775808)
                                                                                                    }
                                                                                                    _ = v113
                                                                                                }
                                                                                                fmt.Println(v4, fi56)
                                                                                            }
                                                                                            if len(v54) >= 2 {
                                                                                                v54 = v54[1:len(v54)]
                                                                                            }
                                                                                            if len(v54) > 0 {
                                                                                                v54[0] = fi56
                                                                                            }
                                                                                            v5 = v5
                                                                                            _ = v112
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v114 float32 = float32(-2.25)
                                                                                _ = v114
                                                                                _ = v72
                                                                            }
                                                                        }
                                                                        var v115 uint32 = 868
                                                                        _ = v115
                                                                    }
                                                                    if (!false) {
                                                                        fmt.Println(v4)
                                                                        fmt.Println(fi56, v2)
                                                                    } else {
                                                                        if (false || (v1 <= v1)) {
                                                                            fmt.Println(fi57, fi56, fi56)
                                                                        }
                                                                    }
                                                                    fmt.Println(v1, v1)
                                                                }
                                                                fmt.Println(fi57)
                                                                if true {
                                                                    _ = len(v54)
                                                                    var v116 float32 = (0.1 + (float32(1.5) + -2.25))
                                                                    _ = v116
                                                                } else {
                                                                    v3 += v3
                                                                }
                                                            } else {
                                                                switch (fi57 % (((v6 ^ fi56)) | 1)) {
                                                                case 32768:
                                                                    v117 := GenS1{F0: 0}
                                                                    _ = v117
                                                                case -447:
                                                                    fmt.Println(fi57)
                                                                case 2147483647:
                                                                    if (42 != -128) {
                                                                        v4 -= v4
                                                                        var v118 float32 = v1
                                                                        _ = v118
                                                                    }
                                                                default:
                                                                    var v119 int64 = 969
                                                                    _ = v119
                                                                }
                                                            }
                                                            if (1e10 >= (0.1 * float32(float64(v1)))) {
                                                                fmt.Println(v40, fi57)
                                                            } else {
                                                                var v120 int32 = 42
                                                                if false {
                                                                    if false {
                                                                        switch 42 {
                                                                        case -560:
                                                                            _ = len(v54)
                                                                        case -1:
                                                                            var v121 uint8 = 123
                                                                            _ = v121
                                                                        case 32767:
                                                                            fmt.Println(v120)
                                                                        }
                                                                        v4 -= (v2 | (v2 % ((1) | 1)))
                                                                        v40 = 0
                                                                    } else {
                                                                        v2 += 1
                                                                        v3 += (v3 * 32767)
                                                                        if len(v54) >= 2 {
                                                                            v54 = v54[1:len(v54)]
                                                                        }
                                                                        if false {
                                                                            fmt.Println(v6, v40)
                                                                            var v122 int8 = 5
                                                                            if ((v3 | 1) > (v3 * 32767)) {
                                                                                var v123 uint16 = uint16(v5)
                                                                                if false {
                                                                                    var v124 int = fi56
                                                                                    _ = v124
                                                                                }
                                                                                var v125 int32 = v120
                                                                                _ = v123
                                                                                _ = v125
                                                                            } else {
                                                                                v4 = uint16(v5)
                                                                                v126 := GenS9{F0: uint16(v40), F1: 183}
                                                                                if (false || (!(genFn1((v4 >> ((genFn1(178, (v120 << (((v40 >> ((1) & 63))) & 63)))) & 63)), int32(v122)) < (^(v40))))) {
                                                                                    fmt.Println(v1)
                                                                                    if false {
                                                                                        fmt.Println(v2)
                                                                                        v3 = (v3 - (v3 ^ 1000))
                                                                                    }
                                                                                }
                                                                                _ = v126
                                                                            }
                                                                            _ = v122
                                                                        }
                                                                        var v127 uint32 = (v5 & 2147483648)
                                                                        _ = v127
                                                                    }
                                                                }
                                                                var v128 uint8 = 40
                                                                _ = v120
                                                                _ = v128
                                                            }
                                                            var v129 uint8 = uint8(v5)
                                                            _ = v129
                                                        }
                                                    }
                                                    v40 = (v40 & 0)
                                                }
                                            }
                                        }
                                        _ = len(v54)
                                    } else {
                                        if true {
                                            fmt.Println(fi56)
                                        }
                                        v2 = v2
                                    }
                                    v54 = append(v54, 203)
                                    if false {
                                        v6 -= fi57
                                        var v130 uint8 = 1
                                        switch (v6 + fi56) {
                                        case -138:
                                            v131 := GenS12{F0: 1, F1: (v3 - (v3 ^ v3)), F2: (fi57 ^ (fi57 - int(v40)))}
                                            _ = v131
                                        case 127:
                                            var v132 uint = v40
                                            _ = v132
                                        }
                                        _ = v130
                                    }
                                }
                            default:
                                v40 -= 832
                                switch fi56 {
                                case 32767:
                                    if false {
                                        switch (fi56 | fi57) {
                                        case -862:
                                            if true {
                                                v133 := GenS13{F0: 299, F1: int8(v3)}
                                                v2 = uint16(v5)
                                                _ = v133
                                            } else {
                                                fmt.Println(v4)
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v134 uint32 = (v5 ^ v5)
                                                    _ = v134
                                                }
                                                if len(v54) > 0 {
                                                    v54[0] = (-(fi57))
                                                }
                                            }
                                        case 263:
                                            v5 += v5
                                        }
                                    } else {
                                        v135 := GenS10{F0: v40, F1: v5}
                                        v1 = (v1 / 1e10)
                                        _ = v135
                                    }
                                default:
                                    v5 = 977
                                }
                            }
                            _ = fi57
                        }
                    }
                    var v136 int16 = (v3 | int16(fi56))
                    _ = fi56
                    _ = v136
                }
            default:
                v137 := GenS1{F0: 1}
                _ = v137
            }
            var v138 float32 = v1
            _ = v138
            _ = v54
        }
    }
    switch 32767 {
    case 9223372036854775807:
        if (127 != uint8(v40)) {
            for fi139 := 0; fi139 < 11; fi139++ {
                v140 := GenS14{F0: (^(v5))}
                var v141 bool = false
                _ = fi139
                _ = v140
                _ = v141
            }
            fmt.Println(v2, v2)
            switch "a" {
            case "a":
                switch (v6 % ((127) | 1)) {
                case 4294967295:
                    if true {
                        if false {
                            for fi142 := 0; fi142 < 4; fi142++ {
                                fmt.Println(v6, v2)
                                var v143 uint8 = 2
                                _ = fi142
                                _ = v143
                            }
                        }
                        for fi144 := 0; fi144 < 8; fi144++ {
                            fmt.Println(v3, v4)
                            v4 -= v2
                            _ = fi144
                        }
                        v145 := GenS5{F0: (^(v40)), F1: (v6 - v6), F2: 79}
                        _ = v145
                    }
                default:
                    if true {
                        var v146 float32 = v1
                        v147 := []int{-1, v6}
                        _ = v146
                        _ = v147
                    } else {
                        v148 := []int{(v6 / ((v6) | 1))}
                        var v149 int8 = 100
                        _ = v148
                        _ = v149
                    }
                }
            case "":
                v3 = (^(v3))
            }
        }
    case 1:
        for fi150 := 0; fi150 < 7; fi150++ {
            for fi151 := 0; fi151 < 14; fi151++ {
                var v152 float32 = 0.0
                v153 := GenS4{F0: 576}
                _ = fi151
                _ = v152
                _ = v153
            }
            switch "hello" {
            case "":
                v5 = (^(v5))
            }
            _ = fi150
        }
    case 773:
        v2 -= v4
    default:
        v154 := GenS8{F0: 388}
        _ = v154
    }
    switch "a" {
    case "foo bar":
        var v155 uint = genFn1(v2, 2147483647)
        _ = v155
    case "hello":
        for fi156 := 0; fi156 < 8; fi156++ {
            v40 = v40
            _ = fi156
        }
    case "\n\t":
        var v157 string = "\n\t"
        _ = v157
    default:
        v6 = v6
    }
    var v158 int = (v6 / (((v6 ^ (v6 - v6))) | 1))
    if ((1.5 - (((v1 + (float32(3.14159) * ((((v1 / v1) - v1) / (v1 * (v1 - (1.5 - v1)))) / float32(float64(((-2.25 - v1) * 3.14159)))))) + (((v1 - (v1 + v1)) * v1) - 1.5)) * (-2.25 * 0.0))) <= ((0.0 - (float32((float64(float32((-2.25 - 3.14159))) * 0.1)) / v1)) / 1.0)) {
        fmt.Println(v2)
        v5 = 1
        var v159 string = "hello"
        _ = v159
    } else {
        for fi160 := 0; fi160 < 19; fi160++ {
            if (true || true) {
                switch (-(v6)) {
                case 0:
                    var v161 int = int(v40)
                    _ = v161
                case 788:
                    fmt.Println(v158, v3)
                default:
                    v158 = v158
                }
                v158 = 32767
            }
            v162 := GenS15{F0: 100, F1: 63}
            for fi163 := 0; fi163 < 19; fi163++ {
                v164 := []int{(fi160 + (fi163 - int(v3))), (fi160 << ((genFn1(0, 2147483647)) & 63))}
                _ = fi163
                _ = v164
            }
            _ = fi160
            _ = v162
        }
        var v165 int8 = 125
        _ = v165
    }
    v4 -= (v2 << ((18446744073709551615) & 63))
    if false {
        v5 = v5
        v158 = v6
    } else {
        v1 += 1.5
        if false {
            for fi166 := 0; fi166 < 20; fi166++ {
                for fi167 := 0; fi167 < 3; fi167++ {
                    var v168 uint64 = 2
                    switch (-(v6)) {
                    case -2147483648:
                        var v169 float64 = 0.1
                        _ = v169
                    }
                    _ = fi167
                    _ = v168
                }
                var v170 uint8 = 183
                _ = fi166
                _ = v170
            }
        } else {
            for fi171 := 0; fi171 < 15; fi171++ {
                switch fi171 {
                case 32768:
                    if false {
                        v2 -= 425
                        if (!(false || (v3 > v3))) {
                            if true {
                                for fi172 := 0; fi172 < 16; fi172++ {
                                    var v173 uint16 = v4
                                    _ = fi172
                                    _ = v173
                                }
                            }
                            v174 := GenS16{F0: v158, F1: 0}
                            _ = v174
                        } else {
                            for fi175 := 0; fi175 < 14; fi175++ {
                                var v176 bool = true
                                v40 = genFn1((v2 ^ 42), -2147483648)
                                switch "foo bar" {
                                case "\n\t":
                                    v177 := []int{}
                                    _ = v177
                                case "hello":
                                    v176 = v176
                                case "foo bar":
                                    if v176 {
                                        fmt.Println(fi171, v176, fi175)
                                    }
                                default:
                                    fmt.Println(v40, fi171)
                                }
                                _ = fi175
                                _ = v176
                            }
                        }
                    }
                }
                var v178 uint8 = 2
                _ = fi171
                _ = v178
            }
        }
        v4 += (^(v4))
    }
    var v179 uint64 = 676
    fmt.Println("v1:", v1)
    fmt.Println("v158:", v158)
    fmt.Println("v179:", v179)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v158
    _ = v179
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
