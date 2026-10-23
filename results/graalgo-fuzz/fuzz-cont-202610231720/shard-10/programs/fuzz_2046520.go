package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 int; F2 int8 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int16; F1 string; F2 int8 }
type GenS12 struct { F0 int64; F1 int }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint16; F1 int64; F2 string }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint32; F1 uint8; F2 bool }
type GenS18 struct { F0 int64; F1 string; F2 uint16 }
type GenS19 struct { F0 uint32; F1 uint8 }
type GenS20 struct { F0 bool; F1 uint64 }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 uint32 }
type GenS23 struct { F0 uint16; F1 uint16; F2 uint }
func genFn0(p0 string) uint64 {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 0, F1: 42}
        v7 := []int{(fi5 * (fi5 ^ fi5))}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint32 = 2147483648
    p0 = p0
    for fi9 := 0; fi9 < 3; fi9++ {
        if false {
            v10 := []int{(^(fi9)), -1, -617}
            if (0 >= 42) {
                for fi11 := 0; fi11 < 6; fi11++ {
                    v12 := GenS2{F0: 778, F1: 42, F2: 45}
                    _ = fi11
                    _ = v12
                }
            } else {
                if (32767 < 42) {
                    for fi13 := 0; fi13 < 18; fi13++ {
                        var v14 bool = true
                        if (1 < 69) {
                            fmt.Println(fi9, fi13)
                            if v14 {
                                var v15 int8 = -128
                                if false {
                                    var v16 int16 = 1000
                                    _ = v16
                                }
                                switch fi9 {
                                case -627:
                                    v17 := GenS2{F0: -1, F1: (v8 * 2147483648), F2: 42}
                                    _ = v17
                                }
                                _ = v15
                            }
                            var v18 uint32 = (v8 & 455)
                            _ = v18
                        }
                        v8 += (v8 - (v8 / (((v8 | v8)) | 1)))
                        _ = fi13
                        _ = v14
                    }
                    if false {
                        var v19 int8 = 0
                        if (true && (false && true)) {
                            var v20 float32 = 0.1
                            switch (fi9 << ((65535) & 63)) {
                            case -1:
                                _ = len(v10)
                            default:
                                if (false && false) {
                                    if ((4294967295 > 937) || (v8 < (^(v8)))) {
                                        var v21 uint = 42
                                        v20 += v20
                                        var v22 float64 = -2.25
                                        _ = v21
                                        _ = v22
                                    } else {
                                        var v23 uint16 = 42
                                        _ = v23
                                    }
                                    v20 -= 1e10
                                }
                            }
                            switch fi9 {
                            case 4294967295:
                                var v24 uint32 = v8
                                _ = v24
                            case 624:
                                switch p0 {
                                case "hello":
                                    if (v8 == (v8 + v8)) {
                                        var v25 int32 = 0
                                        _ = v25
                                    }
                                case "foo bar":
                                    v26 := GenS3{F0: 255}
                                    _ = v26
                                }
                            }
                            _ = v20
                        } else {
                            v19 = -128
                        }
                        switch 32768 {
                        case 9223372036854775807:
                            if false {
                                v19 -= (v19 ^ (v19 + (v19 >> ((1) & 63))))
                                if ((fi9 > fi9) || ((p0 + ((p0 + ((p0 + p0) + p0)) + ("hello" + p0))) == p0)) {
                                    for fi27 := 0; fi27 < 6; fi27++ {
                                        p0 = p0
                                        _ = fi27
                                    }
                                    if true {
                                        if (42 < v19) {
                                            var v28 uint = 866
                                            fmt.Println(v8, fi9, v19)
                                            var v29 uint8 = 42
                                            _ = v28
                                            _ = v29
                                        } else {
                                            for fi30 := 0; fi30 < 5; fi30++ {
                                                if false {
                                                    var v31 uint16 = 65535
                                                    v31 -= v31
                                                    if true {
                                                        fmt.Println(v8, p0)
                                                        v31 = v31
                                                    } else {
                                                        var v32 int = 32768
                                                        if len(v10) > 0 {
                                                            v10[0] = (fi30 % (((v32 >> ((42) & 63))) | 1))
                                                        }
                                                        _ = v32
                                                    }
                                                    _ = len(v10)
                                                    fmt.Println(v19, fi9)
                                                    v10 = append(v10, fi30)
                                                    v33 := GenS4{F0: -859, F1: v8}
                                                    if (true || (((54 <= 18446744073709551615) || (18 > 128)) && (false || false))) {
                                                        if (!(-707 > (fi30 * (fi9 ^ (fi9 ^ (fi30 % ((-780) | 1))))))) {
                                                            fmt.Println(fi30, fi30, p0)
                                                            v8 -= 2
                                                        } else {
                                                            if len(v10) > 0 {
                                                                v10[0] = (-(fi9))
                                                            }
                                                            if false {
                                                                if (18446744073709551615 < 42) {
                                                                    if len(v10) > 0 {
                                                                        v10[0] = fi30
                                                                    }
                                                                }
                                                                fmt.Println(v8, p0)
                                                                if false {
                                                                    v34 := GenS4{F0: -483, F1: (v8 * v8)}
                                                                    if len(v10) >= 2 {
                                                                        v10 = v10[1:len(v10)]
                                                                    }
                                                                    _ = v34
                                                                } else {
                                                                    v35 := GenS3{F0: 2}
                                                                    _ = v35
                                                                }
                                                                var v36 int32 = 395
                                                                _ = v36
                                                            } else {
                                                                var v37 uint8 = 1
                                                                var v38 string = (((p0 + p0) + "") + p0)
                                                                var v39 int64 = 411
                                                                _ = v37
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            if (9223372036854775807 > (fi9 / ((fi9) | 1))) {
                                                                v31 = (v31 / ((795) | 1))
                                                                _ = len(v10)
                                                                var v40 int64 = 1
                                                                _ = v40
                                                            }
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                        }
                                                        var v41 int = fi30
                                                        v41 = (fi9 | (fi9 + 354))
                                                        var v42 float32 = -2.25
                                                        var v43 uint = 255
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    } else {
                                                        var v44 int32 = 1
                                                        v45 := GenS1{F0: 42, F1: 42}
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    _ = v31
                                                    _ = v33
                                                }
                                                v46 := GenS5{F0: 128, F1: "foo bar"}
                                                _ = fi30
                                                _ = v46
                                            }
                                            var v47 bool = (!true)
                                            _ = v47
                                        }
                                    }
                                    for fi48 := 0; fi48 < 1; fi48++ {
                                        if (32767 == 65535) {
                                            if (2 < 0) {
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                                if (2 <= 430) {
                                                    if (-1 < 0) {
                                                        if len(v10) > 0 {
                                                            v10[0] = fi48
                                                        }
                                                        if (2147483647 <= (fi48 ^ fi9)) {
                                                            var v49 int8 = v19
                                                            if ((-1 <= (v19 << ((258) & 63))) || (490 != 42)) {
                                                                _ = len(v10)
                                                                _ = len(v10)
                                                            } else {
                                                                p0 = (p0 + p0)
                                                                fmt.Println(p0, fi9, v49)
                                                                fmt.Println(p0)
                                                                if (42 < (v8 << ((425) & 63))) {
                                                                    if (-2147483648 > -325) {
                                                                        v19 += (v49 << ((1) & 63))
                                                                        switch p0 {
                                                                        case "a":
                                                                            var v50 bool = (p0 >= "\n\t")
                                                                            _ = v50
                                                                        case "\n\t":
                                                                            p0 = p0
                                                                        case "foo bar":
                                                                            switch p0 {
                                                                            case "\n\t":
                                                                                v51 := GenS5{F0: 42, F1: ((p0 + p0) + "foo bar")}
                                                                                _ = v51
                                                                            default:
                                                                                switch ("\n\t" + "foo bar") {
                                                                                case "foo bar":
                                                                                    fmt.Println(fi48, v49, v19)
                                                                                    if ((v49 < (v19 - (v19 % (((v19 + (v49 | (v19 % ((v19) | 1))))) | 1)))) && true) {
                                                                                        fmt.Println(fi9, v49)
                                                                                    }
                                                                                case "a":
                                                                                    var v52 uint32 = (^(v8))
                                                                                    _ = v52
                                                                                case "hello":
                                                                                    var v53 uint32 = (v8 & v8)
                                                                                    v54 := GenS6{F0: 1}
                                                                                    v55 := GenS7{F0: v8, F1: -112, F2: (v49 - int8(v53))}
                                                                                    _ = v53
                                                                                    _ = v54
                                                                                    _ = v55
                                                                                default:
                                                                                    v56 := GenS2{F0: 42, F1: v8, F2: 93}
                                                                                    _ = v56
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v57 float64 = (((float64(((-2.25 - 0.1) + float32(1.5))) / 1.0) * 0.0) - (((0.1 + (-2.25 / 1e10)) * -2.25) - (3.14159 / 1.0)))
                                                                            _ = v57
                                                                        }
                                                                    }
                                                                    v49 = 0
                                                                    switch p0 {
                                                                    case "foo bar":
                                                                        _ = len(v10)
                                                                    case "hello":
                                                                        var v58 bool = (2 > 4294967295)
                                                                        _ = v58
                                                                    }
                                                                    if len(v10) > 0 {
                                                                        v10[0] = (^(fi9))
                                                                    }
                                                                    if false {
                                                                        fmt.Println(v8)
                                                                        v59 := GenS4{F0: 513, F1: (^(v8))}
                                                                        _ = v59
                                                                    }
                                                                } else {
                                                                    if (95 > v19) {
                                                                        var v60 bool = (!true)
                                                                        _ = v60
                                                                    } else {
                                                                        if (fi48 > 42) {
                                                                            v61 := GenS6{F0: 18446744073709551615}
                                                                            switch fi48 {
                                                                            case 811:
                                                                                if ("" <= "") {
                                                                                    var v62 bool = ((fi9 << ((18446744073709551615) & 63)) >= -9223372036854775808)
                                                                                    v63 := GenS6{F0: 200}
                                                                                    fmt.Println(v49, v19)
                                                                                    fmt.Println(p0)
                                                                                    if len(v10) >= 2 {
                                                                                        v10 = v10[1:len(v10)]
                                                                                    }
                                                                                    v19 -= v19
                                                                                    _ = v62
                                                                                    _ = v63
                                                                                } else {
                                                                                    if false {
                                                                                        fmt.Println(v8, v8)
                                                                                        if false {
                                                                                            switch fi48 {
                                                                                            case 16:
                                                                                                var v64 uint32 = 2
                                                                                                _ = v64
                                                                                            default:
                                                                                                var v65 uint8 = 255
                                                                                                _ = v65
                                                                                            }
                                                                                            if false {
                                                                                                if (32767 > 2) {
                                                                                                    p0 = ("" + "a")
                                                                                                    var v66 int32 = 247
                                                                                                    v66 += (v66 + v66)
                                                                                                    _ = v66
                                                                                                }
                                                                                                if len(v10) > 0 {
                                                                                                    v10[0] = fi9
                                                                                                }
                                                                                                v19 = (v49 & v19)
                                                                                                v67 := GenS4{F0: -32768, F1: v8}
                                                                                                var v68 int64 = 42
                                                                                                v68 += 1
                                                                                                var v69 int8 = 83
                                                                                                _ = v67
                                                                                                _ = v68
                                                                                                _ = v69
                                                                                            } else {
                                                                                                var v70 uint16 = 65535
                                                                                                _ = v70
                                                                                            }
                                                                                            p0 = (p0 + (p0 + (p0 + p0)))
                                                                                        }
                                                                                        var v71 uint = 18446744073709551615
                                                                                        _ = v71
                                                                                    } else {
                                                                                        switch (fi9 / ((fi9) | 1)) {
                                                                                        case 990:
                                                                                            var v72 int32 = -2147483648
                                                                                            _ = v72
                                                                                        case -234:
                                                                                            fmt.Println(p0, p0, fi48)
                                                                                        }
                                                                                    }
                                                                                    var v73 float32 = (-2.25 + 0.1)
                                                                                    _ = v73
                                                                                }
                                                                            case 32768:
                                                                                p0 = p0
                                                                                switch p0 {
                                                                                case "foo bar":
                                                                                    v19 -= (-(v19))
                                                                                    p0 = ((p0 + p0) + p0)
                                                                                    switch p0 {
                                                                                    case "\n\t":
                                                                                        p0 = p0
                                                                                    case "a":
                                                                                        if false {
                                                                                            v74 := GenS8{F0: true, F1: -520}
                                                                                            fmt.Println(v8, fi9, v8)
                                                                                            var v75 float32 = ((3.14159 - -2.25) + 1e10)
                                                                                            _ = v74
                                                                                            _ = v75
                                                                                        } else {
                                                                                            switch (fi9 & (fi48 >> ((65535) & 63))) {
                                                                                            case 42:
                                                                                                v19 -= v19
                                                                                            case 599:
                                                                                                fmt.Println(v49, fi9)
                                                                                            default:
                                                                                                v19 += (v49 - (v19 << ((18446744073709551615) & 63)))
                                                                                            }
                                                                                            if false {
                                                                                                if len(v10) > 0 {
                                                                                                    v10[0] = fi48
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        switch 4294967295 {
                                                                                        case 740:
                                                                                            switch 0 {
                                                                                            case -73:
                                                                                                v10 = append(v10, fi48)
                                                                                                var v76 float64 = -2.25
                                                                                                v10 = append(v10, -85)
                                                                                                _ = v76
                                                                                            case -9223372036854775808:
                                                                                                var v77 int16 = -32768
                                                                                                _ = v77
                                                                                            default:
                                                                                                if (-161 >= 42) {
                                                                                                    var v78 float32 = 1.5
                                                                                                    v79 := GenS9{F0: "hello"}
                                                                                                    v80 := GenS4{F0: 32767, F1: v8}
                                                                                                    _ = v78
                                                                                                    _ = v79
                                                                                                    _ = v80
                                                                                                } else {
                                                                                                    var v81 float32 = (0.0 - ((1e10 + 3.14159) * (3.14159 - (0.1 / float32(0.1)))))
                                                                                                    fmt.Println(v81, v49)
                                                                                                    _ = v81
                                                                                                }
                                                                                                var v82 int = fi48
                                                                                                v82 = fi48
                                                                                                _ = v82
                                                                                            }
                                                                                        case -2147483648:
                                                                                            v83 := GenS10{F0: 42}
                                                                                            _ = v83
                                                                                        }
                                                                                    }
                                                                                    var v84 string = (p0 + (p0 + p0))
                                                                                    _ = v84
                                                                                case "":
                                                                                    var v85 int32 = -2147483648
                                                                                    _ = v85
                                                                                default:
                                                                                    var v86 uint64 = 0
                                                                                    var v87 int32 = 42
                                                                                    if (3.14159 >= 1e10) {
                                                                                        if ((!false) && false) {
                                                                                            var v88 uint32 = v8
                                                                                            fmt.Println(v86)
                                                                                            v19 -= (v49 * v19)
                                                                                            _ = v88
                                                                                        } else {
                                                                                            v86 = 4294967295
                                                                                            fmt.Println(v19, fi9, p0)
                                                                                        }
                                                                                        var v89 uint8 = 24
                                                                                        _ = v89
                                                                                    }
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                }
                                                                                if (true && (120 <= 2)) {
                                                                                    v19 -= v49
                                                                                    switch "a" {
                                                                                    case "":
                                                                                        fmt.Println(fi48)
                                                                                        fmt.Println(p0, v8)
                                                                                        v19 = v49
                                                                                    case "hello":
                                                                                        if len(v10) >= 2 {
                                                                                            v10 = v10[1:len(v10)]
                                                                                        }
                                                                                    case "foo bar":
                                                                                        p0 = p0
                                                                                    default:
                                                                                        switch ((p0 + (p0 + p0)) + "a") {
                                                                                        case "hello":
                                                                                            if (!(fi9 >= fi48)) {
                                                                                                fmt.Println(v49, v19)
                                                                                                v49 -= int8(v8)
                                                                                                p0 = ((p0 + "a") + "a")
                                                                                                switch 977 {
                                                                                                case -9223372036854775808:
                                                                                                    v90 := GenS4{F0: 32767, F1: 261}
                                                                                                    _ = v90
                                                                                                case 590:
                                                                                                    fmt.Println(v8, v8)
                                                                                                case 2147483647:
                                                                                                    v49 = (-(v19))
                                                                                                default:
                                                                                                    v10 = append(v10, (fi9 % (((fi48 / ((fi48) | 1))) | 1)))
                                                                                                }
                                                                                                p0 = (("foo bar" + p0) + p0)
                                                                                            } else {
                                                                                                var v91 int = fi48
                                                                                                _ = v91
                                                                                            }
                                                                                        default:
                                                                                            if ((42 < 981) && (!true)) {
                                                                                                var v92 int64 = -158
                                                                                                if false {
                                                                                                    _ = len(v10)
                                                                                                } else {
                                                                                                    if len(v10) > 0 {
                                                                                                        v10[0] = (^(fi48))
                                                                                                    }
                                                                                                    switch fi48 {
                                                                                                    case 127:
                                                                                                        v93 := GenS4{F0: 1, F1: (v8 >> ((2) & 63))}
                                                                                                        _ = v93
                                                                                                    }
                                                                                                }
                                                                                                if (false || false) {
                                                                                                    p0 = "a"
                                                                                                    _ = len(v10)
                                                                                                    v49 += (^(v19))
                                                                                                } else {
                                                                                                    v92 -= v92
                                                                                                }
                                                                                                _ = v92
                                                                                            } else {
                                                                                                var v94 uint = 2
                                                                                                switch "\n\t" {
                                                                                                case "":
                                                                                                    if true {
                                                                                                        v95 := GenS2{F0: 42, F1: v8, F2: 223}
                                                                                                        var v96 uint64 = 1
                                                                                                        if (-2.25 != float64(1.5)) {
                                                                                                            var v97 float32 = float32(1.5)
                                                                                                            fmt.Println(p0, v96, v19)
                                                                                                            var v98 uint16 = 0
                                                                                                            _ = v97
                                                                                                            _ = v98
                                                                                                        }
                                                                                                        _ = v95
                                                                                                        _ = v96
                                                                                                    } else {
                                                                                                        fmt.Println(v49)
                                                                                                        p0 = ""
                                                                                                    }
                                                                                                default:
                                                                                                    if (9223372036854775807 != -9223372036854775808) {
                                                                                                        if (v94 >= 255) {
                                                                                                            fmt.Println(p0, v49, fi9)
                                                                                                            var v99 uint16 = 2
                                                                                                            _ = v99
                                                                                                        } else {
                                                                                                            p0 = p0
                                                                                                            if len(v10) >= 2 {
                                                                                                                v10 = v10[1:len(v10)]
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                p0 = ""
                                                                                                _ = v94
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            if len(v10) >= 2 {
                                                                                v10 = v10[1:len(v10)]
                                                                            }
                                                                            v100 := GenS2{F0: -1, F1: (v8 % ((154) | 1)), F2: 127}
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                var v101 int32 = 0
                                                                                fmt.Println(v101, p0, fi9)
                                                                                _ = v101
                                                                            default:
                                                                                var v102 float32 = float32(0.0)
                                                                                _ = v102
                                                                            }
                                                                            _ = v100
                                                                            _ = v61
                                                                        } else {
                                                                            v103 := GenS11{F0: 1000, F1: (p0 + "\n\t"), F2: v19}
                                                                            switch 32767 {
                                                                            case 1:
                                                                                if ((v8 + v8) == 654) {
                                                                                    var v104 int32 = 1
                                                                                    var v105 int = (fi9 & (fi48 * fi9))
                                                                                    _ = v104
                                                                                    _ = v105
                                                                                }
                                                                            }
                                                                            if true {
                                                                                var v106 int64 = 9223372036854775807
                                                                                _ = v106
                                                                            } else {
                                                                                var v107 uint32 = v8
                                                                                v108 := GenS12{F0: 42, F1: int(v8)}
                                                                                _ = v107
                                                                                _ = v108
                                                                            }
                                                                            _ = v103
                                                                        }
                                                                        fmt.Println(fi48, v19, v49)
                                                                        if (!false) {
                                                                            var v109 uint64 = 42
                                                                            switch (fi48 ^ (fi9 + fi48)) {
                                                                            case 42:
                                                                                if (127 >= 0) {
                                                                                    v49 -= int8(fi9)
                                                                                    fmt.Println(fi48)
                                                                                    var v110 uint32 = (v8 & v8)
                                                                                    _ = v110
                                                                                } else {
                                                                                    switch (p0 + p0) {
                                                                                    case "foo bar":
                                                                                        if (!false) {
                                                                                            _ = len(v10)
                                                                                            _ = len(v10)
                                                                                        } else {
                                                                                            if (0 != 9223372036854775807) {
                                                                                                var v111 float64 = (3.14159 * 3.14159)
                                                                                                v112 := GenS6{F0: v109}
                                                                                                _ = v111
                                                                                                _ = v112
                                                                                            } else {
                                                                                                switch (fi9 << ((4294967295) & 63)) {
                                                                                                case 1:
                                                                                                    v113 := GenS5{F0: 1, F1: p0}
                                                                                                    _ = v113
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case "a":
                                                                                        if len(v10) >= 2 {
                                                                                            v10 = v10[1:len(v10)]
                                                                                        }
                                                                                    default:
                                                                                        v109 += v109
                                                                                    }
                                                                                    v19 = (v49 % (((^(v19))) | 1))
                                                                                }
                                                                            case 1:
                                                                                v114 := GenS4{F0: -348, F1: (^(v8))}
                                                                                if len(v10) > 0 {
                                                                                    v10[0] = int(v109)
                                                                                }
                                                                                p0 = "foo bar"
                                                                                if (255 >= 32767) {
                                                                                    var v115 uint16 = 65535
                                                                                    if false {
                                                                                        switch -2147483648 {
                                                                                        case 32767:
                                                                                            var v116 int16 = -763
                                                                                            if true {
                                                                                                switch p0 {
                                                                                                case "a":
                                                                                                    switch ("" + "a") {
                                                                                                    case "\n\t":
                                                                                                        switch 1 {
                                                                                                        case -9223372036854775808:
                                                                                                            switch p0 {
                                                                                                            case "hello":
                                                                                                                if (9223372036854775807 != -869) {
                                                                                                                    var v117 bool = false
                                                                                                                    var v118 uint8 = 128
                                                                                                                    var v119 float32 = (1e10 * (1.5 / (0.1 + -2.25)))
                                                                                                                    _ = v117
                                                                                                                    _ = v118
                                                                                                                    _ = v119
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                var v120 float32 = (-2.25 + 0.0)
                                                                                                                _ = v120
                                                                                                            case "":
                                                                                                                var v121 int32 = -1
                                                                                                                _ = v121
                                                                                                            }
                                                                                                            _ = len(v10)
                                                                                                        case -1:
                                                                                                            v122 := GenS8{F0: (!true), F1: (fi9 & (fi9 & fi9))}
                                                                                                            _ = v122
                                                                                                        case -261:
                                                                                                            if len(v10) > 0 {
                                                                                                                v10[0] = fi48
                                                                                                            }
                                                                                                        default:
                                                                                                            switch -868 {
                                                                                                            case -2147483648:
                                                                                                                fmt.Println(v19)
                                                                                                                if len(v10) > 0 {
                                                                                                                    v10[0] = -626
                                                                                                                }
                                                                                                            case 2147483647:
                                                                                                                if len(v10) >= 2 {
                                                                                                                    v10 = v10[1:len(v10)]
                                                                                                                }
                                                                                                                v123 := GenS1{F0: v115, F1: (^(v116))}
                                                                                                                var v124 int64 = -9223372036854775808
                                                                                                                _ = v123
                                                                                                                _ = v124
                                                                                                            case 127:
                                                                                                                fmt.Println(v116, v115)
                                                                                                            default:
                                                                                                                var v125 int64 = 193
                                                                                                                _ = v125
                                                                                                            }
                                                                                                        }
                                                                                                        if false {
                                                                                                            var v126 bool = ((1e10 * (-2.25 / 1.5)) != ((1.5 / 1.5) / 1.5))
                                                                                                            if v126 {
                                                                                                                var v127 int16 = -32768
                                                                                                                fmt.Println(fi9, v116, fi9)
                                                                                                                var v128 int8 = 69
                                                                                                                _ = v127
                                                                                                                _ = v128
                                                                                                            } else {
                                                                                                                v129 := GenS13{F0: 42}
                                                                                                                switch 42 {
                                                                                                                case 128:
                                                                                                                    v10 = append(v10, fi48)
                                                                                                                case 1:
                                                                                                                    v10 = append(v10, fi48)
                                                                                                                default:
                                                                                                                    v115 -= (v115 ^ (v115 + v115))
                                                                                                                }
                                                                                                                _ = v129
                                                                                                            }
                                                                                                            if true {
                                                                                                                p0 = ((("a" + p0) + p0) + "foo bar")
                                                                                                                v109 = v109
                                                                                                                v10 = append(v10, (fi48 ^ fi9))
                                                                                                                var v130 uint = 0
                                                                                                                _ = v130
                                                                                                            } else {
                                                                                                                if len(v10) >= 2 {
                                                                                                                    v10 = v10[1:len(v10)]
                                                                                                                }
                                                                                                                if (uint8(v8) < 128) {
                                                                                                                    var v131 float32 = (-2.25 - 0.0)
                                                                                                                    var v132 uint32 = v8
                                                                                                                    v131 += float32(-2.25)
                                                                                                                    _ = v131
                                                                                                                    _ = v132
                                                                                                                }
                                                                                                                v109 += (v109 << ((42) & 63))
                                                                                                                switch ("foo bar" + "") {
                                                                                                                case "\n\t":
                                                                                                                    if len(v10) >= 2 {
                                                                                                                        v10 = v10[1:len(v10)]
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    var v133 uint8 = 127
                                                                                                                    _ = v133
                                                                                                                case "":
                                                                                                                    if len(v10) > 0 {
                                                                                                                        v10[0] = 26
                                                                                                                    }
                                                                                                                }
                                                                                                                var v134 uint8 = 0
                                                                                                                v109 += 42
                                                                                                                _ = v134
                                                                                                            }
                                                                                                            _ = v126
                                                                                                        }
                                                                                                    }
                                                                                                    switch ((p0 + p0) + p0) {
                                                                                                    case "hello":
                                                                                                        v116 += v116
                                                                                                    case "foo bar":
                                                                                                        v135 := GenS14{F0: 970}
                                                                                                        _ = v135
                                                                                                    case "\n\t":
                                                                                                        if len(v10) >= 2 {
                                                                                                            v10 = v10[1:len(v10)]
                                                                                                        }
                                                                                                    default:
                                                                                                        v115 = 865
                                                                                                    }
                                                                                                case "hello":
                                                                                                    if false {
                                                                                                        var v136 int32 = 1
                                                                                                        v137 := GenS15{F0: v115, F1: 0, F2: p0}
                                                                                                        var v138 float32 = -2.25
                                                                                                        _ = v136
                                                                                                        _ = v137
                                                                                                        _ = v138
                                                                                                    }
                                                                                                    var v139 int = fi48
                                                                                                    if (!(0 == -2147483648)) {
                                                                                                        v140 := GenS16{F0: 703}
                                                                                                        var v141 uint64 = v109
                                                                                                        _ = v140
                                                                                                        _ = v141
                                                                                                    }
                                                                                                    _ = v139
                                                                                                }
                                                                                                v142 := GenS17{F0: v8, F1: 127, F2: ((!false) || (!(-518 <= -2147483648)))}
                                                                                                v49 += v49
                                                                                                _ = v142
                                                                                            }
                                                                                            _ = v116
                                                                                        case 2147483647:
                                                                                            fmt.Println(p0, fi48)
                                                                                        }
                                                                                    } else {
                                                                                        v143 := GenS12{F0: 0, F1: int(v49)}
                                                                                        _ = v143
                                                                                    }
                                                                                    _ = v115
                                                                                } else {
                                                                                    if false {
                                                                                        if ((float32(float64(0.0)) != 0.0) || (503 > 4294967295)) {
                                                                                            v10 = append(v10, fi48)
                                                                                            if len(v10) > 0 {
                                                                                                v10[0] = fi48
                                                                                            }
                                                                                            if (1 >= 0) {
                                                                                                v144 := GenS1{F0: 318, F1: -708}
                                                                                                v109 -= (v109 + v109)
                                                                                                v49 = v19
                                                                                                var v145 uint32 = v8
                                                                                                v8 += (^(v8))
                                                                                                _ = v144
                                                                                                _ = v145
                                                                                            }
                                                                                        } else {
                                                                                            v49 += v49
                                                                                        }
                                                                                        v146 := GenS9{F0: "\n\t"}
                                                                                        _ = v146
                                                                                    }
                                                                                }
                                                                                switch p0 {
                                                                                case "":
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        var v147 int8 = int8(v8)
                                                                                        _ = v147
                                                                                    case "foo bar":
                                                                                        if len(v10) >= 2 {
                                                                                            v10 = v10[1:len(v10)]
                                                                                        }
                                                                                    case "hello":
                                                                                        var v148 uint8 = 255
                                                                                        _ = v148
                                                                                    }
                                                                                case "\n\t":
                                                                                    v149 := GenS1{F0: 489, F1: 467}
                                                                                    _ = v149
                                                                                case "a":
                                                                                    if false {
                                                                                        v8 += 1
                                                                                    } else {
                                                                                        p0 = "foo bar"
                                                                                        v8 -= v8
                                                                                        switch p0 {
                                                                                        case "a":
                                                                                            switch (p0 + p0) {
                                                                                            case "a":
                                                                                                if true {
                                                                                                    if false {
                                                                                                        v8 += (v8 & v8)
                                                                                                        var v150 string = p0
                                                                                                        _ = v150
                                                                                                    }
                                                                                                    if (!false) {
                                                                                                        v151 := GenS18{F0: 0, F1: ((("hello" + p0) + "") + ("\n\t" + p0)), F2: 82}
                                                                                                        switch p0 {
                                                                                                        case "a":
                                                                                                            v10 = append(v10, fi48)
                                                                                                        case "hello":
                                                                                                            switch (fi48 | int(v8)) {
                                                                                                            case -1:
                                                                                                                v152 := GenS18{F0: 42, F1: p0, F2: 32767}
                                                                                                                _ = v152
                                                                                                            case -433:
                                                                                                                switch fi9 {
                                                                                                                case -9223372036854775808:
                                                                                                                    p0 = p0
                                                                                                                case -2147483648:
                                                                                                                    switch 944 {
                                                                                                                    case 2147483647:
                                                                                                                        switch (fi9 << ((4294967295) & 63)) {
                                                                                                                        case -9223372036854775808:
                                                                                                                            if false {
                                                                                                                                if len(v10) >= 2 {
                                                                                                                                    v10 = v10[1:len(v10)]
                                                                                                                                }
                                                                                                                                var v153 int32 = 0
                                                                                                                                var v154 float32 = 0.0
                                                                                                                                _ = v153
                                                                                                                                _ = v154
                                                                                                                            }
                                                                                                                        case 1:
                                                                                                                            var v155 float32 = (1.5 * -2.25)
                                                                                                                            _ = v155
                                                                                                                        default:
                                                                                                                            var v156 uint16 = 1
                                                                                                                            _ = v156
                                                                                                                        }
                                                                                                                    case 1:
                                                                                                                        v157 := GenS19{F0: v8, F1: 67}
                                                                                                                        _ = v157
                                                                                                                    case 128:
                                                                                                                        v8 = (v8 + 210)
                                                                                                                    }
                                                                                                                }
                                                                                                            case 32768:
                                                                                                                var v158 uint8 = 128
                                                                                                                _ = v158
                                                                                                            default:
                                                                                                                v8 += v8
                                                                                                            }
                                                                                                        default:
                                                                                                            v19 += v49
                                                                                                        }
                                                                                                        fmt.Println(v19, v8, fi48)
                                                                                                        var v159 uint8 = 0
                                                                                                        fmt.Println(v19, fi48)
                                                                                                        switch (-(fi48)) {
                                                                                                        case -270:
                                                                                                            if len(v10) >= 2 {
                                                                                                                v10 = v10[1:len(v10)]
                                                                                                            }
                                                                                                        case -409:
                                                                                                            if len(v10) > 0 {
                                                                                                                v10[0] = fi9
                                                                                                            }
                                                                                                        case -1:
                                                                                                            v160 := GenS20{F0: true, F1: uint64(v159)}
                                                                                                            _ = v160
                                                                                                        default:
                                                                                                            if len(v10) > 0 {
                                                                                                                v10[0] = -9223372036854775808
                                                                                                            }
                                                                                                        }
                                                                                                        if len(v10) >= 2 {
                                                                                                            v10 = v10[1:len(v10)]
                                                                                                        }
                                                                                                        _ = v151
                                                                                                        _ = v159
                                                                                                    } else {
                                                                                                        switch (fi9 * fi9) {
                                                                                                        case 157:
                                                                                                            v49 += (v49 << ((4294967295) & 63))
                                                                                                        default:
                                                                                                            v10 = append(v10, (fi9 & -1))
                                                                                                        }
                                                                                                        var v161 float32 = (1.5 + 1e10)
                                                                                                        var v162 int64 = -9223372036854775808
                                                                                                        _ = v161
                                                                                                        _ = v162
                                                                                                    }
                                                                                                    var v163 string = p0
                                                                                                    _ = v163
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v164 := GenS21{F0: 0}
                                                                                                _ = v164
                                                                                            case "":
                                                                                                v165 := GenS18{F0: 1, F1: (p0 + p0), F2: 659}
                                                                                                _ = v165
                                                                                            default:
                                                                                                v8 += v8
                                                                                            }
                                                                                        case "":
                                                                                            fmt.Println(v8)
                                                                                            var v166 float32 = float32(0.0)
                                                                                            _ = v166
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v167 float64 = (-2.25 - 0.0)
                                                                                    _ = v167
                                                                                }
                                                                                _ = v114
                                                                            default:
                                                                                v168 := GenS22{F0: (^(v8))}
                                                                                _ = v168
                                                                            }
                                                                            _ = v109
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            if (true || (-123 <= 1)) {
                                                                fmt.Println(v19, v19, fi48)
                                                                if false {
                                                                    fmt.Println(fi48, fi48, p0)
                                                                } else {
                                                                    v49 = 100
                                                                    fmt.Println(p0, v19)
                                                                    fmt.Println(v49)
                                                                }
                                                            }
                                                            _ = v49
                                                        }
                                                        var v169 int64 = -9223372036854775808
                                                        _ = v169
                                                    } else {
                                                        v170 := GenS11{F0: 0, F1: p0, F2: (v19 >> ((1) & 63))}
                                                        var v171 float32 = float32(((1e10 / 3.14159) - (float64((float32(1e10) * (3.14159 * (((1e10 - 3.14159) + 3.14159) - 1.5)))) + 0.1)))
                                                        if len(v10) >= 2 {
                                                            v10 = v10[1:len(v10)]
                                                        }
                                                        _ = v170
                                                        _ = v171
                                                    }
                                                    fmt.Println(fi9)
                                                    p0 = (p0 + p0)
                                                }
                                                fmt.Println(v19)
                                            } else {
                                                fmt.Println(p0)
                                                var v172 float32 = (-2.25 + float32(0.0))
                                                _ = v172
                                            }
                                            v19 -= (v19 | v19)
                                            if ((3.14159 == 1e10) && ((!(0 <= 1000)) && (v19 == (v19 * 119)))) {
                                                v8 -= v8
                                            }
                                            v8 = (v8 % (((^(v8))) | 1))
                                        }
                                        fmt.Println(fi9, fi48, v19)
                                        _ = fi48
                                    }
                                }
                            } else {
                                fmt.Println(fi9, v8, fi9)
                                var v173 float32 = (3.14159 * ((float32(0.0) / 1.5) + 1.5))
                                var v174 int8 = (-(v19))
                                _ = v173
                                _ = v174
                            }
                        }
                        _ = v19
                    }
                }
            }
            _ = v10
        }
        var v175 string = p0
        _ = fi9
        _ = v175
    }
    switch -1 {
    case 9223372036854775807:
        p0 = (p0 + (p0 + "hello"))
    case -1:
        v176 := GenS17{F0: 2147483648, F1: 128, F2: (42 >= 0)}
        _ = v176
    case 4294967295:
        v177 := GenS23{F0: 42, F1: 42, F2: 0}
        _ = v177
    }
    _ = p0
    _ = v8
    return 2
}
func genFn1(p0 uint64, p1 int32) string {
    var v178 string = "hello"
    v179 := []int{-885, -9223372036854775808, -1, 6}
    p1 = p1
    _ = p0
    _ = p1
    _ = v178
    _ = v179
    return "foo bar"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 100
    var v3 uint64 = 711
    var v4 int64 = -9223372036854775808
    if (!true) {
        if true {
            var v180 uint64 = genFn0("hello")
            var v181 uint8 = 127
            for fi182 := 0; fi182 < 14; fi182++ {
                v180 += v180
                _ = fi182
            }
            _ = v180
            _ = v181
        } else {
            if true {
                v1["a"] = 2147483648
            }
        }
        switch "a" {
        case "hello":
            switch "foo bar" {
            case "\n\t":
                var v183 bool = (!false)
                _ = v183
            case "":
                switch "foo bar" {
                case "":
                    switch -2147483648 {
                    case -640:
                        v2 += 0
                    default:
                        v184 := []int{128, -2147483648, -954}
                        _ = v184
                    }
                case "\n\t":
                    fmt.Println(v2, v2, v4)
                case "a":
                    var v185 bool = ((0.1 - 3.14159) != float64(1e10))
                    _ = v185
                default:
                    switch 1 {
                    case 42:
                        v4 += v4
                    case -9223372036854775808:
                        if (0 < -844) {
                            fmt.Println(v4, v3, v4)
                            v186 := GenS19{F0: 2147483648, F1: 2}
                            v3 += (v3 ^ (v3 * (v3 | (^(v3)))))
                            _ = v186
                        }
                    case 1:
                        var v187 bool = true
                        _ = v187
                    }
                }
            }
        case "foo bar":
            var v188 int8 = (^(v2))
            _ = v188
        case "a":
            var v189 int8 = 103
            _ = v189
        default:
            v190 := []int{}
            _ = v190
        }
    } else {
        v4 += (-(v4))
    }
    for fi191 := 0; fi191 < 17; fi191++ {
        fmt.Println(fi191, v4, v3)
        for fi192 := 0; fi192 < 16; fi192++ {
            if (("hello" + "a") <= "hello") {
                var v193 int16 = 0
                v3 += genFn0("\n\t")
                v2 = (v2 >> ((630) & 63))
                _ = v193
            }
            var v194 int32 = 0
            _ = fi192
            _ = v194
        }
        if false {
            for fi195 := 0; fi195 < 2; fi195++ {
                var v196 int16 = 32767
                if true {
                    var v197 uint64 = (v3 - (v3 / ((v3) | 1)))
                    v198 := GenS1{F0: 1, F1: (v196 % ((v196) | 1))}
                    if false {
                        if ((32767 <= (fi195 & (fi191 << ((18446744073709551615) & 63)))) || false) {
                            if false {
                                v4 -= (v4 * (v4 ^ (v4 / (((v4 ^ (v4 & v4))) | 1))))
                            }
                            fmt.Println(v196)
                        } else {
                            var v199 string = (("hello" + ("hello" + "\n\t")) + (("hello" + "") + ("a" + "\n\t")))
                            _ = v199
                        }
                        v200 := GenS16{F0: 0}
                        var v201 int16 = 32767
                        v202 := GenS1{F0: 255, F1: v196}
                        _ = v200
                        _ = v201
                        _ = v202
                    }
                    _ = v197
                    _ = v198
                }
                _ = fi195
                _ = v196
            }
            delete(v1, ("a" + ((("" + ("a" + ("foo bar" + "hello"))) + "hello") + ("" + ""))))
        }
        _ = fi191
    }
    fmt.Println(v3, v3, v2)
    var v203 uint32 = 526
    if false {
        var v204 int16 = -1
        _ = v204
    } else {
        v2 -= (-(v2))
        for fi205 := 0; fi205 < 13; fi205++ {
            var v206 int8 = int8(v203)
            if false {
                for fi207 := 0; fi207 < 8; fi207++ {
                    var v208 string = "a"
                    fmt.Println(fi205, v208, v203)
                    _ = fi207
                    _ = v208
                }
            }
            v209 := GenS16{F0: 971}
            _ = fi205
            _ = v206
            _ = v209
        }
    }
    var v210 uint16 = 65535
    switch 42 {
    case 2147483647:
        var v211 int16 = 830
        _ = v211
    default:
        for fi212 := 0; fi212 < 11; fi212++ {
            v3 = (v3 ^ (v3 >> ((942) & 63)))
            v1["hello"] = uint32(v210)
            _ = fi212
        }
    }
    for fi213 := 0; fi213 < 20; fi213++ {
        var v214 float64 = ((-2.25 - 1.5) * 1.5)
        var v215 uint8 = uint8(v210)
        v2 -= 0
        _ = fi213
        _ = v214
        _ = v215
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v203:", v203)
    fmt.Println("v210:", v210)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v203
    _ = v210
    _ = v3
    _ = v4
}
