package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint16; F1 int64; F2 int64 }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS6 struct { F0 bool; F1 bool; F2 bool }
type GenS7 struct { F0 int32; F1 int32; F2 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 bool; F1 uint16; F2 int }
type GenS11 struct { F0 uint; F1 bool; F2 uint32 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint; F1 int16; F2 uint16 }
type GenS14 struct { F0 uint; F1 int16 }
type GenS15 struct { F0 int32; F1 int; F2 int32 }
type GenS16 struct { F0 int32; F1 bool }
type GenS17 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS18 struct { F0 uint64; F1 int64; F2 int }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 int8; F1 int; F2 string }
type GenS21 struct { F0 int64; F1 int32; F2 int16 }
type GenS22 struct { F0 uint64; F1 bool; F2 bool }
type GenS23 struct { F0 bool; F1 int8; F2 bool }
type GenS24 struct { F0 int16; F1 uint8; F2 int8 }
type GenS25 struct { F0 string }
func genFn0(p0 uint16, p1 uint8, p2 int) uint32 {
    v6 := []int{p2, (p2 % ((32767) | 1)), p2, p2}
    var v7 uint8 = p1
    if (109 <= 18446744073709551615) {
        p1 = (v7 << ((2) & 63))
    } else {
        v8 := GenS1{F0: (v7 + p1)}
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            if false {
                var v11 uint = 4294967295
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v12 int8 = -128
                _ = v11
                _ = v12
            } else {
                fmt.Println(p2)
                if (false && (!(!false))) {
                    var v13 int8 = -1
                    switch (fi9 >> ((18446744073709551615) & 63)) {
                    case 32767:
                        var v14 float32 = -2.25
                        _ = v14
                    case 42:
                        if false {
                            switch -9223372036854775808 {
                            case 4294967295:
                                switch "" {
                                case "a":
                                    if true {
                                        v15 := GenS1{F0: (^(v7))}
                                        if true {
                                            if ((-(p2)) >= (-(fi10))) {
                                                var v16 int16 = 42
                                                _ = v16
                                            }
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        fmt.Println(p2, fi10)
                                        _ = v15
                                    }
                                case "":
                                    if false {
                                        v13 = (v13 * v13)
                                        fmt.Println(p0, p0, p0)
                                        switch "foo bar" {
                                        case "hello":
                                            if true {
                                                v7 = (p1 * 117)
                                                if (1 > -9223372036854775808) {
                                                    if (!false) {
                                                        fmt.Println(v7, fi10)
                                                        var v17 int16 = int16(p0)
                                                        _ = v17
                                                    } else {
                                                        v18 := GenS2{F0: 1}
                                                        if false {
                                                            var v19 uint = 1
                                                            switch (fi10 + (p2 * (^(p2)))) {
                                                            case 128:
                                                                v20 := GenS3{F0: v7, F1: 32767}
                                                                _ = v20
                                                            default:
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                            }
                                                            if false {
                                                                v21 := GenS4{F0: (p0 - p0), F1: -1, F2: 104}
                                                                _ = v21
                                                            }
                                                            fmt.Println(p2)
                                                            _ = v19
                                                        } else {
                                                            if (!true) {
                                                                p0 = 250
                                                                fmt.Println(v13, p0, p1)
                                                            } else {
                                                                v6 = append(v6, (p2 * 0))
                                                            }
                                                            switch "" {
                                                            case "hello":
                                                                _ = len(v6)
                                                            }
                                                        }
                                                        v22 := GenS3{F0: (v7 + p1), F1: 1000}
                                                        v23 := GenS4{F0: uint16(p1), F1: 9223372036854775807, F2: 0}
                                                        if (3.14159 <= 0.1) {
                                                            var v24 int16 = 1000
                                                            v6 = append(v6, (-(fi10)))
                                                            _ = v24
                                                        } else {
                                                            switch "" {
                                                            case "":
                                                                v7 = (p1 % ((215) | 1))
                                                            default:
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                            }
                                                            fmt.Println(p2, v7, v13)
                                                            if len(v6) > 0 {
                                                                v6[0] = p2
                                                            }
                                                        }
                                                        _ = v18
                                                        _ = v22
                                                        _ = v23
                                                    }
                                                } else {
                                                    switch 42 {
                                                    case -242:
                                                        p0 -= (^(p0))
                                                    case 32767:
                                                        var v25 uint64 = 1
                                                        _ = v25
                                                    default:
                                                        switch (fi9 % (((-(p2))) | 1)) {
                                                        case 42:
                                                            switch "" {
                                                            case "":
                                                                var v26 int32 = 689
                                                                _ = v26
                                                            default:
                                                                v27 := GenS5{F0: 891, F1: (p0 - p0), F2: 128}
                                                                _ = v27
                                                            }
                                                        case 0:
                                                            v28 := GenS6{F0: (0 < 1), F1: (true && (!(false || ((!(!true)) && (4294967295 >= 18446744073709551615))))), F2: (-128 == (v13 << ((0) & 63)))}
                                                            _ = v28
                                                        default:
                                                            v29 := GenS7{F0: -869, F1: -647, F2: (^(p0))}
                                                            _ = v29
                                                        }
                                                    }
                                                    var v30 int16 = 1
                                                    p2 = (fi9 >> ((273) & 63))
                                                    _ = v30
                                                }
                                            }
                                        case "foo bar":
                                            fmt.Println(p0, fi10, v13)
                                        case "\n\t":
                                            if len(v6) > 0 {
                                                v6[0] = 166
                                            }
                                        }
                                    }
                                case "foo bar":
                                    fmt.Println(p0, v7)
                                }
                            case 128:
                                var v31 string = ((((("\n\t" + "hello") + "a") + "hello") + "a") + "foo bar")
                                _ = v31
                            default:
                                if true {
                                    switch fi10 {
                                    case 1:
                                        if (!(135 == (v7 >> ((18446744073709551615) & 63)))) {
                                            var v32 int64 = 42
                                            var v33 uint = 160
                                            _ = v32
                                            _ = v33
                                        } else {
                                            v7 += 225
                                            if true {
                                                fmt.Println(fi9, p2, fi10)
                                                var v34 int = (^(fi10))
                                                switch "a" {
                                                case "":
                                                    var v35 uint32 = 228
                                                    _ = v35
                                                case "foo bar":
                                                    switch 4294967295 {
                                                    case -303:
                                                        var v36 uint16 = p0
                                                        _ = v36
                                                    case -9223372036854775808:
                                                        var v37 float32 = (-2.25 - (((-2.25 * (-2.25 - ((0.0 + (float32(float64((3.14159 / 0.1))) / 1.0)) * 0.0))) * 1.5) - float32(1.5)))
                                                        _ = v37
                                                    default:
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    }
                                                case "\n\t":
                                                    switch (p2 >> ((839) & 63)) {
                                                    case 674:
                                                        fmt.Println(p1)
                                                        p2 += fi9
                                                        fmt.Println(v34, p1)
                                                    }
                                                }
                                                switch (("" + ("hello" + "hello")) + "foo bar") {
                                                case "a":
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                case "foo bar":
                                                    v38 := GenS1{F0: v7}
                                                    _ = v38
                                                }
                                                var v39 int16 = 951
                                                _ = v34
                                                _ = v39
                                            } else {
                                                if true {
                                                    var v40 bool = (1 != 42)
                                                    switch ("hello" + "a") {
                                                    case "foo bar":
                                                        v6 = append(v6, (p2 - (p2 / (((-(p2))) | 1))))
                                                        if ((v40 || true) || ((v7 | (v7 * p1)) < 2)) {
                                                            v40 = (((0.1 - -2.25) + ((0.0 + 0.0) - 0.1)) != 1e10)
                                                        } else {
                                                            v13 -= -128
                                                            var v41 float64 = -2.25
                                                            p1 = (^(v7))
                                                            fmt.Println(fi9, v13, v7)
                                                            p2 = fi9
                                                            _ = v41
                                                        }
                                                    default:
                                                        if v40 {
                                                            p0 -= (p0 << ((964) & 63))
                                                            if ((v40 && v40) || false) {
                                                                fmt.Println(p2, p0)
                                                                var v42 float64 = 0.1
                                                                v43 := GenS8{F0: 42}
                                                                if ((216 >= 884) && true) {
                                                                    if v40 {
                                                                        var v44 float64 = (float64(0.1) - -2.25)
                                                                        fmt.Println(fi9, fi10)
                                                                        var v45 int64 = -9223372036854775808
                                                                        _ = v44
                                                                        _ = v45
                                                                    }
                                                                    v7 = p1
                                                                }
                                                                if v40 {
                                                                    switch fi9 {
                                                                    case 2147483647:
                                                                        v46 := GenS9{F0: 32767}
                                                                        _ = v46
                                                                    case -831:
                                                                        var v47 int16 = -1
                                                                        _ = v47
                                                                    default:
                                                                        v48 := GenS10{F0: false, F1: p0, F2: -2147483648}
                                                                        _ = v48
                                                                    }
                                                                    var v49 int8 = (-(v13))
                                                                    v6 = append(v6, (fi9 - (fi9 >> ((2) & 63))))
                                                                    switch 16 {
                                                                    case -214:
                                                                        fmt.Println(v7, fi9)
                                                                    case 32768:
                                                                        v6 = append(v6, p2)
                                                                    case 0:
                                                                        v50 := GenS11{F0: 2, F1: false, F2: 2147483648}
                                                                        _ = v50
                                                                    default:
                                                                        var v51 uint = 0
                                                                        _ = v51
                                                                    }
                                                                    _ = v49
                                                                }
                                                                fmt.Println(p2, fi9, p0)
                                                                var v52 uint = 0
                                                                _ = v42
                                                                _ = v43
                                                                _ = v52
                                                            } else {
                                                                var v53 float32 = (-2.25 * 0.1)
                                                                _ = v53
                                                            }
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            var v54 float32 = 1e10
                                                            _ = v54
                                                        } else {
                                                            p1 += (v7 % ((v7) | 1))
                                                            v7 = v7
                                                            fmt.Println(p0, fi9, fi9)
                                                        }
                                                    }
                                                    _ = v40
                                                } else {
                                                    if true {
                                                        if len(v6) > 0 {
                                                            v6[0] = 367
                                                        }
                                                        v55 := GenS3{F0: p1, F1: 822}
                                                        switch -982 {
                                                        case 32767:
                                                            if (!false) {
                                                                var v56 bool = false
                                                                v13 += 0
                                                                v57 := GenS12{F0: false}
                                                                if v56 {
                                                                    p2 -= fi9
                                                                    fmt.Println(p2, p0)
                                                                } else {
                                                                    if false {
                                                                        if (true || (42 != 0)) {
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            if v56 {
                                                                                v7 = v7
                                                                                var v58 bool = (v56 || false)
                                                                                v59 := GenS13{F0: 42, F1: 1000, F2: 666}
                                                                                _ = v58
                                                                                _ = v59
                                                                            } else {
                                                                                if (v56 || ((v56 || false) || (0 >= -278))) {
                                                                                    var v60 int8 = (^(v13))
                                                                                    var v61 int64 = -543
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                } else {
                                                                                    v62 := GenS9{F0: -164}
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                    var v63 bool = (!(false && ((!((1e10 / (((-2.25 / (0.1 / ((3.14159 / 1.0) + (1.5 * -2.25)))) + (float32(float64(((0.1 - 0.1) + float32((((float64(float32(float64((-2.25 * -2.25)))) + 0.0) + (0.1 / ((float64(3.14159) + (-2.25 + 1e10)) - 0.0))) / 3.14159))))) - 1.5)) * ((0.1 - (3.14159 + float32((-2.25 * float64(1e10))))) + -2.25))) < 3.14159)) && (((0.0 - ((1e10 * 1e10) / (float64(0.0) + 1e10))) * float64(0.0)) > 3.14159))))
                                                                                    _ = v62
                                                                                    _ = v63
                                                                                }
                                                                                var v64 int16 = int16(v7)
                                                                                var v65 uint16 = (p0 / (((p0 - p0)) | 1))
                                                                                _ = v64
                                                                                _ = v65
                                                                            }
                                                                            if v56 {
                                                                                switch ("" + ("hello" + ("foo bar" + (("foo bar" + "\n\t") + "\n\t")))) {
                                                                                case "":
                                                                                    v66 := GenS9{F0: 42}
                                                                                    _ = v66
                                                                                default:
                                                                                    v7 = (v7 % (((^(v7))) | 1))
                                                                                }
                                                                                if v56 {
                                                                                    switch 0 {
                                                                                    case 127:
                                                                                        _ = len(v6)
                                                                                    default:
                                                                                        if (true || v56) {
                                                                                            var v67 uint = 148
                                                                                            v67 = 442
                                                                                            v6 = append(v6, (fi10 * fi10))
                                                                                            switch p2 {
                                                                                            case 32767:
                                                                                                v68 := GenS1{F0: p1}
                                                                                                _ = v68
                                                                                            default:
                                                                                                if false {
                                                                                                    if ((float32(0.0) * float32((-2.25 - 0.0))) > 0.0) {
                                                                                                        p2 -= (-(fi10))
                                                                                                    }
                                                                                                    _ = len(v6)
                                                                                                } else {
                                                                                                    var v69 int16 = -32768
                                                                                                    v69 -= -32768
                                                                                                    _ = v69
                                                                                                }
                                                                                            }
                                                                                            if len(v6) >= 2 {
                                                                                                v6 = v6[1:len(v6)]
                                                                                            }
                                                                                            if len(v6) > 0 {
                                                                                                v6[0] = fi10
                                                                                            }
                                                                                            v70 := GenS14{F0: 18446744073709551615, F1: 1000}
                                                                                            switch ("a" + "hello") {
                                                                                            case "foo bar":
                                                                                                if v56 {
                                                                                                    v7 = (p1 << ((v67) & 63))
                                                                                                } else {
                                                                                                    fmt.Println(p2, v67, p1)
                                                                                                    if (1e10 <= (0.1 + -2.25)) {
                                                                                                        var v71 uint64 = 49
                                                                                                        switch "foo bar" {
                                                                                                        case "\n\t":
                                                                                                            p2 += 0
                                                                                                        default:
                                                                                                            v72 := GenS15{F0: 2147483647, F1: (^(p2)), F2: 1}
                                                                                                            _ = v72
                                                                                                        }
                                                                                                        if v56 {
                                                                                                            v73 := GenS2{F0: 42}
                                                                                                            v74 := GenS2{F0: 551}
                                                                                                            _ = v73
                                                                                                            _ = v74
                                                                                                        } else {
                                                                                                            var v75 uint8 = (v7 << ((v67) & 63))
                                                                                                            v6 = append(v6, int(p1))
                                                                                                            v6 = append(v6, p2)
                                                                                                            _ = v75
                                                                                                        }
                                                                                                        _ = v71
                                                                                                    } else {
                                                                                                        var v76 uint = (v67 & uint(v7))
                                                                                                        var v77 int64 = 9223372036854775807
                                                                                                        _ = v76
                                                                                                        _ = v77
                                                                                                    }
                                                                                                    switch "a" {
                                                                                                    case "foo bar":
                                                                                                        v56 = (1e10 >= ((float64(0.1) - ((float64(1e10) - 3.14159) * 1.5)) / (-2.25 / (0.1 / ((3.14159 + -2.25) + 3.14159)))))
                                                                                                    default:
                                                                                                        switch "hello" {
                                                                                                        case "a":
                                                                                                            var v78 int64 = -9223372036854775808
                                                                                                            _ = v78
                                                                                                        case "":
                                                                                                            fmt.Println(v56)
                                                                                                        default:
                                                                                                            var v79 uint8 = v7
                                                                                                            _ = v79
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            case "hello":
                                                                                                v6 = append(v6, (-(p2)))
                                                                                            case "\n\t":
                                                                                                var v80 int = (fi9 + (-(fi9)))
                                                                                                if len(v6) > 0 {
                                                                                                    v6[0] = fi9
                                                                                                }
                                                                                                var v81 int = (p2 % (((^(fi10))) | 1))
                                                                                                _ = v80
                                                                                                _ = v81
                                                                                            default:
                                                                                                var v82 float32 = -2.25
                                                                                                _ = v82
                                                                                            }
                                                                                            switch 32768 {
                                                                                            case 558:
                                                                                                p1 = v7
                                                                                            }
                                                                                            _ = v67
                                                                                            _ = v70
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v83 := GenS5{F0: 783, F1: 65535, F2: v7}
                                                                                    if v56 {
                                                                                        fmt.Println(p2)
                                                                                        _ = len(v6)
                                                                                        var v84 int64 = -859
                                                                                        _ = v84
                                                                                    }
                                                                                    p0 += p0
                                                                                    var v85 float32 = (0.1 / ((-2.25 / (((1.5 / (1e10 + 0.1)) - 1e10) - 1e10)) * ((3.14159 + float32(1e10)) * (-2.25 * 3.14159))))
                                                                                    _ = v83
                                                                                    _ = v85
                                                                                }
                                                                                if (698 >= 1) {
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                } else {
                                                                                    var v86 float64 = (0.0 / 1.5)
                                                                                    v87 := GenS16{F0: 42, F1: v56}
                                                                                    fmt.Println(fi9)
                                                                                    if (v56 || ((!(257 >= 1)) || false)) {
                                                                                        v88 := GenS17{F0: -128, F1: (^(v7)), F2: 255}
                                                                                        var v89 uint = uint(p0)
                                                                                        var v90 int = (fi9 >> (((^(v89))) & 63))
                                                                                        _ = v88
                                                                                        _ = v89
                                                                                        _ = v90
                                                                                    } else {
                                                                                        if true {
                                                                                            if (0.0 != (((float32(v86) + 0.0) / (1e10 / 1.5)) * float32(v86))) {
                                                                                                var v91 int8 = (^(v13))
                                                                                                var v92 uint16 = (p0 / ((p0) | 1))
                                                                                                fmt.Println(fi10, v7, v92)
                                                                                                _ = v91
                                                                                                _ = v92
                                                                                            }
                                                                                        } else {
                                                                                            var v93 float64 = (1e10 + 1e10)
                                                                                            v94 := GenS3{F0: p1, F1: -370}
                                                                                            v95 := GenS18{F0: 2, F1: 1, F2: 42}
                                                                                            _ = v93
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                        }
                                                                                        v13 = 43
                                                                                    }
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if len(v6) > 0 {
                                                                                v6[0] = fi10
                                                                            }
                                                                            p1 += 236
                                                                        }
                                                                    }
                                                                }
                                                                var v96 int = -988
                                                                _ = v56
                                                                _ = v57
                                                                _ = v96
                                                            }
                                                        case 840:
                                                            fmt.Println(p0)
                                                        case -567:
                                                            if len(v6) > 0 {
                                                                v6[0] = (fi9 - 202)
                                                            }
                                                        default:
                                                            v97 := GenS8{F0: int16(p0)}
                                                            _ = v97
                                                        }
                                                        _ = v55
                                                    } else {
                                                        switch ("hello" + "\n\t") {
                                                        case "a":
                                                            v98 := GenS13{F0: 570, F1: int16(v13), F2: (^(p0))}
                                                            _ = v98
                                                        default:
                                                            v7 = (v7 + (v7 - (v7 - 42)))
                                                        }
                                                        var v99 float32 = 0.1
                                                        _ = v99
                                                    }
                                                    if true {
                                                        switch 2147483647 {
                                                        case 128:
                                                            var v100 uint16 = (p0 + p0)
                                                            _ = v100
                                                        case 32768:
                                                            switch 0 {
                                                            case -9223372036854775808:
                                                                v6 = append(v6, (fi9 | (p2 % ((fi9) | 1))))
                                                            case 0:
                                                                v6 = append(v6, int(p0))
                                                            default:
                                                                var v101 uint8 = (v7 - p1)
                                                                _ = v101
                                                            }
                                                        }
                                                    } else {
                                                        var v102 string = ("" + "a")
                                                        p1 = p1
                                                        var v103 uint32 = 1
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    var v104 uint32 = 600
                                    _ = v104
                                } else {
                                    var v105 bool = false
                                    p0 -= (p0 + (p0 & 42))
                                    v7 = v7
                                    v106 := GenS19{F0: 2}
                                    var v107 uint8 = p1
                                    _ = v105
                                    _ = v106
                                    _ = v107
                                }
                            }
                            if true {
                                p1 = (p1 * uint8(p0))
                            } else {
                                _ = len(v6)
                                if (true && (2 != 2)) {
                                    var v108 uint8 = p1
                                    p0 += p0
                                    _ = v108
                                } else {
                                    var v109 uint16 = 42
                                    var v110 int8 = v13
                                    _ = v109
                                    _ = v110
                                }
                                var v111 uint = 952
                                _ = v111
                            }
                            if true {
                                var v112 int32 = int32(fi9)
                                fmt.Println(fi9, v112)
                                if len(v6) > 0 {
                                    v6[0] = 32768
                                }
                                _ = v112
                            } else {
                                var v113 uint = 989
                                switch fi10 {
                                case 2147483647:
                                    _ = len(v6)
                                case 32768:
                                    if true {
                                        switch "a" {
                                        case "foo bar":
                                            if false {
                                                _ = len(v6)
                                                if (-514 >= 0) {
                                                    var v114 string = ((((("\n\t" + (("hello" + (("" + ("" + "a")) + "")) + "a")) + "hello") + (("a" + "\n\t") + "")) + "foo bar") + "")
                                                    if len(v6) > 0 {
                                                        v6[0] = fi10
                                                    }
                                                    _ = v114
                                                }
                                            }
                                        case "hello":
                                            fmt.Println(v13, p2, v113)
                                        case "":
                                            var v115 float64 = -2.25
                                            _ = v115
                                        default:
                                            v116 := GenS4{F0: (p0 >> (((v113 | v113)) & 63)), F1: 1, F2: 0}
                                            _ = v116
                                        }
                                    } else {
                                        if true {
                                            switch 4294967295 {
                                            case 262:
                                                var v117 float32 = (3.14159 / 1.0)
                                                _ = v117
                                            case 2147483647:
                                                if (-1 != -54) {
                                                    var v118 string = "\n\t"
                                                    fmt.Println(v7, v113)
                                                    _ = v118
                                                } else {
                                                    v7 = 100
                                                    if len(v6) > 0 {
                                                        v6[0] = (-(p2))
                                                    }
                                                    var v119 int64 = 878
                                                    fmt.Println(fi9, fi10)
                                                    var v120 float64 = 0.0
                                                    _ = v119
                                                    _ = v120
                                                }
                                            default:
                                                switch (fi9 * p2) {
                                                case 128:
                                                    if true {
                                                        v121 := GenS2{F0: 2147483647}
                                                        var v122 uint16 = 338
                                                        p1 = 128
                                                        _ = v121
                                                        _ = v122
                                                    }
                                                case 0:
                                                    fmt.Println(v113)
                                                default:
                                                    var v123 float64 = ((((3.14159 - ((0.1 * 1.5) * -2.25)) - 3.14159) - 1.5) - (float64(0.1) / 1.0))
                                                    _ = v123
                                                }
                                            }
                                            v124 := GenS20{F0: (v13 >> ((65535) & 63)), F1: (fi9 >> (((v113 % ((v113) | 1))) & 63)), F2: "\n\t"}
                                            switch "foo bar" {
                                            case "a":
                                                _ = len(v6)
                                            case "\n\t":
                                                fmt.Println(fi9)
                                            default:
                                                p1 = p1
                                            }
                                            v125 := GenS19{F0: 1}
                                            _ = v124
                                            _ = v125
                                        }
                                        var v126 int = fi10
                                        _ = v126
                                    }
                                }
                                v6 = append(v6, fi9)
                                _ = v113
                            }
                            if true {
                                fmt.Println(v7)
                                v127 := GenS11{F0: 80, F1: true, F2: 61}
                                var v128 int8 = v13
                                _ = v127
                                _ = v128
                            }
                        }
                    case 127:
                        p0 = (p0 & 255)
                    default:
                        if true {
                            var v129 uint64 = 2
                            v130 := GenS4{F0: (p0 ^ 42), F1: -532, F2: 773}
                            _ = v129
                            _ = v130
                        }
                    }
                    _ = v13
                }
            }
            var v131 uint = 255
            if (9223372036854775807 <= 9223372036854775807) {
                p0 += 840
            } else {
                v131 -= v131
            }
            _ = fi10
            _ = v131
        }
        _ = fi9
    }
    switch (p2 - -706) {
    case -762:
        var v132 uint32 = 2147483647
        _ = v132
    case 2147483647:
        fmt.Println(v7, v7, p2)
    default:
        v133 := GenS2{F0: 635}
        _ = v133
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 978
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{711}
    var v3 int = 450
    v4 := []int{}
    v5 := make(map[string]string)
    if false {
        switch v3 {
        case 9223372036854775807:
            var v134 int = v3
            _ = v134
        case 0:
            v135 := GenS21{F0: -1, F1: 42, F2: 1}
            _ = v135
        }
    } else {
        if (3.14159 == (3.14159 * 1e10)) {
            switch -2147483648 {
            case 0:
                switch v3 {
                case -2147483648:
                    var v136 uint32 = 0
                    _ = v136
                case 657:
                    fmt.Println(v3, v3)
                }
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            fmt.Println(v1, v3)
            var v137 uint64 = (v1 & 18446744073709551615)
            _ = v137
        }
        fmt.Println(v3, v3)
    }
    var v138 string = ("a" + "")
    if len(v4) > 0 {
        v4[0] = v3
    }
    for fi139 := 0; fi139 < 14; fi139++ {
        fmt.Println(fi139, v1)
        for fi140 := 0; fi140 < 6; fi140++ {
            if (false || (v1 < v1)) {
                v3 = 2147483647
            }
            _ = len(v4)
            _ = fi140
        }
        delete(v5, (v138 + v138))
        _ = fi139
    }
    for fi141 := 0; fi141 < 6; fi141++ {
        var v142 uint32 = genFn0(255, 42, int(v1))
        v3 -= (v3 ^ v3)
        switch v138 {
        case "foo bar":
            var v143 int32 = -2147483648
            _ = v143
        case "\n\t":
            for fi144 := 0; fi144 < 12; fi144++ {
                v1 = (v1 * (v1 << ((4294967295) & 63)))
                var v145 int8 = 100
                var v146 int8 = (-(v145))
                _ = fi144
                _ = v145
                _ = v146
            }
        case "":
            if (uint8(v142) == 127) {
                v3 -= (-(fi141))
            } else {
                fmt.Println(v1, v3, v1)
                for fi147 := 0; fi147 < 15; fi147++ {
                    v148 := GenS22{F0: 42, F1: (true && (!false)), F2: false}
                    switch fi141 {
                    case -452:
                        var v149 uint16 = 276
                        _ = v149
                    case 9223372036854775807:
                        if (205 >= 2147483647) {
                            v3 = -1
                        }
                        var v150 uint32 = genFn0(0, 128, (fi147 - fi147))
                        _ = v150
                    default:
                        v151 := GenS23{F0: true, F1: -128, F2: ((("\n\t" + (v138 + (("" + "foo bar") + ((((v138 + v138) + v138) + "a") + (("foo bar" + ((v138 + v138) + (v138 + (v138 + (v138 + v138))))) + ("" + v138)))))) + (v138 + v138)) > ((((v138 + "hello") + ("" + "\n\t")) + "hello") + "\n\t"))}
                        _ = v151
                    }
                    v1 = (v1 >> ((624) & 63))
                    _ = fi147
                    _ = v148
                }
            }
        default:
            v152 := GenS16{F0: 1, F1: false}
            _ = v152
        }
        _ = fi141
        _ = v142
    }
    v2 = append(v2, genFn0(0, 42, 32768))
    if true {
        v153 := GenS24{F0: 0, F1: 1, F2: 42}
        var v154 string = "hello"
        _ = v153
        _ = v154
    } else {
        for fi155 := 0; fi155 < 20; fi155++ {
            var v156 uint8 = 127
            var v157 bool = true
            if v157 {
                delete(v5, ("\n\t" + ((v138 + "a") + "a")))
                v158 := GenS25{F0: v138}
                var v159 float32 = -2.25
                _ = v158
                _ = v159
            }
            _ = fi155
            _ = v156
            _ = v157
        }
        var v160 int16 = -540
        v161 := GenS21{F0: -208, F1: 0, F2: v160}
        _ = v160
        _ = v161
    }
    fmt.Println(v3, v3, v138)
    for fi162 := 0; fi162 < 15; fi162++ {
        delete(v5, (v138 + v138))
        _ = fi162
    }
    v138 = (("a" + v138) + v138)
    var v163 uint32 = 0
    var v164 int8 = 42
    fmt.Println(v163, v138)
    v165 := GenS7{F0: 0, F1: -676, F2: 42}
    _ = len(v5)
    v138 = v138
    fmt.Println("v1:", v1)
    fmt.Println("v138:", v138)
    fmt.Println("v163:", v163)
    fmt.Println("v164:", v164)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v138
    _ = v163
    _ = v164
    _ = v165
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
