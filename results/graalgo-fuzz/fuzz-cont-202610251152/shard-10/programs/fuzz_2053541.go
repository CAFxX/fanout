package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int64 }
type GenS2 struct { F0 int; F1 int16; F2 int16 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 int16; F2 int32 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 uint64; F1 uint64 }
type GenS9 struct { F0 uint16; F1 uint8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int16; F1 uint8; F2 bool }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint64; F1 int8 }
type GenS14 struct { F0 int16; F1 uint32 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int16 }
func genFn0(p0 int32) uint {
    var v6 string = "hello"
    if false {
        v6 = (("hello" + "hello") + "foo bar")
        v6 = ""
    }
    var v7 int8 = -118
    var v8 int = -646
    v7 = v7
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    var v3 uint = 42
    var v4 int32 = int32(v3)
    var v5 bool = (!true)
    for fi9 := 0; fi9 < 8; fi9++ {
        if v5 {
            for fi10 := 0; fi10 < 14; fi10++ {
                v3 -= v3
                _ = fi10
            }
            v4 -= v4
            v11 := GenS1{F0: v3, F1: 128, F2: -579}
            _ = v11
        } else {
            v12 := GenS2{F0: int(v4), F1: -32768, F2: -746}
            v5 = (0 > -620)
            v3 += (v3 / ((255) | 1))
            _ = v12
        }
        if (1 != uint16(v3)) {
            switch (fi9 * fi9) {
            case 769:
                fmt.Println(v3)
            case 542:
                for fi13 := 0; fi13 < 1; fi13++ {
                    v4 -= (v4 ^ v4)
                    if true {
                        if v5 {
                            var v14 uint = (v3 % (((v3 % ((384) | 1))) | 1))
                            _ = v14
                        } else {
                            _ = len(v1)
                            v3 = v3
                        }
                        fmt.Println(v5, fi9)
                        switch fi13 {
                        case 42:
                            switch ("hello" + "a") {
                            case "":
                                v5 = v5
                            }
                        default:
                            if true {
                                if (false && true) {
                                    if (0 >= 785) {
                                        if v5 {
                                            fmt.Println(fi9, fi13, fi13)
                                            fmt.Println(fi9)
                                            v15 := GenS1{F0: genFn0(42), F1: 217, F2: 0}
                                            _ = v15
                                        } else {
                                            fmt.Println(v5, v5)
                                            v5 = v5
                                            v16 := v1[65535]
                                            _ = v16
                                        }
                                        delete(v1, 2)
                                        fmt.Println(fi13, v5)
                                    } else {
                                        var v17 uint32 = 0
                                        var v18 int8 = -128
                                        fmt.Println(v4)
                                        _ = v17
                                        _ = v18
                                    }
                                } else {
                                    v19 := GenS3{F0: 0, F1: 42, F2: 128}
                                    var v20 uint = 4294967295
                                    _ = v19
                                    _ = v20
                                }
                                var v21 uint16 = 706
                                var v22 uint16 = (^(v21))
                                _ = v21
                                _ = v22
                            } else {
                                var v23 float32 = 0.0
                                _ = v23
                            }
                        }
                    } else {
                        v4 -= v4
                        var v24 int16 = 1
                        v3 += 42
                        _ = v24
                    }
                    _ = fi13
                }
            }
            if (genFn0((v4 % ((2147483647) | 1))) > v3) {
                for fi25 := 0; fi25 < 1; fi25++ {
                    fmt.Println(v4)
                    if false {
                        var v26 int = 128
                        var v27 int32 = (v4 & v4)
                        fmt.Println(v5, fi25)
                        v4 = -451
                        if v5 {
                            var v28 float32 = 0.1
                            _ = v28
                        }
                        _ = v26
                        _ = v27
                    } else {
                        v5 = (900 > 42)
                        fmt.Println(v4, v4)
                        fmt.Println(v5, v5, v3)
                        v4 = (^(v4))
                        if (((^(v3)) < (^(v3))) && ("a" <= ("foo bar" + ""))) {
                            switch "foo bar" {
                            case "a":
                                switch -1 {
                                case 127:
                                    v5 = (-77 > 1)
                                case -76:
                                    if false {
                                        var v29 float64 = ((-2.25 / 1.0) - 1.5)
                                        v4 += v4
                                        _ = v29
                                    } else {
                                        fmt.Println(v5)
                                        fmt.Println(v5, fi25)
                                    }
                                }
                            case "\n\t":
                                var v30 uint64 = 566
                                _ = v30
                            default:
                                var v31 uint32 = 121
                                _ = v31
                            }
                            if v5 {
                                var v32 uint8 = 2
                                if (true && (uint32(v32) > 2147483648)) {
                                    var v33 uint16 = 42
                                    _ = v33
                                }
                                v34 := GenS3{F0: 2, F1: -1, F2: v32}
                                if (-2.25 > ((3.14159 * 0.0) - (float32((3.14159 + float64(1e10))) / (float32((1.5 - (0.0 - 1e10))) * 1.5)))) {
                                    var v35 int64 = 42
                                    v36 := []int{fi9}
                                    var v37 float32 = float32(3.14159)
                                    fmt.Println(v37, v4)
                                    var v38 float64 = (1.5 - (0.0 + 0.0))
                                    if v5 {
                                        var v39 int8 = 100
                                        var v40 uint8 = v32
                                        _ = len(v36)
                                        _ = v39
                                        _ = v40
                                    } else {
                                        switch (fi9 << ((genFn0(42)) & 63)) {
                                        case 1:
                                            _ = len(v36)
                                        case 32767:
                                            fmt.Println(v5, v3, v37)
                                            var v41 int = fi9
                                            fmt.Println(v37, v38, v35)
                                            _ = v41
                                        case 2147483647:
                                            v37 -= (3.14159 - (v37 / (-2.25 * 1.5)))
                                        default:
                                            var v42 int64 = (v35 ^ (v35 / ((v35) | 1)))
                                            _ = v42
                                        }
                                        v35 = (v35 * v35)
                                        if len(v36) >= 2 {
                                            v36 = v36[1:len(v36)]
                                        }
                                        delete(v2, 1)
                                        switch "" {
                                        case "hello":
                                            v37 = (0.0 / (1e10 + -2.25))
                                        case "\n\t":
                                            fmt.Println(fi9, v35)
                                        }
                                    }
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                                _ = v32
                                _ = v34
                            }
                        } else {
                            v43 := []int{fi9, fi9, (fi9 | (fi25 * fi25)), fi9}
                            fmt.Println(v5)
                            _ = v43
                        }
                    }
                    v2[(fi25 << (((v3 & v3)) & 63))] = fi25
                    _ = fi25
                }
            } else {
                if false {
                    switch "hello" {
                    case "a":
                        v3 += v3
                    case "\n\t":
                        if (33 == 0) {
                            v2[4294967295] = (fi9 + fi9)
                        }
                    case "":
                        for fi44 := 0; fi44 < 1; fi44++ {
                            switch (-(fi9)) {
                            case -9223372036854775808:
                                var v45 float32 = 1e10
                                _ = v45
                            case 32768:
                                var v46 uint = v3
                                _ = v46
                            }
                            _ = fi44
                        }
                    default:
                        var v47 uint8 = 1
                        _ = v47
                    }
                    fmt.Println(v4, v3, fi9)
                    if v5 {
                        if v5 {
                            if (!v5) {
                                for fi48 := 0; fi48 < 2; fi48++ {
                                    if (!(v5 || false)) {
                                        var v49 uint8 = 1
                                        _ = v49
                                    }
                                    v3 = genFn0((^(v4)))
                                    switch (-(fi9)) {
                                    case 4294967295:
                                        if v5 {
                                            switch "foo bar" {
                                            case "\n\t":
                                                if (("foo bar" > ("hello" + "a")) || (0.0 <= 1e10)) {
                                                    v4 -= (v4 | 42)
                                                    fmt.Println(fi9, v5)
                                                    v50 := GenS4{F0: ("foo bar" != "\n\t")}
                                                    _ = v50
                                                }
                                            case "foo bar":
                                                var v51 bool = true
                                                _ = v51
                                            }
                                            v52 := []int{fi9, fi48}
                                            var v53 float64 = 3.14159
                                            _ = v52
                                            _ = v53
                                        }
                                    case -892:
                                        v54 := []int{}
                                        _ = v54
                                    case 571:
                                        fmt.Println(fi48, v5)
                                        v55 := GenS1{F0: (v3 ^ (v3 << ((42) & 63))), F1: 197, F2: 0}
                                        v4 = (v4 - v4)
                                        _ = v55
                                    }
                                    _ = fi48
                                }
                            }
                            if v5 {
                                switch fi9 {
                                case 2147483647:
                                    v3 += 962
                                default:
                                    var v56 int = 127
                                    _ = v56
                                }
                                fmt.Println(v5)
                                fmt.Println(v5, v3)
                            } else {
                                for fi57 := 0; fi57 < 13; fi57++ {
                                    var v58 uint32 = 0
                                    v4 = v4
                                    v59 := []int{(^(fi57)), 4294967295, fi57}
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                }
                                fmt.Println(v5)
                            }
                            var v60 int32 = -29
                            _ = v60
                        } else {
                            switch 1 {
                            case -2147483648:
                                switch fi9 {
                                case 2147483647:
                                    v5 = ((-9223372036854775808 == (fi9 ^ (fi9 - int(v3)))) && ((("" != "\n\t") || false) || true))
                                }
                            default:
                                var v61 bool = (((((v5 || v5) && v5) || (749 >= 1)) || true) && ((30 > 18446744073709551615) && v5))
                                _ = v61
                            }
                            v5 = (v5 && (!(false && v5)))
                        }
                        v62 := GenS1{F0: (v3 + (v3 & (v3 % (((v3 & (v3 / (((^(v3))) | 1)))) | 1)))), F1: 19, F2: -9223372036854775808}
                        _ = len(v2)
                        _ = v62
                    } else {
                        if ((!v5) && (true || (!v5))) {
                            fmt.Println(v3, v3, v4)
                            var v63 float32 = (1e10 - -2.25)
                            v64 := GenS4{F0: v5}
                            _ = v63
                            _ = v64
                        }
                        fmt.Println(v5, fi9, v3)
                        var v65 float64 = (-2.25 / -2.25)
                        _ = v65
                    }
                }
                v4 += v4
            }
            _ = len(v1)
        }
        _ = fi9
    }
    var v66 int64 = -1
    v66 -= (v66 << (((v3 | 65535)) & 63))
    if v5 {
        if (128 >= 135) {
            if (v5 || (3.14159 <= ((3.14159 + -2.25) - ((float64((((-2.25 * float32(-2.25)) - -2.25) - 0.0)) + 0.1) / 1e10)))) {
                switch ("" + "foo bar") {
                case "\n\t":
                    v3 = (v3 >> ((42) & 63))
                case "":
                    var v67 bool = (42 > 42)
                    _ = v67
                case "a":
                    v5 = v5
                }
                var v68 uint = v3
                _ = v68
            } else {
                switch ("a" + "foo bar") {
                case "":
                    switch "\n\t" {
                    case "\n\t":
                        switch "" {
                        case "\n\t":
                            v5 = ((uint32(v3) == 4294967295) && false)
                        case "foo bar":
                            v69 := []int{2147483647, 534, 1, 129}
                            _ = v69
                        default:
                            if (!(204 < 2)) {
                                v70 := GenS5{F0: 1, F1: 32767, F2: (v4 + v4)}
                                switch 32767 {
                                case 4294967295:
                                    var v71 uint64 = 42
                                    _ = v71
                                case -1:
                                    if (4294967295 == 32768) {
                                        for fi72 := 0; fi72 < 13; fi72++ {
                                            var v73 float32 = (0.1 / 1e10)
                                            _ = fi72
                                            _ = v73
                                        }
                                        v74 := []int{}
                                        v4 = (v4 * v4)
                                        _ = v74
                                    } else {
                                        v5 = (-9223372036854775808 > (v66 | v66))
                                    }
                                default:
                                    for fi75 := 0; fi75 < 10; fi75++ {
                                        for fi76 := 0; fi76 < 13; fi76++ {
                                            v5 = v5
                                            var v77 uint32 = 2
                                            if (42 >= -32768) {
                                                if v5 {
                                                    var v78 uint32 = 4294967295
                                                    var v79 float32 = float32(1.5)
                                                    _ = v78
                                                    _ = v79
                                                } else {
                                                    if v5 {
                                                        v80 := []int{}
                                                        fmt.Println(fi75, v3, v77)
                                                        v3 += v3
                                                        if (v5 && true) {
                                                            if v5 {
                                                                if (0.0 >= 0.0) {
                                                                    var v81 int8 = -128
                                                                    _ = v81
                                                                }
                                                            } else {
                                                                fmt.Println(v3)
                                                                v82 := GenS5{F0: 2, F1: -32768, F2: (^(v4))}
                                                                var v83 uint32 = (^(v77))
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                            fmt.Println(v3)
                                                            v84 := GenS6{F0: (-(v66)), F1: (v77 % ((42) | 1))}
                                                            _ = v84
                                                        } else {
                                                            v85 := GenS7{F0: uint8(v3), F1: ((42 <= 57) || v5)}
                                                            var v86 float32 = 1.5
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                        _ = v80
                                                    } else {
                                                        var v87 int64 = v66
                                                        v88 := GenS8{F0: 18446744073709551615, F1: 4294967295}
                                                        v89 := GenS7{F0: 120, F1: v5}
                                                        _ = v87
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                    v90 := v2[9223372036854775807]
                                                    if (v4 == (v4 & (v4 - (^(v4))))) {
                                                        v91 := []int{(fi76 / ((v90) | 1)), (^(fi75))}
                                                        fmt.Println(v3, v66)
                                                        v90 -= v90
                                                        _ = v91
                                                    } else {
                                                        switch fi76 {
                                                        case 790:
                                                            v92 := GenS9{F0: 42, F1: 1}
                                                            v93 := GenS6{F0: (^(v66)), F1: 65535}
                                                            _ = v92
                                                            _ = v93
                                                        default:
                                                            v90 += fi75
                                                        }
                                                        var v94 int8 = 0
                                                        _ = v94
                                                    }
                                                    v95 := GenS8{F0: uint64(v3), F1: 0}
                                                    _ = v90
                                                    _ = v95
                                                }
                                                v77 += uint32(v3)
                                            }
                                            _ = fi76
                                            _ = v77
                                        }
                                        for fi96 := 0; fi96 < 12; fi96++ {
                                            v97 := []int{}
                                            _ = fi96
                                            _ = v97
                                        }
                                        var v98 string = ("foo bar" + ("a" + "hello"))
                                        _ = fi75
                                        _ = v98
                                    }
                                }
                                _ = v70
                            } else {
                                for fi99 := 0; fi99 < 11; fi99++ {
                                    switch (fi99 ^ 415) {
                                    case 321:
                                        fmt.Println(fi99)
                                    case 32768:
                                        v66 -= (v66 ^ 658)
                                    case 42:
                                        switch 554 {
                                        case 9223372036854775807:
                                            v3 = (^(v3))
                                        default:
                                            var v100 uint64 = 298
                                            _ = v100
                                        }
                                    default:
                                        switch ("hello" + "\n\t") {
                                        case "\n\t":
                                            if false {
                                                switch fi99 {
                                                case 128:
                                                    v66 = v66
                                                case -2147483648:
                                                    v66 -= v66
                                                }
                                                for fi101 := 0; fi101 < 18; fi101++ {
                                                    if v5 {
                                                        v102 := GenS2{F0: (fi99 << (((^(v3))) & 63)), F1: 42, F2: -32768}
                                                        if (float64(3.14159) == 1e10) {
                                                            v1[574] = 263
                                                            v103 := GenS10{F0: (-(v4))}
                                                            v104 := GenS7{F0: 255, F1: (uint16(v3) < 32767)}
                                                            _ = v103
                                                            _ = v104
                                                        } else {
                                                            v66 = (v66 ^ (v66 << ((v3) & 63)))
                                                            switch "foo bar" {
                                                            case "a":
                                                                v105 := GenS5{F0: 2, F1: 42, F2: v4}
                                                                _ = v105
                                                            case "hello":
                                                                v4 = (v4 >> ((genFn0(0)) & 63))
                                                            case "foo bar":
                                                                v106 := GenS6{F0: v66, F1: 541}
                                                                _ = v106
                                                            }
                                                        }
                                                        _ = v102
                                                    }
                                                    _ = fi101
                                                }
                                                switch "\n\t" {
                                                case "a":
                                                    switch ("foo bar" + "hello") {
                                                    case "foo bar":
                                                        switch fi99 {
                                                        case 307:
                                                            for fi107 := 0; fi107 < 2; fi107++ {
                                                                var v108 float64 = (-2.25 * float64(0.1))
                                                                _ = fi107
                                                                _ = v108
                                                            }
                                                        case 127:
                                                            v3 += (v3 & v3)
                                                        case 1:
                                                            v66 = (v66 * (v66 | (v66 - 558)))
                                                        default:
                                                            var v109 uint8 = 1
                                                            _ = v109
                                                        }
                                                    case "\n\t":
                                                        for fi110 := 0; fi110 < 10; fi110++ {
                                                            _ = len(v2)
                                                            _ = fi110
                                                        }
                                                    default:
                                                        delete(v2, (fi99 / (((fi99 | (fi99 + fi99))) | 1)))
                                                    }
                                                default:
                                                    var v111 int = fi99
                                                    _ = v111
                                                }
                                            } else {
                                                if false {
                                                    v112 := GenS8{F0: 2, F1: 856}
                                                    var v113 uint64 = 2
                                                    _ = v112
                                                    _ = v113
                                                }
                                                for fi114 := 0; fi114 < 12; fi114++ {
                                                    if ((65535 < 2147483647) && v5) {
                                                        fmt.Println(v66, v4, v3)
                                                        if (0 < v3) {
                                                            v66 = (v66 << (((v3 >> (((v3 % ((v3) | 1))) & 63))) & 63))
                                                            _ = len(v2)
                                                        }
                                                        var v115 int64 = v66
                                                        _ = v115
                                                    } else {
                                                        v116 := []int{(fi99 | (-(fi114))), (fi114 + 101)}
                                                        fmt.Println(v66)
                                                        _ = v116
                                                    }
                                                    var v117 int64 = (v66 % ((v66) | 1))
                                                    _ = fi114
                                                    _ = v117
                                                }
                                                v118 := GenS11{F0: -32768, F1: 206, F2: v5}
                                                _ = v118
                                            }
                                        case "":
                                            for fi119 := 0; fi119 < 4; fi119++ {
                                                v4 -= 921
                                                _ = fi119
                                            }
                                        case "foo bar":
                                            v3 += (v3 - (^(v3)))
                                        }
                                    }
                                    v5 = v5
                                    _ = fi99
                                }
                                for fi120 := 0; fi120 < 13; fi120++ {
                                    fmt.Println(v5, v4)
                                    fmt.Println(v3)
                                    if false {
                                        v5 = v5
                                        var v121 int8 = 42
                                        _ = v121
                                    } else {
                                        v122 := v2[fi120]
                                        _ = v122
                                    }
                                    _ = fi120
                                }
                            }
                        }
                    }
                case "a":
                    var v123 int8 = 1
                    _ = v123
                }
                v124 := []int{9223372036854775807, 32768, -708, 127}
                _ = v124
            }
        } else {
            v125 := []int{67, 127, int(v66)}
            v126 := GenS12{F0: (((3.14159 + (float32(0.1) * -2.25)) + 0.0) >= 1.5)}
            _ = v125
            _ = v126
        }
        switch 127 {
        case 902:
            var v127 uint16 = 0
            _ = v127
        case 396:
            var v128 string = "a"
            _ = v128
        case -773:
            v4 = (v4 & (v4 * (^(v4))))
        default:
            v129 := GenS12{F0: (v5 && v5)}
            _ = v129
        }
        fmt.Println(v66)
    }
    delete(v1, 65535)
    fmt.Println(v3, v5, v3)
    if (!(genFn0((v4 - (^(v4)))) > v3)) {
        v5 = ("a" >= "foo bar")
    } else {
        v130 := GenS13{F0: 18446744073709551615, F1: 1}
        var v131 float32 = (1.5 + -2.25)
        for fi132 := 0; fi132 < 15; fi132++ {
            var v133 uint8 = 127
            _ = fi132
            _ = v133
        }
        _ = v130
        _ = v131
    }
    for fi134 := 0; fi134 < 18; fi134++ {
        if (-1 <= v66) {
            if (v3 != (v3 >> (((v3 % (((v3 / (((v3 / ((879) | 1))) | 1))) | 1))) & 63))) {
                v66 += v66
                v5 = true
            }
            var v135 uint64 = 249
            _ = v135
        }
        var v136 float32 = 3.14159
        _ = fi134
        _ = v136
    }
    switch "\n\t" {
    case "foo bar":
        for fi137 := 0; fi137 < 8; fi137++ {
            var v138 uint8 = 148
            v139 := []int{-946, (fi137 - fi137), (fi137 / ((fi137) | 1))}
            if v5 {
                fmt.Println(v3, v3)
                v139 = append(v139, -2147483648)
                var v140 string = ""
                _ = v140
            } else {
                if v5 {
                    v141 := GenS9{F0: uint16(v3), F1: 2}
                    v142 := GenS14{F0: 1, F1: 2147483648}
                    if true {
                        fmt.Println(v5, v138)
                        v2[fi137] = int(v138)
                        for fi143 := 0; fi143 < 14; fi143++ {
                            fmt.Println(v5, v138)
                            v4 += (v4 >> ((529) & 63))
                            if (v138 <= (v138 & 255)) {
                                if v5 {
                                    v138 = v138
                                } else {
                                    v5 = ((831 >= 2) && v5)
                                    if len(v139) >= 2 {
                                        v139 = v139[1:len(v139)]
                                    }
                                    fmt.Println(fi137, v4)
                                }
                            } else {
                                v66 += v66
                                fmt.Println(v138, v3)
                                fmt.Println(v4, fi143, v138)
                            }
                            _ = fi143
                        }
                    }
                    _ = v141
                    _ = v142
                } else {
                    v139 = append(v139, fi137)
                    var v144 int = int(v138)
                    _ = v144
                }
            }
            _ = fi137
            _ = v138
            _ = v139
        }
    case "a":
        v145 := GenS15{F0: (v3 << (((v3 / ((89) | 1))) & 63))}
        _ = v145
    }
    switch -2147483648 {
    case -9223372036854775808:
        v146 := []int{}
        _ = v146
    case -2147483648:
        var v147 int16 = 969
        _ = v147
    case 0:
        if false {
            v148 := []int{}
            _ = v148
        }
    }
    var v149 uint = v3
    if v5 {
        v3 = genFn0(0)
    } else {
        if (796 != 844) {
            fmt.Println(v66, v4, v5)
        } else {
            for fi150 := 0; fi150 < 14; fi150++ {
                for fi151 := 0; fi151 < 7; fi151++ {
                    v149 = 2
                    v152 := GenS16{F0: -32768}
                    if true {
                        v4 = (v4 | (^(v4)))
                        v153 := GenS5{F0: 255, F1: 42, F2: v4}
                        _ = v153
                    }
                    _ = fi151
                    _ = v152
                }
                _ = fi150
            }
        }
        fmt.Println(v149)
        v3 = (v3 ^ v149)
    }
    for fi154 := 0; fi154 < 8; fi154++ {
        if (-32768 == -368) {
            switch (fi154 + 9223372036854775807) {
            case -418:
                v5 = v5
            }
            var v155 float64 = 1e10
            _ = v155
        }
        _ = fi154
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v149:", v149)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v149
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
}
