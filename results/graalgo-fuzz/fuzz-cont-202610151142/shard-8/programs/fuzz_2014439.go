package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 int8; F1 bool; F2 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint64; F1 int }
type GenS11 struct { F0 int32; F1 uint8; F2 int64 }
type GenS12 struct { F0 uint; F1 int8; F2 int64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int8; F1 string; F2 int32 }
type GenS15 struct { F0 uint32; F1 uint; F2 string }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 uint16; F1 uint32 }
type GenS18 struct { F0 bool; F1 uint16 }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 int32; F1 int32; F2 int32 }
type GenS21 struct { F0 uint64; F1 string; F2 int }
type GenS22 struct { F0 uint; F1 uint }
func genFn0() uint16 {
    for fi4 := 0; fi4 < 19; fi4++ {
        if (0 != 4294967295) {
            switch 127 {
            case -1:
                var v5 uint8 = 0
                _ = v5
            case -2147483648:
                fmt.Println(fi4)
            }
            var v6 uint16 = 0
            _ = v6
        }
        _ = fi4
    }
    switch "a" {
    case "":
        var v7 float64 = -2.25
        _ = v7
    case "\n\t":
        var v8 uint16 = 255
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: (4294967295 > 685)}
            var v12 float32 = (float32(1.5) / (1e10 + 0.0))
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    if false {
        var v13 uint = 591
        v13 = 2
        _ = v13
    }
    var v14 int = -2147483648
    _ = v14
    return 255
}
func genFn1() uint64 {
    switch (("hello" + "foo bar") + "") {
    case "hello":
        if (-2147483648 < 171) {
            v15 := GenS1{F0: false}
            v16 := []int{}
            _ = v15
            _ = v16
        } else {
            var v17 float32 = 1e10
            v18 := GenS2{F0: 0, F1: 4294967295}
            _ = v17
            _ = v18
        }
    case "foo bar":
        for fi19 := 0; fi19 < 8; fi19++ {
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 int16 = -556
                v21 -= (v21 + 42)
                if true {
                    v22 := []int{fi19}
                    var v23 int16 = 0
                    _ = v22
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            if (!(false || ((!(true && (true || false))) || (true && (42 != 0))))) {
                var v24 uint = 18446744073709551615
                v24 -= v24
                _ = v24
            } else {
                v25 := GenS3{F0: (655 <= 255)}
                v26 := GenS1{F0: (!true)}
                _ = v25
                _ = v26
            }
            _ = fi19
        }
    case "\n\t":
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 int = 4294967295
            _ = fi27
            _ = v28
        }
    }
    v29 := []int{2147483647, -2147483648}
    _ = v29
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 127
    v3 := []string{(("hello" + "foo bar") + "hello"), ("\n\t" + ("" + "\n\t"))}
    v3 = append(v3, "")
    if true {
        for fi30 := 0; fi30 < 13; fi30++ {
            if true {
                if true {
                    v31 := v1[(fi30 * fi30)]
                    _ = v31
                }
                if (!(float32(0.0) == (-2.25 - 0.0))) {
                    if false {
                        if true {
                            fmt.Println(v2, v2)
                            if false {
                                var v32 float64 = 3.14159
                                _ = v32
                            } else {
                                v2 -= (v2 << ((202) & 63))
                            }
                        }
                    } else {
                        v33 := GenS4{F0: fi30}
                        v2 = 42
                        switch ("a" + "foo bar") {
                        case "foo bar":
                            v2 += (v2 ^ v2)
                        case "\n\t":
                            switch fi30 {
                            case -1:
                                var v34 float32 = 1e10
                                _ = v34
                            case 32768:
                                v35 := GenS2{F0: -71, F1: -2147483648}
                                _ = v35
                            case 9223372036854775807:
                                for fi36 := 0; fi36 < 14; fi36++ {
                                    fmt.Println(fi30, fi30)
                                    if false {
                                        switch "a" {
                                        case "foo bar":
                                            var v37 int8 = -24
                                            v38 := GenS1{F0: (!(!true))}
                                            _ = v37
                                            _ = v38
                                        case "a":
                                            var v39 int = fi36
                                            _ = v39
                                        default:
                                            v40 := v1[fi36]
                                            _ = v40
                                        }
                                        var v41 int32 = -159
                                        switch "foo bar" {
                                        case "foo bar":
                                            if true {
                                                if (!(42 == int32(fi36))) {
                                                    fmt.Println(v41)
                                                    fmt.Println(fi30, fi36)
                                                    var v42 uint32 = 4294967295
                                                    v42 = 2
                                                    _ = v42
                                                }
                                                var v43 string = ("a" + "foo bar")
                                                _ = v43
                                            }
                                        }
                                        _ = v41
                                    }
                                    var v44 bool = false
                                    _ = fi36
                                    _ = v44
                                }
                            default:
                                for fi45 := 0; fi45 < 1; fi45++ {
                                    if (true && (32767 > -32768)) {
                                        v2 += 214
                                        v46 := GenS4{F0: fi45}
                                        _ = v46
                                    } else {
                                        v2 = (^(v2))
                                        if len(v3) > 0 {
                                            v3[0] = "a"
                                        }
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                    v2 = v2
                                    _ = fi45
                                }
                            }
                        case "":
                            fmt.Println(fi30, v2)
                        }
                        _ = v33
                    }
                    fmt.Println(fi30, fi30)
                    v2 = v2
                }
            } else {
                if false {
                    _ = len(v1)
                } else {
                    if true {
                        var v47 float32 = 3.14159
                        _ = v47
                    }
                    switch fi30 {
                    case 42:
                        var v48 uint32 = 1
                        _ = v48
                    case 207:
                        fmt.Println(fi30, v2)
                    case 0:
                        for fi49 := 0; fi49 < 7; fi49++ {
                            if len(v3) > 0 {
                                v3[0] = ""
                            }
                            var v50 float64 = 1.5
                            _ = fi49
                            _ = v50
                        }
                    default:
                        fmt.Println(fi30, fi30)
                    }
                    v2 += (v2 & v2)
                }
                v2 = (v2 % (((v2 - v2)) | 1))
            }
            v2 -= v2
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch 611 {
                case -467:
                    if (1 <= -9223372036854775808) {
                        var v51 uint8 = 130
                        if false {
                            fmt.Println(v51, v51)
                            if false {
                                v2 -= 158
                                if true {
                                    fmt.Println(v2, v2, v2)
                                    var v52 int64 = 0
                                    _ = v52
                                } else {
                                    var v53 uint32 = 370
                                    var v54 int64 = -9223372036854775808
                                    _ = v53
                                    _ = v54
                                }
                                fmt.Println(v51, v51)
                            }
                            if len(v3) > 0 {
                                v3[0] = ("foo bar" + "")
                            }
                        } else {
                            switch "hello" {
                            case "\n\t":
                                var v55 bool = (false || (!(false || (100 == 10))))
                                _ = v55
                            case "":
                                v2 = v2
                            }
                        }
                        _ = v51
                    }
                case 0:
                    if (true && false) {
                        fmt.Println(fi30, v2)
                    }
                case 42:
                    v56 := GenS5{F0: 18446744073709551615, F1: v2}
                    _ = v56
                default:
                    fmt.Println(fi30)
                }
            } else {
                for fi57 := 0; fi57 < 9; fi57++ {
                    v2 -= v2
                    _ = fi57
                }
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                if true {
                    fmt.Println(v2, v2)
                    switch fi30 {
                    case 1:
                        if (806 <= fi30) {
                            for fi58 := 0; fi58 < 1; fi58++ {
                                var v59 bool = (-743 > -337)
                                if v59 {
                                    v60 := GenS2{F0: -128, F1: (fi30 + (fi30 << ((964) & 63)))}
                                    _ = v60
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        if v59 {
                                            v61 := GenS6{F0: "", F1: -8}
                                            switch fi58 {
                                            case -9223372036854775808:
                                                v2 -= v2
                                            default:
                                                switch int(v2) {
                                                case -2147483648:
                                                    v59 = ((v59 || (v59 && (v59 && (v59 && (-1 > 1))))) && v59)
                                                case 166:
                                                    var v62 int8 = 100
                                                    _ = v62
                                                default:
                                                    fmt.Println(fi58)
                                                    if ("" >= ("foo bar" + "foo bar")) {
                                                        v2 -= (v2 / ((v2) | 1))
                                                    } else {
                                                        v2 = (v2 >> ((uint(v2)) & 63))
                                                        fmt.Println(v2, v59)
                                                        if v59 {
                                                            fmt.Println(fi30)
                                                            switch (^(fi58)) {
                                                            case 9223372036854775807:
                                                                fmt.Println(fi58, v2)
                                                                if v59 {
                                                                    if (2 < 2) {
                                                                        v63 := GenS1{F0: false}
                                                                        _ = v63
                                                                    } else {
                                                                        v64 := GenS4{F0: (fi58 | 128)}
                                                                        _ = v64
                                                                    }
                                                                    if (-222 < 42) {
                                                                        v2 += (^(v2))
                                                                        v65 := GenS6{F0: ("\n\t" + "hello"), F1: -40}
                                                                        _ = v65
                                                                    } else {
                                                                        v2 = (^(v2))
                                                                        fmt.Println(fi30, v2, v2)
                                                                        if (genFn0() != 709) {
                                                                            if v59 {
                                                                                var v66 uint32 = 0
                                                                                _ = v66
                                                                            } else {
                                                                                v67 := GenS2{F0: 1, F1: -9223372036854775808}
                                                                                var v68 uint16 = 32767
                                                                                v68 = 556
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                            v69 := GenS5{F0: 144, F1: (^(v2))}
                                                                            _ = v69
                                                                        }
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "hello"
                                                                        }
                                                                        switch "" {
                                                                        case "a":
                                                                            if len(v3) > 0 {
                                                                                v3[0] = "hello"
                                                                            }
                                                                        case "\n\t":
                                                                            v70 := GenS7{F0: 127, F1: v59, F2: 42}
                                                                            _ = v70
                                                                        case "foo bar":
                                                                            _ = len(v3)
                                                                        default:
                                                                            v71 := GenS4{F0: 424}
                                                                            _ = v71
                                                                        }
                                                                    }
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        v72 := GenS1{F0: (!v59)}
                                                                        _ = v72
                                                                    case "\n\t":
                                                                        var v73 uint = 656
                                                                        _ = v73
                                                                    default:
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            fmt.Println(fi30, v59, fi58)
                                                                            var v74 uint64 = 249
                                                                            _ = v74
                                                                        default:
                                                                            v59 = true
                                                                        }
                                                                    }
                                                                }
                                                                v2 = (^(v2))
                                                                var v75 uint8 = v2
                                                                _ = v75
                                                            }
                                                            if true {
                                                                v2 = (^(v2))
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            v76 := GenS5{F0: 4294967295, F1: 255}
                                            if v59 {
                                                v77 := GenS7{F0: 66, F1: true, F2: fi30}
                                                var v78 int = (-(fi58))
                                                _ = v77
                                                _ = v78
                                            } else {
                                                var v79 float64 = -2.25
                                                _ = v79
                                            }
                                            _ = v61
                                            _ = v76
                                        }
                                    case "foo bar":
                                        var v80 uint32 = 71
                                        _ = v80
                                    case "hello":
                                        v2 = (^(v2))
                                    }
                                    v59 = (-1 <= -128)
                                }
                                var v81 uint16 = 479
                                _ = fi58
                                _ = v59
                                _ = v81
                            }
                        }
                    case 42:
                        if (true || true) {
                            switch "\n\t" {
                            case "a":
                                if false {
                                    v82 := GenS4{F0: (fi30 & (fi30 >> ((4294967295) & 63)))}
                                    for fi83 := 0; fi83 < 11; fi83++ {
                                        if false {
                                            v2 -= v2
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if (-9223372036854775808 != 797) {
                                                switch fi83 {
                                                case 42:
                                                    switch (fi30 + 32768) {
                                                    case 660:
                                                        var v84 int64 = -9223372036854775808
                                                        _ = v84
                                                    case -2147483648:
                                                        switch (-(fi30)) {
                                                        case 32768:
                                                            v2 = (v2 << ((406) & 63))
                                                            var v85 string = "hello"
                                                            v2 += v2
                                                            _ = v85
                                                        case -1:
                                                            switch fi30 {
                                                            case 56:
                                                                v86 := GenS8{F0: 2}
                                                                _ = v86
                                                            }
                                                        case -842:
                                                            var v87 uint8 = 1
                                                            _ = v87
                                                        default:
                                                            v88 := GenS8{F0: 1}
                                                            _ = v88
                                                        }
                                                        var v89 string = "foo bar"
                                                        var v90 int32 = 421
                                                        _ = v89
                                                        _ = v90
                                                    default:
                                                        var v91 float64 = 1e10
                                                        _ = v91
                                                    }
                                                case 128:
                                                    v2 = 1
                                                case 4294967295:
                                                    if ((0.1 >= 0.1) && false) {
                                                        switch fi83 {
                                                        case 804:
                                                            v92 := GenS2{F0: 101, F1: (fi83 >> ((2) & 63))}
                                                            _ = v92
                                                        case 127:
                                                            v2 = (v2 - (v2 / ((v2) | 1)))
                                                        case -372:
                                                            v93 := v1[-960]
                                                            _ = v93
                                                        default:
                                                            if (!false) {
                                                                if len(v3) > 0 {
                                                                    v3[0] = ("" + ("a" + ("" + "hello")))
                                                                }
                                                            } else {
                                                                v2 += v2
                                                                v2 = (v2 | 255)
                                                            }
                                                        }
                                                    } else {
                                                        v1[(fi30 | 400)] = (fi83 / ((fi30) | 1))
                                                        v94 := GenS3{F0: false}
                                                        _ = v94
                                                    }
                                                default:
                                                    v2 += 1
                                                }
                                                v95 := GenS9{F0: (4294967295 < 903)}
                                                if true {
                                                    if (!(!(42 <= 42))) {
                                                        v2 += (v2 / ((v2) | 1))
                                                        v2 -= 0
                                                        v2 = (v2 % ((128) | 1))
                                                    } else {
                                                        v1[fi83] = fi83
                                                        switch fi30 {
                                                        case -250:
                                                            var v96 bool = true
                                                            _ = v96
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if false {
                                                            v2 += (v2 % ((v2) | 1))
                                                            switch "a" {
                                                            case "":
                                                                var v97 bool = (0.0 != -2.25)
                                                                _ = v97
                                                            case "hello":
                                                                v98 := GenS4{F0: 32768}
                                                                _ = v98
                                                            default:
                                                                v2 = (^(v2))
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(fi83, fi30, fi83)
                                                    _ = len(v1)
                                                    v2 = v2
                                                    var v99 int32 = 42
                                                    _ = v99
                                                }
                                                var v100 uint16 = 231
                                                _ = v100
                                                _ = v95
                                            } else {
                                                if (!(!false)) {
                                                    v101 := GenS10{F0: 0, F1: -9223372036854775808}
                                                    v2 += (v2 / (((^(v2))) | 1))
                                                    switch "a" {
                                                    case "foo bar":
                                                        if (64 >= 900) {
                                                            var v102 int8 = -19
                                                            var v103 string = "foo bar"
                                                            switch v103 {
                                                            case "foo bar":
                                                                v102 -= 100
                                                            case "":
                                                                if true {
                                                                    if false {
                                                                        v102 = (v102 << ((65535) & 63))
                                                                    } else {
                                                                        var v104 uint16 = 32767
                                                                        if ((v2 | (v2 / (((v2 ^ v2)) | 1))) > 4) {
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                        }
                                                                        _ = v104
                                                                    }
                                                                } else {
                                                                    if ((true || (255 != (v2 | v2))) || (true && true)) {
                                                                        if true {
                                                                            v105 := GenS11{F0: 0, F1: 42, F2: -9223372036854775808}
                                                                            _ = v105
                                                                        } else {
                                                                            v103 = v103
                                                                            v102 = v102
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v103)
                                                                        var v106 uint8 = 128
                                                                        var v107 uint32 = 2147483647
                                                                        _ = v106
                                                                        _ = v107
                                                                    }
                                                                    var v108 int = (fi30 % (((fi30 << ((18446744073709551615) & 63))) | 1))
                                                                    _ = v108
                                                                }
                                                            default:
                                                                v102 += (v102 >> ((1) & 63))
                                                            }
                                                            _ = v102
                                                            _ = v103
                                                        } else {
                                                            v109 := GenS2{F0: -1, F1: -2147483648}
                                                            v110 := GenS6{F0: ("hello" + "hello"), F1: 42}
                                                            _ = v109
                                                            _ = v110
                                                        }
                                                    case "a":
                                                        var v111 int16 = 42
                                                        _ = v111
                                                    case "\n\t":
                                                        fmt.Println(v2)
                                                    }
                                                    _ = v101
                                                } else {
                                                    if false {
                                                        delete(v1, 998)
                                                    } else {
                                                        var v112 int = (fi83 + (fi83 * -2147483648))
                                                        _ = v112
                                                    }
                                                    switch ("hello" + ("foo bar" + "a")) {
                                                    case "":
                                                        v2 = (^(v2))
                                                        v113 := GenS2{F0: 1, F1: 9223372036854775807}
                                                        v2 = (v2 % ((v2) | 1))
                                                        _ = v113
                                                    case "hello":
                                                        if false {
                                                            fmt.Println(v2)
                                                            if true {
                                                                v114 := v1[fi83]
                                                                _ = v114
                                                            } else {
                                                                var v115 bool = (1.5 <= -2.25)
                                                                _ = v115
                                                            }
                                                            v116 := GenS12{F0: 454, F1: 17, F2: 566}
                                                            _ = v116
                                                        } else {
                                                            if false {
                                                                fmt.Println(fi30, fi30)
                                                            } else {
                                                                fmt.Println(v2, fi30, fi83)
                                                            }
                                                            var v117 int64 = 1
                                                            _ = v117
                                                        }
                                                    }
                                                }
                                                if false {
                                                    v2 += (v2 ^ v2)
                                                } else {
                                                    var v118 uint16 = genFn0()
                                                    _ = v118
                                                }
                                                var v119 float64 = ((0.1 - 3.14159) / -2.25)
                                                var v120 float64 = 3.14159
                                                v119 = (0.0 / v119)
                                                _ = v119
                                                _ = v120
                                            }
                                            v2 += v2
                                        }
                                        v121 := v1[fi30]
                                        _ = fi83
                                        _ = v121
                                    }
                                    switch 2147483647 {
                                    case 0:
                                        var v122 uint = 1
                                        _ = v122
                                    case 499:
                                        var v123 int16 = 42
                                        _ = v123
                                    case 1:
                                        v2 += (^(v2))
                                    }
                                    _ = v82
                                }
                            default:
                                fmt.Println(v2)
                            }
                            v2 -= v2
                        } else {
                            if true {
                                if (!(-483 >= -32768)) {
                                    v2 = v2
                                } else {
                                    fmt.Println(fi30, fi30)
                                    v124 := v1[(^(fi30))]
                                    _ = v124
                                }
                                var v125 float64 = 0.0
                                _ = v125
                            }
                            for fi126 := 0; fi126 < 16; fi126++ {
                                fmt.Println(v2, v2, fi126)
                                v2 -= (v2 >> ((42) & 63))
                                _ = fi126
                            }
                            v2 = v2
                        }
                    case 9223372036854775807:
                        var v127 float64 = float64(0.1)
                        _ = v127
                    default:
                        var v128 int = fi30
                        _ = v128
                    }
                    for fi129 := 0; fi129 < 19; fi129++ {
                        if true {
                            v3 = append(v3, "hello")
                            var v130 float32 = (1.5 / -2.25)
                            _ = v130
                        }
                        if true {
                            v131 := v1[fi30]
                            _ = v131
                        }
                        var v132 int = (^(fi129))
                        _ = fi129
                        _ = v132
                    }
                } else {
                    if ((!true) && true) {
                        v133 := v1[fi30]
                        _ = v133
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v2 = (v2 % ((v2) | 1))
                        var v134 uint8 = v2
                        _ = v134
                    }
                    fmt.Println(fi30, v2)
                }
            }
            _ = fi30
        }
        fmt.Println(v2)
    } else {
        if true {
            fmt.Println(v2)
        } else {
            var v135 uint32 = 557
            v1[-924] = -1
            _ = len(v1)
            _ = v135
        }
        for fi136 := 0; fi136 < 18; fi136++ {
            switch 9223372036854775807 {
            case 2147483647:
                if ((fi136 + fi136) <= fi136) {
                    v2 -= 255
                } else {
                    fmt.Println(v2)
                    var v137 uint16 = genFn0()
                    for fi138 := 0; fi138 < 3; fi138++ {
                        if (!true) {
                            var v139 uint = uint(v2)
                            _ = v139
                        }
                        _ = fi138
                    }
                    _ = v137
                }
            case 238:
                v140 := GenS10{F0: 2, F1: fi136}
                _ = v140
            case 4294967295:
                if ((((1.5 / 1.0) + 1.5) * (0.0 + 0.1)) > 0.1) {
                    v2 -= (v2 ^ 1)
                    var v141 bool = true
                    _ = v141
                } else {
                    var v142 uint16 = 32767
                    _ = v142
                }
            }
            v2 += v2
            _ = fi136
        }
        v2 -= v2
    }
    if true {
        var v143 int64 = 42
        v2 = 163
        _ = v143
    } else {
        if (!false) {
            v2 = 228
            if false {
                for fi144 := 0; fi144 < 6; fi144++ {
                    fmt.Println(fi144, fi144)
                    _ = fi144
                }
            }
        } else {
            var v145 string = "hello"
            var v146 int8 = -1
            fmt.Println(v146)
            _ = v145
            _ = v146
        }
    }
    v147 := GenS2{F0: 42, F1: -113}
    delete(v1, 42)
    v2 = v2
    v2 += (v2 >> ((18446744073709551615) & 63))
    if true {
        var v148 string = ("a" + "foo bar")
        var v149 int64 = 1
        _ = v148
        _ = v149
    }
    v1[9223372036854775807] = -646
    var v150 int32 = -965
    if true {
        var v151 int = 1
        var v152 int64 = -9223372036854775808
        if false {
            switch v151 {
            case 1:
                v2 -= 2
            case 42:
                for fi153 := 0; fi153 < 4; fi153++ {
                    var v154 float32 = 3.14159
                    var v155 string = ((("hello" + "foo bar") + ("hello" + (("a" + "a") + ("foo bar" + ("hello" + "\n\t"))))) + "")
                    var v156 uint8 = (v2 | 2)
                    _ = fi153
                    _ = v154
                    _ = v155
                    _ = v156
                }
            }
            for fi157 := 0; fi157 < 4; fi157++ {
                for fi158 := 0; fi158 < 5; fi158++ {
                    v151 -= (fi157 | fi157)
                    if ((!false) || true) {
                        if true {
                            v150 = v150
                            var v159 int8 = -48
                            v150 = (v150 / ((int32(v152)) | 1))
                            if (genFn0() >= 2) {
                                var v160 uint = 1
                                v152 = v152
                                var v161 string = "a"
                                if false {
                                    var v162 bool = (true || false)
                                    var v163 float64 = 0.1
                                    _ = v162
                                    _ = v163
                                } else {
                                    if (v2 <= (v2 / ((v2) | 1))) {
                                        switch v161 {
                                        case "\n\t":
                                            v164 := GenS12{F0: (v160 & (v160 % (((v160 % (((v160 * 356)) | 1))) | 1))), F1: v159, F2: (v152 << ((v160) & 63))}
                                            _ = v164
                                        default:
                                            v165 := GenS9{F0: (float32((0.0 - 1.5)) > 0.0)}
                                            _ = v165
                                        }
                                    } else {
                                        v150 += v150
                                        v166 := GenS13{F0: uint64(v160)}
                                        _ = v166
                                    }
                                    v167 := GenS11{F0: v150, F1: 128, F2: v152}
                                    _ = v167
                                }
                                _ = v160
                                _ = v161
                            }
                            _ = v159
                        }
                        v152 -= v152
                        if len(v3) > 0 {
                            v3[0] = ""
                        }
                    } else {
                        fmt.Println(v152, v2, v152)
                        v168 := GenS9{F0: (!(!false))}
                        var v169 uint64 = 0
                        _ = v168
                        _ = v169
                    }
                    var v170 uint8 = (v2 & 128)
                    _ = fi158
                    _ = v170
                }
                v171 := GenS8{F0: 255}
                if (916 > 326) {
                    switch "\n\t" {
                    case "":
                        if true {
                            switch (v151 % (((fi157 << ((65535) & 63))) | 1)) {
                            case -2147483648:
                                var v172 float64 = -2.25
                                _ = v172
                            case 1:
                                v152 += (-(v152))
                            case 332:
                                v2 = (v2 - 0)
                            default:
                                v152 -= (v152 * (v152 + (v152 * v152)))
                            }
                            v173 := GenS5{F0: 947, F1: v2}
                            _ = v173
                        }
                    case "\n\t":
                        for fi174 := 0; fi174 < 6; fi174++ {
                            switch (fi157 / ((42) | 1)) {
                            case 2147483647:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        fmt.Println(v150)
                                    case "foo bar":
                                        if (0 == (fi174 % (((fi157 / (((fi174 % ((fi157) | 1))) | 1))) | 1))) {
                                            fmt.Println(v152, fi174, v151)
                                        }
                                    default:
                                        fmt.Println(v2, fi157, v2)
                                    }
                                    v175 := GenS14{F0: -126, F1: "a", F2: (v150 ^ v150)}
                                    _ = v175
                                } else {
                                    v150 -= v150
                                    var v176 float64 = (((1.5 + 1.5) - 1.5) - 1e10)
                                    v177 := GenS5{F0: 42, F1: 0}
                                    var v178 float32 = (1e10 - -2.25)
                                    if ((-(v152)) != (v152 >> ((236) & 63))) {
                                        switch "foo bar" {
                                        case "hello":
                                            v176 -= float64(0.1)
                                        default:
                                            v179 := GenS15{F0: 65535, F1: 640, F2: "\n\t"}
                                            _ = v179
                                        }
                                    }
                                    fmt.Println(fi174, fi174)
                                    _ = v176
                                    _ = v177
                                    _ = v178
                                }
                                v151 -= v151
                            case 127:
                                switch (fi157 % ((fi157) | 1)) {
                                case 0:
                                    v150 = (v150 ^ 1)
                                default:
                                    if (((true && true) && ((^(v150)) <= 42)) && false) {
                                        _ = len(v1)
                                    }
                                }
                            }
                            v151 += -2147483648
                            if true {
                                var v180 int = -1
                                delete(v1, fi174)
                                v181 := GenS16{F0: 662}
                                _ = v180
                                _ = v181
                            }
                            _ = fi174
                        }
                    default:
                        var v182 bool = true
                        _ = v182
                    }
                    fmt.Println(v151, v152)
                    for fi183 := 0; fi183 < 9; fi183++ {
                        v150 += (v150 >> ((65535) & 63))
                        v1[(fi157 * (-(fi183)))] = (-(fi183))
                        _ = fi183
                    }
                }
                _ = fi157
                _ = v171
            }
            v151 += v151
        } else {
            switch (("" + (("\n\t" + "a") + "")) + ("a" + "hello")) {
            case "foo bar":
                var v184 uint64 = 42
                _ = v184
            case "hello":
                v185 := v1[v151]
                _ = v185
            default:
                switch (("hello" + "a") + ("" + ("hello" + ("\n\t" + "")))) {
                case "a":
                    v186 := GenS16{F0: 1}
                    _ = v186
                case "":
                    var v187 uint32 = 4294967295
                    _ = v187
                case "\n\t":
                    switch (v151 + v151) {
                    case 0:
                        v3 = append(v3, "hello")
                    case -536:
                        v151 -= (v151 - v151)
                    default:
                        var v188 uint8 = (^(v2))
                        _ = v188
                    }
                }
            }
        }
        _ = v151
        _ = v152
    } else {
        if false {
            v189 := GenS7{F0: -114, F1: (((1.5 - 0.0) < 3.14159) || true), F2: -2147483648}
            v190 := GenS10{F0: 4294967295, F1: -456}
            _ = v189
            _ = v190
        } else {
            if false {
                v191 := GenS9{F0: false}
                switch 4294967295 {
                case 128:
                    v192 := GenS17{F0: 28, F1: 311}
                    _ = v192
                case 32767:
                    if (42 == 54) {
                        v2 = (v2 % ((2) | 1))
                        if true {
                            if (-2.25 > 0.0) {
                                for fi193 := 0; fi193 < 17; fi193++ {
                                    if ((((0.0 / 0.1) + 1.5) >= 0.0) || false) {
                                        for fi194 := 0; fi194 < 6; fi194++ {
                                            if false {
                                                v195 := GenS5{F0: 0, F1: (v2 * v2)}
                                                v196 := GenS18{F0: true, F1: 32767}
                                                v2 += 42
                                                _ = v195
                                                _ = v196
                                            }
                                            _ = fi194
                                        }
                                        var v197 string = ""
                                        v2 -= v2
                                        _ = v197
                                    } else {
                                        if (v2 != (v2 * 127)) {
                                            if true {
                                                if (-1 >= v150) {
                                                    if (!true) {
                                                        v2 = v2
                                                        switch ("foo bar" + ("hello" + "\n\t")) {
                                                        case "hello":
                                                            v198 := GenS19{F0: 65535}
                                                            _ = v198
                                                        }
                                                    } else {
                                                        v150 += v150
                                                        switch "hello" {
                                                        case "hello":
                                                            var v199 int8 = 1
                                                            _ = v199
                                                        }
                                                    }
                                                }
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                var v200 int32 = v150
                                                _ = v200
                                            }
                                            var v201 int16 = 1000
                                            _ = v201
                                        }
                                        for fi202 := 0; fi202 < 20; fi202++ {
                                            v150 += v150
                                            _ = fi202
                                        }
                                        for fi203 := 0; fi203 < 4; fi203++ {
                                            var v204 int16 = 0
                                            if (!((fi193 % ((fi193) | 1)) == fi203)) {
                                                delete(v1, fi203)
                                                v205 := GenS11{F0: v150, F1: (^(v2)), F2: -9223372036854775808}
                                                _ = v205
                                            } else {
                                                fmt.Println(v204, fi193, v204)
                                            }
                                            if true {
                                                var v206 uint32 = uint32(v2)
                                                _ = v206
                                            }
                                            _ = fi203
                                            _ = v204
                                        }
                                    }
                                    v207 := GenS15{F0: 2147483648, F1: 817, F2: ""}
                                    if false {
                                        v2 = (v2 % ((152) | 1))
                                        if false {
                                            v2 = (v2 >> ((761) & 63))
                                        } else {
                                            var v208 uint = 65535
                                            _ = v208
                                        }
                                    } else {
                                        var v209 float32 = (((-2.25 * 3.14159) + (1e10 / 1e10)) * 1.5)
                                        _ = v209
                                    }
                                    _ = fi193
                                    _ = v207
                                }
                            } else {
                                for fi210 := 0; fi210 < 4; fi210++ {
                                    for fi211 := 0; fi211 < 14; fi211++ {
                                        v150 -= int32(v2)
                                        v150 = (v150 << ((uint(v2)) & 63))
                                        v150 += v150
                                        _ = fi211
                                    }
                                    _ = len(v1)
                                    if false {
                                        v2 = (v2 - 0)
                                    }
                                    _ = fi210
                                }
                                v150 += (v150 & v150)
                                for fi212 := 0; fi212 < 16; fi212++ {
                                    if (true && false) {
                                        for fi213 := 0; fi213 < 5; fi213++ {
                                            v214 := GenS13{F0: 18446744073709551615}
                                            v150 += (v150 ^ v150)
                                            _ = len(v3)
                                            _ = fi213
                                            _ = v214
                                        }
                                        var v215 uint8 = (^(v2))
                                        if true {
                                            fmt.Println(v215)
                                            v216 := v1[fi212]
                                            _ = v216
                                        }
                                        _ = v215
                                    } else {
                                        v2 -= v2
                                        switch -2147483648 {
                                        case -965:
                                            var v217 int32 = int32(fi212)
                                            _ = v217
                                        case 948:
                                            for fi218 := 0; fi218 < 1; fi218++ {
                                                fmt.Println(fi218, fi218, v150)
                                                _ = fi218
                                            }
                                        case -100:
                                            if (false && (42 > 42)) {
                                                _ = len(v3)
                                                fmt.Println(fi212, v150)
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = "a"
                                                }
                                                v150 -= 2147483647
                                            }
                                        }
                                        switch int(v2) {
                                        case -711:
                                            var v219 uint16 = 2
                                            _ = v219
                                        default:
                                            if true {
                                                v2 += v2
                                            }
                                        }
                                    }
                                    v150 -= (^(v150))
                                    var v220 uint16 = 0
                                    _ = fi212
                                    _ = v220
                                }
                            }
                        } else {
                            var v221 int32 = (v150 - v150)
                            _ = len(v3)
                            _ = v221
                        }
                        for fi222 := 0; fi222 < 2; fi222++ {
                            var v223 uint32 = 2
                            _ = fi222
                            _ = v223
                        }
                    } else {
                        _ = len(v1)
                        v2 = (v2 * (v2 << ((613) & 63)))
                        v224 := GenS20{F0: v150, F1: -1, F2: v150}
                        _ = v224
                    }
                case -9223372036854775808:
                    switch 2147483647 {
                    case 127:
                        var v225 float64 = (-2.25 / float64((0.1 - 1e10)))
                        _ = v225
                    case -612:
                        if false {
                            v150 += 42
                            for fi226 := 0; fi226 < 3; fi226++ {
                                var v227 float64 = (float64(-2.25) * 0.1)
                                for fi228 := 0; fi228 < 17; fi228++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi228
                                }
                                _ = fi226
                                _ = v227
                            }
                            for fi229 := 0; fi229 < 20; fi229++ {
                                var v230 int64 = 9223372036854775807
                                _ = fi229
                                _ = v230
                            }
                        } else {
                            if true {
                                if true {
                                    if false {
                                        v2 += v2
                                    } else {
                                        v150 = 0
                                    }
                                    v2 = (v2 / (((v2 | (v2 % ((v2) | 1)))) | 1))
                                } else {
                                    v2 = 2
                                    var v231 int32 = v150
                                    _ = v231
                                }
                            }
                            var v232 uint64 = 824
                            v2 = uint8(v232)
                            _ = v232
                        }
                    default:
                        switch -26 {
                        case 0:
                            if true {
                                for fi233 := 0; fi233 < 8; fi233++ {
                                    v234 := GenS21{F0: genFn1(), F1: ("foo bar" + (("a" + "a") + "")), F2: (fi233 ^ -203)}
                                    v235 := GenS22{F0: 255, F1: 517}
                                    var v236 uint32 = 105
                                    _ = fi233
                                    _ = v234
                                    _ = v235
                                    _ = v236
                                }
                                for fi237 := 0; fi237 < 3; fi237++ {
                                    for fi238 := 0; fi238 < 5; fi238++ {
                                        fmt.Println(fi238)
                                        v239 := v1[fi238]
                                        _ = fi238
                                        _ = v239
                                    }
                                    var v240 int16 = -197
                                    _ = fi237
                                    _ = v240
                                }
                                if (2 == 1) {
                                    v150 -= -81
                                } else {
                                    v2 += v2
                                    for fi241 := 0; fi241 < 20; fi241++ {
                                        v2 += (v2 / (((v2 + 1)) | 1))
                                        v150 = v150
                                        v2 -= 128
                                        _ = fi241
                                    }
                                }
                            }
                        }
                    }
                }
                _ = v191
            }
        }
        fmt.Println(v2, v150)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v150:", v150)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v147
    _ = v150
    _ = v2
    _ = v3
}
