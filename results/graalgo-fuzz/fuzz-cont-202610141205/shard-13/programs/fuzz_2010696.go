package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 uint64; F2 int }
type GenS7 struct { F0 uint8; F1 int; F2 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint; F1 int; F2 int32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS14 struct { F0 int; F1 uint16 }
type GenS15 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int64; F1 int8; F2 string }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 int16; F1 bool; F2 bool }
type GenS20 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS21 struct { F0 uint16; F1 int8; F2 bool }
type GenS22 struct { F0 uint64; F1 int16; F2 int }
type GenS23 struct { F0 uint32 }
type GenS24 struct { F0 uint8; F1 uint8 }
type GenS25 struct { F0 int8 }
type GenS26 struct { F0 string; F1 int16 }
type GenS27 struct { F0 bool; F1 uint8; F2 string }
type GenS28 struct { F0 uint32; F1 uint16 }
type GenS29 struct { F0 uint32; F1 int8 }
type GenS30 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS31 struct { F0 int16 }
type GenS32 struct { F0 uint16; F1 uint }
type GenS33 struct { F0 int32; F1 int8; F2 int }
type GenS34 struct { F0 int16; F1 int32; F2 int }
type GenS35 struct { F0 uint; F1 int }
type GenS36 struct { F0 int64 }
type GenS37 struct { F0 int32 }
type GenS38 struct { F0 int16 }
type GenS39 struct { F0 int64; F1 uint16 }
type GenS40 struct { F0 int }
type GenS41 struct { F0 string; F1 int16 }
type GenS42 struct { F0 int32 }
type GenS43 struct { F0 int8 }
type GenS44 struct { F0 uint64 }
type GenS45 struct { F0 uint8 }
type GenS46 struct { F0 int32 }
type GenS47 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS48 struct { F0 uint16 }
type GenS49 struct { F0 string; F1 int32; F2 int64 }
type GenS50 struct { F0 uint8 }
type GenS51 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS52 struct { F0 int64 }
type GenS53 struct { F0 int8; F1 uint8; F2 int }
type GenS54 struct { F0 int32; F1 int }
type GenS55 struct { F0 int8 }
type GenS56 struct { F0 int64; F1 uint8 }
func genFn0() bool {
    v4 := GenS1{F0: 121, F1: 128}
    var v5 int8 = 15
    _ = v4
    _ = v5
    return true
}

func main() {
    var v1 string = (("" + ("" + "hello")) + ("hello" + "foo bar"))
    v2 := []uint32{149, 1}
    var v3 int = -101
    switch (v3 + (v3 + (v3 - v3))) {
    case 42:
        for fi6 := 0; fi6 < 7; fi6++ {
            switch "foo bar" {
            case "hello":
                var v7 bool = true
                _ = v7
            case "":
                if false {
                    var v8 uint8 = 1
                    if ((true && true) || true) {
                        for fi9 := 0; fi9 < 7; fi9++ {
                            var v10 int16 = -1
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                            _ = fi9
                            _ = v10
                        }
                        var v11 int16 = 382
                        fmt.Println(fi6, fi6, v8)
                        _ = v11
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v8
                } else {
                    var v12 float64 = ((-2.25 - 3.14159) * 1e10)
                    var v13 float32 = (-2.25 / float32((v12 * v12)))
                    _ = v12
                    _ = v13
                }
            case "\n\t":
                v3 = (fi6 << ((740) & 63))
            default:
                v2 = append(v2, 2147483647)
            }
            _ = fi6
        }
    case -632:
        for fi14 := 0; fi14 < 8; fi14++ {
            if false {
                for fi15 := 0; fi15 < 20; fi15++ {
                    switch v3 {
                    case -2147483648:
                        if (1.5 == 1.5) {
                            v3 -= (v3 >> ((652) & 63))
                        }
                    case 373:
                        v16 := GenS1{F0: 128, F1: (^(fi15))}
                        _ = v16
                    case 428:
                        if (127 <= 0) {
                            v3 = 128
                            var v17 uint = 0
                            var v18 uint8 = 1
                            v19 := GenS2{F0: 65535, F1: -436}
                            _ = v17
                            _ = v18
                            _ = v19
                        } else {
                            v20 := GenS2{F0: 317, F1: -846}
                            v21 := GenS2{F0: 64, F1: -32768}
                            v1 = v1
                            switch v1 {
                            case "a":
                                var v22 int16 = 32767
                                _ = v22
                            default:
                                v3 = fi14
                            }
                            _ = v20
                            _ = v21
                        }
                    }
                    _ = fi15
                }
                if false {
                    switch (fi14 * (v3 >> ((2) & 63))) {
                    case -169:
                        for fi23 := 0; fi23 < 1; fi23++ {
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                            fmt.Println(v3, fi14)
                            v1 = ((v1 + "a") + v1)
                            _ = fi23
                        }
                    case -9223372036854775808:
                        v2 = append(v2, 1)
                    }
                    switch (fi14 * 128) {
                    case 149:
                        for fi24 := 0; fi24 < 8; fi24++ {
                            v3 = (v3 ^ (v3 << ((1) & 63)))
                            v1 = "foo bar"
                            _ = fi24
                        }
                    case 127:
                        var v25 uint32 = 320
                        _ = v25
                    case -1:
                        if false {
                            for fi26 := 0; fi26 < 16; fi26++ {
                                v2 = append(v2, 0)
                                var v27 uint64 = 42
                                v28 := GenS3{F0: 1}
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        }
                    default:
                        for fi29 := 0; fi29 < 11; fi29++ {
                            var v30 int = fi29
                            if (1000 > -32768) {
                                v30 += fi14
                            }
                            if false {
                                var v31 uint16 = 803
                                fmt.Println(v3, v31, v1)
                                var v32 uint16 = 533
                                _ = v31
                                _ = v32
                            } else {
                                var v33 uint8 = 2
                                v33 -= (v33 * 137)
                                if (fi14 > fi14) {
                                    fmt.Println(fi29, v33)
                                    v1 = "foo bar"
                                }
                                _ = v33
                            }
                            switch "foo bar" {
                            case "a":
                                var v34 int = (v30 * (fi29 << ((1) & 63)))
                                _ = v34
                            }
                            var v35 uint16 = 1
                            _ = fi29
                            _ = v30
                            _ = v35
                        }
                    }
                } else {
                    if false {
                        v3 += 2147483647
                    } else {
                        var v36 uint16 = 0
                        if len(v2) > 0 {
                            v2[0] = 805
                        }
                        _ = v36
                    }
                    fmt.Println(v3)
                    fmt.Println(v3)
                }
                for fi37 := 0; fi37 < 11; fi37++ {
                    fmt.Println(fi14)
                    fmt.Println(fi37)
                    v3 += fi37
                    _ = fi37
                }
            } else {
                v1 = "hello"
                if (!false) {
                    v38 := GenS1{F0: 255, F1: 628}
                    _ = v38
                } else {
                    if true {
                        fmt.Println(fi14, fi14)
                    } else {
                        v3 = -9223372036854775808
                        if (!(0 >= 1)) {
                            fmt.Println(v3, fi14)
                            v39 := GenS3{F0: 1}
                            _ = v39
                        }
                        v1 = (v1 + "hello")
                    }
                }
            }
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch fi14 {
                case 2147483647:
                    fmt.Println(v3, v3)
                case -9223372036854775808:
                    switch v3 {
                    case 4294967295:
                        v40 := GenS1{F0: 128, F1: 9223372036854775807}
                        _ = v40
                    case 9223372036854775807:
                        _ = len(v2)
                    case 42:
                        for fi41 := 0; fi41 < 16; fi41++ {
                            fmt.Println(v3, fi14)
                            if ((((((v1 + "hello") + v1) + v1) + (v1 + v1)) + (v1 + v1)) <= v1) {
                                if (0 >= 674) {
                                    var v42 int = v3
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v42, fi41)
                                    _ = v42
                                } else {
                                    v43 := GenS4{F0: 402, F1: 206, F2: 1}
                                    _ = v43
                                }
                                v1 = (v1 + (v1 + v1))
                            }
                            var v44 uint32 = 65535
                            if false {
                                switch 127 {
                                case 127:
                                    v45 := GenS2{F0: 255, F1: 42}
                                    _ = v45
                                case 32767:
                                    if (1e10 == 1e10) {
                                        var v46 int = 676
                                        var v47 int = (v46 | v3)
                                        v2 = append(v2, (v44 - v44))
                                        _ = v46
                                        _ = v47
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if false {
                                            if false {
                                                if true {
                                                    switch fi14 {
                                                    case 9223372036854775807:
                                                        v48 := GenS5{F0: "foo bar"}
                                                        _ = v48
                                                    case -646:
                                                        v49 := GenS5{F0: v1}
                                                        _ = v49
                                                    }
                                                    switch v1 {
                                                    case "":
                                                        var v50 float64 = (float64(0.0) - 3.14159)
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v50
                                                    default:
                                                        v51 := GenS4{F0: 1, F1: 0, F2: 65535}
                                                        _ = v51
                                                    }
                                                }
                                                var v52 int8 = 100
                                                fmt.Println(fi14, fi41)
                                                _ = v52
                                            }
                                        }
                                    }
                                case -1:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                if true {
                                    v53 := GenS6{F0: v44, F1: 0, F2: fi41}
                                    var v54 int = -2147483648
                                    _ = v53
                                    _ = v54
                                } else {
                                    fmt.Println(v3, fi14, v44)
                                    if false {
                                        switch "hello" {
                                        case "hello":
                                            if len(v2) > 0 {
                                                v2[0] = v44
                                            }
                                        case "foo bar":
                                            var v55 uint32 = v44
                                            _ = v55
                                        default:
                                            switch (v1 + v1) {
                                            case "hello":
                                                v56 := GenS7{F0: 42, F1: (-(v3)), F2: 42}
                                                _ = v56
                                            case "foo bar":
                                                fmt.Println(fi14, fi14)
                                            }
                                        }
                                    }
                                }
                            } else {
                                v44 = 65535
                                if true {
                                    v57 := GenS8{F0: 51}
                                    _ = v57
                                }
                                v3 = (-(fi14))
                            }
                            var v58 float32 = -2.25
                            var v59 int32 = -1
                            fmt.Println(v1)
                            var v60 uint8 = 78
                            switch (v3 >> ((0) & 63)) {
                            case 32768:
                                v61 := GenS8{F0: 0}
                                _ = v61
                            case -1:
                                if false {
                                    v62 := GenS2{F0: 65535, F1: 224}
                                    var v63 uint16 = 32767
                                    var v64 uint = 0
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                } else {
                                    _ = len(v2)
                                    var v65 uint8 = 184
                                    if (true && false) {
                                        v66 := GenS9{F0: 18446744073709551615, F1: 817, F2: 2147483647}
                                        v67 := GenS7{F0: v65, F1: v3, F2: -1}
                                        _ = v66
                                        _ = v67
                                    } else {
                                        if true {
                                            fmt.Println(v59, v58)
                                            var v68 uint32 = v44
                                            _ = v68
                                        }
                                    }
                                    _ = v65
                                }
                            case 449:
                                if false {
                                    v60 = (v60 & v60)
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v59)
                                    if (true && false) {
                                        if false {
                                            v69 := GenS3{F0: v44}
                                            _ = len(v2)
                                            fmt.Println(fi41)
                                            _ = v69
                                        }
                                    } else {
                                        switch v1 {
                                        case "hello":
                                            switch 127 {
                                            case 660:
                                                v44 -= v44
                                                var v70 bool = (65535 <= 32767)
                                                var v71 int8 = 5
                                                _ = v70
                                                _ = v71
                                            case 32768:
                                                v72 := GenS6{F0: v44, F1: 42, F2: v3}
                                                _ = v72
                                            default:
                                                var v73 float32 = ((0.1 + 1.5) - v58)
                                                v74 := GenS10{F0: (v59 & (v59 >> ((2) & 63)))}
                                                _ = v73
                                                _ = v74
                                            }
                                        case "":
                                            v44 += (v44 * v44)
                                        default:
                                            switch v1 {
                                            case "hello":
                                                var v75 int = v3
                                                _ = v75
                                            }
                                        }
                                        if true {
                                            v59 = v59
                                            var v76 bool = (42 == 127)
                                            if v76 {
                                                v3 = 32767
                                            }
                                            _ = v76
                                        }
                                    }
                                }
                            }
                            if (!false) {
                                switch fi41 {
                                case -545:
                                    v58 -= v58
                                case -2147483648:
                                    switch (fi14 | (fi14 * fi41)) {
                                    case 32768:
                                        var v77 float32 = 3.14159
                                        _ = v77
                                    case 136:
                                        if true {
                                            switch ("\n\t" + ("\n\t" + (v1 + "a"))) {
                                            case "hello":
                                                var v78 uint64 = 2
                                                _ = v78
                                            default:
                                                if true {
                                                    switch (fi41 * v3) {
                                                    case 1:
                                                        v60 += (v60 * (v60 * (v60 >> ((65535) & 63))))
                                                    case 128:
                                                        switch 578 {
                                                        case -2147483648:
                                                            v79 := GenS11{F0: v1}
                                                            _ = v79
                                                        case 323:
                                                            switch v3 {
                                                            case 981:
                                                                var v80 string = v1
                                                                _ = v80
                                                            case -9223372036854775808:
                                                                var v81 int32 = (v59 - v59)
                                                                _ = v81
                                                            }
                                                        }
                                                    case 626:
                                                        var v82 float64 = -2.25
                                                        _ = v82
                                                    default:
                                                        if false {
                                                            v58 -= float32(float64(float32(1.5)))
                                                            var v83 string = v1
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            _ = v83
                                                        }
                                                    }
                                                    var v84 bool = (!(!(!(9223372036854775807 > 966))))
                                                    switch ("hello" + v1) {
                                                    case "":
                                                        switch ("hello" + (("" + (v1 + v1)) + "")) {
                                                        case "\n\t":
                                                            switch 32768 {
                                                            case -240:
                                                                v1 = v1
                                                            default:
                                                                v60 = (v60 >> ((18446744073709551615) & 63))
                                                            }
                                                        case "a":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        case "foo bar":
                                                            v3 -= 1
                                                        default:
                                                            v1 = v1
                                                        }
                                                    case "\n\t":
                                                        v59 = v59
                                                    case "a":
                                                        var v85 uint32 = 4294967295
                                                        switch 4294967295 {
                                                        case 9223372036854775807:
                                                            switch "hello" {
                                                            case "":
                                                                switch v1 {
                                                                case "foo bar":
                                                                    switch (-(v3)) {
                                                                    case 2147483647:
                                                                        if len(v2) > 0 {
                                                                            v2[0] = v44
                                                                        }
                                                                    case -1:
                                                                        v59 = (^(v59))
                                                                    case 9223372036854775807:
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    default:
                                                                        if len(v2) > 0 {
                                                                            v2[0] = v44
                                                                        }
                                                                    }
                                                                case "":
                                                                    v86 := GenS9{F0: 35, F1: fi14, F2: v59}
                                                                    _ = v86
                                                                case "a":
                                                                    if v84 {
                                                                        v1 = ""
                                                                    }
                                                                    var v87 int32 = int32(v60)
                                                                    if len(v2) > 0 {
                                                                        v2[0] = (v85 & v85)
                                                                    }
                                                                    _ = v87
                                                                default:
                                                                    v84 = v84
                                                                }
                                                            }
                                                        case 127:
                                                            switch ((((v1 + ((((v1 + v1) + "\n\t") + (v1 + ("" + (v1 + "")))) + v1)) + v1) + "\n\t") + v1) {
                                                            case "\n\t":
                                                                if (false || v84) {
                                                                    v60 = 255
                                                                    v1 = v1
                                                                } else {
                                                                    fmt.Println(v3)
                                                                }
                                                            }
                                                        case 0:
                                                            var v88 bool = v84
                                                            var v89 int64 = -445
                                                            _ = v88
                                                            _ = v89
                                                        default:
                                                            var v90 uint64 = 18446744073709551615
                                                            if ((^(v60)) < v60) {
                                                                var v91 uint = 18446744073709551615
                                                                v90 += v90
                                                                _ = v91
                                                            }
                                                            var v92 int = (^(fi41))
                                                            _ = v90
                                                            _ = v92
                                                        }
                                                        fmt.Println(v44, v58)
                                                        _ = v85
                                                    default:
                                                        var v93 int32 = v59
                                                        fmt.Println(v60)
                                                        v60 += (v60 % ((v60) | 1))
                                                        _ = v93
                                                    }
                                                    if ((65535 >= 255) || ((-(v59)) < v59)) {
                                                        var v94 int16 = -32768
                                                        fmt.Println(v1, v1, fi41)
                                                        _ = v94
                                                    } else {
                                                        if (523 >= 0) {
                                                            v95 := GenS12{F0: (v59 + (v59 << ((18446744073709551615) & 63)))}
                                                            _ = v95
                                                        } else {
                                                            if ((v84 || true) || (v84 && v84)) {
                                                                v96 := GenS8{F0: 2}
                                                                var v97 uint = 18446744073709551615
                                                                _ = v96
                                                                _ = v97
                                                            } else {
                                                                switch fi41 {
                                                                case 9223372036854775807:
                                                                    v58 = 1.5
                                                                }
                                                            }
                                                            fmt.Println(v60)
                                                            v98 := GenS9{F0: 65535, F1: fi41, F2: 2147483647}
                                                            _ = v98
                                                        }
                                                        if ((v84 || (-1 < -128)) || (32768 > (v3 - -2147483648))) {
                                                            if (0.0 > float64(v58)) {
                                                                switch v1 {
                                                                case "":
                                                                    if (v84 && ((!true) && false)) {
                                                                        v99 := GenS9{F0: 42, F1: fi41, F2: (v59 >> ((255) & 63))}
                                                                        var v100 int64 = 9223372036854775807
                                                                        fmt.Println(v44, v58)
                                                                        _ = v100
                                                                        _ = v99
                                                                    } else {
                                                                        if ("" >= (v1 + v1)) {
                                                                            switch v1 {
                                                                            case "a":
                                                                                if (!(1 == 756)) {
                                                                                    var v101 uint = 65535
                                                                                    _ = v101
                                                                                } else {
                                                                                    var v102 int16 = 42
                                                                                    v60 += 127
                                                                                    _ = v102
                                                                                }
                                                                            case "":
                                                                                v103 := GenS13{F0: v44, F1: 42, F2: 0}
                                                                                v44 = (^(v44))
                                                                                fmt.Println(v1, v84)
                                                                                _ = v103
                                                                            }
                                                                        } else {
                                                                            switch v3 {
                                                                            case 1:
                                                                                switch (fi41 << ((68) & 63)) {
                                                                                case 4294967295:
                                                                                    var v104 int64 = -9223372036854775808
                                                                                    _ = v104
                                                                                }
                                                                            }
                                                                        }
                                                                        if (false && v84) {
                                                                            if ((v44 | 859) >= v44) {
                                                                                v105 := GenS8{F0: 973}
                                                                                fmt.Println(fi14, fi14)
                                                                                if (-2.25 > float32((1e10 / ((-2.25 * (float64(((-2.25 / (((v58 + v58) / float32((3.14159 + (((((1.5 - (0.0 + 0.0)) + 0.0) * 1.5) * 0.0) - 1e10)))) / v58)) + v58)) / 3.14159)) * float64(v58))))) {
                                                                                    fmt.Println(v60, v3)
                                                                                    v2 = append(v2, v44)
                                                                                    v3 -= (v3 / (((v3 / (((v3 * fi41)) | 1))) | 1))
                                                                                } else {
                                                                                    fmt.Println(v58, v44, v58)
                                                                                    v106 := GenS14{F0: (^(fi41)), F1: 0}
                                                                                    switch (("" + v1) + ("hello" + "\n\t")) {
                                                                                    case "hello":
                                                                                        v2 = append(v2, v44)
                                                                                    case "\n\t":
                                                                                        var v107 bool = (v84 && true)
                                                                                        _ = v107
                                                                                    default:
                                                                                        switch ((v1 + (v1 + "foo bar")) + v1) {
                                                                                        case "foo bar":
                                                                                            switch (fi14 + (fi14 % ((723) | 1))) {
                                                                                            case -2147483648:
                                                                                                switch -9223372036854775808 {
                                                                                                case 765:
                                                                                                    v3 = (fi14 * -142)
                                                                                                }
                                                                                            case -251:
                                                                                                if (127 != v60) {
                                                                                                    var v108 uint32 = 2147483648
                                                                                                    var v109 int64 = -9223372036854775808
                                                                                                    _ = v108
                                                                                                    _ = v109
                                                                                                } else {
                                                                                                    v59 -= 1
                                                                                                }
                                                                                            }
                                                                                            v84 = v84
                                                                                            var v110 float64 = (0.1 - 0.1)
                                                                                            _ = v110
                                                                                        case "hello":
                                                                                            v2 = append(v2, (^(v44)))
                                                                                        case "":
                                                                                            v1 = (v1 + v1)
                                                                                        }
                                                                                    }
                                                                                    _ = v106
                                                                                }
                                                                                _ = v105
                                                                            }
                                                                            switch v3 {
                                                                            case 0:
                                                                                v2 = append(v2, (v44 + 4294967295))
                                                                            }
                                                                        }
                                                                        var v111 float64 = 0.1
                                                                        _ = v111
                                                                    }
                                                                case "a":
                                                                    v44 -= v44
                                                                    var v112 uint8 = 195
                                                                    if (!(360 != 4294967295)) {
                                                                        var v113 uint64 = 260
                                                                        v1 = v1
                                                                        if v84 {
                                                                            var v114 string = "a"
                                                                            v113 = (^(v113))
                                                                            _ = v114
                                                                        } else {
                                                                            if true {
                                                                                var v115 uint64 = 18446744073709551615
                                                                                var v116 float32 = (v58 - (3.14159 * (-2.25 * (float32(float64((1.5 - v58))) + 1.5))))
                                                                                v112 = 149
                                                                                _ = v115
                                                                                _ = v116
                                                                            } else {
                                                                                fmt.Println(v3)
                                                                                switch v1 {
                                                                                case "hello":
                                                                                    v117 := GenS11{F0: v1}
                                                                                    _ = v117
                                                                                default:
                                                                                    switch (v1 + (v1 + (v1 + ((("" + "foo bar") + "hello") + v1)))) {
                                                                                    case "\n\t":
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = v44
                                                                                        }
                                                                                    case "foo bar":
                                                                                        var v118 bool = v84
                                                                                        _ = v118
                                                                                    default:
                                                                                        v119 := GenS8{F0: 2}
                                                                                        _ = v119
                                                                                    }
                                                                                }
                                                                                var v120 int64 = 0
                                                                                _ = v120
                                                                            }
                                                                            v3 = fi41
                                                                            var v121 string = (v1 + ("a" + "\n\t"))
                                                                            _ = v121
                                                                        }
                                                                        _ = v113
                                                                    }
                                                                    _ = v112
                                                                case "hello":
                                                                    var v122 uint = 2
                                                                    _ = v122
                                                                default:
                                                                    var v123 string = v1
                                                                    var v124 uint32 = v44
                                                                    switch ("a" + v1) {
                                                                    case "":
                                                                        v59 = v59
                                                                        switch fi14 {
                                                                        case -464:
                                                                            v125 := GenS8{F0: 697}
                                                                            _ = v125
                                                                        case 127:
                                                                            if v84 {
                                                                                var v126 uint16 = 559
                                                                                var v127 int16 = -429
                                                                                _ = v126
                                                                                _ = v127
                                                                            } else {
                                                                                if (-72 <= 42) {
                                                                                    if (false || (538 > 935)) {
                                                                                        if ((((3.14159 / 1.0) * 3.14159) * (0.0 / float64(v58))) >= ((-2.25 * 0.1) * ((3.14159 - ((3.14159 / 0.1) / 0.1)) + (-2.25 - (3.14159 * (1.5 / 1.0)))))) {
                                                                                            if false {
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = v44
                                                                                                }
                                                                                            } else {
                                                                                                v128 := GenS15{F0: (42 >= 1), F1: 818, F2: 65535}
                                                                                                v84 = v84
                                                                                                fmt.Println(v44, v58)
                                                                                                _ = v128
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(fi41)
                                                                                        var v129 uint64 = 1
                                                                                        fmt.Println(v60)
                                                                                        v130 := GenS16{F0: -884}
                                                                                        _ = v129
                                                                                        _ = v130
                                                                                    }
                                                                                    switch v1 {
                                                                                    case "hello":
                                                                                        var v131 bool = v84
                                                                                        _ = v131
                                                                                    case "foo bar":
                                                                                        v3 += fi41
                                                                                    case "a":
                                                                                        v44 += (v44 | 787)
                                                                                    default:
                                                                                        v132 := GenS17{F0: -9223372036854775808, F1: 1, F2: ("hello" + ((v123 + (v123 + "\n\t")) + v1))}
                                                                                        _ = v132
                                                                                    }
                                                                                    fmt.Println(v84, v124)
                                                                                }
                                                                                var v133 uint = 0
                                                                                _ = v133
                                                                            }
                                                                        default:
                                                                            var v134 int32 = v59
                                                                            var v135 string = ("" + "hello")
                                                                            if (-2147483648 == (v3 << ((4294967295) & 63))) {
                                                                                v123 = v1
                                                                            }
                                                                            _ = v134
                                                                            _ = v135
                                                                        }
                                                                        var v136 uint16 = uint16(v60)
                                                                        _ = v136
                                                                    case "\n\t":
                                                                        fmt.Println(v60, v123)
                                                                    case "hello":
                                                                        var v137 string = v1
                                                                        _ = v137
                                                                    }
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v124
                                                                    }
                                                                    _ = v123
                                                                    _ = v124
                                                                }
                                                                v58 = (float32((1e10 - -2.25)) - 3.14159)
                                                                var v138 float64 = (float64((v58 / (-2.25 * 1e10))) - float64((0.0 - v58)))
                                                                if v84 {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v44
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        switch ((("a" + v1) + "") + v1) {
                                                                        case "hello":
                                                                            v1 = "hello"
                                                                        }
                                                                    } else {
                                                                        v60 = 0
                                                                    }
                                                                }
                                                                _ = len(v2)
                                                                v139 := GenS1{F0: v60, F1: 127}
                                                                _ = v138
                                                                _ = v139
                                                            } else {
                                                                var v140 uint8 = uint8(v44)
                                                                _ = v140
                                                            }
                                                        }
                                                        if v84 {
                                                            if ((!(v84 || v84)) && false) {
                                                                v141 := GenS6{F0: v44, F1: 948, F2: (-(fi14))}
                                                                v142 := GenS13{F0: 2, F1: (v44 >> ((42) & 63)), F2: -468}
                                                                _ = v141
                                                                _ = v142
                                                            }
                                                            if (!false) {
                                                                fmt.Println(v44, v1, v3)
                                                                v60 = v60
                                                            }
                                                            if (42 <= 0) {
                                                                var v143 uint32 = v44
                                                                var v144 uint32 = (^(v143))
                                                                var v145 int = (v3 * (fi14 >> ((0) & 63)))
                                                                _ = v143
                                                                _ = v144
                                                                _ = v145
                                                            } else {
                                                                fmt.Println(v59, v59)
                                                                switch (v1 + v1) {
                                                                case "a":
                                                                    var v146 uint16 = 0
                                                                    _ = v146
                                                                case "foo bar":
                                                                    switch v1 {
                                                                    case "foo bar":
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    case "\n\t":
                                                                        v147 := GenS7{F0: (v60 / (((v60 >> ((4294967295) & 63))) | 1)), F1: (fi14 | (fi14 | fi14)), F2: 9223372036854775807}
                                                                        _ = v147
                                                                    case "hello":
                                                                        if (fi41 == 0) {
                                                                            v148 := GenS18{F0: 907}
                                                                            switch 128 {
                                                                            case -596:
                                                                                v58 = -2.25
                                                                            default:
                                                                                v44 = 90
                                                                            }
                                                                            var v149 int64 = -9223372036854775808
                                                                            _ = v148
                                                                            _ = v149
                                                                        } else {
                                                                            v60 += v60
                                                                            switch (v1 + v1) {
                                                                            case "foo bar":
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = (v44 * (v44 >> ((917) & 63)))
                                                                                }
                                                                            case "\n\t":
                                                                                v150 := GenS19{F0: -713, F1: (((0.1 + -2.25) + 0.0) <= (-2.25 - ((1e10 / 1e10) * float64((((v58 / ((v58 / (1e10 + ((v58 * (-2.25 / (v58 * v58))) / v58))) + (v58 + v58))) * (v58 / float32((0.1 / ((float64(((v58 * v58) * v58)) + float64(v58)) * float64(float32((3.14159 / 0.1)))))))) - float32((0.1 - (1e10 + 1.5)))))))), F2: v84}
                                                                                _ = v150
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    v151 := GenS20{F0: 1, F1: uint16(v44), F2: 690}
                                                                    _ = v151
                                                                }
                                                                var v152 uint = uint(v44)
                                                                _ = v152
                                                            }
                                                        }
                                                    }
                                                    v153 := GenS21{F0: 2, F1: 0, F2: true}
                                                    switch 1 {
                                                    case 2147483647:
                                                        var v154 float32 = float32((0.0 + 1.5))
                                                        _ = v154
                                                    case 0:
                                                        if v84 {
                                                            v58 -= v58
                                                        }
                                                    case 32768:
                                                        fmt.Println(v3)
                                                    default:
                                                        var v155 int = fi41
                                                        _ = v155
                                                    }
                                                    if v84 {
                                                        fmt.Println(v1, v44)
                                                        v156 := GenS22{F0: 42, F1: -32768, F2: 128}
                                                        switch fi41 {
                                                        case 2147483647:
                                                            v157 := GenS22{F0: 18446744073709551615, F1: -321, F2: v3}
                                                            var v158 string = (v1 + v1)
                                                            _ = v157
                                                            _ = v158
                                                        case -1:
                                                            var v159 int16 = -378
                                                            _ = v159
                                                        default:
                                                            v1 = v1
                                                        }
                                                        _ = v156
                                                    } else {
                                                        v60 = 2
                                                    }
                                                    _ = v153
                                                    _ = v84
                                                } else {
                                                    var v160 bool = ((false || true) || false)
                                                    _ = v160
                                                }
                                            }
                                            var v161 string = (v1 + "hello")
                                            _ = v161
                                        }
                                    }
                                default:
                                    var v162 uint32 = 875
                                    _ = v162
                                }
                            } else {
                                v1 = (v1 + v1)
                                switch v1 {
                                case "foo bar":
                                    v163 := GenS17{F0: 1, F1: 1, F2: v1}
                                    _ = v163
                                case "\n\t":
                                    v58 = (float32(float64(v58)) + (v58 * float32(3.14159)))
                                }
                                fmt.Println(v58)
                            }
                            _ = fi41
                            _ = v44
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                    default:
                        switch ("hello" + v1) {
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                        case "hello":
                            v164 := GenS20{F0: 0, F1: 0, F2: 2}
                            _ = v164
                        case "foo bar":
                            var v165 uint16 = 255
                            _ = v165
                        default:
                            for fi166 := 0; fi166 < 16; fi166++ {
                                v167 := GenS23{F0: 2147483647}
                                var v168 int32 = 0
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v169 := GenS24{F0: 127, F1: 127}
                                if false {
                                    fmt.Println(v3, fi14, v3)
                                    fmt.Println(v3, v1, fi14)
                                    v2 = append(v2, 4294967295)
                                } else {
                                    v170 := GenS25{F0: -121}
                                    if len(v2) > 0 {
                                        v2[0] = 2147483648
                                    }
                                    _ = v170
                                }
                                switch fi14 {
                                case 2147483647:
                                    if true {
                                        v168 -= (v168 + (v168 & (v168 | v168)))
                                        v171 := GenS5{F0: (v1 + v1)}
                                        _ = v171
                                    } else {
                                        if false {
                                            var v172 float64 = -2.25
                                            _ = v172
                                        } else {
                                            v168 += (v168 * -335)
                                            v173 := GenS26{F0: v1, F1: 42}
                                            v168 = -2147483648
                                            var v174 float64 = 1.5
                                            _ = v173
                                            _ = v174
                                        }
                                        fmt.Println(fi166, v3, fi14)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                case 1:
                                    if len(v2) > 0 {
                                        v2[0] = 42
                                    }
                                }
                                fmt.Println(v3, fi14)
                                v1 = ((v1 + v1) + (v1 + ("\n\t" + v1)))
                                var v175 int8 = 127
                                _ = fi166
                                _ = v167
                                _ = v168
                                _ = v169
                                _ = v175
                            }
                        }
                    }
                default:
                    v1 = ("hello" + ("\n\t" + (((v1 + "foo bar") + (((("\n\t" + (v1 + "")) + (v1 + v1)) + v1) + v1)) + (v1 + ("a" + "hello")))))
                }
                v3 += fi14
            } else {
                if ((!(42 != 32767)) || (-196 > 0)) {
                    v176 := GenS25{F0: 1}
                    switch -143 {
                    case -2147483648:
                        v1 = ""
                    case 990:
                        v3 += (v3 >> ((891) & 63))
                    case 32768:
                        if (534 > 831) {
                            if len(v2) > 0 {
                                v2[0] = 2
                            }
                            v1 = v1
                            var v177 int = v3
                            _ = v177
                        }
                    }
                    _ = v176
                }
                if len(v2) > 0 {
                    v2[0] = 42
                }
            }
            var v178 bool = (2147483648 <= 518)
            _ = fi14
            _ = v178
        }
    default:
        var v179 int = (v3 & (-(v3)))
        _ = v179
    }
    if true {
        var v180 float32 = (0.0 * (0.0 + 0.1))
        _ = v180
    } else {
        fmt.Println(v3, v3)
        switch (v3 >> ((65535) & 63)) {
        case 128:
            v181 := GenS27{F0: true, F1: 220, F2: ("a" + ("" + (v1 + (v1 + v1))))}
            _ = v181
        default:
            fmt.Println(v3, v3)
        }
        v3 = (v3 * v3)
    }
    v182 := GenS13{F0: 2147483648, F1: 42, F2: 1000}
    v1 = v1
    _ = len(v2)
    var v183 int16 = 1
    v184 := GenS17{F0: -9223372036854775808, F1: 127, F2: "a"}
    v2 = append(v2, 1)
    if len(v2) > 0 {
        v2[0] = 507
    }
    for fi185 := 0; fi185 < 16; fi185++ {
        fmt.Println(v1)
        var v186 int8 = 100
        _ = fi185
        _ = v186
    }
    if false {
        fmt.Println(v1, v183, v3)
        if true {
            v1 = v1
            if ((!(!true)) || true) {
                fmt.Println(v1)
                fmt.Println(v3, v1)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                _ = len(v2)
                for fi187 := 0; fi187 < 18; fi187++ {
                    if (42 < 2147483648) {
                        v188 := GenS14{F0: fi187, F1: 32767}
                        _ = v188
                    }
                    _ = len(v2)
                    switch (v1 + ("hello" + ("" + (v1 + ("hello" + v1))))) {
                    case "hello":
                        for fi189 := 0; fi189 < 3; fi189++ {
                            v190 := GenS28{F0: 2147483647, F1: 437}
                            var v191 uint = 754
                            _ = fi189
                            _ = v190
                            _ = v191
                        }
                    }
                    _ = fi187
                }
            }
        }
        if false {
            var v192 uint = 18446744073709551615
            _ = v192
        } else {
            v193 := GenS29{F0: 1, F1: -1}
            fmt.Println(v183)
            _ = v193
        }
    } else {
        if false {
            if false {
                v194 := GenS25{F0: 42}
                var v195 int32 = 0
                _ = v194
                _ = v195
            } else {
                v2 = append(v2, 0)
                switch ((v1 + ((v1 + "") + ("a" + ((((v1 + v1) + (((v1 + v1) + v1) + (v1 + v1))) + (v1 + "foo bar")) + v1)))) + "a") {
                case "foo bar":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "a":
                    switch (v3 | (v3 % ((32767) | 1))) {
                    case 32768:
                        var v196 float64 = -2.25
                        _ = v196
                    case -9223372036854775808:
                        switch ((v1 + v1) + "\n\t") {
                        case "\n\t":
                            var v197 int16 = (v183 << ((18446744073709551615) & 63))
                            _ = v197
                        case "hello":
                            if true {
                                v3 += (v3 >> ((42) & 63))
                                v183 = (v183 * v183)
                            } else {
                                for fi198 := 0; fi198 < 8; fi198++ {
                                    if true {
                                        v1 = "a"
                                    } else {
                                        v2 = append(v2, 2)
                                    }
                                    var v199 float32 = 1e10
                                    v199 = (v199 * float32(-2.25))
                                    _ = fi198
                                    _ = v199
                                }
                                v1 = (v1 + v1)
                                for fi200 := 0; fi200 < 5; fi200++ {
                                    if false {
                                        if false {
                                            v2 = append(v2, 989)
                                        }
                                        var v201 uint16 = 371
                                        _ = v201
                                    } else {
                                        fmt.Println(v3)
                                        v202 := GenS30{F0: 0, F1: 2, F2: 79}
                                        _ = v202
                                    }
                                    var v203 uint16 = 42
                                    if (1 > 382) {
                                        if ((!(v1 >= (v1 + v1))) || true) {
                                            if false {
                                                v2 = append(v2, 65535)
                                                fmt.Println(v203)
                                                if ((42 < 46) && true) {
                                                    if (2147483648 >= 461) {
                                                        for fi204 := 0; fi204 < 12; fi204++ {
                                                            var v205 string = v1
                                                            switch 2147483647 {
                                                            case 9223372036854775807:
                                                                fmt.Println(v1, v3, fi204)
                                                            case 42:
                                                                if false {
                                                                    var v206 bool = (false && false)
                                                                    _ = v206
                                                                } else {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 505
                                                                    }
                                                                    v207 := GenS31{F0: (v183 - (v183 + v183))}
                                                                    _ = v207
                                                                }
                                                            }
                                                            v3 -= v3
                                                            _ = fi204
                                                            _ = v205
                                                        }
                                                        if (1 > 4294967295) {
                                                            if len(v2) > 0 {
                                                                v2[0] = 0
                                                            }
                                                        } else {
                                                            v183 += v183
                                                            fmt.Println(v183)
                                                            if false {
                                                                v1 = v1
                                                                for fi208 := 0; fi208 < 16; fi208++ {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        v183 = (^(v183))
                                                                    }
                                                                    v209 := GenS7{F0: 210, F1: v3, F2: -9223372036854775808}
                                                                    var v210 uint8 = 0
                                                                    _ = fi208
                                                                    _ = v209
                                                                    _ = v210
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        v183 = (^(v183))
                                                        var v211 bool = false
                                                        _ = v211
                                                    }
                                                } else {
                                                    if true {
                                                        if (18446744073709551615 <= 42) {
                                                            v3 = v3
                                                            if (2 == 0) {
                                                                v3 = fi200
                                                            }
                                                        } else {
                                                            if len(v2) > 0 {
                                                                v2[0] = 879
                                                            }
                                                            var v212 bool = ((v3 % (((fi200 % ((v3) | 1))) | 1)) == (^(v3)))
                                                            _ = v212
                                                        }
                                                        v183 -= -1
                                                        switch (v3 & 4294967295) {
                                                        case 42:
                                                            var v213 bool = false
                                                            _ = v213
                                                        case -1:
                                                            for fi214 := 0; fi214 < 14; fi214++ {
                                                                if true {
                                                                    if ((-2.25 - -2.25) != 0.1) {
                                                                        v1 = ""
                                                                        var v215 float64 = 0.1
                                                                        v183 = v183
                                                                        _ = v215
                                                                    } else {
                                                                        v183 += -1
                                                                        if (66 > 0) {
                                                                            var v216 float32 = 0.0
                                                                            fmt.Println(fi200, fi200, v216)
                                                                            _ = v216
                                                                        }
                                                                        _ = len(v2)
                                                                        if false {
                                                                            var v217 uint64 = 4294967295
                                                                            _ = v217
                                                                        } else {
                                                                            v203 = v203
                                                                            if (!((true && ((float64(3.14159) + ((((1.5 * 1.5) + 0.1) - 1e10) / float64(0.1))) < 0.0)) || false)) {
                                                                                v1 = v1
                                                                                if true {
                                                                                    fmt.Println(v3, fi214)
                                                                                } else {
                                                                                    if true {
                                                                                        switch v3 {
                                                                                        case -430:
                                                                                            if true {
                                                                                                var v218 int32 = 318
                                                                                                if false {
                                                                                                    if false {
                                                                                                        var v219 float64 = (((float64(1e10) * ((3.14159 / float64((float32(-2.25) - 1e10))) * float64((0.0 * 3.14159)))) + -2.25) - 0.1)
                                                                                                        _ = len(v2)
                                                                                                        _ = v219
                                                                                                    }
                                                                                                    v220 := GenS8{F0: 32767}
                                                                                                    _ = v220
                                                                                                } else {
                                                                                                    if true {
                                                                                                        var v221 uint64 = 1
                                                                                                        fmt.Println(v221)
                                                                                                        _ = v221
                                                                                                    } else {
                                                                                                        v218 = v218
                                                                                                        switch v3 {
                                                                                                        case 0:
                                                                                                            v218 += v218
                                                                                                            fmt.Println(v3, v183, v3)
                                                                                                        default:
                                                                                                            v222 := GenS32{F0: v203, F1: 4294967295}
                                                                                                            _ = v222
                                                                                                        }
                                                                                                        fmt.Println(fi214)
                                                                                                    }
                                                                                                    var v223 bool = true
                                                                                                    switch (fi214 | (-(v3))) {
                                                                                                    case 2147483647:
                                                                                                        switch (("" + "hello") + "a") {
                                                                                                        case "\n\t":
                                                                                                            if v223 {
                                                                                                                if (1e10 == 1e10) {
                                                                                                                    if len(v2) > 0 {
                                                                                                                        v2[0] = 4294967295
                                                                                                                    }
                                                                                                                    var v224 int64 = 9223372036854775807
                                                                                                                    if (fi200 != -1) {
                                                                                                                        if len(v2) > 0 {
                                                                                                                            v2[0] = 2147483647
                                                                                                                        }
                                                                                                                        v203 = v203
                                                                                                                        switch (fi200 >> ((4294967295) & 63)) {
                                                                                                                        case 32768:
                                                                                                                            v224 = v224
                                                                                                                        case 32767:
                                                                                                                            v218 = 2147483647
                                                                                                                        default:
                                                                                                                            if true {
                                                                                                                                var v225 int8 = 83
                                                                                                                                var v226 uint16 = 2
                                                                                                                                _ = v225
                                                                                                                                _ = v226
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch int(v224) {
                                                                                                                    case -631:
                                                                                                                        v227 := GenS9{F0: 858, F1: int(v224), F2: v218}
                                                                                                                        _ = v227
                                                                                                                    case 9223372036854775807:
                                                                                                                        var v228 int32 = v218
                                                                                                                        _ = v228
                                                                                                                    default:
                                                                                                                        if true {
                                                                                                                            v229 := GenS14{F0: fi200, F1: v203}
                                                                                                                            if (true || false) {
                                                                                                                                fmt.Println(v218, v218)
                                                                                                                                v224 -= (v224 & 1)
                                                                                                                                fmt.Println(fi214, fi214, v224)
                                                                                                                            }
                                                                                                                            _ = v229
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if false {
                                                                                                                        v223 = false
                                                                                                                        v230 := GenS15{F0: true, F1: 1, F2: (v203 & 656)}
                                                                                                                        var v231 uint32 = 563
                                                                                                                        _ = v230
                                                                                                                        _ = v231
                                                                                                                    } else {
                                                                                                                        fmt.Println(v183)
                                                                                                                    }
                                                                                                                    _ = v224
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            v232 := GenS33{F0: 0, F1: -1, F2: -902}
                                                                                                            v1 = (v1 + "hello")
                                                                                                            _ = v232
                                                                                                        }
                                                                                                    case 9223372036854775807:
                                                                                                        fmt.Println(v1, v223)
                                                                                                    case -9223372036854775808:
                                                                                                        v218 = 1
                                                                                                    }
                                                                                                    _ = v223
                                                                                                }
                                                                                                v233 := GenS34{F0: (v183 * 1000), F1: (v218 ^ 42), F2: fi200}
                                                                                                if false {
                                                                                                    if (true || false) {
                                                                                                        var v234 int64 = 850
                                                                                                        if false {
                                                                                                            v235 := GenS35{F0: 255, F1: (-(v3))}
                                                                                                            _ = v235
                                                                                                        }
                                                                                                        var v236 int8 = -120
                                                                                                        _ = v234
                                                                                                        _ = v236
                                                                                                    }
                                                                                                }
                                                                                                var v237 int32 = v218
                                                                                                if false {
                                                                                                    v3 = v3
                                                                                                    v238 := GenS3{F0: 860}
                                                                                                    v2 = append(v2, 508)
                                                                                                    _ = v238
                                                                                                } else {
                                                                                                    v239 := GenS36{F0: 378}
                                                                                                    v237 = v237
                                                                                                    var v240 uint64 = uint64(v203)
                                                                                                    _ = v239
                                                                                                    _ = v240
                                                                                                }
                                                                                                _ = v218
                                                                                                _ = v233
                                                                                                _ = v237
                                                                                            } else {
                                                                                                v203 += v203
                                                                                                v3 -= v3
                                                                                                var v241 uint32 = 63
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                _ = v241
                                                                                            }
                                                                                        case 127:
                                                                                            fmt.Println(v203, v3)
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v3, v183)
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi214, fi214, v183)
                                                                            }
                                                                        }
                                                                        var v242 uint16 = 224
                                                                        _ = v242
                                                                    }
                                                                } else {
                                                                    fmt.Println(v203, v183, v183)
                                                                    switch v1 {
                                                                    case "a":
                                                                        var v243 uint = 655
                                                                        _ = v243
                                                                    case "\n\t":
                                                                        fmt.Println(v183, v183)
                                                                    case "":
                                                                        var v244 int8 = -1
                                                                        _ = v244
                                                                    }
                                                                }
                                                                var v245 uint16 = v203
                                                                if true {
                                                                    _ = len(v2)
                                                                    switch v1 {
                                                                    case "":
                                                                        if false {
                                                                            if true {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = 1
                                                                                }
                                                                                if false {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = 42
                                                                                    }
                                                                                    v183 += v183
                                                                                }
                                                                                v1 = v1
                                                                            } else {
                                                                                switch ("hello" + (v1 + ((v1 + v1) + "\n\t"))) {
                                                                                case "a":
                                                                                    fmt.Println(v183, fi200)
                                                                                case "foo bar":
                                                                                    v203 = v245
                                                                                }
                                                                                fmt.Println(v3, v1, v183)
                                                                                v183 = (v183 / ((v183) | 1))
                                                                                fmt.Println(v203)
                                                                                if ((fi200 * (^(fi200))) > fi200) {
                                                                                    fmt.Println(v3)
                                                                                    v183 -= v183
                                                                                    var v246 int64 = 9223372036854775807
                                                                                    if true {
                                                                                        if false {
                                                                                            v203 -= 255
                                                                                            if false {
                                                                                                var v247 uint32 = uint32(v203)
                                                                                                var v248 uint32 = (v247 >> ((4294967295) & 63))
                                                                                                var v249 uint = 718
                                                                                                _ = v247
                                                                                                _ = v248
                                                                                                _ = v249
                                                                                            }
                                                                                        }
                                                                                        v250 := GenS37{F0: 2147483647}
                                                                                        var v251 uint8 = 215
                                                                                        v252 := GenS33{F0: 42, F1: 104, F2: fi214}
                                                                                        _ = v250
                                                                                        _ = v251
                                                                                        _ = v252
                                                                                    } else {
                                                                                        fmt.Println(fi214, v245)
                                                                                        if (723 > 42) {
                                                                                            v203 -= v203
                                                                                            switch (fi200 - -2147483648) {
                                                                                            case -598:
                                                                                                v203 = v203
                                                                                            default:
                                                                                                v253 := GenS18{F0: 4294967295}
                                                                                                _ = v253
                                                                                            }
                                                                                            if false {
                                                                                                fmt.Println(fi214, v1, v1)
                                                                                            }
                                                                                            var v254 float64 = 1e10
                                                                                            var v255 string = v1
                                                                                            _ = v254
                                                                                            _ = v255
                                                                                        }
                                                                                    }
                                                                                    _ = v246
                                                                                }
                                                                            }
                                                                            var v256 uint8 = 1
                                                                            _ = v256
                                                                        } else {
                                                                            v1 = v1
                                                                            switch v1 {
                                                                            case "":
                                                                                v183 = v183
                                                                            default:
                                                                                if true {
                                                                                    var v257 int16 = (v183 / ((-32768) | 1))
                                                                                    _ = v257
                                                                                }
                                                                            }
                                                                        }
                                                                    case "foo bar":
                                                                        if false {
                                                                            v258 := GenS38{F0: 42}
                                                                            var v259 int8 = -1
                                                                            _ = v258
                                                                            _ = v259
                                                                        } else {
                                                                            var v260 bool = true
                                                                            if ((0.0 * 1e10) >= float64(1.5)) {
                                                                                _ = len(v2)
                                                                                if ((v203 <= v245) || v260) {
                                                                                    var v261 bool = v260
                                                                                    v262 := GenS39{F0: int64(v183), F1: 208}
                                                                                    v183 = v183
                                                                                    _ = v261
                                                                                    _ = v262
                                                                                }
                                                                                var v263 int64 = 936
                                                                                _ = v263
                                                                            }
                                                                            var v264 uint16 = v245
                                                                            v265 := GenS7{F0: 2, F1: 42, F2: 9223372036854775807}
                                                                            _ = v260
                                                                            _ = v264
                                                                            _ = v265
                                                                        }
                                                                    default:
                                                                        var v266 int = (fi214 ^ (fi200 - (-(fi214))))
                                                                        _ = v266
                                                                    }
                                                                    fmt.Println(v183)
                                                                } else {
                                                                    v267 := GenS24{F0: uint8(v245), F1: 255}
                                                                    v245 -= 65535
                                                                    var v268 uint8 = 255
                                                                    if false {
                                                                        v2 = append(v2, 4294967295)
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    } else {
                                                                        var v269 string = v1
                                                                        _ = v269
                                                                    }
                                                                    _ = v267
                                                                    _ = v268
                                                                }
                                                                var v270 int16 = v183
                                                                if false {
                                                                    _ = len(v2)
                                                                } else {
                                                                    fmt.Println(v1, fi214)
                                                                    v245 = 1
                                                                }
                                                                _ = fi214
                                                                _ = v245
                                                                _ = v270
                                                            }
                                                        }
                                                    } else {
                                                        var v271 string = (v1 + v1)
                                                        _ = v271
                                                    }
                                                }
                                            } else {
                                                var v272 string = v1
                                                _ = v272
                                            }
                                            for fi273 := 0; fi273 < 20; fi273++ {
                                                switch fi273 {
                                                case 9223372036854775807:
                                                    if (v203 == (v203 % (((v203 % ((0) | 1))) | 1))) {
                                                        if false {
                                                            v274 := GenS17{F0: 1, F1: -128, F2: v1}
                                                            fmt.Println(v1)
                                                            var v275 bool = (51 >= 4294967295)
                                                            v275 = false
                                                            _ = len(v2)
                                                            v203 += v203
                                                            switch v1 {
                                                            case "\n\t":
                                                                if true {
                                                                    v275 = false
                                                                } else {
                                                                    var v276 bool = v275
                                                                    v1 = v1
                                                                    if true {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    } else {
                                                                        v276 = v276
                                                                    }
                                                                    _ = v276
                                                                }
                                                            case "hello":
                                                                var v277 int32 = 42
                                                                _ = v277
                                                            case "foo bar":
                                                                if true {
                                                                    if (0 > 0) {
                                                                        if v275 {
                                                                            if v275 {
                                                                                if true {
                                                                                    v278 := GenS32{F0: v203, F1: 66}
                                                                                    _ = v278
                                                                                } else {
                                                                                    fmt.Println(fi200, v203, v183)
                                                                                    v275 = v275
                                                                                }
                                                                            }
                                                                            if len(v2) > 0 {
                                                                                v2[0] = 782
                                                                            }
                                                                            switch v1 {
                                                                            case "foo bar":
                                                                                fmt.Println(fi200, fi200)
                                                                            case "\n\t":
                                                                                fmt.Println(v275, v3, v3)
                                                                            case "hello":
                                                                                if (!(-116 >= 1)) {
                                                                                    v279 := GenS22{F0: 42, F1: int16(v3), F2: (v3 % ((fi200) | 1))}
                                                                                    _ = v279
                                                                                } else {
                                                                                    switch -187 {
                                                                                    case 1:
                                                                                        v280 := GenS6{F0: 327, F1: 4294967295, F2: fi200}
                                                                                        _ = v280
                                                                                    case -9223372036854775808:
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                    case 128:
                                                                                        switch ((("hello" + v1) + "\n\t") + (v1 + ((v1 + (v1 + "")) + ("a" + (("" + ((((v1 + v1) + v1) + (("a" + "a") + v1)) + ("" + v1))) + "a"))))) {
                                                                                        case "":
                                                                                            switch v1 {
                                                                                            case "\n\t":
                                                                                                if v275 {
                                                                                                    if v275 {
                                                                                                        var v281 bool = v275
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                        var v282 uint64 = 0
                                                                                                        if len(v2) > 0 {
                                                                                                            v2[0] = 65535
                                                                                                        }
                                                                                                        _ = v281
                                                                                                        _ = v282
                                                                                                    }
                                                                                                    v275 = (874 >= 1)
                                                                                                    if len(v2) >= 2 {
                                                                                                        v2 = v2[1:len(v2)]
                                                                                                    }
                                                                                                }
                                                                                            case "a":
                                                                                                var v283 int32 = 1
                                                                                                _ = v283
                                                                                            case "":
                                                                                                v284 := GenS23{F0: 1}
                                                                                                _ = v284
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v1)
                                                                                    switch ("hello" + "foo bar") {
                                                                                    case "\n\t":
                                                                                        if v275 {
                                                                                            var v285 uint8 = 128
                                                                                            if v275 {
                                                                                                if v275 {
                                                                                                    _ = len(v2)
                                                                                                    if len(v2) > 0 {
                                                                                                        v2[0] = uint32(v285)
                                                                                                    }
                                                                                                    if (!(((v285 <= 128) || (!(714 < 42))) && v275)) {
                                                                                                        switch (v1 + v1) {
                                                                                                        case "foo bar":
                                                                                                            v286 := GenS11{F0: v1}
                                                                                                            _ = v286
                                                                                                        default:
                                                                                                            v285 -= 128
                                                                                                        }
                                                                                                        if v275 {
                                                                                                            v287 := GenS22{F0: 1, F1: v183, F2: -779}
                                                                                                            v288 := GenS40{F0: (v3 | -9223372036854775808)}
                                                                                                            var v289 uint64 = 825
                                                                                                            if v275 {
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = 2147483648
                                                                                                                }
                                                                                                                var v290 uint64 = (v289 % (((v289 / (((v289 / (((^(v289))) | 1))) | 1))) | 1))
                                                                                                                v183 -= v183
                                                                                                                var v291 string = ((v1 + v1) + (v1 + v1))
                                                                                                                v292 := GenS18{F0: v290}
                                                                                                                _ = v290
                                                                                                                _ = v291
                                                                                                                _ = v292
                                                                                                            }
                                                                                                            _ = v287
                                                                                                            _ = v288
                                                                                                            _ = v289
                                                                                                        } else {
                                                                                                            v203 += 90
                                                                                                            v1 = v1
                                                                                                            if true {
                                                                                                                v293 := GenS33{F0: 0, F1: 42, F2: (fi200 + -617)}
                                                                                                                v285 += (^(v285))
                                                                                                                v3 -= -1
                                                                                                                _ = v293
                                                                                                            } else {
                                                                                                                if (-238 <= (fi200 * fi273)) {
                                                                                                                    if v275 {
                                                                                                                        if v275 {
                                                                                                                            if v275 {
                                                                                                                                v1 = v1
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v294 := GenS41{F0: "hello", F1: v183}
                                                                                                                            var v295 uint64 = 377
                                                                                                                            _ = v294
                                                                                                                            _ = v295
                                                                                                                        }
                                                                                                                        if true {
                                                                                                                            switch v1 {
                                                                                                                            case "a":
                                                                                                                                if ((float64(3.14159) - 0.1) == 1.5) {
                                                                                                                                    if false {
                                                                                                                                        v285 = (v285 & 2)
                                                                                                                                        if (-2147483648 <= (^(fi273))) {
                                                                                                                                            v296 := GenS32{F0: (v203 << ((255) & 63)), F1: 18446744073709551615}
                                                                                                                                            _ = v296
                                                                                                                                        } else {
                                                                                                                                            v2 = append(v2, 65535)
                                                                                                                                            v297 := GenS42{F0: -2147483648}
                                                                                                                                            _ = v297
                                                                                                                                        }
                                                                                                                                        if (true && (v275 || (v275 && (((-2.25 != 0.1) && (!((v285 - v285) != (v285 & v285)))) && true)))) {
                                                                                                                                            if (-1 < -672) {
                                                                                                                                                v298 := GenS43{F0: -1}
                                                                                                                                                _ = v298
                                                                                                                                            }
                                                                                                                                            v183 = v183
                                                                                                                                        } else {
                                                                                                                                            v1 = v1
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            case "foo bar":
                                                                                                                                if len(v2) >= 2 {
                                                                                                                                    v2 = v2[1:len(v2)]
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                fmt.Println(v285, v3, v285)
                                                                                                                            }
                                                                                                                            switch fi273 {
                                                                                                                            case 127:
                                                                                                                                var v299 int = (fi273 & 0)
                                                                                                                                _ = v299
                                                                                                                            case 32768:
                                                                                                                                switch v1 {
                                                                                                                                case "":
                                                                                                                                    var v300 uint32 = 42
                                                                                                                                    _ = v300
                                                                                                                                case "\n\t":
                                                                                                                                    v3 += v3
                                                                                                                                default:
                                                                                                                                    fmt.Println(v203, v275)
                                                                                                                                }
                                                                                                                            case 9223372036854775807:
                                                                                                                                var v301 uint = 320
                                                                                                                                _ = v301
                                                                                                                            default:
                                                                                                                                v183 += (v183 % ((v183) | 1))
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            if v275 {
                                                                                                                                if (-1 > 0) {
                                                                                                                                    if false {
                                                                                                                                        if v275 {
                                                                                                                                            var v302 string = "hello"
                                                                                                                                            if (100 != -1) {
                                                                                                                                                var v303 string = ("hello" + "a")
                                                                                                                                                v303 = (v302 + v302)
                                                                                                                                                v275 = (!v275)
                                                                                                                                                v304 := GenS43{F0: 1}
                                                                                                                                                if (v275 || (3.14159 != 1.5)) {
                                                                                                                                                    v2 = append(v2, 917)
                                                                                                                                                    if v275 {
                                                                                                                                                        v305 := GenS44{F0: 42}
                                                                                                                                                        switch v1 {
                                                                                                                                                        case "\n\t":
                                                                                                                                                            if len(v2) >= 2 {
                                                                                                                                                                v2 = v2[1:len(v2)]
                                                                                                                                                            }
                                                                                                                                                        case "a":
                                                                                                                                                            fmt.Println(v275, v303, v302)
                                                                                                                                                            v306 := GenS45{F0: (v285 * 1)}
                                                                                                                                                            _ = v306
                                                                                                                                                        }
                                                                                                                                                        _ = v305
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    v302 = v1
                                                                                                                                                }
                                                                                                                                                _ = v303
                                                                                                                                                _ = v304
                                                                                                                                            } else {
                                                                                                                                                fmt.Println(fi200, v285, v302)
                                                                                                                                                v307 := GenS16{F0: (v183 >> ((0) & 63))}
                                                                                                                                                v3 = fi200
                                                                                                                                                _ = v307
                                                                                                                                            }
                                                                                                                                            _ = v302
                                                                                                                                        } else {
                                                                                                                                            var v308 int32 = -1
                                                                                                                                            v203 += (v203 / (((v203 * (v203 - 1))) | 1))
                                                                                                                                            v309 := GenS25{F0: -54}
                                                                                                                                            _ = v308
                                                                                                                                            _ = v309
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        v3 += fi200
                                                                                                                                        var v310 int16 = 1000
                                                                                                                                        _ = v310
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    if v275 {
                                                                                                                                        v203 -= (v203 << ((360) & 63))
                                                                                                                                        var v311 float64 = (float64(((3.14159 - (1.5 * 0.0)) - 3.14159)) * 0.1)
                                                                                                                                        _ = v311
                                                                                                                                    }
                                                                                                                                    if true {
                                                                                                                                        var v312 string = v1
                                                                                                                                        var v313 int = fi273
                                                                                                                                        _ = v312
                                                                                                                                        _ = v313
                                                                                                                                    } else {
                                                                                                                                        var v314 string = "a"
                                                                                                                                        _ = v314
                                                                                                                                    }
                                                                                                                                    var v315 string = "\n\t"
                                                                                                                                    _ = v315
                                                                                                                                }
                                                                                                                                v275 = ((v285 / ((v285) | 1)) >= v285)
                                                                                                                                var v316 int32 = 1
                                                                                                                                _ = v316
                                                                                                                            }
                                                                                                                            v1 = v1
                                                                                                                            if len(v2) > 0 {
                                                                                                                                v2[0] = 1
                                                                                                                            }
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v317 uint64 = 42
                                                                                                                        _ = v317
                                                                                                                    }
                                                                                                                    v3 = 4294967295
                                                                                                                } else {
                                                                                                                    fmt.Println(fi200)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        var v318 int16 = v183
                                                                                                        _ = v318
                                                                                                    }
                                                                                                }
                                                                                                v319 := GenS46{F0: 0}
                                                                                                v1 = ("foo bar" + v1)
                                                                                                switch (("\n\t" + v1) + v1) {
                                                                                                case "hello":
                                                                                                    switch (v3 ^ v3) {
                                                                                                    case 0:
                                                                                                        switch 487 {
                                                                                                        case 39:
                                                                                                            var v320 int8 = -128
                                                                                                            _ = v320
                                                                                                        case 2147483647:
                                                                                                            var v321 uint64 = uint64(v203)
                                                                                                            _ = v321
                                                                                                        case 1:
                                                                                                            v2 = append(v2, 2147483648)
                                                                                                        }
                                                                                                    case -9223372036854775808:
                                                                                                        switch fi273 {
                                                                                                        case -339:
                                                                                                            v285 += (v285 & v285)
                                                                                                        }
                                                                                                    case 42:
                                                                                                        v3 -= 2147483647
                                                                                                    default:
                                                                                                        v322 := GenS6{F0: 643, F1: 42, F2: -1}
                                                                                                        _ = v322
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    var v323 uint8 = 183
                                                                                                    _ = v323
                                                                                                }
                                                                                                v324 := GenS15{F0: (479 == 4294967295), F1: v203, F2: (v203 >> ((42) & 63))}
                                                                                                _ = v319
                                                                                                _ = v324
                                                                                            }
                                                                                            var v325 uint = 2
                                                                                            _ = v285
                                                                                            _ = v325
                                                                                        } else {
                                                                                            v326 := GenS20{F0: 100, F1: (v203 % ((255) | 1)), F2: 32767}
                                                                                            v327 := GenS21{F0: (v203 - (v203 >> ((0) & 63))), F1: 127, F2: true}
                                                                                            v2 = append(v2, 1)
                                                                                            _ = v326
                                                                                            _ = v327
                                                                                        }
                                                                                    case "foo bar":
                                                                                        v183 = v183
                                                                                    case "":
                                                                                        v183 += (v183 + v183)
                                                                                    }
                                                                                    switch v3 {
                                                                                    case 32768:
                                                                                        var v328 float32 = 3.14159
                                                                                        _ = v328
                                                                                    }
                                                                                }
                                                                                switch (v3 | 32767) {
                                                                                case 32767:
                                                                                    v2 = append(v2, 0)
                                                                                default:
                                                                                    v203 += v203
                                                                                }
                                                                            default:
                                                                                var v329 uint64 = 0
                                                                                _ = v329
                                                                            }
                                                                            v203 = (v203 >> ((221) & 63))
                                                                            var v330 float64 = 1e10
                                                                            _ = v330
                                                                        } else {
                                                                            _ = len(v2)
                                                                            v331 := GenS5{F0: "\n\t"}
                                                                            _ = v331
                                                                        }
                                                                    } else {
                                                                        v3 += fi200
                                                                        if (42 < -38) {
                                                                            v332 := GenS38{F0: v183}
                                                                            v183 = v183
                                                                            _ = v332
                                                                        }
                                                                        v2 = append(v2, 2147483648)
                                                                        fmt.Println(fi273, v1)
                                                                        v203 += (v203 / (((v203 << ((255) & 63))) | 1))
                                                                    }
                                                                    var v333 int64 = 396
                                                                    var v334 uint32 = 42
                                                                    _ = v333
                                                                    _ = v334
                                                                } else {
                                                                    var v335 uint32 = 2147483648
                                                                    if v275 {
                                                                        if false {
                                                                            switch (v3 - (-(v3))) {
                                                                            case 1:
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = v335
                                                                                }
                                                                                if false {
                                                                                    switch fi200 {
                                                                                    case 32767:
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                    case 128:
                                                                                        if v275 {
                                                                                            var v336 uint16 = (v203 + v203)
                                                                                            _ = v336
                                                                                        } else {
                                                                                            var v337 uint64 = 780
                                                                                            var v338 int = (-(v3))
                                                                                            v183 = 1000
                                                                                            _ = v337
                                                                                            _ = v338
                                                                                        }
                                                                                    case -1:
                                                                                        var v339 float32 = float32(((0.1 / 1.0) - 3.14159))
                                                                                        v340 := GenS25{F0: -128}
                                                                                        _ = v339
                                                                                        _ = v340
                                                                                    }
                                                                                    _ = len(v2)
                                                                                } else {
                                                                                    var v341 float64 = 0.0
                                                                                    _ = v341
                                                                                }
                                                                            case 9223372036854775807:
                                                                                var v342 uint16 = (v203 << ((757) & 63))
                                                                                _ = v342
                                                                            case -255:
                                                                                v335 -= v335
                                                                            default:
                                                                                if (((v203 >> ((uint(v203)) & 63)) != (v203 / (((v203 ^ 521)) | 1))) && false) {
                                                                                    v3 = v3
                                                                                    switch (v1 + v1) {
                                                                                    case "":
                                                                                        v343 := GenS47{F0: v183, F1: v335, F2: 255}
                                                                                        _ = v343
                                                                                    }
                                                                                    fmt.Println(fi200, v335, fi273)
                                                                                    if false {
                                                                                        var v344 uint16 = 32767
                                                                                        switch v1 {
                                                                                        case "a":
                                                                                            switch v1 {
                                                                                            case "a":
                                                                                                var v345 float32 = 0.0
                                                                                                _ = v345
                                                                                            case "":
                                                                                                var v346 uint = 0
                                                                                                _ = v346
                                                                                            default:
                                                                                                var v347 int64 = -415
                                                                                                _ = v347
                                                                                            }
                                                                                        default:
                                                                                            if false {
                                                                                                if (false || v275) {
                                                                                                    if v275 {
                                                                                                        v3 += v3
                                                                                                        var v348 float64 = 3.14159
                                                                                                        if true {
                                                                                                            if len(v2) > 0 {
                                                                                                                v2[0] = v335
                                                                                                            }
                                                                                                            var v349 uint64 = 167
                                                                                                            switch fi200 {
                                                                                                            case 353:
                                                                                                                if v275 {
                                                                                                                    if len(v2) > 0 {
                                                                                                                        v2[0] = 4294967295
                                                                                                                    }
                                                                                                                    v350 := GenS20{F0: 1, F1: 42, F2: v344}
                                                                                                                    _ = v350
                                                                                                                } else {
                                                                                                                    var v351 uint = 2
                                                                                                                    fmt.Println(v348, v275)
                                                                                                                    if v275 {
                                                                                                                        fmt.Println(v275, v1, v335)
                                                                                                                        if v275 {
                                                                                                                            v2 = append(v2, v335)
                                                                                                                            v344 -= v344
                                                                                                                            v335 = v335
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v351
                                                                                                                }
                                                                                                            case 759:
                                                                                                                v349 -= (^(v349))
                                                                                                            case 32768:
                                                                                                                v352 := GenS48{F0: (v203 | v203)}
                                                                                                                _ = v352
                                                                                                            }
                                                                                                            fmt.Println(v344)
                                                                                                            _ = v349
                                                                                                        } else {
                                                                                                            v344 -= 65535
                                                                                                            if len(v2) > 0 {
                                                                                                                v2[0] = v335
                                                                                                            }
                                                                                                            var v353 int16 = -361
                                                                                                            var v354 uint32 = (v335 + v335)
                                                                                                            if (!(v275 || v275)) {
                                                                                                                v3 = int(v335)
                                                                                                                fmt.Println(v354, fi273)
                                                                                                            } else {
                                                                                                                v355 := GenS10{F0: 42}
                                                                                                                fmt.Println(v353, v354)
                                                                                                                _ = v355
                                                                                                            }
                                                                                                            _ = v353
                                                                                                            _ = v354
                                                                                                        }
                                                                                                        _ = v348
                                                                                                    } else {
                                                                                                        v356 := GenS25{F0: 116}
                                                                                                        _ = v356
                                                                                                    }
                                                                                                    switch int(v335) {
                                                                                                    case 32767:
                                                                                                        v357 := GenS49{F0: ("a" + v1), F1: -125, F2: 1}
                                                                                                        _ = v357
                                                                                                    case 1:
                                                                                                        if len(v2) > 0 {
                                                                                                            v2[0] = (v335 - (v335 ^ 2))
                                                                                                        }
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                var v358 float32 = float32((((((1e10 / 0.1) * 0.1) - float64(0.1)) - 3.14159) / 1.0))
                                                                                                if v275 {
                                                                                                    if v275 {
                                                                                                        v359 := GenS44{F0: 42}
                                                                                                        _ = v359
                                                                                                    } else {
                                                                                                        var v360 bool = (v275 || true)
                                                                                                        v361 := GenS10{F0: int32(fi200)}
                                                                                                        if (-32768 < (v183 | (v183 - (v183 >> ((18446744073709551615) & 63))))) {
                                                                                                            var v362 uint32 = 1
                                                                                                            switch "a" {
                                                                                                            case "foo bar":
                                                                                                                v363 := GenS50{F0: 1}
                                                                                                                v335 = (^(v335))
                                                                                                                if len(v2) >= 2 {
                                                                                                                    v2 = v2[1:len(v2)]
                                                                                                                }
                                                                                                                _ = v363
                                                                                                            case "hello":
                                                                                                                if v360 {
                                                                                                                    v362 = 0
                                                                                                                    if len(v2) > 0 {
                                                                                                                        v2[0] = v362
                                                                                                                    }
                                                                                                                    v335 = v362
                                                                                                                } else {
                                                                                                                    fmt.Println(v344, v275, v3)
                                                                                                                    var v364 uint16 = v203
                                                                                                                    var v365 int64 = 9223372036854775807
                                                                                                                    _ = v364
                                                                                                                    _ = v365
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = (v362 >> ((42) & 63))
                                                                                                                }
                                                                                                            }
                                                                                                            var v366 int64 = -865
                                                                                                            _ = v362
                                                                                                            _ = v366
                                                                                                        }
                                                                                                        v344 += v203
                                                                                                        switch (-(fi273)) {
                                                                                                        case -1:
                                                                                                            if ((v344 | 491) < v203) {
                                                                                                                if v275 {
                                                                                                                    v367 := GenS28{F0: 2147483648, F1: (^(v344))}
                                                                                                                    _ = v367
                                                                                                                }
                                                                                                                var v368 int8 = int8(v183)
                                                                                                                _ = v368
                                                                                                            } else {
                                                                                                                v360 = v360
                                                                                                                var v369 int = int(v183)
                                                                                                                _ = v369
                                                                                                            }
                                                                                                        default:
                                                                                                            _ = len(v2)
                                                                                                        }
                                                                                                        _ = v360
                                                                                                        _ = v361
                                                                                                    }
                                                                                                    var v370 uint16 = (v203 >> ((18446744073709551615) & 63))
                                                                                                    if v275 {
                                                                                                        v275 = (277 < 42)
                                                                                                        var v371 float64 = float64((v358 + ((1.5 + v358) / (1e10 + 0.1))))
                                                                                                        _ = v371
                                                                                                    }
                                                                                                    _ = v370
                                                                                                }
                                                                                                _ = v358
                                                                                            } else {
                                                                                                switch (^(fi273)) {
                                                                                                case 172:
                                                                                                    switch fi200 {
                                                                                                    case -403:
                                                                                                        v372 := GenS51{F0: v344, F1: 4294967295, F2: (v203 + v344)}
                                                                                                        _ = v372
                                                                                                    case 4294967295:
                                                                                                        if v275 {
                                                                                                            fmt.Println(fi273)
                                                                                                        }
                                                                                                    }
                                                                                                case -156:
                                                                                                    v373 := GenS52{F0: 280}
                                                                                                    _ = v373
                                                                                                default:
                                                                                                    switch (-(fi200)) {
                                                                                                    case 1:
                                                                                                        v335 = v335
                                                                                                    case 170:
                                                                                                        var v374 float32 = 3.14159
                                                                                                        _ = v374
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v183 = 1000
                                                                                        _ = v344
                                                                                    } else {
                                                                                        v375 := GenS33{F0: 2147483647, F1: -128, F2: fi200}
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        v376 := GenS15{F0: (v275 || true), F1: 331, F2: v203}
                                                                                        _ = v375
                                                                                        _ = v376
                                                                                    }
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = 2
                                                                                    }
                                                                                    var v377 uint64 = 675
                                                                                    var v378 string = v1
                                                                                    _ = v377
                                                                                    _ = v378
                                                                                } else {
                                                                                    fmt.Println(v183, v203)
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = (^(v335))
                                                                                    }
                                                                                    var v379 int = (fi200 * 127)
                                                                                    var v380 int32 = 42
                                                                                    _ = v379
                                                                                    _ = v380
                                                                                }
                                                                                var v381 uint32 = 1
                                                                                _ = v381
                                                                            }
                                                                            v1 = v1
                                                                            var v382 int8 = 0
                                                                            _ = v382
                                                                        }
                                                                        if (v335 > (v335 & 2147483647)) {
                                                                            v1 = v1
                                                                            switch (v3 | (fi273 >> ((42) & 63))) {
                                                                            case -266:
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = (v335 ^ (v335 << ((42) & 63)))
                                                                                }
                                                                            case -983:
                                                                                var v383 int8 = 127
                                                                                _ = v383
                                                                            default:
                                                                                v384 := GenS53{F0: -26, F1: 255, F2: v3}
                                                                                _ = v384
                                                                            }
                                                                            _ = len(v2)
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (^(v335))
                                                                            }
                                                                            v275 = true
                                                                        } else {
                                                                            fmt.Println(v335)
                                                                            switch (fi273 % (((fi273 - -1)) | 1)) {
                                                                            case -75:
                                                                                switch (v3 + fi273) {
                                                                                case -937:
                                                                                    if (v275 && ((v183 - v183) <= v183)) {
                                                                                        if false {
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            _ = len(v2)
                                                                                            switch v1 {
                                                                                            case "hello":
                                                                                                v183 = -32768
                                                                                            case "foo bar":
                                                                                                if v275 {
                                                                                                    fmt.Println(v203, v275)
                                                                                                } else {
                                                                                                    if len(v2) >= 2 {
                                                                                                        v2 = v2[1:len(v2)]
                                                                                                    }
                                                                                                    _ = len(v2)
                                                                                                    var v385 int64 = 42
                                                                                                    _ = v385
                                                                                                }
                                                                                            default:
                                                                                                v275 = (9 == 407)
                                                                                            }
                                                                                            fmt.Println(fi273, v183)
                                                                                        } else {
                                                                                            var v386 uint = 340
                                                                                            if v275 {
                                                                                                var v387 uint32 = (v335 & v335)
                                                                                                _ = v387
                                                                                            } else {
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                v388 := GenS54{F0: 1, F1: int(v335)}
                                                                                                _ = v388
                                                                                            }
                                                                                            fmt.Println(v275, v183)
                                                                                            _ = v386
                                                                                        }
                                                                                        v1 = v1
                                                                                    }
                                                                                default:
                                                                                    v389 := GenS45{F0: 72}
                                                                                    _ = v389
                                                                                }
                                                                            case -1:
                                                                                v390 := GenS43{F0: int8(v203)}
                                                                                _ = v390
                                                                            default:
                                                                                _ = len(v2)
                                                                            }
                                                                            var v391 uint8 = 42
                                                                            _ = v391
                                                                        }
                                                                    }
                                                                    v183 = (v183 / ((v183) | 1))
                                                                    fmt.Println(fi200, v203)
                                                                    _ = v335
                                                                }
                                                            }
                                                            _ = v274
                                                            _ = v275
                                                        } else {
                                                            var v392 int64 = -9223372036854775808
                                                            _ = v392
                                                        }
                                                        v1 = (v1 + (((((((("hello" + ("" + "hello")) + "a") + ((v1 + ("" + v1)) + v1)) + ((v1 + ((v1 + "hello") + v1)) + "foo bar")) + "hello") + v1) + (v1 + ((v1 + "foo bar") + "\n\t"))) + v1))
                                                        v3 -= fi200
                                                        var v393 uint32 = 65535
                                                        if len(v2) > 0 {
                                                            v2[0] = v393
                                                        }
                                                        var v394 int8 = 1
                                                        v1 = ""
                                                        var v395 int64 = 0
                                                        if false {
                                                            if (!false) {
                                                                v393 = v393
                                                            }
                                                        } else {
                                                            v396 := GenS52{F0: v395}
                                                            _ = v396
                                                        }
                                                        v183 -= v183
                                                        if (true && ((v3 != -9223372036854775808) && (!(247 < 42)))) {
                                                            v397 := GenS55{F0: (-(v394))}
                                                            var v398 int = (v3 + (fi200 % (((-(fi200))) | 1)))
                                                            switch (v3 >> ((0) & 63)) {
                                                            case 165:
                                                                switch v3 {
                                                                case -253:
                                                                    _ = len(v2)
                                                                    var v399 float64 = 3.14159
                                                                    _ = v399
                                                                case -9223372036854775808:
                                                                    fmt.Println(v183, fi200, v393)
                                                                default:
                                                                    var v400 uint64 = 617
                                                                    _ = v400
                                                                }
                                                                switch -902 {
                                                                case 32767:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = (v393 & v393)
                                                                    }
                                                                case 32768:
                                                                    if (!(v398 < fi200)) {
                                                                        var v401 uint32 = 0
                                                                        v393 = (v401 | v401)
                                                                        _ = v401
                                                                    } else {
                                                                        if ((!(true && (851 > fi273))) || (v394 < v394)) {
                                                                            v398 = fi200
                                                                            if (false && (true && false)) {
                                                                                v203 = 0
                                                                                var v402 uint8 = uint8(v393)
                                                                                _ = v402
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = len(v2)
                                                                case 2147483647:
                                                                    if true {
                                                                        v203 = v203
                                                                        v2 = append(v2, (v393 & (v393 % ((v393) | 1))))
                                                                        v403 := GenS56{F0: (v395 << ((712) & 63)), F1: 42}
                                                                        _ = v403
                                                                    } else {
                                                                        switch v1 {
                                                                        case "\n\t":
                                                                            v3 = (v398 << ((0) & 63))
                                                                        case "foo bar":
                                                                            if true {
                                                                                var v404 string = ""
                                                                                _ = v404
                                                                            }
                                                                        case "hello":
                                                                            var v405 int = v3
                                                                            _ = v405
                                                                        default:
                                                                            var v406 uint8 = 171
                                                                            _ = v406
                                                                        }
                                                                        v393 = v393
                                                                        if false {
                                                                            var v407 uint8 = 255
                                                                            var v408 uint16 = 8
                                                                            var v409 int16 = (v183 & v183)
                                                                            var v410 uint8 = v407
                                                                            _ = v407
                                                                            _ = v408
                                                                            _ = v409
                                                                            _ = v410
                                                                        }
                                                                        v411 := GenS17{F0: int64(v398), F1: v394, F2: v1}
                                                                        _ = v411
                                                                    }
                                                                }
                                                            }
                                                            _ = v397
                                                            _ = v398
                                                        } else {
                                                            fmt.Println(v1, v395)
                                                        }
                                                        _ = v393
                                                        _ = v394
                                                        _ = v395
                                                    } else {
                                                        _ = len(v2)
                                                    }
                                                case 127:
                                                    fmt.Println(fi200, fi273, v3)
                                                default:
                                                    v203 += (^(v203))
                                                }
                                                _ = fi273
                                            }
                                            var v412 int8 = -64
                                            _ = v412
                                        }
                                        fmt.Println(v3, v183, fi200)
                                        var v413 int16 = (v183 >> ((9) & 63))
                                        _ = v413
                                    }
                                    _ = fi200
                                    _ = v203
                                }
                            }
                        default:
                            if (255 == 65535) {
                                v1 = v1
                                switch (v3 - v3) {
                                case 32768:
                                    v2 = append(v2, 1)
                                default:
                                    v1 = (v1 + "")
                                }
                            } else {
                                switch (v3 << ((0) & 63)) {
                                case 42:
                                    v2 = append(v2, 42)
                                }
                            }
                        }
                    case 9223372036854775807:
                        v414 := GenS21{F0: 535, F1: 100, F2: (0.1 == 3.14159)}
                        _ = v414
                    }
                default:
                    v183 = v183
                }
            }
            v415 := GenS50{F0: 2}
            _ = v415
        } else {
            fmt.Println(v183)
            if true {
                if (1 != 294) {
                    v416 := GenS40{F0: 300}
                    if false {
                        var v417 bool = true
                        _ = v417
                    } else {
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                    }
                    _ = v416
                } else {
                    if (0 < 127) {
                        v3 = int(v183)
                        v3 = v3
                    }
                    var v418 int32 = 0
                    _ = v418
                }
                var v419 uint = 255
                var v420 int16 = v183
                _ = v419
                _ = v420
            } else {
                v421 := GenS32{F0: 2, F1: 42}
                _ = len(v2)
                _ = v421
            }
            v1 = v1
        }
        v3 += v3
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    }
    var v422 uint8 = 151
    var v423 bool = false
    if (127 != 127) {
        _ = len(v2)
        var v424 float64 = 1.5
        _ = v424
    } else {
        v1 = v1
        var v425 uint16 = 959
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = v425
    }
    var v426 uint8 = 1
    var v427 uint32 = 852
    fmt.Println("v1:", v1)
    fmt.Println("v183:", v183)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v422:", v422)
    fmt.Println("v423:", v423)
    fmt.Println("v426:", v426)
    fmt.Println("v427:", v427)
    _ = v1
    _ = v182
    _ = v183
    _ = v184
    _ = v2
    _ = v3
    _ = v422
    _ = v423
    _ = v426
    _ = v427
}
