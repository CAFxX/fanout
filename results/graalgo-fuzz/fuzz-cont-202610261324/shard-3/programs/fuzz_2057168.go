package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint }
type GenS3 struct { F0 string; F1 int32; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 uint32; F1 uint16 }
type GenS9 struct { F0 int16; F1 int16; F2 int16 }
type GenS10 struct { F0 int8; F1 uint64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint32; F1 int32 }
type GenS13 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS14 struct { F0 int32; F1 uint32; F2 int8 }
type GenS15 struct { F0 uint64; F1 int64 }
type GenS16 struct { F0 uint16; F1 uint16 }
type GenS17 struct { F0 int8; F1 uint32; F2 int }
type GenS18 struct { F0 int16; F1 int32; F2 uint16 }
type GenS19 struct { F0 string }
type GenS20 struct { F0 uint; F1 uint32 }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 uint; F1 bool; F2 int8 }
type GenS23 struct { F0 bool; F1 int32; F2 uint32 }
type GenS24 struct { F0 int16 }
type GenS25 struct { F0 uint }
type GenS26 struct { F0 string; F1 uint64; F2 uint16 }
type GenS27 struct { F0 string; F1 uint32 }
type GenS28 struct { F0 bool; F1 int64; F2 uint64 }
type GenS29 struct { F0 bool; F1 uint; F2 string }
type GenS30 struct { F0 int16; F1 int16 }
type GenS31 struct { F0 uint; F1 uint64 }
type GenS32 struct { F0 int16; F1 int64; F2 uint }
type GenS33 struct { F0 uint; F1 int8 }
func genFn0() uint {
    if false {
        var v5 int64 = -9223372036854775808
        _ = v5
    }
    v6 := GenS1{F0: false, F1: -2147483648, F2: 32767}
    switch "hello" {
    case "hello":
        var v7 int8 = 67
        _ = v7
    case "a":
        var v8 int = -702
        _ = v8
    case "foo bar":
        var v9 uint8 = 2
        _ = v9
    }
    if false {
        switch (("a" + "hello") + ("hello" + "hello")) {
        case "foo bar":
            v10 := GenS2{F0: 2, F1: 4294967295, F2: 344}
            _ = v10
        case "hello":
            var v11 int = -2147483648
            _ = v11
        case "a":
            for fi12 := 0; fi12 < 9; fi12++ {
                if false {
                    v13 := GenS3{F0: "a", F1: 1, F2: 18446744073709551615}
                    _ = v13
                }
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 bool = (-905 >= (fi14 % ((9223372036854775807) | 1)))
                    v16 := GenS4{F0: 0, F1: 0, F2: 255}
                    var v17 int = fi12
                    v18 := GenS5{F0: 625}
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                for fi19 := 0; fi19 < 5; fi19++ {
                    v20 := []int{fi12, (fi12 << ((42) & 63)), (fi12 ^ (fi12 >> ((0) & 63)))}
                    var v21 uint64 = 1
                    v22 := GenS1{F0: false, F1: -2147483648, F2: (fi19 + (^(fi12)))}
                    if (!false) {
                        v21 = (v21 + 42)
                        if (((!(!(("" + "foo bar") <= "\n\t"))) || true) && false) {
                            var v23 int64 = -1
                            var v24 uint32 = 42
                            if false {
                                v25 := GenS3{F0: "a", F1: -1, F2: v21}
                                switch -738 {
                                case 9223372036854775807:
                                    if true {
                                        fmt.Println(v24, v24, v23)
                                        v26 := GenS6{F0: 0, F1: 1, F2: 562}
                                        var v27 float32 = -2.25
                                        _ = v26
                                        _ = v27
                                    }
                                case 32767:
                                    v24 = v24
                                case 538:
                                    v28 := GenS1{F0: (42 >= 255), F1: 42, F2: (^(fi12))}
                                    fmt.Println(fi19, fi12, v24)
                                    var v29 string = "a"
                                    _ = v28
                                    _ = v29
                                }
                                _ = v25
                            } else {
                                fmt.Println(fi12)
                                switch "a" {
                                case "a":
                                    fmt.Println(fi19, v23, v21)
                                    fmt.Println(v24, v21)
                                    fmt.Println(v24, fi12, v23)
                                default:
                                    var v30 uint = 1
                                    _ = v30
                                }
                                var v31 int = fi19
                                _ = v31
                            }
                            _ = v23
                            _ = v24
                        }
                        v32 := GenS3{F0: "foo bar", F1: -1, F2: v21}
                        _ = v32
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = fi12
            }
        default:
            v33 := GenS2{F0: 52, F1: 938, F2: 124}
            _ = v33
        }
    } else {
        var v34 int = 4294967295
        for fi35 := 0; fi35 < 9; fi35++ {
            for fi36 := 0; fi36 < 8; fi36++ {
                fmt.Println(v34)
                fmt.Println(fi35, v34, v34)
                _ = fi36
            }
            for fi37 := 0; fi37 < 2; fi37++ {
                switch "" {
                case "a":
                    if false {
                        fmt.Println(v34)
                        v38 := GenS7{F0: -124, F1: 196}
                        var v39 uint8 = 1
                        v40 := []int{42}
                        var v41 int64 = 353
                        _ = v38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                case "foo bar":
                    if false {
                        if (964 != -353) {
                            v42 := GenS4{F0: 2147483647, F1: 0, F2: 143}
                            _ = v42
                        } else {
                            if true {
                                switch fi35 {
                                case 4294967295:
                                    var v43 uint32 = 899
                                    _ = v43
                                default:
                                    v44 := []int{2147483647, fi37, (^(v34)), (-(fi35))}
                                    _ = v44
                                }
                            }
                            v34 = v34
                        }
                        if true {
                            if (!true) {
                                if (-515 > 1) {
                                    switch v34 {
                                    case 741:
                                        v34 = (v34 | (fi37 / ((127) | 1)))
                                    case 481:
                                        v34 -= fi35
                                    case 4294967295:
                                        if false {
                                            if (true && (925 >= 915)) {
                                                var v45 float64 = 1.5
                                                fmt.Println(v45, fi37)
                                                _ = v45
                                            }
                                        }
                                    default:
                                        switch "hello" {
                                        case "foo bar":
                                            v34 += fi37
                                        case "a":
                                            fmt.Println(v34, fi35, fi35)
                                        case "\n\t":
                                            switch ("" + (("foo bar" + "a") + "\n\t")) {
                                            case "foo bar":
                                                v46 := GenS5{F0: 4294967295}
                                                _ = v46
                                            case "a":
                                                fmt.Println(fi35, fi35, fi37)
                                            case "":
                                                fmt.Println(fi35)
                                            default:
                                                var v47 uint = 18446744073709551615
                                                var v48 uint64 = 2
                                                if (2 <= 885) {
                                                    v49 := GenS2{F0: v48, F1: 42, F2: (v47 << (((v47 / ((v47) | 1))) & 63))}
                                                    var v50 float64 = 1.5
                                                    switch (fi37 + fi35) {
                                                    case 32767:
                                                        if ((3.14159 * 3.14159) == 0.0) {
                                                            switch ("a" + ("\n\t" + "\n\t")) {
                                                            case "":
                                                                v51 := GenS1{F0: false, F1: -697, F2: fi37}
                                                                _ = v51
                                                            case "hello":
                                                                var v52 uint8 = 255
                                                                _ = v52
                                                            case "a":
                                                                v53 := GenS5{F0: 270}
                                                                _ = v53
                                                            }
                                                            if (!true) {
                                                                v48 -= v48
                                                                if (186 > 2147483647) {
                                                                    v54 := []int{}
                                                                    _ = v54
                                                                } else {
                                                                    v50 = v50
                                                                    fmt.Println(v47, v48)
                                                                }
                                                                v55 := []int{2147483647, -2147483648, fi35}
                                                                _ = v55
                                                            }
                                                            if (65535 != 32767) {
                                                                fmt.Println(v34)
                                                                if false {
                                                                    v34 = (fi37 * 369)
                                                                    v56 := []int{-1, fi37, fi37, (^(fi37))}
                                                                    var v57 float32 = 3.14159
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                                switch "a" {
                                                                case "\n\t":
                                                                    if true {
                                                                        var v58 uint64 = (^(v48))
                                                                        if ((false || true) || (fi35 > (fi35 & v34))) {
                                                                            v48 = (v48 + 0)
                                                                            v59 := GenS4{F0: -997, F1: 2147483647, F2: 1}
                                                                            _ = v59
                                                                        }
                                                                        var v60 string = "foo bar"
                                                                        _ = v58
                                                                        _ = v60
                                                                    } else {
                                                                        v61 := GenS8{F0: 0, F1: uint16(v47)}
                                                                        _ = v61
                                                                    }
                                                                case "":
                                                                    v62 := []int{}
                                                                    _ = v62
                                                                default:
                                                                    var v63 uint8 = 2
                                                                    _ = v63
                                                                }
                                                                v64 := GenS9{F0: 262, F1: 1, F2: 1}
                                                                if true {
                                                                    fmt.Println(fi35, v50, v47)
                                                                    var v65 float64 = (v50 + (float64(3.14159) - (3.14159 * (0.1 - float64(3.14159)))))
                                                                    _ = v65
                                                                } else {
                                                                    v66 := GenS10{F0: 1, F1: (v48 * (v48 | (v48 << (((^(v47))) & 63))))}
                                                                    switch "hello" {
                                                                    case "":
                                                                        var v67 uint16 = 255
                                                                        _ = v67
                                                                    case "\n\t":
                                                                        if true {
                                                                            var v68 uint16 = 42
                                                                            v34 = 0
                                                                            var v69 int16 = -79
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                    }
                                                                    _ = v66
                                                                }
                                                                _ = v64
                                                            } else {
                                                                if true {
                                                                    switch fi37 {
                                                                    case 4294967295:
                                                                        var v70 int32 = 1
                                                                        _ = v70
                                                                    case 134:
                                                                        v50 += v50
                                                                    }
                                                                    v71 := []int{fi35}
                                                                    _ = v71
                                                                } else {
                                                                    if true {
                                                                        v72 := []int{(fi35 & (^(fi35))), fi35}
                                                                        var v73 float64 = v50
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                    if true {
                                                                        var v74 int32 = -399
                                                                        _ = v74
                                                                    }
                                                                    fmt.Println(v47, v48)
                                                                    if true {
                                                                        v75 := GenS11{F0: -660}
                                                                        _ = v75
                                                                    }
                                                                    if false {
                                                                        v76 := []int{}
                                                                        _ = v76
                                                                    }
                                                                }
                                                                if (127 > 1) {
                                                                    fmt.Println(fi37)
                                                                    fmt.Println(v47, fi35)
                                                                    switch "" {
                                                                    case "hello":
                                                                        v50 = -2.25
                                                                    case "\n\t":
                                                                        v47 = (v47 * v47)
                                                                    default:
                                                                        v47 += (v47 * v47)
                                                                    }
                                                                    if (2147483647 < -2147483648) {
                                                                        var v77 float32 = -2.25
                                                                        if true {
                                                                            v78 := GenS5{F0: (v47 % (((v47 + v47)) | 1))}
                                                                            fmt.Println(v34, v48)
                                                                            var v79 int32 = 831
                                                                            _ = v78
                                                                            _ = v79
                                                                        }
                                                                        _ = v77
                                                                    } else {
                                                                        if false {
                                                                            v47 = v47
                                                                            var v80 float64 = v50
                                                                            if ("a" == "") {
                                                                                v50 = v50
                                                                            } else {
                                                                                v81 := []int{-9223372036854775808, (v34 | (v34 ^ (fi35 % ((fi35) | 1))))}
                                                                                if len(v81) >= 2 {
                                                                                    v81 = v81[1:len(v81)]
                                                                                }
                                                                                v34 -= fi37
                                                                                var v82 uint8 = 255
                                                                                switch ("\n\t" + "a") {
                                                                                case "\n\t":
                                                                                    var v83 int = (fi35 / ((v34) | 1))
                                                                                    _ = v83
                                                                                case "":
                                                                                    if (!(!(((1.5 - v80) + ((-2.25 / v50) / 1e10)) > (float64(1.5) / 0.1)))) {
                                                                                        var v84 uint32 = 764
                                                                                        var v85 string = "hello"
                                                                                        v86 := GenS10{F0: 117, F1: v48}
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                        _ = v86
                                                                                    } else {
                                                                                        var v87 int8 = 1
                                                                                        v50 = (v50 * v80)
                                                                                        if false {
                                                                                            if (!(false && true)) {
                                                                                                switch (^(fi35)) {
                                                                                                case -638:
                                                                                                    if false {
                                                                                                        var v88 float32 = 1e10
                                                                                                        if (false || (0 == 676)) {
                                                                                                            switch int(v48) {
                                                                                                            case 128:
                                                                                                                v89 := GenS7{F0: 1, F1: v47}
                                                                                                                _ = v89
                                                                                                            default:
                                                                                                                var v90 int8 = 118
                                                                                                                _ = v90
                                                                                                            }
                                                                                                            if (4294967295 < (fi35 % ((-631) | 1))) {
                                                                                                                if true {
                                                                                                                    switch (v34 ^ 128) {
                                                                                                                    case 4294967295:
                                                                                                                        if (false && false) {
                                                                                                                            if true {
                                                                                                                                v91 := GenS12{F0: 465, F1: 42}
                                                                                                                                _ = v91
                                                                                                                            } else {
                                                                                                                                v87 = 42
                                                                                                                            }
                                                                                                                            if false {
                                                                                                                                var v92 int = 4294967295
                                                                                                                                v93 := GenS13{F0: 1, F1: (^(v82)), F2: (v82 * v82)}
                                                                                                                                if false {
                                                                                                                                    var v94 uint = 0
                                                                                                                                    v95 := GenS14{F0: 0, F1: 200, F2: 0}
                                                                                                                                    if true {
                                                                                                                                        v92 -= int(v94)
                                                                                                                                        fmt.Println(fi37, v87)
                                                                                                                                        v96 := GenS15{F0: 4294967295, F1: 0}
                                                                                                                                        _ = v96
                                                                                                                                    }
                                                                                                                                    _ = v94
                                                                                                                                    _ = v95
                                                                                                                                }
                                                                                                                                _ = v92
                                                                                                                                _ = v93
                                                                                                                            } else {
                                                                                                                                if true {
                                                                                                                                    if (uint32(v47) >= 2147483647) {
                                                                                                                                        var v97 int16 = -466
                                                                                                                                        _ = v97
                                                                                                                                    } else {
                                                                                                                                        var v98 float32 = -2.25
                                                                                                                                        _ = v98
                                                                                                                                    }
                                                                                                                                    var v99 int16 = 42
                                                                                                                                    fmt.Println(v48, v88, v88)
                                                                                                                                    _ = v99
                                                                                                                                } else {
                                                                                                                                    var v100 int8 = 127
                                                                                                                                    if (!true) {
                                                                                                                                        var v101 uint16 = 317
                                                                                                                                        v102 := GenS16{F0: v101, F1: v101}
                                                                                                                                        v80 -= (((float64(v88) / 1e10) + (v50 + (v80 * (v80 - (0.0 * v50))))) - 0.1)
                                                                                                                                        _ = v101
                                                                                                                                        _ = v102
                                                                                                                                    }
                                                                                                                                    _ = v100
                                                                                                                                }
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            if false {
                                                                                                                                v103 := GenS3{F0: ("\n\t" + ""), F1: 0, F2: v48}
                                                                                                                                v82 = uint8(v47)
                                                                                                                                var v104 uint16 = 415
                                                                                                                                _ = v103
                                                                                                                                _ = v104
                                                                                                                            }
                                                                                                                            v87 = v87
                                                                                                                            switch (("hello" + "\n\t") + ("hello" + ("" + ("foo bar" + ("foo bar" + "foo bar"))))) {
                                                                                                                            case "a":
                                                                                                                                _ = len(v81)
                                                                                                                            default:
                                                                                                                                v81 = append(v81, -786)
                                                                                                                                v81 = append(v81, (fi37 | int(v87)))
                                                                                                                                v87 += v87
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case -9223372036854775808:
                                                                                                                        var v105 uint16 = 1
                                                                                                                        _ = v105
                                                                                                                    case 42:
                                                                                                                        if false {
                                                                                                                            if (true && false) {
                                                                                                                                v106 := GenS13{F0: 0, F1: v82, F2: (v82 | 93)}
                                                                                                                                _ = v106
                                                                                                                            } else {
                                                                                                                                switch (fi35 ^ -41) {
                                                                                                                                case 32767:
                                                                                                                                    v81 = append(v81, (fi37 << (((v47 & v47)) & 63)))
                                                                                                                                case 780:
                                                                                                                                    switch "" {
                                                                                                                                    case "hello":
                                                                                                                                        v82 = (v82 | v82)
                                                                                                                                    case "\n\t":
                                                                                                                                        v107 := GenS17{F0: v87, F1: 478, F2: v34}
                                                                                                                                        _ = v107
                                                                                                                                    case "":
                                                                                                                                        fmt.Println(v47, fi35)
                                                                                                                                    default:
                                                                                                                                        fmt.Println(v87, fi35, fi37)
                                                                                                                                    }
                                                                                                                                case 79:
                                                                                                                                    var v108 float32 = (v88 * 1.5)
                                                                                                                                    _ = v108
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v80 = ((v50 - 1.5) / (0.0 + (v80 / ((v80 / (v50 + 0.0)) - 0.0))))
                                                                                                                        } else {
                                                                                                                            var v109 uint8 = (v82 ^ (v82 % ((v82) | 1)))
                                                                                                                            v110 := GenS6{F0: uint16(v82), F1: -496, F2: 467}
                                                                                                                            v81 = append(v81, 1)
                                                                                                                            _ = v109
                                                                                                                            _ = v110
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch ("hello" + "foo bar") {
                                                                                                                    case "\n\t":
                                                                                                                        var v111 uint = v47
                                                                                                                        _ = v111
                                                                                                                    }
                                                                                                                    switch "" {
                                                                                                                    case "":
                                                                                                                        v88 -= 3.14159
                                                                                                                    default:
                                                                                                                        var v112 int16 = 1
                                                                                                                        _ = v112
                                                                                                                    }
                                                                                                                }
                                                                                                                var v113 int = fi37
                                                                                                                fmt.Println(v34, v34)
                                                                                                                _ = v113
                                                                                                            } else {
                                                                                                                if len(v81) >= 2 {
                                                                                                                    v81 = v81[1:len(v81)]
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v88
                                                                                                    }
                                                                                                case 567:
                                                                                                    fmt.Println(v34, v47, v80)
                                                                                                }
                                                                                                if (1 == v48) {
                                                                                                    var v114 uint = v47
                                                                                                    fmt.Println(v80, fi37, v82)
                                                                                                    v82 += (v82 << ((v114) & 63))
                                                                                                    _ = v114
                                                                                                }
                                                                                                v87 -= (^(v87))
                                                                                            } else {
                                                                                                fmt.Println(v48, v80)
                                                                                                if len(v81) > 0 {
                                                                                                    v81[0] = fi37
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if false {
                                                                                                v115 := GenS14{F0: -1, F1: uint32(v82), F2: 42}
                                                                                                var v116 uint64 = v48
                                                                                                _ = v115
                                                                                                _ = v116
                                                                                            }
                                                                                            _ = len(v81)
                                                                                        }
                                                                                        _ = v87
                                                                                    }
                                                                                default:
                                                                                    var v117 float32 = -2.25
                                                                                    _ = v117
                                                                                }
                                                                                v118 := GenS18{F0: 0, F1: 754, F2: 0}
                                                                                var v119 uint64 = v48
                                                                                v120 := GenS19{F0: "\n\t"}
                                                                                _ = v118
                                                                                _ = v119
                                                                                _ = v120
                                                                                _ = v81
                                                                                _ = v82
                                                                            }
                                                                            _ = v80
                                                                        } else {
                                                                            v121 := []int{(v34 << (((v47 ^ (^(v47)))) & 63)), (fi35 + v34), (v34 << ((v47) & 63))}
                                                                            if false {
                                                                                v47 = v47
                                                                                var v122 uint16 = 442
                                                                                var v123 int16 = 524
                                                                                var v124 uint8 = 255
                                                                                _ = v122
                                                                                _ = v123
                                                                                _ = v124
                                                                            } else {
                                                                                v121 = append(v121, (-(fi37)))
                                                                                v48 = (v48 | v48)
                                                                            }
                                                                            _ = v121
                                                                        }
                                                                    }
                                                                    v34 += 32767
                                                                    switch 127 {
                                                                    case 32768:
                                                                        v48 += v48
                                                                    case 1:
                                                                        var v125 int32 = 1
                                                                        _ = v125
                                                                    default:
                                                                        if false {
                                                                            fmt.Println(v34, fi37)
                                                                            var v126 float64 = (3.14159 * (1e10 + 0.1))
                                                                            v127 := []int{v34, -844, fi35, 128}
                                                                            _ = v126
                                                                            _ = v127
                                                                        }
                                                                    }
                                                                }
                                                                fmt.Println(fi35)
                                                            }
                                                        }
                                                    case 4294967295:
                                                        if (42 == v48) {
                                                            v128 := []int{-897, 4294967295, (v34 << ((930) & 63)), 841}
                                                            _ = v128
                                                        }
                                                    case 326:
                                                        v34 = -307
                                                    }
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    v34 = v34
                                                    v129 := GenS3{F0: "a", F1: 0, F2: v48}
                                                    var v130 int = -2147483648
                                                    _ = v129
                                                    _ = v130
                                                }
                                                _ = v47
                                                _ = v48
                                            }
                                        default:
                                            switch v34 {
                                            case -477:
                                                if true {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v34 = (fi35 >> ((42) & 63))
                                                    }
                                                    var v131 int16 = 32767
                                                    _ = v131
                                                } else {
                                                    v132 := GenS9{F0: 42, F1: 1, F2: 1000}
                                                    _ = v132
                                                }
                                            case 32767:
                                                var v133 float32 = (-2.25 / 3.14159)
                                                _ = v133
                                            default:
                                                v34 += (fi35 * (fi37 - -9223372036854775808))
                                            }
                                        }
                                    }
                                    if true {
                                        var v134 uint16 = 0
                                        v134 += v134
                                        _ = v134
                                    } else {
                                        switch "hello" {
                                        case "hello":
                                            switch 42 {
                                            case 42:
                                                var v135 bool = false
                                                _ = v135
                                            case -702:
                                                v136 := []int{fi37, (fi37 & v34), (fi35 ^ fi37), v34}
                                                _ = v136
                                            }
                                        case "a":
                                            if (68 >= -1) {
                                                var v137 uint64 = 0
                                                _ = v137
                                            }
                                        default:
                                            if (!false) {
                                                fmt.Println(v34)
                                            } else {
                                                var v138 int32 = 0
                                                v138 = 1
                                                _ = v138
                                            }
                                            var v139 int16 = 42
                                            var v140 int8 = 0
                                            _ = v139
                                            _ = v140
                                        }
                                        v141 := []int{v34, (fi35 >> ((1) & 63)), -399}
                                        if false {
                                            _ = len(v141)
                                        }
                                        _ = v141
                                    }
                                } else {
                                    switch ("a" + ("foo bar" + "foo bar")) {
                                    case "a":
                                        v34 = 127
                                    case "\n\t":
                                        if true {
                                            fmt.Println(fi37)
                                        }
                                    case "hello":
                                        v142 := GenS20{F0: 453, F1: 248}
                                        _ = v142
                                    default:
                                        switch (v34 << ((464) & 63)) {
                                        case 762:
                                            var v143 int = -786
                                            _ = v143
                                        case -269:
                                            switch ("a" + "") {
                                            case "hello":
                                                v144 := GenS19{F0: (("" + "foo bar") + "")}
                                                _ = v144
                                            default:
                                                var v145 string = "a"
                                                _ = v145
                                            }
                                        case 2147483647:
                                            if false {
                                                v34 += fi35
                                            }
                                        }
                                    }
                                }
                                fmt.Println(v34, v34)
                            } else {
                                if (false && false) {
                                    if false {
                                        var v146 float64 = (((3.14159 + 3.14159) / (-2.25 + (0.0 / (1e10 + 1.5)))) - 1e10)
                                        v147 := GenS21{F0: 1}
                                        v148 := GenS5{F0: 18446744073709551615}
                                        _ = v146
                                        _ = v147
                                        _ = v148
                                    } else {
                                        var v149 int = (v34 - fi37)
                                        v150 := []int{v149, v149, fi35, (v149 * fi37)}
                                        _ = v149
                                        _ = v150
                                    }
                                } else {
                                    v151 := []int{(v34 / ((fi35) | 1)), (fi35 % (((fi35 * 388)) | 1)), (fi37 + (fi37 << ((4294967295) & 63))), fi37}
                                    v152 := GenS1{F0: (94 > -128), F1: 0, F2: fi35}
                                    _ = v151
                                    _ = v152
                                }
                                if (0 <= 953) {
                                    v153 := GenS22{F0: 0, F1: true, F2: -7}
                                    var v154 int16 = -969
                                    v155 := []int{fi37}
                                    v156 := GenS23{F0: true, F1: 265, F2: 427}
                                    _ = len(v155)
                                    v157 := GenS16{F0: 255, F1: 530}
                                    v34 -= 46
                                    _ = v153
                                    _ = v154
                                    _ = v155
                                    _ = v156
                                    _ = v157
                                } else {
                                    if false {
                                        if (!(0.0 != 1.5)) {
                                            if (0.1 < 3.14159) {
                                                v34 = fi35
                                                v158 := []int{v34, (fi35 >> ((634) & 63))}
                                                v34 += (^(fi35))
                                                v34 += -9223372036854775808
                                                _ = v158
                                            }
                                            if true {
                                                var v159 uint = 255
                                                v34 += (fi37 * (v34 << ((65535) & 63)))
                                                fmt.Println(fi37, v159, v34)
                                                _ = v159
                                            } else {
                                                v34 += fi35
                                                v34 = fi35
                                                fmt.Println(fi35, v34)
                                            }
                                            v34 = 907
                                        } else {
                                            v160 := GenS24{F0: 32767}
                                            if (!true) {
                                                var v161 int32 = 2147483647
                                                var v162 int16 = 1000
                                                _ = v161
                                                _ = v162
                                            }
                                            _ = v160
                                        }
                                        var v163 int16 = -1
                                        _ = v163
                                    } else {
                                        v164 := GenS25{F0: 42}
                                        if false {
                                            if ((4294967295 == 2) && false) {
                                                v165 := []int{(fi37 + 0), v34}
                                                switch "foo bar" {
                                                case "\n\t":
                                                    fmt.Println(v34, fi37, v34)
                                                }
                                                _ = v165
                                            }
                                            v166 := GenS26{F0: "\n\t", F1: 18446744073709551615, F2: 360}
                                            v167 := []int{-906, fi37, fi35, -297}
                                            var v168 uint16 = 1
                                            _ = v166
                                            _ = v167
                                            _ = v168
                                        } else {
                                            v169 := GenS27{F0: (("" + ("" + "hello")) + "\n\t"), F1: 1}
                                            _ = v169
                                        }
                                        _ = v164
                                    }
                                    var v170 int8 = 1
                                    _ = v170
                                }
                                fmt.Println(fi35, fi35)
                            }
                            fmt.Println(fi35, fi37)
                        } else {
                            v34 = -9223372036854775808
                        }
                    }
                default:
                    v34 -= (v34 >> ((4294967295) & 63))
                }
                if false {
                    fmt.Println(fi37, fi37)
                    v34 += (v34 - (v34 | 32768))
                }
                var v171 float32 = 1.5
                v172 := []int{fi37, v34, (fi35 << ((2) & 63)), -9223372036854775808}
                _ = fi37
                _ = v171
                _ = v172
            }
            _ = fi35
        }
        v34 = 4294967295
        _ = v34
    }
    _ = v6
    return 638
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 42
    v3 := make(map[int]uint32)
    var v4 int16 = -531
    v4 -= v4
    v4 -= (v4 ^ 1000)
    var v173 int8 = -106
    for fi174 := 0; fi174 < 4; fi174++ {
        fmt.Println(v4)
        if false {
            var v175 float32 = float32(float64(1.5))
            _ = v175
        }
        _ = fi174
    }
    if false {
        var v176 float64 = 0.1
        v176 = (float64(1.5) * 0.1)
        _ = v176
    } else {
        v4 += v4
        if (false || false) {
            var v177 int32 = 2147483647
            switch ("a" + "a") {
            case "a":
                var v178 int64 = -1
                _ = v178
            }
            switch -1 {
            case -1:
                fmt.Println(v2, v177)
            case 1:
                if false {
                    v1["foo bar"] = 252
                    fmt.Println(v2)
                }
            case -9223372036854775808:
                var v179 uint8 = 70
                _ = v179
            default:
                v4 -= v4
            }
            _ = v177
        }
    }
    v173 = 0
    for fi180 := 0; fi180 < 3; fi180++ {
        if (true && (1 <= (v173 & v173))) {
            for fi181 := 0; fi181 < 17; fi181++ {
                v182 := GenS28{F0: true, F1: 0, F2: 0}
                _ = fi181
                _ = v182
            }
        } else {
            fmt.Println(fi180)
            v4 += (v4 + (v4 % ((v4) | 1)))
        }
        v183 := GenS29{F0: true, F1: (v2 / ((0) | 1)), F2: "a"}
        v184 := GenS29{F0: (42 != 255), F1: 0, F2: ""}
        _ = fi180
        _ = v183
        _ = v184
    }
    fmt.Println(v2, v4)
    if true {
        for fi185 := 0; fi185 < 1; fi185++ {
            switch "a" {
            case "foo bar":
                var v186 int8 = v173
                _ = v186
            default:
                switch fi185 {
                case 9223372036854775807:
                    v187 := GenS9{F0: v4, F1: -1, F2: v4}
                    _ = v187
                }
            }
            var v188 uint32 = 2147483648
            _ = fi185
            _ = v188
        }
        v173 = (v173 - v173)
    } else {
        v173 = (^(v173))
        if ((255 != 2) && true) {
            switch -2147483648 {
            case 787:
                switch ((("foo bar" + "\n\t") + (("hello" + ("\n\t" + ((("foo bar" + "foo bar") + "a") + ("hello" + "a")))) + "foo bar")) + (("a" + ("\n\t" + ("hello" + ""))) + (("a" + "a") + ("hello" + ("foo bar" + ("hello" + "\n\t")))))) {
                case "\n\t":
                    v2 = v2
                case "":
                    if false {
                        if true {
                            v2 -= (v2 / ((genFn0()) | 1))
                            v189 := []int{int(v173), 4294967295, -230}
                            _ = v189
                        }
                    }
                case "a":
                    fmt.Println(v2, v173)
                default:
                    if (42 <= 2) {
                        for fi190 := 0; fi190 < 2; fi190++ {
                            var v191 string = "a"
                            _ = fi190
                            _ = v191
                        }
                    } else {
                        var v192 string = "\n\t"
                        var v193 uint16 = 1
                        _ = v192
                        _ = v193
                    }
                }
            case 4294967295:
                v194 := []int{127}
                _ = v194
            case -454:
                if false {
                    v173 += v173
                    var v195 string = "foo bar"
                    v196 := GenS30{F0: (v4 & -1), F1: -1}
                    _ = v195
                    _ = v196
                }
            default:
                var v197 uint64 = 1
                _ = v197
            }
        } else {
            v1["hello"] = 2147483648
            if (!(false || false)) {
                v2 = (v2 << ((637) & 63))
                for fi198 := 0; fi198 < 15; fi198++ {
                    v2 += v2
                    if true {
                        v199 := GenS31{F0: v2, F1: 674}
                        _ = v199
                    }
                    _ = fi198
                }
            }
            if true {
                _ = len(v1)
                v200 := []int{}
                for fi201 := 0; fi201 < 6; fi201++ {
                    if true {
                        v2 = v2
                    }
                    fmt.Println(v2, v2, v2)
                    for fi202 := 0; fi202 < 15; fi202++ {
                        _ = len(v1)
                        _ = fi202
                    }
                    _ = fi201
                }
                _ = v200
            }
        }
        var v203 uint = v2
        _ = v203
    }
    v173 = (v173 | v173)
    for fi204 := 0; fi204 < 4; fi204++ {
        v205 := GenS32{F0: (v4 / ((42) | 1)), F1: -9223372036854775808, F2: 111}
        fmt.Println(fi204)
        _ = fi204
        _ = v205
    }
    if (!true) {
        switch 209 {
        case -483:
            if (65535 <= 2147483648) {
                if true {
                    v173 = (v173 % (((v173 << ((genFn0()) & 63))) | 1))
                    v206 := GenS33{F0: genFn0(), F1: v173}
                    _ = v206
                } else {
                    v207 := GenS14{F0: 688, F1: 2147483647, F2: v173}
                    if (0 == 1) {
                        var v208 int16 = (v4 & v4)
                        _ = v208
                    } else {
                        for fi209 := 0; fi209 < 16; fi209++ {
                            v173 = -128
                            var v210 int8 = -19
                            _ = fi209
                            _ = v210
                        }
                    }
                    _ = v207
                }
            } else {
                v211 := GenS10{F0: (-(v173)), F1: 1}
                _ = v211
            }
        }
    } else {
        if true {
            fmt.Println(v4, v173, v173)
        }
        var v212 int32 = -655
        v173 -= v173
        _ = v212
    }
    var v213 float32 = ((0.0 / (1e10 + (0.0 * (((0.1 * (3.14159 / ((float32((3.14159 * float64(float32(1.5)))) / 1.0) - (1.5 + 1e10)))) / 1.5) / 3.14159)))) / 1.5)
    fmt.Println(v2, v173, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v173:", v173)
    fmt.Println("v2:", v2)
    fmt.Println("v213:", v213)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v173
    _ = v2
    _ = v213
    _ = v3
    _ = v4
}
