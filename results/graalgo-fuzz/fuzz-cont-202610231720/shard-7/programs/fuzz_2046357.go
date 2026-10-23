package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint8; F2 int8 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int16; F1 int64; F2 int32 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int8; F1 uint64; F2 int8 }
type GenS10 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS11 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS12 struct { F0 uint8; F1 int32; F2 string }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int; F1 int64; F2 uint16 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int32; F1 uint }

func main() {
    var v1 uint = 42
    var v2 uint32 = 0
    var v3 int8 = 127
    var v4 string = "a"
    v5 := []int{int(v3)}
    switch "\n\t" {
    case "\n\t":
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int = fi6
            for fi8 := 0; fi8 < 7; fi8++ {
                if true {
                    switch (v4 + v4) {
                    case "foo bar":
                        switch "a" {
                        case "hello":
                            if ((v1 ^ 65535) <= v1) {
                                if (128 >= 255) {
                                    var v9 uint16 = 760
                                    v4 = v4
                                    var v10 int64 = -872
                                    _ = v10
                                    _ = v9
                                } else {
                                    fmt.Println(fi6, v4, v4)
                                    fmt.Println(v4, fi8)
                                    var v11 string = "\n\t"
                                    switch "a" {
                                    case "hello":
                                        var v12 uint8 = 234
                                        _ = v12
                                    case "foo bar":
                                        v13 := GenS1{F0: 1, F1: (v3 & v3)}
                                        _ = v13
                                    case "a":
                                        if (true && (!false)) {
                                            v7 = 2147483647
                                            if true {
                                                switch v4 {
                                                case "":
                                                    switch v7 {
                                                    case 9223372036854775807:
                                                        v14 := GenS1{F0: -1, F1: (^(v3))}
                                                        _ = v14
                                                    }
                                                    if (int16(v2) <= -32768) {
                                                        var v15 int64 = 0
                                                        _ = v15
                                                    }
                                                case "a":
                                                    v16 := GenS1{F0: -1, F1: 100}
                                                    _ = v16
                                                }
                                                if (!true) {
                                                    var v17 float32 = ((3.14159 * -2.25) - 1e10)
                                                    switch (v7 | -445) {
                                                    case 610:
                                                        if false {
                                                            switch (v4 + "") {
                                                            case "\n\t":
                                                                _ = len(v5)
                                                            case "":
                                                                if len(v5) > 0 {
                                                                    v5[0] = 616
                                                                }
                                                                var v18 uint64 = 683
                                                                switch -1 {
                                                                case -779:
                                                                    switch v11 {
                                                                    case "hello":
                                                                        v5 = append(v5, (fi8 % ((fi6) | 1)))
                                                                    case "foo bar":
                                                                        var v19 int8 = (v3 + (v3 + v3))
                                                                        var v20 uint32 = uint32(v18)
                                                                        v1 = uint(v18)
                                                                        _ = v19
                                                                        _ = v20
                                                                    }
                                                                case -2147483648:
                                                                    if (!(93 == 2)) {
                                                                        if ((v7 | (fi8 ^ (v7 & (v7 & fi8)))) < (^(v7))) {
                                                                            if true {
                                                                                v3 += v3
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    switch (v11 + v11) {
                                                                                    case "\n\t":
                                                                                        var v21 int8 = v3
                                                                                        _ = v21
                                                                                    }
                                                                                default:
                                                                                    v4 = "foo bar"
                                                                                    v17 += v17
                                                                                    var v22 int8 = (v3 * (v3 / ((v3) | 1)))
                                                                                    v23 := GenS2{F0: 127, F1: 1}
                                                                                    _ = v22
                                                                                    _ = v23
                                                                                }
                                                                                v24 := GenS3{F0: -32768, F1: 2, F2: (v3 << ((2) & 63))}
                                                                                _ = v24
                                                                            }
                                                                        } else {
                                                                            var v25 int64 = -1
                                                                            fmt.Println(v3)
                                                                            var v26 int = 127
                                                                            _ = v25
                                                                            _ = v26
                                                                        }
                                                                        switch (v7 - fi8) {
                                                                        case 18:
                                                                            fmt.Println(v3, v11, v17)
                                                                        case -84:
                                                                            var v27 string = "foo bar"
                                                                            _ = v27
                                                                        case -1:
                                                                            fmt.Println(v18)
                                                                        default:
                                                                            if (-2.25 != 0.1) {
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = v7
                                                                                }
                                                                                fmt.Println(v7)
                                                                                var v28 uint16 = 1
                                                                                _ = v28
                                                                            }
                                                                            if (float32(1.5) == -2.25) {
                                                                                v11 = ("a" + v4)
                                                                            }
                                                                        }
                                                                        var v29 uint32 = v2
                                                                        _ = v29
                                                                    }
                                                                case 32768:
                                                                    var v30 bool = ((v1 == v1) || false)
                                                                    _ = v30
                                                                default:
                                                                    if false {
                                                                        if len(v5) > 0 {
                                                                            v5[0] = -129
                                                                        }
                                                                        switch 4294967295 {
                                                                        case -661:
                                                                            switch 42 {
                                                                            case 0:
                                                                                v3 -= 127
                                                                            case -594:
                                                                                v5 = append(v5, (v7 << ((v1) & 63)))
                                                                            case 127:
                                                                                switch (fi8 % ((9223372036854775807) | 1)) {
                                                                                case 4294967295:
                                                                                    v1 = (v1 ^ (v1 % ((v1) | 1)))
                                                                                }
                                                                            }
                                                                        case -9223372036854775808:
                                                                            v1 = (^(v1))
                                                                        case 679:
                                                                            if false {
                                                                                var v31 int = 9223372036854775807
                                                                                _ = v31
                                                                            }
                                                                        default:
                                                                            if len(v5) > 0 {
                                                                                v5[0] = fi8
                                                                            }
                                                                        }
                                                                        var v32 int64 = 9223372036854775807
                                                                        if (924 <= 983) {
                                                                            v1 = v1
                                                                            v33 := GenS4{F0: true, F1: uint8(v2), F2: 0}
                                                                            v5 = append(v5, fi8)
                                                                            _ = v33
                                                                        }
                                                                        _ = v32
                                                                    } else {
                                                                        var v34 uint64 = 405
                                                                        _ = v34
                                                                    }
                                                                }
                                                                _ = v18
                                                            default:
                                                                var v35 float64 = float64((v17 / 1.5))
                                                                _ = v35
                                                            }
                                                            if len(v5) > 0 {
                                                                v5[0] = -9223372036854775808
                                                            }
                                                            var v36 string = v4
                                                            var v37 float64 = 1e10
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        v38 := GenS1{F0: -330, F1: 42}
                                                        if true {
                                                            switch v4 {
                                                            case "a":
                                                                fmt.Println(v11, fi6, v11)
                                                            case "hello":
                                                                switch v11 {
                                                                case "":
                                                                    switch fi6 {
                                                                    case -568:
                                                                        v5 = append(v5, fi6)
                                                                    case 510:
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    default:
                                                                        var v39 bool = (42 < 1)
                                                                        _ = v39
                                                                    }
                                                                case "a":
                                                                    if true {
                                                                        var v40 uint8 = 91
                                                                        _ = v40
                                                                    }
                                                                case "\n\t":
                                                                    if false {
                                                                        var v41 int = (-(v7))
                                                                        if ((^(v1)) <= v1) {
                                                                            v17 -= 1.5
                                                                            v3 = v3
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            v42 := GenS5{F0: 32767, F1: -1, F2: 1}
                                                                            _ = v42
                                                                        } else {
                                                                            fmt.Println(v7, v41)
                                                                            v5 = append(v5, int(v3))
                                                                        }
                                                                        _ = v41
                                                                    } else {
                                                                        switch 42 {
                                                                        case -347:
                                                                            if true {
                                                                                if (float64(v17) > 0.1) {
                                                                                    if false {
                                                                                        var v43 uint32 = 452
                                                                                        v1 += v1
                                                                                        _ = v43
                                                                                    }
                                                                                    var v44 int16 = -1
                                                                                    _ = v44
                                                                                } else {
                                                                                    if (-546 != 1) {
                                                                                        if true {
                                                                                            var v45 int8 = (^(v3))
                                                                                            v1 = 42
                                                                                            v4 = "a"
                                                                                            if len(v5) >= 2 {
                                                                                                v5 = v5[1:len(v5)]
                                                                                            }
                                                                                            if false {
                                                                                                v11 = v11
                                                                                                v7 = fi6
                                                                                                v46 := GenS5{F0: int16(v3), F1: -1, F2: -2147483648}
                                                                                                v45 += v45
                                                                                                v47 := GenS3{F0: -816, F1: 255, F2: (v45 * (v45 << ((v1) & 63)))}
                                                                                                var v48 float32 = (((v17 * v17) - v17) * v17)
                                                                                                v2 = v2
                                                                                                switch fi8 {
                                                                                                case 4294967295:
                                                                                                    if (!true) {
                                                                                                        if len(v5) > 0 {
                                                                                                            v5[0] = (v7 & (-(fi6)))
                                                                                                        }
                                                                                                        _ = len(v5)
                                                                                                        var v49 int = (fi6 & fi6)
                                                                                                        _ = v49
                                                                                                    }
                                                                                                }
                                                                                                v50 := GenS1{F0: 32767, F1: v45}
                                                                                                if false {
                                                                                                    v45 = 42
                                                                                                }
                                                                                                v51 := GenS6{F0: false, F1: 42}
                                                                                                _ = v46
                                                                                                _ = v47
                                                                                                _ = v48
                                                                                                _ = v50
                                                                                                _ = v51
                                                                                            } else {
                                                                                                switch fi8 {
                                                                                                case 2147483647:
                                                                                                    v4 = "a"
                                                                                                case 127:
                                                                                                    var v52 uint64 = 318
                                                                                                    _ = v52
                                                                                                case -320:
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = 0
                                                                                                    }
                                                                                                }
                                                                                                var v53 int8 = 0
                                                                                                if len(v5) > 0 {
                                                                                                    v5[0] = v7
                                                                                                }
                                                                                                _ = v53
                                                                                            }
                                                                                            _ = v45
                                                                                        }
                                                                                        if true {
                                                                                            fmt.Println(v4)
                                                                                        }
                                                                                    } else {
                                                                                        var v54 uint16 = 65535
                                                                                        var v55 float64 = 0.1
                                                                                        var v56 uint8 = 237
                                                                                        _ = v54
                                                                                        _ = v55
                                                                                        _ = v56
                                                                                    }
                                                                                    v3 += -128
                                                                                    switch v11 {
                                                                                    case "foo bar":
                                                                                        _ = len(v5)
                                                                                    }
                                                                                }
                                                                                v5 = append(v5, (fi8 << ((uint(v2)) & 63)))
                                                                                v7 -= -1
                                                                            }
                                                                        case 4294967295:
                                                                            if (42 > -9223372036854775808) {
                                                                                v4 = v4
                                                                                v57 := GenS7{F0: 0}
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = (-(fi6))
                                                                                }
                                                                                _ = v57
                                                                            } else {
                                                                                if (!(1 == 9223372036854775807)) {
                                                                                    var v58 uint32 = v2
                                                                                    var v59 uint16 = 1
                                                                                    _ = v58
                                                                                    _ = v59
                                                                                } else {
                                                                                    switch v4 {
                                                                                    case "hello":
                                                                                        var v60 uint64 = 4294967295
                                                                                        _ = v60
                                                                                    case "foo bar":
                                                                                        var v61 int16 = 0
                                                                                        _ = v61
                                                                                    }
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = fi6
                                                                                    }
                                                                                }
                                                                                v17 = (v17 + 0.1)
                                                                            }
                                                                        }
                                                                        var v62 int = fi8
                                                                        v63 := GenS3{F0: -269, F1: 42, F2: (v3 << (((v1 / ((4294967295) | 1))) & 63))}
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                        v64 := GenS2{F0: 42, F1: 42}
                                                                        v65 := GenS4{F0: false, F1: 180, F2: 1}
                                                                        var v66 bool = true
                                                                        var v67 uint = v1
                                                                        var v68 int8 = (v3 << ((v1) & 63))
                                                                        _ = v62
                                                                        _ = v63
                                                                        _ = v64
                                                                        _ = v65
                                                                        _ = v66
                                                                        _ = v67
                                                                        _ = v68
                                                                    }
                                                                default:
                                                                    switch v7 {
                                                                    case 763:
                                                                        if ((32767 <= -596) && true) {
                                                                            fmt.Println(v17, v17)
                                                                            if (!(false && (!false))) {
                                                                                var v69 int64 = 0
                                                                                fmt.Println(fi8, v69, v4)
                                                                                if ((fi8 == 770) && (v1 == 65535)) {
                                                                                    var v70 uint32 = v2
                                                                                    _ = v70
                                                                                }
                                                                                v69 = v69
                                                                                _ = v69
                                                                            } else {
                                                                                v4 = v4
                                                                                v1 = (^(v1))
                                                                                switch v11 {
                                                                                case "a":
                                                                                    v2 = (v2 + v2)
                                                                                case "\n\t":
                                                                                    if (v1 == v1) {
                                                                                        if ((fi8 ^ fi6) != (-(fi8))) {
                                                                                            _ = len(v5)
                                                                                        } else {
                                                                                            fmt.Println(v1, v2)
                                                                                            v71 := GenS4{F0: (1 > 1), F1: 90, F2: 42}
                                                                                            v72 := GenS6{F0: true, F1: 2}
                                                                                            var v73 uint8 = 0
                                                                                            _ = v71
                                                                                            _ = v72
                                                                                            _ = v73
                                                                                        }
                                                                                        if false {
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = v7
                                                                                            }
                                                                                            var v74 int32 = int32(v1)
                                                                                            v11 = v11
                                                                                            _ = v74
                                                                                        } else {
                                                                                            v4 = "a"
                                                                                            if (false || true) {
                                                                                                if true {
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = (v7 >> (((v1 ^ v1)) & 63))
                                                                                                    }
                                                                                                    var v75 uint64 = uint64(v2)
                                                                                                    switch 4294967295 {
                                                                                                    case 408:
                                                                                                        v5 = append(v5, (fi6 / ((v7) | 1)))
                                                                                                    default:
                                                                                                        var v76 uint = (v1 * uint(v75))
                                                                                                        _ = v76
                                                                                                    }
                                                                                                    _ = v75
                                                                                                }
                                                                                                if (v1 == (v1 ^ (v1 + (^(v1))))) {
                                                                                                    v77 := GenS4{F0: false, F1: uint8(v1), F2: 0}
                                                                                                    var v78 int16 = 1
                                                                                                    switch (-(v7)) {
                                                                                                    case 75:
                                                                                                        fmt.Println(fi6)
                                                                                                    case -473:
                                                                                                        v2 = v2
                                                                                                        if (!(v3 != v3)) {
                                                                                                            v79 := GenS4{F0: (!(!(!(!true)))), F1: 44, F2: 92}
                                                                                                            _ = v79
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v77
                                                                                                    _ = v78
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v80 := GenS5{F0: -291, F1: 1, F2: 1}
                                                                                    v81 := GenS4{F0: (324 > 560), F1: 0, F2: 127}
                                                                                    _ = v80
                                                                                    _ = v81
                                                                                default:
                                                                                    fmt.Println(v11, v4, v1)
                                                                                }
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v2, v7, fi6)
                                                                            if len(v5) > 0 {
                                                                                v5[0] = (fi8 >> ((v1) & 63))
                                                                            }
                                                                            v5 = append(v5, (^(fi8)))
                                                                        }
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                if true {
                                                                    v5 = append(v5, 9223372036854775807)
                                                                    v2 = (v2 / (((v2 % (((v2 % ((v2) | 1))) | 1))) | 1))
                                                                    var v82 uint8 = 31
                                                                    switch (fi8 + fi8) {
                                                                    case 213:
                                                                        v17 -= ((float32((-2.25 / (1.5 - 0.1))) - ((v17 - v17) - (v17 - 0.0))) + float32(3.14159))
                                                                    case -9223372036854775808:
                                                                        v83 := GenS1{F0: 42, F1: 53}
                                                                        _ = v83
                                                                    case 42:
                                                                        v17 = (0.0 + float32(-2.25))
                                                                        switch v4 {
                                                                        case "hello":
                                                                            v11 = ("foo bar" + (v4 + v11))
                                                                        case "a":
                                                                            var v84 uint16 = 42
                                                                            switch v4 {
                                                                            case "":
                                                                                switch (fi6 * (fi6 ^ 110)) {
                                                                                case 32768:
                                                                                    var v85 int = -448
                                                                                    _ = v85
                                                                                default:
                                                                                    var v86 float64 = 1.5
                                                                                    _ = v86
                                                                                }
                                                                            case "\n\t":
                                                                                var v87 uint8 = 1
                                                                                _ = v87
                                                                            case "foo bar":
                                                                                fmt.Println(v3, v7, fi8)
                                                                            default:
                                                                                v2 -= (v2 >> ((4294967295) & 63))
                                                                            }
                                                                            _ = v84
                                                                        }
                                                                    }
                                                                    _ = v82
                                                                }
                                                            default:
                                                                fmt.Println(v7, v2, v7)
                                                            }
                                                            var v88 int16 = 1000
                                                            _ = v88
                                                        }
                                                        if (((v3 + 127) <= -128) && (588 <= 9223372036854775807)) {
                                                            v89 := GenS4{F0: ((v3 ^ (v3 % ((v3) | 1))) < v3), F1: 128, F2: 1}
                                                            var v90 int8 = 0
                                                            if false {
                                                                switch (v4 + v4) {
                                                                case "hello":
                                                                    var v91 uint16 = 65535
                                                                    _ = v91
                                                                case "":
                                                                    v92 := GenS1{F0: -1, F1: v3}
                                                                    _ = v92
                                                                }
                                                                v90 = (v90 + (^(v90)))
                                                                v7 = (v7 + fi6)
                                                            }
                                                            v93 := GenS6{F0: (!(1 <= 42)), F1: 228}
                                                            _ = v89
                                                            _ = v90
                                                            _ = v93
                                                        }
                                                        var v94 int16 = -1
                                                        if (v2 <= v2) {
                                                            v95 := GenS8{F0: 127}
                                                            v94 = (v94 << (((v1 >> ((v1) & 63))) & 63))
                                                            _ = v95
                                                        } else {
                                                            var v96 uint = (v1 % (((v1 + (^(v1)))) | 1))
                                                            v97 := GenS9{F0: (v3 + (v3 & (-(v3)))), F1: 42, F2: (v3 + v3)}
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        _ = v38
                                                        _ = v94
                                                    case -814:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    _ = v17
                                                }
                                                if ((v3 + v3) <= (-(v3))) {
                                                    var v98 int8 = (v3 ^ (v3 * (v3 ^ (v3 | (v3 & v3)))))
                                                    if false {
                                                        var v99 uint32 = 2147483647
                                                        v5 = append(v5, (v7 << ((v1) & 63)))
                                                        if len(v5) > 0 {
                                                            v5[0] = (^(fi6))
                                                        }
                                                        _ = v99
                                                    } else {
                                                        var v100 int = (fi8 | fi8)
                                                        var v101 int8 = v98
                                                        var v102 int = 128
                                                        _ = v100
                                                        _ = v101
                                                        _ = v102
                                                    }
                                                    _ = v98
                                                } else {
                                                    v103 := GenS6{F0: false, F1: 42}
                                                    _ = v103
                                                }
                                            } else {
                                                _ = len(v5)
                                                v5 = append(v5, -1)
                                                fmt.Println(v7, v11, fi8)
                                            }
                                            switch -2147483648 {
                                            case 87:
                                                if len(v5) > 0 {
                                                    v5[0] = (v7 ^ (fi8 * fi8))
                                                }
                                            case 32768:
                                                if (42 > -9223372036854775808) {
                                                    var v104 float64 = (0.1 - float64(((float32((1e10 + 3.14159)) - (3.14159 / 1.5)) + -2.25)))
                                                    var v105 uint8 = 36
                                                    _ = v104
                                                    _ = v105
                                                } else {
                                                    v106 := GenS10{F0: 225, F1: 128, F2: 0}
                                                    v1 = (v1 >> ((v1) & 63))
                                                    _ = v106
                                                }
                                            case -779:
                                                var v107 int8 = v3
                                                v1 = (v1 - 2)
                                                v107 -= v107
                                                _ = v107
                                            }
                                        }
                                    default:
                                        _ = len(v5)
                                    }
                                    switch v11 {
                                    case "hello":
                                        _ = len(v5)
                                    case "":
                                        v108 := GenS10{F0: 42, F1: uint8(v2), F2: 9223372036854775807}
                                        _ = v108
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if (false && false) {
                                        var v109 uint = (v1 % ((v1) | 1))
                                        v110 := GenS9{F0: (v3 * v3), F1: 2, F2: (-(v3))}
                                        v3 -= (v3 / ((0) | 1))
                                        _ = v109
                                        _ = v110
                                    }
                                    _ = v11
                                }
                            }
                        default:
                            var v111 int32 = 2147483647
                            _ = v111
                        }
                    }
                } else {
                    v112 := GenS11{F0: 0, F1: 910, F2: 2}
                    var v113 uint16 = 1
                    fmt.Println(fi8, v1, v1)
                    _ = v112
                    _ = v113
                }
                switch fi6 {
                case -981:
                    v3 -= (^(v3))
                }
                _ = fi8
            }
            if true {
                var v114 int = fi6
                if true {
                    v115 := GenS12{F0: 177, F1: int32(v3), F2: ("" + v4)}
                    var v116 uint8 = 0
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v115
                    _ = v116
                }
                _ = v114
            }
            _ = fi6
            _ = v7
        }
    case "":
        if len(v5) > 0 {
            v5[0] = -500
        }
    }
    var v117 float64 = 3.14159
    v118 := GenS4{F0: (1.5 > (1.5 + 3.14159)), F1: 128, F2: 39}
    switch "" {
    case "a":
        if false {
            var v119 int64 = 0
            _ = v119
        }
    case "":
        if false {
            var v120 int16 = 32767
            for fi121 := 0; fi121 < 10; fi121++ {
                var v122 int32 = 0
                _ = fi121
                _ = v122
            }
            switch v4 {
            case "":
                var v123 uint8 = 2
                _ = v123
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = -663
                }
            case "hello":
                v124 := GenS13{F0: v2}
                _ = v124
            default:
                var v125 uint = v1
                _ = v125
            }
            _ = v120
        }
    case "foo bar":
        v126 := GenS6{F0: (true || (2 <= 975)), F1: 2}
        _ = v126
    default:
        fmt.Println(v117, v117)
    }
    var v127 int32 = -1
    for fi128 := 0; fi128 < 8; fi128++ {
        v4 = "hello"
        _ = fi128
    }
    v117 = v117
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v127, v2, v1)
    case "":
        switch "" {
        case "a":
            v129 := GenS12{F0: 92, F1: (v127 | 613), F2: (v4 + v4)}
            _ = v129
        case "foo bar":
            switch (v4 + "a") {
            case "\n\t":
                v127 = (v127 * v127)
            case "a":
                v3 -= v3
            default:
                switch -9223372036854775808 {
                case 566:
                    v130 := GenS14{F0: 32767, F1: 9223372036854775807, F2: 1}
                    _ = v130
                case -1:
                    for fi131 := 0; fi131 < 10; fi131++ {
                        var v132 float64 = (-2.25 + v117)
                        v2 = (v2 & (v2 * 2147483647))
                        _ = fi131
                        _ = v132
                    }
                case 2147483647:
                    var v133 int8 = v3
                    _ = v133
                }
            }
        default:
            v4 = v4
        }
    case "a":
        v134 := GenS7{F0: 65535}
        _ = v134
    default:
        var v135 uint8 = 255
        _ = v135
    }
    fmt.Println(v127, v1, v117)
    switch int(v1) {
    case 982:
        v1 -= v1
    case -571:
        v117 += v117
    default:
        v136 := GenS15{F0: 632}
        _ = v136
    }
    for fi137 := 0; fi137 < 2; fi137++ {
        if true {
            if len(v5) > 0 {
                v5[0] = fi137
            }
            v138 := GenS15{F0: 42}
            _ = v138
        }
        var v139 uint16 = 255
        v117 += ((0.0 / 1.0) * v117)
        _ = fi137
        _ = v139
    }
    var v140 uint32 = 42
    if (313 > 42) {
        v141 := GenS16{F0: v127, F1: v1}
        v127 = v127
        var v142 uint8 = 255
        _ = v141
        _ = v142
    } else {
        switch v4 {
        case "hello":
            var v143 string = "\n\t"
            _ = v143
        case "":
            v127 += v127
        case "a":
            var v144 int32 = v127
            _ = v144
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        for fi145 := 0; fi145 < 19; fi145++ {
            var v146 uint16 = 1
            v147 := GenS16{F0: (^(v127)), F1: 255}
            _ = fi145
            _ = v146
            _ = v147
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v127:", v127)
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v117
    _ = v118
    _ = v127
    _ = v140
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
