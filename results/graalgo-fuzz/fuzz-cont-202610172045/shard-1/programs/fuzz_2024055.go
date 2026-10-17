package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 uint; F1 int32; F2 int32 }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 uint64; F1 int8; F2 int16 }
type GenS10 struct { F0 string; F1 int8; F2 int }
type GenS11 struct { F0 int16; F1 uint8 }
type GenS12 struct { F0 int16; F1 uint32 }
type GenS13 struct { F0 int32; F1 int32; F2 uint64 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 bool; F1 int64; F2 int8 }
type GenS16 struct { F0 int8; F1 bool; F2 int }
type GenS17 struct { F0 string; F1 uint; F2 uint16 }
type GenS18 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS19 struct { F0 uint32; F1 int64; F2 int16 }
type GenS20 struct { F0 int32; F1 uint; F2 uint64 }
type GenS21 struct { F0 int8; F1 int32 }
type GenS22 struct { F0 uint32; F1 int }
type GenS23 struct { F0 int; F1 bool }
func genFn0() uint16 {
    v7 := []int{272, 127, -340, 1}
    var v8 uint16 = 32767
    if false {
        for fi9 := 0; fi9 < 18; fi9++ {
            v8 += (v8 % ((v8) | 1))
            v8 += v8
            if (!false) {
                for fi10 := 0; fi10 < 4; fi10++ {
                    v7 = append(v7, int(v8))
                    _ = fi10
                }
            } else {
                var v11 bool = (42 != 2147483647)
                fmt.Println(v8)
                fmt.Println(fi9)
                _ = v11
            }
            _ = fi9
        }
        var v12 uint16 = (v8 >> ((921) & 63))
        _ = v12
    } else {
        var v13 uint64 = 4294967295
        for fi14 := 0; fi14 < 11; fi14++ {
            if false {
                var v15 bool = (true || (65 > -20))
                _ = v15
            } else {
                for fi16 := 0; fi16 < 17; fi16++ {
                    _ = len(v7)
                    v7 = append(v7, 9223372036854775807)
                    _ = fi16
                }
                var v17 int64 = -1
                _ = v17
            }
            _ = len(v7)
            _ = fi14
        }
        _ = v13
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            switch (fi18 << ((4294967295) & 63)) {
            case -9223372036854775808:
                var v20 uint8 = 182
                _ = v20
            case -2147483648:
                switch fi19 {
                case 4294967295:
                    if (!(!true)) {
                        switch fi18 {
                        case -9223372036854775808:
                            fmt.Println(fi18, fi18, fi18)
                        }
                        switch fi18 {
                        case 413:
                            var v21 float32 = float32((float64(1.5) / (1e10 + 3.14159)))
                            _ = v21
                        case -689:
                            var v22 int32 = -1
                            _ = v22
                        }
                        if len(v7) > 0 {
                            v7[0] = (fi18 / ((fi19) | 1))
                        }
                        var v23 int8 = 30
                        _ = v23
                    } else {
                        var v24 uint8 = 128
                        switch fi19 {
                        case 307:
                            if ((fi19 - fi18) < 32767) {
                                _ = len(v7)
                            } else {
                                if false {
                                    var v25 uint8 = 0
                                    switch -2147483648 {
                                    case 1:
                                        if (255 <= (^(v24))) {
                                            if (!(2 < v8)) {
                                                v25 += 139
                                                v25 = 42
                                                var v26 uint = uint(v8)
                                                _ = v26
                                            }
                                            if (!false) {
                                                switch ("" + ("hello" + "hello")) {
                                                case "\n\t":
                                                    v8 += (v8 >> ((551) & 63))
                                                case "":
                                                    switch fi19 {
                                                    case 9223372036854775807:
                                                        v8 = v8
                                                    case 766:
                                                        switch "" {
                                                        case "\n\t":
                                                            var v27 uint16 = 955
                                                            _ = v27
                                                        case "":
                                                            var v28 uint32 = uint32(v8)
                                                            _ = v28
                                                        case "a":
                                                            v29 := GenS1{F0: 2, F1: 232}
                                                            _ = v29
                                                        default:
                                                            v8 = uint16(v25)
                                                        }
                                                    }
                                                case "foo bar":
                                                    if false {
                                                        if (v24 != 0) {
                                                            var v30 int8 = 100
                                                            if ((v8 * (v8 << ((88) & 63))) != (v8 - 255)) {
                                                                if len(v7) >= 2 {
                                                                    v7 = v7[1:len(v7)]
                                                                }
                                                                v8 = v8
                                                            } else {
                                                                fmt.Println(fi19, v25, v25)
                                                                fmt.Println(v8, v25, fi19)
                                                                v7 = append(v7, 757)
                                                            }
                                                            var v31 uint16 = v8
                                                            switch fi19 {
                                                            case -1:
                                                                var v32 uint8 = (v25 + v25)
                                                                _ = v32
                                                            case -450:
                                                                fmt.Println(v30)
                                                            default:
                                                                var v33 string = "hello"
                                                                _ = v33
                                                            }
                                                            var v34 bool = true
                                                            _ = v30
                                                            _ = v31
                                                            _ = v34
                                                        } else {
                                                            if len(v7) >= 2 {
                                                                v7 = v7[1:len(v7)]
                                                            }
                                                        }
                                                        v24 = (v24 | v24)
                                                        v25 = (v24 * (v24 >> ((821) & 63)))
                                                    } else {
                                                        v8 += v8
                                                    }
                                                }
                                            }
                                            fmt.Println(v25, v25)
                                            if true {
                                                switch ((("" + ("hello" + "")) + "hello") + ("a" + ("" + ("\n\t" + ("hello" + "hello"))))) {
                                                case "a":
                                                    fmt.Println(v8)
                                                    if (255 == v8) {
                                                        v8 += (v8 + (^(v8)))
                                                        fmt.Println(fi18, v24, v8)
                                                    }
                                                    v35 := GenS1{F0: 4294967295, F1: 607}
                                                    _ = v35
                                                case "hello":
                                                    v25 = v24
                                                default:
                                                    fmt.Println(v8)
                                                }
                                            } else {
                                                var v36 uint16 = v8
                                                _ = v36
                                            }
                                            v8 += 1
                                            switch "foo bar" {
                                            case "\n\t":
                                                switch (fi19 >> ((255) & 63)) {
                                                case 8:
                                                    _ = len(v7)
                                                case 1:
                                                    v37 := GenS1{F0: 1, F1: 4294967295}
                                                    _ = v37
                                                default:
                                                    var v38 uint16 = uint16(v25)
                                                    v39 := GenS2{F0: true}
                                                    if false {
                                                        v8 += v8
                                                        v40 := GenS1{F0: 255, F1: 572}
                                                        _ = v40
                                                    } else {
                                                        switch 32767 {
                                                        case -1:
                                                            v41 := GenS3{F0: -71}
                                                            _ = v41
                                                        default:
                                                            switch (^(fi19)) {
                                                            case 336:
                                                                v24 += 0
                                                            case 2147483647:
                                                                var v42 uint8 = v25
                                                                _ = v42
                                                            case 128:
                                                                var v43 int = -609
                                                                _ = v43
                                                            default:
                                                                var v44 string = ("hello" + ("\n\t" + "a"))
                                                                _ = v44
                                                            }
                                                        }
                                                        v7 = append(v7, (^(fi18)))
                                                        v25 = v24
                                                    }
                                                    v8 = uint16(v24)
                                                    _ = v38
                                                    _ = v39
                                                }
                                            default:
                                                fmt.Println(v24, fi19, v25)
                                            }
                                            v8 = (v8 >> ((271) & 63))
                                            var v45 int16 = 1000
                                            var v46 float64 = -2.25
                                            var v47 int64 = 365
                                            var v48 int = (fi19 ^ (fi19 - 2147483647))
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        }
                                    case -601:
                                        fmt.Println(v8)
                                    default:
                                        if false {
                                            v7 = append(v7, -243)
                                            fmt.Println(v8)
                                            fmt.Println(v25, fi18)
                                        }
                                    }
                                    var v49 uint = 65535
                                    _ = v25
                                    _ = v49
                                }
                                v50 := GenS2{F0: true}
                                if ((true || (!(2147483647 <= fi19))) || true) {
                                    var v51 int32 = 1
                                    _ = v51
                                }
                                _ = v50
                            }
                            if (-1 < 9223372036854775807) {
                                v8 = uint16(v24)
                            }
                            switch ("foo bar" + (("foo bar" + "\n\t") + ("hello" + "a"))) {
                            case "":
                                if (float64(1.5) >= ((-2.25 - -2.25) - 0.1)) {
                                    v24 = v24
                                    if false {
                                        if false {
                                            v52 := GenS4{F0: 350, F1: true}
                                            _ = v52
                                        }
                                        v24 -= (v24 % (((^(v24))) | 1))
                                        var v53 int64 = 42
                                        if (true && (2 == 2147483648)) {
                                            if (true && true) {
                                                var v54 uint = 65535
                                                _ = v54
                                            }
                                            if true {
                                                if (400 == 591) {
                                                    var v55 int64 = v53
                                                    if (false || false) {
                                                        var v56 uint = 18446744073709551615
                                                        switch (fi18 - fi19) {
                                                        case -2147483648:
                                                            v56 -= v56
                                                            if ((^(v56)) < (^(v56))) {
                                                                var v57 string = "hello"
                                                                switch fi18 {
                                                                case -2147483648:
                                                                    var v58 string = "\n\t"
                                                                    _ = v58
                                                                case 622:
                                                                    v8 = uint16(v24)
                                                                default:
                                                                    if true {
                                                                        v59 := GenS5{F0: (v8 % ((v8) | 1)), F1: v57}
                                                                        v53 = v55
                                                                        fmt.Println(v24, v57, v24)
                                                                        switch -2147483648 {
                                                                        case 899:
                                                                            switch fi18 {
                                                                            case 127:
                                                                                v53 = (v55 << (((v56 | (v56 >> (((v56 - 0)) & 63)))) & 63))
                                                                            case 1:
                                                                                v8 = (v8 - v8)
                                                                            case 0:
                                                                                fmt.Println(v24)
                                                                            }
                                                                        case 2147483647:
                                                                            if len(v7) >= 2 {
                                                                                v7 = v7[1:len(v7)]
                                                                            }
                                                                            var v60 float64 = 1e10
                                                                            _ = v60
                                                                        case 128:
                                                                            v61 := GenS2{F0: (0.0 >= 1e10)}
                                                                            _ = v61
                                                                        }
                                                                        _ = v59
                                                                    }
                                                                }
                                                                v62 := GenS1{F0: 299, F1: 828}
                                                                _ = v57
                                                                _ = v62
                                                            } else {
                                                                if (3.14159 < ((1e10 + (3.14159 + (1.5 - -2.25))) - float32(1.5))) {
                                                                    fmt.Println(v55)
                                                                } else {
                                                                    switch fi19 {
                                                                    case -641:
                                                                        if false {
                                                                            var v63 int16 = 32767
                                                                            _ = len(v7)
                                                                            var v64 int32 = 2147483647
                                                                            _ = v63
                                                                            _ = v64
                                                                        }
                                                                    case 9223372036854775807:
                                                                        fmt.Println(v8, v55)
                                                                    case 0:
                                                                        switch fi18 {
                                                                        case -634:
                                                                            switch fi19 {
                                                                            case -541:
                                                                                v55 -= 1
                                                                            }
                                                                        default:
                                                                            v65 := GenS1{F0: v56, F1: v56}
                                                                            _ = v65
                                                                        }
                                                                    default:
                                                                        v56 -= (v56 - (v56 / (((^(v56))) | 1)))
                                                                    }
                                                                    var v66 uint = (v56 * (^(v56)))
                                                                    _ = v66
                                                                }
                                                            }
                                                            fmt.Println(v56)
                                                        case -21:
                                                            var v67 string = ("" + "\n\t")
                                                            _ = v67
                                                        }
                                                        var v68 uint8 = 124
                                                        _ = v56
                                                        _ = v68
                                                    }
                                                    _ = v55
                                                }
                                                if false {
                                                    switch (^(fi18)) {
                                                    case 0:
                                                        fmt.Println(fi19, v8, v8)
                                                    case 1:
                                                        switch "hello" {
                                                        case "hello":
                                                            v69 := GenS3{F0: -86}
                                                            _ = v69
                                                        case "foo bar":
                                                            v24 -= v24
                                                        case "":
                                                            fmt.Println(fi19, fi18)
                                                        }
                                                    }
                                                } else {
                                                    if (!true) {
                                                        var v70 bool = false
                                                        var v71 int8 = -28
                                                        fmt.Println(fi19, v24, fi18)
                                                        if v70 {
                                                            fmt.Println(v71, v24)
                                                            _ = len(v7)
                                                            var v72 int16 = 1000
                                                            _ = v72
                                                        }
                                                        var v73 bool = (473 > (v8 & v8))
                                                        _ = v70
                                                        _ = v71
                                                        _ = v73
                                                    } else {
                                                        fmt.Println(fi19)
                                                        if len(v7) > 0 {
                                                            v7[0] = (fi18 ^ fi18)
                                                        }
                                                        v7 = append(v7, fi18)
                                                    }
                                                    v74 := GenS6{F0: 255, F1: (v8 % (((v8 & (v8 / ((42) | 1)))) | 1))}
                                                    v53 -= v53
                                                    _ = v74
                                                }
                                                if len(v7) > 0 {
                                                    v7[0] = 128
                                                }
                                            }
                                            switch "\n\t" {
                                            case "\n\t":
                                                v75 := GenS2{F0: false}
                                                _ = v75
                                            default:
                                                if ((v8 << ((255) & 63)) >= 65535) {
                                                    v53 = (-(v53))
                                                } else {
                                                    fmt.Println(v53, fi19, fi18)
                                                }
                                            }
                                            var v76 int16 = -32768
                                            _ = v76
                                        }
                                        switch -658 {
                                        case -2147483648:
                                            v77 := GenS1{F0: 65535, F1: 42}
                                            _ = v77
                                        case -593:
                                            var v78 uint = 4294967295
                                            _ = v78
                                        default:
                                            switch (-(fi18)) {
                                            case 0:
                                                v8 += (v8 + (v8 / ((v8) | 1)))
                                            case -1:
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                            }
                                        }
                                        _ = v53
                                    } else {
                                        v24 += (v24 ^ 42)
                                        var v79 int32 = -2147483648
                                        _ = v79
                                    }
                                }
                            case "\n\t":
                                if false {
                                    var v80 uint16 = 608
                                    _ = v80
                                } else {
                                    if false {
                                        v81 := GenS4{F0: 18446744073709551615, F1: ((v24 >> ((65535) & 63)) != v24)}
                                        var v82 bool = (true || false)
                                        _ = v81
                                        _ = v82
                                    }
                                    switch "" {
                                    case "hello":
                                        v7 = append(v7, (-(fi19)))
                                    }
                                }
                            default:
                                if ((float32((1.5 + 1e10)) + 3.14159) < 1e10) {
                                    var v83 string = (("\n\t" + ("\n\t" + ("a" + ""))) + "a")
                                    if len(v7) > 0 {
                                        v7[0] = -9223372036854775808
                                    }
                                    switch fi18 {
                                    case 32768:
                                        if true {
                                            if false {
                                                var v84 uint = 42
                                                _ = v84
                                            } else {
                                                if (3.14159 != 0.0) {
                                                    if false {
                                                        var v85 int32 = 648
                                                        v7 = append(v7, 1)
                                                        if true {
                                                            if len(v7) >= 2 {
                                                                v7 = v7[1:len(v7)]
                                                            }
                                                            v7 = append(v7, (fi19 | (fi19 * (fi19 << ((717) & 63)))))
                                                            var v86 float64 = float64((0.1 * ((-2.25 / (-2.25 - (((1.5 + (0.0 / -2.25)) - 0.1) * ((3.14159 * 0.0) - ((3.14159 * -2.25) * -2.25))))) + 0.1)))
                                                            _ = v86
                                                        } else {
                                                            v85 -= (-(v85))
                                                        }
                                                        if (true && (!true)) {
                                                            fmt.Println(v85, fi19, v85)
                                                            if len(v7) > 0 {
                                                                v7[0] = fi18
                                                            }
                                                            v83 = v83
                                                            if ((1 < 0) || (-1 > 84)) {
                                                                if (130 >= 2) {
                                                                    if (2147483648 > 0) {
                                                                        fmt.Println(v24, fi18)
                                                                        if (!true) {
                                                                            v8 = v8
                                                                        } else {
                                                                            var v87 int64 = -1
                                                                            var v88 uint16 = v8
                                                                            var v89 int32 = v85
                                                                            if false {
                                                                                v90 := GenS7{F0: 255, F1: (v85 ^ v85), F2: (v85 | (v85 ^ -915))}
                                                                                _ = v90
                                                                            } else {
                                                                                v8 = 1
                                                                                switch (fi18 % ((-494) | 1)) {
                                                                                case -9223372036854775808:
                                                                                    switch fi18 {
                                                                                    case -9223372036854775808:
                                                                                        switch 571 {
                                                                                        case 2147483647:
                                                                                            if false {
                                                                                                switch (v83 + v83) {
                                                                                                case "a":
                                                                                                    v7 = append(v7, (fi19 ^ fi18))
                                                                                                case "hello":
                                                                                                    v8 = (^(v88))
                                                                                                case "foo bar":
                                                                                                    switch fi18 {
                                                                                                    case -9223372036854775808:
                                                                                                        v85 += (-(v89))
                                                                                                        var v91 int8 = 0
                                                                                                        _ = v91
                                                                                                    case 2147483647:
                                                                                                        if true {
                                                                                                            v24 += 171
                                                                                                            v24 = v24
                                                                                                            if (!(!((v85 - -2147483648) > -434))) {
                                                                                                                if (v83 > (v83 + v83)) {
                                                                                                                    v92 := GenS8{F0: 42, F1: 100}
                                                                                                                    v7 = append(v7, 1)
                                                                                                                    fmt.Println(v87)
                                                                                                                    v93 := GenS9{F0: 4294967295, F1: -128, F2: -428}
                                                                                                                    _ = v92
                                                                                                                    _ = v93
                                                                                                                }
                                                                                                                v94 := GenS2{F0: false}
                                                                                                                _ = v94
                                                                                                            }
                                                                                                        }
                                                                                                    case -100:
                                                                                                        v89 += (v89 % ((int32(v87)) | 1))
                                                                                                    }
                                                                                                }
                                                                                                if len(v7) >= 2 {
                                                                                                    v7 = v7[1:len(v7)]
                                                                                                }
                                                                                            }
                                                                                        case 127:
                                                                                            var v95 int32 = 42
                                                                                            _ = v95
                                                                                        default:
                                                                                            v83 = ("\n\t" + v83)
                                                                                        }
                                                                                        v87 -= v87
                                                                                    case 0:
                                                                                        if false {
                                                                                            fmt.Println(fi19)
                                                                                            fmt.Println(fi18, v87, fi19)
                                                                                        } else {
                                                                                            var v96 float32 = 3.14159
                                                                                            _ = v96
                                                                                        }
                                                                                    case 4294967295:
                                                                                        var v97 int = (-(fi18))
                                                                                        _ = v97
                                                                                    }
                                                                                case -1:
                                                                                    v98 := GenS4{F0: 90, F1: false}
                                                                                    _ = v98
                                                                                }
                                                                                if false {
                                                                                    fmt.Println(v87, fi19, v24)
                                                                                    v24 = 127
                                                                                    if true {
                                                                                        v88 -= v88
                                                                                        v99 := GenS3{F0: 42}
                                                                                        var v100 string = ("" + "a")
                                                                                        _ = v100
                                                                                        _ = v99
                                                                                    } else {
                                                                                        if ((v24 ^ v24) != v24) {
                                                                                            var v101 float64 = 3.14159
                                                                                            fmt.Println(v85)
                                                                                            var v102 uint = 18446744073709551615
                                                                                            _ = v101
                                                                                            _ = v102
                                                                                        } else {
                                                                                            v103 := GenS10{F0: ((((v83 + "a") + (v83 + ("\n\t" + v83))) + v83) + ("a" + v83)), F1: int8(v88), F2: fi19}
                                                                                            var v104 uint64 = 9
                                                                                            if len(v7) >= 2 {
                                                                                                v7 = v7[1:len(v7)]
                                                                                            }
                                                                                            v105 := GenS5{F0: 42, F1: v83}
                                                                                            fmt.Println(v89)
                                                                                            _ = v103
                                                                                            _ = v104
                                                                                            _ = v105
                                                                                        }
                                                                                        var v106 int8 = 1
                                                                                        var v107 int8 = -128
                                                                                        _ = v106
                                                                                        _ = v107
                                                                                    }
                                                                                } else {
                                                                                    v83 = v83
                                                                                    v24 -= 128
                                                                                    v108 := GenS11{F0: 788, F1: (^(v24))}
                                                                                    v85 = 102
                                                                                    if false {
                                                                                        v109 := GenS2{F0: (!(v87 == (v87 % (((v87 >> ((956) & 63))) | 1))))}
                                                                                        _ = len(v7)
                                                                                        v110 := GenS10{F0: ("foo bar" + v83), F1: 1, F2: (-(fi18))}
                                                                                        _ = v109
                                                                                        _ = v110
                                                                                    }
                                                                                    if (false || (0.1 > 0.1)) {
                                                                                        fmt.Println(v85)
                                                                                        var v111 uint16 = (v88 >> ((0) & 63))
                                                                                        _ = v111
                                                                                    } else {
                                                                                        var v112 uint = 255
                                                                                        _ = v112
                                                                                    }
                                                                                    _ = v108
                                                                                }
                                                                            }
                                                                            var v113 uint = 104
                                                                            _ = v113
                                                                            _ = v87
                                                                            _ = v88
                                                                            _ = v89
                                                                        }
                                                                        if true {
                                                                            v85 = 522
                                                                            var v114 uint64 = 18446744073709551615
                                                                            if (1 < 587) {
                                                                                if (!true) {
                                                                                    v24 = (^(v24))
                                                                                    var v115 float32 = (1.5 / float32(1e10))
                                                                                    _ = v115
                                                                                } else {
                                                                                    if true {
                                                                                        switch ("foo bar" + "hello") {
                                                                                        case "":
                                                                                            var v116 int = fi19
                                                                                            v85 += -1
                                                                                            _ = v116
                                                                                        case "a":
                                                                                            var v117 uint64 = 0
                                                                                            _ = v117
                                                                                        }
                                                                                        var v118 float32 = 1e10
                                                                                        _ = v118
                                                                                    } else {
                                                                                        v119 := GenS12{F0: 42, F1: 648}
                                                                                        switch -124 {
                                                                                        case -825:
                                                                                            _ = len(v7)
                                                                                        case 641:
                                                                                            var v120 int8 = -128
                                                                                            _ = v120
                                                                                        }
                                                                                        _ = v119
                                                                                    }
                                                                                    v83 = v83
                                                                                    if true {
                                                                                        var v121 uint64 = (v114 | (v114 ^ (^(v114))))
                                                                                        var v122 int8 = -1
                                                                                        switch -694 {
                                                                                        case -2147483648:
                                                                                            fmt.Println(fi19, fi19, v121)
                                                                                        case 4294967295:
                                                                                            v123 := GenS13{F0: 1, F1: v85, F2: 757}
                                                                                            _ = v123
                                                                                        case 128:
                                                                                            switch ("" + "foo bar") {
                                                                                            case "\n\t":
                                                                                                v83 = v83
                                                                                            }
                                                                                        default:
                                                                                            switch v83 {
                                                                                            case "hello":
                                                                                                fmt.Println(v85)
                                                                                            case "foo bar":
                                                                                                v114 = v114
                                                                                                if false {
                                                                                                    var v124 uint = 60
                                                                                                    _ = len(v7)
                                                                                                    v125 := GenS14{F0: -2147483648}
                                                                                                    _ = v124
                                                                                                    _ = v125
                                                                                                } else {
                                                                                                    v8 = v8
                                                                                                    v126 := GenS12{F0: -529, F1: uint32(v114)}
                                                                                                    _ = v126
                                                                                                }
                                                                                                if true {
                                                                                                    if (((v24 & 127) >= v24) || (1 < v121)) {
                                                                                                        if false {
                                                                                                            if ((("" + "hello") >= v83) || ((!(42 == 2147483648)) && (!(!("\n\t" == "foo bar"))))) {
                                                                                                                if len(v7) > 0 {
                                                                                                                    v7[0] = fi18
                                                                                                                }
                                                                                                                var v127 int8 = (v122 << ((42) & 63))
                                                                                                                v85 = v85
                                                                                                                _ = v127
                                                                                                            } else {
                                                                                                                if (!false) {
                                                                                                                    if (!(18446744073709551615 > 42)) {
                                                                                                                        v128 := GenS1{F0: uint(v121), F1: uint(v24)}
                                                                                                                        if false {
                                                                                                                            v129 := GenS7{F0: 1, F1: v85, F2: v85}
                                                                                                                            var v130 bool = ((float64(0.0) / float64(1.5)) == 1e10)
                                                                                                                            var v131 string = ""
                                                                                                                            _ = v129
                                                                                                                            _ = v130
                                                                                                                            _ = v131
                                                                                                                        } else {
                                                                                                                            v132 := GenS15{F0: true, F1: 818, F2: v122}
                                                                                                                            v24 = 2
                                                                                                                            _ = v132
                                                                                                                        }
                                                                                                                        _ = v128
                                                                                                                    } else {
                                                                                                                        switch v83 {
                                                                                                                        case "hello":
                                                                                                                            if true {
                                                                                                                                fmt.Println(v8, fi18, v122)
                                                                                                                                switch v83 {
                                                                                                                                case "\n\t":
                                                                                                                                    if (-9223372036854775808 < 9223372036854775807) {
                                                                                                                                        if true {
                                                                                                                                            v8 += (v8 % ((v8) | 1))
                                                                                                                                            v7 = append(v7, fi19)
                                                                                                                                            var v133 int8 = (^(v122))
                                                                                                                                            _ = v133
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                case "foo bar":
                                                                                                                                    v24 += (v24 * 75)
                                                                                                                                default:
                                                                                                                                    v134 := GenS11{F0: 32767, F1: 249}
                                                                                                                                    _ = v134
                                                                                                                                }
                                                                                                                                var v135 string = ((("foo bar" + ((v83 + v83) + v83)) + v83) + v83)
                                                                                                                                _ = v135
                                                                                                                            } else {
                                                                                                                                var v136 uint64 = (v114 << ((4294967295) & 63))
                                                                                                                                v137 := GenS5{F0: 296, F1: ("a" + (v83 + v83))}
                                                                                                                                _ = v136
                                                                                                                                _ = v137
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            v138 := GenS4{F0: 79, F1: (1.5 >= (0.1 / -2.25))}
                                                                                                                            _ = v138
                                                                                                                        }
                                                                                                                        fmt.Println(v8)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v7 = append(v7, fi19)
                                                                                                                    v114 = (^(v121))
                                                                                                                    v8 -= v8
                                                                                                                    v139 := GenS4{F0: v114, F1: true}
                                                                                                                    _ = v139
                                                                                                                }
                                                                                                            }
                                                                                                            v140 := GenS2{F0: false}
                                                                                                            _ = v140
                                                                                                        } else {
                                                                                                            if len(v7) >= 2 {
                                                                                                                v7 = v7[1:len(v7)]
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v83, fi18, fi18)
                                                                                                } else {
                                                                                                    fmt.Println(v114, v114, fi19)
                                                                                                }
                                                                                            case "a":
                                                                                                v114 -= (v121 * v114)
                                                                                            }
                                                                                        }
                                                                                        _ = v121
                                                                                        _ = v122
                                                                                    } else {
                                                                                        if (!(v83 != "hello")) {
                                                                                            v141 := GenS13{F0: (v85 >> ((598) & 63)), F1: v85, F2: uint64(v24)}
                                                                                            _ = v141
                                                                                        } else {
                                                                                            var v142 float64 = -2.25
                                                                                            if false {
                                                                                                v114 = (v114 / (((v114 % ((2) | 1))) | 1))
                                                                                                if len(v7) > 0 {
                                                                                                    v7[0] = (fi18 / (((^(fi18))) | 1))
                                                                                                }
                                                                                            } else {
                                                                                                v24 = v24
                                                                                                if true {
                                                                                                    switch v83 {
                                                                                                    case "\n\t":
                                                                                                        var v143 string = v83
                                                                                                        _ = v143
                                                                                                    case "hello":
                                                                                                        v144 := GenS13{F0: v85, F1: 1, F2: (v114 + v114)}
                                                                                                        _ = v144
                                                                                                    case "foo bar":
                                                                                                        var v145 int = (-(fi19))
                                                                                                        _ = v145
                                                                                                    }
                                                                                                } else {
                                                                                                    fmt.Println(v24)
                                                                                                }
                                                                                                if (4294967295 < 42) {
                                                                                                    v146 := GenS16{F0: -106, F1: true, F2: (fi19 << ((2) & 63))}
                                                                                                    _ = v146
                                                                                                } else {
                                                                                                    v114 -= 2
                                                                                                    if true {
                                                                                                        _ = len(v7)
                                                                                                    } else {
                                                                                                        var v147 int32 = 0
                                                                                                        v8 = v8
                                                                                                        v85 = -1
                                                                                                        var v148 uint8 = v24
                                                                                                        _ = v147
                                                                                                        _ = v148
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            switch (v83 + v83) {
                                                                                            case "hello":
                                                                                                fmt.Println(v142)
                                                                                            default:
                                                                                                var v149 string = v83
                                                                                                _ = v149
                                                                                            }
                                                                                            _ = v142
                                                                                        }
                                                                                        v85 += v85
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                if (true && (42 < 21)) {
                                                                                    var v150 float32 = -2.25
                                                                                    if (!(!(((-852 > (-(v85))) && false) || true))) {
                                                                                        fmt.Println(fi18, v24)
                                                                                        var v151 int = (-(fi18))
                                                                                        _ = v151
                                                                                    }
                                                                                    _ = v150
                                                                                }
                                                                                var v152 uint16 = (^(v8))
                                                                                v24 = v24
                                                                                fmt.Println(v85, v152)
                                                                                fmt.Println(fi18, v83)
                                                                                var v153 string = ""
                                                                                v154 := GenS12{F0: -61, F1: 2147483648}
                                                                                switch v83 {
                                                                                case "a":
                                                                                    v155 := GenS17{F0: (((v83 + (v153 + ("a" + (v153 + "a")))) + v153) + v83), F1: 1, F2: v8}
                                                                                    _ = v155
                                                                                case "foo bar":
                                                                                    v7 = append(v7, int(v24))
                                                                                default:
                                                                                    v156 := GenS15{F0: true, F1: -1, F2: 81}
                                                                                    _ = v156
                                                                                }
                                                                                _ = v152
                                                                                _ = v153
                                                                                _ = v154
                                                                            }
                                                                            _ = v114
                                                                        } else {
                                                                            switch (fi19 ^ fi19) {
                                                                            case -446:
                                                                                switch v83 {
                                                                                case "hello":
                                                                                    fmt.Println(v85, v8)
                                                                                    v8 -= (v8 << ((4294967295) & 63))
                                                                                    if len(v7) > 0 {
                                                                                        v7[0] = (fi18 % (((^(fi18))) | 1))
                                                                                    }
                                                                                case "foo bar":
                                                                                    var v157 int16 = -735
                                                                                    _ = v157
                                                                                default:
                                                                                    if len(v7) > 0 {
                                                                                        v7[0] = (fi19 << ((42) & 63))
                                                                                    }
                                                                                    var v158 uint64 = 0
                                                                                    v158 = 18
                                                                                    _ = v158
                                                                                }
                                                                            case 804:
                                                                                v24 += v24
                                                                            default:
                                                                                fmt.Println(fi18)
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v159 int64 = -1
                                                                        var v160 int32 = v85
                                                                        var v161 int32 = 0
                                                                        _ = v159
                                                                        _ = v160
                                                                        _ = v161
                                                                    }
                                                                    fmt.Println(v83, v83)
                                                                }
                                                                var v162 int = -2147483648
                                                                _ = v162
                                                            }
                                                        }
                                                        v163 := GenS18{F0: -1, F1: 2, F2: 1}
                                                        _ = len(v7)
                                                        _ = v163
                                                        _ = v85
                                                    } else {
                                                        v8 += (v8 & (^(v8)))
                                                        v24 += uint8(v8)
                                                        _ = len(v7)
                                                    }
                                                    v83 = v83
                                                }
                                                fmt.Println(fi18, v8)
                                            }
                                            var v164 uint8 = v24
                                            switch fi19 {
                                            case 1:
                                                var v165 uint32 = 322
                                                _ = v165
                                            case -14:
                                                v7 = append(v7, 128)
                                            case 0:
                                                v7 = append(v7, (fi19 | -1))
                                            default:
                                                v8 += (v8 & v8)
                                            }
                                            var v166 uint = 4294967295
                                            _ = v164
                                            _ = v166
                                        }
                                    }
                                    _ = v83
                                }
                            }
                            var v167 uint = 4294967295
                            _ = v167
                        case 4294967295:
                            if ((!false) && (true || (18446744073709551615 < 4294967295))) {
                                var v168 uint32 = 1
                                _ = v168
                            }
                        }
                        _ = v24
                    }
                default:
                    if true {
                        if (!true) {
                            var v169 uint8 = 1
                            if (2147483647 <= 42) {
                                fmt.Println(v8, fi19)
                                fmt.Println(fi19, v8)
                            } else {
                                var v170 uint32 = 2
                                v171 := GenS13{F0: 42, F1: -90, F2: 1}
                                _ = v170
                                _ = v171
                            }
                            if false {
                                if false {
                                    v8 = v8
                                    var v172 int16 = -817
                                    _ = v172
                                } else {
                                    var v173 int = (fi19 | (^(fi18)))
                                    _ = v173
                                }
                                var v174 float32 = ((0.0 * float32(1.5)) + 1e10)
                                v174 -= 0.1
                                var v175 bool = false
                                _ = len(v7)
                                _ = v174
                                _ = v175
                            }
                            switch fi18 {
                            case 42:
                                var v176 int8 = -26
                                _ = v176
                            case 128:
                                var v177 float64 = float64(0.1)
                                _ = v177
                            case 32767:
                                if (!false) {
                                    v8 += 0
                                    var v178 uint16 = v8
                                    _ = v178
                                }
                            default:
                                v8 = 410
                            }
                            _ = v169
                        }
                        switch (fi18 << ((454) & 63)) {
                        case 2147483647:
                            if len(v7) > 0 {
                                v7[0] = fi18
                            }
                        default:
                            v179 := GenS17{F0: ("" + "hello"), F1: 255, F2: v8}
                            _ = v179
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    } else {
                        v8 = (v8 % ((v8) | 1))
                        var v180 uint8 = 125
                        fmt.Println(fi19, fi19)
                        _ = v180
                    }
                }
                switch (("foo bar" + "") + ("" + ("hello" + "a"))) {
                case "a":
                    if true {
                        switch ("foo bar" + (("\n\t" + "a") + ("\n\t" + ((("" + "a") + "") + ("\n\t" + ("hello" + "hello")))))) {
                        case "foo bar":
                            v8 += v8
                        default:
                            v181 := GenS3{F0: -56}
                            _ = v181
                        }
                        v8 -= v8
                        var v182 int64 = -1
                        v183 := GenS10{F0: "", F1: -1, F2: (fi19 + fi18)}
                        _ = v182
                        _ = v183
                    }
                case "":
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
                var v184 uint64 = 42
                _ = v184
            }
            v7 = append(v7, fi18)
            v185 := GenS10{F0: "hello", F1: 1, F2: fi19}
            _ = fi19
            _ = v185
        }
        for fi186 := 0; fi186 < 3; fi186++ {
            var v187 uint64 = 18446744073709551615
            if ((2147483648 == 34) && (!(false && false))) {
                if (0 != -63) {
                    if true {
                        _ = len(v7)
                        v187 = v187
                        v187 = (v187 & v187)
                    } else {
                        v188 := GenS11{F0: 0, F1: 255}
                        _ = v188
                    }
                } else {
                    v187 -= v187
                    v189 := GenS19{F0: 2147483648, F1: 0, F2: -1}
                    _ = v189
                }
                fmt.Println(fi18, v187)
                v8 = (^(v8))
            }
            _ = fi186
            _ = v187
        }
        _ = fi18
    }
    for fi190 := 0; fi190 < 20; fi190++ {
        if len(v7) > 0 {
            v7[0] = fi190
        }
        v191 := GenS10{F0: (("\n\t" + "") + "foo bar"), F1: 0, F2: (fi190 / ((fi190) | 1))}
        for fi192 := 0; fi192 < 20; fi192++ {
            v8 -= v8
            var v193 string = ("a" + "hello")
            if ((fi192 & fi192) < (fi190 >> ((1) & 63))) {
                switch ("" + "hello") {
                case "a":
                    var v194 int32 = 2147483647
                    _ = v194
                case "hello":
                    switch v193 {
                    case "hello":
                        v193 = (("\n\t" + v193) + "foo bar")
                    case "foo bar":
                        var v195 bool = false
                        _ = v195
                    case "\n\t":
                        v8 += (v8 % (((v8 - v8)) | 1))
                    }
                default:
                    v193 = "hello"
                    v8 = (^(v8))
                    var v196 int = (fi190 << ((255) & 63))
                    var v197 uint8 = 0
                    _ = v196
                    _ = v197
                }
                if (42 != 1) {
                    var v198 int = -736
                    var v199 uint64 = 599
                    _ = v198
                    _ = v199
                } else {
                    fmt.Println(fi192, v8)
                }
                var v200 int32 = 2147483647
                if len(v7) > 0 {
                    v7[0] = (fi190 | (fi192 | (fi192 >> ((255) & 63))))
                }
                _ = v200
            }
            _ = fi192
            _ = v193
        }
        _ = fi190
        _ = v191
    }
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{4294967295, 2147483647, 370}
    v3 := make(map[int]int)
    v4 := []uint32{0, 2147483647, 816, 2147483647}
    var v5 bool = true
    v6 := make(map[int]int)
    fmt.Println(v5)
    var v201 int = 128
    v202 := GenS17{F0: ("\n\t" + "foo bar"), F1: 488, F2: 80}
    v203 := v6[(v201 + 417)]
    v201 -= v201
    if (32768 == v203) {
        for fi204 := 0; fi204 < 15; fi204++ {
            var v205 int64 = -9223372036854775808
            v201 = (v203 << ((255) & 63))
            _ = fi204
            _ = v205
        }
    }
    _ = len(v3)
    var v206 int8 = 127
    var v207 int16 = 42
    for fi208 := 0; fi208 < 20; fi208++ {
        v209 := GenS20{F0: -380, F1: 42, F2: 628}
        fmt.Println(fi208, v207)
        _ = fi208
        _ = v209
    }
    var v210 int16 = v207
    for fi211 := 0; fi211 < 14; fi211++ {
        switch fi211 {
        case 288:
            v203 -= 11
        case 9223372036854775807:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        for fi212 := 0; fi212 < 14; fi212++ {
            v213 := GenS21{F0: (v206 | v206), F1: 749}
            v214 := GenS22{F0: 422, F1: v203}
            var v215 string = ("hello" + ("a" + "\n\t"))
            _ = len(v4)
            _ = fi212
            _ = v213
            _ = v214
            _ = v215
        }
        _ = fi211
    }
    var v216 float32 = float32((float64(3.14159) + float64((3.14159 / 1.5))))
    var v217 int64 = 0
    fmt.Println(v5)
    v201 = v203
    var v218 int32 = -2147483648
    for fi219 := 0; fi219 < 15; fi219++ {
        v210 += 31
        for fi220 := 0; fi220 < 19; fi220++ {
            v221 := GenS12{F0: v210, F1: 2}
            if (255 < 128) {
                v4 = append(v4, 65535)
                if len(v2) > 0 {
                    v2[0] = 42
                }
            } else {
                switch ("hello" + "") {
                case "foo bar":
                    v222 := GenS5{F0: 42, F1: "\n\t"}
                    _ = v222
                }
            }
            if (!v5) {
                if false {
                    v217 = 9223372036854775807
                    _ = len(v4)
                    var v223 int16 = v207
                    _ = v223
                } else {
                    var v224 float32 = ((1e10 * 0.0) - v216)
                    _ = v224
                }
                v201 -= 290
                if v5 {
                    _ = len(v4)
                } else {
                    v6[1] = 9223372036854775807
                    var v225 int16 = (-(v210))
                    _ = v225
                }
            }
            _ = fi220
            _ = v221
        }
        v226 := GenS23{F0: (fi219 << ((0) & 63)), F1: (0 <= 128)}
        _ = fi219
        _ = v226
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v201:", v201)
    fmt.Println("v203:", v203)
    fmt.Println("v206:", v206)
    fmt.Println("v207:", v207)
    fmt.Println("v210:", v210)
    fmt.Println("v216:", v216)
    fmt.Println("v217:", v217)
    fmt.Println("v218:", v218)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v201
    _ = v202
    _ = v203
    _ = v206
    _ = v207
    _ = v210
    _ = v216
    _ = v217
    _ = v218
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
