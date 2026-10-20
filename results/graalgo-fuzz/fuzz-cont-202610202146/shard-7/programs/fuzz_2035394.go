package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint64; F1 int; F2 string }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 int64; F1 uint64; F2 int32 }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int; F1 int; F2 uint8 }
type GenS10 struct { F0 int64; F1 int16; F2 string }
type GenS11 struct { F0 int; F1 int16 }
type GenS12 struct { F0 int32; F1 uint16 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int64; F1 int8; F2 uint32 }
type GenS16 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS17 struct { F0 uint32; F1 uint16 }
type GenS18 struct { F0 uint32; F1 int }
type GenS19 struct { F0 int }
type GenS20 struct { F0 int; F1 string; F2 int8 }
type GenS21 struct { F0 string; F1 int16; F2 uint16 }
type GenS22 struct { F0 bool; F1 string; F2 uint64 }
type GenS23 struct { F0 int32; F1 int8 }
type GenS24 struct { F0 bool; F1 uint16 }
type GenS25 struct { F0 int32; F1 int; F2 int8 }
type GenS26 struct { F0 int8; F1 uint; F2 int16 }
type GenS27 struct { F0 uint; F1 string }
type GenS28 struct { F0 uint16; F1 int; F2 string }
type GenS29 struct { F0 uint16 }
type GenS30 struct { F0 uint; F1 int8 }
type GenS31 struct { F0 uint; F1 uint; F2 uint8 }
type GenS32 struct { F0 string; F1 uint8 }
type GenS33 struct { F0 string }
type GenS34 struct { F0 uint8; F1 uint; F2 int8 }
type GenS35 struct { F0 uint32; F1 uint; F2 int64 }
type GenS36 struct { F0 string; F1 bool; F2 int32 }
type GenS37 struct { F0 int8; F1 int32; F2 string }
type GenS38 struct { F0 uint32; F1 int8 }
type GenS39 struct { F0 uint }
type GenS40 struct { F0 int32; F1 bool; F2 int32 }
type GenS41 struct { F0 int64 }
func genFn0(p0 string) uint16 {
    v7 := GenS1{F0: (("" + "foo bar") + (p0 + (p0 + "a")))}
    var v8 int = 60
    var v9 uint = 2
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 317
}
func genFn1(p0 int) int64 {
    var v10 int = p0
    p0 += 42
    var v11 float32 = (3.14159 * 1e10)
    var v12 string = "\n\t"
    v13 := []int{(v10 % (((p0 % (((v10 % (((p0 * -668)) | 1))) | 1))) | 1)), p0}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -2147483648
    v3 := []int{88, 42}
    v4 := []uint32{}
    var v5 int16 = 1
    var v6 int16 = v5
    v5 = -730
    switch 32767 {
    case 42:
        fmt.Println(v6)
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 uint32 = 739
        var v16 uint = 282
        _ = fi14
        _ = v15
        _ = v16
    }
    if ((!true) || true) {
        if len(v3) > 0 {
            v3[0] = 32767
        }
        var v17 int64 = 9223372036854775807
        var v18 int16 = (v6 / ((42) | 1))
        _ = v17
        _ = v18
    }
    switch 32768 {
    case -1:
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 float64 = 1e10
            v5 -= v6
            delete(v1, 2)
            _ = fi19
            _ = v20
        }
    case 9223372036854775807:
        fmt.Println(v6)
    case 140:
        v1[4294967295] = (32768 > -9223372036854775808)
    default:
        v5 = v6
    }
    if false {
        if false {
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 int8 = 122
                _ = fi21
                _ = v22
            }
            if (true || true) {
                switch ("hello" + ("\n\t" + "")) {
                case "hello":
                    var v23 uint16 = 42
                    _ = v23
                case "a":
                    v5 = 1
                default:
                    var v24 uint64 = 18446744073709551615
                    _ = v24
                }
                var v25 uint = 998
                switch ("foo bar" + "foo bar") {
                case "hello":
                    v25 += 65535
                case "a":
                    var v26 int32 = (v2 - 103)
                    _ = v26
                }
                _ = v25
            } else {
                delete(v1, 0)
                var v27 int16 = v6
                _ = v27
            }
        }
        var v28 int16 = (-(v6))
        delete(v1, 65535)
        _ = v28
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        for fi30 := 0; fi30 < 10; fi30++ {
            var v31 uint64 = 18446744073709551615
            if ("hello" <= ("\n\t" + ("foo bar" + "\n\t"))) {
                switch fi29 {
                case 42:
                    if (fi30 <= (fi30 % (((fi29 & 247)) | 1))) {
                        fmt.Println(v6)
                    } else {
                        v32 := v1[uint32(v31)]
                        _ = v32
                    }
                    fmt.Println(v6)
                }
            }
            _ = fi30
            _ = v31
        }
        var v33 int8 = 20
        _ = fi29
        _ = v33
    }
    if false {
        var v34 int16 = v6
        for fi35 := 0; fi35 < 7; fi35++ {
            if false {
                switch (("hello" + (("" + "a") + "\n\t")) + "foo bar") {
                case "foo bar":
                    v36 := GenS1{F0: ("a" + "\n\t")}
                    _ = v36
                case "":
                    v37 := GenS1{F0: "a"}
                    _ = v37
                default:
                    if (215 > 42) {
                        v38 := GenS2{F0: 65535, F1: 2}
                        fmt.Println(fi35)
                        _ = len(v1)
                        _ = v38
                    } else {
                        switch "foo bar" {
                        case "hello":
                            var v39 bool = true
                            _ = v39
                        case "foo bar":
                            var v40 string = ("hello" + "")
                            _ = v40
                        case "":
                            v41 := GenS1{F0: ""}
                            _ = v41
                        }
                        var v42 bool = false
                        var v43 bool = v42
                        _ = v42
                        _ = v43
                    }
                }
                if (-2.25 == float64(float32((1.5 + 1.5)))) {
                    v34 -= v34
                    for fi44 := 0; fi44 < 1; fi44++ {
                        switch fi35 {
                        case -624:
                            if ((0.0 > 0.0) && false) {
                                var v45 float32 = 3.14159
                                _ = v45
                            } else {
                                switch -327 {
                                case -234:
                                    var v46 uint64 = 42
                                    _ = v46
                                case -9223372036854775808:
                                    fmt.Println(fi44, v2, v6)
                                    v5 = v34
                                    switch "foo bar" {
                                    case "hello":
                                        switch "a" {
                                        case "\n\t":
                                            if (!true) {
                                                v47 := GenS3{F0: (fi35 - fi35), F1: 2147483648, F2: (v2 * (-(v2)))}
                                                if true {
                                                    if false {
                                                        v34 = (-(v34))
                                                        v1[42] = false
                                                        v48 := v1[2147483648]
                                                        _ = v48
                                                    } else {
                                                        var v49 float64 = 0.1
                                                        _ = v49
                                                    }
                                                    v6 += v34
                                                } else {
                                                    if true {
                                                        v2 -= (v2 ^ (v2 << ((111) & 63)))
                                                        fmt.Println(v2)
                                                        v34 = (v5 / ((382) | 1))
                                                        delete(v1, 2147483647)
                                                        if (2 != 832) {
                                                            v2 = 232
                                                            switch "hello" {
                                                            case "":
                                                                v2 = 214
                                                            case "a":
                                                                v5 -= int16(fi35)
                                                            }
                                                            if true {
                                                                v50 := GenS3{F0: fi44, F1: 495, F2: -2147483648}
                                                                _ = v50
                                                            } else {
                                                                var v51 int8 = -115
                                                                v52 := GenS1{F0: "a"}
                                                                if false {
                                                                    v3 = append(v3, (fi35 ^ -845))
                                                                } else {
                                                                    var v53 int32 = 42
                                                                    var v54 float64 = (0.1 * ((((0.1 - 3.14159) / 1e10) + -2.25) / ((-2.25 / 3.14159) / 0.1)))
                                                                    v55 := v1[2]
                                                                    _ = v53
                                                                    _ = v54
                                                                    _ = v55
                                                                }
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                            v5 = 1
                                                        } else {
                                                            v56 := GenS2{F0: 18446744073709551615, F1: 42}
                                                            var v57 bool = (false || false)
                                                            switch fi35 {
                                                            case 32768:
                                                                v3 = append(v3, fi44)
                                                            case 2147483647:
                                                                switch fi44 {
                                                                case 0:
                                                                    if v57 {
                                                                        v5 = v6
                                                                    } else {
                                                                        if (0 >= -199) {
                                                                            fmt.Println(v57, fi35, v5)
                                                                        } else {
                                                                            v58 := v1[2]
                                                                            _ = v58
                                                                        }
                                                                        v2 -= v2
                                                                        v1[65535] = true
                                                                        if (!true) {
                                                                            v57 = v57
                                                                        }
                                                                    }
                                                                case -427:
                                                                    var v59 uint32 = 0
                                                                    _ = v59
                                                                case 59:
                                                                    v2 -= (v2 >> ((716) & 63))
                                                                default:
                                                                    v60 := GenS4{F0: 2, F1: (fi35 ^ fi44), F2: ("foo bar" + "a")}
                                                                    _ = v60
                                                                }
                                                            default:
                                                                if (1 < 1) {
                                                                    v61 := GenS2{F0: 4294967295, F1: 0}
                                                                    fmt.Println(v5, v5)
                                                                    _ = v61
                                                                }
                                                            }
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        var v62 int = (fi44 % ((fi35) | 1))
                                                        if false {
                                                            fmt.Println(fi35, fi35, v5)
                                                        }
                                                        fmt.Println(fi35, v6, v62)
                                                        v5 += v34
                                                        var v63 uint16 = 347
                                                        if false {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                var v64 bool = false
                                                                _ = v64
                                                            case "":
                                                                if true {
                                                                    switch "a" {
                                                                    case "a":
                                                                        v65 := GenS5{F0: 4294967295, F1: 961}
                                                                        _ = v65
                                                                    }
                                                                    if ((42 > 42) && true) {
                                                                        v63 = (v63 | v63)
                                                                        switch int(v6) {
                                                                        case -581:
                                                                            v63 -= v63
                                                                        case -525:
                                                                            v66 := GenS5{F0: 8, F1: 408}
                                                                            _ = v66
                                                                        case 127:
                                                                            v67 := v1[4294967295]
                                                                            _ = v67
                                                                        default:
                                                                            var v68 bool = (!false)
                                                                            _ = v68
                                                                        }
                                                                    }
                                                                    v69 := GenS6{F0: 42, F1: 1, F2: (v2 ^ 564)}
                                                                    v70 := GenS7{F0: 4294967295, F1: 65535}
                                                                    _ = v69
                                                                    _ = v70
                                                                } else {
                                                                    v6 = v34
                                                                    var v71 int8 = -128
                                                                    _ = len(v4)
                                                                    var v72 uint64 = 18446744073709551615
                                                                    _ = v71
                                                                    _ = v72
                                                                }
                                                            case "foo bar":
                                                                v73 := GenS3{F0: int(v6), F1: 1, F2: (v2 & -889)}
                                                                _ = v73
                                                            default:
                                                                v2 += v2
                                                            }
                                                        }
                                                        v2 = v2
                                                        fmt.Println(v63)
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                    _ = len(v3)
                                                }
                                                v74 := GenS5{F0: 1, F1: 461}
                                                _ = v47
                                                _ = v74
                                            } else {
                                                var v75 string = ("hello" + (("" + (("" + ("hello" + "")) + "hello")) + ("hello" + "\n\t")))
                                                fmt.Println(v75, v2)
                                                _ = v75
                                            }
                                        }
                                    case "\n\t":
                                        var v76 int8 = 42
                                        _ = v76
                                    }
                                }
                            }
                        case 32768:
                            v34 = int16(v2)
                        default:
                            var v77 uint = 4294967295
                            _ = v77
                        }
                        _ = fi44
                    }
                    if false {
                        v78 := GenS8{F0: (!true)}
                        delete(v1, 282)
                        _ = v78
                    }
                } else {
                    if false {
                        for fi79 := 0; fi79 < 4; fi79++ {
                            if (1 == v2) {
                                if false {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if false {
                                        if true {
                                            if true {
                                                switch (^(fi35)) {
                                                case 128:
                                                    var v80 int16 = v5
                                                    _ = v80
                                                case -592:
                                                    switch fi35 {
                                                    case 4294967295:
                                                        var v81 float64 = (1e10 * ((1.5 / (1.5 + (float64(3.14159) / (0.0 + 1.5)))) + 0.0))
                                                        _ = v81
                                                    }
                                                case 488:
                                                    v5 = (v5 >> ((0) & 63))
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if ((false && false) && (((true && true) || (!true)) || false)) {
                                                            v82 := GenS9{F0: int(v6), F1: 0, F2: 1}
                                                            var v83 float64 = (0.0 + (1.5 * 1.5))
                                                            switch (fi35 / (((fi79 + 128)) | 1)) {
                                                            case 127:
                                                                if true {
                                                                    var v84 int = fi35
                                                                    _ = v84
                                                                } else {
                                                                    var v85 int = (fi79 * -9223372036854775808)
                                                                    _ = v85
                                                                }
                                                            default:
                                                                var v86 int = fi79
                                                                _ = v86
                                                            }
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                    default:
                                                        _ = len(v3)
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                default:
                                                    var v87 uint = 808
                                                    _ = v87
                                                }
                                                if true {
                                                    fmt.Println(fi79, v2)
                                                } else {
                                                    if (("" < "") || ((42 != 2) && false)) {
                                                        if true {
                                                            var v88 int16 = (v5 | v6)
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            if false {
                                                                v2 += (v2 * v2)
                                                                var v89 int16 = 0
                                                                _ = v89
                                                            } else {
                                                                if (100 != 42) {
                                                                    v2 = v2
                                                                    var v90 string = ("" + "foo bar")
                                                                    _ = v90
                                                                } else {
                                                                    v91 := GenS10{F0: 42, F1: 0, F2: "foo bar"}
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 2147483648
                                                                    }
                                                                    fmt.Println(v5, v34)
                                                                    v92 := GenS5{F0: 2, F1: 245}
                                                                    v93 := GenS11{F0: (-(fi79)), F1: (v5 >> ((894) & 63))}
                                                                    _ = v91
                                                                    _ = v92
                                                                    _ = v93
                                                                }
                                                                if (v34 != -32768) {
                                                                    var v94 float64 = ((float64(3.14159) * 3.14159) - 0.0)
                                                                    if (!true) {
                                                                        var v95 uint64 = 524
                                                                        _ = v95
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            v2 = (v2 % ((v2) | 1))
                                                                        case "hello":
                                                                            v34 = v5
                                                                        case "\n\t":
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                var v96 uint = 4294967295
                                                                                _ = v96
                                                                            case "foo bar":
                                                                                var v97 string = ("hello" + "foo bar")
                                                                                _ = v97
                                                                            case "a":
                                                                                var v98 uint8 = 149
                                                                                _ = v98
                                                                            default:
                                                                                var v99 uint = 65535
                                                                                _ = v99
                                                                            }
                                                                        }
                                                                        switch ("a" + ("a" + "foo bar")) {
                                                                        case "\n\t":
                                                                            switch fi35 {
                                                                            case 42:
                                                                                var v100 int16 = (v6 << ((255) & 63))
                                                                                _ = v100
                                                                            }
                                                                        default:
                                                                            if (2 != 1) {
                                                                                v1[2147483647] = (!true)
                                                                            }
                                                                        }
                                                                    }
                                                                    v34 -= v88
                                                                    _ = v94
                                                                } else {
                                                                    var v101 int64 = 146
                                                                    _ = v101
                                                                }
                                                                if ("\n\t" != "hello") {
                                                                    v2 = v2
                                                                    if (2 == 0) {
                                                                        v5 = (-(v34))
                                                                        var v102 int16 = (v34 & (v5 & (^(v5))))
                                                                        _ = v102
                                                                    } else {
                                                                        var v103 int8 = -38
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        v88 = (v5 | -197)
                                                                        _ = v103
                                                                    }
                                                                }
                                                                if (false || ((false && (!true)) || (!((^(v88)) < 42)))) {
                                                                    if false {
                                                                        v104 := GenS5{F0: 347, F1: 2}
                                                                        if true {
                                                                            delete(v1, 2)
                                                                            if false {
                                                                                var v105 uint = 1
                                                                                _ = v105
                                                                            } else {
                                                                                v106 := GenS12{F0: v2, F1: 32767}
                                                                                if (42 <= 2) {
                                                                                    var v107 bool = false
                                                                                    var v108 int32 = (v2 & v2)
                                                                                    _ = v107
                                                                                    _ = v108
                                                                                }
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = (-(fi35))
                                                                                }
                                                                                if (true && (false || ((true && (1e10 < float32((0.0 / -2.25)))) && (false || (v88 == (-(v34))))))) {
                                                                                    if true {
                                                                                        var v109 int = (fi35 % ((9223372036854775807) | 1))
                                                                                        _ = v109
                                                                                    } else {
                                                                                        var v110 uint32 = 2
                                                                                        fmt.Println(fi35)
                                                                                        _ = v110
                                                                                    }
                                                                                    _ = len(v1)
                                                                                } else {
                                                                                    fmt.Println(v88)
                                                                                    var v111 float32 = (0.1 + 0.1)
                                                                                    _ = v111
                                                                                }
                                                                                _ = v106
                                                                            }
                                                                        }
                                                                        _ = v104
                                                                    } else {
                                                                        v88 += v88
                                                                        fmt.Println(v88, v6)
                                                                    }
                                                                } else {
                                                                    if (42 != 0) {
                                                                        var v112 int64 = genFn1((fi35 >> ((4294967295) & 63)))
                                                                        v88 += (v5 % (((v34 + (v6 * v88))) | 1))
                                                                        _ = len(v3)
                                                                        v113 := GenS13{F0: 42}
                                                                        var v114 uint16 = 0
                                                                        _ = v112
                                                                        _ = v113
                                                                        _ = v114
                                                                    } else {
                                                                        if ("\n\t" > (((("" + ("" + "a")) + ("hello" + ((("hello" + "foo bar") + "\n\t") + (((("\n\t" + "hello") + "hello") + ("hello" + "foo bar")) + ("hello" + "\n\t"))))) + "") + ("foo bar" + "hello"))) {
                                                                            v6 += 0
                                                                            v6 += (v88 | v34)
                                                                        }
                                                                        if (42 <= 1) {
                                                                            if ((0.1 / ((-2.25 + ((0.1 / -2.25) - 1.5)) + ((-2.25 * 1e10) - float32(float64((-2.25 - ((0.0 - ((0.0 * 0.0) - 0.1)) + 1.5))))))) != (1.5 - (0.1 - float32(-2.25)))) {
                                                                                v5 = v5
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v34 = v88
                                                                        }
                                                                        switch (fi35 * 680) {
                                                                        case 1:
                                                                            fmt.Println(v34, v88)
                                                                        case -9223372036854775808:
                                                                            if true {
                                                                                fmt.Println(v5, v34, v5)
                                                                                fmt.Println(v88, v6, v2)
                                                                                v5 += -446
                                                                            } else {
                                                                                v115 := GenS7{F0: 18446744073709551615, F1: 4294967295}
                                                                                switch -111 {
                                                                                case 4294967295:
                                                                                    v34 += (v34 << ((1) & 63))
                                                                                case -333:
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        var v116 int16 = (v6 ^ (v5 | v6))
                                                                                        _ = v116
                                                                                    case "a":
                                                                                        v117 := GenS6{F0: 995, F1: 2, F2: (v2 & v2)}
                                                                                        switch "a" {
                                                                                        case "":
                                                                                            switch fi35 {
                                                                                            case 0:
                                                                                                if (true || false) {
                                                                                                    var v118 uint16 = 114
                                                                                                    v6 -= v6
                                                                                                    _ = v118
                                                                                                }
                                                                                            }
                                                                                        case "\n\t":
                                                                                            if (0.0 > ((0.0 / 1e10) + 0.0)) {
                                                                                                fmt.Println(v88, v6)
                                                                                                switch "" {
                                                                                                case "a":
                                                                                                    v88 -= v34
                                                                                                }
                                                                                            } else {
                                                                                                if false {
                                                                                                    var v119 int16 = (v5 / (((v5 * (v88 * 765))) | 1))
                                                                                                    _ = v119
                                                                                                } else {
                                                                                                    v88 += v6
                                                                                                    v120 := GenS12{F0: v2, F1: 755}
                                                                                                    _ = v120
                                                                                                }
                                                                                                v121 := GenS14{F0: (496 >= 600)}
                                                                                                _ = v121
                                                                                            }
                                                                                        case "foo bar":
                                                                                            v122 := GenS4{F0: 404, F1: fi35, F2: "\n\t"}
                                                                                            _ = v122
                                                                                        }
                                                                                        if (915 == 1) {
                                                                                            v4 = append(v4, 2147483648)
                                                                                        }
                                                                                        _ = v117
                                                                                    }
                                                                                }
                                                                                _ = len(v3)
                                                                                _ = v115
                                                                            }
                                                                        case 2147483647:
                                                                            delete(v1, 2)
                                                                        default:
                                                                            v123 := GenS15{F0: 1, F1: int8(v6), F2: 1}
                                                                            _ = v123
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = v88
                                                        } else {
                                                            var v124 bool = false
                                                            if v124 {
                                                                v125 := GenS16{F0: 118, F1: (v6 ^ v5), F2: 2147483648}
                                                                _ = v125
                                                            }
                                                            v124 = (v124 || (218 == 0))
                                                            fmt.Println(fi79, v2)
                                                            fmt.Println(v2, fi35, v2)
                                                            _ = v124
                                                        }
                                                        v126 := GenS17{F0: 2147483648, F1: genFn0("")}
                                                        v127 := GenS4{F0: 4294967295, F1: fi79, F2: ("foo bar" + "foo bar")}
                                                        v128 := GenS18{F0: 0, F1: int(v6)}
                                                        var v129 int16 = (v34 ^ (v5 & (v6 << ((1) & 63))))
                                                        _ = v126
                                                        _ = v127
                                                        _ = v128
                                                        _ = v129
                                                    }
                                                }
                                                switch "" {
                                                case "foo bar":
                                                    fmt.Println(v6, v5, v6)
                                                case "":
                                                    if ("hello" >= "\n\t") {
                                                        var v130 string = "foo bar"
                                                        _ = v130
                                                    } else {
                                                        fmt.Println(fi35)
                                                        fmt.Println(v2, fi35, v5)
                                                    }
                                                case "\n\t":
                                                    switch fi79 {
                                                    case -9223372036854775808:
                                                        v131 := GenS19{F0: (fi79 ^ fi35)}
                                                        _ = v131
                                                    }
                                                }
                                            }
                                            if (!true) {
                                                var v132 int8 = 1
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v132
                                            } else {
                                                v133 := GenS20{F0: (fi79 * -2147483648), F1: "", F2: 123}
                                                v134 := GenS21{F0: "", F1: v5, F2: 42}
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v133
                                                _ = v134
                                            }
                                        } else {
                                            if ((255 == 161) && (true && true)) {
                                                v6 = 32767
                                                if true {
                                                    var v135 uint8 = 127
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v5 = v34
                                                    case "foo bar":
                                                        var v136 int16 = (v5 >> ((4294967295) & 63))
                                                        _ = v136
                                                    case "a":
                                                        v137 := GenS9{F0: (fi35 ^ (-(fi79))), F1: (fi79 << ((420) & 63)), F2: (v135 * v135)}
                                                        _ = v137
                                                    default:
                                                        var v138 uint32 = 4294967295
                                                        _ = v138
                                                    }
                                                    if (1e10 < 0.0) {
                                                        v135 += 189
                                                    } else {
                                                        if true {
                                                            fmt.Println(fi79, fi79)
                                                            if true {
                                                                fmt.Println(v135, v34, v5)
                                                                switch 9223372036854775807 {
                                                                case -1:
                                                                    v6 += v6
                                                                }
                                                                fmt.Println(fi79, v34, fi35)
                                                            } else {
                                                                var v139 uint8 = 60
                                                                _ = v139
                                                            }
                                                            var v140 float32 = float32(float64(((1e10 * 1.5) * ((0.1 - float32(0.1)) / 0.1))))
                                                            _ = v140
                                                        }
                                                        var v141 string = "\n\t"
                                                        _ = v141
                                                    }
                                                    _ = v135
                                                } else {
                                                    v6 = 1000
                                                    switch "a" {
                                                    case "foo bar":
                                                        v5 = v6
                                                    case "a":
                                                        v2 -= (v2 * v2)
                                                    }
                                                    if true {
                                                        var v142 int = fi79
                                                        if (!(((!(false && (!(!(!true))))) && (65535 <= 787)) && false)) {
                                                            v34 += (v6 / ((v6) | 1))
                                                            if (4294967295 < 1) {
                                                                switch 128 {
                                                                case -1:
                                                                    v143 := GenS22{F0: (42 <= 2), F1: "\n\t", F2: 4294967295}
                                                                    _ = v143
                                                                case 1:
                                                                    if false {
                                                                        switch fi79 {
                                                                        case -1:
                                                                            var v144 int64 = -9223372036854775808
                                                                            _ = v144
                                                                        default:
                                                                            fmt.Println(v2, v142)
                                                                        }
                                                                        v145 := GenS19{F0: fi35}
                                                                        var v146 uint = 42
                                                                        _ = v145
                                                                        _ = v146
                                                                    }
                                                                }
                                                            }
                                                            var v147 int32 = -602
                                                            _ = v147
                                                        } else {
                                                            if false {
                                                                v1[0] = false
                                                            }
                                                            var v148 bool = true
                                                            _ = v148
                                                        }
                                                        var v149 uint = 394
                                                        _ = v142
                                                        _ = v149
                                                    }
                                                }
                                                v6 = (v34 % (((v34 ^ -32768)) | 1))
                                            }
                                        }
                                        if (!(false && true)) {
                                            v6 += v6
                                        }
                                        var v150 int16 = (v5 / ((v34) | 1))
                                        _ = v150
                                    } else {
                                        if false {
                                            var v151 int64 = 839
                                            var v152 string = ("a" + ("" + "foo bar"))
                                            switch ("foo bar" + ((((("foo bar" + v152) + (("" + v152) + ("a" + v152))) + (v152 + "hello")) + v152) + "\n\t")) {
                                            case "a":
                                                var v153 float32 = 0.0
                                                _ = v153
                                            case "foo bar":
                                                v154 := GenS23{F0: -2147483648, F1: 1}
                                                _ = v154
                                            default:
                                                var v155 uint16 = 0
                                                _ = v155
                                            }
                                            _ = v151
                                            _ = v152
                                        }
                                    }
                                    if true {
                                        if (!false) {
                                            switch 32767 {
                                            case 42:
                                                v156 := GenS6{F0: 42, F1: 917, F2: (v2 - (v2 | (^(v2))))}
                                                if true {
                                                    var v157 string = "\n\t"
                                                    fmt.Println(fi35, fi79, fi79)
                                                    _ = v157
                                                } else {
                                                    v2 = (v2 + (^(v2)))
                                                    if false {
                                                        v158 := GenS24{F0: false, F1: 1}
                                                        switch ("a" + ((((((("a" + "a") + "") + ("\n\t" + ("hello" + "\n\t"))) + "a") + "hello") + "\n\t") + ((("hello" + ("a" + "\n\t")) + "a") + "foo bar"))) {
                                                        case "a":
                                                            switch "" {
                                                            case "a":
                                                                var v159 int32 = (v2 - v2)
                                                                _ = v159
                                                            case "\n\t":
                                                                var v160 uint16 = 32767
                                                                _ = v160
                                                            case "foo bar":
                                                                v161 := GenS21{F0: "a", F1: (v6 / ((v5) | 1)), F2: 145}
                                                                _ = v161
                                                            }
                                                        }
                                                        _ = v158
                                                    } else {
                                                        v4 = append(v4, 2147483647)
                                                    }
                                                }
                                                v6 = 1
                                                v162 := GenS19{F0: 1}
                                                v163 := GenS25{F0: 891, F1: (fi35 << ((255) & 63)), F2: 1}
                                                var v164 uint32 = 514
                                                _ = v156
                                                _ = v162
                                                _ = v163
                                                _ = v164
                                            default:
                                                if false {
                                                    v165 := GenS17{F0: 835, F1: 32767}
                                                    v2 = v2
                                                    var v166 int64 = -1
                                                    var v167 int = fi35
                                                    fmt.Println(v34, v167, v6)
                                                    v168 := v1[4294967295]
                                                    var v169 string = ""
                                                    delete(v1, 946)
                                                    var v170 int64 = (v166 - (v166 ^ int64(v167)))
                                                    if false {
                                                        var v171 bool = (v168 || (1 < 580))
                                                        _ = v171
                                                    } else {
                                                        v2 = (v2 & (v2 % (((v2 - (v2 << ((42) & 63)))) | 1)))
                                                        switch v169 {
                                                        case "\n\t":
                                                            var v172 uint8 = 2
                                                            _ = v172
                                                        default:
                                                            var v173 uint8 = 59
                                                            _ = v173
                                                        }
                                                        if v168 {
                                                            fmt.Println(v169, v2)
                                                            var v174 uint = 255
                                                            fmt.Println(v166, v6, fi35)
                                                            _ = v174
                                                        } else {
                                                            var v175 int = (fi79 / ((32767) | 1))
                                                            v3 = append(v3, int(v166))
                                                            _ = v175
                                                        }
                                                    }
                                                    if false {
                                                        _ = len(v3)
                                                        v176 := GenS26{F0: 42, F1: 4294967295, F2: -401}
                                                        _ = v176
                                                    } else {
                                                        if true {
                                                            var v177 int8 = -128
                                                            fmt.Println(v2, v170, v6)
                                                            _ = v177
                                                        }
                                                        v2 = (-(v2))
                                                        if (863 <= 42) {
                                                            v178 := GenS27{F0: 1, F1: v169}
                                                            fmt.Println(v6)
                                                            if ((v169 + (("\n\t" + v169) + v169)) > v169) {
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                v3 = append(v3, fi79)
                                                                if (65535 >= 42) {
                                                                    if v168 {
                                                                        if v168 {
                                                                            var v179 uint16 = 104
                                                                            _ = v179
                                                                        } else {
                                                                            switch v169 {
                                                                            case "\n\t":
                                                                                var v180 float64 = float64(float32(0.1))
                                                                                _ = v180
                                                                            default:
                                                                                _ = len(v3)
                                                                            }
                                                                        }
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (v167 - fi79)
                                                                        }
                                                                    }
                                                                    v2 -= (v2 * -443)
                                                                }
                                                            }
                                                            var v181 uint64 = 18446744073709551615
                                                            if v168 {
                                                                v182 := GenS11{F0: 42, F1: v34}
                                                                var v183 bool = (v168 && true)
                                                                _ = v182
                                                                _ = v183
                                                            } else {
                                                                var v184 float64 = 0.0
                                                                _ = v184
                                                            }
                                                            _ = v178
                                                            _ = v181
                                                        }
                                                    }
                                                    _ = v165
                                                    _ = v166
                                                    _ = v167
                                                    _ = v168
                                                    _ = v169
                                                    _ = v170
                                                }
                                                v2 = -1
                                                v2 += (v2 << ((0) & 63))
                                            }
                                            if (!(42 > genFn0("a"))) {
                                                _ = len(v1)
                                                v4 = append(v4, 0)
                                            } else {
                                                v185 := GenS13{F0: 42}
                                                switch fi79 {
                                                case 2147483647:
                                                    fmt.Println(fi35, v5)
                                                case -151:
                                                    switch (fi35 / (((^(fi79))) | 1)) {
                                                    case 32768:
                                                        var v186 int32 = (v2 ^ 2147483647)
                                                        _ = v186
                                                    default:
                                                        v5 = v34
                                                    }
                                                case 4294967295:
                                                    switch ("foo bar" + "hello") {
                                                    case "":
                                                        var v187 int = fi35
                                                        _ = v187
                                                    }
                                                default:
                                                    if (true && (false && false)) {
                                                        v188 := GenS28{F0: genFn0((("foo bar" + "foo bar") + ("" + "foo bar"))), F1: 773, F2: "a"}
                                                        _ = v188
                                                    } else {
                                                        v34 = v6
                                                        v189 := GenS18{F0: 788, F1: fi79}
                                                        _ = v189
                                                    }
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = 65535
                                                }
                                                _ = v185
                                            }
                                        } else {
                                            v190 := GenS29{F0: 65535}
                                            v191 := GenS30{F0: 1, F1: 1}
                                            _ = v190
                                            _ = v191
                                        }
                                    } else {
                                        v5 -= (v5 - (v34 - v6))
                                    }
                                }
                                var v192 int16 = v6
                                _ = v192
                            }
                            var v193 uint64 = 18446744073709551615
                            fmt.Println(v5, fi79, v5)
                            _ = fi79
                            _ = v193
                        }
                    } else {
                        var v194 uint16 = 255
                        v194 -= v194
                        _ = v194
                    }
                    v6 = (^(v34))
                    if false {
                        for fi195 := 0; fi195 < 19; fi195++ {
                            v2 += v2
                            switch ("foo bar" + "\n\t") {
                            case "hello":
                                _ = len(v1)
                            case "a":
                                fmt.Println(fi195)
                            }
                            _ = fi195
                        }
                        v6 += v5
                        switch 32768 {
                        case 246:
                            for fi196 := 0; fi196 < 17; fi196++ {
                                if false {
                                    var v197 int64 = -1
                                    v197 = v197
                                    _ = v197
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                switch fi196 {
                                case -453:
                                    if false {
                                        v6 -= (v34 % ((v34) | 1))
                                        v198 := GenS31{F0: 0, F1: 4294967295, F2: 2}
                                        fmt.Println(v34, fi196)
                                        _ = v198
                                    } else {
                                        if (((!(-525 != 9223372036854775807)) || (!(!(!false)))) || true) {
                                            switch fi35 {
                                            case 829:
                                                v3 = append(v3, (^(fi35)))
                                            case 4294967295:
                                                _ = len(v3)
                                            case 32767:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            if true {
                                                var v199 int = (fi35 ^ (fi196 + fi196))
                                                _ = v199
                                            }
                                            v6 -= -1
                                            v200 := GenS28{F0: 1, F1: (fi35 + -1), F2: ""}
                                            _ = v200
                                        }
                                        var v201 uint16 = genFn0("\n\t")
                                        var v202 int32 = 2147483647
                                        v202 -= (v202 % (((-(v202))) | 1))
                                        if (42 != 1) {
                                            _ = len(v4)
                                            var v203 float64 = 0.1
                                            v5 -= 0
                                            v6 -= 32767
                                            _ = v203
                                        }
                                        if (!true) {
                                            var v204 uint32 = 93
                                            v1[(v204 % ((v204) | 1))] = ((!(!true)) || (!false))
                                            switch "\n\t" {
                                            case "":
                                                if (!(1 != 131)) {
                                                    if true {
                                                        v4 = append(v4, 187)
                                                        if true {
                                                            switch 32768 {
                                                            case 42:
                                                                v205 := v1[(v204 ^ (^(v204)))]
                                                                v206 := GenS15{F0: 0, F1: 1, F2: 1}
                                                                _ = v205
                                                                _ = v206
                                                            case -1:
                                                                switch (^(fi196)) {
                                                                case -1:
                                                                    v6 += v5
                                                                case -795:
                                                                    v6 -= (v5 | (v5 % (((v6 << ((2) & 63))) | 1)))
                                                                    switch 32767 {
                                                                    case -2147483648:
                                                                        if true {
                                                                            v202 -= (v2 / ((v202) | 1))
                                                                            var v207 int64 = int64(v204)
                                                                            _ = v207
                                                                        }
                                                                    default:
                                                                        if (false && (-128 == -128)) {
                                                                            v208 := GenS32{F0: ("" + ""), F1: 2}
                                                                            _ = v208
                                                                        }
                                                                    }
                                                                    delete(v1, v204)
                                                                    v202 = -2147483648
                                                                case 32767:
                                                                    if (!true) {
                                                                        var v209 bool = true
                                                                        var v210 bool = v209
                                                                        var v211 float64 = (float64(0.0) * -2.25)
                                                                        _ = v209
                                                                        _ = v210
                                                                        _ = v211
                                                                    }
                                                                }
                                                            case -379:
                                                                _ = len(v3)
                                                            default:
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            }
                                                            if ((0 <= -1) && (false && false)) {
                                                                fmt.Println(fi196, fi196)
                                                                fmt.Println(v6, v6)
                                                                switch (fi196 % ((-158) | 1)) {
                                                                case 128:
                                                                    delete(v1, (v204 + (v204 & (v204 - v204))))
                                                                case -102:
                                                                    var v212 int8 = 1
                                                                    _ = v212
                                                                default:
                                                                    var v213 uint32 = v204
                                                                    _ = v213
                                                                }
                                                            } else {
                                                                v6 = -1
                                                                switch fi35 {
                                                                case 106:
                                                                    v214 := GenS33{F0: "hello"}
                                                                    _ = v214
                                                                case -773:
                                                                    var v215 uint32 = v204
                                                                    _ = v215
                                                                case 2147483647:
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                default:
                                                                    var v216 bool = (127 >= 64)
                                                                    _ = v216
                                                                }
                                                                if false {
                                                                    v5 = (v6 * 165)
                                                                    var v217 uint32 = 4294967295
                                                                    _ = v217
                                                                } else {
                                                                    if (42 < 1) {
                                                                        v2 = v2
                                                                    }
                                                                    var v218 uint64 = 4294967295
                                                                    if (2147483647 >= (v2 << ((18446744073709551615) & 63))) {
                                                                        v219 := GenS34{F0: 255, F1: 18446744073709551615, F2: 100}
                                                                        _ = v219
                                                                    } else {
                                                                        fmt.Println(v6, fi196, v204)
                                                                        v2 = -2147483648
                                                                    }
                                                                    _ = v218
                                                                }
                                                            }
                                                        }
                                                        _ = len(v1)
                                                    }
                                                } else {
                                                    v220 := GenS3{F0: fi196, F1: v204, F2: -492}
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v221 uint16 = v201
                                                    _ = v220
                                                    _ = v221
                                                }
                                            case "foo bar":
                                                v202 -= v202
                                            case "hello":
                                                switch fi35 {
                                                case -615:
                                                    var v222 uint64 = 18446744073709551615
                                                    var v223 int32 = v202
                                                    _ = v222
                                                    _ = v223
                                                case -815:
                                                    switch "foo bar" {
                                                    case "":
                                                        if (!true) {
                                                            switch "hello" {
                                                            case "":
                                                                switch fi196 {
                                                                case 1:
                                                                    v3 = append(v3, (fi35 | 32768))
                                                                }
                                                            case "foo bar":
                                                                var v224 int16 = (^(v6))
                                                                _ = v224
                                                            }
                                                            v34 -= v6
                                                            if (true || false) {
                                                                v6 = v6
                                                                fmt.Println(fi196)
                                                            } else {
                                                                var v225 int16 = (v6 ^ (v5 ^ 32767))
                                                                if (!("hello" != ("\n\t" + "hello"))) {
                                                                    if ((v201 + 954) != 255) {
                                                                        if true {
                                                                            switch "hello" {
                                                                            case "a":
                                                                                var v226 uint = 2
                                                                                v34 = -32768
                                                                                _ = v226
                                                                            }
                                                                            if len(v4) > 0 {
                                                                                v4[0] = v204
                                                                            }
                                                                            switch (fi196 | fi196) {
                                                                            case -198:
                                                                                var v227 float32 = 0.0
                                                                                _ = v227
                                                                            default:
                                                                                switch fi196 {
                                                                                case -1:
                                                                                    var v228 uint16 = 32767
                                                                                    _ = v228
                                                                                case 128:
                                                                                    if (v6 >= (v225 ^ 42)) {
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            var v229 bool = (true && (int64(v204) < 846))
                                                                                            _ = v229
                                                                                        case "":
                                                                                            v230 := v1[v204]
                                                                                            _ = v230
                                                                                        case "hello":
                                                                                            v2 = -2147483648
                                                                                        }
                                                                                        v2 = (v202 + (-(v2)))
                                                                                    } else {
                                                                                        fmt.Println(v201, v204)
                                                                                        v204 -= (v204 + uint32(v201))
                                                                                    }
                                                                                case 127:
                                                                                    v231 := GenS28{F0: (v201 & (v201 ^ v201)), F1: 32768, F2: "hello"}
                                                                                    _ = v231
                                                                                default:
                                                                                    v34 = v5
                                                                                    switch 32767 {
                                                                                    case -755:
                                                                                        v232 := GenS35{F0: v204, F1: 693, F2: genFn1((-(fi196)))}
                                                                                        _ = v232
                                                                                    case 2147483647:
                                                                                        switch ("foo bar" + ("a" + "a")) {
                                                                                        case "a":
                                                                                            v233 := GenS36{F0: "foo bar", F1: (!true), F2: (v2 % ((int32(v201)) | 1))}
                                                                                            var v234 uint64 = 2
                                                                                            var v235 uint32 = 0
                                                                                            _ = v233
                                                                                            _ = v234
                                                                                            _ = v235
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = len(v1)
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "":
                                                                            v236 := GenS37{F0: 100, F1: -2147483648, F2: ("hello" + "hello")}
                                                                            _ = v236
                                                                        case "a":
                                                                            var v237 int32 = v202
                                                                            _ = v237
                                                                        case "foo bar":
                                                                            var v238 uint16 = genFn0("\n\t")
                                                                            _ = v238
                                                                        default:
                                                                            v239 := GenS4{F0: 2, F1: fi35, F2: "hello"}
                                                                            _ = v239
                                                                        }
                                                                        switch ("\n\t" + "") {
                                                                        case "\n\t":
                                                                            var v240 bool = (127 <= 42)
                                                                            _ = v240
                                                                        case "hello":
                                                                            var v241 float64 = -2.25
                                                                            _ = v241
                                                                        }
                                                                        if false {
                                                                            if true {
                                                                                if false {
                                                                                    v5 += (v6 | (-(v225)))
                                                                                }
                                                                                v242 := v1[(v204 - 2)]
                                                                                _ = v242
                                                                            } else {
                                                                                v5 -= v5
                                                                                switch (fi35 | (fi196 >> ((341) & 63))) {
                                                                                case -792:
                                                                                    switch "a" {
                                                                                    case "":
                                                                                        switch "foo bar" {
                                                                                        case "":
                                                                                            var v243 string = "hello"
                                                                                            v3 = append(v3, fi196)
                                                                                            v6 += (v6 * v225)
                                                                                            v204 = uint32(v201)
                                                                                            var v244 uint64 = 383
                                                                                            _ = v243
                                                                                            _ = v244
                                                                                        case "foo bar":
                                                                                            var v245 int8 = 35
                                                                                            _ = v245
                                                                                        }
                                                                                    case "hello":
                                                                                        v246 := GenS38{F0: 4294967295, F1: -36}
                                                                                        _ = v246
                                                                                    case "\n\t":
                                                                                        switch (fi196 << ((4294967295) & 63)) {
                                                                                        case 32767:
                                                                                            v201 -= (v201 % ((484) | 1))
                                                                                        case 128:
                                                                                            if (true && (0.0 == 3.14159)) {
                                                                                                var v247 string = ""
                                                                                                fmt.Println(fi35)
                                                                                                v247 = "foo bar"
                                                                                                _ = v247
                                                                                            } else {
                                                                                                v248 := GenS39{F0: 28}
                                                                                                v2 = (v2 ^ v2)
                                                                                                _ = v248
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v249 uint = 42
                                                                                _ = v249
                                                                            }
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            v201 = 2
                                                                        } else {
                                                                            delete(v1, v204)
                                                                            var v250 int16 = v225
                                                                            _ = v250
                                                                        }
                                                                    }
                                                                    v251 := GenS4{F0: 4294967295, F1: (fi35 / ((fi35) | 1)), F2: "foo bar"}
                                                                    _ = v251
                                                                }
                                                                v252 := GenS18{F0: 42, F1: (fi196 - 9223372036854775807)}
                                                                _ = v225
                                                                _ = v252
                                                            }
                                                        }
                                                    case "foo bar":
                                                        v202 = (v2 + v202)
                                                    default:
                                                        v253 := GenS1{F0: "a"}
                                                        v254 := GenS40{F0: -2147483648, F1: (65535 <= 441), F2: int32(fi196)}
                                                        var v255 int32 = (v202 >> ((42) & 63))
                                                        _ = v253
                                                        _ = v254
                                                        _ = v255
                                                    }
                                                }
                                            }
                                            _ = v204
                                        }
                                        _ = v201
                                        _ = v202
                                    }
                                case 2147483647:
                                    v5 -= v34
                                default:
                                    v6 = (v6 * -79)
                                    var v256 uint64 = 794
                                    v257 := GenS3{F0: fi35, F1: 65535, F2: (^(v2))}
                                    _ = v256
                                    _ = v257
                                }
                                _ = fi196
                            }
                        case -696:
                            if (3.14159 == -2.25) {
                                for fi258 := 0; fi258 < 14; fi258++ {
                                    v259 := GenS3{F0: fi35, F1: 865, F2: (v2 % (((v2 - (v2 * 42))) | 1))}
                                    v260 := v1[593]
                                    _ = fi258
                                    _ = v259
                                    _ = v260
                                }
                                v34 -= v5
                                var v261 int8 = 1
                                _ = v261
                            } else {
                                switch ("foo bar" + "") {
                                case "hello":
                                    for fi262 := 0; fi262 < 3; fi262++ {
                                        v4 = append(v4, 2)
                                        _ = fi262
                                    }
                                case "\n\t":
                                    var v263 float64 = 0.0
                                    _ = v263
                                case "a":
                                    _ = len(v4)
                                default:
                                    delete(v1, 42)
                                }
                                for fi264 := 0; fi264 < 6; fi264++ {
                                    v5 += v6
                                    _ = fi264
                                }
                            }
                        case 877:
                            fmt.Println(fi35)
                        }
                    } else {
                        _ = len(v1)
                        fmt.Println(v5, v6, v34)
                        v2 += 2147483647
                    }
                }
            }
            for fi265 := 0; fi265 < 18; fi265++ {
                var v266 string = ("foo bar" + "")
                _ = fi265
                _ = v266
            }
            v267 := GenS39{F0: 42}
            _ = fi35
            _ = v267
        }
        _ = v34
    } else {
        var v268 int = int(v2)
        v2 += (v2 & (^(v2)))
        var v269 uint8 = 1
        _ = v268
        _ = v269
    }
    switch 42 {
    case 127:
        for fi270 := 0; fi270 < 20; fi270++ {
            switch "" {
            case "foo bar":
                v271 := GenS29{F0: 1}
                _ = v271
            case "hello":
                fmt.Println(v5)
            }
            fmt.Println(v2, v5)
            switch 128 {
            case 32767:
                var v272 uint16 = 2
                _ = v272
            case 0:
                fmt.Println(v2, v6)
            case 42:
                if (false || true) {
                    _ = len(v3)
                    var v273 float32 = (((3.14159 - 1.5) - (3.14159 + float32(0.0))) - (float32(1e10) + (1e10 / (1.5 / 1.0))))
                    v274 := GenS41{F0: 9223372036854775807}
                    _ = v273
                    _ = v274
                }
            default:
                var v275 uint16 = 32767
                _ = v275
            }
            _ = fi270
        }
    default:
        for fi276 := 0; fi276 < 16; fi276++ {
            v277 := GenS17{F0: 4294967295, F1: 97}
            var v278 string = ""
            _ = fi276
            _ = v277
            _ = v278
        }
    }
    fmt.Println(v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v279 int32 = (v2 / ((-1) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v279:", v279)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v279
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
