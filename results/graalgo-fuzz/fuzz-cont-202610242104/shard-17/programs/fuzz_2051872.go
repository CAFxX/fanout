package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint; F2 uint8 }
type GenS4 struct { F0 int64; F1 int64; F2 int8 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS7 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 int64; F1 uint64 }
type GenS10 struct { F0 bool; F1 uint64 }
type GenS11 struct { F0 uint32; F1 string }
type GenS12 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS13 struct { F0 int; F1 uint16; F2 uint64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 string; F1 int32 }
func genFn0(p0 int64, p1 bool) string {
    fmt.Println(p1, p0)
    v4 := GenS1{F0: 4294967295, F1: 0, F2: 0}
    _ = p0
    _ = p1
    _ = v4
    return "\n\t"
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]bool)
    v3 := []uint32{0}
    var v5 uint64 = 0
    v5 = 0
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int64 = (v1 / ((v1) | 1))
            _ = fi7
            _ = v8
        }
        v5 = (^(v5))
        _ = fi6
    }
    if true {
        v5 = (v5 * v5)
        if (42 >= v5) {
            v5 = (v5 % (((^(v5))) | 1))
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            switch "foo bar" {
            case "foo bar":
                for fi10 := 0; fi10 < 15; fi10++ {
                    switch (-(fi9)) {
                    case 1:
                        var v11 uint8 = 42
                        _ = v11
                    case -9223372036854775808:
                        var v12 bool = (-404 < -2147483648)
                        _ = v12
                    }
                    v1 = v1
                    if true {
                        v5 += 42
                        var v13 int32 = 363
                        _ = len(v3)
                        _ = v13
                    }
                    _ = fi10
                }
            default:
                fmt.Println(v1)
            }
            _ = fi9
        }
    } else {
        v14 := GenS1{F0: 2147483647, F1: 255, F2: (^(v5))}
        _ = v14
    }
    var v15 float64 = 1.5
    if (!false) {
        _ = len(v2)
        v1 += v1
    }
    if false {
        v16 := GenS1{F0: -313, F1: 18446744073709551615, F2: 0}
        var v17 uint64 = (v5 & v5)
        if (!((!false) && (1 == (v1 + (v1 % ((-957) | 1)))))) {
            var v18 float64 = v15
            v5 = v5
            var v19 float64 = v15
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 5; fi20++ {
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 bool = false
                    if (32767 != 367) {
                        v3 = append(v3, 2147483647)
                        v17 += v17
                    } else {
                        if v22 {
                            fmt.Println(fi21, v5)
                            switch ("hello" + "\n\t") {
                            case "foo bar":
                                if (127 < 60) {
                                    var v23 float32 = float32((v15 / (v15 - v15)))
                                    v17 = 981
                                    v24 := GenS2{F0: 12}
                                    _ = v23
                                    _ = v24
                                } else {
                                    v5 = (v17 + v5)
                                }
                            case "":
                                var v25 uint32 = 891
                                _ = v25
                            case "a":
                                var v26 float32 = float32(-2.25)
                                _ = v26
                            default:
                                var v27 uint64 = 1
                                _ = v27
                            }
                        } else {
                            fmt.Println(fi20)
                            delete(v2, fi20)
                        }
                    }
                    fmt.Println(v17, fi20)
                    v28 := GenS3{F0: v22, F1: 1, F2: uint8(v17)}
                    v5 = (v17 + 0)
                    _ = fi21
                    _ = v22
                    _ = v28
                }
                v17 -= (v5 + v17)
                _ = fi20
            }
            var v29 float64 = v15
            _ = v29
        }
        _ = v16
        _ = v17
    } else {
        var v30 int64 = v1
        if false {
            switch 4294967295 {
            case 128:
                _ = len(v3)
            default:
                for fi31 := 0; fi31 < 2; fi31++ {
                    v3 = append(v3, 2147483647)
                    v2[(fi31 % ((-247) | 1))] = false
                    switch (fi31 >> ((765) & 63)) {
                    case 978:
                        v15 += v15
                    case -9223372036854775808:
                        for fi32 := 0; fi32 < 9; fi32++ {
                            fmt.Println(v15, fi31, v1)
                            fmt.Println(v1)
                            v5 -= v5
                            _ = fi32
                        }
                    default:
                        v33 := GenS1{F0: (^(fi31)), F1: 2, F2: (v5 ^ v5)}
                        _ = v33
                    }
                    _ = fi31
                }
            }
            fmt.Println(v1, v1, v15)
        } else {
            v3 = append(v3, 65535)
            var v34 int32 = int32(v5)
            for fi35 := 0; fi35 < 15; fi35++ {
                if (true || true) {
                    for fi36 := 0; fi36 < 19; fi36++ {
                        switch fi36 {
                        case -1:
                            v5 += (v5 / ((42) | 1))
                        case 32767:
                            if false {
                                var v37 int = fi36
                                var v38 uint16 = 997
                                _ = v37
                                _ = v38
                            }
                        case 9223372036854775807:
                            v34 = 0
                        default:
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                        }
                        v30 -= v1
                        if false {
                            switch (fi36 ^ fi35) {
                            case 687:
                                var v39 float64 = (v15 - ((1.5 / (float64(-2.25) + v15)) - -2.25))
                                fmt.Println(v30, v5, fi36)
                                fmt.Println(fi35)
                                _ = v39
                            case -364:
                                switch fi36 {
                                case -9223372036854775808:
                                    v5 = (^(v5))
                                case -140:
                                    fmt.Println(fi36, v1)
                                    v3 = append(v3, 494)
                                    if (true || (!false)) {
                                        var v40 uint = 2
                                        var v41 uint32 = 0
                                        _ = v40
                                        _ = v41
                                    } else {
                                        var v42 uint = 0
                                        v43 := v2[(fi36 - 128)]
                                        var v44 uint32 = 2147483647
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                case -323:
                                    v45 := GenS4{F0: v1, F1: v1, F2: -1}
                                    _ = v45
                                default:
                                    v5 = (v5 ^ 4294967295)
                                }
                            case 117:
                                fmt.Println(fi36)
                            }
                            if true {
                                v34 = (v34 / (((-(v34))) | 1))
                            }
                            if len(v3) > 0 {
                                v3[0] = 42
                            }
                        }
                        _ = fi36
                    }
                    switch (fi35 >> ((42) & 63)) {
                    case 128:
                        if ((-(v34)) > v34) {
                            switch ("" + ("hello" + ("a" + "a"))) {
                            case "":
                                v15 = v15
                            case "foo bar":
                                v5 = (v5 * (v5 / ((v5) | 1)))
                            case "a":
                                var v46 int16 = -152
                                _ = v46
                            }
                            var v47 float32 = 1.5
                            _ = v47
                        } else {
                            fmt.Println(v30, v15)
                            var v48 uint16 = 1
                            fmt.Println(v30, v34, fi35)
                            _ = v48
                        }
                    case 0:
                        v15 = v15
                    default:
                        if (32767 != 920) {
                            for fi49 := 0; fi49 < 17; fi49++ {
                                switch 127 {
                                case 32767:
                                    v34 += 290
                                    v50 := GenS5{F0: 65535, F1: 42}
                                    _ = v50
                                case 128:
                                    _ = len(v3)
                                case -2147483648:
                                    if false {
                                        v34 = (v34 >> ((0) & 63))
                                        v51 := GenS4{F0: -9223372036854775808, F1: v1, F2: -128}
                                        var v52 int32 = (v34 + -1)
                                        if ((("" + (((("\n\t" + "\n\t") + "") + ("" + "a")) + "hello")) + (("\n\t" + "") + ((("hello" + ("\n\t" + ("\n\t" + "a"))) + "foo bar") + ("a" + "hello")))) > ((("foo bar" + ("\n\t" + "")) + (("" + "\n\t") + ("a" + (("" + ("foo bar" + ("\n\t" + "hello"))) + ("foo bar" + ("a" + (("foo bar" + ("foo bar" + "\n\t")) + ""))))))) + (("" + "a") + (("\n\t" + "hello") + "hello")))) {
                                            fmt.Println(v15, v5)
                                            v53 := GenS5{F0: 212, F1: v1}
                                            _ = v53
                                        }
                                        _ = len(v3)
                                        var v54 int8 = 99
                                        switch "\n\t" {
                                        case "foo bar":
                                            switch fi35 {
                                            case -2147483648:
                                                switch -286 {
                                                case 9223372036854775807:
                                                    v55 := GenS6{F0: 255, F1: 1, F2: 1}
                                                    _ = v55
                                                default:
                                                    v2[42] = true
                                                }
                                            case 4294967295:
                                                delete(v2, (fi49 % ((fi49) | 1)))
                                            case 32768:
                                                if true {
                                                    if false {
                                                        switch (("foo bar" + ("foo bar" + ("a" + "a"))) + (("" + "foo bar") + ("" + ("\n\t" + "a")))) {
                                                        case "":
                                                            if false {
                                                                fmt.Println(fi49, fi35)
                                                                v54 = v54
                                                            } else {
                                                                v56 := GenS7{F0: 32767, F1: 495, F2: (v34 >> ((255) & 63))}
                                                                _ = v56
                                                            }
                                                            if true {
                                                                if true {
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        _ = len(v3)
                                                                    case "a":
                                                                        fmt.Println(v30)
                                                                    case "":
                                                                        v30 = (-(v1))
                                                                    }
                                                                }
                                                                if ((float32(v15) / (3.14159 * (3.14159 - 0.1))) == 0.0) {
                                                                    var v57 uint16 = 788
                                                                    _ = v57
                                                                }
                                                                v58 := GenS3{F0: (4294967295 >= 0), F1: 254, F2: 194}
                                                                _ = v58
                                                            } else {
                                                                switch fi49 {
                                                                case -1:
                                                                    var v59 int64 = (v1 >> ((374) & 63))
                                                                    _ = v59
                                                                case 1:
                                                                    switch fi49 {
                                                                    case -2147483648:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 806
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            switch fi49 {
                                                            case 128:
                                                                var v60 uint8 = 42
                                                                _ = v60
                                                            case 127:
                                                                v5 += (v5 | v5)
                                                            case -166:
                                                                switch fi35 {
                                                                case 1:
                                                                    v15 -= v15
                                                                case -699:
                                                                    v61 := GenS7{F0: 65535, F1: 811, F2: v34}
                                                                    _ = v61
                                                                case 0:
                                                                    delete(v2, 4294967295)
                                                                }
                                                            default:
                                                                if true {
                                                                    v62 := GenS3{F0: (float32(float64(0.0)) > 0.0), F1: 2, F2: 13}
                                                                    fmt.Println(v52, v34, fi35)
                                                                    var v63 string = "hello"
                                                                    _ = v62
                                                                    _ = v63
                                                                } else {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                            }
                                                        case "foo bar":
                                                            var v64 int64 = -1
                                                            _ = v64
                                                        case "\n\t":
                                                            var v65 int32 = int32(v5)
                                                            _ = v65
                                                        }
                                                        fmt.Println(v34, v34, v52)
                                                        v30 = v30
                                                    }
                                                    switch 301 {
                                                    case 4294967295:
                                                        var v66 int8 = v54
                                                        _ = v66
                                                    case -9223372036854775808:
                                                        v67 := GenS3{F0: true, F1: 18446744073709551615, F2: 0}
                                                        _ = v67
                                                    }
                                                    if false {
                                                        v15 += v15
                                                        fmt.Println(fi49)
                                                        v52 += (v52 ^ (v52 >> ((0) & 63)))
                                                        if true {
                                                            v68 := GenS6{F0: 255, F1: 65535, F2: 1}
                                                            var v69 int8 = (v54 & v54)
                                                            _ = v68
                                                            _ = v69
                                                        } else {
                                                            if (!(18446744073709551615 <= 42)) {
                                                                v30 = (-(v30))
                                                                switch (fi35 * 9223372036854775807) {
                                                                case 0:
                                                                    fmt.Println(v30)
                                                                case -9223372036854775808:
                                                                    switch 9223372036854775807 {
                                                                    case 0:
                                                                        var v70 int32 = 631
                                                                        _ = v70
                                                                    case -382:
                                                                        v2[fi49] = true
                                                                    case 128:
                                                                        var v71 int16 = 1
                                                                        _ = v71
                                                                    default:
                                                                        var v72 float64 = 1e10
                                                                        _ = v72
                                                                    }
                                                                case -916:
                                                                    v30 = (-(v1))
                                                                }
                                                                v73 := v2[(fi49 % ((fi49) | 1))]
                                                                _ = v73
                                                            }
                                                            switch (("foo bar" + ("a" + "\n\t")) + "foo bar") {
                                                            case "foo bar":
                                                                v54 += (v54 / (((v54 ^ v54)) | 1))
                                                            case "":
                                                                if true {
                                                                    v74 := GenS5{F0: 479, F1: v30}
                                                                    _ = v74
                                                                } else {
                                                                    v30 -= v1
                                                                    fmt.Println(v1, v34)
                                                                    if false {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        switch ("a" + "\n\t") {
                                                                        case "\n\t":
                                                                            v5 -= 42
                                                                        default:
                                                                            v75 := GenS2{F0: 517}
                                                                            _ = v75
                                                                        }
                                                                        v76 := GenS8{F0: v5, F1: -599}
                                                                        v30 = 235
                                                                        v54 -= int8(fi35)
                                                                        _ = len(v2)
                                                                        _ = v76
                                                                    } else {
                                                                        fmt.Println(v34, v52)
                                                                    }
                                                                    v1 -= 0
                                                                    switch fi49 {
                                                                    case 1:
                                                                        v1 = (v1 * (v30 ^ (v30 - v30)))
                                                                    default:
                                                                        v34 -= v52
                                                                    }
                                                                }
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    switch -737 {
                                                    case 1:
                                                        switch 1 {
                                                        case 0:
                                                            var v77 int = fi35
                                                            _ = v77
                                                        case 32767:
                                                            if ((1 <= -14) || false) {
                                                                if (255 >= 0) {
                                                                    fmt.Println(v30, v1, v34)
                                                                    v78 := GenS9{F0: (-(v30)), F1: (v5 << ((255) & 63))}
                                                                    _ = v78
                                                                }
                                                                v79 := GenS8{F0: v5, F1: (fi49 << ((42) & 63))}
                                                                _ = v79
                                                            } else {
                                                                if true {
                                                                    delete(v2, (fi35 | fi49))
                                                                    switch (fi49 << ((495) & 63)) {
                                                                    case 655:
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        v80 := v2[(fi49 - (fi49 % ((fi49) | 1)))]
                                                                        _ = v80
                                                                    case 157:
                                                                        v54 += -128
                                                                    case 421:
                                                                        v1 -= (v30 ^ v30)
                                                                    default:
                                                                        var v81 int8 = 0
                                                                        _ = v81
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        switch fi35 {
                                                                        case 127:
                                                                            switch fi35 {
                                                                            case 832:
                                                                                v82 := GenS7{F0: 761, F1: 32767, F2: 304}
                                                                                _ = v82
                                                                            case 365:
                                                                                v34 = v34
                                                                            case 1:
                                                                                v34 -= (v34 >> ((4294967295) & 63))
                                                                            }
                                                                        default:
                                                                            v83 := GenS2{F0: 709}
                                                                            _ = v83
                                                                        }
                                                                    case "foo bar":
                                                                        var v84 int32 = (v52 / ((v34) | 1))
                                                                        _ = v84
                                                                    }
                                                                } else {
                                                                    var v85 int8 = 1
                                                                    _ = v85
                                                                }
                                                                fmt.Println(v1, v1, v54)
                                                                if true {
                                                                    fmt.Println(fi35, v52, v54)
                                                                    v15 = v15
                                                                    if true {
                                                                        _ = len(v3)
                                                                    } else {
                                                                        if (!(true && (296 >= 308))) {
                                                                            v86 := GenS1{F0: 2147483647, F1: 18446744073709551615, F2: (v5 % (((v5 | 895)) | 1))}
                                                                            fmt.Println(v34, v15)
                                                                            _ = v86
                                                                        } else {
                                                                            v87 := GenS10{F0: false, F1: v5}
                                                                            v88 := v2[fi35]
                                                                            _ = v87
                                                                            _ = v88
                                                                        }
                                                                        var v89 uint16 = 65535
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        _ = v89
                                                                    }
                                                                } else {
                                                                    v15 -= (1e10 - float64(1e10))
                                                                    var v90 int64 = (v1 + v30)
                                                                    _ = v90
                                                                }
                                                            }
                                                        case -97:
                                                            var v91 int32 = v34
                                                            _ = v91
                                                        }
                                                        if true {
                                                            v52 = 42
                                                        }
                                                        v15 += (v15 - v15)
                                                        var v92 string = ("foo bar" + "hello")
                                                        _ = v92
                                                    case -856:
                                                        v93 := GenS11{F0: 2147483647, F1: ("\n\t" + ("a" + ("hello" + ("" + "foo bar"))))}
                                                        _ = v93
                                                    case -2147483648:
                                                        var v94 int = (fi35 - (fi49 << ((18446744073709551615) & 63)))
                                                        v94 -= (fi35 / ((32767) | 1))
                                                        if len(v3) > 0 {
                                                            v3[0] = 2147483647
                                                        }
                                                        _ = v94
                                                    default:
                                                        var v95 string = "hello"
                                                        _ = v95
                                                    }
                                                    v96 := GenS12{F0: 1, F1: uint8(v5), F2: 255}
                                                    _ = v96
                                                }
                                            }
                                        }
                                        _ = v51
                                        _ = v52
                                        _ = v54
                                    } else {
                                        switch "hello" {
                                        case "a":
                                            if (255 > 210) {
                                                _ = len(v3)
                                                fmt.Println(v34, v30)
                                                v15 = 1e10
                                                fmt.Println(v30, fi35, v5)
                                                if true {
                                                    v97 := GenS7{F0: 656, F1: 2, F2: (v34 << ((197) & 63))}
                                                    if (false && true) {
                                                        v98 := GenS13{F0: fi35, F1: 32767, F2: 42}
                                                        _ = v98
                                                    }
                                                    _ = v97
                                                } else {
                                                    var v99 uint64 = (v5 + (v5 | 345))
                                                    if true {
                                                        v5 -= (v5 - (v5 + (v99 / ((0) | 1))))
                                                        if len(v3) > 0 {
                                                            v3[0] = 966
                                                        }
                                                        switch 468 {
                                                        case -9223372036854775808:
                                                            v1 = v1
                                                        default:
                                                            v100 := GenS14{F0: 42}
                                                            var v101 int32 = 1
                                                            _ = v100
                                                            _ = v101
                                                        }
                                                    } else {
                                                        if (!(0.0 == (1e10 / float32(v15)))) {
                                                            if true {
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    v15 -= v15
                                                                case "\n\t":
                                                                    fmt.Println(v34, v1, v99)
                                                                case "hello":
                                                                    v99 = v99
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v15, fi35, v34)
                                                    }
                                                    v34 = (v34 - (v34 - (v34 << ((415) & 63))))
                                                    _ = v99
                                                }
                                            }
                                        case "foo bar":
                                            if true {
                                                v102 := GenS7{F0: 2, F1: 629, F2: -1}
                                                v103 := GenS10{F0: (255 >= 127), F1: v5}
                                                _ = v102
                                                _ = v103
                                            } else {
                                                v5 -= 4294967295
                                                v104 := v2[fi49]
                                                v2[fi35] = (0.0 < 1.5)
                                                v15 = (0.1 / (v15 / v15))
                                                _ = v104
                                            }
                                        case "\n\t":
                                            _ = len(v3)
                                        }
                                    }
                                }
                                _ = fi49
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (true || (0.0 == v15)) {
                                v15 -= v15
                            } else {
                                v5 = 719
                                v105 := GenS4{F0: v30, F1: v1, F2: -1}
                                _ = v105
                            }
                        }
                    }
                }
                if false {
                    if (v34 > (v34 & v34)) {
                        var v106 int16 = 1000
                        _ = v106
                    } else {
                        var v107 uint64 = (v5 << ((4294967295) & 63))
                        _ = v107
                    }
                    v1 = 9223372036854775807
                    switch 0 {
                    case 127:
                        var v108 int8 = -128
                        _ = v108
                    case 0:
                        v3 = append(v3, 2)
                    }
                }
                var v109 int32 = (v34 % (((v34 | -1)) | 1))
                _ = fi35
                _ = v109
            }
            _ = v34
        }
        _ = v30
    }
    switch 500 {
    case -1:
        for fi110 := 0; fi110 < 11; fi110++ {
            var v111 float32 = 0.1
            for fi112 := 0; fi112 < 20; fi112++ {
                if (!true) {
                    v15 = v15
                    v113 := GenS13{F0: (fi110 + (fi112 / ((fi110) | 1))), F1: 1, F2: (v5 + 42)}
                    var v114 string = ("" + "a")
                    _ = v113
                    _ = v114
                }
                delete(v2, 1)
                v1 -= int64(v5)
                v5 += v5
                v15 -= (((v15 * float64(-2.25)) / v15) - float64(v111))
                v15 -= (v15 + (((-2.25 * v15) + 1e10) * ((v15 / ((v15 * 1.5) / v15)) / (3.14159 - v15))))
                _ = fi112
            }
            v111 += (v111 + float32((v15 - (((v15 / v15) / 1e10) * -2.25))))
            _ = fi110
            _ = v111
        }
    default:
        v5 += 59
    }
    switch -9223372036854775808 {
    case 22:
        for fi115 := 0; fi115 < 12; fi115++ {
            var v116 bool = true
            v15 -= ((v15 - v15) * (v15 / 0.1))
            _ = fi115
            _ = v116
        }
    }
    var v117 int16 = 175
    switch (("hello" + "a") + "a") {
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        if (false && true) {
            v118 := GenS15{F0: 100}
            var v119 int = -1
            _ = v118
            _ = v119
        }
    case "":
        if (true || (((0.1 / v15) / 0.1) > v15)) {
            switch -1 {
            case -9223372036854775808:
                v120 := GenS16{F0: (("foo bar" + "a") + "a"), F1: 1}
                _ = v120
            }
            if len(v3) > 0 {
                v3[0] = 476
            }
        }
    }
    if (512 != v5) {
        for fi121 := 0; fi121 < 1; fi121++ {
            for fi122 := 0; fi122 < 20; fi122++ {
                switch fi121 {
                case 127:
                    if false {
                        fmt.Println(fi121)
                        v123 := GenS10{F0: (42 <= 65535), F1: (v5 ^ (v5 | 18446744073709551615))}
                        _ = v123
                    }
                case -130:
                    var v124 int16 = (v117 << ((912) & 63))
                    fmt.Println(fi121)
                    _ = v124
                case 32767:
                    fmt.Println(v15, v15, v15)
                default:
                    v15 = v15
                    v1 = (-(v1))
                    v125 := v2[fi122]
                    _ = v125
                }
                if false {
                    var v126 float32 = (3.14159 / 0.1)
                    var v127 float32 = v126
                    var v128 uint = 42
                    _ = v126
                    _ = v127
                    _ = v128
                }
                _ = fi122
            }
            _ = fi121
        }
        v1 += 42
    } else {
        var v129 bool = ((false || false) || (42 >= 128))
        fmt.Println(v129, v5, v1)
        _ = v129
    }
    v15 = v15
    fmt.Println(v5, v15)
    for fi130 := 0; fi130 < 19; fi130++ {
        if ((v117 << ((255) & 63)) <= v117) {
            for fi131 := 0; fi131 < 17; fi131++ {
                _ = len(v2)
                v2[2147483647] = false
                v5 = (v5 + (v5 >> ((258) & 63)))
                fmt.Println(v117)
                _ = len(v3)
                _ = fi131
            }
            v132 := GenS1{F0: (fi130 % (((fi130 >> ((496) & 63))) | 1)), F1: 1, F2: 434}
            _ = v132
        } else {
            switch "\n\t" {
            case "":
                v2[fi130] = ((v117 % ((v117) | 1)) > v117)
            case "hello":
                var v133 float32 = 1e10
                _ = v133
            case "foo bar":
                var v134 int32 = -977
                _ = v134
            }
        }
        v117 = v117
        fmt.Println(fi130, v5, fi130)
        _ = fi130
    }
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v117
    _ = v15
    _ = v2
    _ = v3
    _ = v5
}
