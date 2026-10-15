package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 int32; F1 int8; F2 int64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint32; F1 string; F2 bool }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint16; F1 int64; F2 int8 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint32; F1 int }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int64; F1 uint; F2 int64 }
type GenS18 struct { F0 bool; F1 int64; F2 int32 }
func genFn0(p0 int16) uint16 {
    switch 1 {
    case -219:
        v4 := GenS1{F0: 780}
        _ = v4
    case 9223372036854775807:
        var v5 uint64 = 0
        _ = v5
    default:
        for fi6 := 0; fi6 < 14; fi6++ {
            switch -1 {
            case -983:
                for fi7 := 0; fi7 < 4; fi7++ {
                    p0 -= 1
                    _ = fi7
                }
            case 2147483647:
                switch fi6 {
                case -9223372036854775808:
                    var v8 bool = true
                    _ = v8
                case 42:
                    var v9 uint64 = 0
                    _ = v9
                default:
                    switch (("foo bar" + "foo bar") + ("a" + "")) {
                    case "a":
                        v10 := GenS1{F0: 42}
                        _ = v10
                    default:
                        switch (("hello" + ((("" + (((("foo bar" + "a") + ("foo bar" + "\n\t")) + "hello") + (("foo bar" + "") + "hello"))) + ("" + "foo bar")) + "hello")) + "hello") {
                        case "\n\t":
                            var v11 uint16 = 42
                            _ = v11
                        case "a":
                            var v12 int64 = 42
                            _ = v12
                        default:
                            for fi13 := 0; fi13 < 13; fi13++ {
                                p0 -= p0
                                _ = fi13
                            }
                        }
                    }
                }
            default:
                var v14 int64 = -9223372036854775808
                _ = v14
            }
            var v15 uint = 65535
            _ = fi6
            _ = v15
        }
    }
    fmt.Println(p0)
    for fi16 := 0; fi16 < 17; fi16++ {
        fmt.Println(p0)
        fmt.Println(fi16, fi16)
        _ = fi16
    }
    _ = p0
    return 65535
}
func genFn1() uint16 {
    v17 := GenS2{F0: 42, F1: -1}
    v18 := []int{1, 449}
    for fi19 := 0; fi19 < 10; fi19++ {
        if (!true) {
            var v20 uint64 = 42
            if ((316 <= 18446744073709551615) || false) {
                v21 := GenS2{F0: 42, F1: 0}
                _ = v21
            } else {
                if false {
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    v20 = (^(v20))
                } else {
                    if false {
                        v18 = append(v18, (fi19 | 128))
                        v20 = (^(v20))
                    } else {
                        var v22 float32 = (1e10 - ((0.0 + 0.1) - -2.25))
                        var v23 int64 = 1
                        v24 := GenS2{F0: -1, F1: (v23 % ((42) | 1))}
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    if ((!((false || (!("a" >= "a"))) && false)) || true) {
                        switch "" {
                        case "\n\t":
                            if true {
                                var v25 int16 = 42
                                _ = v25
                            } else {
                                var v26 bool = true
                                var v27 int = (fi19 - (fi19 + (-(fi19))))
                                _ = v26
                                _ = v27
                            }
                        case "hello":
                            v20 += v20
                        default:
                            v20 -= v20
                        }
                        v20 += v20
                        v20 = (v20 ^ 308)
                    } else {
                        var v28 float64 = 3.14159
                        v29 := GenS3{F0: 128}
                        if ((true && false) && false) {
                            for fi30 := 0; fi30 < 17; fi30++ {
                                switch -9223372036854775808 {
                                case -2147483648:
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                case 32768:
                                    var v31 float64 = 1e10
                                    _ = v31
                                }
                                if false {
                                    v20 -= v20
                                    if (2147483647 > 42) {
                                        v20 += (v20 | 18446744073709551615)
                                    } else {
                                        var v32 int8 = 42
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                        _ = v32
                                    }
                                } else {
                                    if (fi30 == fi19) {
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                        if ((-2147483648 >= 1) && false) {
                                            fmt.Println(fi30, fi19)
                                            v28 = float64((3.14159 / 0.1))
                                            if true {
                                                var v33 uint16 = 1
                                                _ = v33
                                            }
                                        } else {
                                            v20 = 42
                                            v28 = 1.5
                                        }
                                        var v34 uint16 = 2
                                        if ((v28 > float64(-2.25)) && true) {
                                            v35 := GenS3{F0: 17}
                                            if len(v18) > 0 {
                                                v18[0] = (fi30 ^ fi19)
                                            }
                                            _ = v35
                                        }
                                        v20 += (v20 % ((1) | 1))
                                        _ = v34
                                    } else {
                                        if true {
                                            var v36 uint = 0
                                            switch ("" + ((("" + "foo bar") + "a") + ("" + ("hello" + "")))) {
                                            case "foo bar":
                                                if false {
                                                    fmt.Println(v20, v36)
                                                    var v37 int8 = -128
                                                    var v38 uint16 = uint16(v20)
                                                    _ = v37
                                                    _ = v38
                                                } else {
                                                    fmt.Println(v20)
                                                    v36 -= v36
                                                    v28 = (((v28 + v28) * -2.25) - (1.5 * 0.1))
                                                    switch (fi30 + 782) {
                                                    case 1:
                                                        if ((211 > 42) || (42 <= 2)) {
                                                            v28 = (((v28 + v28) / ((float64(3.14159) * v28) * float64(float32(3.14159)))) + -2.25)
                                                            if ((!false) && ((v36 & 624) >= (v36 | v36))) {
                                                                var v39 float64 = float64(0.1)
                                                                if len(v18) >= 2 {
                                                                    v18 = v18[1:len(v18)]
                                                                }
                                                                var v40 uint64 = v20
                                                                if (1 != 0) {
                                                                    var v41 int = 130
                                                                    v28 += v28
                                                                    _ = v41
                                                                }
                                                                _ = v39
                                                                _ = v40
                                                            } else {
                                                                v28 -= v28
                                                                v28 = v28
                                                            }
                                                            var v42 uint = v36
                                                            v42 = (v42 * v36)
                                                            v28 += (-2.25 - v28)
                                                            var v43 int = fi30
                                                            var v44 int32 = -631
                                                            _ = v42
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                    }
                                                    v18 = append(v18, (fi19 ^ (fi30 | fi19)))
                                                    if ((!false) || true) {
                                                        v45 := GenS4{F0: fi30}
                                                        fmt.Println(fi19, fi19, v20)
                                                        fmt.Println(v28, v36)
                                                        _ = v45
                                                    } else {
                                                        fmt.Println(v28, v28)
                                                        v46 := GenS5{F0: 674}
                                                        _ = v46
                                                    }
                                                }
                                            case "":
                                                if true {
                                                    v47 := GenS6{F0: true, F1: 287}
                                                    _ = v47
                                                } else {
                                                    v48 := GenS1{F0: 1}
                                                    _ = v48
                                                }
                                            case "a":
                                                if true {
                                                    v18 = append(v18, 4294967295)
                                                } else {
                                                    var v49 bool = true
                                                    _ = v49
                                                }
                                            default:
                                                v50 := GenS6{F0: false, F1: 4294967295}
                                                _ = v50
                                            }
                                            _ = v36
                                        }
                                        switch fi19 {
                                        case -260:
                                            var v51 int8 = 100
                                            _ = v51
                                        case 42:
                                            v52 := GenS7{F0: "a", F1: 1}
                                            _ = v52
                                        default:
                                            if (628 > 0) {
                                                if (!(255 < 193)) {
                                                    var v53 int32 = 960
                                                    _ = v53
                                                }
                                                if true {
                                                    v28 = (v28 * (float64(((3.14159 / -2.25) / 3.14159)) - v28))
                                                    fmt.Println(v20)
                                                    if len(v18) >= 2 {
                                                        v18 = v18[1:len(v18)]
                                                    }
                                                    fmt.Println(fi30, v20, fi30)
                                                } else {
                                                    v54 := GenS3{F0: 7}
                                                    v28 = ((((((v28 * v28) * (v28 * ((v28 / float64(-2.25)) - float64((3.14159 / (0.1 * 0.1)))))) * v28) * v28) - float64(1.5)) / ((1e10 - (((v28 - ((float64(1e10) + v28) / v28)) + float64(-2.25)) / 1.0)) - ((v28 - ((v28 / v28) * (3.14159 + v28))) / (float64(1e10) + v28))))
                                                    if len(v18) >= 2 {
                                                        v18 = v18[1:len(v18)]
                                                    }
                                                    _ = v54
                                                }
                                                var v55 uint = 777
                                                _ = v55
                                            } else {
                                                if false {
                                                    var v56 uint32 = 1
                                                    _ = v56
                                                } else {
                                                    if false {
                                                        v28 = float64(3.14159)
                                                        var v57 uint16 = 65535
                                                        var v58 uint16 = 863
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                }
                                                v20 = 1
                                                if true {
                                                    if true {
                                                        v59 := GenS8{F0: -2147483648, F1: 100, F2: 427}
                                                        _ = v59
                                                    }
                                                    var v60 int8 = 47
                                                    var v61 int16 = -32768
                                                    _ = v60
                                                    _ = v61
                                                }
                                            }
                                        }
                                        var v62 int16 = 0
                                        _ = v62
                                    }
                                }
                                v63 := GenS9{F0: 1000}
                                _ = fi30
                                _ = v63
                            }
                            v20 -= v20
                        }
                        _ = v28
                        _ = v29
                    }
                    v20 -= (v20 + v20)
                }
            }
            _ = v20
        }
        _ = fi19
    }
    v64 := GenS10{F0: genFn0(-464)}
    _ = v17
    _ = v18
    _ = v64
    return 0
}

func main() {
    v1 := []int{2147483647, 1, -19, 42}
    var v2 int32 = -1
    var v3 int64 = 0
    v3 -= 42
    if (100 != 42) {
        for fi65 := 0; fi65 < 18; fi65++ {
            for fi66 := 0; fi66 < 3; fi66++ {
                switch ("hello" + "hello") {
                case "foo bar":
                    var v67 uint64 = 1
                    _ = v67
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    switch (fi65 / (((fi65 / ((42) | 1))) | 1)) {
                    case -984:
                        var v68 uint16 = 88
                        _ = v68
                    case 2147483647:
                        if (!(-2.25 < -2.25)) {
                            var v69 int16 = -392
                            fmt.Println(v69, v69)
                            v70 := GenS3{F0: 255}
                            _ = v69
                            _ = v70
                        }
                    default:
                        v71 := GenS10{F0: 65535}
                        _ = v71
                    }
                }
                _ = fi66
            }
            var v72 float64 = (3.14159 * 1e10)
            _ = fi65
            _ = v72
        }
        for fi73 := 0; fi73 < 1; fi73++ {
            if true {
                v2 = (v2 ^ 0)
                if (!(127 < 0)) {
                    v74 := GenS3{F0: 0}
                    var v75 int64 = v3
                    v76 := GenS4{F0: (fi73 + fi73)}
                    _ = v74
                    _ = v75
                    _ = v76
                } else {
                    for fi77 := 0; fi77 < 7; fi77++ {
                        if (-32768 != -238) {
                            if true {
                                fmt.Println(fi77)
                                v3 = v3
                                v78 := GenS9{F0: 32767}
                                _ = v78
                            } else {
                                var v79 int8 = 127
                                v2 = v2
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v79
                            }
                            if true {
                                var v80 uint32 = 2147483647
                                if false {
                                    v1 = append(v1, fi77)
                                } else {
                                    v2 = (v2 / (((v2 >> ((850) & 63))) | 1))
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v80
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = fi77
                        }
                        v2 = (^(v2))
                        if true {
                            v2 -= (v2 * v2)
                            if len(v1) > 0 {
                                v1[0] = fi77
                            }
                            var v81 uint16 = genFn0(1000)
                            _ = v81
                        } else {
                            v82 := GenS11{F0: 389, F1: ((("hello" + ((("foo bar" + "hello") + "a") + "")) + ("foo bar" + "")) + "\n\t"), F2: (-2.25 < 3.14159)}
                            _ = v82
                        }
                        v83 := GenS12{F0: 41}
                        v84 := GenS9{F0: 934}
                        _ = fi77
                        _ = v83
                        _ = v84
                    }
                    fmt.Println(fi73)
                    var v85 float32 = 1e10
                    _ = v85
                }
                v86 := GenS13{F0: 363, F1: (v3 >> ((255) & 63)), F2: 0}
                _ = v86
            } else {
                if len(v1) > 0 {
                    v1[0] = (fi73 | 1)
                }
            }
            v87 := GenS2{F0: v3, F1: (v3 * v3)}
            _ = fi73
            _ = v87
        }
    }
    var v88 uint8 = 1
    if ((65535 > 4294967295) && (3.14159 > (0.1 / 1e10))) {
        for fi89 := 0; fi89 < 1; fi89++ {
            v2 += v2
            var v90 bool = false
            _ = fi89
            _ = v90
        }
        v3 = v3
        switch ("\n\t" + ("" + "")) {
        case "hello":
            v88 -= 157
        default:
            v91 := GenS9{F0: -720}
            _ = v91
        }
    }
    v2 = v2
    for fi92 := 0; fi92 < 19; fi92++ {
        for fi93 := 0; fi93 < 6; fi93++ {
            fmt.Println(v2, v88, v3)
            var v94 uint64 = 18446744073709551615
            _ = fi93
            _ = v94
        }
        for fi95 := 0; fi95 < 3; fi95++ {
            var v96 int = (-(fi92))
            switch "a" {
            case "\n\t":
                fmt.Println(v3, v88, v3)
            case "hello":
                v1 = append(v1, -688)
            }
            switch "" {
            case "":
                fmt.Println(v3, v96)
            case "foo bar":
                v97 := GenS14{F0: false}
                _ = v97
            case "a":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (9223372036854775807 <= -9223372036854775808) {
                    fmt.Println(v88, fi95, fi95)
                    if (!false) {
                        if (!((1.5 + 3.14159) >= 0.0)) {
                            v2 -= v2
                            v1 = append(v1, (v96 >> ((42) & 63)))
                            v98 := GenS15{F0: 2, F1: (fi95 >> ((65535) & 63))}
                            _ = len(v1)
                            if (v2 > (v2 % (((v2 / ((v2) | 1))) | 1))) {
                                if true {
                                    var v99 string = (("hello" + "a") + ((("a" + ("hello" + "a")) + ((("foo bar" + "") + "hello") + "hello")) + "foo bar"))
                                    var v100 int16 = 1000
                                    v101 := GenS16{F0: 4294967295}
                                    _ = v100
                                    _ = v101
                                    _ = v99
                                } else {
                                    fmt.Println(fi95, v3, fi92)
                                    v3 = v3
                                }
                                switch fi92 {
                                case -861:
                                    if false {
                                        var v102 int16 = 42
                                        _ = v102
                                    } else {
                                        fmt.Println(v3, v3)
                                        v88 = (v88 >> ((247) & 63))
                                        var v103 int16 = 0
                                        v2 = v2
                                        _ = v103
                                    }
                                case 4294967295:
                                    switch (fi92 & 2147483647) {
                                    case 128:
                                        v88 += 27
                                    case 127:
                                        _ = len(v1)
                                    case 2147483647:
                                        v104 := GenS15{F0: 2, F1: (fi92 / ((-1) | 1))}
                                        _ = v104
                                    default:
                                        var v105 bool = true
                                        _ = v105
                                    }
                                case 0:
                                    if false {
                                        var v106 uint32 = 501
                                        v2 += (v2 * (v2 + v2))
                                        _ = v106
                                    }
                                default:
                                    v1 = append(v1, (fi95 * (fi95 | fi92)))
                                }
                                if true {
                                    var v107 int16 = 566
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v107
                                } else {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v2 = v2
                                    case "\n\t":
                                        if ((("hello" + "foo bar") <= "") && ((1 > 2) || false)) {
                                            var v108 uint64 = 42
                                            switch v96 {
                                            case 127:
                                                if false {
                                                    if false {
                                                        var v109 uint64 = 42
                                                        _ = v109
                                                    } else {
                                                        fmt.Println(v2, v96, v88)
                                                        if (-24 < -1) {
                                                            var v110 bool = ((0.1 >= float64(((3.14159 * 1e10) - -2.25))) || true)
                                                            _ = v110
                                                        }
                                                    }
                                                }
                                            case 56:
                                                v111 := GenS9{F0: -588}
                                                _ = v111
                                            default:
                                                v108 += 1
                                            }
                                            var v112 int8 = 42
                                            _ = v108
                                            _ = v112
                                        } else {
                                            fmt.Println(v3, v2)
                                            v113 := GenS17{F0: 1, F1: 271, F2: v3}
                                            _ = v113
                                        }
                                        switch (fi95 | (fi95 | (v96 % (((^(fi92))) | 1)))) {
                                        case -2147483648:
                                            switch "" {
                                            case "":
                                                v96 = 4294967295
                                            case "hello":
                                                v114 := GenS8{F0: v2, F1: -95, F2: 1}
                                                _ = v114
                                            }
                                        case 124:
                                            var v115 uint32 = 42
                                            _ = v115
                                        default:
                                            v2 -= -960
                                        }
                                    case "a":
                                        var v116 int32 = int32(v88)
                                        _ = v116
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v117 := GenS18{F0: false, F1: (v3 | v3), F2: (^(v2))}
                                    _ = v117
                                }
                                v118 := GenS17{F0: (v3 | v3), F1: 65535, F2: v3}
                                switch "\n\t" {
                                case "foo bar":
                                    switch (^(fi95)) {
                                    case 127:
                                        fmt.Println(v96)
                                    }
                                case "\n\t":
                                    switch ("" + ("hello" + "hello")) {
                                    case "":
                                        var v119 float64 = float64((1.5 / 0.1))
                                        _ = v119
                                    case "a":
                                        if false {
                                            switch (fi95 ^ v96) {
                                            case 1:
                                                switch "foo bar" {
                                                case "hello":
                                                    switch (fi92 & 128) {
                                                    case 0:
                                                        var v120 string = ""
                                                        _ = v120
                                                    default:
                                                        v2 = v2
                                                    }
                                                case "foo bar":
                                                    switch (fi92 | 32767) {
                                                    case 128:
                                                        v96 = v96
                                                    case -124:
                                                        v121 := GenS11{F0: 2147483648, F1: "hello", F2: (v88 < (v88 & 127))}
                                                        _ = v121
                                                    case 9223372036854775807:
                                                        fmt.Println(v2, v88)
                                                    default:
                                                        var v122 float64 = 0.1
                                                        _ = v122
                                                    }
                                                    var v123 string = ""
                                                    _ = v123
                                                default:
                                                    var v124 float32 = 0.0
                                                    _ = v124
                                                }
                                            case -2147483648:
                                                fmt.Println(v2)
                                            }
                                        }
                                    }
                                case "a":
                                    if len(v1) > 0 {
                                        v1[0] = fi92
                                    }
                                default:
                                    if false {
                                        if (!(true || (v2 <= (v2 ^ -1)))) {
                                            v1 = append(v1, int(v88))
                                            fmt.Println(fi92, v96)
                                            var v125 uint32 = 4294967295
                                            if (!(!false)) {
                                                v125 = v125
                                                var v126 uint32 = v125
                                                v125 = (v126 % ((112) | 1))
                                                _ = v126
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v127 uint = 742
                                            _ = v125
                                            _ = v127
                                        } else {
                                            var v128 float32 = 0.1
                                            var v129 int64 = 9223372036854775807
                                            fmt.Println(v3, fi95, v3)
                                            v130 := GenS1{F0: 1}
                                            if (false || true) {
                                                v129 = (v3 ^ (v129 & 9223372036854775807))
                                            }
                                            _ = v128
                                            _ = v129
                                            _ = v130
                                        }
                                    } else {
                                        switch (fi92 * int(v3)) {
                                        case -1:
                                            v1 = append(v1, (fi95 >> ((4294967295) & 63)))
                                            fmt.Println(v96)
                                        case 2147483647:
                                            if false {
                                                fmt.Println(fi92)
                                                fmt.Println(v3, v96)
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v3 += v3
                                                if len(v1) > 0 {
                                                    v1[0] = fi92
                                                }
                                            }
                                        default:
                                            v131 := GenS17{F0: v3, F1: 914, F2: (v3 * (v3 >> ((278) & 63)))}
                                            _ = v131
                                        }
                                    }
                                }
                                _ = v118
                            } else {
                                var v132 string = (("\n\t" + "hello") + "\n\t")
                                var v133 string = "foo bar"
                                _ = v132
                                _ = v133
                            }
                            _ = v98
                        } else {
                            v1 = append(v1, v96)
                        }
                        v3 = (v3 | (v3 & int64(v2)))
                        var v134 int = v96
                        _ = v134
                    }
                } else {
                    if (3.14159 > ((3.14159 + 3.14159) / -2.25)) {
                        v96 = v96
                    }
                    if ((float64(1.5) * ((1.5 - 1.5) + (0.1 * 0.0))) != float64(3.14159)) {
                        var v135 string = ("\n\t" + "")
                        _ = v135
                    } else {
                        switch fi92 {
                        case -2147483648:
                            _ = len(v1)
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v136 uint = 65535
                        _ = v136
                    }
                }
            }
            _ = fi95
            _ = v96
        }
        var v137 int8 = 78
        _ = fi92
        _ = v137
    }
    if (-549 <= 42) {
        var v138 int32 = v2
        _ = v138
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (-9223372036854775808 >= (-(v3))) {
        var v139 uint = 80
        var v140 uint = 2
        var v141 uint16 = genFn0(-1)
        _ = v139
        _ = v140
        _ = v141
    }
    v3 = v3
    if true {
        var v142 uint16 = 2
        v143 := GenS4{F0: -1}
        _ = v142
        _ = v143
    }
    v2 = v2
    v1 = append(v1, -9223372036854775808)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v88
}
