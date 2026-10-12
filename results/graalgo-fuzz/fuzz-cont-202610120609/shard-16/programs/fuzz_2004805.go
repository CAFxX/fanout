package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 string; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 int; F2 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 uint; F1 uint16; F2 int8 }
type GenS9 struct { F0 int8; F1 int64 }
type GenS10 struct { F0 bool; F1 uint8; F2 int8 }
type GenS11 struct { F0 int8; F1 uint16 }
type GenS12 struct { F0 uint8; F1 bool }
type GenS13 struct { F0 int32; F1 bool }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 bool; F1 uint }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 int16; F1 int32; F2 uint }
type GenS18 struct { F0 uint16; F1 string }
func genFn0(p0 int64) int64 {
    v7 := GenS1{F0: 802}
    var v8 int8 = 127
    var v9 bool = false
    v9 = false
    var v10 uint64 = 400
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}
func genFn1() uint {
    v11 := []int{127, 32768, -975, 861}
    switch ("\n\t" + ((("foo bar" + "hello") + ("hello" + "")) + "a")) {
    case "foo bar":
        var v12 float32 = float32(0.1)
        _ = v12
    case "hello":
        var v13 uint32 = 2147483647
        _ = v13
    case "a":
        var v14 bool = (("hello" > ("foo bar" + "a")) && (426 <= 42))
        _ = v14
    default:
        if len(v11) > 0 {
            v11[0] = 4294967295
        }
    }
    var v15 bool = true
    _ = v11
    _ = v15
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 42
    var v3 int8 = 0
    var v4 string = ""
    var v5 int = -1
    var v6 int = -7
    var v16 int8 = v3
    switch v6 {
    case 32767:
        v2 = -1
    case 32768:
        v5 += (v5 + v6)
    }
    v17 := []int{}
    var v18 uint32 = 0
    if (true || true) {
        switch "foo bar" {
        case "hello":
            _ = len(v1)
        case "":
            v2 -= v2
        }
        v4 = v4
        for fi19 := 0; fi19 < 18; fi19++ {
            v20 := GenS2{F0: v4, F1: ("hello" + ("" + v4)), F2: 4294967295}
            v6 -= -288
            _ = fi19
            _ = v20
        }
    }
    v5 = int(v18)
    v18 = v18
    var v21 uint = 0
    v18 += (v18 / ((897) | 1))
    switch 2147483647 {
    case 127:
        v22 := GenS1{F0: 1}
        _ = v22
    case -2147483648:
        for fi23 := 0; fi23 < 16; fi23++ {
            switch fi23 {
            case 607:
                if false {
                    for fi24 := 0; fi24 < 20; fi24++ {
                        v5 += (fi24 - (fi23 / (((v6 | fi24)) | 1)))
                        var v25 int64 = v2
                        var v26 int32 = 2147483647
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    fmt.Println(fi23)
                } else {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        fmt.Println(v3, v2)
                        v28 := GenS2{F0: (v4 + "\n\t"), F1: "a", F2: 18446744073709551615}
                        _ = fi27
                        _ = v28
                    }
                }
            case 47:
                v3 = (-(v16))
            }
            for fi29 := 0; fi29 < 13; fi29++ {
                switch v4 {
                case "":
                    switch "foo bar" {
                    case "a":
                        if false {
                            fmt.Println(v4, v18)
                            var v30 int16 = 42
                            v2 = (v2 - (v2 % ((v2) | 1)))
                            if (1e10 >= (1e10 / 1.5)) {
                                v21 += genFn1()
                                var v31 bool = ((v21 | v21) <= v21)
                                _ = v31
                            } else {
                                v6 = (v5 - int(v2))
                                if ((((2147483647 < 124) || true) || false) || false) {
                                    var v32 float64 = ((1.5 * ((1.5 * 0.0) / (1.5 * ((1e10 + 1e10) / 1.5)))) - 3.14159)
                                    if (858 > 42) {
                                        fmt.Println(v6, fi23)
                                        fmt.Println(v21, v16, v2)
                                        v16 = -128
                                        v16 -= v3
                                        _ = len(v1)
                                        v33 := GenS3{F0: (^(v30)), F1: 602}
                                        if (!((v18 + 0) >= (v18 | (v18 >> ((42) & 63))))) {
                                            v34 := GenS2{F0: "\n\t", F1: v4, F2: 2}
                                            v35 := GenS4{F0: v30}
                                            _ = v34
                                            _ = v35
                                        } else {
                                            v4 = v4
                                            var v36 int32 = 2147483647
                                            v30 = -798
                                            v5 = (v6 / (((fi23 * v6)) | 1))
                                            var v37 string = (v4 + v4)
                                            var v38 string = (((v37 + (v37 + (v37 + v4))) + v4) + "foo bar")
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        }
                                        fmt.Println(v4, v21, v18)
                                        var v39 uint16 = uint16(v21)
                                        v40 := GenS2{F0: v4, F1: "", F2: 4294967295}
                                        var v41 float32 = 1.5
                                        switch -992 {
                                        case -9223372036854775808:
                                            v42 := GenS4{F0: (v30 + 42)}
                                            _ = v42
                                        case 1:
                                            v43 := GenS3{F0: (v30 - v30), F1: (v18 % (((v18 & 1)) | 1))}
                                            _ = v43
                                        default:
                                            var v44 float32 = 3.14159
                                            _ = v44
                                        }
                                        v45 := GenS5{F0: int64(v30), F1: (fi23 - fi23), F2: -1}
                                        v46 := v1[(v5 - v6)]
                                        _ = v33
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                        _ = v45
                                        _ = v46
                                    } else {
                                        if false {
                                            fmt.Println(v5)
                                        }
                                    }
                                    switch 32767 {
                                    case 4294967295:
                                        if true {
                                            switch "" {
                                            case "\n\t":
                                                v18 += v18
                                            }
                                            if false {
                                                var v47 uint16 = uint16(v21)
                                                v48 := GenS6{F0: (v2 * (v2 | v2))}
                                                v30 += int16(v47)
                                                v32 = 3.14159
                                                _ = v47
                                                _ = v48
                                            }
                                            var v49 uint = uint(v18)
                                            switch "foo bar" {
                                            case "hello":
                                                fmt.Println(v32, v32, v21)
                                            }
                                            _ = v49
                                        } else {
                                            if false {
                                                if ((v2 == (v2 - v2)) || true) {
                                                    fmt.Println(v3, v30)
                                                    _ = len(v1)
                                                    v50 := GenS7{F0: v2, F1: v18}
                                                    _ = v50
                                                }
                                                var v51 uint8 = 72
                                                _ = v51
                                            }
                                            v30 += v30
                                        }
                                    case -587:
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                    }
                                    _ = v32
                                } else {
                                    switch "hello" {
                                    case "a":
                                        var v52 float64 = float64(((-2.25 / 1.0) + 3.14159))
                                        var v53 int8 = v16
                                        _ = v52
                                        _ = v53
                                    case "\n\t":
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                    case "hello":
                                        fmt.Println(v16)
                                    default:
                                        var v54 int32 = 0
                                        _ = v54
                                    }
                                }
                                var v55 uint16 = 405
                                _ = v55
                            }
                            _ = v30
                        }
                    }
                }
                switch fi29 {
                case 0:
                    var v56 float32 = (0.0 - (3.14159 * 0.1))
                    _ = v56
                case -2147483648:
                    if ((float64((0.0 + (((0.1 / 3.14159) + (0.1 - 1e10)) / 1.0))) != float64(1e10)) || false) {
                        switch "hello" {
                        case "hello":
                            if len(v17) > 0 {
                                v17[0] = v5
                            }
                        }
                        var v57 int = (fi29 | (v5 % (((fi23 | fi23)) | 1)))
                        v1[v6] = 1
                        var v58 int64 = 9223372036854775807
                        v17 = append(v17, (fi29 * fi23))
                        _ = v57
                        _ = v58
                    }
                case -1:
                    if false {
                        if false {
                            switch fi29 {
                            case 1:
                                v6 = v5
                            default:
                                switch v4 {
                                case "a":
                                    fmt.Println(v21)
                                case "hello":
                                    switch (-(fi29)) {
                                    case -9223372036854775808:
                                        switch v5 {
                                        case -2147483648:
                                            v59 := GenS3{F0: -32768, F1: (v18 & v18)}
                                            _ = v59
                                        case -752:
                                            var v60 uint16 = 998
                                            _ = v60
                                        case -9223372036854775808:
                                            v61 := GenS8{F0: genFn1(), F1: 32767, F2: v3}
                                            _ = v61
                                        }
                                    case 4294967295:
                                        v18 -= (v18 * v18)
                                    }
                                case "":
                                    v62 := GenS6{F0: (v2 + v2)}
                                    _ = v62
                                }
                            }
                            if (!true) {
                                var v63 int = (v5 ^ (^(fi23)))
                                _ = v63
                            }
                            v64 := GenS9{F0: 0, F1: v2}
                            fmt.Println(v16)
                            if (int16(fi29) > -32768) {
                                var v65 uint16 = 255
                                _ = v65
                            }
                            v16 = v16
                            _ = v64
                        }
                        fmt.Println(v3)
                    }
                default:
                    v5 = (fi29 % (((v5 ^ (fi23 - (v6 >> ((2) & 63))))) | 1))
                }
                _ = fi29
            }
            if true {
                v66 := GenS10{F0: false, F1: 250, F2: -33}
                v67 := GenS11{F0: v16, F1: 2}
                _ = v66
                _ = v67
            } else {
                var v68 uint8 = 255
                _ = v68
            }
            _ = fi23
        }
    }
    fmt.Println(v4)
    switch (v6 % ((-247) | 1)) {
    case -2147483648:
        switch v4 {
        case "":
            var v69 int8 = v16
            _ = v69
        case "a":
            fmt.Println(v18, v5, v18)
        default:
            switch v6 {
            case 885:
                v5 = (v6 << ((2) & 63))
            default:
                if false {
                    for fi70 := 0; fi70 < 14; fi70++ {
                        v6 += -215
                        v2 = 42
                        _ = fi70
                    }
                } else {
                    fmt.Println(v3, v3)
                    var v71 uint = 2
                    _ = v71
                }
            }
        }
    }
    switch (v4 + v4) {
    case "\n\t":
        var v72 bool = false
        _ = v72
    default:
        var v73 uint16 = 2
        _ = v73
    }
    for fi74 := 0; fi74 < 12; fi74++ {
        fmt.Println(v3, v16)
        _ = fi74
    }
    for fi75 := 0; fi75 < 13; fi75++ {
        switch (fi75 >> (((v21 | genFn1())) & 63)) {
        case 4294967295:
            v21 = (v21 << (((v21 * (v21 ^ (v21 - v21)))) & 63))
        case 42:
            var v76 float64 = (3.14159 - (1e10 + 1e10))
            _ = v76
        case -2147483648:
            if (!true) {
                v21 = (^(v21))
            }
        }
        if false {
            var v77 int8 = v16
            if (true && false) {
                fmt.Println(v77)
                var v78 uint8 = 128
                _ = v78
            } else {
                if false {
                    v5 = 9223372036854775807
                }
                var v79 string = v4
                v3 = 100
                _ = v79
            }
            _ = v77
        } else {
            var v80 float32 = -2.25
            _ = len(v17)
            if false {
                if (!false) {
                    if (!(-541 >= 1)) {
                        v81 := v1[v6]
                        _ = v81
                    } else {
                        if false {
                            var v82 uint64 = 18446744073709551615
                            for fi83 := 0; fi83 < 17; fi83++ {
                                v18 = uint32(v82)
                                if (1 > 0) {
                                    switch v4 {
                                    case "foo bar":
                                        var v84 int16 = -1
                                        _ = v84
                                    case "hello":
                                        if (!((v16 ^ (v16 % (((v3 / ((v16) | 1))) | 1))) < -128)) {
                                            var v85 bool = false
                                            switch (fi75 / ((fi75) | 1)) {
                                            case 32768:
                                                var v86 uint64 = uint64(v18)
                                                _ = v86
                                            default:
                                                fmt.Println(v2, v80)
                                            }
                                            _ = v85
                                        }
                                    case "":
                                        v80 += 1e10
                                        var v87 string = v4
                                        v2 = (-(v2))
                                        _ = v87
                                    default:
                                        v17 = append(v17, v6)
                                    }
                                    v6 -= v5
                                }
                                fmt.Println(v80, v3, v5)
                                _ = fi83
                            }
                            _ = v82
                        } else {
                            var v88 int = (^(fi75))
                            delete(v1, (v5 % ((42) | 1)))
                            switch (-(fi75)) {
                            case 127:
                                if false {
                                    var v89 int16 = 933
                                    v21 = (v21 ^ (v21 | v21))
                                    if len(v17) > 0 {
                                        v17[0] = (v88 / (((v6 << ((574) & 63))) | 1))
                                    }
                                    _ = v89
                                }
                            case 117:
                                if (v3 == (v16 << (((v21 | 0)) & 63))) {
                                    var v90 string = ((v4 + v4) + ("" + v4))
                                    v91 := GenS7{F0: (-(v2)), F1: v18}
                                    v17 = append(v17, v88)
                                    _ = v90
                                    _ = v91
                                }
                            }
                            _ = v88
                        }
                        for fi92 := 0; fi92 < 7; fi92++ {
                            fmt.Println(v4, v4, v2)
                            _ = fi92
                        }
                    }
                }
                if false {
                    for fi93 := 0; fi93 < 18; fi93++ {
                        if (false && (!true)) {
                            var v94 bool = false
                            fmt.Println(fi75)
                            if v94 {
                                var v95 uint16 = 0
                                v21 -= v21
                                switch "hello" {
                                case "hello":
                                    var v96 int16 = -113
                                    _ = v96
                                case "\n\t":
                                    switch "hello" {
                                    case "a":
                                        v97 := GenS5{F0: -7, F1: v5, F2: 42}
                                        _ = v97
                                    default:
                                        switch v5 {
                                        case 988:
                                            v98 := GenS2{F0: ((v4 + "foo bar") + v4), F1: "a", F2: 4294967295}
                                            _ = v98
                                        case 127:
                                            if len(v17) > 0 {
                                                v17[0] = (-(v6))
                                            }
                                        case 4294967295:
                                            fmt.Println(fi93)
                                            if v94 {
                                                v99 := GenS12{F0: 63, F1: ((v95 ^ v95) < (^(v95)))}
                                                if ((v21 <= v21) || ((^(v6)) == (fi75 | fi75))) {
                                                    if (v94 || (1000 == 1000)) {
                                                        switch v6 {
                                                        case -695:
                                                            v16 = v3
                                                        case 32767:
                                                            var v100 uint16 = v95
                                                            _ = v100
                                                        case -9223372036854775808:
                                                            _ = len(v17)
                                                        }
                                                        v80 += ((-2.25 * v80) + v80)
                                                        v101 := GenS3{F0: 42, F1: v18}
                                                        _ = v101
                                                    }
                                                    var v102 int16 = 0
                                                    _ = v102
                                                } else {
                                                    v16 = -30
                                                    v103 := GenS7{F0: v2, F1: 2}
                                                    _ = v103
                                                }
                                                v95 = v95
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v21 -= 1
                                                case "a":
                                                    if ((v2 + (v2 ^ int64(v95))) >= int64(fi75)) {
                                                        switch v4 {
                                                        case "foo bar":
                                                            v104 := GenS13{F0: 0, F1: v94}
                                                            _ = v104
                                                        case "hello":
                                                            if (3.14159 != (0.0 + -2.25)) {
                                                                var v105 int8 = 1
                                                                _ = v105
                                                            }
                                                        case "":
                                                            v1[9223372036854775807] = v18
                                                        default:
                                                            switch "a" {
                                                            case "\n\t":
                                                                v94 = (47 < 18446744073709551615)
                                                            case "a":
                                                                if (!false) {
                                                                    var v106 int16 = int16(v95)
                                                                    v21 = uint(v18)
                                                                    if len(v17) >= 2 {
                                                                        v17 = v17[1:len(v17)]
                                                                    }
                                                                    _ = v106
                                                                } else {
                                                                    if v94 {
                                                                        switch int(v95) {
                                                                        case 167:
                                                                            v107 := v1[v6]
                                                                            _ = v107
                                                                        case -576:
                                                                            _ = len(v1)
                                                                        default:
                                                                            var v108 int16 = 982
                                                                            _ = v108
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v4)
                                                                    }
                                                                    delete(v1, v5)
                                                                    v109 := GenS1{F0: 2}
                                                                    _ = v109
                                                                }
                                                            case "hello":
                                                                if true {
                                                                    var v110 uint8 = 127
                                                                    _ = v110
                                                                } else {
                                                                    var v111 int16 = 630
                                                                    var v112 bool = false
                                                                    _ = v111
                                                                    _ = v112
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v16 = (v16 << ((v21) & 63))
                                                    switch ((v4 + v4) + "a") {
                                                    case "\n\t":
                                                        v6 += v6
                                                    case "hello":
                                                        switch "" {
                                                        case "":
                                                            v94 = v94
                                                        case "a":
                                                            if v94 {
                                                                var v113 uint8 = 47
                                                                v1[(^(fi93))] = (v18 + (^(v18)))
                                                                v114 := GenS14{F0: v113}
                                                                _ = v113
                                                                _ = v114
                                                            }
                                                        case "hello":
                                                            if v94 {
                                                                var v115 uint64 = 18446744073709551615
                                                                v116 := GenS1{F0: 910}
                                                                var v117 string = "a"
                                                                _ = v115
                                                                _ = v116
                                                                _ = v117
                                                            }
                                                        }
                                                    }
                                                }
                                                fmt.Println(v4, v16, v5)
                                                var v118 uint8 = 128
                                                if (!false) {
                                                    if false {
                                                        var v119 uint8 = (^(v118))
                                                        v119 = (^(v118))
                                                        _ = v119
                                                    } else {
                                                        if v94 {
                                                            v2 = genFn0(9223372036854775807)
                                                            v80 = v80
                                                            v6 = 32768
                                                        } else {
                                                            v3 = v3
                                                            v6 = fi93
                                                        }
                                                        fmt.Println(v95, v118, v95)
                                                        v2 = (v2 << (((v21 << (((^(v21))) & 63))) & 63))
                                                        fmt.Println(v16, v95, v94)
                                                    }
                                                    v120 := GenS7{F0: genFn0((v2 + v2)), F1: (v18 << ((2) & 63))}
                                                    v2 = v2
                                                    _ = v120
                                                }
                                                _ = v118
                                                _ = v99
                                            }
                                        }
                                    }
                                default:
                                    v121 := GenS5{F0: -1, F1: (-(fi93)), F2: v16}
                                    _ = v121
                                }
                                _ = v95
                            } else {
                                _ = len(v17)
                                fmt.Println(v3)
                                var v122 int = fi75
                                if ((v2 * (v2 + genFn0((v2 ^ v2)))) == v2) {
                                    var v123 int8 = (v3 + (v3 | (v3 % ((0) | 1))))
                                    switch (((v4 + v4) + "\n\t") + "hello") {
                                    case "":
                                        v21 -= 253
                                    case "foo bar":
                                        v124 := GenS1{F0: 2}
                                        _ = v124
                                    }
                                    var v125 uint64 = 99
                                    _ = v123
                                    _ = v125
                                }
                                _ = v122
                            }
                            var v126 int64 = (-(v2))
                            if (2 != 4294967295) {
                                v126 = (v126 & (v126 - v126))
                                if (1000 > 194) {
                                    switch v4 {
                                    case "hello":
                                        var v127 float32 = v80
                                        _ = v127
                                    case "":
                                        v94 = ("a" <= (v4 + (v4 + "hello")))
                                    case "foo bar":
                                        _ = len(v17)
                                    default:
                                        if v94 {
                                            v94 = v94
                                            v128 := GenS15{F0: ((true || (!(42 >= 773))) && v94), F1: 65535}
                                            _ = v128
                                        }
                                    }
                                    var v129 uint64 = 42
                                    _ = v129
                                } else {
                                    var v130 bool = true
                                    _ = v130
                                }
                            } else {
                                v3 = v3
                                v5 -= int(v16)
                                v131 := GenS16{F0: 255}
                                _ = v131
                            }
                            v132 := GenS10{F0: ((1 >= -2147483648) && false), F1: 128, F2: -1}
                            v133 := GenS1{F0: 4294967295}
                            _ = v126
                            _ = v132
                            _ = v133
                            _ = v94
                        } else {
                            _ = len(v1)
                            if (v18 < v18) {
                                v80 += v80
                                v5 = (v6 | (fi75 % (((v5 & (v5 * 136))) | 1)))
                            } else {
                                var v134 int8 = v16
                                fmt.Println(fi93, v2)
                                var v135 int = 827
                                _ = v134
                                _ = v135
                            }
                        }
                        var v136 int = v5
                        if (v80 == (3.14159 + v80)) {
                            var v137 int32 = 0
                            v138 := GenS17{F0: -396, F1: (v137 % ((v137) | 1)), F2: (^(v21))}
                            _ = v137
                            _ = v138
                        }
                        _ = fi93
                        _ = v136
                    }
                } else {
                    v139 := v1[2147483647]
                    v140 := GenS4{F0: -32768}
                    var v141 int64 = v2
                    _ = v139
                    _ = v140
                    _ = v141
                }
                if (true || (true || true)) {
                    var v142 float32 = v80
                    v18 -= (v18 / (((v18 | v18)) | 1))
                    _ = v142
                }
            }
            _ = v80
        }
        if ("\n\t" == v4) {
            for fi143 := 0; fi143 < 19; fi143++ {
                switch 685 {
                case -238:
                    v144 := GenS18{F0: 65535, F1: "hello"}
                    _ = v144
                case -9223372036854775808:
                    fmt.Println(v4)
                default:
                    var v145 float32 = (0.0 / 3.14159)
                    _ = v145
                }
                _ = fi143
            }
        }
        _ = fi75
    }
    for fi146 := 0; fi146 < 11; fi146++ {
        v147 := GenS4{F0: 32767}
        v18 += 65535
        _ = fi146
        _ = v147
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
