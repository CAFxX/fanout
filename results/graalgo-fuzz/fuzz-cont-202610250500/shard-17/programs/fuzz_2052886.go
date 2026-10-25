package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint; F2 string }
type GenS3 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint8; F1 string; F2 uint32 }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 uint16; F1 int; F2 bool }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int16; F1 int16; F2 uint16 }
type GenS13 struct { F0 string; F1 int8 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int32; F1 uint16; F2 int8 }
type GenS16 struct { F0 uint; F1 int; F2 int8 }
type GenS17 struct { F0 bool; F1 uint; F2 bool }
type GenS18 struct { F0 uint16; F1 string }
type GenS19 struct { F0 int32; F1 bool; F2 uint16 }
type GenS20 struct { F0 uint8; F1 int8 }
type GenS21 struct { F0 int8; F1 int32 }
type GenS22 struct { F0 uint8; F1 int16 }
type GenS23 struct { F0 uint32; F1 string }

func main() {
    var v1 float64 = 3.14159
    var v2 uint = 42
    var v3 uint = (^(v2))
    v4 := []int{4294967295}
    var v5 uint = 1
    v6 := []uint32{2147483648, 1, 223}
    v3 = (^(v5))
    var v7 uint16 = 0
    switch 42 {
    case 797:
        var v8 int64 = 42
        _ = v8
    default:
        var v9 uint64 = 42
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: (-32768 < -32768)}
            if true {
                switch (("a" + "\n\t") + ("hello" + ("a" + "a"))) {
                case "foo bar":
                    if (!(128 > 116)) {
                        var v13 uint16 = (v7 % ((215) | 1))
                        _ = v13
                    } else {
                        if (true || false) {
                            var v14 float64 = float64(float32(0.0))
                            _ = v14
                        } else {
                            if ((!(-237 == 9223372036854775807)) && (169 < uint8(v7))) {
                                if len(v6) > 0 {
                                    v6[0] = 408
                                }
                                v1 = (v1 + v1)
                                var v15 uint = (v5 >> (((v2 / ((v5) | 1))) & 63))
                                v15 += (v15 + v3)
                                if (279 != -2147483648) {
                                    if (96 != 0) {
                                        v4 = append(v4, 632)
                                        fmt.Println(v3, v5, v7)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    } else {
                                        var v16 string = "\n\t"
                                        _ = len(v4)
                                        _ = v16
                                    }
                                    var v17 int64 = 42
                                    _ = v17
                                }
                                var v18 int16 = -186
                                v15 -= 342
                                v1 = ((3.14159 * v1) / 1.5)
                                switch "foo bar" {
                                case "foo bar":
                                    fmt.Println(v7, v1, v5)
                                case "a":
                                    var v19 uint = (v2 & v3)
                                    _ = v19
                                case "hello":
                                    _ = len(v6)
                                }
                                switch (fi11 / ((-9223372036854775808) | 1)) {
                                case 127:
                                    v15 -= (v3 * (^(v5)))
                                    v20 := GenS1{F0: (-668 == (v18 & (v18 | v18)))}
                                    if (true || (!false)) {
                                        var v21 bool = false
                                        var v22 float64 = 0.0
                                        _ = v21
                                        _ = v22
                                    } else {
                                        var v23 string = ("a" + "foo bar")
                                        v24 := GenS2{F0: 184, F1: 4294967295, F2: v23}
                                        switch v23 {
                                        case "":
                                            var v25 float32 = 1e10
                                            _ = v25
                                        default:
                                            _ = len(v4)
                                        }
                                        if false {
                                            v7 = v7
                                            if true {
                                                var v26 uint64 = 4294967295
                                                var v27 uint16 = (v7 >> (((v15 << (((v5 / ((v15) | 1))) & 63))) & 63))
                                                if true {
                                                    var v28 uint64 = v26
                                                    _ = v28
                                                } else {
                                                    switch v23 {
                                                    case "\n\t":
                                                        var v29 int32 = 42
                                                        _ = v29
                                                    }
                                                    fmt.Println(v26, fi10)
                                                    if (false && false) {
                                                        if false {
                                                            v30 := GenS2{F0: 0, F1: v5, F2: "foo bar"}
                                                            v31 := GenS3{F0: (fi11 == (fi11 | fi10)), F1: v7, F2: 2}
                                                            _ = v30
                                                            _ = v31
                                                        } else {
                                                            fmt.Println(fi10, v18)
                                                        }
                                                        switch (((v23 + (v23 + (("foo bar" + v23) + ("hello" + v23)))) + "\n\t") + (("a" + v23) + v23)) {
                                                        case "\n\t":
                                                            v32 := GenS4{F0: 4294967295}
                                                            _ = v32
                                                        case "hello":
                                                            var v33 uint32 = 2147483647
                                                            _ = v33
                                                        case "foo bar":
                                                            switch (v23 + "") {
                                                            case "\n\t":
                                                                var v34 int16 = (v18 ^ (v18 - -831))
                                                                _ = v34
                                                            default:
                                                                v35 := GenS5{F0: 255, F1: (v26 * 4294967295), F2: 42}
                                                                v4 = append(v4, fi10)
                                                                switch v23 {
                                                                case "hello":
                                                                    if ((v18 & (v18 % (((v18 % (((v18 * 32767)) | 1))) | 1))) != v18) {
                                                                        fmt.Println(v15, v3)
                                                                        switch fi10 {
                                                                        case 4294967295:
                                                                            v23 = ((v23 + v23) + ("" + v23))
                                                                        case 0:
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                        case 128:
                                                                            v27 -= 32767
                                                                        default:
                                                                            var v36 float32 = 1.5
                                                                            _ = v36
                                                                        }
                                                                    }
                                                                    switch v23 {
                                                                    case "":
                                                                        if false {
                                                                            v18 = (v18 + 42)
                                                                        } else {
                                                                            v37 := GenS6{F0: 215}
                                                                            _ = v37
                                                                        }
                                                                    case "foo bar":
                                                                        v38 := GenS1{F0: (!((v2 ^ (v2 | 18446744073709551615)) >= 569))}
                                                                        _ = v38
                                                                    case "hello":
                                                                        if true {
                                                                            if ((1 == fi10) || (true && true)) {
                                                                                _ = len(v6)
                                                                                v6 = append(v6, 2)
                                                                                var v39 uint = v2
                                                                                _ = v39
                                                                            } else {
                                                                                var v40 string = (v23 + v23)
                                                                                var v41 string = (v40 + (("a" + v23) + v23))
                                                                                switch fi10 {
                                                                                case 32768:
                                                                                    fmt.Println(v5, v27)
                                                                                    v4 = append(v4, int(v7))
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                default:
                                                                                    v4 = append(v4, (-(fi10)))
                                                                                }
                                                                                _ = v40
                                                                                _ = v41
                                                                            }
                                                                            var v42 float64 = -2.25
                                                                            if (1 > 0) {
                                                                                if (4294967295 < 65535) {
                                                                                    v6 = append(v6, 2147483648)
                                                                                    v5 += 65535
                                                                                    switch 2147483647 {
                                                                                    case 2147483647:
                                                                                        var v43 int32 = -2147483648
                                                                                        _ = v43
                                                                                    case -9223372036854775808:
                                                                                        if true {
                                                                                            var v44 int32 = int32(fi11)
                                                                                            if true {
                                                                                                if ((v1 - v1) <= 1.5) {
                                                                                                    var v45 int32 = v44
                                                                                                    _ = v45
                                                                                                }
                                                                                            } else {
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                                v27 -= (v7 * v7)
                                                                                            }
                                                                                            v26 -= (v26 + 4294967295)
                                                                                            _ = v44
                                                                                        }
                                                                                    case 32767:
                                                                                        v4 = append(v4, (-(fi10)))
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                switch fi10 {
                                                                                case -1:
                                                                                    if (false || (v5 < (v5 % (((v3 + v15)) | 1)))) {
                                                                                        if (false && ((2 == 2147483648) || true)) {
                                                                                            _ = len(v4)
                                                                                            var v46 uint32 = 287
                                                                                            if false {
                                                                                                v42 = 3.14159
                                                                                                var v47 uint16 = v7
                                                                                                var v48 float32 = 3.14159
                                                                                                _ = v47
                                                                                                _ = v48
                                                                                            } else {
                                                                                                v23 = v23
                                                                                            }
                                                                                            _ = v46
                                                                                        } else {
                                                                                            v49 := GenS7{F0: ""}
                                                                                            v15 += uint(v7)
                                                                                            if false {
                                                                                                v18 = v18
                                                                                                if ((4294967295 == (v3 / (((^(v2))) | 1))) && true) {
                                                                                                    if (!(true && true)) {
                                                                                                        fmt.Println(v27, v27, v18)
                                                                                                        v15 = uint(v7)
                                                                                                        switch (fi11 - fi11) {
                                                                                                        case -728:
                                                                                                            v5 = 698
                                                                                                        case 128:
                                                                                                            v2 -= 2
                                                                                                        default:
                                                                                                            var v50 int32 = 42
                                                                                                            _ = v50
                                                                                                        }
                                                                                                    } else {
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = len(v6)
                                                                                            }
                                                                                            _ = v49
                                                                                        }
                                                                                        v51 := GenS3{F0: (fi10 <= fi10), F1: v27, F2: 326}
                                                                                        if false {
                                                                                            v52 := GenS8{F0: 0, F1: v23, F2: 4294967295}
                                                                                            if false {
                                                                                                _ = len(v4)
                                                                                            }
                                                                                            if ((765 == 2147483647) && true) {
                                                                                                switch v23 {
                                                                                                case "\n\t":
                                                                                                    v3 = (^(v3))
                                                                                                case "foo bar":
                                                                                                    v15 = (v3 + v5)
                                                                                                case "":
                                                                                                    if len(v6) > 0 {
                                                                                                        v6[0] = 65535
                                                                                                    }
                                                                                                }
                                                                                                if len(v6) >= 2 {
                                                                                                    v6 = v6[1:len(v6)]
                                                                                                }
                                                                                            } else {
                                                                                                _ = len(v6)
                                                                                                v53 := GenS9{F0: int16(v7), F1: 584}
                                                                                                var v54 int8 = 84
                                                                                                _ = v53
                                                                                                _ = v54
                                                                                            }
                                                                                            v55 := GenS2{F0: 42, F1: (v3 * (v15 >> ((255) & 63))), F2: v23}
                                                                                            v23 = ((("hello" + v23) + ("\n\t" + v23)) + "")
                                                                                            if true {
                                                                                                v1 -= -2.25
                                                                                                var v56 uint = v5
                                                                                                var v57 string = "\n\t"
                                                                                                fmt.Println(v27, v23, v15)
                                                                                                fmt.Println(v57)
                                                                                                _ = len(v4)
                                                                                                _ = v56
                                                                                                _ = v57
                                                                                            } else {
                                                                                                v6 = append(v6, 2)
                                                                                            }
                                                                                            if (127 < 127) {
                                                                                                v58 := GenS6{F0: uint8(v27)}
                                                                                                _ = v58
                                                                                            }
                                                                                            _ = v52
                                                                                            _ = v55
                                                                                        }
                                                                                        _ = v51
                                                                                    } else {
                                                                                        v23 = ((v23 + v23) + v23)
                                                                                        v59 := GenS3{F0: (-2.25 >= -2.25), F1: (v27 / ((293) | 1)), F2: 556}
                                                                                        _ = len(v6)
                                                                                        _ = v59
                                                                                    }
                                                                                case 4294967295:
                                                                                    switch ("\n\t" + v23) {
                                                                                    case "a":
                                                                                        v18 = v18
                                                                                    case "foo bar":
                                                                                        var v60 int = fi11
                                                                                        _ = v60
                                                                                    default:
                                                                                        if false {
                                                                                            var v61 float64 = -2.25
                                                                                            var v62 int16 = v18
                                                                                            _ = v61
                                                                                            _ = v62
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v63 := GenS2{F0: 1, F1: v15, F2: v23}
                                                                                v1 += (float64((float32((((-2.25 + 0.1) * ((v42 / v42) - v42)) - (3.14159 + float64(-2.25)))) + ((1.5 * 0.0) * -2.25))) * (1e10 - 1e10))
                                                                                _ = v63
                                                                            }
                                                                            _ = v42
                                                                        }
                                                                    default:
                                                                        switch v23 {
                                                                        case "\n\t":
                                                                            v7 += (^(v7))
                                                                        case "hello":
                                                                            v18 += -1
                                                                        case "":
                                                                            var v64 string = "hello"
                                                                            _ = v64
                                                                        }
                                                                    }
                                                                case "":
                                                                    v65 := GenS10{F0: (v7 ^ 42), F1: fi11, F2: false}
                                                                    _ = v65
                                                                case "\n\t":
                                                                    var v66 int = (fi11 ^ 32768)
                                                                    _ = v66
                                                                }
                                                                _ = v35
                                                            }
                                                        }
                                                        if (false || (true || true)) {
                                                            if false {
                                                                v27 = (v7 % (((v7 % (((v7 * (^(v7)))) | 1))) | 1))
                                                            }
                                                            if false {
                                                                v4 = append(v4, (fi10 & fi10))
                                                                var v67 float32 = (-2.25 - ((0.0 + -2.25) / 1.5))
                                                                _ = v67
                                                            } else {
                                                                var v68 int32 = 2147483647
                                                                var v69 uint32 = 0
                                                                var v70 uint64 = 2
                                                                v4 = append(v4, (fi10 >> ((1) & 63)))
                                                                var v71 uint8 = 255
                                                                v71 -= (^(v71))
                                                                _ = v68
                                                                _ = v69
                                                                _ = v70
                                                                _ = v71
                                                            }
                                                        } else {
                                                            switch ((v23 + "") + (v23 + (v23 + v23))) {
                                                            case "":
                                                                fmt.Println(v7)
                                                            case "a":
                                                                fmt.Println(v15, v26)
                                                            default:
                                                                var v72 float64 = v1
                                                                _ = v72
                                                            }
                                                            var v73 uint64 = v26
                                                            if len(v4) > 0 {
                                                                v4[0] = 299
                                                            }
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            v3 = v15
                                                            _ = v73
                                                        }
                                                    }
                                                    v74 := GenS11{F0: 0}
                                                    _ = v74
                                                }
                                                _ = v26
                                                _ = v27
                                            }
                                            v7 += v7
                                        } else {
                                            if (true && ((v18 + 0) == (v18 & (v18 - v18)))) {
                                                if ((-2.25 * float32((v1 - (1.5 * 1.5)))) != (0.0 - 3.14159)) {
                                                    fmt.Println(fi10, fi11)
                                                    if true {
                                                        v2 += 0
                                                        if (!(0 >= v18)) {
                                                            fmt.Println(v23, v5)
                                                            if false {
                                                                v15 -= 827
                                                                if ((true && false) || true) {
                                                                    var v75 int8 = -1
                                                                    _ = v75
                                                                } else {
                                                                    switch 954 {
                                                                    case 1:
                                                                        if false {
                                                                            switch ("a" + ("hello" + (v23 + ("a" + v23)))) {
                                                                            case "hello":
                                                                                v5 += (v15 << ((4294967295) & 63))
                                                                            case "\n\t":
                                                                                v18 = v18
                                                                            case "foo bar":
                                                                                if (-9223372036854775808 >= 1) {
                                                                                    switch v23 {
                                                                                    case "a":
                                                                                        var v76 bool = false
                                                                                        if len(v6) >= 2 {
                                                                                            v6 = v6[1:len(v6)]
                                                                                        }
                                                                                        _ = v76
                                                                                    case "":
                                                                                        var v77 float32 = 3.14159
                                                                                        _ = v77
                                                                                    }
                                                                                    v18 += (-(v18))
                                                                                }
                                                                            default:
                                                                                v78 := GenS12{F0: (v18 & (v18 | v18)), F1: 241, F2: (v7 / ((v7) | 1))}
                                                                                _ = v78
                                                                            }
                                                                            if true {
                                                                                var v79 string = v23
                                                                                v80 := GenS13{F0: "\n\t", F1: 60}
                                                                                _ = v79
                                                                                _ = v80
                                                                            } else {
                                                                                switch v23 {
                                                                                case "":
                                                                                    v3 -= v2
                                                                                default:
                                                                                    if false {
                                                                                        v81 := GenS13{F0: "", F1: 0}
                                                                                        fmt.Println(v3, v18, v23)
                                                                                        v82 := GenS14{F0: 0}
                                                                                        _ = v81
                                                                                        _ = v82
                                                                                    }
                                                                                }
                                                                                v5 += (v15 ^ (v15 ^ v15))
                                                                            }
                                                                            var v83 uint32 = 65535
                                                                            _ = v83
                                                                        }
                                                                    case 32768:
                                                                        var v84 uint8 = 127
                                                                        _ = v84
                                                                    case 561:
                                                                        v7 += v7
                                                                    }
                                                                    if true {
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            if (true || (true || false)) {
                                                                                var v85 int16 = 1000
                                                                                var v86 uint = 0
                                                                                _ = v85
                                                                                _ = v86
                                                                            }
                                                                        }
                                                                        switch (v23 + "hello") {
                                                                        case "a":
                                                                            var v87 uint8 = 143
                                                                            _ = v87
                                                                        case "\n\t":
                                                                            if (v18 < 646) {
                                                                                var v88 uint8 = 124
                                                                                if true {
                                                                                    var v89 int16 = (^(v18))
                                                                                    if (false || (!true)) {
                                                                                        v4 = append(v4, (-(fi11)))
                                                                                    } else {
                                                                                        switch (-(fi11)) {
                                                                                        case 32768:
                                                                                            if false {
                                                                                                var v90 uint64 = 4294967295
                                                                                                v89 += v18
                                                                                                _ = v90
                                                                                            }
                                                                                        case -1:
                                                                                            v91 := GenS2{F0: (v88 >> (((v5 % ((v2) | 1))) & 63)), F1: 542, F2: v23}
                                                                                            _ = v91
                                                                                        }
                                                                                        v4 = append(v4, 9223372036854775807)
                                                                                    }
                                                                                    _ = v89
                                                                                }
                                                                                _ = v88
                                                                            } else {
                                                                                if true {
                                                                                    if true {
                                                                                        _ = len(v6)
                                                                                    } else {
                                                                                        v92 := GenS11{F0: -2147483648}
                                                                                        v3 = v5
                                                                                        fmt.Println(v3, v3, v7)
                                                                                        _ = v92
                                                                                    }
                                                                                    v93 := GenS10{F0: (v7 * (v7 | (v7 | (v7 & v7)))), F1: 1, F2: true}
                                                                                    fmt.Println(v5, v5)
                                                                                    fmt.Println(v7)
                                                                                    v5 -= v5
                                                                                    var v94 float64 = ((v1 + (v1 / v1)) + v1)
                                                                                    v18 += (v18 / ((475) | 1))
                                                                                    _ = v93
                                                                                    _ = v94
                                                                                }
                                                                                if (100 != 127) {
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = 42
                                                                                    }
                                                                                    var v95 int64 = 455
                                                                                    _ = v95
                                                                                }
                                                                                fmt.Println(v3)
                                                                            }
                                                                        case "":
                                                                            var v96 int = 32768
                                                                            _ = v96
                                                                        default:
                                                                            v23 = v23
                                                                        }
                                                                    } else {
                                                                        v1 -= 3.14159
                                                                        v97 := GenS1{F0: (2147483647 >= 0)}
                                                                        var v98 int16 = v18
                                                                        _ = v97
                                                                        _ = v98
                                                                    }
                                                                }
                                                                var v99 uint16 = 0
                                                                _ = v99
                                                            } else {
                                                                var v100 int8 = 0
                                                                if false {
                                                                    v6 = append(v6, 404)
                                                                    v101 := GenS1{F0: (!(!false))}
                                                                    _ = v101
                                                                } else {
                                                                    v18 = (^(v18))
                                                                    v100 += 42
                                                                    var v102 uint64 = 18446744073709551615
                                                                    fmt.Println(fi11, v1, v23)
                                                                    if true {
                                                                        v102 += (v102 % ((v102) | 1))
                                                                    }
                                                                    _ = v102
                                                                }
                                                                _ = v100
                                                            }
                                                            v18 = -686
                                                        }
                                                    }
                                                }
                                                v23 = v23
                                                fmt.Println(v23, fi11, fi11)
                                            }
                                            if false {
                                                var v103 uint8 = 0
                                                v23 = v23
                                                v103 = v103
                                                v104 := GenS15{F0: -919, F1: (v7 & v7), F2: -1}
                                                _ = v103
                                                _ = v104
                                            }
                                            v105 := GenS16{F0: (v2 / (((v2 >> (((v5 - (v3 % (((v15 & v15)) | 1)))) & 63))) | 1)), F1: int(v18), F2: -10}
                                            _ = v105
                                        }
                                        v2 = (^(v2))
                                        _ = len(v6)
                                        _ = v23
                                        _ = v24
                                    }
                                    _ = v20
                                }
                                if false {
                                    var v106 bool = (2 < 2147483648)
                                    v1 -= v1
                                    switch "foo bar" {
                                    case "hello":
                                        var v107 float32 = (-2.25 / 1.5)
                                        _ = v107
                                    case "\n\t":
                                        if ((1 == 142) && true) {
                                            v106 = v106
                                            v108 := GenS1{F0: (((v1 > v1) && (42 >= 786)) && (float64(0.0) <= 1e10))}
                                            _ = v108
                                        } else {
                                            if (((("" + ((("hello" + "foo bar") + "foo bar") + "a")) + "\n\t") + "a") <= "foo bar") {
                                                fmt.Println(v106, v18, v3)
                                                v109 := GenS17{F0: true, F1: v5, F2: v106}
                                                v15 = v3
                                                _ = v109
                                            } else {
                                                var v110 uint = v2
                                                var v111 float32 = 1e10
                                                v111 = (v111 - v111)
                                                v112 := GenS8{F0: 42, F1: "\n\t", F2: 696}
                                                _ = v110
                                                _ = v111
                                                _ = v112
                                            }
                                            switch "a" {
                                            case "a":
                                                var v113 uint = (v15 / (((v15 * 206)) | 1))
                                                _ = v113
                                            case "":
                                                v114 := GenS9{F0: 1, F1: -705}
                                                _ = v114
                                            case "\n\t":
                                                fmt.Println(v3, fi11)
                                            }
                                        }
                                    }
                                    _ = v106
                                }
                                var v115 float64 = (1e10 - v1)
                                _ = v115
                                _ = v15
                                _ = v18
                            } else {
                                v1 += v1
                                fmt.Println(fi11, v5, v7)
                                var v116 uint64 = 0
                                _ = v116
                            }
                            v117 := GenS18{F0: 255, F1: "foo bar"}
                            v1 = (v1 / v1)
                            if false {
                                var v118 int16 = 972
                                var v119 uint8 = 0
                                fmt.Println(v119, v119, fi10)
                                _ = v118
                                _ = v119
                            }
                            _ = v117
                        }
                    }
                case "a":
                    fmt.Println(v1, v5, v3)
                case "":
                    v5 += v3
                default:
                    v120 := GenS2{F0: 1, F1: v2, F2: (("\n\t" + "a") + ((("\n\t" + ("a" + "a")) + ("" + "")) + ("" + "foo bar")))}
                    _ = v120
                }
                var v121 float32 = 0.0
                _ = v121
            } else {
                if true {
                    v7 = v7
                    switch fi11 {
                    case 2147483647:
                        v122 := GenS13{F0: "hello", F1: 127}
                        _ = v122
                    case 793:
                        var v123 string = "foo bar"
                        _ = v123
                    default:
                        var v124 bool = (((-2.25 + (float64((((0.0 - (-2.25 * float32(float64(0.1)))) + (float32(float64(3.14159)) * (1.5 + ((3.14159 * float32(v1)) * 1.5)))) / float32(1.5))) / float64(-2.25))) / v1) == 1.5)
                        _ = v124
                    }
                    var v125 int8 = 1
                    _ = v125
                }
                v2 -= (^(v5))
                switch "a" {
                case "hello":
                    switch "a" {
                    case "a":
                        if false {
                            v1 += v1
                            fmt.Println(v3, fi11)
                            v5 = v5
                        }
                    case "\n\t":
                        var v126 int64 = -1
                        _ = v126
                    case "":
                        var v127 uint = (^(v2))
                        _ = v127
                    }
                case "":
                    fmt.Println(v7, v1)
                case "a":
                    v5 += (v3 ^ v3)
                    v3 += 42
                    _ = len(v4)
                default:
                    var v128 uint = v2
                    _ = v128
                }
            }
            _ = fi11
            _ = v12
        }
        var v129 int32 = 1
        for fi130 := 0; fi130 < 5; fi130++ {
            var v131 int16 = 1000
            if true {
                v132 := GenS19{F0: (-(v129)), F1: (false && false), F2: (v7 ^ 65535)}
                v1 = ((float64(1.5) - v1) / 1.0)
                if true {
                    v133 := GenS20{F0: 104, F1: -1}
                    _ = v133
                } else {
                    v5 -= 0
                    if ((!true) && (v5 >= (v2 << ((v5) & 63)))) {
                        if false {
                            var v134 bool = (v129 >= 0)
                            _ = v134
                        }
                        v135 := GenS20{F0: 214, F1: 1}
                        var v136 float64 = v1
                        _ = v135
                        _ = v136
                    } else {
                        var v137 string = "a"
                        _ = v137
                    }
                    var v138 uint = 4294967295
                    _ = v138
                }
                var v139 float64 = v1
                _ = v132
                _ = v139
            } else {
                var v140 uint = (^(v2))
                if false {
                    var v141 float32 = 0.0
                    _ = v141
                } else {
                    fmt.Println(v7, v2, v2)
                    v142 := GenS14{F0: 1}
                    _ = v142
                }
                if len(v6) > 0 {
                    v6[0] = uint32(v2)
                }
                _ = v140
            }
            _ = fi130
            _ = v131
        }
        _ = fi10
        _ = v129
    }
    v4 = append(v4, 127)
    v143 := GenS11{F0: -46}
    switch 32768 {
    case 812:
        switch "\n\t" {
        case "foo bar":
            if (2 > 2) {
                if true {
                    v3 -= (v5 * v3)
                }
                var v144 int32 = -460
                _ = v144
            }
        }
    case -488:
        for fi145 := 0; fi145 < 20; fi145++ {
            var v146 float64 = (v1 / -2.25)
            switch ("" + (("\n\t" + "a") + "\n\t")) {
            case "\n\t":
                var v147 string = ("\n\t" + "")
                _ = v147
            case "hello":
                fmt.Println(v1, v3)
            default:
                v6 = append(v6, 2147483647)
            }
            if len(v4) > 0 {
                v4[0] = (fi145 << ((uint(v7)) & 63))
            }
            _ = fi145
            _ = v146
        }
    default:
        for fi148 := 0; fi148 < 7; fi148++ {
            v149 := GenS21{F0: 1, F1: 1}
            var v150 uint8 = 127
            _ = fi148
            _ = v149
            _ = v150
        }
    }
    v151 := GenS22{F0: 255, F1: -32768}
    if false {
        var v152 int64 = 0
        _ = v152
    } else {
        for fi153 := 0; fi153 < 1; fi153++ {
            v154 := GenS23{F0: 98, F1: ""}
            fmt.Println(v7, v2, v2)
            _ = fi153
            _ = v154
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v5, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v143
    _ = v151
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
