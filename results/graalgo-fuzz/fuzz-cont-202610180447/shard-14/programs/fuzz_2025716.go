package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS10 struct { F0 uint64; F1 uint32 }
type GenS11 struct { F0 int16; F1 int32 }
type GenS12 struct { F0 bool; F1 int8; F2 int }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int8; F1 string; F2 int32 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint64; F1 string }
type GenS18 struct { F0 bool; F1 uint64 }
type GenS19 struct { F0 int; F1 int8 }
type GenS20 struct { F0 int }
type GenS21 struct { F0 int32; F1 int }
type GenS22 struct { F0 uint8 }
type GenS23 struct { F0 uint; F1 int64 }
type GenS24 struct { F0 int8; F1 uint64 }
type GenS25 struct { F0 uint8 }
type GenS26 struct { F0 int8 }
type GenS27 struct { F0 int8; F1 int16; F2 int64 }
type GenS28 struct { F0 int16; F1 int8 }
type GenS29 struct { F0 int8; F1 uint }
type GenS30 struct { F0 uint8; F1 uint }
type GenS31 struct { F0 uint; F1 bool; F2 int16 }
type GenS32 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS33 struct { F0 uint16 }
type GenS34 struct { F0 int16; F1 int16 }
type GenS35 struct { F0 bool }
type GenS36 struct { F0 int16; F1 int8 }
type GenS37 struct { F0 uint32; F1 int64 }
type GenS38 struct { F0 int64; F1 int16 }
type GenS39 struct { F0 uint64 }
type GenS40 struct { F0 uint16; F1 uint }
type GenS41 struct { F0 uint32; F1 int; F2 int8 }
type GenS42 struct { F0 int64 }
type GenS43 struct { F0 uint8; F1 int64 }
type GenS44 struct { F0 bool; F1 int; F2 int16 }
type GenS45 struct { F0 int32 }
type GenS46 struct { F0 int32 }
type GenS47 struct { F0 uint32 }
type GenS48 struct { F0 uint32; F1 uint32 }
type GenS49 struct { F0 int16; F1 int32; F2 int64 }
type GenS50 struct { F0 int32 }
type GenS51 struct { F0 int64; F1 int; F2 uint16 }
type GenS52 struct { F0 bool; F1 uint }
type GenS53 struct { F0 int16; F1 uint64; F2 bool }
type GenS54 struct { F0 bool; F1 int }
type GenS55 struct { F0 bool }
type GenS56 struct { F0 int8 }
type GenS57 struct { F0 int16 }
type GenS58 struct { F0 int16; F1 int64 }
type GenS59 struct { F0 int64; F1 int32 }
type GenS60 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS61 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS62 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS63 struct { F0 int }

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    var v4 int8 = -128
    fmt.Println(v4, v1)
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int32 = 42
        var v7 int64 = 1
        v7 += v7
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 uint = 410
            fmt.Println(fi8, v1)
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            switch "foo bar" {
            case "\n\t":
                var v11 int32 = 1
                _ = v11
            }
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 int8 = (-(v4))
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 17; fi14++ {
                v4 -= 1
                if (true || false) {
                    v4 += (v4 * 118)
                } else {
                    if true {
                        fmt.Println(v1, fi14)
                        v15 := GenS1{F0: 1}
                        _ = v15
                    } else {
                        v16 := GenS1{F0: 3}
                        v17 := []int{32767, fi10}
                        if (!false) {
                            v4 = v4
                            if true {
                                if (!(false || true)) {
                                    fmt.Println(v1, fi10)
                                    fmt.Println(fi10, v4)
                                    fmt.Println(fi10, fi14, fi14)
                                    var v18 int = 127
                                    _ = v18
                                }
                                switch (fi10 + fi10) {
                                case 127:
                                    fmt.Println(fi10)
                                case 32768:
                                    fmt.Println(fi10, fi14)
                                case 45:
                                    fmt.Println(fi14)
                                }
                                fmt.Println(v4, fi10, fi10)
                            } else {
                                var v19 float64 = 0.1
                                v20 := GenS1{F0: 1}
                                v1 = 1
                                if true {
                                    if (561 == v1) {
                                        fmt.Println(v1, fi14)
                                        v21 := GenS1{F0: 104}
                                        if true {
                                            var v22 uint = 618
                                            _ = v22
                                        }
                                        _ = v21
                                    }
                                    var v23 int32 = 2147483647
                                    _ = v23
                                } else {
                                    if len(v17) > 0 {
                                        v17[0] = 32768
                                    }
                                }
                                v1 -= v1
                                _ = v19
                                _ = v20
                            }
                            if true {
                                if true {
                                    var v24 int = (fi14 % (((fi10 ^ fi14)) | 1))
                                    if (false && true) {
                                        v24 += (fi10 + (fi14 / ((fi14) | 1)))
                                        if false {
                                            v3[-2147483648] = 454
                                            var v25 float64 = (0.1 - (1e10 + (1e10 * 3.14159)))
                                            if (0.0 <= 3.14159) {
                                                v4 = (v4 % ((v4) | 1))
                                                var v26 uint16 = 346
                                                _ = v26
                                            } else {
                                                var v27 int = fi14
                                                _ = len(v17)
                                                v1 += v1
                                                var v28 bool = true
                                                if (v28 && (32767 == 2)) {
                                                    v27 = 32767
                                                    v27 = 128
                                                }
                                                _ = v27
                                                _ = v28
                                            }
                                            _ = v25
                                        } else {
                                            switch v24 {
                                            case 4294967295:
                                                fmt.Println(v4, v1)
                                            default:
                                                delete(v3, fi14)
                                            }
                                            delete(v3, (^(fi14)))
                                            fmt.Println(v4)
                                            switch (fi10 % ((v24) | 1)) {
                                            case 42:
                                                switch fi14 {
                                                case 1:
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v29 uint64 = 0
                                                        _ = v29
                                                    default:
                                                        if len(v17) > 0 {
                                                            v17[0] = 1
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    if ((v4 == v4) || ((float32(-2.25) >= (1e10 * (1.5 - (0.0 / (3.14159 + ((1.5 * (1.5 - 0.0)) + 0.0)))))) || true)) {
                                        switch ("\n\t" + "\n\t") {
                                        case "a":
                                            if len(v17) > 0 {
                                                v17[0] = (^(fi14))
                                            }
                                        }
                                    }
                                    _ = v24
                                } else {
                                    v1 -= v1
                                    v30 := GenS1{F0: 112}
                                    if ((false && true) || false) {
                                        switch ("a" + ((("a" + "") + "") + ((("foo bar" + "a") + (("" + "") + "a")) + ""))) {
                                        case "hello":
                                            v4 -= (v4 >> ((332) & 63))
                                        case "":
                                            v1 += v1
                                        case "a":
                                            var v31 int64 = -1
                                            _ = v31
                                        default:
                                            v32 := GenS2{F0: (v4 ^ (v4 >> ((509) & 63)))}
                                            _ = v32
                                        }
                                        var v33 int16 = 32767
                                        v4 += -1
                                        _ = v33
                                    }
                                    var v34 int32 = 1
                                    if true {
                                        var v35 uint16 = v1
                                        _ = v35
                                    } else {
                                        var v36 uint = 65535
                                        switch "foo bar" {
                                        case "foo bar":
                                            v1 += 65535
                                        case "a":
                                            switch "a" {
                                            case "foo bar":
                                                var v37 string = ("foo bar" + ("a" + "foo bar"))
                                                _ = v37
                                            }
                                        }
                                        if len(v17) > 0 {
                                            v17[0] = fi10
                                        }
                                        _ = v36
                                    }
                                    switch 1 {
                                    case 32767:
                                        if (!("" > "a")) {
                                            if false {
                                                var v38 bool = (67 <= 42)
                                                _ = v38
                                            } else {
                                                fmt.Println(fi14)
                                                if (2147483648 == 2147483648) {
                                                    v4 = v4
                                                }
                                                v1 = v1
                                            }
                                            fmt.Println(fi14, fi10, fi10)
                                        } else {
                                            switch fi14 {
                                            case 157:
                                                var v39 int32 = (v34 & (-(v34)))
                                                _ = v39
                                            case 4294967295:
                                                switch -994 {
                                                case -2147483648:
                                                    var v40 uint32 = 42
                                                    _ = v40
                                                case 1:
                                                    v41 := GenS3{F0: (true || true), F1: 128}
                                                    _ = v41
                                                case 750:
                                                    var v42 uint32 = 276
                                                    _ = v42
                                                default:
                                                    v17 = append(v17, fi14)
                                                }
                                            default:
                                                if (1 < 1) {
                                                    var v43 int64 = 259
                                                    _ = v43
                                                } else {
                                                    switch (^(fi10)) {
                                                    case 9223372036854775807:
                                                        v44 := GenS1{F0: 2}
                                                        _ = v44
                                                    case 431:
                                                        switch 32767 {
                                                        case -2147483648:
                                                            switch fi14 {
                                                            case 32767:
                                                                if true {
                                                                    if len(v17) > 0 {
                                                                        v17[0] = (fi10 << ((0) & 63))
                                                                    }
                                                                    fmt.Println(v1, v1)
                                                                    var v45 int = (fi10 & fi10)
                                                                    _ = v45
                                                                }
                                                            default:
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    fmt.Println(fi14, v4)
                                                                }
                                                            }
                                                        case 328:
                                                            v4 += 76
                                                            var v46 int8 = (v4 % ((v4) | 1))
                                                            fmt.Println(v1, fi14)
                                                            delete(v3, fi14)
                                                            _ = len(v3)
                                                            v47 := GenS3{F0: (65535 != 0), F1: 13}
                                                            _ = v46
                                                            _ = v47
                                                        case -1:
                                                            if len(v17) > 0 {
                                                                v17[0] = fi10
                                                            }
                                                        default:
                                                            v4 -= (v4 ^ v4)
                                                        }
                                                    default:
                                                        if (0 == 9223372036854775807) {
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v48 string = "a"
                                                                _ = v48
                                                            case "":
                                                                switch 127 {
                                                                case 128:
                                                                    var v49 float32 = 1.5
                                                                    var v50 uint64 = 349
                                                                    if (!(-1 > -1)) {
                                                                        if len(v17) >= 2 {
                                                                            v17 = v17[1:len(v17)]
                                                                        }
                                                                        v4 -= (v4 << ((4294967295) & 63))
                                                                    } else {
                                                                        v51 := GenS3{F0: false, F1: 0}
                                                                        _ = v51
                                                                    }
                                                                    _ = v49
                                                                    _ = v50
                                                                case 0:
                                                                    v17 = append(v17, fi10)
                                                                default:
                                                                    if (333 <= 65535) {
                                                                        var v52 float32 = ((0.1 + 1.5) + 1.5)
                                                                        var v53 int = 1
                                                                        switch int(v4) {
                                                                        case 127:
                                                                            v54 := GenS1{F0: 0}
                                                                            fmt.Println(v4)
                                                                            if (v4 < (v4 >> ((2) & 63))) {
                                                                                v1 += 1
                                                                            }
                                                                            _ = v54
                                                                        case 0:
                                                                            var v55 int16 = 740
                                                                            _ = v55
                                                                        default:
                                                                            delete(v3, (fi10 * (fi10 - fi14)))
                                                                        }
                                                                        _ = v52
                                                                        _ = v53
                                                                    } else {
                                                                        v56 := GenS1{F0: 255}
                                                                        var v57 int64 = 406
                                                                        _ = v56
                                                                        _ = v57
                                                                    }
                                                                }
                                                            case "a":
                                                                switch 9223372036854775807 {
                                                                case -190:
                                                                    var v58 uint64 = 42
                                                                    _ = v58
                                                                default:
                                                                    if false {
                                                                        var v59 int64 = 0
                                                                        if len(v17) > 0 {
                                                                            v17[0] = fi14
                                                                        }
                                                                        switch fi14 {
                                                                        case -2147483648:
                                                                            fmt.Println(v59, v59, v1)
                                                                        case -9223372036854775808:
                                                                            var v60 uint = 868
                                                                            _ = v60
                                                                        case -49:
                                                                            fmt.Println(v4, v34, fi14)
                                                                        }
                                                                        _ = v59
                                                                    }
                                                                }
                                                            default:
                                                                v1 = v1
                                                            }
                                                            if false {
                                                                var v61 bool = (1 >= 1)
                                                                if (!false) {
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        v4 += v4
                                                                    case "\n\t":
                                                                        _ = len(v17)
                                                                    case "hello":
                                                                        v62 := GenS4{F0: 0}
                                                                        _ = v62
                                                                    }
                                                                    v2[42] = 32768
                                                                    switch ("foo bar" + ("hello" + (("\n\t" + "hello") + "\n\t"))) {
                                                                    case "\n\t":
                                                                        fmt.Println(v4)
                                                                        switch (fi10 & (fi14 + (fi14 << ((430) & 63)))) {
                                                                        case 639:
                                                                            v1 = (^(v1))
                                                                        }
                                                                        var v63 int8 = (^(v4))
                                                                        _ = v63
                                                                    case "foo bar":
                                                                        _ = len(v17)
                                                                    default:
                                                                        if (!v61) {
                                                                            v1 = (v1 % ((42) | 1))
                                                                        } else {
                                                                            v3[(^(fi10))] = 65535
                                                                            var v64 int = 1
                                                                            v65 := GenS2{F0: v4}
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                    }
                                                                }
                                                                var v66 int64 = 9223372036854775807
                                                                v67 := GenS2{F0: v4}
                                                                v68 := v3[42]
                                                                if ((0.0 + ((3.14159 * float32(0.1)) - 0.1)) < -2.25) {
                                                                    fmt.Println(fi14, fi10)
                                                                    var v69 uint32 = v68
                                                                    fmt.Println(v4, v66, fi10)
                                                                    _ = v69
                                                                }
                                                                if (-32768 > 618) {
                                                                    if v61 {
                                                                        var v70 int8 = v4
                                                                        _ = v70
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        if v61 {
                                                                            v71 := GenS5{F0: (-(fi10))}
                                                                            var v72 uint32 = (v68 + v68)
                                                                            if len(v17) >= 2 {
                                                                                v17 = v17[1:len(v17)]
                                                                            }
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                    case "a":
                                                                        if ((!v61) && v61) {
                                                                            v66 = int64(v34)
                                                                            if ((v66 / (((v66 | (v66 >> ((42) & 63)))) | 1)) != (^(v66))) {
                                                                                if (42 <= v68) {
                                                                                    _ = len(v17)
                                                                                    if ((!v61) || ((2147483647 > (^(v68))) || v61)) {
                                                                                        v73 := GenS2{F0: v4}
                                                                                        v74 := GenS6{F0: (v4 << ((773) & 63))}
                                                                                        v3[2147483647] = (v68 % (((v68 >> ((326) & 63))) | 1))
                                                                                        _ = v73
                                                                                        _ = v74
                                                                                    }
                                                                                    v17 = append(v17, 0)
                                                                                } else {
                                                                                    v75 := GenS7{F0: -1}
                                                                                    v76 := v2[v68]
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                }
                                                                                if (((-32768 >= 1000) || ((v66 % (((v66 << ((1) & 63))) | 1)) == v66)) && false) {
                                                                                    v34 -= v34
                                                                                    if v61 {
                                                                                        fmt.Println(v34, v4)
                                                                                        v34 += 1
                                                                                        delete(v2, (v68 & 2))
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    fmt.Println(v34)
                                                                                    switch (fi10 << ((2) & 63)) {
                                                                                    case 42:
                                                                                        v77 := GenS8{F0: (v1 >> ((1) & 63))}
                                                                                        _ = v77
                                                                                    case -720:
                                                                                        switch ("" + "foo bar") {
                                                                                        case "foo bar":
                                                                                            if false {
                                                                                                if ((1.5 * 3.14159) >= ((1e10 / 1.0) + 0.1)) {
                                                                                                    v4 -= v4
                                                                                                    v66 += v66
                                                                                                    var v78 string = "hello"
                                                                                                    var v79 int64 = (-(v66))
                                                                                                    _ = v78
                                                                                                    _ = v79
                                                                                                } else {
                                                                                                    var v80 int64 = (v66 & v66)
                                                                                                    var v81 int64 = v66
                                                                                                    v80 -= (^(v66))
                                                                                                    _ = v80
                                                                                                    _ = v81
                                                                                                }
                                                                                                var v82 uint = 18446744073709551615
                                                                                                var v83 int64 = (v66 / ((42) | 1))
                                                                                                _ = v82
                                                                                                _ = v83
                                                                                            }
                                                                                        case "":
                                                                                            v34 = (v34 << ((4294967295) & 63))
                                                                                        default:
                                                                                            v84 := GenS9{F0: 4294967295, F1: 463, F2: v1}
                                                                                            _ = v84
                                                                                        }
                                                                                    default:
                                                                                        v3[fi14] = v68
                                                                                    }
                                                                                    switch fi10 {
                                                                                    case 128:
                                                                                        _ = len(v2)
                                                                                    case 32768:
                                                                                        v61 = v61
                                                                                    case 0:
                                                                                        var v85 float64 = 1e10
                                                                                        v34 -= v34
                                                                                        v86 := GenS10{F0: 4294967295, F1: 0}
                                                                                        _ = v85
                                                                                        _ = v86
                                                                                    }
                                                                                } else {
                                                                                    if (fi10 < fi14) {
                                                                                        v61 = ((v34 + (v34 / ((v34) | 1))) > -636)
                                                                                    } else {
                                                                                        switch "foo bar" {
                                                                                        case "hello":
                                                                                            switch "a" {
                                                                                            case "a":
                                                                                                fmt.Println(v61)
                                                                                            case "":
                                                                                                var v87 uint16 = v1
                                                                                                _ = v87
                                                                                            }
                                                                                        case "a":
                                                                                            var v88 float32 = (-2.25 + (1e10 / 1e10))
                                                                                            _ = v88
                                                                                        case "":
                                                                                            fmt.Println(fi14, v68, v68)
                                                                                        }
                                                                                        if v61 {
                                                                                            v89 := GenS5{F0: fi14}
                                                                                            v61 = false
                                                                                            if v61 {
                                                                                                v90 := v2[(v68 << ((18446744073709551615) & 63))]
                                                                                                if (v66 < v66) {
                                                                                                    v68 = 0
                                                                                                } else {
                                                                                                    v2[0] = (v90 / (((fi10 | v90)) | 1))
                                                                                                    fmt.Println(v68, fi10, v66)
                                                                                                }
                                                                                                _ = v90
                                                                                            }
                                                                                            _ = v89
                                                                                        } else {
                                                                                            v91 := v3[fi10]
                                                                                            v92 := GenS11{F0: 0, F1: v34}
                                                                                            v61 = true
                                                                                            if ((v34 - v34) <= v34) {
                                                                                                delete(v2, 0)
                                                                                            }
                                                                                            if v61 {
                                                                                                switch fi10 {
                                                                                                case 77:
                                                                                                    _ = len(v3)
                                                                                                case 283:
                                                                                                    v93 := GenS12{F0: (int32(v4) >= (v34 ^ v34)), F1: (v4 << ((18446744073709551615) & 63)), F2: (fi14 * fi14)}
                                                                                                    _ = v93
                                                                                                case -914:
                                                                                                    if true {
                                                                                                        v4 = (^(v4))
                                                                                                        fmt.Println(v68)
                                                                                                        switch fi10 {
                                                                                                        case -9223372036854775808:
                                                                                                            v66 = -9223372036854775808
                                                                                                        case 1:
                                                                                                            v94 := GenS13{F0: v4}
                                                                                                            _ = v94
                                                                                                        case 946:
                                                                                                            v66 = -9223372036854775808
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    if len(v17) > 0 {
                                                                                                        v17[0] = (fi14 | 2147483647)
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                switch fi14 {
                                                                                                case 42:
                                                                                                    v34 = (v34 >> ((18446744073709551615) & 63))
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "":
                                                                                                    if len(v17) >= 2 {
                                                                                                        v17 = v17[1:len(v17)]
                                                                                                    }
                                                                                                case "a":
                                                                                                    fmt.Println(v68, v68, v34)
                                                                                                case "foo bar":
                                                                                                    v1 += v1
                                                                                                }
                                                                                                v17 = append(v17, -820)
                                                                                            }
                                                                                            _ = v91
                                                                                            _ = v92
                                                                                        }
                                                                                        switch fi10 {
                                                                                        case 4294967295:
                                                                                            v66 = (v66 >> ((42) & 63))
                                                                                        case -2147483648:
                                                                                            if (385 <= 2) {
                                                                                                v95 := GenS9{F0: 2, F1: 0, F2: (v1 % ((137) | 1))}
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v34, fi10)
                                                                                                case "\n\t":
                                                                                                    if true {
                                                                                                        var v96 float32 = 1.5
                                                                                                        var v97 uint16 = 1
                                                                                                        var v98 int16 = -940
                                                                                                        var v99 bool = (v66 != v66)
                                                                                                        _ = v96
                                                                                                        _ = v97
                                                                                                        _ = v98
                                                                                                        _ = v99
                                                                                                    } else {
                                                                                                        delete(v2, (^(v68)))
                                                                                                        v1 += v1
                                                                                                        var v100 int8 = v4
                                                                                                        v101 := GenS14{F0: (v68 * (^(v68)))}
                                                                                                        _ = v100
                                                                                                        _ = v101
                                                                                                    }
                                                                                                }
                                                                                                v102 := GenS7{F0: 32767}
                                                                                                _ = v102
                                                                                                _ = v95
                                                                                            }
                                                                                        default:
                                                                                            v2[(v68 | (v68 << ((939) & 63)))] = fi10
                                                                                            v17 = append(v17, 0)
                                                                                            v103 := GenS12{F0: (1.5 <= 3.14159), F1: (^(v4)), F2: fi10}
                                                                                            _ = v103
                                                                                        }
                                                                                    }
                                                                                    v17 = append(v17, (fi14 / ((798) | 1)))
                                                                                }
                                                                                var v104 int64 = v66
                                                                                switch (fi14 >> ((896) & 63)) {
                                                                                case -9223372036854775808:
                                                                                    if ((-9223372036854775808 >= fi14) || (0 >= v104)) {
                                                                                        var v105 uint32 = (^(v68))
                                                                                        if len(v17) > 0 {
                                                                                            v17[0] = int(v68)
                                                                                        }
                                                                                        _ = v105
                                                                                    } else {
                                                                                        if true {
                                                                                            var v106 uint8 = 42
                                                                                            v107 := GenS2{F0: (^(v4))}
                                                                                            var v108 uint32 = uint32(v1)
                                                                                            _ = v106
                                                                                            _ = v107
                                                                                            _ = v108
                                                                                        } else {
                                                                                            var v109 uint = 1
                                                                                            _ = v109
                                                                                        }
                                                                                    }
                                                                                    if true {
                                                                                        var v110 uint16 = 65535
                                                                                        _ = v110
                                                                                    }
                                                                                    v61 = (("" + "a") == "")
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        if false {
                                                                                            var v111 uint64 = 123
                                                                                            v1 = v1
                                                                                            var v112 int8 = v4
                                                                                            _ = v111
                                                                                            _ = v112
                                                                                        } else {
                                                                                            if (!(!(25 <= 0))) {
                                                                                                v113 := GenS13{F0: (v4 - (-(v4)))}
                                                                                                switch "hello" {
                                                                                                case "foo bar":
                                                                                                    var v114 bool = v61
                                                                                                    _ = v114
                                                                                                case "\n\t":
                                                                                                    var v115 float32 = -2.25
                                                                                                    var v116 uint16 = v1
                                                                                                    if true {
                                                                                                        v104 = 42
                                                                                                        v117 := GenS2{F0: v4}
                                                                                                        fmt.Println(v116, v115)
                                                                                                        _ = v117
                                                                                                    }
                                                                                                    _ = v115
                                                                                                    _ = v116
                                                                                                default:
                                                                                                    _ = len(v17)
                                                                                                }
                                                                                                var v118 int16 = -563
                                                                                                _ = v113
                                                                                                _ = v118
                                                                                            } else {
                                                                                                switch "foo bar" {
                                                                                                case "hello":
                                                                                                    switch (("a" + "") + "hello") {
                                                                                                    case "\n\t":
                                                                                                        if v61 {
                                                                                                            v17 = append(v17, (fi10 & fi14))
                                                                                                            v119 := v3[(fi10 + (^(fi14)))]
                                                                                                            fmt.Println(fi10, v119)
                                                                                                            _ = v119
                                                                                                        }
                                                                                                    default:
                                                                                                        v61 = (false || (!(1 <= 116)))
                                                                                                        var v120 uint16 = (v1 % (((v1 / ((1) | 1))) | 1))
                                                                                                        if (2 > (v68 + v68)) {
                                                                                                            v121 := GenS13{F0: (^(v4))}
                                                                                                            v122 := GenS1{F0: 2}
                                                                                                            if (v61 || (4294967295 <= uint(v120))) {
                                                                                                                var v123 int = fi14
                                                                                                                if false {
                                                                                                                    v66 += (v66 >> ((0) & 63))
                                                                                                                    switch "\n\t" {
                                                                                                                    case "\n\t":
                                                                                                                        fmt.Println(v66, fi14)
                                                                                                                    case "a":
                                                                                                                        v124 := GenS4{F0: int64(v120)}
                                                                                                                        _ = v124
                                                                                                                    default:
                                                                                                                        if (0 >= 65535) {
                                                                                                                            v125 := GenS15{F0: v4, F1: "a", F2: 42}
                                                                                                                            switch "\n\t" {
                                                                                                                            case "":
                                                                                                                                _ = len(v3)
                                                                                                                            }
                                                                                                                            var v126 int32 = v34
                                                                                                                            v126 = -1
                                                                                                                            if (-16 > (^(v4))) {
                                                                                                                                fmt.Println(v68)
                                                                                                                                if true {
                                                                                                                                    fmt.Println(v68, v66)
                                                                                                                                    fmt.Println(v66, fi14, v34)
                                                                                                                                    v123 -= v123
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                delete(v2, v68)
                                                                                                                            }
                                                                                                                            var v127 uint64 = uint64(v1)
                                                                                                                            _ = v125
                                                                                                                            _ = v126
                                                                                                                            _ = v127
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch fi10 {
                                                                                                                    case 128:
                                                                                                                        var v128 int = 215
                                                                                                                        _ = v128
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    switch -9223372036854775808 {
                                                                                                                    case -2147483648:
                                                                                                                        if (537 > v34) {
                                                                                                                            var v129 int64 = v104
                                                                                                                            _ = v129
                                                                                                                        }
                                                                                                                    case 127:
                                                                                                                        fmt.Println(v104, fi10, v34)
                                                                                                                    case 627:
                                                                                                                        var v130 uint8 = 0
                                                                                                                        _ = v130
                                                                                                                    default:
                                                                                                                        v1 = 1
                                                                                                                        v131 := GenS16{F0: v34}
                                                                                                                        v61 = v61
                                                                                                                        _ = v131
                                                                                                                    }
                                                                                                                    v132 := GenS17{F0: 4294967295, F1: ("a" + (("a" + "a") + "foo bar"))}
                                                                                                                    if (240 >= 127) {
                                                                                                                        var v133 float64 = 0.0
                                                                                                                        _ = v133
                                                                                                                    } else {
                                                                                                                        fmt.Println(v61)
                                                                                                                        var v134 int16 = -835
                                                                                                                        var v135 uint64 = 4294967295
                                                                                                                        fmt.Println(v134, v123, v120)
                                                                                                                        _ = v134
                                                                                                                        _ = v135
                                                                                                                    }
                                                                                                                    v136 := GenS18{F0: (v61 || false), F1: 42}
                                                                                                                    switch (fi10 * int(v68)) {
                                                                                                                    case 127:
                                                                                                                        v137 := GenS13{F0: (v4 % ((-100) | 1))}
                                                                                                                        _ = v137
                                                                                                                    case -885:
                                                                                                                        fmt.Println(v34, v123)
                                                                                                                    case 4294967295:
                                                                                                                        switch ("hello" + "\n\t") {
                                                                                                                        case "a":
                                                                                                                            var v138 uint32 = v68
                                                                                                                            _ = v138
                                                                                                                        case "\n\t":
                                                                                                                            fmt.Println(v68, v68, v120)
                                                                                                                        case "hello":
                                                                                                                            _ = len(v3)
                                                                                                                            v139 := GenS19{F0: (v123 / (((fi10 & fi10)) | 1)), F1: int8(v104)}
                                                                                                                            _ = v139
                                                                                                                        default:
                                                                                                                            v17 = append(v17, 0)
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        if v61 {
                                                                                                                            var v140 int8 = v4
                                                                                                                            _ = v140
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v132
                                                                                                                    _ = v136
                                                                                                                }
                                                                                                                v141 := v2[v68]
                                                                                                                v142 := GenS19{F0: 0, F1: 42}
                                                                                                                switch "a" {
                                                                                                                case "hello":
                                                                                                                    if (float32((0.0 * ((float64((3.14159 - ((1e10 / 1.0) + float32(0.1)))) + -2.25) + float64((float32(0.1) - 0.0))))) != (0.0 / 0.1)) {
                                                                                                                        if v61 {
                                                                                                                            if v61 {
                                                                                                                                var v143 uint64 = uint64(v68)
                                                                                                                                _ = v143
                                                                                                                            }
                                                                                                                            delete(v2, (^(v68)))
                                                                                                                            v66 -= (v66 - -9223372036854775808)
                                                                                                                        } else {
                                                                                                                            if v61 {
                                                                                                                                if len(v17) > 0 {
                                                                                                                                    v17[0] = 4294967295
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v144 := GenS20{F0: (-(v141))}
                                                                                                                            _ = v144
                                                                                                                        }
                                                                                                                        if v61 {
                                                                                                                            var v145 int16 = -43
                                                                                                                            var v146 float64 = 1e10
                                                                                                                            _ = v145
                                                                                                                            _ = v146
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        if ((float64(1.5) - float64((-2.25 * 0.0))) != -2.25) {
                                                                                                                            v147 := GenS21{F0: v34, F1: (v141 >> ((18446744073709551615) & 63))}
                                                                                                                            var v148 uint32 = (v68 - 0)
                                                                                                                            var v149 uint16 = (v120 >> ((2) & 63))
                                                                                                                            _ = v147
                                                                                                                            _ = v148
                                                                                                                            _ = v149
                                                                                                                        } else {
                                                                                                                            switch "foo bar" {
                                                                                                                            case "a":
                                                                                                                                v120 = v120
                                                                                                                            case "hello":
                                                                                                                                switch v123 {
                                                                                                                                case 32767:
                                                                                                                                    if v61 {
                                                                                                                                        if len(v17) > 0 {
                                                                                                                                            v17[0] = v141
                                                                                                                                        }
                                                                                                                                        if len(v17) >= 2 {
                                                                                                                                            v17 = v17[1:len(v17)]
                                                                                                                                        }
                                                                                                                                        if v61 {
                                                                                                                                            v68 += v68
                                                                                                                                            var v150 int16 = 32767
                                                                                                                                            _ = v150
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        v151 := GenS16{F0: v34}
                                                                                                                                        _ = v151
                                                                                                                                    }
                                                                                                                                case 42:
                                                                                                                                    v152 := v3[v123]
                                                                                                                                    _ = v152
                                                                                                                                case 0:
                                                                                                                                    if ((42 != v4) && ((v104 * (v66 % ((814) | 1))) >= v104)) {
                                                                                                                                        switch ("foo bar" + "a") {
                                                                                                                                        case "a":
                                                                                                                                            if true {
                                                                                                                                                var v153 float64 = float64(1e10)
                                                                                                                                                _ = v153
                                                                                                                                            }
                                                                                                                                        case "\n\t":
                                                                                                                                            var v154 int = 9223372036854775807
                                                                                                                                            _ = v154
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        v155 := v3[(-(fi10))]
                                                                                                                                        var v156 bool = ((v4 | 42) < (v4 + 96))
                                                                                                                                        var v157 int64 = (^(v104))
                                                                                                                                        _ = v155
                                                                                                                                        _ = v156
                                                                                                                                        _ = v157
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                if v61 {
                                                                                                                                    if (!v61) {
                                                                                                                                        v123 += fi14
                                                                                                                                    } else {
                                                                                                                                        var v158 uint = 665
                                                                                                                                        _ = v158
                                                                                                                                    }
                                                                                                                                    if v61 {
                                                                                                                                        fmt.Println(v123)
                                                                                                                                    } else {
                                                                                                                                        if ((fi10 * (fi14 % (((fi14 >> ((65535) & 63))) | 1))) > fi10) {
                                                                                                                                            var v159 int64 = v104
                                                                                                                                            _ = v159
                                                                                                                                        } else {
                                                                                                                                            switch fi10 {
                                                                                                                                            case -2147483648:
                                                                                                                                                switch "" {
                                                                                                                                                case "hello":
                                                                                                                                                    if ((^(v68)) <= 821) {
                                                                                                                                                        v3[fi14] = 2147483647
                                                                                                                                                        var v160 int32 = v34
                                                                                                                                                        var v161 int8 = v4
                                                                                                                                                        _ = v160
                                                                                                                                                        _ = v161
                                                                                                                                                    } else {
                                                                                                                                                        if v61 {
                                                                                                                                                            v162 := GenS22{F0: 128}
                                                                                                                                                            if (42 <= 865) {
                                                                                                                                                                delete(v2, (v68 * 4294967295))
                                                                                                                                                            } else {
                                                                                                                                                                if v61 {
                                                                                                                                                                    v66 = v104
                                                                                                                                                                    switch "a" {
                                                                                                                                                                    case "":
                                                                                                                                                                        v141 -= fi14
                                                                                                                                                                    case "a":
                                                                                                                                                                        switch 4294967295 {
                                                                                                                                                                        case 4294967295:
                                                                                                                                                                            var v163 int16 = -1
                                                                                                                                                                            _ = v163
                                                                                                                                                                        case -9223372036854775808:
                                                                                                                                                                            v66 -= 0
                                                                                                                                                                        case 0:
                                                                                                                                                                            var v164 uint32 = v68
                                                                                                                                                                            _ = v164
                                                                                                                                                                        }
                                                                                                                                                                    default:
                                                                                                                                                                        _ = len(v17)
                                                                                                                                                                    }
                                                                                                                                                                    if ((v61 && (3.14159 == 3.14159)) || false) {
                                                                                                                                                                        v4 += v4
                                                                                                                                                                    }
                                                                                                                                                                } else {
                                                                                                                                                                    if v61 {
                                                                                                                                                                        fmt.Println(v4, v141, v61)
                                                                                                                                                                        v61 = ((!v61) && (!true))
                                                                                                                                                                        switch -18 {
                                                                                                                                                                        case -1:
                                                                                                                                                                            var v165 bool = false
                                                                                                                                                                            _ = v165
                                                                                                                                                                        case -2147483648:
                                                                                                                                                                            var v166 string = "hello"
                                                                                                                                                                            _ = v166
                                                                                                                                                                        case 42:
                                                                                                                                                                            v141 -= -937
                                                                                                                                                                            if true {
                                                                                                                                                                                if ((^(v68)) < (^(v68))) {
                                                                                                                                                                                    switch ((("foo bar" + "hello") + "a") + "\n\t") {
                                                                                                                                                                                    case "\n\t":
                                                                                                                                                                                        var v167 int8 = 107
                                                                                                                                                                                        _ = v167
                                                                                                                                                                                    case "a":
                                                                                                                                                                                        v120 = 42
                                                                                                                                                                                    }
                                                                                                                                                                                } else {
                                                                                                                                                                                    v168 := GenS23{F0: uint(v120), F1: (-(v104))}
                                                                                                                                                                                    _ = v168
                                                                                                                                                                                }
                                                                                                                                                                                var v169 uint16 = v120
                                                                                                                                                                                _ = v169
                                                                                                                                                                            }
                                                                                                                                                                            var v170 int16 = int16(fi10)
                                                                                                                                                                            _ = v170
                                                                                                                                                                        }
                                                                                                                                                                    } else {
                                                                                                                                                                        v34 += (v34 / ((v34) | 1))
                                                                                                                                                                        fmt.Println(fi10)
                                                                                                                                                                    }
                                                                                                                                                                    v61 = (-753 < int32(v4))
                                                                                                                                                                }
                                                                                                                                                                switch "a" {
                                                                                                                                                                case "a":
                                                                                                                                                                    var v171 int64 = (v104 ^ v66)
                                                                                                                                                                    _ = v171
                                                                                                                                                                case "foo bar":
                                                                                                                                                                    v4 -= (v4 ^ -1)
                                                                                                                                                                default:
                                                                                                                                                                    if v61 {
                                                                                                                                                                        if v61 {
                                                                                                                                                                            if (false || false) {
                                                                                                                                                                                v66 = (v104 ^ (v66 - (v66 >> ((42) & 63))))
                                                                                                                                                                            } else {
                                                                                                                                                                                if v61 {
                                                                                                                                                                                    v141 = fi10
                                                                                                                                                                                } else {
                                                                                                                                                                                    var v172 float64 = 0.1
                                                                                                                                                                                    switch (v141 ^ fi10) {
                                                                                                                                                                                    case -2147483648:
                                                                                                                                                                                        switch fi10 {
                                                                                                                                                                                        case 42:
                                                                                                                                                                                            var v173 int16 = -1
                                                                                                                                                                                            _ = v173
                                                                                                                                                                                        default:
                                                                                                                                                                                            var v174 int8 = v4
                                                                                                                                                                                            _ = v174
                                                                                                                                                                                        }
                                                                                                                                                                                    case 127:
                                                                                                                                                                                        if v61 {
                                                                                                                                                                                            var v175 float64 = ((float64(1.5) / (v172 + (float64(1.5) / v172))) / (v172 - (0.0 - 3.14159)))
                                                                                                                                                                                            _ = v175
                                                                                                                                                                                        }
                                                                                                                                                                                    case 844:
                                                                                                                                                                                        v123 += (fi14 - v141)
                                                                                                                                                                                        if len(v17) >= 2 {
                                                                                                                                                                                            v17 = v17[1:len(v17)]
                                                                                                                                                                                        }
                                                                                                                                                                                        var v176 float32 = (0.1 / 3.14159)
                                                                                                                                                                                        if (v176 >= float32(v172)) {
                                                                                                                                                                                            var v177 bool = v61
                                                                                                                                                                                            v178 := GenS24{F0: (v4 | (v4 + (v4 % ((v4) | 1)))), F1: 2}
                                                                                                                                                                                            v4 += 91
                                                                                                                                                                                            _ = v177
                                                                                                                                                                                            _ = v178
                                                                                                                                                                                        }
                                                                                                                                                                                        _ = v176
                                                                                                                                                                                    }
                                                                                                                                                                                    fmt.Println(v4)
                                                                                                                                                                                    _ = v172
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                            fmt.Println(v120)
                                                                                                                                                                            v179 := GenS25{F0: 42}
                                                                                                                                                                            _ = v179
                                                                                                                                                                        } else {
                                                                                                                                                                            v180 := GenS26{F0: -128}
                                                                                                                                                                            var v181 int32 = v34
                                                                                                                                                                            var v182 int64 = -1
                                                                                                                                                                            _ = v180
                                                                                                                                                                            _ = v181
                                                                                                                                                                            _ = v182
                                                                                                                                                                        }
                                                                                                                                                                    } else {
                                                                                                                                                                        var v183 float32 = 1e10
                                                                                                                                                                        fmt.Println(v34, v123)
                                                                                                                                                                        if (3.14159 != -2.25) {
                                                                                                                                                                            var v184 float64 = 0.1
                                                                                                                                                                            _ = len(v3)
                                                                                                                                                                            _ = v184
                                                                                                                                                                        }
                                                                                                                                                                        _ = v183
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                                switch "\n\t" {
                                                                                                                                                                case "hello":
                                                                                                                                                                    v3[(v123 % ((v141) | 1))] = v68
                                                                                                                                                                    v185 := GenS19{F0: 32768, F1: (v4 % ((v4) | 1))}
                                                                                                                                                                    fmt.Println(v4)
                                                                                                                                                                    switch "hello" {
                                                                                                                                                                    case "hello":
                                                                                                                                                                        v66 += (v104 - 9223372036854775807)
                                                                                                                                                                    case "foo bar":
                                                                                                                                                                        v1 = 152
                                                                                                                                                                    case "":
                                                                                                                                                                        fmt.Println(v104, v141)
                                                                                                                                                                    default:
                                                                                                                                                                        v186 := GenS14{F0: v68}
                                                                                                                                                                        _ = v186
                                                                                                                                                                    }
                                                                                                                                                                    if v61 {
                                                                                                                                                                        v120 += v1
                                                                                                                                                                        if (!v61) {
                                                                                                                                                                            if false {
                                                                                                                                                                                v187 := GenS14{F0: (^(v68))}
                                                                                                                                                                                _ = v187
                                                                                                                                                                            } else {
                                                                                                                                                                                var v188 uint32 = (v68 << ((2) & 63))
                                                                                                                                                                                v61 = false
                                                                                                                                                                                v189 := GenS22{F0: 128}
                                                                                                                                                                                _ = v188
                                                                                                                                                                                _ = v189
                                                                                                                                                                            }
                                                                                                                                                                            v3[(fi10 % (((^(fi14))) | 1))] = 267
                                                                                                                                                                        } else {
                                                                                                                                                                            if (32767 != 1000) {
                                                                                                                                                                                if true {
                                                                                                                                                                                    v123 = (fi14 + v141)
                                                                                                                                                                                    var v190 int16 = 1
                                                                                                                                                                                    _ = v190
                                                                                                                                                                                }
                                                                                                                                                                            } else {
                                                                                                                                                                                var v191 int8 = (v4 | (-(v4)))
                                                                                                                                                                                v4 = v191
                                                                                                                                                                                fmt.Println(v104, v1)
                                                                                                                                                                                _ = v191
                                                                                                                                                                            }
                                                                                                                                                                            var v192 int8 = v4
                                                                                                                                                                            _ = v192
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    _ = v185
                                                                                                                                                                case "\n\t":
                                                                                                                                                                    if true {
                                                                                                                                                                        var v193 int8 = (^(v4))
                                                                                                                                                                        _ = v193
                                                                                                                                                                    }
                                                                                                                                                                default:
                                                                                                                                                                    if (true && ((v120 - uint16(v68)) != v1)) {
                                                                                                                                                                        v194 := GenS27{F0: v4, F1: 222, F2: (v104 << ((65535) & 63))}
                                                                                                                                                                        _ = v194
                                                                                                                                                                    }
                                                                                                                                                                    v123 = (v141 & (^(v141)))
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            fmt.Println(fi10)
                                                                                                                                                            _ = v162
                                                                                                                                                        } else {
                                                                                                                                                            if (!(!true)) {
                                                                                                                                                                var v195 float32 = -2.25
                                                                                                                                                                fmt.Println(v141, v61, fi10)
                                                                                                                                                                _ = v195
                                                                                                                                                            }
                                                                                                                                                            v61 = ((true && (1.5 != float64((1e10 + (0.1 * (((float32((-2.25 * 1.5)) / 1.5) * 1.5) - 0.0)))))) || v61)
                                                                                                                                                            var v196 bool = false
                                                                                                                                                            _ = v196
                                                                                                                                                        }
                                                                                                                                                        if (!(!true)) {
                                                                                                                                                            v68 -= v68
                                                                                                                                                            v68 = v68
                                                                                                                                                            if v61 {
                                                                                                                                                                switch "hello" {
                                                                                                                                                                case "hello":
                                                                                                                                                                    v123 += (fi10 ^ 352)
                                                                                                                                                                case "a":
                                                                                                                                                                    switch (^(v141)) {
                                                                                                                                                                    case 32767:
                                                                                                                                                                        if (v61 || v61) {
                                                                                                                                                                            switch "a" {
                                                                                                                                                                            case "hello":
                                                                                                                                                                                delete(v3, (fi10 | 4294967295))
                                                                                                                                                                            case "":
                                                                                                                                                                                var v197 int16 = int16(v34)
                                                                                                                                                                                _ = v197
                                                                                                                                                                            default:
                                                                                                                                                                                v141 += 127
                                                                                                                                                                            }
                                                                                                                                                                            v66 -= 9223372036854775807
                                                                                                                                                                        } else {
                                                                                                                                                                            if len(v17) > 0 {
                                                                                                                                                                                v17[0] = (fi14 - (fi14 - v123))
                                                                                                                                                                            }
                                                                                                                                                                            v198 := v3[127]
                                                                                                                                                                            if (42 <= v4) {
                                                                                                                                                                                v2[v198] = (v141 << ((987) & 63))
                                                                                                                                                                                v199 := GenS20{F0: -2147483648}
                                                                                                                                                                                switch (("hello" + "") + "\n\t") {
                                                                                                                                                                                case "a":
                                                                                                                                                                                    fmt.Println(v61)
                                                                                                                                                                                    fmt.Println(v104, fi10)
                                                                                                                                                                                default:
                                                                                                                                                                                    var v200 string = "\n\t"
                                                                                                                                                                                    _ = v200
                                                                                                                                                                                }
                                                                                                                                                                                switch (fi10 ^ 0) {
                                                                                                                                                                                case 232:
                                                                                                                                                                                    fmt.Println(v104, v4, v1)
                                                                                                                                                                                case 127:
                                                                                                                                                                                    _ = len(v17)
                                                                                                                                                                                }
                                                                                                                                                                                fmt.Println(v1, v61)
                                                                                                                                                                                v201 := GenS1{F0: 14}
                                                                                                                                                                                _ = v199
                                                                                                                                                                                _ = v201
                                                                                                                                                                            }
                                                                                                                                                                            _ = v198
                                                                                                                                                                        }
                                                                                                                                                                    case -450:
                                                                                                                                                                        v202 := v2[(v68 - v68)]
                                                                                                                                                                        _ = v202
                                                                                                                                                                    default:
                                                                                                                                                                        switch v141 {
                                                                                                                                                                        case -9223372036854775808:
                                                                                                                                                                            switch v141 {
                                                                                                                                                                            case 9223372036854775807:
                                                                                                                                                                                v203 := GenS28{F0: 413, F1: v4}
                                                                                                                                                                                _ = v203
                                                                                                                                                                            default:
                                                                                                                                                                                fmt.Println(fi14, v104)
                                                                                                                                                                            }
                                                                                                                                                                        default:
                                                                                                                                                                            v66 = (v66 | v104)
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                case "foo bar":
                                                                                                                                                                    fmt.Println(fi14, v66, v120)
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            var v204 uint32 = 4294967295
                                                                                                                                                            if v61 {
                                                                                                                                                                if v61 {
                                                                                                                                                                    var v205 int64 = v66
                                                                                                                                                                    _ = v205
                                                                                                                                                                }
                                                                                                                                                                fmt.Println(v61, v61, v34)
                                                                                                                                                            } else {
                                                                                                                                                                v204 += 65535
                                                                                                                                                            }
                                                                                                                                                            _ = v204
                                                                                                                                                        }
                                                                                                                                                        var v206 uint16 = 42
                                                                                                                                                        if (3.14159 <= 1e10) {
                                                                                                                                                            v61 = v61
                                                                                                                                                            var v207 bool = (0 > 255)
                                                                                                                                                            var v208 uint8 = 128
                                                                                                                                                            switch int(v1) {
                                                                                                                                                            case 0:
                                                                                                                                                                _ = len(v17)
                                                                                                                                                                var v209 uint32 = v68
                                                                                                                                                                _ = v209
                                                                                                                                                            case 1:
                                                                                                                                                                v210 := v2[2147483647]
                                                                                                                                                                _ = v210
                                                                                                                                                            default:
                                                                                                                                                                v211 := GenS8{F0: v206}
                                                                                                                                                                var v212 int16 = int16(v4)
                                                                                                                                                                var v213 bool = true
                                                                                                                                                                _ = v211
                                                                                                                                                                _ = v212
                                                                                                                                                                _ = v213
                                                                                                                                                            }
                                                                                                                                                            switch ("\n\t" + (("" + "") + (("a" + "foo bar") + "a"))) {
                                                                                                                                                            case "":
                                                                                                                                                                var v214 int64 = v104
                                                                                                                                                                _ = v214
                                                                                                                                                            case "foo bar":
                                                                                                                                                                var v215 uint16 = (v1 | 911)
                                                                                                                                                                _ = v215
                                                                                                                                                            case "a":
                                                                                                                                                                var v216 uint16 = (v1 << ((255) & 63))
                                                                                                                                                                _ = v216
                                                                                                                                                            }
                                                                                                                                                            _ = v207
                                                                                                                                                            _ = v208
                                                                                                                                                        }
                                                                                                                                                        if len(v17) >= 2 {
                                                                                                                                                            v17 = v17[1:len(v17)]
                                                                                                                                                        }
                                                                                                                                                        v120 = 32767
                                                                                                                                                        _ = v206
                                                                                                                                                    }
                                                                                                                                                case "foo bar":
                                                                                                                                                    fmt.Println(v123, v104)
                                                                                                                                                default:
                                                                                                                                                    fmt.Println(v141, v61)
                                                                                                                                                }
                                                                                                                                            case -1:
                                                                                                                                                fmt.Println(v104, fi10, v61)
                                                                                                                                            }
                                                                                                                                            if true {
                                                                                                                                                switch v123 {
                                                                                                                                                case -9223372036854775808:
                                                                                                                                                    if ((((v141 & fi14) == fi10) || v61) || v61) {
                                                                                                                                                        v217 := GenS11{F0: -46, F1: (v34 * (v34 >> ((382) & 63)))}
                                                                                                                                                        switch ("\n\t" + ("a" + "\n\t")) {
                                                                                                                                                        case "a":
                                                                                                                                                            v218 := GenS9{F0: 4294967295, F1: 18446744073709551615, F2: (v1 * (^(v1)))}
                                                                                                                                                            _ = v218
                                                                                                                                                        default:
                                                                                                                                                            if false {
                                                                                                                                                                var v219 int16 = 994
                                                                                                                                                                if (false && (!false)) {
                                                                                                                                                                    v219 -= (v219 << ((42) & 63))
                                                                                                                                                                }
                                                                                                                                                                _ = len(v17)
                                                                                                                                                                if true {
                                                                                                                                                                    var v220 float64 = -2.25
                                                                                                                                                                    fmt.Println(v4)
                                                                                                                                                                    v66 = v66
                                                                                                                                                                    switch "foo bar" {
                                                                                                                                                                    case "foo bar":
                                                                                                                                                                        v221 := GenS5{F0: (^(fi10))}
                                                                                                                                                                        _ = v221
                                                                                                                                                                    case "a":
                                                                                                                                                                        var v222 int32 = (v34 + -2147483648)
                                                                                                                                                                        _ = v222
                                                                                                                                                                    case "hello":
                                                                                                                                                                        if v61 {
                                                                                                                                                                            if (!true) {
                                                                                                                                                                                v120 = (v120 + v120)
                                                                                                                                                                                fmt.Println(v123, fi10, v141)
                                                                                                                                                                                var v223 uint8 = 255
                                                                                                                                                                                var v224 int8 = (v4 - int8(v141))
                                                                                                                                                                                _ = v223
                                                                                                                                                                                _ = v224
                                                                                                                                                                            } else {
                                                                                                                                                                                v104 += v104
                                                                                                                                                                            }
                                                                                                                                                                            v2[v68] = int(v1)
                                                                                                                                                                            if v61 {
                                                                                                                                                                                v141 = (fi10 & (-(v123)))
                                                                                                                                                                            } else {
                                                                                                                                                                                var v225 bool = v61
                                                                                                                                                                                var v226 int = (^(v123))
                                                                                                                                                                                _ = v225
                                                                                                                                                                                _ = v226
                                                                                                                                                                            }
                                                                                                                                                                            var v227 uint = 866
                                                                                                                                                                            _ = v227
                                                                                                                                                                        } else {
                                                                                                                                                                            v2[(^(v68))] = int(v104)
                                                                                                                                                                            fmt.Println(v61)
                                                                                                                                                                            v68 = v68
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    _ = v220
                                                                                                                                                                }
                                                                                                                                                                var v228 int8 = v4
                                                                                                                                                                _ = v219
                                                                                                                                                                _ = v228
                                                                                                                                                            } else {
                                                                                                                                                                _ = len(v17)
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        fmt.Println(v123, fi10, v34)
                                                                                                                                                        _ = v217
                                                                                                                                                    } else {
                                                                                                                                                        v229 := GenS9{F0: 998, F1: 1, F2: v120}
                                                                                                                                                        _ = v229
                                                                                                                                                    }
                                                                                                                                                default:
                                                                                                                                                    v230 := GenS16{F0: int32(v141)}
                                                                                                                                                    _ = v230
                                                                                                                                                }
                                                                                                                                                var v231 string = "foo bar"
                                                                                                                                                v232 := GenS29{F0: (v4 - v4), F1: 65535}
                                                                                                                                                _ = v231
                                                                                                                                                _ = v232
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    if (42 != (v104 % (((v104 | (v104 - (v104 << ((255) & 63))))) | 1))) {
                                                                                                                                        if (-2.25 > float32((1e10 - 1e10))) {
                                                                                                                                            v233 := GenS17{F0: 279, F1: "hello"}
                                                                                                                                            _ = v233
                                                                                                                                        } else {
                                                                                                                                            switch -29 {
                                                                                                                                            case 10:
                                                                                                                                                v234 := GenS30{F0: 119, F1: 18446744073709551615}
                                                                                                                                                _ = v234
                                                                                                                                            case 4294967295:
                                                                                                                                                v235 := v2[(v68 % (((v68 | v68)) | 1))]
                                                                                                                                                _ = v235
                                                                                                                                            }
                                                                                                                                            v236 := GenS31{F0: 160, F1: (36 > 128), F2: 42}
                                                                                                                                            _ = v236
                                                                                                                                        }
                                                                                                                                        fmt.Println(v68, v1)
                                                                                                                                        v61 = (v61 || (v34 == (v34 - (v34 / (((v34 << ((88) & 63))) | 1)))))
                                                                                                                                    } else {
                                                                                                                                        v237 := v2[v68]
                                                                                                                                        switch (fi14 << ((255) & 63)) {
                                                                                                                                        case 9223372036854775807:
                                                                                                                                            fmt.Println(v123, v4)
                                                                                                                                        case 0:
                                                                                                                                            switch (v123 * 127) {
                                                                                                                                            case 0:
                                                                                                                                                v104 = (v104 / (((v66 | v66)) | 1))
                                                                                                                                            case -249:
                                                                                                                                                var v238 uint = 46
                                                                                                                                                _ = v238
                                                                                                                                            case 2147483647:
                                                                                                                                                var v239 string = ""
                                                                                                                                                _ = v239
                                                                                                                                            }
                                                                                                                                        case -9223372036854775808:
                                                                                                                                            if (v61 || (!(0 > v34))) {
                                                                                                                                                var v240 bool = v61
                                                                                                                                                v237 -= fi14
                                                                                                                                                fmt.Println(v1)
                                                                                                                                                _ = v240
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        switch (-(v237)) {
                                                                                                                                        case 127:
                                                                                                                                            var v241 int = v237
                                                                                                                                            _ = v241
                                                                                                                                        }
                                                                                                                                        _ = v237
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v242 := GenS23{F0: 2, F1: v104}
                                                                                                                            fmt.Println(fi14)
                                                                                                                            _ = v242
                                                                                                                        }
                                                                                                                        v243 := GenS11{F0: 32767, F1: (v34 & (v34 % ((v34) | 1)))}
                                                                                                                        _ = len(v17)
                                                                                                                        v61 = v61
                                                                                                                        switch v141 {
                                                                                                                        case 4294967295:
                                                                                                                            v244 := GenS7{F0: 32767}
                                                                                                                            _ = v244
                                                                                                                        case 42:
                                                                                                                            var v245 uint64 = 42
                                                                                                                            _ = v245
                                                                                                                        default:
                                                                                                                            v1 = (v120 / ((v120) | 1))
                                                                                                                        }
                                                                                                                        _ = v243
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v123
                                                                                                                _ = v141
                                                                                                                _ = v142
                                                                                                            } else {
                                                                                                                var v246 bool = v61
                                                                                                                v61 = (v61 || v246)
                                                                                                                _ = len(v17)
                                                                                                                _ = v246
                                                                                                            }
                                                                                                            fmt.Println(v34, fi10)
                                                                                                            switch ("a" + "") {
                                                                                                            case "":
                                                                                                                if (!v61) {
                                                                                                                    var v247 uint = 1
                                                                                                                    _ = v247
                                                                                                                } else {
                                                                                                                    fmt.Println(fi10, v68, v34)
                                                                                                                    if len(v17) > 0 {
                                                                                                                        v17[0] = 2147483647
                                                                                                                    }
                                                                                                                    fmt.Println(v61)
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                var v248 bool = v61
                                                                                                                _ = v248
                                                                                                            default:
                                                                                                                v66 = v66
                                                                                                            }
                                                                                                            _ = v121
                                                                                                            _ = v122
                                                                                                        }
                                                                                                        _ = v120
                                                                                                    }
                                                                                                default:
                                                                                                    v68 -= 2
                                                                                                }
                                                                                            }
                                                                                            _ = len(v2)
                                                                                            v249 := GenS7{F0: 1}
                                                                                            if v61 {
                                                                                                if (fi10 > (fi14 - (fi10 | fi10))) {
                                                                                                    if (715 != 828) {
                                                                                                        var v250 int32 = v34
                                                                                                        if (!false) {
                                                                                                            var v251 int = fi10
                                                                                                            var v252 bool = v61
                                                                                                            switch "\n\t" {
                                                                                                            case "foo bar":
                                                                                                                switch fi10 {
                                                                                                                case 60:
                                                                                                                    var v253 bool = false
                                                                                                                    _ = v253
                                                                                                                case -585:
                                                                                                                    var v254 int16 = 772
                                                                                                                    _ = v254
                                                                                                                }
                                                                                                            case "":
                                                                                                                var v255 int8 = v4
                                                                                                                _ = v255
                                                                                                            default:
                                                                                                                var v256 uint = uint(v68)
                                                                                                                _ = v256
                                                                                                            }
                                                                                                            var v257 float32 = (0.0 * 0.1)
                                                                                                            fmt.Println(fi14, v61)
                                                                                                            _ = v251
                                                                                                            _ = v252
                                                                                                            _ = v257
                                                                                                        }
                                                                                                        _ = v250
                                                                                                    } else {
                                                                                                        var v258 float32 = (float32((1.5 - ((1.5 / (0.0 - -2.25)) / 1.0))) + (0.1 + ((1e10 - 1e10) + 1e10)))
                                                                                                        if len(v17) > 0 {
                                                                                                            v17[0] = (^(fi14))
                                                                                                        }
                                                                                                        v104 += v66
                                                                                                        _ = v258
                                                                                                    }
                                                                                                    var v259 int64 = 42
                                                                                                    v66 += v66
                                                                                                    var v260 float64 = 0.1
                                                                                                    if true {
                                                                                                        v68 -= (^(v68))
                                                                                                        var v261 uint64 = 371
                                                                                                        v66 -= 569
                                                                                                        _ = v261
                                                                                                    } else {
                                                                                                        var v262 bool = (!true)
                                                                                                        _ = v262
                                                                                                    }
                                                                                                    _ = v259
                                                                                                    _ = v260
                                                                                                } else {
                                                                                                    switch "hello" {
                                                                                                    case "":
                                                                                                        v66 = -9223372036854775808
                                                                                                    case "foo bar":
                                                                                                        var v263 bool = (-1 <= -1)
                                                                                                        _ = v263
                                                                                                    case "\n\t":
                                                                                                        v264 := GenS32{F0: 1, F1: (v104 / ((0) | 1)), F2: 115}
                                                                                                        if true {
                                                                                                            if v61 {
                                                                                                                fmt.Println(v68, fi14, v66)
                                                                                                                var v265 int32 = v34
                                                                                                                fmt.Println(v61, fi10)
                                                                                                                _ = v265
                                                                                                            } else {
                                                                                                                var v266 int8 = (v4 >> ((65535) & 63))
                                                                                                                _ = v266
                                                                                                            }
                                                                                                        }
                                                                                                        v267 := GenS33{F0: (^(v1))}
                                                                                                        _ = v264
                                                                                                        _ = v267
                                                                                                    default:
                                                                                                        v17 = append(v17, (fi14 * (^(fi14))))
                                                                                                        v268 := GenS2{F0: (v4 / ((v4) | 1))}
                                                                                                        _ = v268
                                                                                                    }
                                                                                                    fmt.Println(v104)
                                                                                                }
                                                                                                var v269 float64 = (-2.25 + 0.0)
                                                                                                _ = v269
                                                                                            } else {
                                                                                                v68 -= (^(v68))
                                                                                            }
                                                                                            v270 := GenS27{F0: v4, F1: 32767, F2: v104}
                                                                                            fmt.Println(v34, fi10, v4)
                                                                                            _ = v249
                                                                                            _ = v270
                                                                                        }
                                                                                    }
                                                                                    var v271 string = (((("a" + "") + ("" + ("hello" + "a"))) + "hello") + ("" + "hello"))
                                                                                    _ = v271
                                                                                case 9223372036854775807:
                                                                                    if (!true) {
                                                                                        var v272 bool = (v1 <= 32)
                                                                                        if v272 {
                                                                                            var v273 bool = v61
                                                                                            _ = v273
                                                                                        }
                                                                                        v274 := GenS8{F0: 32767}
                                                                                        _ = v272
                                                                                        _ = v274
                                                                                    } else {
                                                                                        if ((56 != int16(v34)) && true) {
                                                                                            v4 = (v4 & (v4 % ((-128) | 1)))
                                                                                            if (!(v61 || v61)) {
                                                                                                var v275 uint32 = v68
                                                                                                _ = v275
                                                                                            }
                                                                                            if (v104 == (v66 | v104)) {
                                                                                                var v276 int64 = (v104 * 42)
                                                                                                _ = v276
                                                                                            }
                                                                                        } else {
                                                                                            switch fi14 {
                                                                                            case -639:
                                                                                                if v61 {
                                                                                                    var v277 int = -9223372036854775808
                                                                                                    switch fi14 {
                                                                                                    case -1:
                                                                                                        switch "" {
                                                                                                        case "hello":
                                                                                                            if v61 {
                                                                                                                v68 -= v68
                                                                                                                if (!v61) {
                                                                                                                    var v278 float32 = (-2.25 - 1e10)
                                                                                                                    v66 = (v104 << ((42) & 63))
                                                                                                                    if (0 != 0) {
                                                                                                                        var v279 int16 = 561
                                                                                                                        var v280 uint8 = 0
                                                                                                                        _ = v279
                                                                                                                        _ = v280
                                                                                                                    }
                                                                                                                    v34 = (v34 >> ((42) & 63))
                                                                                                                    _ = v278
                                                                                                                } else {
                                                                                                                    switch "\n\t" {
                                                                                                                    case "":
                                                                                                                        switch "\n\t" {
                                                                                                                        case "hello":
                                                                                                                            v104 = (v66 ^ -9223372036854775808)
                                                                                                                            var v281 int32 = v34
                                                                                                                            if ("\n\t" > ("foo bar" + "hello")) {
                                                                                                                                v282 := GenS17{F0: 130, F1: "\n\t"}
                                                                                                                                _ = v282
                                                                                                                            }
                                                                                                                            fmt.Println(v1, v104)
                                                                                                                            _ = v281
                                                                                                                        case "\n\t":
                                                                                                                            fmt.Println(fi14)
                                                                                                                            var v283 uint = 2
                                                                                                                            _ = v283
                                                                                                                        default:
                                                                                                                            fmt.Println(v104, v1)
                                                                                                                        }
                                                                                                                    case "hello":
                                                                                                                        if len(v17) >= 2 {
                                                                                                                            v17 = v17[1:len(v17)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v284 uint32 = (v68 * 2)
                                                                                                                    switch "" {
                                                                                                                    case "":
                                                                                                                        v34 -= v34
                                                                                                                    case "hello":
                                                                                                                        var v285 int16 = 126
                                                                                                                        _ = v285
                                                                                                                    }
                                                                                                                    _ = v284
                                                                                                                }
                                                                                                                var v286 int32 = v34
                                                                                                                _ = v286
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            if v61 {
                                                                                                                fmt.Println(v61, v66)
                                                                                                                v287 := GenS34{F0: 1, F1: 42}
                                                                                                                _ = v287
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        v2[v68] = -9223372036854775808
                                                                                                    }
                                                                                                    v34 = (v34 >> ((65535) & 63))
                                                                                                    _ = v277
                                                                                                } else {
                                                                                                    v288 := GenS29{F0: int8(v34), F1: 255}
                                                                                                    v289 := GenS33{F0: v1}
                                                                                                    _ = v288
                                                                                                    _ = v289
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v290 uint16 = (v1 << ((18446744073709551615) & 63))
                                                                                        v66 = (v66 | (v104 % ((v66) | 1)))
                                                                                        _ = v290
                                                                                    }
                                                                                case -1:
                                                                                    v1 += (^(v1))
                                                                                    var v291 int64 = v66
                                                                                    fmt.Println(v66, v68, v4)
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        v1 = (v1 / ((1) | 1))
                                                                                    case "hello":
                                                                                        switch "\n\t" {
                                                                                        case "foo bar":
                                                                                            v1 = 42
                                                                                        case "\n\t":
                                                                                            switch ("hello" + ((("hello" + ((("a" + "") + (("hello" + "\n\t") + "")) + (("a" + ("\n\t" + (("foo bar" + "") + "a"))) + ("\n\t" + "a")))) + "\n\t") + "a")) {
                                                                                            case "hello":
                                                                                                v292 := GenS19{F0: int(v34), F1: (v4 % ((-103) | 1))}
                                                                                                _ = v292
                                                                                            case "":
                                                                                                v1 = 0
                                                                                            default:
                                                                                                var v293 float64 = 3.14159
                                                                                                _ = v293
                                                                                            }
                                                                                            if v61 {
                                                                                                v1 += (v1 | 0)
                                                                                                v294 := GenS35{F0: (true || v61)}
                                                                                                v295 := GenS9{F0: 0, F1: 758, F2: 218}
                                                                                                _ = v294
                                                                                                _ = v295
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v68, v68, fi10)
                                                                                    _ = v291
                                                                                default:
                                                                                    v61 = v61
                                                                                }
                                                                                v296 := GenS35{F0: (719 <= 0)}
                                                                                _ = v104
                                                                                _ = v296
                                                                            }
                                                                        }
                                                                    case "":
                                                                        var v297 uint32 = (v68 + (v68 * v68))
                                                                        _ = v297
                                                                    default:
                                                                        if v61 {
                                                                            v298 := GenS36{F0: 32767, F1: (v4 | int8(v1))}
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                v299 := GenS26{F0: v4}
                                                                                v61 = (!false)
                                                                                var v300 int = (fi10 + (fi10 | (fi10 & (fi14 >> ((894) & 63)))))
                                                                                _ = v299
                                                                                _ = v300
                                                                            case "hello":
                                                                                var v301 int = fi14
                                                                                _ = v301
                                                                            case "\n\t":
                                                                                var v302 int = (fi14 >> ((18446744073709551615) & 63))
                                                                                _ = v302
                                                                            default:
                                                                                v303 := GenS37{F0: 91, F1: v66}
                                                                                _ = v303
                                                                            }
                                                                            fmt.Println(v68, v66, v34)
                                                                            _ = v298
                                                                        }
                                                                    }
                                                                    if false {
                                                                        v34 = (^(v34))
                                                                        if ((float64(float32((float64(-2.25) * float64(0.0)))) * float64((0.1 + -2.25))) == (((1e10 * -2.25) - 0.0) * 3.14159)) {
                                                                            if len(v17) > 0 {
                                                                                v17[0] = (fi10 * int(v68))
                                                                            }
                                                                        } else {
                                                                            var v304 uint = 65535
                                                                            _ = v304
                                                                        }
                                                                        _ = len(v2)
                                                                    } else {
                                                                        var v305 int16 = 275
                                                                        _ = v305
                                                                    }
                                                                }
                                                                if len(v17) >= 2 {
                                                                    v17 = v17[1:len(v17)]
                                                                }
                                                                v306 := v3[(fi14 + fi10)]
                                                                v307 := GenS16{F0: 2147483647}
                                                                _ = v306
                                                                _ = v307
                                                                _ = v61
                                                                _ = v66
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                            v4 += v4
                                                        }
                                                    }
                                                    var v308 uint8 = 171
                                                    _ = v308
                                                }
                                            }
                                            v2[65535] = fi10
                                            if true {
                                                if (!false) {
                                                    var v309 uint = 317
                                                    _ = v309
                                                } else {
                                                    if true {
                                                        if len(v17) >= 2 {
                                                            v17 = v17[1:len(v17)]
                                                        }
                                                    }
                                                    v1 = v1
                                                }
                                            } else {
                                                if (18446744073709551615 >= 255) {
                                                    switch -1 {
                                                    case 32767:
                                                        var v310 string = "a"
                                                        _ = v310
                                                    case 128:
                                                        if true {
                                                            v34 = (^(v34))
                                                        } else {
                                                            v34 -= (v34 ^ v34)
                                                            v1 = 279
                                                            fmt.Println(v1)
                                                            if len(v17) > 0 {
                                                                v17[0] = int(v4)
                                                            }
                                                        }
                                                    }
                                                }
                                                var v311 float64 = 0.0
                                                _ = v311
                                            }
                                        }
                                    case 530:
                                        if false {
                                            if ((v34 >> ((802) & 63)) != v34) {
                                                delete(v2, 1)
                                                if true {
                                                    if true {
                                                        var v312 uint16 = 0
                                                        _ = v312
                                                    } else {
                                                        v4 = (-(v4))
                                                        if false {
                                                            v313 := GenS38{F0: -1, F1: 1000}
                                                            if len(v17) > 0 {
                                                                v17[0] = fi10
                                                            }
                                                            _ = v313
                                                        } else {
                                                            fmt.Println(v34, fi10)
                                                            v34 -= v34
                                                        }
                                                        var v314 int8 = v4
                                                        if (("\n\t" + "foo bar") > "foo bar") {
                                                            switch "a" {
                                                            case "a":
                                                                v1 = v1
                                                            case "\n\t":
                                                                var v315 int = fi10
                                                                _ = v315
                                                            }
                                                            var v316 uint8 = 82
                                                            if ((!(-573 > 0)) && (1 < 2)) {
                                                                fmt.Println(fi14, v34, v1)
                                                            } else {
                                                                v4 = (v4 - -128)
                                                            }
                                                            _ = v316
                                                        }
                                                        _ = v314
                                                    }
                                                    delete(v2, 65535)
                                                    var v317 uint = 255
                                                    _ = v317
                                                }
                                                fmt.Println(v34)
                                            } else {
                                                if ((false || (false && ((int32(v4) <= v34) && true))) && false) {
                                                    var v318 uint64 = 1
                                                    _ = v318
                                                }
                                                v34 += v34
                                                v34 += (v34 | 2147483647)
                                            }
                                            switch (((("hello" + "foo bar") + "a") + "a") + ("hello" + "hello")) {
                                            case "a":
                                                v4 -= (v4 | 0)
                                            case "hello":
                                                if true {
                                                    v4 = (v4 - v4)
                                                    var v319 int8 = (v4 | (v4 % (((v4 * (v4 + (v4 * int8(v1))))) | 1)))
                                                    var v320 bool = (true || (true || false))
                                                    _ = v319
                                                    _ = v320
                                                } else {
                                                    if true {
                                                        switch fi10 {
                                                        case 128:
                                                            v4 = v4
                                                        default:
                                                            v321 := GenS39{F0: 2}
                                                            _ = v321
                                                        }
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v322 := GenS40{F0: 142, F1: 0}
                                                            if ((3.14159 / (3.14159 / -2.25)) != 0.0) {
                                                                if false {
                                                                    v323 := GenS31{F0: 2, F1: false, F2: -1}
                                                                    if false {
                                                                        fmt.Println(v34, fi10, v4)
                                                                    } else {
                                                                        v324 := GenS27{F0: 0, F1: 1000, F2: -1}
                                                                        v34 = -932
                                                                        var v325 uint64 = 398
                                                                        _ = v324
                                                                        _ = v325
                                                                    }
                                                                    _ = v323
                                                                } else {
                                                                    _ = len(v2)
                                                                    var v326 float32 = (float32(0.1) + 1.5)
                                                                    _ = v326
                                                                }
                                                                v327 := GenS30{F0: 0, F1: 2}
                                                                _ = v327
                                                            }
                                                            _ = v322
                                                        case "hello":
                                                            v4 = (v4 % ((-1) | 1))
                                                        case "\n\t":
                                                            v2[65535] = fi10
                                                        }
                                                        var v328 int32 = (v34 * v34)
                                                        _ = v328
                                                    }
                                                }
                                            case "foo bar":
                                                if len(v17) > 0 {
                                                    v17[0] = -96
                                                }
                                            }
                                            v34 = (v34 & (v34 << ((18446744073709551615) & 63)))
                                        }
                                    case -9223372036854775808:
                                        fmt.Println(v1, v4, v34)
                                    default:
                                        switch fi14 {
                                        case -165:
                                            v4 = (-(v4))
                                        case 127:
                                            var v329 int8 = 0
                                            _ = v329
                                        default:
                                            switch "a" {
                                            case "":
                                                var v330 float32 = (1e10 / (((1e10 / 3.14159) - 3.14159) + (((1e10 + -2.25) / 1e10) * ((1.5 * (0.1 + ((0.1 / (1e10 * 1.5)) + (3.14159 / 3.14159)))) * -2.25))))
                                                _ = v330
                                            case "a":
                                                if false {
                                                    switch "a" {
                                                    case "\n\t":
                                                        v1 = (v1 >> ((2) & 63))
                                                    case "a":
                                                        var v331 int16 = 1000
                                                        _ = v331
                                                    default:
                                                        switch (fi10 | 1) {
                                                        case 295:
                                                            v332 := GenS8{F0: v1}
                                                            _ = v332
                                                        case 0:
                                                            var v333 int32 = -2147483648
                                                            _ = v333
                                                        case 4294967295:
                                                            v334 := GenS41{F0: 1, F1: fi14, F2: (v4 << ((0) & 63))}
                                                            _ = v334
                                                        default:
                                                            fmt.Println(fi10, v1, v1)
                                                            if (true || true) {
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    var v335 int16 = 0
                                                                    _ = v335
                                                                }
                                                                fmt.Println(fi10, fi10)
                                                                var v336 int16 = -905
                                                                _ = v336
                                                            } else {
                                                                var v337 uint = 532
                                                                fmt.Println(fi10)
                                                                v4 = (v4 | v4)
                                                                _ = v337
                                                            }
                                                            var v338 int32 = v34
                                                            _ = v338
                                                        }
                                                    }
                                                } else {
                                                    if true {
                                                        fmt.Println(v1, v1, fi10)
                                                        var v339 uint8 = 128
                                                        _ = v339
                                                    }
                                                    fmt.Println(v34, v1, v1)
                                                    var v340 bool = (0 >= 375)
                                                    _ = v340
                                                }
                                            case "foo bar":
                                                v17 = append(v17, fi14)
                                            default:
                                                fmt.Println(v4, v1)
                                            }
                                            fmt.Println(fi10)
                                        }
                                    }
                                    switch "" {
                                    case "a":
                                        v1 = (v1 * (v1 << ((2) & 63)))
                                        _ = len(v17)
                                        fmt.Println(fi14)
                                    case "foo bar":
                                        var v341 int8 = v4
                                        switch "" {
                                        case "a":
                                            if false {
                                                v4 -= (v341 % ((-128) | 1))
                                                v342 := GenS5{F0: (fi14 & 4294967295)}
                                                _ = v342
                                            }
                                        case "\n\t":
                                            v2[2147483647] = (-(fi10))
                                        }
                                        switch "foo bar" {
                                        case "":
                                            if ((748 <= 113) || (true || (42 <= -32768))) {
                                                fmt.Println(v341)
                                            }
                                        case "a":
                                            if len(v17) >= 2 {
                                                v17 = v17[1:len(v17)]
                                            }
                                            if (!(!(1 >= 1))) {
                                                var v343 int16 = 1
                                                if (!(0.0 == 1.5)) {
                                                    v341 += v341
                                                    switch (fi10 + -1) {
                                                    case 32767:
                                                        v1 = (v1 * 65535)
                                                    case 89:
                                                        var v344 float32 = 0.0
                                                        _ = v344
                                                    case 242:
                                                        var v345 uint = 10
                                                        _ = v345
                                                    default:
                                                        v341 = (v341 / ((-30) | 1))
                                                    }
                                                    v17 = append(v17, fi10)
                                                    if len(v17) >= 2 {
                                                        v17 = v17[1:len(v17)]
                                                    }
                                                    if (float32(3.14159) == (1.5 / -2.25)) {
                                                        switch ("" + "foo bar") {
                                                        case "foo bar":
                                                            var v346 uint16 = (v1 * v1)
                                                            _ = v346
                                                        case "hello":
                                                            v17 = append(v17, fi10)
                                                        default:
                                                            if ((float32((3.14159 - (1.5 * 3.14159))) / 1.5) != ((0.1 * 1e10) * 3.14159)) {
                                                                v347 := GenS13{F0: v4}
                                                                var v348 float64 = 0.1
                                                                _ = v347
                                                                _ = v348
                                                            }
                                                        }
                                                        fmt.Println(v34)
                                                    } else {
                                                        fmt.Println(v4, fi10, v1)
                                                        var v349 uint8 = 34
                                                        v341 -= 127
                                                        _ = v349
                                                    }
                                                    if false {
                                                        switch (("foo bar" + ("foo bar" + "a")) + "\n\t") {
                                                        case "foo bar":
                                                            switch "a" {
                                                            case "a":
                                                                var v350 uint8 = 128
                                                                _ = v350
                                                            }
                                                        case "":
                                                            v4 += -45
                                                        case "a":
                                                            v351 := GenS15{F0: -128, F1: "foo bar", F2: (v34 << ((560) & 63))}
                                                            _ = v351
                                                        default:
                                                            var v352 uint64 = 1
                                                            _ = v352
                                                        }
                                                        var v353 bool = false
                                                        _ = v353
                                                    } else {
                                                        v354 := GenS42{F0: 0}
                                                        var v355 uint = 65535
                                                        _ = v354
                                                        _ = v355
                                                    }
                                                    v356 := GenS21{F0: v34, F1: fi10}
                                                    _ = v356
                                                }
                                                _ = v343
                                            } else {
                                                if len(v17) > 0 {
                                                    v17[0] = -9223372036854775808
                                                }
                                            }
                                        default:
                                            switch "a" {
                                            case "foo bar":
                                                if ((("foo bar" + "foo bar") < "hello") && true) {
                                                    var v357 uint32 = 794
                                                    v34 += 170
                                                    _ = v357
                                                }
                                            case "":
                                                v358 := GenS13{F0: (v341 | 127)}
                                                _ = v358
                                            default:
                                                v34 += int32(fi10)
                                                v4 += -96
                                                fmt.Println(fi10)
                                            }
                                        }
                                        _ = v341
                                    }
                                    _ = v30
                                    _ = v34
                                }
                                if ((false && false) && false) {
                                    switch -711 {
                                    case 127:
                                        v1 += v1
                                    case -139:
                                        var v359 float32 = 1e10
                                        var v360 uint16 = (v1 % (((v1 * 65535)) | 1))
                                        fmt.Println(v4, fi10, fi10)
                                        _ = v359
                                        _ = v360
                                    }
                                    var v361 int32 = 2147483647
                                    v361 = -32
                                    _ = v361
                                } else {
                                    v362 := GenS43{F0: 255, F1: 511}
                                    var v363 uint64 = uint64(v1)
                                    v3[-1] = 0
                                    _ = v362
                                    _ = v363
                                }
                            } else {
                                v364 := GenS39{F0: 18446744073709551615}
                                v4 = (v4 >> ((4294967295) & 63))
                                _ = v364
                            }
                        } else {
                            if (927 > -671) {
                                switch "a" {
                                case "foo bar":
                                    v1 = v1
                                }
                                var v365 bool = false
                                _ = v365
                            }
                        }
                        _ = v16
                        _ = v17
                    }
                }
                v366 := []int{(fi10 | fi14)}
                switch "hello" {
                case "\n\t":
                    if (!(-128 != 127)) {
                        v367 := GenS28{F0: 42, F1: v4}
                        _ = v367
                    } else {
                        var v368 uint = 255
                        switch (fi10 ^ 32767) {
                        case 249:
                            fmt.Println(fi10, fi10, fi10)
                        }
                        switch fi10 {
                        case -9223372036854775808:
                            v369 := v3[(-(fi14))]
                            _ = v369
                        case 1:
                            delete(v3, (fi10 + (fi14 - (fi10 * (fi14 - fi14)))))
                        case 32768:
                            v370 := v3[fi14]
                            v370 += 1
                            _ = v370
                        }
                        _ = v368
                    }
                case "a":
                    var v371 uint32 = 65535
                    _ = v371
                case "hello":
                    v372 := v2[1]
                    _ = v372
                default:
                    fmt.Println(fi14, fi14)
                }
                _ = fi14
                _ = v366
            }
            _ = fi10
        }
    } else {
        if false {
            fmt.Println(v1)
            switch int(v4) {
            case 1:
                for fi373 := 0; fi373 < 12; fi373++ {
                    if ("\n\t" == "\n\t") {
                        switch -9223372036854775808 {
                        case 540:
                            var v374 uint64 = 734
                            _ = v374
                        default:
                            fmt.Println(v4, v4)
                        }
                    } else {
                        v4 += v4
                        var v375 int64 = 479
                        _ = v375
                    }
                    var v376 int32 = 42
                    for fi377 := 0; fi377 < 3; fi377++ {
                        v378 := GenS44{F0: ((v1 >> ((65535) & 63)) == v1), F1: -175, F2: 0}
                        v4 = 49
                        _ = fi377
                        _ = v378
                    }
                    _ = fi373
                    _ = v376
                }
            default:
                if (2 <= 255) {
                    v1 = (v1 ^ 2)
                    switch 128 {
                    case -9223372036854775808:
                        v379 := GenS34{F0: 635, F1: 32767}
                        _ = v379
                    case 134:
                        var v380 int8 = v4
                        _ = v380
                    }
                }
            }
        }
        var v381 int8 = (v4 ^ v4)
        _ = v381
    }
    switch "foo bar" {
    case "foo bar":
        var v382 int = int(v1)
        _ = v382
    case "hello":
        if (((0.1 / (3.14159 / 3.14159)) / (0.1 - float32(1e10))) <= (((1.5 * 1.5) + ((3.14159 * 0.1) / 1.0)) / (1.5 * ((-2.25 / 1.0) + 1e10)))) {
            v383 := []int{128, 0, -858}
            v384 := GenS12{F0: false, F1: v4, F2: -9223372036854775808}
            fmt.Println(v4, v1)
            _ = v383
            _ = v384
        }
    case "\n\t":
        v4 += 127
    default:
        if (323 >= 127) {
            v385 := v2[294]
            _ = v385
        } else {
            for fi386 := 0; fi386 < 20; fi386++ {
                v1 = 32767
                _ = fi386
            }
        }
    }
    fmt.Println(v1)
    for fi387 := 0; fi387 < 6; fi387++ {
        for fi388 := 0; fi388 < 17; fi388++ {
            v389 := []int{fi387, fi387, fi387, fi387}
            _ = fi388
            _ = v389
        }
        for fi390 := 0; fi390 < 17; fi390++ {
            var v391 uint8 = 31
            if true {
                var v392 uint32 = 342
                var v393 string = "foo bar"
                if true {
                    if true {
                        switch 1 {
                        case 32767:
                            switch (v393 + v393) {
                            case "a":
                                var v394 uint = 4294967295
                                switch v393 {
                                case "":
                                    if false {
                                        var v395 int64 = 9223372036854775807
                                        var v396 string = (("" + v393) + v393)
                                        _ = v395
                                        _ = v396
                                    }
                                case "\n\t":
                                    v397 := GenS45{F0: 0}
                                    _ = v397
                                case "hello":
                                    switch "hello" {
                                    case "":
                                        if (!(-620 >= -9223372036854775808)) {
                                            v392 += v392
                                            if (1e10 <= 3.14159) {
                                                v394 = v394
                                                v394 = (v394 ^ (v394 - v394))
                                                if ((v392 >> (((v394 / (((^(v394))) | 1))) & 63)) == 377) {
                                                    v398 := GenS46{F0: 42}
                                                    if false {
                                                        if (!(-1 >= 2147483647)) {
                                                            if true {
                                                                v4 -= v4
                                                                switch (((((v393 + v393) + ("" + v393)) + "foo bar") + (v393 + v393)) + (v393 + v393)) {
                                                                case "\n\t":
                                                                    if true {
                                                                        var v399 string = v393
                                                                        delete(v2, v392)
                                                                        _ = v399
                                                                    }
                                                                default:
                                                                    if (uint(v391) <= v394) {
                                                                        fmt.Println(v394)
                                                                    } else {
                                                                        if true {
                                                                            switch 203 {
                                                                            case -1:
                                                                                var v400 uint64 = 42
                                                                                _ = v400
                                                                            }
                                                                            v391 += 0
                                                                            var v401 float32 = -2.25
                                                                            _ = v401
                                                                        } else {
                                                                            var v402 int = fi390
                                                                            v403 := []int{}
                                                                            _ = v402
                                                                            _ = v403
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            v391 = (v391 * (v391 + (^(v391))))
                                                        }
                                                        v404 := []int{(fi390 * (fi387 / ((fi390) | 1)))}
                                                        _ = len(v404)
                                                        _ = v404
                                                    }
                                                    v392 = 2
                                                    _ = v398
                                                } else {
                                                    v405 := v3[(fi390 % ((fi390) | 1))]
                                                    if false {
                                                        v406 := GenS42{F0: 9223372036854775807}
                                                        fmt.Println(fi390, v405, fi390)
                                                        var v407 float32 = ((3.14159 - -2.25) * 0.1)
                                                        if (((v407 * v407) * 1.5) != 3.14159) {
                                                            switch fi387 {
                                                            case -8:
                                                                if true {
                                                                    var v408 int32 = 42
                                                                    if (false && ((v4 % ((v4) | 1)) > v4)) {
                                                                        _ = len(v2)
                                                                        var v409 uint8 = (^(v391))
                                                                        if true {
                                                                            var v410 int64 = int64(v408)
                                                                            if (false && false) {
                                                                                if false {
                                                                                    fmt.Println(v405)
                                                                                    if ((v1 <= v1) || true) {
                                                                                        v409 -= (v391 % ((128) | 1))
                                                                                        switch fi390 {
                                                                                        case 623:
                                                                                            v405 += v405
                                                                                        }
                                                                                        var v411 uint = v394
                                                                                        _ = v411
                                                                                    } else {
                                                                                        var v412 int = fi387
                                                                                        var v413 uint8 = v409
                                                                                        if ((v413 | v409) == 0) {
                                                                                            _ = len(v2)
                                                                                        }
                                                                                        _ = v412
                                                                                        _ = v413
                                                                                    }
                                                                                    v405 -= 4294967295
                                                                                    if true {
                                                                                        var v414 string = v393
                                                                                        var v415 int = (fi390 * fi387)
                                                                                        if (true && (!true)) {
                                                                                            var v416 uint8 = (v409 - 126)
                                                                                            v417 := []int{}
                                                                                            fmt.Println(v391, v405)
                                                                                            _ = v416
                                                                                            _ = v417
                                                                                        } else {
                                                                                            var v418 int32 = v408
                                                                                            v392 = v392
                                                                                            _ = v418
                                                                                        }
                                                                                        _ = v414
                                                                                        _ = v415
                                                                                    } else {
                                                                                        v405 = (v405 & (v392 % (((v392 ^ (v405 << ((v394) & 63)))) | 1)))
                                                                                        var v419 string = "hello"
                                                                                        _ = v419
                                                                                    }
                                                                                    if true {
                                                                                        v420 := GenS47{F0: 11}
                                                                                        var v421 int32 = (v408 & 280)
                                                                                        switch ((((v393 + v393) + ((v393 + (v393 + ("hello" + "hello"))) + ((("a" + v393) + v393) + v393))) + v393) + "hello") {
                                                                                        case "":
                                                                                            var v422 uint32 = v405
                                                                                            _ = v422
                                                                                        case "hello":
                                                                                            var v423 uint32 = (v405 << ((v394) & 63))
                                                                                            _ = v423
                                                                                        }
                                                                                        v424 := []int{-1, fi387, -9223372036854775808}
                                                                                        _ = v420
                                                                                        _ = v421
                                                                                        _ = v424
                                                                                    }
                                                                                } else {
                                                                                    if (true || true) {
                                                                                        switch v393 {
                                                                                        case "\n\t":
                                                                                            var v425 int64 = v410
                                                                                            _ = v425
                                                                                        default:
                                                                                            v426 := GenS3{F0: true, F1: (v409 | (v391 % ((v409) | 1)))}
                                                                                            _ = v426
                                                                                        }
                                                                                        v410 = (v410 >> (((v394 >> ((v394) & 63))) & 63))
                                                                                        v427 := []int{fi387, (-(fi390)), (-(fi390))}
                                                                                        _ = v427
                                                                                    } else {
                                                                                        v2[v405] = 2147483647
                                                                                        fmt.Println(fi390, v405, fi387)
                                                                                    }
                                                                                }
                                                                                v392 = (v392 % ((621) | 1))
                                                                                delete(v3, (fi387 ^ 127))
                                                                                _ = len(v3)
                                                                                fmt.Println(v405)
                                                                                v428 := []int{fi390, (fi390 & -194), (fi387 | (fi387 % ((int(v409)) | 1))), (fi387 << (((v394 * 2)) & 63))}
                                                                                var v429 uint = (v394 ^ (v394 - (v394 << (((v394 + v394)) & 63))))
                                                                                v2[v405] = 844
                                                                                _ = v428
                                                                                _ = v429
                                                                            } else {
                                                                                if (-1 > int64(v409)) {
                                                                                    var v430 float32 = v407
                                                                                    if false {
                                                                                        var v431 int = 4294967295
                                                                                        var v432 float64 = float64(v407)
                                                                                        _ = v431
                                                                                        _ = v432
                                                                                    } else {
                                                                                        v433 := GenS19{F0: fi387, F1: v4}
                                                                                        switch -9223372036854775808 {
                                                                                        case -697:
                                                                                            v393 = ("\n\t" + v393)
                                                                                        case 374:
                                                                                            delete(v2, (^(v392)))
                                                                                        case 2147483647:
                                                                                            delete(v3, 2147483647)
                                                                                        }
                                                                                        _ = v433
                                                                                    }
                                                                                    _ = v430
                                                                                } else {
                                                                                    if ((v4 << ((v394) & 63)) < v4) {
                                                                                        var v434 float32 = 0.1
                                                                                        v1 += v1
                                                                                        if (!(v408 >= v408)) {
                                                                                            if true {
                                                                                                v393 = "hello"
                                                                                            } else {
                                                                                                v435 := GenS34{F0: 965, F1: int16(v1)}
                                                                                                _ = v435
                                                                                            }
                                                                                            v409 = v391
                                                                                            switch fi390 {
                                                                                            case 2147483647:
                                                                                                v393 = ("" + v393)
                                                                                                var v436 uint64 = 822
                                                                                                if true {
                                                                                                    v409 = v409
                                                                                                    if ((v410 << ((v394) & 63)) < (^(v410))) {
                                                                                                        if ((v436 - (v436 + uint64(v394))) < v436) {
                                                                                                            if false {
                                                                                                                v437 := GenS29{F0: v4, F1: (v394 % ((uint(v436)) | 1))}
                                                                                                                v434 += (float32(1.5) * float32(0.1))
                                                                                                                var v438 int8 = v4
                                                                                                                _ = v437
                                                                                                                _ = v438
                                                                                                            } else {
                                                                                                                switch v393 {
                                                                                                                case "":
                                                                                                                    v439 := GenS10{F0: (v436 / ((uint64(v394)) | 1)), F1: (v405 << ((18446744073709551615) & 63))}
                                                                                                                    _ = v439
                                                                                                                }
                                                                                                                v407 = v407
                                                                                                                fmt.Println(v393, fi390, v434)
                                                                                                                switch (("hello" + v393) + "hello") {
                                                                                                                case "foo bar":
                                                                                                                    v393 = "\n\t"
                                                                                                                default:
                                                                                                                    var v440 float64 = -2.25
                                                                                                                    _ = v440
                                                                                                                }
                                                                                                                switch v393 {
                                                                                                                case "foo bar":
                                                                                                                    v436 = v436
                                                                                                                }
                                                                                                            }
                                                                                                            v441 := GenS13{F0: (v4 - 127)}
                                                                                                            _ = v441
                                                                                                        } else {
                                                                                                            switch v393 {
                                                                                                            case "hello":
                                                                                                                var v442 int64 = (v410 * v410)
                                                                                                                _ = v442
                                                                                                            case "foo bar":
                                                                                                                var v443 int8 = v4
                                                                                                                _ = v443
                                                                                                            }
                                                                                                            var v444 int32 = int32(v394)
                                                                                                            switch (v393 + "") {
                                                                                                            case "a":
                                                                                                                switch "a" {
                                                                                                                case "\n\t":
                                                                                                                    switch fi390 {
                                                                                                                    case -9223372036854775808:
                                                                                                                        if ((801 <= (v394 ^ (v394 ^ (^(v394))))) && true) {
                                                                                                                            v445 := GenS48{F0: (v405 ^ uint32(v409)), F1: (v392 % (((^(v405))) | 1))}
                                                                                                                            v4 += v4
                                                                                                                            _ = v445
                                                                                                                        }
                                                                                                                        var v446 float32 = 0.0
                                                                                                                        if true {
                                                                                                                            var v447 int16 = 0
                                                                                                                            _ = v447
                                                                                                                        }
                                                                                                                        _ = v446
                                                                                                                    default:
                                                                                                                        if (v436 > (v436 >> ((4294967295) & 63))) {
                                                                                                                            var v448 uint8 = (v391 / ((v409) | 1))
                                                                                                                            _ = v448
                                                                                                                        } else {
                                                                                                                            var v449 float32 = float32((0.1 / 1.5))
                                                                                                                            switch (fi387 & fi387) {
                                                                                                                            case 32768:
                                                                                                                                if false {
                                                                                                                                    v2[1] = 165
                                                                                                                                    v450 := GenS49{F0: -32768, F1: (^(v408)), F2: -9223372036854775808}
                                                                                                                                    switch "foo bar" {
                                                                                                                                    case "a":
                                                                                                                                        switch (^(fi387)) {
                                                                                                                                        case 799:
                                                                                                                                            fmt.Println(fi390)
                                                                                                                                        }
                                                                                                                                    case "\n\t":
                                                                                                                                        v451 := GenS33{F0: 1}
                                                                                                                                        _ = v451
                                                                                                                                    case "foo bar":
                                                                                                                                        _ = len(v3)
                                                                                                                                    }
                                                                                                                                    _ = v450
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                fmt.Println(v392, v434)
                                                                                                                            }
                                                                                                                            _ = v449
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                var v452 string = v393
                                                                                                                v436 += (v436 * v436)
                                                                                                                _ = v452
                                                                                                            case "\n\t":
                                                                                                                if (!(true && true)) {
                                                                                                                    if ((v405 << (((v394 + 42)) & 63)) != v392) {
                                                                                                                        fmt.Println(v444)
                                                                                                                        if true {
                                                                                                                            switch int(v436) {
                                                                                                                            case -1:
                                                                                                                                v453 := GenS7{F0: 42}
                                                                                                                                _ = v453
                                                                                                                            default:
                                                                                                                                v405 -= 2
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v454 := []int{(fi387 % (((fi387 * (^(fi390)))) | 1)), (fi390 << ((v394) & 63)), -9223372036854775808}
                                                                                                                            var v455 string = "\n\t"
                                                                                                                            switch fi387 {
                                                                                                                            case 4294967295:
                                                                                                                                v456 := GenS8{F0: (v1 % (((v1 << ((v394) & 63))) | 1))}
                                                                                                                                _ = v456
                                                                                                                            case 127:
                                                                                                                                v457 := GenS50{F0: (v444 >> (((v394 % ((v394) | 1))) & 63))}
                                                                                                                                _ = v457
                                                                                                                            case -705:
                                                                                                                                if true {
                                                                                                                                    fmt.Println(v455, v455)
                                                                                                                                    v444 -= (v408 - -316)
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v458 int = fi387
                                                                                                                            var v459 string = (v393 + v393)
                                                                                                                            _ = v454
                                                                                                                            _ = v455
                                                                                                                            _ = v458
                                                                                                                            _ = v459
                                                                                                                        }
                                                                                                                        if (v393 > v393) {
                                                                                                                            if (false || (fi390 > (fi390 ^ fi387))) {
                                                                                                                                var v460 uint16 = v1
                                                                                                                                delete(v2, (v392 * v392))
                                                                                                                                _ = v460
                                                                                                                            }
                                                                                                                            switch fi390 {
                                                                                                                            case 660:
                                                                                                                                var v461 uint8 = (v409 * (v409 << (((v394 >> ((v394) & 63))) & 63)))
                                                                                                                                _ = v461
                                                                                                                            case 128:
                                                                                                                                fmt.Println(v405, v1)
                                                                                                                            case 601:
                                                                                                                                v462 := GenS30{F0: v391, F1: (^(v394))}
                                                                                                                                _ = v462
                                                                                                                            default:
                                                                                                                                var v463 int16 = 0
                                                                                                                                v463 = v463
                                                                                                                                _ = v463
                                                                                                                            }
                                                                                                                            v1 = (v1 - (v1 % ((65535) | 1)))
                                                                                                                        }
                                                                                                                        v407 = (1.5 - v434)
                                                                                                                    } else {
                                                                                                                        v410 -= v410
                                                                                                                        v444 = int32(v409)
                                                                                                                    }
                                                                                                                    v394 -= (v394 + v394)
                                                                                                                    switch (fi387 >> ((v394) & 63)) {
                                                                                                                    case 32768:
                                                                                                                        v464 := GenS51{F0: v410, F1: fi390, F2: (v1 ^ (v1 & (v1 + v1)))}
                                                                                                                        _ = v464
                                                                                                                    default:
                                                                                                                        var v465 uint32 = v392
                                                                                                                        _ = v465
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                _ = len(v3)
                                                                                                            }
                                                                                                            _ = v444
                                                                                                        }
                                                                                                        if ((v392 + (v405 % (((v405 % ((v392) | 1))) | 1))) != v392) {
                                                                                                            var v466 bool = (!(!(true || false)))
                                                                                                            var v467 int = (fi387 + fi387)
                                                                                                            _ = v466
                                                                                                            _ = v467
                                                                                                        } else {
                                                                                                            v392 = 210
                                                                                                        }
                                                                                                        if false {
                                                                                                            var v468 int16 = 0
                                                                                                            var v469 int32 = -628
                                                                                                            _ = v468
                                                                                                            _ = v469
                                                                                                        }
                                                                                                        v470 := GenS52{F0: (false || true), F1: (v394 % ((v394) | 1))}
                                                                                                        v405 = 1
                                                                                                        _ = v470
                                                                                                    } else {
                                                                                                        v471 := []int{(-(fi390)), (-(fi390)), int(v408), (fi390 | (^(fi390)))}
                                                                                                        fmt.Println(v1, v394)
                                                                                                        _ = v471
                                                                                                    }
                                                                                                    var v472 int8 = v4
                                                                                                    _ = v472
                                                                                                }
                                                                                                _ = v436
                                                                                            case 1:
                                                                                                switch "foo bar" {
                                                                                                case "a":
                                                                                                    v1 -= 2
                                                                                                }
                                                                                                if ((-(v410)) <= 9223372036854775807) {
                                                                                                    switch v393 {
                                                                                                    case "hello":
                                                                                                        switch fi390 {
                                                                                                        case 9223372036854775807:
                                                                                                            if (!true) {
                                                                                                                switch 4294967295 {
                                                                                                                case 32768:
                                                                                                                    var v473 int64 = 42
                                                                                                                    _ = v473
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v474 int32 = (^(v408))
                                                                                                                var v475 uint8 = (^(v409))
                                                                                                                _ = v474
                                                                                                                _ = v475
                                                                                                            }
                                                                                                        case 790:
                                                                                                            if true {
                                                                                                                v476 := GenS4{F0: -117}
                                                                                                                fmt.Println(v393, v410)
                                                                                                                _ = v476
                                                                                                            } else {
                                                                                                                v405 = 750
                                                                                                            }
                                                                                                        }
                                                                                                        if (435 != 4294967295) {
                                                                                                            v477 := []int{fi387}
                                                                                                            v391 += (v409 | (v409 & (v409 + (v391 / ((1) | 1)))))
                                                                                                            v478 := GenS39{F0: 548}
                                                                                                            _ = v477
                                                                                                            _ = v478
                                                                                                        } else {
                                                                                                            switch 4294967295 {
                                                                                                            case -9223372036854775808:
                                                                                                                var v479 uint8 = v409
                                                                                                                _ = v479
                                                                                                            case 32768:
                                                                                                                v480 := []int{fi390, fi390, (fi387 + (fi387 + -9223372036854775808)), (fi390 & fi387)}
                                                                                                                _ = v480
                                                                                                            case 66:
                                                                                                                v407 += v407
                                                                                                            }
                                                                                                            switch fi387 {
                                                                                                            case -634:
                                                                                                                v481 := GenS22{F0: (v391 << (((v394 - v394)) & 63))}
                                                                                                                _ = v481
                                                                                                            case 1:
                                                                                                                fmt.Println(fi390, v4, v393)
                                                                                                                if ((!false) && (false && true)) {
                                                                                                                    var v482 uint16 = (v1 - v1)
                                                                                                                    _ = v482
                                                                                                                } else {
                                                                                                                    v483 := GenS2{F0: -128}
                                                                                                                    v484 := GenS28{F0: 0, F1: v4}
                                                                                                                    if ("a" < v393) {
                                                                                                                        fmt.Println(v392)
                                                                                                                        v394 -= 2
                                                                                                                    }
                                                                                                                    _ = v483
                                                                                                                    _ = v484
                                                                                                                }
                                                                                                            }
                                                                                                            v409 -= v409
                                                                                                        }
                                                                                                        switch v393 {
                                                                                                        case "hello":
                                                                                                            v392 -= v392
                                                                                                        case "foo bar":
                                                                                                            v485 := GenS53{F0: -1, F1: 42, F2: false}
                                                                                                            _ = v485
                                                                                                        default:
                                                                                                            fmt.Println(v409, fi390, fi390)
                                                                                                        }
                                                                                                    default:
                                                                                                        v3[(fi387 * -20)] = (^(v405))
                                                                                                    }
                                                                                                } else {
                                                                                                    v408 = (v408 << (((v394 / (((v394 >> ((65535) & 63))) | 1))) & 63))
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v4 = (v4 % (((^(v4))) | 1))
                                                                                        }
                                                                                        _ = v434
                                                                                    } else {
                                                                                        var v486 uint64 = 811
                                                                                        var v487 uint32 = uint32(v486)
                                                                                        _ = v486
                                                                                        _ = v487
                                                                                    }
                                                                                }
                                                                                v392 += 1
                                                                                var v488 uint8 = 1
                                                                                v410 = (v410 * (v410 * (v410 + v410)))
                                                                                switch (v393 + ((v393 + v393) + v393)) {
                                                                                case "":
                                                                                    v394 -= (v394 * v394)
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        switch v393 {
                                                                                        case "hello":
                                                                                            v489 := []int{9223372036854775807, 42, fi390}
                                                                                            _ = v489
                                                                                        case "":
                                                                                            if ((!false) || ((^(v392)) < v405)) {
                                                                                                var v490 float32 = v407
                                                                                                v491 := GenS23{F0: (v394 + (v394 ^ (^(v394)))), F1: (v410 | v410)}
                                                                                                _ = v490
                                                                                                _ = v491
                                                                                            }
                                                                                        case "a":
                                                                                            v492 := v2[(v392 ^ 1)]
                                                                                            v2[v405] = v492
                                                                                            v4 += (v4 / ((v4) | 1))
                                                                                            _ = v492
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(fi390)
                                                                                        switch (^(fi387)) {
                                                                                        case 9223372036854775807:
                                                                                            if ((false || (4294967295 == v405)) && (true && false)) {
                                                                                                v391 = v409
                                                                                                if (true && (0.1 != (float64(v407) / 1.0))) {
                                                                                                    v394 = (v394 * 18446744073709551615)
                                                                                                    v392 -= v405
                                                                                                } else {
                                                                                                    v493 := GenS54{F0: ((v4 <= v4) && false), F1: (fi390 >> ((v394) & 63))}
                                                                                                    _ = v493
                                                                                                }
                                                                                                var v494 int32 = (v408 - v408)
                                                                                                _ = v494
                                                                                            } else {
                                                                                                var v495 uint16 = (v1 << ((690) & 63))
                                                                                                v409 += (^(v488))
                                                                                                _ = v495
                                                                                            }
                                                                                        case 32767:
                                                                                            v496 := GenS10{F0: 1, F1: v392}
                                                                                            var v497 uint32 = (v392 ^ 920)
                                                                                            fmt.Println(v407, v393)
                                                                                            _ = v496
                                                                                            _ = v497
                                                                                        case 341:
                                                                                            fmt.Println(fi390)
                                                                                        default:
                                                                                            var v498 string = (v393 + ("a" + v393))
                                                                                            _ = v498
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v499 := v2[v392]
                                                                                    _ = v499
                                                                                }
                                                                                _ = v488
                                                                            }
                                                                            _ = v410
                                                                        }
                                                                        _ = v409
                                                                    }
                                                                    _ = v408
                                                                }
                                                            case 803:
                                                                if true {
                                                                    fmt.Println(fi387, v407, v394)
                                                                    v500 := []int{(fi390 >> ((v394) & 63)), (fi390 % ((604) | 1))}
                                                                    _ = v500
                                                                }
                                                            default:
                                                                switch v393 {
                                                                case "\n\t":
                                                                    if true {
                                                                        if ((!(1 == 1)) && false) {
                                                                            var v501 uint32 = (v405 | (v405 ^ v392))
                                                                            _ = v501
                                                                        }
                                                                        if false {
                                                                            v1 = uint16(v405)
                                                                        } else {
                                                                            var v502 float64 = (-2.25 * 1e10)
                                                                            _ = v502
                                                                        }
                                                                        v503 := GenS47{F0: 610}
                                                                        fmt.Println(v391, v1, v407)
                                                                        v392 += v392
                                                                        _ = v503
                                                                    } else {
                                                                        if (3.14159 != 1.5) {
                                                                            var v504 int = (fi387 ^ (fi390 / (((fi387 % ((fi390) | 1))) | 1)))
                                                                            fmt.Println(v1, v393)
                                                                            fmt.Println(v392, v392, v405)
                                                                            var v505 float64 = 3.14159
                                                                            v394 -= v394
                                                                            _ = v504
                                                                            _ = v505
                                                                        }
                                                                        v405 += v392
                                                                    }
                                                                }
                                                            }
                                                            v405 -= uint32(v394)
                                                            v407 -= float32(0.1)
                                                        } else {
                                                            var v506 uint32 = (v392 * 825)
                                                            if false {
                                                                v405 = (v392 / (((v392 & (v405 + uint32(v394)))) | 1))
                                                                v507 := GenS49{F0: 32767, F1: 0, F2: int64(v391)}
                                                                _ = v507
                                                            }
                                                            var v508 uint32 = 2147483648
                                                            v4 -= v4
                                                            _ = v506
                                                            _ = v508
                                                        }
                                                        var v509 uint16 = (^(v1))
                                                        v509 = (^(v509))
                                                        v510 := []int{fi390, fi387, 23}
                                                        _ = v406
                                                        _ = v407
                                                        _ = v509
                                                        _ = v510
                                                    }
                                                    if (((v405 & uint32(v394)) < 0) || (9223372036854775807 != -539)) {
                                                        switch fi390 {
                                                        case 4294967295:
                                                            var v511 int64 = -802
                                                            _ = v511
                                                        case 2147483647:
                                                            if (2147483648 >= uint32(v1)) {
                                                                switch 32767 {
                                                                case -9223372036854775808:
                                                                    var v512 bool = false
                                                                    var v513 int64 = 42
                                                                    if false {
                                                                        var v514 uint = 4294967295
                                                                        _ = v514
                                                                    } else {
                                                                        v515 := GenS6{F0: (-(v4))}
                                                                        var v516 uint8 = (v391 ^ v391)
                                                                        v516 = v516
                                                                        _ = v515
                                                                        _ = v516
                                                                    }
                                                                    _ = v512
                                                                    _ = v513
                                                                case 4294967295:
                                                                    fmt.Println(fi390)
                                                                case 324:
                                                                    if true {
                                                                        v4 += (v4 / ((127) | 1))
                                                                        var v517 string = (v393 + v393)
                                                                        var v518 int = int(v4)
                                                                        _ = v517
                                                                        _ = v518
                                                                    } else {
                                                                        v519 := []int{fi387}
                                                                        _ = v519
                                                                    }
                                                                }
                                                                v393 = (v393 + "foo bar")
                                                                var v520 int8 = v4
                                                                v1 = v1
                                                                if true {
                                                                    v521 := GenS55{F0: (false && false)}
                                                                    _ = v521
                                                                }
                                                                _ = v520
                                                            }
                                                        case 0:
                                                            switch v393 {
                                                            case "foo bar":
                                                                var v522 uint = v394
                                                                v523 := v3[(-(fi387))]
                                                                _ = v522
                                                                _ = v523
                                                            case "\n\t":
                                                                v394 -= 18446744073709551615
                                                            case "a":
                                                                v1 += (v1 + (v1 | uint16(v405)))
                                                            default:
                                                                v391 += (v391 + (v391 & v391))
                                                            }
                                                        }
                                                    }
                                                    _ = v405
                                                }
                                            }
                                        }
                                        if false {
                                            v524 := GenS53{F0: -32768, F1: 239, F2: true}
                                            _ = v524
                                        } else {
                                            _ = len(v3)
                                            fmt.Println(v1, fi387, v391)
                                        }
                                        var v525 bool = true
                                        _ = v525
                                    case "\n\t":
                                        v526 := GenS48{F0: v392, F1: (^(v392))}
                                        _ = v526
                                    case "foo bar":
                                        v394 = (v394 * 0)
                                    default:
                                        v527 := []int{-61, fi387}
                                        _ = v527
                                    }
                                default:
                                    fmt.Println(v392, v392, v394)
                                }
                                _ = v394
                            case "\n\t":
                                var v528 int = (fi387 ^ 32768)
                                _ = v528
                            case "hello":
                                v529 := v2[(v392 ^ 2)]
                                _ = v529
                            default:
                                v530 := []int{(fi390 ^ fi390), 0}
                                _ = v530
                            }
                            v1 = (v1 | v1)
                            var v531 uint16 = (v1 * (^(v1)))
                            _ = v531
                        case 9223372036854775807:
                            v532 := []int{(^(fi387)), int(v1), fi390, 244}
                            _ = v532
                        case -917:
                            switch v393 {
                            case "foo bar":
                                _ = len(v2)
                            case "hello":
                                v533 := []int{fi390, (fi387 ^ fi390), fi390}
                                _ = v533
                            default:
                                var v534 int64 = -1
                                _ = v534
                            }
                        default:
                            if ((-32768 == -992) || (979 != 632)) {
                                v535 := GenS17{F0: 92, F1: v393}
                                fmt.Println(v1, v393, v393)
                                switch "" {
                                case "":
                                    v536 := GenS20{F0: fi390}
                                    _ = v536
                                default:
                                    v4 = int8(v392)
                                }
                                _ = v535
                            } else {
                                switch "a" {
                                case "\n\t":
                                    v1 = (v1 - uint16(v391))
                                case "":
                                    var v537 int8 = v4
                                    fmt.Println(v392)
                                    v537 -= v4
                                    _ = v537
                                default:
                                    v538 := GenS40{F0: (v1 * (^(v1))), F1: uint(v392)}
                                    _ = v538
                                }
                            }
                            if (v4 <= -128) {
                                var v539 uint32 = v392
                                v540 := GenS56{F0: v4}
                                v541 := GenS27{F0: v4, F1: 32767, F2: 1}
                                _ = v539
                                _ = v540
                                _ = v541
                            } else {
                                _ = len(v2)
                                delete(v3, -2147483648)
                            }
                            switch -9223372036854775808 {
                            case 17:
                                if (float32((3.14159 + (0.1 * (((((float64(float32(-2.25)) * (0.1 / ((((1e10 * -2.25) - float64(0.0)) + float64(((-2.25 + -2.25) + (-2.25 + 0.0)))) / (0.1 / ((1.5 + -2.25) - 0.1))))) * ((0.0 - (1e10 * 0.0)) - 3.14159)) - (1.5 - 0.0)) + (3.14159 - ((-2.25 + (0.1 - (-2.25 - 3.14159))) / 1.0))) + 0.1)))) <= -2.25) {
                                    var v542 string = v393
                                    v543 := GenS57{F0: -1}
                                    _ = v542
                                    _ = v543
                                }
                            case 32768:
                                v544 := GenS12{F0: (v392 != (v392 ^ (v392 ^ v392))), F1: v4, F2: 42}
                                _ = v544
                            default:
                                var v545 string = v393
                                _ = v545
                            }
                            switch (^(fi390)) {
                            case -9223372036854775808:
                                v3[32767] = v392
                            case 2147483647:
                                v2[v392] = (^(fi387))
                            case 184:
                                v546 := GenS58{F0: 502, F1: 795}
                                _ = v546
                            default:
                                v4 = 0
                            }
                        }
                        if (false && true) {
                            if (int16(fi387) == -1) {
                                v547 := GenS47{F0: v392}
                                _ = v547
                            }
                            v548 := GenS59{F0: 42, F1: 42}
                            switch fi387 {
                            case -2147483648:
                                v549 := GenS60{F0: (v392 - v392), F1: 4294967295, F2: 116}
                                _ = v549
                            }
                            _ = v548
                        } else {
                            switch fi390 {
                            case 32767:
                                switch fi387 {
                                case 987:
                                    var v550 uint32 = 42
                                    _ = v550
                                case 1:
                                    var v551 uint = 42
                                    _ = v551
                                }
                            case -957:
                                if false {
                                    fmt.Println(v391, v393, fi387)
                                    var v552 bool = false
                                    _ = v552
                                }
                            }
                            if false {
                                v553 := GenS12{F0: true, F1: (v4 - v4), F2: (fi387 + fi390)}
                                v554 := GenS61{F0: v392, F1: (!(fi387 < fi387)), F2: 672}
                                if false {
                                    var v555 float64 = 0.0
                                    var v556 uint = 596
                                    v557 := []int{fi390, fi387, int(v556), int(v4)}
                                    var v558 int64 = 0
                                    _ = v555
                                    _ = v556
                                    _ = v557
                                    _ = v558
                                } else {
                                    var v559 uint32 = v392
                                    _ = v559
                                }
                                _ = v553
                                _ = v554
                            } else {
                                if false {
                                    if (486 >= 0) {
                                        v3[777] = 145
                                        var v560 int32 = 1
                                        _ = v560
                                    }
                                    if true {
                                        v4 += 0
                                    }
                                    v561 := []int{fi390, 4294967295, (fi390 * (fi390 | fi390))}
                                    _ = v561
                                } else {
                                    var v562 uint16 = (v1 + (v1 / ((v1) | 1)))
                                    _ = v562
                                }
                                if true {
                                    v563 := GenS42{F0: 281}
                                    if (!false) {
                                        v564 := v3[61]
                                        if (-228 > -197) {
                                            var v565 int8 = 100
                                            v2[0] = fi390
                                            _ = v565
                                        } else {
                                            fmt.Println(v393, v393)
                                            switch (fi387 | 4294967295) {
                                            case 2147483647:
                                                switch v393 {
                                                case "\n\t":
                                                    if (((!(true || true)) && (1 != 966)) && (v391 > uint8(v392))) {
                                                        v391 = v391
                                                    }
                                                case "":
                                                    v566 := v3[32767]
                                                    _ = v566
                                                case "hello":
                                                    v391 = v391
                                                }
                                            case 665:
                                                switch (fi390 - (^(fi387))) {
                                                case -43:
                                                    v567 := []int{(fi390 / (((fi390 ^ fi387)) | 1)), -368}
                                                    _ = v567
                                                case 9223372036854775807:
                                                    v564 = (v564 >> ((4294967295) & 63))
                                                case 32767:
                                                    switch fi387 {
                                                    case 32767:
                                                        var v568 float32 = float32((0.1 - (3.14159 * (3.14159 + float64((3.14159 - 1.5))))))
                                                        _ = v568
                                                    case 42:
                                                        v569 := []int{1}
                                                        _ = v569
                                                    case 382:
                                                        switch (^(fi387)) {
                                                        case 34:
                                                            switch v393 {
                                                            case "":
                                                                var v570 int32 = -520
                                                                _ = v570
                                                            default:
                                                                var v571 string = (("\n\t" + v393) + v393)
                                                                _ = v571
                                                            }
                                                        case 0:
                                                            fmt.Println(v392, v392)
                                                        case 4294967295:
                                                            if ((-70 <= 690) || (9223372036854775807 < 9223372036854775807)) {
                                                                var v572 float32 = 1e10
                                                                _ = v572
                                                            }
                                                        default:
                                                            var v573 uint16 = (v1 & v1)
                                                            _ = v573
                                                        }
                                                    default:
                                                        v392 -= v564
                                                    }
                                                }
                                            case 1:
                                                if (((-2.25 * (0.0 * ((0.1 / 1.5) - (float32((0.1 + (((1e10 - 1e10) * 1e10) - (float64(((0.1 / 1.5) + -2.25)) * (1.5 / float64(0.1)))))) / ((-2.25 / 0.1) - -2.25))))) / -2.25) <= 0.1) {
                                                    var v574 int = (fi390 << ((505) & 63))
                                                    v4 += (v4 & 100)
                                                    fmt.Println(fi387, v392, v4)
                                                    _ = v574
                                                } else {
                                                    v575 := GenS62{F0: 2, F1: v1, F2: (v564 << ((1) & 63))}
                                                    _ = v575
                                                }
                                            }
                                        }
                                        _ = v564
                                    }
                                    _ = v563
                                } else {
                                    v393 = ((v393 + (("foo bar" + ("\n\t" + v393)) + v393)) + ("foo bar" + (v393 + "")))
                                }
                                _ = len(v3)
                            }
                        }
                    }
                    if (-32768 <= 149) {
                        var v576 float32 = ((3.14159 / 1.5) + 1.5)
                        _ = v576
                    } else {
                        var v577 int32 = 766
                        var v578 string = "a"
                        if true {
                            v4 = (-(v4))
                        } else {
                            switch fi387 {
                            case 4294967295:
                                var v579 uint = 619
                                var v580 uint = (v579 ^ 844)
                                if false {
                                    v581 := GenS63{F0: (fi390 & fi390)}
                                    _ = v581
                                }
                                _ = v579
                                _ = v580
                            case 32768:
                                var v582 uint64 = 42
                                _ = v582
                            }
                        }
                        var v583 uint8 = 128
                        _ = v577
                        _ = v578
                        _ = v583
                    }
                } else {
                    v584 := GenS24{F0: (v4 & (-(v4))), F1: 18446744073709551615}
                    _ = v584
                }
                var v585 uint = 810
                v586 := []int{-2147483648, (fi387 ^ (fi390 + fi387)), int(v391)}
                _ = v392
                _ = v393
                _ = v585
                _ = v586
            }
            _ = fi390
            _ = v391
        }
        _ = fi387
    }
    v1 = 2
    v587 := GenS34{F0: -1, F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v587
}
