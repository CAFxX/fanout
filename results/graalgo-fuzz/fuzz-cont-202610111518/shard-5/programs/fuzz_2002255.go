package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int; F1 int; F2 int16 }
type GenS9 struct { F0 bool; F1 uint8 }
type GenS10 struct { F0 uint8; F1 int8 }
type GenS11 struct { F0 uint32; F1 bool }
type GenS12 struct { F0 uint; F1 int64 }
type GenS13 struct { F0 int32; F1 int64 }
type GenS14 struct { F0 int; F1 bool; F2 uint32 }
type GenS15 struct { F0 uint; F1 uint; F2 int }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 int; F1 uint64; F2 bool }
type GenS18 struct { F0 int8; F1 bool }
type GenS19 struct { F0 int32; F1 int64; F2 uint64 }
type GenS20 struct { F0 int64; F1 uint16 }
type GenS21 struct { F0 uint8 }
type GenS22 struct { F0 uint; F1 bool; F2 uint64 }
type GenS23 struct { F0 uint; F1 int; F2 uint64 }
type GenS24 struct { F0 bool; F1 string }
type GenS25 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS26 struct { F0 int32 }
type GenS27 struct { F0 uint16 }
func genFn0(p0 bool, p1 uint64) uint16 {
    v6 := GenS1{F0: (p1 - 1), F1: 1}
    if (0.1 >= 1e10) {
        var v7 uint16 = 0
        var v8 int32 = 745
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}
func genFn1(p0 int32, p1 uint16) bool {
    v9 := GenS2{F0: 0, F1: 255}
    var v10 int8 = -7
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return false
}

func main() {
    var v1 string = "foo bar"
    var v2 float64 = 1.5
    v3 := []uint32{487, 2147483647, 2}
    v4 := []int{-2147483648, -560}
    var v5 uint16 = 1
    switch ("foo bar" + v1) {
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "foo bar":
        for fi11 := 0; fi11 < 14; fi11++ {
            v1 = v1
            _ = fi11
        }
    case "hello":
        for fi12 := 0; fi12 < 3; fi12++ {
            switch v1 {
            case "foo bar":
                _ = len(v4)
            }
            var v13 int64 = 663
            _ = fi12
            _ = v13
        }
    default:
        if len(v3) > 0 {
            v3[0] = 2
        }
    }
    if true {
        v1 = v1
    }
    var v14 float32 = 1.5
    _ = len(v3)
    v2 += (3.14159 / v2)
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        v1 = "hello"
    }
    if false {
        fmt.Println(v5)
    }
    var v15 uint16 = (^(v5))
    fmt.Println(v15)
    for fi16 := 0; fi16 < 6; fi16++ {
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 float32 = ((((v14 + v14) / 1e10) - (((0.1 - v14) * 0.1) / ((0.1 / -2.25) / float32(v2)))) * float32(v2))
            var v19 uint64 = 2
            _ = fi17
            _ = v18
            _ = v19
        }
        if true {
            fmt.Println(fi16)
            v20 := GenS1{F0: 974, F1: 1000}
            _ = v20
        } else {
            v2 += 1.5
            v15 -= genFn0(true, 4294967295)
        }
        _ = fi16
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        if (!(true || true)) {
            switch 32767 {
            case 9223372036854775807:
                var v22 float32 = v14
                _ = v22
            case 32768:
                switch (((v1 + v1) + "") + v1) {
                case "hello":
                    var v23 uint64 = uint64(v5)
                    _ = v23
                case "foo bar":
                    v24 := GenS3{F0: 69}
                    _ = v24
                default:
                    if true {
                        v14 -= v14
                        var v25 uint32 = 4294967295
                        _ = v25
                    } else {
                        v26 := GenS3{F0: 865}
                        var v27 int32 = -581
                        v28 := GenS4{F0: v15, F1: -32768}
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                }
            default:
                v4 = append(v4, 128)
            }
            v29 := GenS5{F0: 56}
            _ = v29
        }
        if (((true && false) && false) && false) {
            switch v1 {
            case "":
                v2 = v2
            case "a":
                v2 += ((((0.1 + v2) * float64((v14 + v14))) - v2) / (float64(((((0.0 * float32(v2)) / v14) + v14) / -2.25)) / v2))
            case "foo bar":
                if (2147483647 >= 65535) {
                    v14 += 0.1
                }
            }
            if (-2147483648 == -2147483648) {
                fmt.Println(v5, v5)
                if (42 >= -844) {
                    v15 -= (^(v5))
                    for fi30 := 0; fi30 < 7; fi30++ {
                        v5 -= (v5 / ((v5) | 1))
                        if true {
                            v5 += (v5 - 42)
                            var v31 int8 = 100
                            _ = v31
                        }
                        _ = fi30
                    }
                    switch fi21 {
                    case 4294967295:
                        v32 := GenS6{F0: (v1 >= "foo bar"), F1: 748, F2: 0}
                        _ = v32
                    case -9223372036854775808:
                        var v33 uint64 = 18446744073709551615
                        _ = v33
                    }
                } else {
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v35 := GenS7{F0: 65535}
                        v36 := GenS8{F0: fi34, F1: (fi34 - fi21), F2: 1}
                        if true {
                            v37 := GenS9{F0: (!("\n\t" != ("" + v1))), F1: 0}
                            if true {
                                var v38 uint64 = 42
                                var v39 uint = 42
                                v40 := GenS4{F0: (v5 << (((v39 >> ((65535) & 63))) & 63)), F1: 42}
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            if (!true) {
                                if len(v4) > 0 {
                                    v4[0] = fi21
                                }
                                v1 = v1
                            } else {
                                var v41 uint64 = 0
                                _ = v41
                            }
                            _ = v37
                        } else {
                            v15 += 92
                        }
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    v1 = ((v1 + "a") + ("foo bar" + "foo bar"))
                    var v42 uint64 = 18446744073709551615
                    _ = v42
                }
            }
        } else {
            fmt.Println(v2, v15)
            switch 1 {
            case -345:
                var v43 int16 = 1
                _ = v43
            }
            if (float64(v14) != 1e10) {
                if true {
                    if false {
                        _ = len(v3)
                        if false {
                            if (-619 == 0) {
                                for fi44 := 0; fi44 < 14; fi44++ {
                                    if false {
                                        switch fi44 {
                                        case 32768:
                                            v15 = (^(v5))
                                        case 4294967295:
                                            if (false || (-93 < -32768)) {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v45 := GenS10{F0: 32, F1: -1}
                                                if (!(false && true)) {
                                                    if (true || (0 < 2)) {
                                                        v46 := GenS10{F0: uint8(v15), F1: int8(v15)}
                                                        var v47 float64 = v2
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    fmt.Println(v1)
                                                } else {
                                                    v2 += v2
                                                    v3 = append(v3, 1)
                                                    var v48 int16 = 349
                                                    _ = v48
                                                }
                                                _ = v45
                                            } else {
                                                switch (^(fi21)) {
                                                case 32767:
                                                    fmt.Println(fi21)
                                                case 128:
                                                    if true {
                                                        var v49 int16 = -1
                                                        _ = v49
                                                    } else {
                                                        if (0 > 18446744073709551615) {
                                                            var v50 int64 = 1
                                                            if (42 <= (^(v50))) {
                                                                if ((!(!((0.0 / v14) < 1.5))) || false) {
                                                                    if false {
                                                                        var v51 uint8 = 1
                                                                        if (!true) {
                                                                            v3 = append(v3, 2)
                                                                        }
                                                                        _ = v51
                                                                    } else {
                                                                        if true {
                                                                            v50 += -9223372036854775808
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                        }
                                                                        v15 += v15
                                                                    }
                                                                } else {
                                                                    v52 := GenS11{F0: uint32(v15), F1: (!false)}
                                                                    fmt.Println(v50)
                                                                    _ = v52
                                                                }
                                                                v53 := GenS10{F0: 127, F1: 22}
                                                                v2 -= v2
                                                                _ = v53
                                                            }
                                                            var v54 uint8 = 32
                                                            _ = v50
                                                            _ = v54
                                                        } else {
                                                            v1 = ((v1 + v1) + v1)
                                                            v2 -= 1.5
                                                        }
                                                        if (!false) {
                                                            switch (^(fi21)) {
                                                            case 32768:
                                                                v14 = v14
                                                            default:
                                                                fmt.Println(v1)
                                                            }
                                                            v55 := GenS12{F0: 42, F1: -1}
                                                            v56 := GenS6{F0: true, F1: 723, F2: 128}
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    }
                                                case 0:
                                                    v4 = append(v4, int(v15))
                                                }
                                            }
                                        default:
                                            v5 = v5
                                        }
                                        v4 = append(v4, (fi21 * fi21))
                                        v15 -= v15
                                    } else {
                                        v57 := GenS4{F0: (v5 | (v15 >> ((2) & 63))), F1: -406}
                                        _ = v57
                                    }
                                    if false {
                                        _ = len(v3)
                                    } else {
                                        if (!((-2 != fi44) || ((848 == 1) || false))) {
                                            var v58 int32 = 0
                                            var v59 string = v1
                                            fmt.Println(v15, v15, v58)
                                            _ = v58
                                            _ = v59
                                        } else {
                                            v5 = (v5 ^ v15)
                                        }
                                        switch fi44 {
                                        case 42:
                                            v1 = ""
                                        default:
                                            var v60 int32 = 2147483647
                                            _ = v60
                                        }
                                        v61 := GenS10{F0: 255, F1: 0}
                                        _ = v61
                                    }
                                    v62 := GenS13{F0: 0, F1: 283}
                                    _ = fi44
                                    _ = v62
                                }
                                for fi63 := 0; fi63 < 12; fi63++ {
                                    if true {
                                        v64 := GenS8{F0: 127, F1: (fi21 | 4294967295), F2: -771}
                                        _ = v64
                                    } else {
                                        switch "\n\t" {
                                        case "foo bar":
                                            if false {
                                                var v65 int64 = 1
                                                v1 = (v1 + "foo bar")
                                                var v66 float32 = (-2.25 / float32((v2 + ((1.5 / ((3.14159 - float64(v14)) - v2)) / (v2 / v2)))))
                                                _ = v65
                                                _ = v66
                                            } else {
                                                var v67 string = v1
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v67
                                            }
                                        case "\n\t":
                                            if true {
                                                switch (("a" + "") + v1) {
                                                case "":
                                                    v1 = (("foo bar" + v1) + v1)
                                                    v3 = append(v3, 2)
                                                case "hello":
                                                    switch v1 {
                                                    case "":
                                                        fmt.Println(fi21, v1)
                                                        _ = len(v3)
                                                    case "foo bar":
                                                        fmt.Println(v2, fi63)
                                                    case "\n\t":
                                                        if (!true) {
                                                            v3 = append(v3, 331)
                                                        } else {
                                                            v1 = ""
                                                            fmt.Println(v5)
                                                            switch (v1 + "") {
                                                            case "hello":
                                                                _ = len(v4)
                                                            case "":
                                                                if (false && true) {
                                                                    v2 -= (float64(((3.14159 * v14) * v14)) - v2)
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    var v68 int32 = int32(fi21)
                                                                    _ = v68
                                                                }
                                                            case "foo bar":
                                                                var v69 int32 = 2147483647
                                                                _ = v69
                                                            default:
                                                                var v70 uint8 = 128
                                                                _ = v70
                                                            }
                                                            v71 := GenS1{F0: 0, F1: 42}
                                                            _ = v71
                                                        }
                                                    default:
                                                        if (18446744073709551615 > 255) {
                                                            fmt.Println(v1)
                                                            v72 := GenS14{F0: fi63, F1: ((1 != 867) && true), F2: 1}
                                                            v5 += (v15 + v5)
                                                            if true {
                                                                var v73 uint16 = genFn0(false, 2)
                                                                _ = v73
                                                            } else {
                                                                var v74 int64 = 42
                                                                v14 -= -2.25
                                                                var v75 float64 = v2
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                            var v76 float32 = (((v14 + v14) + v14) / v14)
                                                            _ = v72
                                                            _ = v76
                                                        } else {
                                                            if true {
                                                                v5 -= genFn0(false, 2)
                                                                v3 = append(v3, 42)
                                                                var v77 int32 = 2147483647
                                                                if (-1 > -23) {
                                                                    v78 := GenS15{F0: 255, F1: 18446744073709551615, F2: (fi63 % (((fi21 + int(v5))) | 1))}
                                                                    _ = v78
                                                                }
                                                                _ = v77
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v3 = append(v3, 0)
                                                }
                                                if (42 >= 32767) {
                                                    v4 = append(v4, (-(fi21)))
                                                    if len(v4) > 0 {
                                                        v4[0] = 353
                                                    }
                                                } else {
                                                    var v79 int = fi21
                                                    v80 := GenS16{F0: v15}
                                                    var v81 float32 = (v14 + 0.1)
                                                    _ = v79
                                                    _ = v80
                                                    _ = v81
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                if false {
                                                    switch fi63 {
                                                    case -921:
                                                        v3 = append(v3, 0)
                                                    case -9223372036854775808:
                                                        v82 := GenS17{F0: fi21, F1: 0, F2: (!(false || false))}
                                                        _ = v82
                                                    }
                                                    if true {
                                                        var v83 int64 = -784
                                                        v3 = append(v3, 0)
                                                        v84 := GenS13{F0: 1, F1: v83}
                                                        v15 += (v5 & 17)
                                                        _ = len(v4)
                                                        _ = v83
                                                        _ = v84
                                                    } else {
                                                        fmt.Println(fi21, fi21, fi63)
                                                        v5 = (^(v5))
                                                        if len(v3) > 0 {
                                                            v3[0] = 189
                                                        }
                                                    }
                                                    if ((28 <= 0) || false) {
                                                        var v85 int32 = 42
                                                        if false {
                                                            fmt.Println(fi21, v15)
                                                            if (v85 > -1) {
                                                                var v86 uint32 = 908
                                                                _ = v86
                                                            }
                                                            v15 = v15
                                                        }
                                                        switch "a" {
                                                        case "hello":
                                                            var v87 uint32 = 1
                                                            _ = v87
                                                        case "":
                                                            switch fi63 {
                                                            case 4294967295:
                                                                var v88 int64 = -1
                                                                _ = v88
                                                            default:
                                                                if false {
                                                                    var v89 string = (v1 + v1)
                                                                    var v90 float64 = v2
                                                                    var v91 uint8 = 255
                                                                    _ = v89
                                                                    _ = v90
                                                                    _ = v91
                                                                } else {
                                                                    fmt.Println(fi63)
                                                                    fmt.Println(v5, v15)
                                                                }
                                                            }
                                                            v14 += v14
                                                            v2 = v2
                                                        default:
                                                            v4 = append(v4, fi21)
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                        _ = v85
                                                    }
                                                }
                                                v92 := GenS18{F0: -87, F1: false}
                                                _ = v92
                                            }
                                        }
                                        switch fi63 {
                                        case -638:
                                            fmt.Println(v15, fi63, v5)
                                        case 0:
                                            var v93 uint32 = uint32(v5)
                                            _ = v93
                                        default:
                                            fmt.Println(v15, v15, v5)
                                        }
                                        if (-2147483648 < 2147483647) {
                                            v94 := GenS19{F0: 566, F1: -1, F2: 2}
                                            switch fi63 {
                                            case -318:
                                                v15 -= 1
                                                if true {
                                                    if (!(!false)) {
                                                        v14 += (-2.25 * 3.14159)
                                                        if (genFn0((!true), 952) > v5) {
                                                            if (!(fi21 != (fi21 + 578))) {
                                                                v14 -= 0.0
                                                                if ((!false) || (v14 != v14)) {
                                                                    v95 := GenS20{F0: -1, F1: (v5 / ((v15) | 1))}
                                                                    var v96 float64 = float64((v14 * (1.5 / (v14 + v14))))
                                                                    _ = v95
                                                                    _ = v96
                                                                } else {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 2147483648
                                                                    }
                                                                    v15 = v5
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        var v97 uint8 = 42
                                                                        _ = v97
                                                                    case "\n\t":
                                                                        v98 := GenS17{F0: (fi63 + fi21), F1: 4294967295, F2: (988 != 0)}
                                                                        _ = v98
                                                                    default:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 107
                                                                        }
                                                                    }
                                                                    switch -9223372036854775808 {
                                                                    case 1:
                                                                        fmt.Println(fi21)
                                                                    case 128:
                                                                        if false {
                                                                            fmt.Println(v2)
                                                                            v14 = 3.14159
                                                                        }
                                                                    case -595:
                                                                        if false {
                                                                            switch (fi21 + fi21) {
                                                                            case -9223372036854775808:
                                                                                v2 = 0.1
                                                                            case 32768:
                                                                                v99 := GenS8{F0: fi63, F1: (^(fi63)), F2: 1000}
                                                                                _ = v99
                                                                            case 32767:
                                                                                var v100 int32 = 2147483647
                                                                                _ = v100
                                                                            default:
                                                                                if true {
                                                                                    _ = len(v4)
                                                                                    if true {
                                                                                        var v101 uint = 2
                                                                                        v5 += v15
                                                                                        if ("" >= "") {
                                                                                            _ = len(v4)
                                                                                            fmt.Println(fi21, v101)
                                                                                        }
                                                                                        v102 := GenS10{F0: 128, F1: -1}
                                                                                        _ = v101
                                                                                        _ = v102
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = len(v4)
                                                                            if false {
                                                                                switch fi21 {
                                                                                case 42:
                                                                                    v103 := GenS14{F0: (fi21 + (fi63 + fi63)), F1: (!true), F2: 2}
                                                                                    _ = v103
                                                                                }
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 2147483647
                                                                                }
                                                                            } else {
                                                                                if true {
                                                                                    _ = len(v3)
                                                                                    v5 += (v15 << ((196) & 63))
                                                                                    v104 := GenS20{F0: 9223372036854775807, F1: (v15 << ((42) & 63))}
                                                                                    _ = v104
                                                                                } else {
                                                                                    v14 -= v14
                                                                                    v105 := GenS21{F0: 163}
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = (fi21 | fi63)
                                                                                    }
                                                                                    v106 := GenS22{F0: 42, F1: (42 != 98), F2: 1}
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    fmt.Println(fi63)
                                                                                    _ = v105
                                                                                    _ = v106
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v3 = append(v3, 61)
                                                                            var v107 int64 = 196
                                                                            var v108 uint = 65535
                                                                            _ = v107
                                                                            _ = v108
                                                                        }
                                                                    default:
                                                                        var v109 int32 = 1
                                                                        _ = v109
                                                                    }
                                                                    fmt.Println(fi63)
                                                                    if (true || (!(623 == 65535))) {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        var v110 uint32 = 2147483647
                                                                        var v111 uint = 144
                                                                        _ = v110
                                                                        _ = v111
                                                                    } else {
                                                                        v1 = ("hello" + v1)
                                                                        v112 := GenS18{F0: -59, F1: true}
                                                                        _ = v112
                                                                    }
                                                                }
                                                                switch fi21 {
                                                                case 32767:
                                                                    var v113 float32 = ((float32(1e10) * v14) / float32((((0.0 / v2) * (((v2 / 3.14159) - v2) / 3.14159)) + v2)))
                                                                    _ = v113
                                                                case 127:
                                                                    switch (fi63 % ((fi21) | 1)) {
                                                                    case 1:
                                                                        fmt.Println(v2, v5)
                                                                    case 128:
                                                                        v14 -= v14
                                                                    }
                                                                case 9223372036854775807:
                                                                    v5 -= v5
                                                                default:
                                                                    v114 := GenS4{F0: (v15 % (((v15 >> ((2) & 63))) | 1)), F1: int16(fi21)}
                                                                    _ = v114
                                                                }
                                                                if false {
                                                                    var v115 int = (fi21 * (fi21 ^ fi63))
                                                                    v116 := GenS7{F0: 0}
                                                                    v117 := GenS2{F0: 2, F1: 62}
                                                                    v118 := GenS23{F0: 890, F1: v115, F2: 42}
                                                                    _ = v115
                                                                    _ = v116
                                                                    _ = v117
                                                                    _ = v118
                                                                }
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            } else {
                                                                v119 := GenS4{F0: (v15 - genFn0(false, 268)), F1: -999}
                                                                _ = v119
                                                            }
                                                            var v120 uint = 4294967295
                                                            if (true && (v14 <= float32(v2))) {
                                                                if ((fi21 % ((fi21) | 1)) == 111) {
                                                                    var v121 bool = true
                                                                    var v122 uint64 = 18446744073709551615
                                                                    switch (fi21 & (fi63 >> (((^(v120))) & 63))) {
                                                                    case 1:
                                                                        if v121 {
                                                                            if (int(v122) <= fi21) {
                                                                                fmt.Println(fi63)
                                                                                switch v1 {
                                                                                case "a":
                                                                                    var v123 uint16 = genFn0(v121, 1)
                                                                                    _ = v123
                                                                                }
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = 4294967295
                                                                                }
                                                                                v121 = (455 <= 42)
                                                                            } else {
                                                                                var v124 int64 = 9223372036854775807
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    v124 += -899
                                                                                }
                                                                                switch v1 {
                                                                                case "\n\t":
                                                                                    v125 := GenS14{F0: fi21, F1: (!v121), F2: 0}
                                                                                    _ = v125
                                                                                }
                                                                                _ = v124
                                                                            }
                                                                            v15 = (v15 ^ v5)
                                                                            var v126 string = "\n\t"
                                                                            _ = v126
                                                                        } else {
                                                                            var v127 int8 = -128
                                                                            switch 9223372036854775807 {
                                                                            case 32768:
                                                                                fmt.Println(v5, v122)
                                                                            case 0:
                                                                                fmt.Println(v120, v120, v120)
                                                                            case -974:
                                                                                var v128 string = v1
                                                                                _ = v128
                                                                            }
                                                                            var v129 int64 = 9223372036854775807
                                                                            _ = v127
                                                                            _ = v129
                                                                        }
                                                                    case 2147483647:
                                                                        v130 := GenS3{F0: 540}
                                                                        _ = v130
                                                                    }
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 2
                                                                    }
                                                                    fmt.Println(v120, v2)
                                                                    _ = v121
                                                                    _ = v122
                                                                }
                                                                switch ("foo bar" + "\n\t") {
                                                                case "a":
                                                                    v131 := GenS24{F0: (695 > 2147483648), F1: (("a" + v1) + (v1 + v1))}
                                                                    _ = v131
                                                                case "hello":
                                                                    var v132 uint8 = 1
                                                                    _ = v132
                                                                default:
                                                                    v133 := GenS25{F0: 196, F1: 2147483647, F2: v5}
                                                                    _ = v133
                                                                }
                                                            } else {
                                                                var v134 uint64 = 4294967295
                                                                v5 += (v5 - 255)
                                                                var v135 uint64 = v134
                                                                fmt.Println(fi63)
                                                                v2 = ((0.1 - (0.1 - 3.14159)) / v2)
                                                                _ = v134
                                                                _ = v135
                                                            }
                                                            _ = v120
                                                        }
                                                    }
                                                }
                                                if (0 >= 42) {
                                                    var v136 string = v1
                                                    v5 -= genFn0((650 <= 0), 0)
                                                    fmt.Println(fi21, v136, v1)
                                                    v5 += (v5 >> ((255) & 63))
                                                    _ = v136
                                                } else {
                                                    switch fi21 {
                                                    case 621:
                                                        switch (fi21 & -1) {
                                                        case -9223372036854775808:
                                                            switch (("" + ((((((v1 + "foo bar") + v1) + ("" + (v1 + "hello"))) + v1) + v1) + (v1 + ("hello" + ((v1 + v1) + (v1 + "hello")))))) + "hello") {
                                                            case "":
                                                                if false {
                                                                    v2 -= ((3.14159 - ((v2 / float64((float32(float64((v14 + 3.14159))) + 0.1))) - (-2.25 * v2))) + v2)
                                                                    var v137 bool = true
                                                                    var v138 uint32 = 0
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (^(v138))
                                                                    }
                                                                    v138 = v138
                                                                    v139 := GenS2{F0: 4294967295, F1: 955}
                                                                    var v140 string = v1
                                                                    if false {
                                                                        fmt.Println(v137, v138, v15)
                                                                        v15 = (v5 / ((v15) | 1))
                                                                    } else {
                                                                        var v141 float64 = (0.0 * v2)
                                                                        if false {
                                                                            fmt.Println(v138)
                                                                            fmt.Println(v140, fi63, v140)
                                                                        }
                                                                        v142 := GenS1{F0: 128, F1: 0}
                                                                        _ = v141
                                                                        _ = v142
                                                                    }
                                                                    _ = len(v3)
                                                                    _ = v137
                                                                    _ = v138
                                                                    _ = v139
                                                                    _ = v140
                                                                }
                                                            case "foo bar":
                                                                v143 := GenS9{F0: (v1 >= ((v1 + ("hello" + ("a" + (v1 + v1)))) + v1)), F1: 42}
                                                                _ = v143
                                                            case "\n\t":
                                                                fmt.Println(v1)
                                                            }
                                                        case -1:
                                                            if false {
                                                                if true {
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        var v144 int32 = 42
                                                                        _ = v144
                                                                    }
                                                                } else {
                                                                    v5 += (v5 | v5)
                                                                    if (false && (757 >= 2)) {
                                                                        switch v1 {
                                                                        case "hello":
                                                                            v2 += (1e10 * (v2 * float64(float32(v2))))
                                                                        case "a":
                                                                            if true {
                                                                                var v145 int8 = -108
                                                                                fmt.Println(v5, v145)
                                                                                _ = v145
                                                                            }
                                                                        case "foo bar":
                                                                            var v146 string = "\n\t"
                                                                            switch (v1 + v1) {
                                                                            case "\n\t":
                                                                                fmt.Println(v5, fi21, v2)
                                                                            case "hello":
                                                                                fmt.Println(v15, v146)
                                                                            case "foo bar":
                                                                                var v147 uint16 = genFn0((v5 > (v5 - (v5 * 662))), 2)
                                                                                _ = v147
                                                                            default:
                                                                                var v148 string = "\n\t"
                                                                                _ = v148
                                                                            }
                                                                            _ = v146
                                                                        default:
                                                                            var v149 float64 = ((float64(v14) / (1.5 - float64(1e10))) + v2)
                                                                            _ = v149
                                                                        }
                                                                        var v150 uint = 65535
                                                                        switch ((v1 + (v1 + (v1 + ""))) + ((v1 + v1) + v1)) {
                                                                        case "a":
                                                                            switch (fi63 & (fi63 * fi63)) {
                                                                            case 0:
                                                                                v4 = append(v4, (fi63 % ((127) | 1)))
                                                                            default:
                                                                                v150 = 0
                                                                            }
                                                                        }
                                                                        _ = v150
                                                                    }
                                                                }
                                                                fmt.Println(v14)
                                                            } else {
                                                                var v151 uint32 = 4294967295
                                                                _ = v151
                                                            }
                                                        }
                                                    case 32768:
                                                        if (2147483647 < 561) {
                                                            if len(v4) > 0 {
                                                                v4[0] = 32767
                                                            }
                                                            if (!true) {
                                                                if true {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 65535
                                                                    }
                                                                }
                                                                fmt.Println(v1, v2)
                                                            }
                                                            if (42 == -1) {
                                                                v152 := GenS15{F0: 2, F1: uint(v5), F2: (fi63 & (fi21 % ((fi63) | 1)))}
                                                                var v153 string = v1
                                                                _ = v152
                                                                _ = v153
                                                            } else {
                                                                switch (v1 + v1) {
                                                                case "a":
                                                                    v154 := GenS15{F0: 2, F1: 4294967295, F2: (fi63 / ((892) | 1))}
                                                                    _ = v154
                                                                case "\n\t":
                                                                    v3 = append(v3, 65535)
                                                                default:
                                                                    v155 := GenS16{F0: (v15 % (((^(v5))) | 1))}
                                                                    _ = v155
                                                                }
                                                                var v156 uint32 = 42
                                                                _ = v156
                                                            }
                                                        } else {
                                                            v15 += v15
                                                            if false {
                                                                switch 128 {
                                                                case 4294967295:
                                                                    _ = len(v3)
                                                                }
                                                                if (-39 < 0) {
                                                                    switch (fi21 * fi63) {
                                                                    case -584:
                                                                        v157 := GenS21{F0: 127}
                                                                        _ = v157
                                                                    default:
                                                                        var v158 string = ""
                                                                        _ = v158
                                                                    }
                                                                    var v159 uint32 = 1
                                                                    _ = v159
                                                                }
                                                            } else {
                                                                v5 = v5
                                                                v160 := GenS19{F0: 2147483647, F1: 9223372036854775807, F2: 266}
                                                                _ = v160
                                                            }
                                                        }
                                                    case 156:
                                                        if len(v4) > 0 {
                                                            v4[0] = fi21
                                                        }
                                                    }
                                                    var v161 uint8 = 128
                                                    _ = v161
                                                }
                                            default:
                                                if false {
                                                    var v162 float32 = v14
                                                    if (0.0 == v14) {
                                                        v15 = (v5 >> ((968) & 63))
                                                        fmt.Println(v15, v5, v2)
                                                        var v163 int32 = 1
                                                        _ = v163
                                                    } else {
                                                        fmt.Println(fi21, v1, fi21)
                                                        v2 += (float64((0.1 / float32(v2))) - v2)
                                                        v164 := GenS26{F0: 694}
                                                        _ = v164
                                                    }
                                                    _ = v162
                                                }
                                            }
                                            _ = v94
                                        } else {
                                            var v165 int = (fi63 - (fi63 + 1))
                                            v4 = append(v4, fi63)
                                            v5 -= v5
                                            _ = v165
                                        }
                                    }
                                    var v166 int = (-(fi21))
                                    v2 -= ((-2.25 - (v2 - v2)) / (((v2 + v2) / float64(((1.5 * 1e10) * v14))) * v2))
                                    _ = fi63
                                    _ = v166
                                }
                                fmt.Println(v14)
                            } else {
                                v15 = (v15 % ((32767) | 1))
                                v15 += v5
                                var v167 int = -2147483648
                                _ = v167
                            }
                            v2 = 1.5
                        }
                        v168 := GenS8{F0: fi21, F1: fi21, F2: 32767}
                        _ = v168
                    } else {
                        var v169 bool = false
                        for fi170 := 0; fi170 < 8; fi170++ {
                            fmt.Println(v5, fi170, v15)
                            v171 := GenS27{F0: (v15 + v5)}
                            _ = fi170
                            _ = v171
                        }
                        _ = v169
                    }
                }
            }
        }
        _ = fi21
    }
    var v172 uint64 = 136
    if true {
        v5 = 32767
    }
    var v173 bool = false
    for fi174 := 0; fi174 < 7; fi174++ {
        var v175 uint8 = 135
        _ = fi174
        _ = v175
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v172:", v172)
    fmt.Println("v173:", v173)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v172
    _ = v173
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
