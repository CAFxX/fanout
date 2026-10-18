package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 uint; F2 string }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 uint32; F1 int16; F2 int }
type GenS8 struct { F0 int64; F1 int64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32; F1 bool; F2 uint8 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint32; F1 int8; F2 int32 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int8; F1 uint; F2 uint }
type GenS16 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS17 struct { F0 bool; F1 uint32 }
type GenS18 struct { F0 int8; F1 int16 }
type GenS19 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 uint; F1 int }
func genFn0(p0 int16, p1 int, p2 int64) int32 {
    var v6 int32 = -1
    p0 -= p0
    v7 := []int{(p1 & p1), p1, p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 float32 = 0.0
    var v2 uint = 818
    v3 := []int{9223372036854775807, 9223372036854775807}
    var v4 string = ("a" + "")
    var v5 int16 = -1
    v4 = ("a" + v4)
    v3 = append(v3, 9223372036854775807)
    for fi8 := 0; fi8 < 16; fi8++ {
        switch fi8 {
        case -2147483648:
            var v9 int = (fi8 >> ((18446744073709551615) & 63))
            _ = v9
        case 9223372036854775807:
            switch (v4 + ("\n\t" + "\n\t")) {
            case "":
                var v10 string = (v4 + "foo bar")
                _ = v10
            default:
                if (18446744073709551615 <= 469) {
                    fmt.Println(v4, v4)
                } else {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        switch v4 {
                        case "foo bar":
                            v5 -= (v5 | (v5 ^ (v5 - -890)))
                        default:
                            v5 = v5
                        }
                        var v12 int8 = 42
                        _ = fi11
                        _ = v12
                    }
                    v13 := GenS1{F0: 1}
                    _ = v13
                }
            }
        default:
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 uint32 = 4294967295
                v1 = (v1 * (((float32((0.1 / 1.0)) + 1e10) - v1) - (1.5 + 1.5)))
                _ = fi14
                _ = v15
            }
        }
        if true {
            if (1 >= 2) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v1)
            } else {
                v2 -= v2
            }
            fmt.Println(v2, v1)
            if len(v3) > 0 {
                v3[0] = fi8
            }
        } else {
            if false {
                v16 := GenS2{F0: 0, F1: (^(v2)), F2: v2}
                _ = v16
            }
        }
        v4 = v4
        _ = fi8
    }
    var v17 int16 = v5
    if true {
        var v18 uint32 = 42
        if (1e10 == (((v1 - float32(1.5)) / v1) / v1)) {
            var v19 int = 32767
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 uint64 = 782
                switch (fi20 ^ v19) {
                case 74:
                    v2 += v2
                case -788:
                    switch 2147483647 {
                    case 42:
                        var v22 float32 = (((v1 - v1) - float32((((3.14159 * 3.14159) * 0.0) * 3.14159))) / (0.0 * ((0.0 + (0.1 / v1)) + v1)))
                        _ = v22
                    case 4294967295:
                        if (108 >= 1) {
                            if true {
                                v19 -= 149
                            }
                            if false {
                                v17 = v5
                                if (42 <= -1) {
                                    v23 := GenS1{F0: 42}
                                    _ = v23
                                }
                                v5 = 1
                            }
                        } else {
                            for fi24 := 0; fi24 < 10; fi24++ {
                                v25 := GenS1{F0: 42}
                                _ = fi24
                                _ = v25
                            }
                            if false {
                                fmt.Println(v2, v2, v2)
                                switch v4 {
                                case "hello":
                                    if (!(true || false)) {
                                        v1 -= ((((v1 * 3.14159) * (((v1 - 0.1) + float32(float64(v1))) - v1)) * -2.25) + v1)
                                        if (!true) {
                                            v26 := GenS3{F0: v2}
                                            _ = v26
                                        } else {
                                            v27 := GenS3{F0: v2}
                                            fmt.Println(fi20, v5)
                                            v18 += (v18 & (v18 / ((v18) | 1)))
                                            _ = v27
                                        }
                                        if true {
                                            if true {
                                                v28 := GenS2{F0: 2147483647, F1: (v2 % ((v2) | 1)), F2: (v2 + (v2 + (v2 | v2)))}
                                                switch fi20 {
                                                case 205:
                                                    var v29 float64 = ((float64(0.0) - (0.1 - (3.14159 - (((float64((0.1 * v1)) + (-2.25 * (-2.25 * (0.0 + 0.0)))) + (0.1 - -2.25)) - 0.0)))) * ((((0.1 + (((-2.25 / -2.25) / 3.14159) / -2.25)) * (0.0 - (((3.14159 - (3.14159 * -2.25)) - -2.25) / 1.0))) - 0.0) * float64(0.1)))
                                                    _ = v29
                                                case -139:
                                                    if len(v3) > 0 {
                                                        v3[0] = v19
                                                    }
                                                }
                                                _ = v28
                                            }
                                        } else {
                                            if (!true) {
                                                if ((true && (!true)) && false) {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v30 int64 = -545
                                                    v17 -= int16(v21)
                                                    _ = v30
                                                }
                                            }
                                            v31 := GenS3{F0: v2}
                                            switch "\n\t" {
                                            case "foo bar":
                                                switch (fi20 + (^(v19))) {
                                                case 41:
                                                    v32 := GenS4{F0: 774, F1: v2, F2: v4}
                                                    _ = v32
                                                case -861:
                                                    if (1e10 < v1) {
                                                        v17 += -489
                                                        if true {
                                                            v4 = (v4 + ("\n\t" + (v4 + v4)))
                                                            for fi33 := 0; fi33 < 11; fi33++ {
                                                                var v34 bool = true
                                                                _ = fi33
                                                                _ = v34
                                                            }
                                                        } else {
                                                            var v35 bool = (65535 != (v18 ^ 65535))
                                                            if v35 {
                                                                for fi36 := 0; fi36 < 5; fi36++ {
                                                                    var v37 uint64 = (v21 >> (((v2 % ((v2) | 1))) & 63))
                                                                    var v38 uint64 = 2
                                                                    _ = fi36
                                                                    _ = v37
                                                                    _ = v38
                                                                }
                                                                for fi39 := 0; fi39 < 11; fi39++ {
                                                                    var v40 uint64 = 18446744073709551615
                                                                    var v41 int32 = 42
                                                                    _ = fi39
                                                                    _ = v40
                                                                    _ = v41
                                                                }
                                                            }
                                                            _ = v35
                                                        }
                                                        if ((v1 * (float32(1.5) * (v1 - v1))) != v1) {
                                                            for fi42 := 0; fi42 < 20; fi42++ {
                                                                if (2 >= 128) {
                                                                    var v43 uint = 65535
                                                                    v44 := GenS3{F0: (^(v2))}
                                                                    if len(v3) > 0 {
                                                                        v3[0] = v19
                                                                    }
                                                                    _ = v43
                                                                    _ = v44
                                                                } else {
                                                                    fmt.Println(fi20, v1)
                                                                }
                                                                v45 := GenS5{F0: "", F1: (v18 >> ((2) & 63))}
                                                                var v46 int32 = genFn0(340, 127, -263)
                                                                _ = fi42
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                            var v47 int16 = v17
                                                            for fi48 := 0; fi48 < 9; fi48++ {
                                                                if ((-1 > 9223372036854775807) || (!true)) {
                                                                    var v49 bool = false
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    switch v4 {
                                                                    case "\n\t":
                                                                        v50 := GenS6{F0: 42, F1: (v21 ^ v21)}
                                                                        _ = v50
                                                                    case "a":
                                                                        var v51 bool = v49
                                                                        _ = v51
                                                                    default:
                                                                        v5 += v47
                                                                    }
                                                                    _ = v49
                                                                }
                                                                _ = fi48
                                                            }
                                                            _ = v47
                                                        }
                                                    }
                                                case 42:
                                                    v17 -= v17
                                                default:
                                                    v52 := GenS2{F0: -696, F1: 18446744073709551615, F2: v2}
                                                    _ = v52
                                                }
                                            case "a":
                                                if (true && (!true)) {
                                                    fmt.Println(v19, v1)
                                                } else {
                                                    v21 = v21
                                                    switch ((((v4 + "hello") + ((("\n\t" + v4) + ("hello" + v4)) + (("\n\t" + ("" + v4)) + ("\n\t" + v4)))) + v4) + "a") {
                                                    case "":
                                                        v19 -= v19
                                                    case "hello":
                                                        for fi53 := 0; fi53 < 4; fi53++ {
                                                            if false {
                                                                v54 := GenS7{F0: 42, F1: v5, F2: 4294967295}
                                                                _ = v54
                                                            }
                                                            fmt.Println(v4, fi20)
                                                            switch (v4 + v4) {
                                                            case "\n\t":
                                                                fmt.Println(v18, v17)
                                                            case "hello":
                                                                switch v4 {
                                                                case "":
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                case "foo bar":
                                                                    fmt.Println(v5, v21)
                                                                case "hello":
                                                                    if false {
                                                                        if false {
                                                                            v55 := GenS1{F0: 849}
                                                                            if false {
                                                                                v56 := GenS3{F0: v2}
                                                                                v17 -= v17
                                                                                _ = v56
                                                                            }
                                                                            var v57 int64 = 42
                                                                            _ = v55
                                                                            _ = v57
                                                                        }
                                                                        if (0.1 < -2.25) {
                                                                            if true {
                                                                                if false {
                                                                                    switch 729 {
                                                                                    case 9223372036854775807:
                                                                                        var v58 int8 = 100
                                                                                        _ = v58
                                                                                    case 0:
                                                                                        v59 := GenS8{F0: -1, F1: 9223372036854775807}
                                                                                        switch (v4 + "foo bar") {
                                                                                        case "\n\t":
                                                                                            v4 = (v4 + (v4 + v4))
                                                                                        case "foo bar":
                                                                                            if false {
                                                                                                var v60 float32 = 1e10
                                                                                                _ = v60
                                                                                            } else {
                                                                                                switch fi20 {
                                                                                                case -2147483648:
                                                                                                    var v61 bool = false
                                                                                                    _ = v61
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v21 += (v21 + v21)
                                                                                        _ = v59
                                                                                    default:
                                                                                        v17 -= v5
                                                                                    }
                                                                                }
                                                                                v3 = append(v3, fi53)
                                                                            } else {
                                                                                var v62 int16 = 644
                                                                                fmt.Println(v18, v4)
                                                                                var v63 uint64 = (v21 % ((v21) | 1))
                                                                                _ = v62
                                                                                _ = v63
                                                                            }
                                                                            v5 += (-(v5))
                                                                            v64 := GenS9{F0: 1}
                                                                            var v65 uint = (v2 & 830)
                                                                            _ = v64
                                                                            _ = v65
                                                                        } else {
                                                                            v21 += (v21 - v21)
                                                                            if false {
                                                                                switch v4 {
                                                                                case "a":
                                                                                    if (!true) {
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = -9223372036854775808
                                                                                        }
                                                                                        var v66 int64 = 9223372036854775807
                                                                                        v18 += v18
                                                                                        _ = v66
                                                                                    }
                                                                                default:
                                                                                    v67 := GenS10{F0: 2147483647, F1: true, F2: 142}
                                                                                    _ = v67
                                                                                }
                                                                            } else {
                                                                                if (1 >= 127) {
                                                                                    switch (fi20 + -958) {
                                                                                    case 446:
                                                                                        var v68 uint16 = 473
                                                                                        _ = v68
                                                                                    case -48:
                                                                                        fmt.Println(v18, fi20, v18)
                                                                                    default:
                                                                                        v69 := GenS8{F0: 9223372036854775807, F1: 1}
                                                                                        _ = v69
                                                                                    }
                                                                                    if false {
                                                                                        v5 += (v5 | v17)
                                                                                    } else {
                                                                                        if (true || (!false)) {
                                                                                            if false {
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                                var v70 uint64 = (v21 | v21)
                                                                                                switch v19 {
                                                                                                case 4294967295:
                                                                                                    v71 := GenS11{F0: (^(v18))}
                                                                                                    _ = v71
                                                                                                }
                                                                                                _ = v70
                                                                                            } else {
                                                                                                switch "" {
                                                                                                case "\n\t":
                                                                                                    switch (fi20 - 1) {
                                                                                                    case 811:
                                                                                                        var v72 uint = v2
                                                                                                        _ = v72
                                                                                                    case 684:
                                                                                                        v73 := GenS11{F0: (v18 & 1)}
                                                                                                        v19 += fi20
                                                                                                        if (128 <= 255) {
                                                                                                            fmt.Println(v19)
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = v19
                                                                                                            }
                                                                                                            v74 := GenS7{F0: v18, F1: v5, F2: 902}
                                                                                                            fmt.Println(v21, fi53, fi53)
                                                                                                            _ = v74
                                                                                                        } else {
                                                                                                            v75 := GenS1{F0: 0}
                                                                                                            fmt.Println(v5, v5, v21)
                                                                                                            _ = v75
                                                                                                        }
                                                                                                        _ = v73
                                                                                                    default:
                                                                                                        var v76 int16 = (-(v17))
                                                                                                        _ = v76
                                                                                                    }
                                                                                                case "hello":
                                                                                                    var v77 int64 = int64(v18)
                                                                                                    _ = v77
                                                                                                default:
                                                                                                    var v78 int32 = genFn0((v17 / (((-(v17))) | 1)), fi20, -1)
                                                                                                    _ = v78
                                                                                                }
                                                                                                _ = len(v3)
                                                                                            }
                                                                                            var v79 uint32 = 744
                                                                                            _ = v79
                                                                                        } else {
                                                                                            var v80 float64 = (1.5 + float64(0.1))
                                                                                            _ = v80
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v18)
                                                                                }
                                                                                switch (v4 + v4) {
                                                                                case "foo bar":
                                                                                    fmt.Println(v4, v2)
                                                                                case "hello":
                                                                                    _ = len(v3)
                                                                                case "\n\t":
                                                                                    var v81 uint8 = 255
                                                                                    _ = v81
                                                                                default:
                                                                                    v21 -= v21
                                                                                }
                                                                                v82 := GenS12{F0: (v18 - v18)}
                                                                                _ = v82
                                                                            }
                                                                        }
                                                                        switch (v4 + v4) {
                                                                        case "foo bar":
                                                                            var v83 int64 = -1
                                                                            _ = v83
                                                                        }
                                                                    } else {
                                                                        switch (fi53 + 128) {
                                                                        case -739:
                                                                            var v84 string = v4
                                                                            _ = v84
                                                                        case 0:
                                                                            v85 := GenS3{F0: (^(v2))}
                                                                            _ = v85
                                                                        case 4294967295:
                                                                            fmt.Println(v21, v19, fi20)
                                                                        default:
                                                                            if (((1e10 * 1.5) / float64(v1)) == (-2.25 + (-2.25 * (3.14159 + 1e10)))) {
                                                                                v86 := GenS7{F0: 4294967295, F1: (v17 + 1), F2: fi53}
                                                                                _ = v86
                                                                            } else {
                                                                                var v87 int = fi20
                                                                                v4 = v4
                                                                                if (127 <= 128) {
                                                                                    v17 = v5
                                                                                    var v88 int64 = -9223372036854775808
                                                                                    _ = v88
                                                                                } else {
                                                                                    if (!(!(true || (true || (v4 <= ((v4 + v4) + (v4 + (v4 + "")))))))) {
                                                                                        v2 = v2
                                                                                        v18 = uint32(v2)
                                                                                        v21 = 758
                                                                                    }
                                                                                    v19 = 632
                                                                                }
                                                                                var v89 bool = true
                                                                                var v90 int = 98
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                _ = v87
                                                                                _ = v89
                                                                                _ = v90
                                                                            }
                                                                            var v91 int64 = -9223372036854775808
                                                                            _ = v91
                                                                        }
                                                                        if (v18 > v18) {
                                                                            v19 = 128
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            fmt.Println(v19, v4, v19)
                                                                        } else {
                                                                            var v92 uint64 = (v21 & v21)
                                                                            _ = v92
                                                                        }
                                                                    }
                                                                default:
                                                                    v19 = (v19 & fi53)
                                                                }
                                                                fmt.Println(v5)
                                                                _ = len(v3)
                                                            default:
                                                                switch "a" {
                                                                case "\n\t":
                                                                    _ = len(v3)
                                                                }
                                                                v18 = v18
                                                                switch (v4 + v4) {
                                                                case "a":
                                                                    v3 = append(v3, (v19 >> (((v2 * (v2 * 213))) & 63)))
                                                                }
                                                            }
                                                            _ = fi53
                                                        }
                                                    }
                                                    switch "foo bar" {
                                                    case "a":
                                                        fmt.Println(v17, v4, v17)
                                                    case "":
                                                        var v93 uint16 = 1
                                                        _ = v93
                                                    default:
                                                        v94 := GenS13{F0: 0, F1: -52, F2: -1}
                                                        _ = v94
                                                    }
                                                }
                                            default:
                                                switch -333 {
                                                case -2147483648:
                                                    var v95 float32 = 1.5
                                                    _ = v95
                                                case -9223372036854775808:
                                                    fmt.Println(v4)
                                                default:
                                                    var v96 int8 = 34
                                                    _ = v96
                                                }
                                            }
                                            _ = v31
                                        }
                                    }
                                case "foo bar":
                                    var v97 float64 = -2.25
                                    _ = v97
                                case "a":
                                    var v98 int16 = 1000
                                    _ = v98
                                }
                            } else {
                                v99 := GenS14{F0: 2}
                                for fi100 := 0; fi100 < 10; fi100++ {
                                    v18 = v18
                                    _ = fi100
                                }
                                _ = v99
                            }
                            v5 += (v17 % (((v17 | 1)) | 1))
                        }
                    case 32768:
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                    default:
                        var v101 uint64 = (^(v21))
                        _ = v101
                    }
                case -567:
                    var v102 uint16 = 2
                    _ = v102
                default:
                    _ = len(v3)
                }
                _ = fi20
                _ = v21
            }
            if (42 != 808) {
                for fi103 := 0; fi103 < 20; fi103++ {
                    switch v4 {
                    case "foo bar":
                        if (976 >= 42) {
                            var v104 string = (v4 + v4)
                            _ = v104
                        } else {
                            if false {
                                fmt.Println(v17, v5, fi103)
                                if false {
                                    var v105 uint = (v2 & v2)
                                    switch (fi103 << ((65535) & 63)) {
                                    case -9223372036854775808:
                                        switch (v4 + v4) {
                                        case "a":
                                            if false {
                                                v105 = 0
                                                for fi106 := 0; fi106 < 5; fi106++ {
                                                    fmt.Println(v17, v2)
                                                    _ = fi106
                                                }
                                            } else {
                                                v107 := GenS10{F0: 42, F1: (42 != v18), F2: 53}
                                                for fi108 := 0; fi108 < 1; fi108++ {
                                                    if ((-2.25 == (float64(float32((0.1 + -2.25))) / 1.0)) || (v17 > (^(v17)))) {
                                                        v4 = v4
                                                        fmt.Println(fi108)
                                                        if len(v3) > 0 {
                                                            v3[0] = (fi108 >> ((481) & 63))
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    } else {
                                                        v4 = (v4 + v4)
                                                        v19 -= 604
                                                        if false {
                                                            v1 = (v1 + (float32(0.1) / (v1 + 3.14159)))
                                                        }
                                                    }
                                                    var v109 float64 = (-2.25 / 0.1)
                                                    v5 -= -375
                                                    fmt.Println(fi103, v4)
                                                    _ = fi108
                                                    _ = v109
                                                }
                                                _ = v107
                                            }
                                        }
                                    case -280:
                                        if (2 > 135) {
                                            var v110 float64 = 0.0
                                            v4 = "hello"
                                            _ = v110
                                        } else {
                                            v105 = v2
                                            v111 := GenS1{F0: 1}
                                            var v112 int32 = -2147483648
                                            _ = v111
                                            _ = v112
                                        }
                                    }
                                    _ = v105
                                }
                                switch (fi103 + (v19 % ((32767) | 1))) {
                                case 1:
                                    if (!(-32768 == (v5 - v5))) {
                                        for fi113 := 0; fi113 < 11; fi113++ {
                                            if (-26 >= 0) {
                                                v2 += (v2 ^ v2)
                                                v1 -= 0.1
                                            } else {
                                                v19 -= (fi103 / ((fi103) | 1))
                                                v4 = v4
                                                v19 = (v19 - fi113)
                                                v17 += v5
                                            }
                                            if (!true) {
                                                var v114 float64 = 3.14159
                                                v115 := GenS4{F0: 1, F1: (v2 % ((v2) | 1)), F2: v4}
                                                _ = v114
                                                _ = v115
                                            } else {
                                                switch (fi103 << (((v2 >> ((v2) & 63))) & 63)) {
                                                case -961:
                                                    _ = len(v3)
                                                default:
                                                    v3 = append(v3, -357)
                                                }
                                                switch (fi113 >> ((0) & 63)) {
                                                case 2147483647:
                                                    v116 := GenS8{F0: -9223372036854775808, F1: 9223372036854775807}
                                                    _ = v116
                                                default:
                                                    v117 := GenS10{F0: 2147483647, F1: (0 == 32767), F2: 250}
                                                    _ = v117
                                                }
                                            }
                                            fmt.Println(v18)
                                            if true {
                                                v118 := GenS15{F0: 1, F1: (v2 ^ 138), F2: 4294967295}
                                                v5 -= -32768
                                                var v119 int16 = v5
                                                _ = v118
                                                _ = v119
                                            }
                                            v120 := GenS16{F0: 42, F1: 2, F2: (v17 & -1)}
                                            _ = fi113
                                            _ = v120
                                        }
                                        if true {
                                            v1 = float32(float64(v1))
                                        }
                                    } else {
                                        fmt.Println(v1)
                                    }
                                case 128:
                                    switch (v4 + "") {
                                    case "foo bar":
                                        var v121 int16 = -32768
                                        _ = v121
                                    case "":
                                        v122 := GenS5{F0: v4, F1: (v18 * 4294967295)}
                                        _ = v122
                                    default:
                                        fmt.Println(v4, v5)
                                    }
                                default:
                                    v17 -= (^(v5))
                                }
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v123 float64 = (0.1 / 1e10)
                                var v124 uint8 = 128
                                _ = v123
                                _ = v124
                            }
                            if (!false) {
                                var v125 float64 = 0.0
                                _ = v125
                            } else {
                                v126 := GenS17{F0: (false || ((true && (((((((-(v17)) <= int16(v19)) || (!false)) || false) || ((true || true) || (v17 <= (-(v5))))) || true) || false)) && true)), F1: (v18 & (v18 & (v18 * 2147483647)))}
                                v19 += (fi103 & (v19 * (v19 % (((v19 & 32768)) | 1))))
                                v1 = 0.1
                                _ = v126
                            }
                        }
                    case "a":
                        var v127 int16 = 32767
                        _ = v127
                    }
                    var v128 uint32 = (v18 - v18)
                    v19 = (fi103 | (^(v19)))
                    _ = fi103
                    _ = v128
                }
            } else {
                v3 = append(v3, 753)
                var v129 uint = uint(v18)
                v130 := GenS16{F0: v18, F1: 1, F2: v5}
                _ = v129
                _ = v130
            }
            _ = v19
        } else {
            v17 = 1000
            var v131 float32 = 1e10
            _ = v131
        }
        if (false && (true || false)) {
            var v132 uint32 = (v18 << (((v2 << (((v2 | v2)) & 63))) & 63))
            var v133 int8 = -128
            var v134 int64 = -9223372036854775808
            _ = v132
            _ = v133
            _ = v134
        }
        _ = v18
    } else {
        var v135 uint16 = 42
        var v136 float64 = ((0.0 - (1e10 - 0.1)) + 0.0)
        _ = v135
        _ = v136
    }
    _ = len(v3)
    if (-2.25 == v1) {
        v1 = v1
        if true {
            if (4294967295 < 32768) {
                var v137 float64 = 1.5
                _ = v137
            }
            for fi138 := 0; fi138 < 15; fi138++ {
                v4 = (v4 + ((v4 + "a") + (((v4 + "hello") + ("\n\t" + ((("\n\t" + v4) + v4) + "foo bar"))) + "\n\t")))
                var v139 string = v4
                _ = fi138
                _ = v139
            }
            if false {
                v140 := GenS5{F0: (("\n\t" + ("foo bar" + (v4 + v4))) + (v4 + (v4 + ((v4 + "hello") + v4)))), F1: 0}
                v2 += (v2 / ((v2) | 1))
                if false {
                    for fi141 := 0; fi141 < 15; fi141++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi141
                    }
                }
                _ = v140
            } else {
                if (407 < 0) {
                    var v142 bool = (false || (!(65535 > (^(v2)))))
                    switch -361 {
                    case -899:
                        fmt.Println(v2)
                    case 32768:
                        v142 = v142
                    case 4294967295:
                        v17 = 42
                    default:
                        switch 32768 {
                        case -561:
                            for fi143 := 0; fi143 < 10; fi143++ {
                                for fi144 := 0; fi144 < 17; fi144++ {
                                    var v145 float64 = (3.14159 / 0.1)
                                    _ = fi144
                                    _ = v145
                                }
                                if (v142 || (-128 < 34)) {
                                    for fi146 := 0; fi146 < 4; fi146++ {
                                        if v142 {
                                            v147 := GenS2{F0: -304, F1: v2, F2: v2}
                                            var v148 bool = v142
                                            _ = v147
                                            _ = v148
                                        }
                                        _ = fi146
                                    }
                                } else {
                                    switch v4 {
                                    case "a":
                                        for fi149 := 0; fi149 < 13; fi149++ {
                                            if len(v3) > 0 {
                                                v3[0] = (fi143 >> ((v2) & 63))
                                            }
                                            _ = fi149
                                        }
                                    case "hello":
                                        v150 := GenS18{F0: 1, F1: v17}
                                        _ = v150
                                    default:
                                        var v151 string = v4
                                        _ = v151
                                    }
                                    v2 = v2
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi143
                            }
                        case 469:
                            for fi152 := 0; fi152 < 1; fi152++ {
                                v3 = append(v3, fi152)
                                _ = fi152
                            }
                        case 128:
                            var v153 uint16 = 2
                            _ = v153
                        }
                    }
                    v154 := GenS13{F0: 2147483648, F1: -128, F2: int32(v17)}
                    _ = v142
                    _ = v154
                } else {
                    fmt.Println(v4)
                    for fi155 := 0; fi155 < 11; fi155++ {
                        var v156 string = v4
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (int8(v2) >= 0) {
                            switch v156 {
                            case "hello":
                                fmt.Println(v156, v2)
                            default:
                                var v157 bool = true
                                _ = v157
                            }
                            v4 = v156
                            v5 += (-(v17))
                        } else {
                            for fi158 := 0; fi158 < 19; fi158++ {
                                var v159 int64 = 786
                                if (4294967295 < 642) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v159 += (v159 % (((v159 / (((v159 * 9223372036854775807)) | 1))) | 1))
                                } else {
                                    switch (v156 + (v4 + "\n\t")) {
                                    case "a":
                                        v2 = (v2 * v2)
                                    case "":
                                        var v160 int64 = v159
                                        _ = v160
                                    case "\n\t":
                                        var v161 uint32 = 4294967295
                                        _ = v161
                                    }
                                    if (796 < 42) {
                                        switch (fi155 >> ((255) & 63)) {
                                        case 127:
                                            fmt.Println(fi155)
                                        default:
                                            v162 := GenS7{F0: 1, F1: 797, F2: (fi155 >> (((v2 / ((v2) | 1))) & 63))}
                                            _ = v162
                                        }
                                    }
                                }
                                v163 := GenS19{F0: 255, F1: 0, F2: 1}
                                _ = fi158
                                _ = v159
                                _ = v163
                            }
                            var v164 int16 = 1000
                            _ = v164
                        }
                        _ = fi155
                        _ = v156
                    }
                }
                if true {
                    switch v4 {
                    case "hello":
                        if true {
                            fmt.Println(v2, v17, v17)
                        }
                    case "a":
                        v165 := GenS20{F0: 4294967295}
                        _ = v165
                    }
                    fmt.Println(v5)
                    fmt.Println(v2, v4)
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if false {
                        fmt.Println(v5, v5, v17)
                        var v166 float64 = 3.14159
                        v166 += 1e10
                        _ = v166
                    } else {
                        var v167 float32 = v1
                        _ = v167
                    }
                }
            }
        } else {
            for fi168 := 0; fi168 < 2; fi168++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch fi168 {
                case 2147483647:
                    switch v4 {
                    case "hello":
                        v4 = v4
                    case "a":
                        v5 = (v17 - int16(fi168))
                    case "":
                        switch fi168 {
                        case 179:
                            if false {
                                fmt.Println(v5)
                                if ("" == "hello") {
                                    v17 += (v17 + (v5 << ((v2) & 63)))
                                    v169 := GenS13{F0: 2147483647, F1: 100, F2: genFn0(v5, int(v2), 334)}
                                    _ = v169
                                } else {
                                    if (!((true && false) && false)) {
                                        v170 := GenS3{F0: (^(v2))}
                                        var v171 int64 = -9223372036854775808
                                        _ = v170
                                        _ = v171
                                    }
                                }
                            } else {
                                switch 127 {
                                case -257:
                                    if false {
                                        _ = len(v3)
                                        v17 = (v17 ^ (v17 + (^(v17))))
                                    } else {
                                        v172 := GenS6{F0: 464, F1: 1}
                                        _ = len(v3)
                                        fmt.Println(fi168, fi168)
                                        _ = v172
                                    }
                                default:
                                    for fi173 := 0; fi173 < 20; fi173++ {
                                        v174 := GenS13{F0: 1, F1: 75, F2: 2147483647}
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v175 := GenS21{F0: 0, F1: (-(fi173))}
                                        _ = fi173
                                        _ = v174
                                        _ = v175
                                    }
                                }
                            }
                        default:
                            v5 -= (v5 % (((v17 / (((v5 / ((v5) | 1))) | 1))) | 1))
                        }
                    default:
                        var v176 int = (-(fi168))
                        _ = v176
                    }
                case 9223372036854775807:
                    var v177 uint32 = 2
                    _ = v177
                default:
                    switch v4 {
                    case "":
                        if len(v3) > 0 {
                            v3[0] = (fi168 << (((v2 % (((v2 ^ (v2 >> ((42) & 63)))) | 1))) & 63))
                        }
                    default:
                        v5 = (v5 >> (((v2 - v2)) & 63))
                    }
                }
                v5 = (v5 % ((1) | 1))
                _ = fi168
            }
            for fi178 := 0; fi178 < 15; fi178++ {
                v5 += (v5 >> ((2) & 63))
                _ = len(v3)
                _ = fi178
            }
        }
        _ = len(v3)
    }
    for fi179 := 0; fi179 < 12; fi179++ {
        _ = len(v3)
        v3 = append(v3, fi179)
        _ = fi179
    }
    if (3.14159 <= (1.5 + v1)) {
        var v180 float32 = (1.5 * (0.1 / (((v1 + ((v1 / float32(1e10)) / (v1 - (1.5 - v1)))) + -2.25) / float32(((-2.25 * 3.14159) - (0.0 * (1.5 / 3.14159)))))))
        fmt.Println(v180, v17)
        v181 := GenS17{F0: true, F1: 399}
        _ = v180
        _ = v181
    }
    fmt.Println(v5, v1)
    v3 = append(v3, 128)
    var v182 int64 = 9223372036854775807
    fmt.Println(v4, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v182:", v182)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v182
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
