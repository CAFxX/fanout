package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int }
type GenS2 struct { F0 int16; F1 uint32; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 uint16; F2 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool; F1 int; F2 uint64 }
type GenS10 struct { F0 int16; F1 uint8; F2 int }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS13 struct { F0 uint64; F1 string; F2 int16 }
type GenS14 struct { F0 int; F1 uint32 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint; F1 string; F2 uint8 }
type GenS18 struct { F0 uint; F1 int16 }
type GenS19 struct { F0 uint32; F1 int8 }
type GenS20 struct { F0 uint8; F1 uint; F2 int16 }
type GenS21 struct { F0 int; F1 string }
type GenS22 struct { F0 int16 }
type GenS23 struct { F0 uint }
type GenS24 struct { F0 bool; F1 uint64; F2 bool }
type GenS25 struct { F0 int16; F1 int32; F2 bool }
type GenS26 struct { F0 uint64; F1 string }
type GenS27 struct { F0 bool; F1 uint8 }
type GenS28 struct { F0 int; F1 uint; F2 uint8 }
type GenS29 struct { F0 int; F1 uint16 }
type GenS30 struct { F0 uint32 }
type GenS31 struct { F0 bool; F1 int8 }
type GenS32 struct { F0 int8; F1 uint; F2 uint8 }
type GenS33 struct { F0 uint64 }
func genFn0(p0 bool) int8 {
    if (14 < 823) {
        v5 := GenS1{F0: 2, F1: 255, F2: -9223372036854775808}
        _ = v5
    } else {
        var v6 uint32 = 65535
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := []int{(-(fi7)), fi7}
            if ((-103 <= 0) && p0) {
                switch fi7 {
                case 128:
                    var v9 int64 = -1
                    _ = v9
                case 32768:
                    if true {
                        v8 = append(v8, 1)
                        if (!false) {
                            var v10 int32 = -2147483648
                            _ = v10
                        }
                    }
                default:
                    var v11 float64 = 1e10
                    _ = v11
                }
                var v12 string = "foo bar"
                v6 += (v6 / ((678) | 1))
                _ = v12
            } else {
                if true {
                    if ((0.0 / (-2.25 + 0.1)) <= (-2.25 / 1.5)) {
                        if true {
                            var v13 int64 = -370
                            _ = v13
                        }
                    } else {
                        p0 = (0.0 < float32((float64(0.0) * 1e10)))
                        p0 = (73 < 25)
                    }
                    p0 = (p0 || true)
                    v14 := GenS1{F0: 42, F1: 42, F2: fi7}
                    _ = v14
                }
                for fi15 := 0; fi15 < 6; fi15++ {
                    if (false && false) {
                        v6 += v6
                    } else {
                        if len(v8) > 0 {
                            v8[0] = fi7
                        }
                    }
                    switch "foo bar" {
                    case "foo bar":
                        switch fi15 {
                        case 128:
                            switch 9223372036854775807 {
                            case 752:
                                var v16 int64 = 9223372036854775807
                                var v17 uint16 = uint16(v6)
                                _ = v16
                                _ = v17
                            case 0:
                                fmt.Println(fi15)
                            default:
                                if (1 <= -1) {
                                    if true {
                                        switch (fi7 - fi7) {
                                        case -1:
                                            switch "a" {
                                            case "":
                                                var v18 uint64 = 4294967295
                                                _ = v18
                                            case "hello":
                                                if p0 {
                                                    if (32767 > 1) {
                                                        v19 := GenS1{F0: 42, F1: 4294967295, F2: 329}
                                                        if (((42 <= 42) && true) || p0) {
                                                            _ = len(v8)
                                                            var v20 bool = true
                                                            _ = v20
                                                        } else {
                                                            v8 = append(v8, 130)
                                                        }
                                                        fmt.Println(v6, p0, p0)
                                                        _ = v19
                                                    } else {
                                                        if (false || (!p0)) {
                                                            fmt.Println(fi15)
                                                        } else {
                                                            p0 = (947 != 65535)
                                                            v21 := GenS1{F0: 606, F1: 15, F2: (fi7 ^ fi7)}
                                                            p0 = p0
                                                            v8 = append(v8, (^(fi7)))
                                                            switch (^(fi7)) {
                                                            case 32768:
                                                                v6 += v6
                                                            case 128:
                                                                if len(v8) > 0 {
                                                                    v8[0] = 2147483647
                                                                }
                                                            case -1:
                                                                v22 := GenS2{F0: 0, F1: v6, F2: (!(524 != uint16(v6)))}
                                                                _ = v22
                                                            }
                                                            _ = v21
                                                        }
                                                        switch (fi7 & fi15) {
                                                        case 128:
                                                            if (p0 || (!(-1 == 771))) {
                                                                v6 += 4294967295
                                                                if (683 == uint(v6)) {
                                                                    p0 = p0
                                                                    v6 -= (^(v6))
                                                                } else {
                                                                    switch ("hello" + "hello") {
                                                                    case "\n\t":
                                                                        if true {
                                                                            if p0 {
                                                                                v6 += v6
                                                                                var v23 int = (fi7 ^ fi15)
                                                                                _ = v23
                                                                            } else {
                                                                                if p0 {
                                                                                    switch 792 {
                                                                                    case 42:
                                                                                        if (2 < 128) {
                                                                                            var v24 float32 = 1e10
                                                                                            var v25 int32 = 2147483647
                                                                                            if (!false) {
                                                                                                _ = len(v8)
                                                                                            } else {
                                                                                                if false {
                                                                                                    v6 = (v6 << ((489) & 63))
                                                                                                } else {
                                                                                                    if p0 {
                                                                                                        if len(v8) > 0 {
                                                                                                            v8[0] = 2147483647
                                                                                                        }
                                                                                                        if len(v8) >= 2 {
                                                                                                            v8 = v8[1:len(v8)]
                                                                                                        }
                                                                                                    }
                                                                                                    if ((1 >= v25) || (v6 > (v6 * v6))) {
                                                                                                        if true {
                                                                                                            if len(v8) > 0 {
                                                                                                                v8[0] = (fi7 | 787)
                                                                                                            }
                                                                                                            if len(v8) > 0 {
                                                                                                                v8[0] = (fi7 / (((fi15 - (fi7 & fi15))) | 1))
                                                                                                            }
                                                                                                            fmt.Println(v24, fi15)
                                                                                                            v26 := GenS3{F0: 847}
                                                                                                            var v27 uint32 = v6
                                                                                                            _ = v26
                                                                                                            _ = v27
                                                                                                        }
                                                                                                        var v28 float32 = 0.1
                                                                                                        var v29 string = ("\n\t" + ("\n\t" + "\n\t"))
                                                                                                        p0 = (!false)
                                                                                                        if (p0 && (uint64(v6) > 402)) {
                                                                                                            if p0 {
                                                                                                                var v30 uint16 = 2
                                                                                                                v6 -= (^(v6))
                                                                                                                var v31 bool = (p0 && p0)
                                                                                                                _ = v30
                                                                                                                _ = v31
                                                                                                            } else {
                                                                                                                _ = len(v8)
                                                                                                                var v32 int = -259
                                                                                                                _ = len(v8)
                                                                                                                _ = v32
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v33 uint16 = 1
                                                                                                            v29 = ""
                                                                                                            v34 := GenS3{F0: -32768}
                                                                                                            if p0 {
                                                                                                                _ = len(v8)
                                                                                                                switch "foo bar" {
                                                                                                                case "foo bar":
                                                                                                                    var v35 int16 = 0
                                                                                                                    switch (v29 + ("" + "hello")) {
                                                                                                                    case "a":
                                                                                                                        switch (fi7 >> ((18446744073709551615) & 63)) {
                                                                                                                        case 32767:
                                                                                                                            p0 = ((1 == -518) || (!p0))
                                                                                                                        case 671:
                                                                                                                            var v36 int = fi15
                                                                                                                            _ = v36
                                                                                                                        }
                                                                                                                    case "\n\t":
                                                                                                                        var v37 uint16 = v33
                                                                                                                        _ = v37
                                                                                                                    case "hello":
                                                                                                                        if (((!(!false)) || true) || (p0 && p0)) {
                                                                                                                            switch fi7 {
                                                                                                                            case -2147483648:
                                                                                                                                if p0 {
                                                                                                                                    var v38 int32 = (v25 | (-(v25)))
                                                                                                                                    if len(v8) > 0 {
                                                                                                                                        v8[0] = fi7
                                                                                                                                    }
                                                                                                                                    if p0 {
                                                                                                                                        v39 := GenS1{F0: (v33 ^ (v33 - v33)), F1: 1, F2: (fi15 % (((^(fi7))) | 1))}
                                                                                                                                        if p0 {
                                                                                                                                            if len(v8) > 0 {
                                                                                                                                                v8[0] = -899
                                                                                                                                            }
                                                                                                                                            v40 := GenS2{F0: v35, F1: (^(v6)), F2: (!p0)}
                                                                                                                                            _ = v40
                                                                                                                                        } else {
                                                                                                                                            v41 := GenS4{F0: 0}
                                                                                                                                            p0 = (!((v35 * 963) > (v35 >> ((625) & 63))))
                                                                                                                                            _ = v41
                                                                                                                                        }
                                                                                                                                        switch fi15 {
                                                                                                                                        case 128:
                                                                                                                                            v29 = v29
                                                                                                                                        default:
                                                                                                                                            v42 := GenS5{F0: (v25 | int32(v35)), F1: (v33 % ((v33) | 1)), F2: v35}
                                                                                                                                            _ = v42
                                                                                                                                        }
                                                                                                                                        _ = v39
                                                                                                                                    }
                                                                                                                                    _ = v38
                                                                                                                                } else {
                                                                                                                                    var v43 uint8 = 2
                                                                                                                                    var v44 int64 = 1
                                                                                                                                    _ = v43
                                                                                                                                    _ = v44
                                                                                                                                }
                                                                                                                                switch fi7 {
                                                                                                                                case 32768:
                                                                                                                                    switch v29 {
                                                                                                                                    case "\n\t":
                                                                                                                                        switch fi7 {
                                                                                                                                        case 0:
                                                                                                                                            switch v29 {
                                                                                                                                            case "a":
                                                                                                                                                fmt.Println(fi15)
                                                                                                                                            default:
                                                                                                                                                v24 += float32(0.1)
                                                                                                                                            }
                                                                                                                                            var v45 int64 = 968
                                                                                                                                            v46 := GenS5{F0: v25, F1: (v33 - v33), F2: v35}
                                                                                                                                            _ = v45
                                                                                                                                            _ = v46
                                                                                                                                        case 128:
                                                                                                                                            v25 = v25
                                                                                                                                        default:
                                                                                                                                            v35 = 1
                                                                                                                                        }
                                                                                                                                        var v47 string = ("a" + ((((v29 + (v29 + (("a" + "\n\t") + v29))) + v29) + (v29 + "a")) + "hello"))
                                                                                                                                        _ = v47
                                                                                                                                    }
                                                                                                                                case 127:
                                                                                                                                    v48 := GenS1{F0: 65535, F1: 18446744073709551615, F2: (fi7 % (((-(fi15))) | 1))}
                                                                                                                                    _ = v48
                                                                                                                                }
                                                                                                                                v49 := GenS6{F0: v33}
                                                                                                                                _ = v49
                                                                                                                            case 4294967295:
                                                                                                                                var v50 uint64 = 4294967295
                                                                                                                                _ = v50
                                                                                                                            case 8:
                                                                                                                                var v51 uint64 = 18446744073709551615
                                                                                                                                _ = v51
                                                                                                                            }
                                                                                                                            v35 = v35
                                                                                                                            switch ((v29 + "foo bar") + (("\n\t" + "hello") + v29)) {
                                                                                                                            case "\n\t":
                                                                                                                                v8 = append(v8, 553)
                                                                                                                            case "foo bar":
                                                                                                                                var v52 uint16 = v33
                                                                                                                                _ = v52
                                                                                                                            case "hello":
                                                                                                                                v8 = append(v8, fi7)
                                                                                                                                v35 = 42
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v35 = (v35 % ((1) | 1))
                                                                                                                            if len(v8) >= 2 {
                                                                                                                                v8 = v8[1:len(v8)]
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch ((("a" + v29) + v29) + ((v29 + "") + (v29 + v29))) {
                                                                                                                    case "foo bar":
                                                                                                                        switch (v29 + (v29 + v29)) {
                                                                                                                        case "a":
                                                                                                                            var v53 int = (fi7 << ((42) & 63))
                                                                                                                            _ = v53
                                                                                                                        case "hello":
                                                                                                                            if len(v8) >= 2 {
                                                                                                                                v8 = v8[1:len(v8)]
                                                                                                                            }
                                                                                                                        case "":
                                                                                                                            if ((v25 <= 2147483647) && p0) {
                                                                                                                                p0 = p0
                                                                                                                                if len(v8) > 0 {
                                                                                                                                    v8[0] = (fi15 / ((2147483647) | 1))
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch v29 {
                                                                                                                    case "":
                                                                                                                        v8 = append(v8, (fi7 ^ (fi15 % ((fi7) | 1))))
                                                                                                                    case "a":
                                                                                                                        p0 = p0
                                                                                                                    }
                                                                                                                    var v54 int8 = 0
                                                                                                                    _ = v35
                                                                                                                    _ = v54
                                                                                                                case "a":
                                                                                                                    if (100 <= 0) {
                                                                                                                        switch int(v6) {
                                                                                                                        case -9223372036854775808:
                                                                                                                            v28 = ((float32(1.5) + (v24 - v28)) + (0.1 * v28))
                                                                                                                        case -704:
                                                                                                                            v28 = 0.1
                                                                                                                        default:
                                                                                                                            var v55 bool = false
                                                                                                                            _ = v55
                                                                                                                        }
                                                                                                                        switch (fi15 * fi7) {
                                                                                                                        case 0:
                                                                                                                            _ = len(v8)
                                                                                                                            var v56 bool = true
                                                                                                                            _ = v56
                                                                                                                        default:
                                                                                                                            if p0 {
                                                                                                                                v24 += float32(float64(((float32(0.0) - (v24 / ((((v28 - (-2.25 / 3.14159)) - (float32(3.14159) * v24)) / 1.5) + 3.14159))) + (1.5 / v24))))
                                                                                                                                fmt.Println(fi7)
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v8 = append(v8, fi15)
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v33
                                                                                                            _ = v34
                                                                                                        }
                                                                                                        v57 := GenS4{F0: 115}
                                                                                                        _ = v28
                                                                                                        _ = v29
                                                                                                        _ = v57
                                                                                                    } else {
                                                                                                        var v58 uint8 = 2
                                                                                                        if len(v8) >= 2 {
                                                                                                            v8 = v8[1:len(v8)]
                                                                                                        }
                                                                                                        if p0 {
                                                                                                            v59 := GenS2{F0: 42, F1: 1, F2: (754 > 4294967295)}
                                                                                                            _ = v59
                                                                                                        } else {
                                                                                                            v58 = 221
                                                                                                            v60 := GenS4{F0: -50}
                                                                                                            _ = v60
                                                                                                        }
                                                                                                        if p0 {
                                                                                                            v24 = ((float32(1e10) / ((v24 + (v24 - v24)) / (v24 * float32(3.14159)))) / v24)
                                                                                                            if true {
                                                                                                                var v61 int8 = 127
                                                                                                                var v62 int = (fi7 >> ((2) & 63))
                                                                                                                _ = v61
                                                                                                                _ = v62
                                                                                                            } else {
                                                                                                                if false {
                                                                                                                    var v63 uint64 = 423
                                                                                                                    fmt.Println(v58, v58)
                                                                                                                    _ = v63
                                                                                                                } else {
                                                                                                                    fmt.Println(v24, v58)
                                                                                                                    var v64 uint16 = 65535
                                                                                                                    var v65 uint32 = 553
                                                                                                                    _ = v64
                                                                                                                    _ = v65
                                                                                                                }
                                                                                                                v58 = v58
                                                                                                                if len(v8) >= 2 {
                                                                                                                    v8 = v8[1:len(v8)]
                                                                                                                }
                                                                                                                if (fi15 > (^(fi15))) {
                                                                                                                    v66 := GenS4{F0: 100}
                                                                                                                    var v67 uint = 39
                                                                                                                    switch -9223372036854775808 {
                                                                                                                    case 9223372036854775807:
                                                                                                                        fmt.Println(v58)
                                                                                                                    case 4294967295:
                                                                                                                        if (p0 || ((v58 > (v58 - 42)) && (1 > -384))) {
                                                                                                                            var v68 uint = (v67 + 1)
                                                                                                                            switch "a" {
                                                                                                                            case "\n\t":
                                                                                                                                if (42 < -73) {
                                                                                                                                    v69 := GenS7{F0: fi15, F1: v25}
                                                                                                                                    v70 := GenS8{F0: v58}
                                                                                                                                    if len(v8) >= 2 {
                                                                                                                                        v8 = v8[1:len(v8)]
                                                                                                                                    }
                                                                                                                                    v58 -= v58
                                                                                                                                    var v71 uint64 = 2
                                                                                                                                    _ = v69
                                                                                                                                    _ = v70
                                                                                                                                    _ = v71
                                                                                                                                } else {
                                                                                                                                    var v72 int8 = -1
                                                                                                                                    _ = v72
                                                                                                                                }
                                                                                                                            }
                                                                                                                            if len(v8) > 0 {
                                                                                                                                v8[0] = 32768
                                                                                                                            }
                                                                                                                            _ = v68
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v58 -= (^(v58))
                                                                                                                    }
                                                                                                                    _ = v66
                                                                                                                    _ = v67
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        v73 := GenS9{F0: p0, F1: (fi7 | fi7), F2: 1}
                                                                                                        _ = v58
                                                                                                        _ = v73
                                                                                                    }
                                                                                                    if len(v8) >= 2 {
                                                                                                        v8 = v8[1:len(v8)]
                                                                                                    }
                                                                                                    var v74 int32 = v25
                                                                                                    _ = v74
                                                                                                }
                                                                                            }
                                                                                            _ = v24
                                                                                            _ = v25
                                                                                        } else {
                                                                                            if len(v8) >= 2 {
                                                                                                v8 = v8[1:len(v8)]
                                                                                            }
                                                                                        }
                                                                                    case 1:
                                                                                        var v75 uint32 = 357
                                                                                        _ = v75
                                                                                    }
                                                                                    switch "foo bar" {
                                                                                    case "a":
                                                                                        switch "foo bar" {
                                                                                        case "hello":
                                                                                            var v76 bool = (-121 < 0)
                                                                                            _ = v76
                                                                                        case "foo bar":
                                                                                            fmt.Println(p0)
                                                                                        default:
                                                                                            var v77 uint = 2
                                                                                            _ = v77
                                                                                        }
                                                                                    case "\n\t":
                                                                                        if p0 {
                                                                                            if (-9223372036854775808 != 9223372036854775807) {
                                                                                                if len(v8) > 0 {
                                                                                                    v8[0] = 32767
                                                                                                }
                                                                                            }
                                                                                            switch "hello" {
                                                                                            case "hello":
                                                                                                v78 := GenS10{F0: 1, F1: 72, F2: -742}
                                                                                                _ = v78
                                                                                            case "a":
                                                                                                if p0 {
                                                                                                    var v79 float32 = float32(0.0)
                                                                                                    _ = v79
                                                                                                } else {
                                                                                                    v80 := GenS9{F0: true, F1: (fi15 ^ -2147483648), F2: 1}
                                                                                                    _ = v80
                                                                                                }
                                                                                            case "":
                                                                                                if (-820 <= int16(v6)) {
                                                                                                    _ = len(v8)
                                                                                                    if len(v8) >= 2 {
                                                                                                        v8 = v8[1:len(v8)]
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                switch ("" + "hello") {
                                                                                                case "a":
                                                                                                    fmt.Println(p0, fi15, v6)
                                                                                                case "\n\t":
                                                                                                    switch "foo bar" {
                                                                                                    case "a":
                                                                                                        fmt.Println(p0)
                                                                                                    default:
                                                                                                        if p0 {
                                                                                                            v81 := GenS2{F0: 433, F1: 1, F2: ((-768 < -2147483648) && false)}
                                                                                                            v6 -= 42
                                                                                                            switch "hello" {
                                                                                                            case "":
                                                                                                                v82 := GenS8{F0: 35}
                                                                                                                _ = v82
                                                                                                            }
                                                                                                            _ = v81
                                                                                                        } else {
                                                                                                            if len(v8) > 0 {
                                                                                                                v8[0] = fi15
                                                                                                            }
                                                                                                            v6 += v6
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    if ((-456 < 2147483647) && p0) {
                                                                                                        p0 = (781 < 193)
                                                                                                        var v83 uint = 255
                                                                                                        v83 -= (v83 * 255)
                                                                                                        _ = v83
                                                                                                    } else {
                                                                                                        p0 = (!p0)
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if len(v8) > 0 {
                                                                                                v8[0] = (-(fi7))
                                                                                            }
                                                                                            _ = len(v8)
                                                                                            var v84 bool = false
                                                                                            _ = v84
                                                                                        }
                                                                                    default:
                                                                                        v6 += 4294967295
                                                                                    }
                                                                                    switch ("" + "hello") {
                                                                                    case "a":
                                                                                        _ = len(v8)
                                                                                    default:
                                                                                        p0 = p0
                                                                                    }
                                                                                } else {
                                                                                    if true {
                                                                                        var v85 int8 = 42
                                                                                        if p0 {
                                                                                            v86 := GenS3{F0: -1}
                                                                                            if (!(0 >= 42)) {
                                                                                                var v87 uint16 = 0
                                                                                                fmt.Println(fi15)
                                                                                                v6 -= v6
                                                                                                _ = v87
                                                                                            } else {
                                                                                                _ = len(v8)
                                                                                                if (!(3.14159 <= 0.0)) {
                                                                                                    if (!((fi15 >= fi15) || (p0 && (4294967295 > 18446744073709551615)))) {
                                                                                                        if (!p0) {
                                                                                                            fmt.Println(v6, p0, v6)
                                                                                                            if (321 > 4294967295) {
                                                                                                                if p0 {
                                                                                                                    fmt.Println(v6, fi7)
                                                                                                                } else {
                                                                                                                    if (65535 <= 42) {
                                                                                                                        var v88 float32 = 0.0
                                                                                                                        v85 -= -1
                                                                                                                        _ = v88
                                                                                                                    }
                                                                                                                }
                                                                                                                if (128 == 127) {
                                                                                                                    p0 = true
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            v6 += v6
                                                                                                            switch (fi15 * 32768) {
                                                                                                            case 1:
                                                                                                                var v89 int16 = 42
                                                                                                                _ = v89
                                                                                                            case 451:
                                                                                                                var v90 float64 = ((1e10 - float64(((3.14159 * (1.5 * -2.25)) * (0.1 + 0.1)))) - 0.1)
                                                                                                                _ = v90
                                                                                                            case 127:
                                                                                                                var v91 float64 = 0.1
                                                                                                                switch 816 {
                                                                                                                case -2147483648:
                                                                                                                    v91 += 1.5
                                                                                                                case -906:
                                                                                                                    var v92 uint64 = 42
                                                                                                                    _ = v92
                                                                                                                }
                                                                                                                if (4294967295 > 2) {
                                                                                                                    if p0 {
                                                                                                                        if len(v8) > 0 {
                                                                                                                            v8[0] = 990
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v93 float32 = (0.0 + 3.14159)
                                                                                                                    fmt.Println(v85, v6, fi7)
                                                                                                                    _ = v93
                                                                                                                } else {
                                                                                                                    var v94 float64 = float64(0.0)
                                                                                                                    if len(v8) >= 2 {
                                                                                                                        v8 = v8[1:len(v8)]
                                                                                                                    }
                                                                                                                    _ = v94
                                                                                                                }
                                                                                                                _ = v91
                                                                                                            }
                                                                                                            if true {
                                                                                                                if p0 {
                                                                                                                    if p0 {
                                                                                                                        var v95 float32 = 0.1
                                                                                                                        fmt.Println(fi15)
                                                                                                                        _ = v95
                                                                                                                    }
                                                                                                                    v96 := GenS3{F0: -118}
                                                                                                                    _ = v96
                                                                                                                } else {
                                                                                                                    switch "foo bar" {
                                                                                                                    case "foo bar":
                                                                                                                        var v97 uint32 = 65535
                                                                                                                        _ = v97
                                                                                                                    }
                                                                                                                    if (9223372036854775807 < 91) {
                                                                                                                        v6 = (v6 + 2)
                                                                                                                        if true {
                                                                                                                            var v98 uint = 18446744073709551615
                                                                                                                            if (42 != -9223372036854775808) {
                                                                                                                                var v99 bool = true
                                                                                                                                _ = v99
                                                                                                                            }
                                                                                                                            _ = v98
                                                                                                                        } else {
                                                                                                                            if (65535 != 1) {
                                                                                                                                switch ("" + ("" + "")) {
                                                                                                                                case "":
                                                                                                                                    if (-2.25 >= 0.0) {
                                                                                                                                        var v100 int = (^(fi7))
                                                                                                                                        _ = v100
                                                                                                                                    }
                                                                                                                                case "a":
                                                                                                                                    v85 += v85
                                                                                                                                case "hello":
                                                                                                                                    v85 += (v85 - int8(fi15))
                                                                                                                                }
                                                                                                                                if len(v8) > 0 {
                                                                                                                                    v8[0] = (fi7 & fi15)
                                                                                                                                }
                                                                                                                                if (p0 || (2147483648 == v6)) {
                                                                                                                                    if (false && p0) {
                                                                                                                                        if (true || (p0 && p0)) {
                                                                                                                                            if len(v8) > 0 {
                                                                                                                                                v8[0] = fi15
                                                                                                                                            }
                                                                                                                                            var v101 uint32 = (v6 - 730)
                                                                                                                                            var v102 int32 = int32(v85)
                                                                                                                                            var v103 int = fi7
                                                                                                                                            _ = v101
                                                                                                                                            _ = v102
                                                                                                                                            _ = v103
                                                                                                                                        }
                                                                                                                                        if p0 {
                                                                                                                                            _ = len(v8)
                                                                                                                                        }
                                                                                                                                        var v104 int8 = 0
                                                                                                                                        _ = v104
                                                                                                                                    } else {
                                                                                                                                        if (((1.5 + ((float64(1e10) + -2.25) / -2.25)) / (float64(0.1) + 0.0)) <= 3.14159) {
                                                                                                                                            switch (^(fi15)) {
                                                                                                                                            case 32768:
                                                                                                                                                switch fi7 {
                                                                                                                                                case 42:
                                                                                                                                                    if ((2 != 1) || p0) {
                                                                                                                                                        _ = len(v8)
                                                                                                                                                    } else {
                                                                                                                                                        if false {
                                                                                                                                                            v85 = v85
                                                                                                                                                            v105 := GenS11{F0: (fi15 ^ fi15)}
                                                                                                                                                            _ = v105
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                case 9223372036854775807:
                                                                                                                                                    v106 := GenS12{F0: 1, F1: v6, F2: (v85 ^ (-(v85)))}
                                                                                                                                                    _ = v106
                                                                                                                                                case 4294967295:
                                                                                                                                                    v107 := GenS9{F0: (true && true), F1: (fi15 % ((925) | 1)), F2: 0}
                                                                                                                                                    _ = v107
                                                                                                                                                }
                                                                                                                                            case 4294967295:
                                                                                                                                                var v108 int8 = -14
                                                                                                                                                _ = v108
                                                                                                                                            default:
                                                                                                                                                var v109 int = 32767
                                                                                                                                                _ = v109
                                                                                                                                            }
                                                                                                                                            var v110 int64 = -945
                                                                                                                                            _ = v110
                                                                                                                                        }
                                                                                                                                        p0 = true
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    if p0 {
                                                                                                                                        var v111 uint32 = (v6 & (v6 / ((v6) | 1)))
                                                                                                                                        _ = v111
                                                                                                                                    }
                                                                                                                                    v112 := GenS13{F0: 4294967295, F1: "a", F2: -32768}
                                                                                                                                    v113 := GenS14{F0: fi7, F1: (v6 & (v6 * (v6 ^ v6)))}
                                                                                                                                    _ = v112
                                                                                                                                    _ = v113
                                                                                                                                }
                                                                                                                                v114 := GenS15{F0: int64(fi7)}
                                                                                                                                var v115 uint16 = 0
                                                                                                                                _ = v114
                                                                                                                                _ = v115
                                                                                                                            }
                                                                                                                        }
                                                                                                                        var v116 float64 = -2.25
                                                                                                                        var v117 int16 = 1000
                                                                                                                        if ((fi15 << ((275) & 63)) < (fi7 * fi7)) {
                                                                                                                            if (0 >= 973) {
                                                                                                                                p0 = (2 < 702)
                                                                                                                                v117 += v117
                                                                                                                                v118 := GenS5{F0: 42, F1: 0, F2: (v117 ^ v117)}
                                                                                                                                _ = v118
                                                                                                                            } else {
                                                                                                                                var v119 float32 = 1e10
                                                                                                                                switch (("hello" + "foo bar") + ("\n\t" + ("foo bar" + ("hello" + ((("hello" + "\n\t") + "a") + ("foo bar" + (("foo bar" + "hello") + "\n\t"))))))) {
                                                                                                                                case "foo bar":
                                                                                                                                    var v120 int16 = (v117 << ((180) & 63))
                                                                                                                                    _ = v120
                                                                                                                                case "":
                                                                                                                                    v85 += 0
                                                                                                                                }
                                                                                                                                if (-985 > 42) {
                                                                                                                                    v117 += v117
                                                                                                                                    v121 := GenS13{F0: 810, F1: (("hello" + (("" + "foo bar") + "hello")) + ("\n\t" + (("a" + ("" + "hello")) + ""))), F2: -32768}
                                                                                                                                    if (((0 > 4294967295) || false) || p0) {
                                                                                                                                        var v122 float64 = ((((v116 * v116) + v116) + 0.0) / (float64(1e10) + (v116 + (v116 + (1e10 + v116)))))
                                                                                                                                        _ = v122
                                                                                                                                    } else {
                                                                                                                                        v119 += 0.1
                                                                                                                                    }
                                                                                                                                    _ = v121
                                                                                                                                }
                                                                                                                                _ = v119
                                                                                                                            }
                                                                                                                            _ = len(v8)
                                                                                                                            v8 = append(v8, fi15)
                                                                                                                        }
                                                                                                                        _ = v116
                                                                                                                        _ = v117
                                                                                                                    } else {
                                                                                                                        if p0 {
                                                                                                                            p0 = (p0 || ((1e10 != 1.5) || (!true)))
                                                                                                                            v85 = 0
                                                                                                                        } else {
                                                                                                                            v85 += (v85 ^ v85)
                                                                                                                            v6 += (v6 | (^(v6)))
                                                                                                                            fmt.Println(v6, v85)
                                                                                                                        }
                                                                                                                        p0 = p0
                                                                                                                        if (4294967295 >= 2) {
                                                                                                                            if len(v8) > 0 {
                                                                                                                                v8[0] = (fi15 | fi7)
                                                                                                                            }
                                                                                                                            var v123 bool = p0
                                                                                                                            _ = v123
                                                                                                                        }
                                                                                                                    }
                                                                                                                    p0 = p0
                                                                                                                }
                                                                                                                v124 := GenS16{F0: -355}
                                                                                                                fmt.Println(fi7, v85, v85)
                                                                                                                _ = v124
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    var v125 float32 = 1e10
                                                                                                    switch fi7 {
                                                                                                    case -375:
                                                                                                        v8 = append(v8, (fi15 / ((fi15) | 1)))
                                                                                                    case 2147483647:
                                                                                                        v126 := GenS17{F0: 28, F1: ("" + "foo bar"), F2: 42}
                                                                                                        _ = v126
                                                                                                    }
                                                                                                    _ = v125
                                                                                                } else {
                                                                                                    var v127 float32 = (1.5 * 3.14159)
                                                                                                    fmt.Println(v6)
                                                                                                    _ = v127
                                                                                                }
                                                                                                if len(v8) >= 2 {
                                                                                                    v8 = v8[1:len(v8)]
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v85, fi7)
                                                                                            _ = v86
                                                                                        } else {
                                                                                            switch (("hello" + (("a" + "foo bar") + "foo bar")) + "\n\t") {
                                                                                            case "foo bar":
                                                                                                if (p0 && (!("" != ("" + "a")))) {
                                                                                                    v128 := GenS18{F0: 65535, F1: 1000}
                                                                                                    v129 := GenS19{F0: (v6 ^ (v6 | 784)), F1: (v85 ^ v85)}
                                                                                                    fmt.Println(p0, fi7, fi15)
                                                                                                    switch fi15 {
                                                                                                    case 14:
                                                                                                        var v130 uint8 = 28
                                                                                                        _ = v130
                                                                                                    }
                                                                                                    _ = v128
                                                                                                    _ = v129
                                                                                                } else {
                                                                                                    v6 = v6
                                                                                                }
                                                                                            case "":
                                                                                                var v131 uint64 = 0
                                                                                                _ = v131
                                                                                            case "a":
                                                                                                var v132 int8 = (v85 << ((0) & 63))
                                                                                                _ = v132
                                                                                            default:
                                                                                                var v133 bool = p0
                                                                                                _ = v133
                                                                                            }
                                                                                            if len(v8) >= 2 {
                                                                                                v8 = v8[1:len(v8)]
                                                                                            }
                                                                                            if (p0 && false) {
                                                                                                var v134 bool = (0 < 255)
                                                                                                fmt.Println(v6, fi15)
                                                                                                var v135 int = 419
                                                                                                _ = v134
                                                                                                _ = v135
                                                                                            } else {
                                                                                                fmt.Println(v85)
                                                                                                v8 = append(v8, (fi7 >> ((42) & 63)))
                                                                                            }
                                                                                        }
                                                                                        _ = v85
                                                                                    }
                                                                                    if ((fi15 - (fi7 / ((fi7) | 1))) == fi7) {
                                                                                        var v136 int8 = -40
                                                                                        _ = v136
                                                                                    } else {
                                                                                        var v137 float64 = (0.1 * 0.1)
                                                                                        _ = v137
                                                                                    }
                                                                                    v138 := GenS20{F0: 255, F1: 554, F2: 42}
                                                                                    _ = v138
                                                                                }
                                                                                if len(v8) > 0 {
                                                                                    v8[0] = fi7
                                                                                }
                                                                                var v139 bool = (!false)
                                                                                if (2 > 255) {
                                                                                    var v140 uint64 = 18446744073709551615
                                                                                    var v141 uint8 = 0
                                                                                    switch (fi15 + (fi15 % ((fi15) | 1))) {
                                                                                    case 1:
                                                                                        switch int(v140) {
                                                                                        case 128:
                                                                                            switch ("foo bar" + (("foo bar" + ("a" + "a")) + "")) {
                                                                                            case "\n\t":
                                                                                                if false {
                                                                                                    v139 = ((!((v139 || v139) && false)) || (65535 <= uint(v140)))
                                                                                                    if (((v6 * 65535) < (v6 + v6)) && p0) {
                                                                                                        v141 = (v141 & v141)
                                                                                                    } else {
                                                                                                        fmt.Println(v140, fi15, v140)
                                                                                                    }
                                                                                                    v142 := GenS21{F0: fi7, F1: "\n\t"}
                                                                                                    _ = v142
                                                                                                } else {
                                                                                                    var v143 bool = v139
                                                                                                    switch "a" {
                                                                                                    case "a":
                                                                                                        v144 := GenS9{F0: p0, F1: -544, F2: (v140 % (((v140 - v140)) | 1))}
                                                                                                        _ = v144
                                                                                                    case "":
                                                                                                        if (1 > 415) {
                                                                                                            var v145 uint32 = v6
                                                                                                            if true {
                                                                                                                if false {
                                                                                                                    if v139 {
                                                                                                                        v8 = append(v8, fi7)
                                                                                                                        v143 = (42 > 42)
                                                                                                                        if false {
                                                                                                                            v139 = (1 < -360)
                                                                                                                        } else {
                                                                                                                            var v146 int16 = 854
                                                                                                                            v140 = (v140 % (((v140 / (((v140 * (v140 / (((v140 % ((2) | 1))) | 1)))) | 1))) | 1))
                                                                                                                            var v147 float64 = float64(0.1)
                                                                                                                            _ = v146
                                                                                                                            _ = v147
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v148 uint8 = v141
                                                                                                                        var v149 bool = true
                                                                                                                        if (((("foo bar" + "a") + "foo bar") >= ("foo bar" + ("a" + "hello"))) && (-48 < -1)) {
                                                                                                                            var v150 int16 = 0
                                                                                                                            if v143 {
                                                                                                                                switch ("a" + "hello") {
                                                                                                                                case "\n\t":
                                                                                                                                    v145 = (v145 << ((635) & 63))
                                                                                                                                case "a":
                                                                                                                                    if (3.14159 >= 0.1) {
                                                                                                                                        v143 = false
                                                                                                                                        v151 := GenS17{F0: 675, F1: ((("a" + "") + ((("a" + "a") + "foo bar") + (("a" + (("foo bar" + (("hello" + "\n\t") + "foo bar")) + "")) + ""))) + "foo bar"), F2: v141}
                                                                                                                                        switch fi15 {
                                                                                                                                        case 2147483647:
                                                                                                                                            v152 := GenS22{F0: v150}
                                                                                                                                            _ = v152
                                                                                                                                        }
                                                                                                                                        var v153 string = ("" + ("" + ("foo bar" + "\n\t")))
                                                                                                                                        _ = v151
                                                                                                                                        _ = v153
                                                                                                                                    } else {
                                                                                                                                        fmt.Println(v143, v148, v6)
                                                                                                                                        if p0 {
                                                                                                                                            v8 = append(v8, (fi15 & fi15))
                                                                                                                                        } else {
                                                                                                                                            var v154 int16 = v150
                                                                                                                                            _ = v154
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                fmt.Println(fi15, fi7, v140)
                                                                                                                            }
                                                                                                                            if v143 {
                                                                                                                                var v155 float64 = 0.1
                                                                                                                                switch ("" + "\n\t") {
                                                                                                                                case "hello":
                                                                                                                                    v150 += (v150 * v150)
                                                                                                                                default:
                                                                                                                                    if false {
                                                                                                                                        if len(v8) > 0 {
                                                                                                                                            v8[0] = 32768
                                                                                                                                        }
                                                                                                                                        switch (fi15 | 32767) {
                                                                                                                                        case 127:
                                                                                                                                            fmt.Println(v139, v150, v141)
                                                                                                                                        case 306:
                                                                                                                                            v141 = (v141 / ((v148) | 1))
                                                                                                                                        }
                                                                                                                                        _ = len(v8)
                                                                                                                                    } else {
                                                                                                                                        if len(v8) >= 2 {
                                                                                                                                            v8 = v8[1:len(v8)]
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                var v156 int64 = 0
                                                                                                                                _ = v155
                                                                                                                                _ = v156
                                                                                                                            } else {
                                                                                                                                switch (fi15 ^ 32767) {
                                                                                                                                case 32767:
                                                                                                                                    var v157 bool = true
                                                                                                                                    v143 = (true && ((728 <= 18446744073709551615) && false))
                                                                                                                                    v149 = (42 != 946)
                                                                                                                                    _ = v157
                                                                                                                                case 32768:
                                                                                                                                    fmt.Println(v141, fi15, v148)
                                                                                                                                default:
                                                                                                                                    _ = len(v8)
                                                                                                                                }
                                                                                                                                var v158 float64 = 1.5
                                                                                                                                _ = v158
                                                                                                                            }
                                                                                                                            _ = v150
                                                                                                                        } else {
                                                                                                                            v159 := GenS15{F0: -648}
                                                                                                                            _ = v159
                                                                                                                        }
                                                                                                                        p0 = v143
                                                                                                                        _ = v148
                                                                                                                        _ = v149
                                                                                                                    }
                                                                                                                    var v160 string = "foo bar"
                                                                                                                    _ = v160
                                                                                                                } else {
                                                                                                                    fmt.Println(v145, p0)
                                                                                                                    var v161 uint32 = (v6 + uint32(v140))
                                                                                                                    _ = v161
                                                                                                                }
                                                                                                                v162 := GenS23{F0: 18446744073709551615}
                                                                                                                var v163 int64 = -149
                                                                                                                _ = v162
                                                                                                                _ = v163
                                                                                                            } else {
                                                                                                                if (1e10 < ((float32(float64(1e10)) * ((float32(float64(3.14159)) * 0.1) * ((1.5 - 1.5) + 3.14159))) * 1.5)) {
                                                                                                                    var v164 uint8 = (^(v141))
                                                                                                                    _ = v164
                                                                                                                } else {
                                                                                                                    if ((float64(0.1) >= (1.5 / 1.5)) || (((1.5 / float32(1.5)) / 1.0) >= -2.25)) {
                                                                                                                        v143 = v139
                                                                                                                        switch fi15 {
                                                                                                                        case 42:
                                                                                                                            v165 := GenS23{F0: 4294967295}
                                                                                                                            _ = v165
                                                                                                                        }
                                                                                                                        v166 := GenS24{F0: (v139 || v143), F1: v140, F2: p0}
                                                                                                                        if (uint16(v141) <= 905) {
                                                                                                                            var v167 uint32 = 42
                                                                                                                            _ = v167
                                                                                                                        } else {
                                                                                                                            v140 = 459
                                                                                                                            v168 := GenS25{F0: int16(fi7), F1: 286, F2: v139}
                                                                                                                            v140 = v140
                                                                                                                            _ = v168
                                                                                                                        }
                                                                                                                        fmt.Println(v140)
                                                                                                                        _ = v166
                                                                                                                    }
                                                                                                                    if (fi7 >= 127) {
                                                                                                                        v6 -= v6
                                                                                                                    } else {
                                                                                                                        v140 = (v140 * 2)
                                                                                                                    }
                                                                                                                }
                                                                                                                v143 = (-1 < -603)
                                                                                                                var v169 int = -794
                                                                                                                switch v169 {
                                                                                                                case -565:
                                                                                                                    v169 = (fi7 % ((fi7) | 1))
                                                                                                                case 128:
                                                                                                                    fmt.Println(v6, v139, v6)
                                                                                                                }
                                                                                                                _ = v169
                                                                                                            }
                                                                                                            switch fi15 {
                                                                                                            case 9223372036854775807:
                                                                                                                v140 -= 1
                                                                                                            case 167:
                                                                                                                var v170 uint8 = v141
                                                                                                                _ = v170
                                                                                                            }
                                                                                                            v8 = append(v8, (-(fi7)))
                                                                                                            v171 := GenS17{F0: 274, F1: "", F2: (v141 ^ 127)}
                                                                                                            if p0 {
                                                                                                                var v172 uint32 = v145
                                                                                                                fmt.Println(p0)
                                                                                                                var v173 float32 = (-2.25 * -2.25)
                                                                                                                _ = v172
                                                                                                                _ = v173
                                                                                                            }
                                                                                                            _ = v145
                                                                                                            _ = v171
                                                                                                        } else {
                                                                                                            v6 = v6
                                                                                                            var v174 int64 = int64(v6)
                                                                                                            _ = v174
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        if (v139 || (-1 >= -1)) {
                                                                                                            v143 = ((!(v139 || ((v141 % (((v141 % ((255) | 1))) | 1)) == 138))) && (42 != 1))
                                                                                                            v143 = (1 >= 42)
                                                                                                            switch fi15 {
                                                                                                            case -620:
                                                                                                                var v175 int = fi7
                                                                                                                _ = v175
                                                                                                            default:
                                                                                                                fmt.Println(fi7)
                                                                                                            }
                                                                                                        } else {
                                                                                                            v143 = (1 != 510)
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v143
                                                                                                }
                                                                                            default:
                                                                                                var v176 string = "hello"
                                                                                                _ = v176
                                                                                            }
                                                                                        case 734:
                                                                                            fmt.Println(v140)
                                                                                            if p0 {
                                                                                                switch (fi7 / (((fi15 & fi15)) | 1)) {
                                                                                                case -270:
                                                                                                    v177 := GenS23{F0: 271}
                                                                                                    _ = v177
                                                                                                case -194:
                                                                                                    _ = len(v8)
                                                                                                    if len(v8) > 0 {
                                                                                                        v8[0] = fi15
                                                                                                    }
                                                                                                    fmt.Println(p0, v139, v140)
                                                                                                    var v178 int64 = 1
                                                                                                    switch "foo bar" {
                                                                                                    case "\n\t":
                                                                                                        v8 = append(v8, -9223372036854775808)
                                                                                                    case "":
                                                                                                        switch "\n\t" {
                                                                                                        case "":
                                                                                                            v8 = append(v8, -679)
                                                                                                        case "a":
                                                                                                            v179 := GenS21{F0: fi7, F1: "hello"}
                                                                                                            _ = v179
                                                                                                        case "\n\t":
                                                                                                            var v180 int32 = 42
                                                                                                            _ = v180
                                                                                                        }
                                                                                                    default:
                                                                                                        v6 += 1
                                                                                                    }
                                                                                                    _ = v178
                                                                                                case -223:
                                                                                                    v140 = (v140 | (v140 - v140))
                                                                                                default:
                                                                                                    _ = len(v8)
                                                                                                }
                                                                                            } else {
                                                                                                if (42 != -32768) {
                                                                                                    var v181 uint8 = v141
                                                                                                    v141 = 0
                                                                                                    if len(v8) > 0 {
                                                                                                        v8[0] = fi7
                                                                                                    }
                                                                                                    v182 := GenS23{F0: 2}
                                                                                                    v141 -= (v181 * (^(v141)))
                                                                                                    _ = v181
                                                                                                    _ = v182
                                                                                                } else {
                                                                                                    _ = len(v8)
                                                                                                    _ = len(v8)
                                                                                                    var v183 float32 = 3.14159
                                                                                                    _ = v183
                                                                                                }
                                                                                                switch "foo bar" {
                                                                                                case "a":
                                                                                                    fmt.Println(v141)
                                                                                                case "":
                                                                                                    v139 = (!v139)
                                                                                                case "foo bar":
                                                                                                    var v184 uint32 = v6
                                                                                                    _ = v184
                                                                                                }
                                                                                                if len(v8) > 0 {
                                                                                                    v8[0] = 568
                                                                                                }
                                                                                                v140 = (v140 + v140)
                                                                                                v185 := GenS18{F0: 0, F1: 1000}
                                                                                                _ = v185
                                                                                            }
                                                                                            v8 = append(v8, fi15)
                                                                                        case 0:
                                                                                            v139 = ("a" >= "foo bar")
                                                                                        }
                                                                                    }
                                                                                    v140 = v140
                                                                                    v186 := GenS23{F0: 197}
                                                                                    _ = v140
                                                                                    _ = v141
                                                                                    _ = v186
                                                                                }
                                                                                p0 = (uint(v6) > uint(v6))
                                                                                _ = v139
                                                                            }
                                                                            var v187 bool = (0 < (v6 | (v6 ^ (v6 * v6))))
                                                                            v188 := GenS24{F0: true, F1: 295, F2: (0 <= 9223372036854775807)}
                                                                            v189 := GenS25{F0: 32767, F1: 2147483647, F2: (float32((0.1 / (((1.5 * (1.5 + ((0.0 * 3.14159) - 0.1))) - -2.25) / (0.1 - 0.0)))) < 0.1)}
                                                                            var v190 uint16 = 32767
                                                                            _ = v187
                                                                            _ = v188
                                                                            _ = v189
                                                                            _ = v190
                                                                        } else {
                                                                            if (p0 || (p0 || p0)) {
                                                                                fmt.Println(fi7)
                                                                                fmt.Println(fi15, p0, v6)
                                                                                var v191 uint = 687
                                                                                _ = v191
                                                                            } else {
                                                                                v192 := GenS12{F0: 1, F1: (v6 + v6), F2: 42}
                                                                                _ = v192
                                                                            }
                                                                            v193 := GenS9{F0: (168 < uint8(v6)), F1: (^(fi15)), F2: 641}
                                                                            v6 -= (v6 - (v6 << ((1) & 63)))
                                                                            _ = v193
                                                                        }
                                                                    default:
                                                                        var v194 bool = (200 > 812)
                                                                        _ = v194
                                                                    }
                                                                    v6 += (v6 - v6)
                                                                    switch "hello" {
                                                                    case "a":
                                                                        fmt.Println(fi15)
                                                                    case "":
                                                                        v195 := GenS26{F0: 0, F1: "a"}
                                                                        _ = v195
                                                                    }
                                                                }
                                                            }
                                                        case 42:
                                                            p0 = (754 == 159)
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                var v196 bool = p0
                                                _ = v196
                                            }
                                        case 479:
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                        default:
                                            var v197 int64 = -1
                                            _ = v197
                                        }
                                        if len(v8) > 0 {
                                            v8[0] = fi7
                                        }
                                    }
                                } else {
                                    if (-664 == 753) {
                                        if ((fi7 > fi15) || (((1e10 * 3.14159) <= float32((0.1 * 0.1))) && p0)) {
                                            p0 = ((("" != "hello") || p0) && p0)
                                            v198 := GenS27{F0: (!(p0 && (p0 || (111 != 32767)))), F1: 128}
                                            _ = v198
                                        }
                                        v199 := GenS13{F0: 2, F1: ("a" + ("hello" + ("" + ""))), F2: 0}
                                        if (128 != 127) {
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                        }
                                        fmt.Println(fi7)
                                        _ = v199
                                    }
                                }
                            }
                        case 9223372036854775807:
                            v200 := GenS28{F0: (fi15 * fi15), F1: 802, F2: 255}
                            _ = v200
                        case 811:
                            _ = len(v8)
                        default:
                            v201 := GenS23{F0: 0}
                            _ = v201
                        }
                        v202 := GenS29{F0: (fi15 | (fi7 ^ 32768)), F1: 42}
                        if p0 {
                            if ((!false) && (323 != -444)) {
                                var v203 uint8 = 42
                                var v204 int = (fi7 - (fi7 >> ((65535) & 63)))
                                _ = v203
                                _ = v204
                            } else {
                                switch (fi7 << ((506) & 63)) {
                                case -347:
                                    _ = len(v8)
                                case 127:
                                    var v205 uint32 = 0
                                    _ = v205
                                case 1:
                                    v6 = 42
                                }
                                switch "hello" {
                                case "\n\t":
                                    v206 := GenS25{F0: 42, F1: 228, F2: true}
                                    _ = v206
                                case "":
                                    var v207 string = ""
                                    v8 = append(v8, (fi15 << ((362) & 63)))
                                    if len(v8) > 0 {
                                        v8[0] = 9223372036854775807
                                    }
                                    _ = v207
                                case "a":
                                    v8 = append(v8, fi7)
                                }
                            }
                        } else {
                            v6 = v6
                            var v208 float32 = 1e10
                            var v209 string = ("" + ("hello" + ""))
                            v208 = v208
                            _ = v208
                            _ = v209
                        }
                        _ = v202
                    default:
                        var v210 uint16 = 1
                        _ = v210
                    }
                    v211 := GenS12{F0: 2, F1: 2, F2: -1}
                    if ("\n\t" >= ((("\n\t" + "hello") + "a") + "foo bar")) {
                        var v212 string = "foo bar"
                        v213 := GenS20{F0: 42, F1: 1, F2: -709}
                        _ = v212
                        _ = v213
                    }
                    var v214 uint32 = v6
                    _ = fi15
                    _ = v211
                    _ = v214
                }
            }
            _ = fi7
            _ = v8
        }
        var v215 int8 = -1
        _ = v215
        _ = v6
    }
    p0 = (0 < 128)
    p0 = p0
    if p0 {
        p0 = p0
    }
    switch "" {
    case "\n\t":
        var v216 int32 = -990
        _ = v216
    default:
        v217 := []int{}
        _ = v217
    }
    _ = p0
    return 0
}

func main() {
    var v1 float32 = 3.14159
    v2 := []int{9223372036854775807, -808}
    v3 := []int{-9223372036854775808, -2147483648}
    var v4 uint16 = 42
    var v218 uint32 = 4294967295
    var v219 uint8 = 154
    var v220 string = "hello"
    if (820 > 880) {
        for fi221 := 0; fi221 < 12; fi221++ {
            var v222 int32 = -2147483648
            _ = fi221
            _ = v222
        }
    } else {
        var v223 int64 = 42
        switch 2147483647 {
        case 0:
            var v224 uint16 = 32767
            _ = v224
        case 795:
            switch 32767 {
            case 1:
                v220 = v220
            case 404:
                v225 := GenS30{F0: v218}
                _ = v225
            }
        case 2147483647:
            v220 = "foo bar"
        }
        _ = v223
    }
    fmt.Println(v4, v1)
    var v226 float32 = -2.25
    switch "foo bar" {
    case "foo bar":
        switch -221 {
        case -822:
            v1 = ((v1 - v226) - 1.5)
        case 596:
            fmt.Println(v220, v218)
        case 0:
            v227 := GenS31{F0: true, F1: 100}
            _ = v227
        }
    }
    v220 = (("\n\t" + v220) + (v220 + "hello"))
    var v228 bool = (!true)
    if v228 {
        v229 := GenS9{F0: v228, F1: -2147483648, F2: 4294967295}
        _ = v229
    }
    v230 := GenS30{F0: v218}
    var v231 uint64 = 415
    for fi232 := 0; fi232 < 1; fi232++ {
        switch (^(fi232)) {
        case -1:
            switch 0 {
            case -9223372036854775808:
                _ = len(v3)
            default:
                var v233 uint64 = (^(v231))
                _ = v233
            }
        case 128:
            fmt.Println(v220, v226, v219)
        }
        _ = fi232
    }
    v234 := GenS32{F0: 19, F1: 1, F2: 239}
    fmt.Println(v231, v1, v228)
    v235 := GenS33{F0: 887}
    v231 -= (^(v231))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v218:", v218)
    fmt.Println("v219:", v219)
    fmt.Println("v220:", v220)
    fmt.Println("v226:", v226)
    fmt.Println("v228:", v228)
    fmt.Println("v231:", v231)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v218
    _ = v219
    _ = v220
    _ = v226
    _ = v228
    _ = v230
    _ = v231
    _ = v234
    _ = v235
    _ = v3
    _ = v4
}
