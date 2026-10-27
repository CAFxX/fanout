package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS4 struct { F0 int8; F1 string; F2 uint16 }
type GenS5 struct { F0 uint64; F1 uint; F2 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint; F1 int8; F2 uint }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int16 }

func main() {
    var v1 uint16 = 199
    var v2 string = ("hello" + ("" + ""))
    var v3 int32 = 42
    var v4 uint = 102
    var v5 uint64 = 4294967295
    for fi6 := 0; fi6 < 17; fi6++ {
        switch v2 {
        case "hello":
            switch ("" + v2) {
            case "\n\t":
                v7 := GenS1{F0: 2147483647, F1: (v5 + uint64(v4)), F2: uint16(v4)}
                _ = v7
            case "":
                fmt.Println(v5, v3, v2)
            }
        case "foo bar":
            v8 := []int{}
            _ = v8
        case "\n\t":
            fmt.Println(v3, v4)
        default:
            v5 += (v5 * v5)
        }
        _ = fi6
    }
    fmt.Println(v4)
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v5, v2, v3)
        v10 := GenS2{F0: 4294967295, F1: true, F2: uint8(v4)}
        var v11 int8 = 1
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint32 = 2147483647
    var v13 float64 = 3.14159
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(v3, v4)
        switch fi14 {
        case 127:
            var v15 uint64 = 837
            _ = v15
        case -118:
            switch v2 {
            case "foo bar":
                var v16 int32 = v3
                _ = v16
            }
        }
        _ = fi14
    }
    switch 1 {
    case -313:
        switch 0 {
        case 4294967295:
            if true {
                if true {
                    for fi17 := 0; fi17 < 10; fi17++ {
                        for fi18 := 0; fi18 < 1; fi18++ {
                            v19 := []int{fi17, (fi18 * (fi18 >> (((v4 >> (((^(v4))) & 63))) & 63)))}
                            fmt.Println(v4, v4)
                            _ = fi18
                            _ = v19
                        }
                        v13 -= 1e10
                        _ = fi17
                    }
                    v1 = (v1 / (((v1 * v1)) | 1))
                    var v20 int16 = 42
                    _ = v20
                } else {
                    v21 := GenS2{F0: (v12 << (((^(v4))) & 63)), F1: (!true), F2: 1}
                    var v22 uint = (^(v4))
                    _ = v21
                    _ = v22
                }
                var v23 string = "hello"
                fmt.Println(v23, v3, v2)
                _ = v23
            } else {
                v24 := []int{42}
                fmt.Println(v5)
                if true {
                    fmt.Println(v3, v1, v1)
                    fmt.Println(v13)
                } else {
                    switch v2 {
                    case "foo bar":
                        for fi25 := 0; fi25 < 7; fi25++ {
                            var v26 uint8 = 0
                            if len(v24) > 0 {
                                v24[0] = (-(fi25))
                            }
                            for fi27 := 0; fi27 < 14; fi27++ {
                                var v28 int64 = 9223372036854775807
                                _ = fi27
                                _ = v28
                            }
                            _ = fi25
                            _ = v26
                        }
                    case "":
                        if false {
                            v12 = (v12 - (v12 ^ v12))
                            v13 += (((0.0 - 0.0) * v13) - v13)
                            _ = len(v24)
                        }
                    default:
                        fmt.Println(v12, v1)
                    }
                    fmt.Println(v4, v4)
                }
                _ = v24
            }
        default:
            switch 997 {
            case -2147483648:
                for fi29 := 0; fi29 < 1; fi29++ {
                    var v30 int8 = -1
                    v31 := GenS1{F0: v3, F1: v5, F2: (v1 | 42)}
                    if true {
                        switch v2 {
                        case "\n\t":
                            for fi32 := 0; fi32 < 9; fi32++ {
                                switch v2 {
                                case "hello":
                                    if false {
                                        v33 := []int{int(v3), fi32, 93, (^(fi32))}
                                        if true {
                                            v34 := GenS3{F0: v1, F1: -9223372036854775808, F2: uint32(v5)}
                                            if false {
                                                v12 -= (^(v12))
                                            } else {
                                                var v35 int16 = -900
                                                fmt.Println(v1, fi29)
                                                switch int(v3) {
                                                case -2147483648:
                                                    var v36 float32 = -2.25
                                                    _ = v36
                                                case -9223372036854775808:
                                                    v37 := GenS1{F0: v3, F1: 18446744073709551615, F2: 42}
                                                    _ = v37
                                                }
                                                _ = v35
                                            }
                                            var v38 uint = 18446744073709551615
                                            _ = v34
                                            _ = v38
                                        } else {
                                            v39 := GenS2{F0: v12, F1: false, F2: 102}
                                            var v40 int32 = (v3 ^ (v3 >> (((v4 - (v4 % ((65535) | 1)))) & 63)))
                                            var v41 uint16 = (v1 << ((v4) & 63))
                                            switch v2 {
                                            case "a":
                                                v12 += v12
                                            case "\n\t":
                                                if len(v33) > 0 {
                                                    v33[0] = fi32
                                                }
                                            }
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                        switch fi32 {
                                        case 9223372036854775807:
                                            if len(v33) > 0 {
                                                v33[0] = (fi29 >> (((v4 / (((^(v4))) | 1))) & 63))
                                            }
                                            v12 -= v12
                                            _ = len(v33)
                                            var v42 uint16 = (v1 * (v1 - 65535))
                                            if (int8(v3) < 127) {
                                                var v43 uint8 = 1
                                                _ = v43
                                            }
                                            _ = v42
                                        case 0:
                                            if ((1e10 == (3.14159 - 3.14159)) && false) {
                                                v44 := GenS1{F0: v3, F1: (v5 * v5), F2: (^(v1))}
                                                _ = v44
                                            }
                                            v3 -= v3
                                            var v45 uint8 = 1
                                            _ = v45
                                        case 4294967295:
                                            if (!((v3 >> ((2) & 63)) <= (v3 >> ((v4) & 63)))) {
                                                v13 = v13
                                            }
                                        default:
                                            if len(v33) >= 2 {
                                                v33 = v33[1:len(v33)]
                                            }
                                            v30 -= (v30 - v30)
                                        }
                                        _ = v33
                                    }
                                case "\n\t":
                                    v46 := []int{}
                                    _ = v46
                                case "a":
                                    switch -2147483648 {
                                    case 4294967295:
                                        v30 = (v30 - v30)
                                    case -2147483648:
                                        switch v2 {
                                        case "":
                                            v12 = (^(v12))
                                        default:
                                            var v47 int = fi32
                                            _ = v47
                                        }
                                    case 0:
                                        v3 = -20
                                    default:
                                        var v48 uint = (v4 ^ v4)
                                        _ = v48
                                    }
                                }
                                v49 := []int{fi32, int(v12), (^(fi32)), fi32}
                                fmt.Println(fi32)
                                _ = fi32
                                _ = v49
                            }
                        case "":
                            v2 = ((("\n\t" + "a") + "\n\t") + v2)
                        case "a":
                            var v50 uint = uint(v1)
                            _ = v50
                        default:
                            var v51 uint = (^(v4))
                            _ = v51
                        }
                        v1 += v1
                        var v52 uint32 = v12
                        _ = v52
                    } else {
                        v53 := []int{(fi29 ^ 127)}
                        _ = v53
                    }
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            default:
                for fi54 := 0; fi54 < 2; fi54++ {
                    if (1 < -9223372036854775808) {
                        var v55 string = v2
                        var v56 int32 = (v3 + 1)
                        _ = v55
                        _ = v56
                    } else {
                        v1 += (v1 % (((v1 << ((0) & 63))) | 1))
                    }
                    _ = fi54
                }
            }
        }
    default:
        if false {
            var v57 float64 = v13
            _ = v57
        }
    }
    v58 := []int{9223372036854775807, -255, -549, 9223372036854775807}
    for fi59 := 0; fi59 < 3; fi59++ {
        v2 = ((((v2 + ((((v2 + ((("\n\t" + v2) + "a") + "\n\t")) + v2) + "foo bar") + v2)) + (v2 + v2)) + (((v2 + v2) + v2) + (("hello" + v2) + v2))) + v2)
        if false {
            if (true || false) {
                var v60 string = "a"
                if ((1.5 + v13) >= v13) {
                    v61 := GenS3{F0: (v1 << ((uint(v5)) & 63)), F1: 388, F2: uint32(v4)}
                    var v62 float32 = 1e10
                    _ = v61
                    _ = v62
                }
                _ = v60
            } else {
                switch (v2 + v2) {
                case "foo bar":
                    for fi63 := 0; fi63 < 15; fi63++ {
                        fmt.Println(fi63)
                        _ = fi63
                    }
                case "":
                    var v64 uint16 = (v1 % (((v1 >> ((255) & 63))) | 1))
                    _ = v64
                case "a":
                    var v65 float64 = ((v13 + 1.5) + v13)
                    _ = v65
                }
                var v66 uint16 = 0
                if (!(0 > 35)) {
                    v67 := GenS3{F0: (v66 / (((v66 % (((^(v66))) | 1))) | 1)), F1: -1, F2: (v12 % (((v12 << ((v4) & 63))) | 1))}
                    var v68 int = -2147483648
                    v69 := GenS2{F0: uint32(v1), F1: true, F2: 42}
                    _ = v67
                    _ = v68
                    _ = v69
                } else {
                    v2 = "a"
                    switch fi59 {
                    case 4294967295:
                        switch (fi59 >> (((v4 | v4)) & 63)) {
                        case 42:
                            var v70 uint32 = 2147483648
                            _ = v70
                        case -1:
                            var v71 bool = (true && true)
                            _ = v71
                        case 9223372036854775807:
                            var v72 string = "foo bar"
                            _ = v72
                        default:
                            v12 = (v12 & v12)
                        }
                    case 0:
                        if len(v58) > 0 {
                            v58[0] = fi59
                        }
                    default:
                        fmt.Println(v5)
                    }
                }
                _ = v66
            }
            for fi73 := 0; fi73 < 12; fi73++ {
                var v74 bool = false
                v75 := GenS3{F0: (v1 - 889), F1: -1, F2: v12}
                fmt.Println(fi59, v3, v4)
                _ = fi73
                _ = v74
                _ = v75
            }
            var v76 int8 = 0
            _ = v76
        }
        _ = len(v58)
        _ = fi59
    }
    fmt.Println(v2, v4, v5)
    fmt.Println(v12, v5, v5)
    for fi77 := 0; fi77 < 18; fi77++ {
        if ((((float32(3.14159) - 0.1) + (3.14159 / 1.0)) != (-2.25 * (1e10 - 0.1))) || false) {
            v78 := GenS1{F0: -2147483648, F1: 0, F2: 73}
            _ = v78
        } else {
            var v79 int8 = 100
            if false {
                v2 = ((v2 + (v2 + (v2 + (v2 + (((v2 + ((("" + ((v2 + ((((v2 + "hello") + (v2 + v2)) + "foo bar") + "a")) + (v2 + v2))) + "foo bar") + (v2 + "hello"))) + "\n\t") + v2))))) + "hello")
                v5 = v5
            } else {
                for fi80 := 0; fi80 < 11; fi80++ {
                    fmt.Println(v3, fi77, fi80)
                    switch "a" {
                    case "hello":
                        var v81 int32 = v3
                        _ = v81
                    case "\n\t":
                        fmt.Println(v1, v13, v4)
                    default:
                        fmt.Println(v12)
                    }
                    if (31 >= uint8(v1)) {
                        var v82 uint32 = (v12 * v12)
                        _ = v82
                    } else {
                        switch v2 {
                        case "foo bar":
                            var v83 float32 = ((float32(float64((-2.25 + -2.25))) * ((1.5 - 0.1) * (float32(0.1) + 0.1))) * 1e10)
                            var v84 float64 = 3.14159
                            if (42 == -648) {
                                var v85 int8 = 77
                                _ = v85
                            } else {
                                v86 := GenS2{F0: v12, F1: true, F2: 0}
                                if true {
                                    switch fi77 {
                                    case 320:
                                        _ = len(v58)
                                    case -695:
                                        if (!false) {
                                            switch (-(fi80)) {
                                            case 32767:
                                                var v87 string = v2
                                                _ = v87
                                            case 42:
                                                v88 := GenS4{F0: v79, F1: v2, F2: uint16(v12)}
                                                _ = v88
                                            }
                                        } else {
                                            v89 := GenS5{F0: (v5 ^ v5), F1: v4, F2: 42}
                                            if ("foo bar" > (((v2 + "") + "foo bar") + (((((v2 + v2) + ("foo bar" + v2)) + ("\n\t" + v2)) + (v2 + v2)) + ((v2 + v2) + ("hello" + v2))))) {
                                                var v90 float32 = (-2.25 / v83)
                                                _ = len(v58)
                                                _ = v90
                                            } else {
                                                var v91 uint8 = 35
                                                v92 := GenS6{F0: v79}
                                                v91 = uint8(v1)
                                                _ = v91
                                                _ = v92
                                            }
                                            if (127 <= v79) {
                                                v13 = 3.14159
                                                v93 := GenS2{F0: v12, F1: (true || true), F2: 128}
                                                var v94 uint32 = v12
                                                _ = v93
                                                _ = v94
                                            }
                                            _ = v89
                                        }
                                    default:
                                        var v95 uint32 = v12
                                        _ = v95
                                    }
                                }
                                _ = v86
                            }
                            _ = v83
                            _ = v84
                        }
                    }
                    _ = fi80
                }
                var v96 int32 = v3
                for fi97 := 0; fi97 < 16; fi97++ {
                    fmt.Println(v3, v1)
                    var v98 int32 = v96
                    _ = fi97
                    _ = v98
                }
                _ = v96
            }
            if (963 != 0) {
                v2 = ((v2 + "") + ("hello" + "hello"))
                if false {
                    _ = len(v58)
                    if false {
                        fmt.Println(v3, fi77)
                        var v99 uint16 = 32767
                        _ = v99
                    } else {
                        if (v3 > v3) {
                            if (9223372036854775807 <= 1) {
                                fmt.Println(fi77, v5)
                                if false {
                                    var v100 int64 = -1
                                    switch (fi77 - 127) {
                                    case 0:
                                        switch fi77 {
                                        case -9223372036854775808:
                                            v58 = append(v58, 4294967295)
                                        case 794:
                                            for fi101 := 0; fi101 < 8; fi101++ {
                                                if false {
                                                    if (v79 == (v79 % ((v79) | 1))) {
                                                        if (0 <= 42) {
                                                            fmt.Println(v4, v3)
                                                            var v102 uint8 = 37
                                                            _ = v102
                                                        } else {
                                                            switch (^(fi101)) {
                                                            case 4294967295:
                                                                var v103 uint32 = (v12 ^ (v12 & (^(v12))))
                                                                _ = v103
                                                            case -916:
                                                                switch (-(fi77)) {
                                                                case -2147483648:
                                                                    if len(v58) > 0 {
                                                                        v58[0] = (fi77 - -1)
                                                                    }
                                                                case 42:
                                                                    if true {
                                                                        fmt.Println(v4)
                                                                    } else {
                                                                        fmt.Println(fi101, v4)
                                                                        var v104 uint16 = 2
                                                                        _ = v104
                                                                    }
                                                                    switch -140 {
                                                                    case 1:
                                                                        var v105 int64 = 9223372036854775807
                                                                        _ = v105
                                                                    case 32768:
                                                                        v4 -= v4
                                                                    case -9223372036854775808:
                                                                        switch (v2 + ((v2 + "hello") + v2)) {
                                                                        case "\n\t":
                                                                            var v106 float64 = 3.14159
                                                                            _ = v106
                                                                        case "hello":
                                                                            var v107 string = ("\n\t" + v2)
                                                                            _ = v107
                                                                        default:
                                                                            _ = len(v58)
                                                                            var v108 int32 = 42
                                                                            _ = v108
                                                                        }
                                                                    }
                                                                    if (v1 > (v1 ^ (v1 & v1))) {
                                                                        v1 += 42
                                                                        var v109 float32 = 3.14159
                                                                        _ = v109
                                                                    }
                                                                default:
                                                                    var v110 uint16 = v1
                                                                    _ = v110
                                                                }
                                                            default:
                                                                v111 := GenS7{F0: (v12 % ((1) | 1)), F1: 0, F2: v5}
                                                                _ = v111
                                                            }
                                                            var v112 uint32 = 806
                                                            _ = v112
                                                        }
                                                        v113 := GenS8{F0: ((v2 + "hello") + (v2 + v2))}
                                                        v114 := GenS9{F0: 1}
                                                        _ = v113
                                                        _ = v114
                                                    }
                                                    var v115 float64 = (0.0 / v13)
                                                    if true {
                                                        v116 := GenS3{F0: (v1 - (v1 & v1)), F1: v100, F2: uint32(v1)}
                                                        v58 = append(v58, (fi101 / ((fi77) | 1)))
                                                        v1 += uint16(v4)
                                                        _ = v116
                                                    }
                                                    _ = v115
                                                }
                                                v2 = "foo bar"
                                                v1 -= v1
                                                _ = fi101
                                            }
                                        default:
                                            if (true && false) {
                                                if true {
                                                    if true {
                                                        v2 = v2
                                                    }
                                                    for fi117 := 0; fi117 < 6; fi117++ {
                                                        var v118 uint = 894
                                                        v119 := GenS10{F0: true}
                                                        v13 -= v13
                                                        switch fi117 {
                                                        case -1:
                                                            fmt.Println(v13, v3)
                                                        case 438:
                                                            v120 := GenS5{F0: 582, F1: 1, F2: (fi117 % ((int(v100)) | 1))}
                                                            _ = v120
                                                        }
                                                        _ = fi117
                                                        _ = v118
                                                        _ = v119
                                                    }
                                                    for fi121 := 0; fi121 < 7; fi121++ {
                                                        fmt.Println(fi77)
                                                        v122 := GenS11{F0: v4, F1: 127, F2: 255}
                                                        _ = fi121
                                                        _ = v122
                                                    }
                                                }
                                            }
                                        }
                                    default:
                                        var v123 uint32 = 859
                                        _ = v123
                                    }
                                    _ = v100
                                }
                            }
                        }
                    }
                } else {
                    fmt.Println(v5, v5)
                    v79 = (v79 >> ((v4) & 63))
                }
                if (1.5 < 1e10) {
                    switch fi77 {
                    case -1:
                        var v124 int = fi77
                        _ = v124
                    case 0:
                        v125 := GenS12{F0: (v3 >> (((v4 & v4)) & 63))}
                        _ = v125
                    }
                    v4 = v4
                }
            }
            _ = v79
        }
        var v126 int32 = v3
        v127 := GenS4{F0: 78, F1: "", F2: 2}
        _ = fi77
        _ = v126
        _ = v127
    }
    if (int(v1) <= -542) {
        v128 := GenS13{F0: 32767}
        var v129 int64 = -175
        for fi130 := 0; fi130 < 1; fi130++ {
            var v131 int8 = -128
            _ = fi130
            _ = v131
        }
        _ = v128
        _ = v129
    }
    if len(v58) > 0 {
        v58[0] = -262
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v58:", len(v58), v58)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
}
