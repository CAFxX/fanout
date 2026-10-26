package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16; F1 int64 }
type GenS9 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS10 struct { F0 string; F1 uint16 }
type GenS11 struct { F0 int64; F1 int16; F2 string }
type GenS12 struct { F0 int; F1 uint }
type GenS13 struct { F0 int8; F1 int; F2 int32 }
type GenS14 struct { F0 uint16; F1 int64; F2 bool }
type GenS15 struct { F0 uint; F1 uint16; F2 int64 }
type GenS16 struct { F0 int8; F1 int; F2 int64 }
type GenS17 struct { F0 int16; F1 uint16 }
type GenS18 struct { F0 uint8; F1 int; F2 string }
type GenS19 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS20 struct { F0 int64; F1 int64; F2 int8 }

func main() {
    var v1 uint8 = 128
    var v2 string = "\n\t"
    v3 := []uint32{296, 4294967295}
    var v4 int64 = -262
    v5 := []int{}
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int32 = -2147483648
        _ = fi6
        _ = v7
    }
    v1 -= (v1 ^ 42)
    v4 = v4
    if true {
        v3 = append(v3, 0)
        v3 = append(v3, 64)
    } else {
        var v8 float32 = -2.25
        switch v2 {
        case "foo bar":
            switch ((((v2 + (v2 + (v2 + "\n\t"))) + v2) + v2) + (v2 + ((v2 + "a") + "hello"))) {
            case "\n\t":
                if false {
                    for fi9 := 0; fi9 < 2; fi9++ {
                        if true {
                            v10 := GenS1{F0: 20, F1: 42}
                            switch 264 {
                            case 127:
                                for fi11 := 0; fi11 < 13; fi11++ {
                                    if len(v3) > 0 {
                                        v3[0] = 597
                                    }
                                    v3 = append(v3, 2)
                                    fmt.Println(v2, v2, fi9)
                                    fmt.Println(fi9, v4, v4)
                                    v1 += v1
                                    v12 := GenS2{F0: 65535, F1: 117}
                                    _ = fi11
                                    _ = v12
                                }
                            case 32768:
                                v4 += (v4 % ((-141) | 1))
                            }
                            _ = v10
                        }
                        var v13 int32 = 0
                        if (0 <= 32767) {
                            if true {
                                _ = len(v5)
                                switch (fi9 / (((fi9 % ((-9223372036854775808) | 1))) | 1)) {
                                case 154:
                                    var v14 int16 = -620
                                    _ = v14
                                case 570:
                                    for fi15 := 0; fi15 < 9; fi15++ {
                                        switch 361 {
                                        case 9223372036854775807:
                                            if true {
                                                v16 := GenS2{F0: 80, F1: 2}
                                                if len(v3) > 0 {
                                                    v3[0] = 1
                                                }
                                                fmt.Println(v1, v2, v1)
                                                var v17 int32 = 774
                                                v18 := GenS2{F0: 4294967295, F1: 255}
                                                _ = v16
                                                _ = v17
                                                _ = v18
                                            } else {
                                                if (2147483647 != 0) {
                                                    switch fi15 {
                                                    case 415:
                                                        v8 = v8
                                                    default:
                                                        if false {
                                                            fmt.Println(fi9, v4)
                                                            if true {
                                                                if true {
                                                                    if (-814 != int(v4)) {
                                                                        if ((v4 & v4) != v4) {
                                                                            var v19 string = ("hello" + "foo bar")
                                                                            v20 := GenS3{F0: false, F1: 2}
                                                                            _ = v19
                                                                            _ = v20
                                                                        } else {
                                                                            fmt.Println(fi9, v8)
                                                                            if (false || true) {
                                                                                switch fi15 {
                                                                                case 1:
                                                                                    v21 := GenS1{F0: 476, F1: v13}
                                                                                    _ = v21
                                                                                case -605:
                                                                                    _ = len(v3)
                                                                                default:
                                                                                    v22 := GenS1{F0: 177, F1: -1}
                                                                                    _ = v22
                                                                                }
                                                                                var v23 uint64 = 5
                                                                                _ = v23
                                                                            } else {
                                                                                var v24 uint = 4294967295
                                                                                var v25 uint16 = 0
                                                                                var v26 string = "hello"
                                                                                _ = v24
                                                                                _ = v25
                                                                                _ = v26
                                                                            }
                                                                            v27 := GenS1{F0: 42, F1: v13}
                                                                            v13 -= (v13 ^ (v13 | v13))
                                                                            _ = v27
                                                                        }
                                                                    } else {
                                                                        v3 = append(v3, 4294967295)
                                                                    }
                                                                    var v28 uint64 = 894
                                                                    v29 := GenS3{F0: true, F1: (v28 | (v28 >> ((898) & 63)))}
                                                                    _ = v28
                                                                    _ = v29
                                                                }
                                                                if ((v1 + (v1 + (v1 ^ (v1 ^ (v1 & 239))))) >= v1) {
                                                                    switch 4294967295 {
                                                                    case 749:
                                                                        var v30 uint = 0
                                                                        fmt.Println(v4, fi15)
                                                                        var v31 int16 = 1000
                                                                        if true {
                                                                            if (632 != uint32(v30)) {
                                                                                v2 = (v2 + v2)
                                                                                var v32 int32 = (v13 | v13)
                                                                                _ = v32
                                                                            } else {
                                                                                v33 := GenS2{F0: 4294967295, F1: (v30 * v30)}
                                                                                var v34 int8 = 127
                                                                                _ = v33
                                                                                _ = v34
                                                                            }
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            var v35 float32 = v8
                                                                            _ = v35
                                                                        } else {
                                                                            switch (fi15 << (((v30 ^ (v30 ^ 255))) & 63)) {
                                                                            case -410:
                                                                                v8 -= 1.5
                                                                            case 42:
                                                                                v36 := GenS1{F0: 255, F1: -1}
                                                                                _ = v36
                                                                            case -2147483648:
                                                                                v5 = append(v5, fi9)
                                                                            default:
                                                                                _ = len(v3)
                                                                            }
                                                                            var v37 int = 2147483647
                                                                            _ = v37
                                                                        }
                                                                        v38 := GenS2{F0: 82, F1: v30}
                                                                        var v39 string = v2
                                                                        if false {
                                                                            if len(v5) > 0 {
                                                                                v5[0] = (fi15 % ((-2147483648) | 1))
                                                                            }
                                                                            var v40 float64 = 1e10
                                                                            var v41 int8 = -128
                                                                            _ = v40
                                                                            _ = v41
                                                                        } else {
                                                                            if false {
                                                                                _ = len(v3)
                                                                                if ((1e10 - v8) <= v8) {
                                                                                    var v42 uint64 = 4294967295
                                                                                    if (v13 != v13) {
                                                                                        v2 = (((v2 + ("a" + (v2 + (v39 + "foo bar")))) + "foo bar") + "\n\t")
                                                                                        _ = len(v3)
                                                                                    } else {
                                                                                        var v43 int16 = v31
                                                                                        v4 = v4
                                                                                        _ = v43
                                                                                    }
                                                                                    _ = v42
                                                                                }
                                                                                v2 = (v2 + v39)
                                                                            } else {
                                                                                fmt.Println(v8, v8)
                                                                            }
                                                                            v31 -= 1000
                                                                            var v44 int16 = 32767
                                                                            v45 := GenS1{F0: 32767, F1: -2147483648}
                                                                            if false {
                                                                                v46 := GenS2{F0: 2147483648, F1: v30}
                                                                                var v47 float32 = (v8 / (v8 + (v8 - v8)))
                                                                                var v48 uint8 = (v1 % ((v1) | 1))
                                                                                _ = v46
                                                                                _ = v47
                                                                                _ = v48
                                                                            } else {
                                                                                v49 := GenS1{F0: 507, F1: v13}
                                                                                _ = v49
                                                                            }
                                                                            v50 := GenS4{F0: false, F1: (v30 - v30)}
                                                                            _ = v44
                                                                            _ = v45
                                                                            _ = v50
                                                                        }
                                                                        if (true || false) {
                                                                            switch ("foo bar" + (("\n\t" + v39) + "a")) {
                                                                            case "foo bar":
                                                                                switch (fi15 ^ (-(fi9))) {
                                                                                case 128:
                                                                                    v51 := GenS4{F0: (196 == 127), F1: (v30 << ((v30) & 63))}
                                                                                    _ = v51
                                                                                case -1:
                                                                                    v52 := GenS5{F0: 42, F1: 42, F2: (v30 % ((v30) | 1))}
                                                                                    _ = v52
                                                                                default:
                                                                                    var v53 int32 = (v13 / ((v13) | 1))
                                                                                    _ = v53
                                                                                }
                                                                            case "hello":
                                                                                v54 := GenS3{F0: true, F1: 12}
                                                                                _ = v54
                                                                            }
                                                                            v55 := GenS6{F0: 1, F1: v39}
                                                                            if false {
                                                                                v4 += v4
                                                                                if (1.5 != 3.14159) {
                                                                                    if true {
                                                                                        v30 = (v30 & v30)
                                                                                        fmt.Println(v2)
                                                                                        switch (fi15 & (fi9 << ((4294967295) & 63))) {
                                                                                        case 32767:
                                                                                            if (!true) {
                                                                                                var v56 float64 = 0.1
                                                                                                _ = v56
                                                                                            } else {
                                                                                                fmt.Println(v30, v39)
                                                                                                switch "" {
                                                                                                case "a":
                                                                                                    var v57 string = v39
                                                                                                    v1 = v1
                                                                                                    var v58 int64 = (v4 / ((v4) | 1))
                                                                                                    _ = v57
                                                                                                    _ = v58
                                                                                                case "":
                                                                                                    if (false && false) {
                                                                                                        var v59 int64 = (-(v4))
                                                                                                        if (true || false) {
                                                                                                            v8 -= 0.0
                                                                                                            switch 474 {
                                                                                                            case 2147483647:
                                                                                                                fmt.Println(v31)
                                                                                                            case 9223372036854775807:
                                                                                                                var v60 int8 = -128
                                                                                                                _ = v60
                                                                                                            }
                                                                                                            var v61 uint8 = uint8(v30)
                                                                                                            if true {
                                                                                                                if len(v3) > 0 {
                                                                                                                    v3[0] = 161
                                                                                                                }
                                                                                                                var v62 uint32 = 882
                                                                                                                _ = v62
                                                                                                            }
                                                                                                            _ = v61
                                                                                                        } else {
                                                                                                            var v63 string = v39
                                                                                                            v39 = "foo bar"
                                                                                                            _ = v63
                                                                                                        }
                                                                                                        _ = v59
                                                                                                    } else {
                                                                                                        fmt.Println(v13, v13, v4)
                                                                                                        var v64 int8 = 1
                                                                                                        _ = v64
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    v31 += v31
                                                                                                default:
                                                                                                    v13 = v13
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            var v65 uint8 = v1
                                                                                            _ = v65
                                                                                        }
                                                                                        switch (fi15 * (fi9 - fi9)) {
                                                                                        case 542:
                                                                                            if false {
                                                                                                v1 = v1
                                                                                            }
                                                                                        case -1:
                                                                                            var v66 int32 = v13
                                                                                            _ = v66
                                                                                        }
                                                                                        v13 -= v13
                                                                                        var v67 string = v2
                                                                                        if false {
                                                                                            _ = len(v5)
                                                                                            fmt.Println(v30, v8)
                                                                                            v30 -= 1
                                                                                        }
                                                                                        _ = v67
                                                                                    } else {
                                                                                        v68 := GenS7{F0: (v4 ^ -1)}
                                                                                        _ = v68
                                                                                    }
                                                                                } else {
                                                                                    var v69 int32 = v13
                                                                                    if (!(42 < uint16(v1))) {
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            var v70 int16 = 1000
                                                                                            _ = v70
                                                                                        case "foo bar":
                                                                                            var v71 int32 = 1
                                                                                            _ = v71
                                                                                        case "":
                                                                                            switch fi9 {
                                                                                            case 42:
                                                                                                v8 = v8
                                                                                                var v72 int32 = (v69 % (((v69 - (v13 * v13))) | 1))
                                                                                                var v73 uint8 = (v1 % (((v1 % ((v1) | 1))) | 1))
                                                                                                v74 := GenS5{F0: 1, F1: 0, F2: (v30 + v30)}
                                                                                                _ = v72
                                                                                                _ = v73
                                                                                                _ = v74
                                                                                            case 0:
                                                                                                switch fi15 {
                                                                                                case 42:
                                                                                                    var v75 uint8 = v1
                                                                                                    _ = v75
                                                                                                case 0:
                                                                                                    switch 128 {
                                                                                                    case -719:
                                                                                                        if (!((^(v31)) != (^(v31)))) {
                                                                                                            v8 = -2.25
                                                                                                            _ = len(v3)
                                                                                                        }
                                                                                                    case 127:
                                                                                                        switch v39 {
                                                                                                        case "":
                                                                                                            var v76 int8 = int8(fi15)
                                                                                                            _ = v76
                                                                                                        case "a":
                                                                                                            if false {
                                                                                                                if false {
                                                                                                                    if len(v5) >= 2 {
                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    fmt.Println(fi15)
                                                                                                                    v77 := GenS8{F0: (^(v31)), F1: int64(v1)}
                                                                                                                    var v78 bool = false
                                                                                                                    _ = v77
                                                                                                                    _ = v78
                                                                                                                }
                                                                                                                fmt.Println(v39, v2)
                                                                                                                if (278 <= 18446744073709551615) {
                                                                                                                    if ((v4 & v4) != 281) {
                                                                                                                        var v79 float64 = (3.14159 * (0.0 - 0.0))
                                                                                                                        if len(v5) >= 2 {
                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                        }
                                                                                                                        _ = v79
                                                                                                                    }
                                                                                                                    _ = len(v5)
                                                                                                                    v80 := GenS9{F0: 543, F1: 65535, F2: (v1 * (v1 / ((1) | 1)))}
                                                                                                                    var v81 int16 = v31
                                                                                                                    v69 = (v13 ^ int32(v30))
                                                                                                                    var v82 uint16 = 1
                                                                                                                    _ = v80
                                                                                                                    _ = v81
                                                                                                                    _ = v82
                                                                                                                }
                                                                                                                if ((fi9 / ((-177) | 1)) <= 1) {
                                                                                                                    v5 = append(v5, fi9)
                                                                                                                    v8 -= 1e10
                                                                                                                    fmt.Println(v30)
                                                                                                                } else {
                                                                                                                    _ = len(v3)
                                                                                                                    v83 := GenS9{F0: 561, F1: (v30 % ((v30) | 1)), F2: v1}
                                                                                                                    _ = v83
                                                                                                                }
                                                                                                            }
                                                                                                            var v84 uint8 = 42
                                                                                                            _ = v84
                                                                                                        case "hello":
                                                                                                            v85 := GenS9{F0: 871, F1: 189, F2: v1}
                                                                                                            _ = v85
                                                                                                        default:
                                                                                                            if (0 != 42) {
                                                                                                                fmt.Println(v13, v13, v8)
                                                                                                                if true {
                                                                                                                    var v86 int64 = int64(v1)
                                                                                                                    v1 = 127
                                                                                                                    _ = v86
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v87 uint32 = 42
                                                                                                                if (1000 < v31) {
                                                                                                                    v87 += v87
                                                                                                                    v2 = v2
                                                                                                                    v88 := GenS10{F0: "a", F1: 1}
                                                                                                                    _ = v88
                                                                                                                } else {
                                                                                                                    switch "hello" {
                                                                                                                    case "\n\t":
                                                                                                                        fmt.Println(v30)
                                                                                                                    case "":
                                                                                                                        v1 -= (^(v1))
                                                                                                                    default:
                                                                                                                        var v89 int32 = v69
                                                                                                                        _ = v89
                                                                                                                    }
                                                                                                                    var v90 int16 = (v31 | int16(v30))
                                                                                                                    var v91 float32 = (1.5 + 1.5)
                                                                                                                    _ = v90
                                                                                                                    _ = v91
                                                                                                                }
                                                                                                                _ = v87
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        if len(v3) > 0 {
                                                                                                            v3[0] = 2
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v30, v69)
                                                                                                    v39 = "\n\t"
                                                                                                }
                                                                                            case 32768:
                                                                                                var v92 int = fi9
                                                                                                _ = v92
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v93 := GenS9{F0: 4294967295, F1: v30, F2: v1}
                                                                                    _ = v69
                                                                                    _ = v93
                                                                                }
                                                                                v94 := GenS11{F0: int64(fi15), F1: v31, F2: (v2 + v39)}
                                                                                _ = v94
                                                                            } else {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                var v95 int = (^(fi9))
                                                                                var v96 uint64 = 258
                                                                                _ = v95
                                                                                _ = v96
                                                                            }
                                                                            _ = v55
                                                                        } else {
                                                                            var v97 int16 = v31
                                                                            var v98 uint16 = 255
                                                                            _ = v97
                                                                            _ = v98
                                                                        }
                                                                        switch 32767 {
                                                                        case 4294967295:
                                                                            switch fi15 {
                                                                            case -2147483648:
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    var v99 int = fi9
                                                                                    _ = v99
                                                                                case "a":
                                                                                    _ = len(v3)
                                                                                default:
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        v100 := GenS9{F0: 1, F1: v30, F2: 2}
                                                                                        _ = v100
                                                                                    }
                                                                                    v101 := GenS7{F0: 42}
                                                                                    v39 = (v2 + v2)
                                                                                    _ = v101
                                                                                }
                                                                                if false {
                                                                                    if true {
                                                                                        var v102 int64 = v4
                                                                                        if (-372 < (v4 << ((227) & 63))) {
                                                                                            v30 = (v30 / (((^(v30))) | 1))
                                                                                            if (0 > (v1 >> ((1) & 63))) {
                                                                                                switch 127 {
                                                                                                case 169:
                                                                                                    if (v31 >= (v31 << (((v30 * v30)) & 63))) {
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v4)
                                                                                                case 0:
                                                                                                    v8 -= v8
                                                                                                case 127:
                                                                                                    v103 := GenS12{F0: fi9, F1: 255}
                                                                                                    _ = v103
                                                                                                default:
                                                                                                    var v104 float32 = (3.14159 - (v8 + float32(1e10)))
                                                                                                    _ = v104
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v30, fi15)
                                                                                        } else {
                                                                                            var v105 int64 = v4
                                                                                            if (-2.25 < 1e10) {
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = 65535
                                                                                                }
                                                                                                v13 = -512
                                                                                                var v106 uint16 = 1
                                                                                                _ = v106
                                                                                            }
                                                                                            v3 = append(v3, 2147483647)
                                                                                            v2 = ("a" + (v39 + v39))
                                                                                            _ = len(v5)
                                                                                            _ = v105
                                                                                        }
                                                                                        v31 -= -602
                                                                                        v107 := GenS10{F0: v39, F1: 2}
                                                                                        switch fi15 {
                                                                                        case 938:
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = (fi9 * -1)
                                                                                            }
                                                                                        case 32767:
                                                                                            v8 = 1.5
                                                                                        default:
                                                                                            var v108 float64 = (1e10 * 0.0)
                                                                                            _ = v108
                                                                                        }
                                                                                        _ = v102
                                                                                        _ = v107
                                                                                    }
                                                                                    var v109 uint32 = 2147483647
                                                                                    _ = len(v5)
                                                                                    _ = v109
                                                                                }
                                                                            default:
                                                                                if (!true) {
                                                                                    fmt.Println(v1, v39, v4)
                                                                                    v110 := GenS13{F0: -128, F1: fi15, F2: (^(v13))}
                                                                                    _ = v110
                                                                                }
                                                                            }
                                                                        case 2147483647:
                                                                            v1 -= (v1 ^ (v1 & (v1 ^ (v1 - v1))))
                                                                        case 9223372036854775807:
                                                                            var v111 uint64 = 1
                                                                            _ = v111
                                                                        default:
                                                                            switch fi9 {
                                                                            case -738:
                                                                                if (v31 != v31) {
                                                                                    fmt.Println(fi15, v8)
                                                                                    v112 := GenS5{F0: 2147483648, F1: 413, F2: (v30 / ((v30) | 1))}
                                                                                    _ = v112
                                                                                } else {
                                                                                    v1 -= (v1 + 30)
                                                                                }
                                                                                v113 := GenS5{F0: uint32(v1), F1: 338, F2: (v30 << ((1) & 63))}
                                                                                var v114 int = fi15
                                                                                _ = v113
                                                                                _ = v114
                                                                            case -755:
                                                                                v5 = append(v5, fi15)
                                                                            }
                                                                        }
                                                                        _ = v30
                                                                        _ = v31
                                                                        _ = v38
                                                                        _ = v39
                                                                    case 128:
                                                                        switch (v2 + (v2 + (v2 + (("\n\t" + "hello") + (v2 + "\n\t"))))) {
                                                                        case "hello":
                                                                            var v115 int16 = -1
                                                                            _ = v115
                                                                        }
                                                                    }
                                                                    var v116 uint = 944
                                                                    _ = v116
                                                                } else {
                                                                    if false {
                                                                        var v117 uint16 = 42
                                                                        v118 := GenS11{F0: -1, F1: 246, F2: v2}
                                                                        v4 += 0
                                                                        _ = v117
                                                                        _ = v118
                                                                    }
                                                                    fmt.Println(fi9, v2)
                                                                    fmt.Println(v8, v1)
                                                                }
                                                            }
                                                            var v119 uint = 4294967295
                                                            _ = v119
                                                        }
                                                    }
                                                    var v120 int16 = 1000
                                                    v121 := GenS1{F0: 293, F1: v13}
                                                    _ = v120
                                                    _ = v121
                                                } else {
                                                    v8 += float32(float64(((v8 / (v8 / -2.25)) - v8)))
                                                }
                                                v122 := GenS14{F0: 42, F1: v4, F2: (!true)}
                                                _ = v122
                                            }
                                        }
                                        var v123 uint16 = 42
                                        v8 -= ((((0.1 - v8) - float32(3.14159)) / v8) + v8)
                                        fmt.Println(v8, v4, v2)
                                        if ((^(v13)) < (v13 << ((65535) & 63))) {
                                            v124 := GenS15{F0: 570, F1: (v123 % ((0) | 1)), F2: v4}
                                            var v125 float32 = v8
                                            _ = v124
                                            _ = v125
                                        }
                                        if (42 >= 18446744073709551615) {
                                            v8 += float32(((float64(v8) + 0.1) - (float64(float32((1.5 * float64((1e10 * (1.5 - v8)))))) + 1.5)))
                                            var v126 uint32 = 716
                                            _ = v126
                                        }
                                        var v127 int64 = v4
                                        _ = fi15
                                        _ = v123
                                        _ = v127
                                    }
                                default:
                                    v5 = append(v5, (fi9 << ((618) & 63)))
                                }
                            } else {
                                if len(v5) > 0 {
                                    v5[0] = 32767
                                }
                                v13 = v13
                            }
                            v1 -= (^(v1))
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            v128 := GenS6{F0: 342, F1: (v2 + "foo bar")}
                            switch fi9 {
                            case 352:
                                var v129 uint16 = 255
                                _ = v129
                            case 563:
                                switch 127 {
                                case -282:
                                    for fi130 := 0; fi130 < 13; fi130++ {
                                        v2 = ((v2 + v2) + "hello")
                                        if false {
                                            v2 = v2
                                            switch (fi9 - 128) {
                                            case -2147483648:
                                                var v131 uint32 = 535
                                                _ = v131
                                            case 1:
                                                if true {
                                                    if len(v5) > 0 {
                                                        v5[0] = fi130
                                                    }
                                                    v4 = v4
                                                    if (127 != 1) {
                                                        var v132 string = ((("foo bar" + v2) + v2) + v2)
                                                        _ = v132
                                                    }
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                            case 32767:
                                                if false {
                                                    v8 = float32(float64(v8))
                                                    var v133 int32 = 0
                                                    var v134 int = fi9
                                                    _ = v133
                                                    _ = v134
                                                }
                                            }
                                            if true {
                                                var v135 uint32 = 482
                                                if len(v5) > 0 {
                                                    v5[0] = fi130
                                                }
                                                v8 += 0.0
                                                _ = v135
                                            } else {
                                                v8 += 0.1
                                                v2 = v2
                                                var v136 float32 = (v8 / 3.14159)
                                                _ = v136
                                            }
                                        }
                                        switch (fi130 | (fi9 / (((-(fi9))) | 1))) {
                                        case 1:
                                            if (!((false && true) && false)) {
                                                var v137 uint64 = 1
                                                fmt.Println(v8, v137)
                                                var v138 int16 = -1
                                                _ = v137
                                                _ = v138
                                            }
                                        case 127:
                                            if len(v3) > 0 {
                                                v3[0] = 2147483648
                                            }
                                        case 128:
                                            var v139 int = 32768
                                            _ = v139
                                        }
                                        if (false || true) {
                                            var v140 uint16 = 42
                                            v141 := GenS12{F0: -1, F1: 643}
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v140
                                            _ = v141
                                        }
                                        var v142 int8 = 0
                                        _ = fi130
                                        _ = v142
                                    }
                                case 128:
                                    fmt.Println(v2, v4)
                                case -599:
                                    for fi143 := 0; fi143 < 7; fi143++ {
                                        v144 := GenS16{F0: 127, F1: 9223372036854775807, F2: v4}
                                        var v145 int = fi143
                                        _ = fi143
                                        _ = v144
                                        _ = v145
                                    }
                                default:
                                    var v146 uint16 = 1
                                    _ = v146
                                }
                            }
                            _ = v128
                        }
                        _ = fi9
                        _ = v13
                    }
                    v147 := GenS2{F0: 1, F1: 42}
                    v2 = v2
                    _ = v147
                } else {
                    if (360 >= -190) {
                        for fi148 := 0; fi148 < 9; fi148++ {
                            v149 := GenS17{F0: -1, F1: 298}
                            _ = fi148
                            _ = v149
                        }
                        v8 += (v8 + 0.1)
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(v1, v2)
                }
            }
        case "a":
            var v150 uint8 = v1
            _ = v150
        case "\n\t":
            if (42 > -512) {
                v5 = append(v5, 2147483647)
                v5 = append(v5, -1)
            } else {
                for fi151 := 0; fi151 < 9; fi151++ {
                    fmt.Println(v8, v4)
                    _ = fi151
                }
                switch 738 {
                case -892:
                    v1 = v1
                case -259:
                    if (4294967295 > 1) {
                        if (false && (890 == 411)) {
                            switch 165 {
                            case 338:
                                var v152 string = ("foo bar" + (v2 + v2))
                                _ = v152
                            case 37:
                                fmt.Println(v1)
                            case 32767:
                                v8 = -2.25
                            default:
                                for fi153 := 0; fi153 < 20; fi153++ {
                                    switch fi153 {
                                    case -668:
                                        if (float64(v8) >= 1.5) {
                                            for fi154 := 0; fi154 < 11; fi154++ {
                                                if (!false) {
                                                    v4 = (v4 | v4)
                                                    var v155 float64 = 0.1
                                                    v155 -= ((-2.25 + 0.0) * (v155 + v155))
                                                    _ = v155
                                                }
                                                _ = fi154
                                            }
                                        }
                                    case -226:
                                        var v156 int32 = 1
                                        _ = v156
                                    default:
                                        v8 -= v8
                                    }
                                    if ((3.14159 - (float64(((((1e10 * v8) + (v8 * ((1e10 + (((((1.5 / v8) - 0.0) * 1.5) / v8) - (1.5 + 0.0))) * (3.14159 * (1.5 - 3.14159))))) + v8) + v8)) - 1.5)) == 1.5) {
                                        var v157 uint16 = 1
                                        for fi158 := 0; fi158 < 9; fi158++ {
                                            var v159 int8 = 0
                                            var v160 float64 = 0.0
                                            _ = fi158
                                            _ = v159
                                            _ = v160
                                        }
                                        _ = v157
                                    }
                                    _ = fi153
                                }
                            }
                            var v161 uint16 = 632
                            _ = v161
                        }
                    } else {
                        for fi162 := 0; fi162 < 7; fi162++ {
                            switch fi162 {
                            case 4294967295:
                                fmt.Println(fi162, v8, v4)
                            }
                            if (!(!true)) {
                                v2 = "a"
                            } else {
                                if len(v5) > 0 {
                                    v5[0] = (fi162 / ((-785) | 1))
                                }
                                for fi163 := 0; fi163 < 14; fi163++ {
                                    var v164 int32 = 42
                                    _ = fi163
                                    _ = v164
                                }
                                if true {
                                    var v165 uint32 = 42
                                    v5 = append(v5, fi162)
                                    _ = v165
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = fi162
                                    }
                                }
                            }
                            _ = fi162
                        }
                    }
                case -926:
                    var v166 uint = 0
                    _ = v166
                default:
                    v8 = (v8 + (v8 * 1.5))
                }
                for fi167 := 0; fi167 < 17; fi167++ {
                    v168 := GenS18{F0: v1, F1: (fi167 * -2147483648), F2: v2}
                    _ = fi167
                    _ = v168
                }
            }
        }
        var v169 uint = 18446744073709551615
        _ = v169
        _ = v8
    }
    v2 = (v2 + (v2 + "hello"))
    if false {
        switch 1 {
        case -2147483648:
            var v170 string = ""
            _ = v170
        default:
            v2 = v2
        }
        for fi171 := 0; fi171 < 7; fi171++ {
            switch v2 {
            case "hello":
                fmt.Println(v4, v1, v1)
            }
            if (!(true && (0 > 265))) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            v172 := GenS12{F0: (fi171 / (((fi171 + -1)) | 1)), F1: 51}
            _ = fi171
            _ = v172
        }
        var v173 int8 = 127
        _ = v173
    }
    var v174 int8 = 100
    v175 := GenS1{F0: 2, F1: 2147483647}
    for fi176 := 0; fi176 < 3; fi176++ {
        for fi177 := 0; fi177 < 14; fi177++ {
            v178 := GenS6{F0: 111, F1: v2}
            switch fi177 {
            case 414:
                fmt.Println(fi177, fi176, fi177)
            }
            fmt.Println(v174, v174, fi176)
            _ = fi177
            _ = v178
        }
        if ((v1 % ((v1) | 1)) == v1) {
            v174 += 1
            var v179 uint = 2
            var v180 int32 = 1
            _ = v179
            _ = v180
        }
        v3 = append(v3, 2147483647)
        _ = fi176
    }
    if (("a" + v2) == v2) {
        v2 = v2
        if (289 > 42) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            if (v1 != v1) {
                v1 = v1
            }
            var v181 bool = (false && false)
            _ = v181
        }
    }
    v182 := GenS4{F0: true, F1: 932}
    if true {
        if (0.1 > 0.0) {
            if false {
                _ = len(v5)
                var v183 float32 = -2.25
                _ = v183
            }
            v184 := GenS19{F0: 986, F1: (((-704 == 1) && false) && true), F2: 32767}
            var v185 string = v2
            _ = v184
            _ = v185
        } else {
            v186 := GenS20{F0: (v4 | v4), F1: 895, F2: -9}
            v174 = (v174 ^ (-(v174)))
            for fi187 := 0; fi187 < 14; fi187++ {
                var v188 string = "a"
                _ = fi187
                _ = v188
            }
            _ = v186
        }
        if len(v5) > 0 {
            v5[0] = -1
        }
    }
    var v189 float32 = (-2.25 * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v174:", v174)
    fmt.Println("v189:", v189)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v174
    _ = v175
    _ = v182
    _ = v189
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
