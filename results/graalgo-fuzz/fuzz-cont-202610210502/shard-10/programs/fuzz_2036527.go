package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int64; F2 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint; F2 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 int64 }
type GenS8 struct { F0 int8; F1 int; F2 int32 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 int16; F1 uint64; F2 uint }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int32; F1 uint32 }
type GenS14 struct { F0 uint64; F1 uint16 }
type GenS15 struct { F0 int64; F1 int64 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 uint; F1 uint32 }
type GenS18 struct { F0 string; F1 int; F2 uint8 }
type GenS19 struct { F0 int32; F1 int32 }
type GenS20 struct { F0 uint32; F1 string; F2 uint16 }
func genFn0(p0 uint16, p1 int8, p2 int) uint32 {
    p1 = p1
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int64 = 960
        v9 := GenS1{F0: -304}
        var v10 uint32 = 1
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 339
}

func main() {
    var v1 int16 = 1000
    v2 := []string{"", ("" + "hello"), "foo bar", "\n\t"}
    var v3 int64 = 1
    var v4 string = "\n\t"
    v5 := make(map[string]uint32)
    v6 := []int{-11, 9223372036854775807, 672}
    _ = len(v5)
    if true {
        fmt.Println(v1, v3, v1)
        if false {
            if (((0.1 * 3.14159) + (-2.25 / 3.14159)) != ((float64(-2.25) / 1.5) + 1e10)) {
                v1 += v1
                if (!false) {
                    fmt.Println(v3)
                }
                v1 = (v1 * (v1 % (((-(v1))) | 1)))
            } else {
                for fi11 := 0; fi11 < 7; fi11++ {
                    fmt.Println(v4, v3)
                    for fi12 := 0; fi12 < 11; fi12++ {
                        v13 := GenS2{F0: "hello"}
                        var v14 int16 = v1
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = fi11
                }
                v15 := GenS2{F0: v4}
                if (1 >= 2) {
                    v1 -= 42
                    switch 92 {
                    case 42:
                        v3 = 0
                    case -9223372036854775808:
                        v4 = v4
                    default:
                        var v16 int32 = -3
                        _ = v16
                    }
                } else {
                    for fi17 := 0; fi17 < 10; fi17++ {
                        if true {
                            if (!false) {
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    var v19 int64 = (-(v3))
                                    _ = len(v5)
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = fi18
                                    _ = v19
                                }
                                var v20 float64 = 1e10
                                if ((!true) || true) {
                                    delete(v5, (v4 + ("foo bar" + ((v4 + "foo bar") + ""))))
                                    switch ("hello" + "\n\t") {
                                    case "a":
                                        switch (((v4 + "a") + "a") + v4) {
                                        case "foo bar":
                                            v21 := v5[v4]
                                            _ = v21
                                        case "a":
                                            if (182 < 2) {
                                                v22 := GenS2{F0: v4}
                                                v1 += (-(v1))
                                                var v23 uint32 = 4294967295
                                                _ = v22
                                                _ = v23
                                            } else {
                                                v3 = 807
                                            }
                                        case "hello":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    }
                                    v24 := GenS1{F0: v1}
                                    _ = v24
                                } else {
                                    var v25 string = "hello"
                                    _ = v25
                                }
                                _ = v20
                            }
                            v1 = 774
                            fmt.Println(fi17)
                        }
                        _ = fi17
                    }
                    switch 1 {
                    case 4294967295:
                        for fi26 := 0; fi26 < 4; fi26++ {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v27 uint32 = 2
                            v27 -= genFn0(1, -1, fi26)
                            _ = fi26
                            _ = v27
                        }
                    case -9223372036854775808:
                        v3 = (-(v3))
                    case 2147483647:
                        v1 = (v1 / ((v1) | 1))
                    default:
                        switch 32767 {
                        case 42:
                            switch ("foo bar" + (v4 + v4)) {
                            case "foo bar":
                                for fi28 := 0; fi28 < 3; fi28++ {
                                    if true {
                                        switch v4 {
                                        case "hello":
                                            var v29 int8 = 42
                                            _ = v29
                                        case "foo bar":
                                            switch "\n\t" {
                                            case "a":
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            case "\n\t":
                                                v3 = (v3 & (v3 * int64(fi28)))
                                            default:
                                                for fi30 := 0; fi30 < 4; fi30++ {
                                                    v1 = (v1 - v1)
                                                    if true {
                                                        var v31 int = fi28
                                                        _ = v31
                                                    } else {
                                                        v32 := GenS2{F0: v4}
                                                        var v33 uint16 = 50
                                                        if (!((18446744073709551615 >= 586) && (false && (v4 != ("" + v4))))) {
                                                            switch (-(fi28)) {
                                                            case 127:
                                                                fmt.Println(v1, v4)
                                                            case 2147483647:
                                                                var v34 bool = true
                                                                _ = v34
                                                            default:
                                                                _ = len(v5)
                                                                v35 := GenS1{F0: v1}
                                                                _ = v35
                                                            }
                                                        }
                                                        v36 := GenS3{F0: 733, F1: (v3 - v3), F2: (fi30 + (^(fi28)))}
                                                        switch (fi30 * fi30) {
                                                        case 2147483647:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        case -1:
                                                            v3 = 608
                                                        case 127:
                                                            if true {
                                                                var v37 int = (fi28 ^ 2147483647)
                                                                var v38 int = (v37 ^ v37)
                                                                switch ((v4 + "foo bar") + v4) {
                                                                case "\n\t":
                                                                    v37 += -601
                                                                case "":
                                                                    var v39 int8 = -9
                                                                    _ = v39
                                                                case "foo bar":
                                                                    if false {
                                                                        if ((false && false) && (4294967295 >= 237)) {
                                                                            v40 := GenS3{F0: uint(v33), F1: v3, F2: v38}
                                                                            v41 := v5["\n\t"]
                                                                            switch ("hello" + "hello") {
                                                                            case "hello":
                                                                                switch v4 {
                                                                                case "foo bar":
                                                                                    v42 := GenS1{F0: (v1 + (^(v1)))}
                                                                                    _ = v42
                                                                                default:
                                                                                    v43 := GenS3{F0: 398, F1: v3, F2: int(v3)}
                                                                                    _ = v43
                                                                                }
                                                                            case "":
                                                                                if ((-1 != 117) && true) {
                                                                                    v37 -= 46
                                                                                    fmt.Println(v38, v33)
                                                                                }
                                                                            default:
                                                                                fmt.Println(v38, v38)
                                                                            }
                                                                            _ = v40
                                                                            _ = v41
                                                                        }
                                                                    } else {
                                                                        var v44 uint32 = 2147483647
                                                                        _ = v44
                                                                    }
                                                                default:
                                                                    if ((0.1 < (1.5 + 0.1)) || ((0 <= 0) || (9223372036854775807 != v3))) {
                                                                        if len(v6) > 0 {
                                                                            v6[0] = int(v1)
                                                                        }
                                                                        v3 = -717
                                                                        fmt.Println(fi28, v38)
                                                                    } else {
                                                                        var v45 float32 = float32(float64((0.1 * -2.25)))
                                                                        _ = v45
                                                                    }
                                                                    if (!false) {
                                                                        var v46 int8 = 115
                                                                        v38 = -1
                                                                        v47 := GenS4{F0: v4}
                                                                        _ = v46
                                                                        _ = v47
                                                                    }
                                                                }
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                        default:
                                                            var v48 float32 = 1e10
                                                            _ = v48
                                                        }
                                                        v33 += 42
                                                        _ = v32
                                                        _ = v33
                                                        _ = v36
                                                    }
                                                    fmt.Println(fi28, v4, fi28)
                                                    if (0 <= 26) {
                                                        v49 := GenS5{F0: 1, F1: 4294967295, F2: v4}
                                                        _ = v49
                                                    } else {
                                                        var v50 uint32 = 67
                                                        _ = v50
                                                    }
                                                    v51 := GenS6{F0: 127}
                                                    _ = fi30
                                                    _ = v51
                                                }
                                            }
                                        }
                                        switch fi28 {
                                        case 1:
                                            var v52 float64 = 0.0
                                            _ = v52
                                        case 42:
                                            for fi53 := 0; fi53 < 6; fi53++ {
                                                v54 := GenS4{F0: v4}
                                                switch v4 {
                                                case "\n\t":
                                                    if false {
                                                        var v55 uint = 0
                                                        v4 = "\n\t"
                                                        _ = v55
                                                    } else {
                                                        switch fi28 {
                                                        case 32767:
                                                            v56 := GenS7{F0: -32768, F1: (-(v3))}
                                                            v1 = v1
                                                            v3 = v3
                                                            v3 = (v3 / ((v3) | 1))
                                                            switch (v4 + "a") {
                                                            case "foo bar":
                                                                fmt.Println(v1)
                                                            case "hello":
                                                                v4 = "hello"
                                                            case "\n\t":
                                                                v4 = v4
                                                            }
                                                            v57 := GenS8{F0: 0, F1: 0, F2: 42}
                                                            _ = v56
                                                            _ = v57
                                                        case -882:
                                                            switch (fi53 - (fi28 - (fi28 >> ((18446744073709551615) & 63)))) {
                                                            case -1:
                                                                v58 := GenS9{F0: 0, F1: 42}
                                                                _ = v58
                                                            default:
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                        }
                                                        var v59 uint64 = 2
                                                        _ = v59
                                                    }
                                                    v60 := GenS6{F0: -18}
                                                    _ = v60
                                                }
                                                _ = fi53
                                                _ = v54
                                            }
                                        case -877:
                                            var v61 uint8 = 42
                                            _ = v61
                                        default:
                                            switch 925 {
                                            case 32767:
                                                v3 = v3
                                            }
                                        }
                                        fmt.Println(v1)
                                    }
                                    v62 := GenS8{F0: -37, F1: fi28, F2: 0}
                                    for fi63 := 0; fi63 < 7; fi63++ {
                                        switch fi63 {
                                        case -2147483648:
                                            var v64 int32 = -901
                                            _ = v64
                                        }
                                        v65 := GenS2{F0: v4}
                                        _ = fi63
                                        _ = v65
                                    }
                                    _ = fi28
                                    _ = v62
                                }
                            case "":
                                if ((3.14159 < 0.1) && (0 == 2)) {
                                    v1 += v1
                                    for fi66 := 0; fi66 < 1; fi66++ {
                                        v1 = v1
                                        v67 := GenS5{F0: 0, F1: 856, F2: "foo bar"}
                                        _ = fi66
                                        _ = v67
                                    }
                                    v68 := GenS10{F0: v1, F1: 800, F2: 65535}
                                    _ = v68
                                }
                            }
                        }
                    }
                }
                _ = v15
            }
            v5[(("foo bar" + v4) + v4)] = 0
            if ((-2.25 - 0.1) == ((1.5 / 1.0) + 0.0)) {
                v3 = (v3 >> ((1) & 63))
                for fi69 := 0; fi69 < 10; fi69++ {
                    switch fi69 {
                    case 9223372036854775807:
                        if (!true) {
                            v70 := GenS3{F0: 2, F1: (v3 << ((18446744073709551615) & 63)), F2: 519}
                            if (!true) {
                                var v71 bool = (false || (false && true))
                                _ = v71
                            } else {
                                switch (fi69 + fi69) {
                                case 1:
                                    v4 = (v4 + ((("" + (v4 + v4)) + "") + ((v4 + v4) + v4)))
                                default:
                                    for fi72 := 0; fi72 < 14; fi72++ {
                                        if (false && true) {
                                            switch (("hello" + v4) + v4) {
                                            case "":
                                                v73 := GenS5{F0: 127, F1: 42, F2: ""}
                                                _ = v73
                                            default:
                                                v3 -= v3
                                            }
                                            if true {
                                                v4 = v4
                                                v3 -= 1
                                                var v74 int8 = -128
                                                if true {
                                                    v74 -= v74
                                                    if true {
                                                        switch v4 {
                                                        case "foo bar":
                                                            var v75 int16 = v1
                                                            _ = v75
                                                        default:
                                                            v74 += v74
                                                        }
                                                        var v76 int16 = (v1 >> ((65535) & 63))
                                                        fmt.Println(v3, v74)
                                                        _ = v76
                                                    } else {
                                                        var v77 uint32 = 2147483647
                                                        var v78 float32 = float32((3.14159 * 1e10))
                                                        v79 := GenS11{F0: v77}
                                                        _ = v77
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                }
                                                _ = v74
                                            } else {
                                                v4 = ("" + "")
                                                var v80 uint8 = 0
                                                _ = v80
                                            }
                                            if true {
                                                fmt.Println(v1, fi69)
                                                if (0.1 == float32((1e10 / (1.5 + 0.1)))) {
                                                    var v81 float32 = (0.1 / (1e10 + float32((float64(0.0) + (0.0 / (3.14159 * (((((-2.25 * ((-2.25 / 1.0) / float64((0.1 / 1.0)))) - 0.1) / (float64((1.5 * -2.25)) + 0.1)) - 0.0) * 0.1)))))))
                                                    _ = v81
                                                }
                                                if true {
                                                    v82 := GenS2{F0: v4}
                                                    _ = v82
                                                } else {
                                                    v3 = (v3 << ((0) & 63))
                                                    fmt.Println(v4)
                                                }
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v83 uint8 = 1
                                                _ = v83
                                            }
                                            var v84 int8 = 0
                                            _ = v84
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = fi72
                                    }
                                }
                                var v85 int = (fi69 ^ fi69)
                                fmt.Println(v85)
                                _ = v85
                            }
                            var v86 uint32 = 1
                            _ = v70
                            _ = v86
                        } else {
                            var v87 uint = 18446744073709551615
                            if true {
                                if true {
                                    var v88 uint = v87
                                    v3 += v3
                                    _ = v88
                                } else {
                                    if (42 >= -831) {
                                        if len(v6) > 0 {
                                            v6[0] = (fi69 | fi69)
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v89 string = v4
                                        _ = v89
                                    } else {
                                        v5[(v4 + "a")] = 426
                                    }
                                    v3 = (v3 * (v3 + (v3 / (((v3 + (-(v3)))) | 1))))
                                }
                                v90 := GenS4{F0: ("foo bar" + v4)}
                                if (false || false) {
                                    v1 -= (v1 | 42)
                                    if false {
                                        fmt.Println(v4, v3)
                                        v3 = v3
                                    }
                                    if (!true) {
                                        var v91 int8 = 119
                                        var v92 uint = (v87 ^ v87)
                                        for fi93 := 0; fi93 < 14; fi93++ {
                                            fmt.Println(v3, v87)
                                            v92 += v92
                                            _ = fi93
                                        }
                                        _ = v91
                                        _ = v92
                                    }
                                } else {
                                    switch "hello" {
                                    case "\n\t":
                                        var v94 int16 = int16(fi69)
                                        _ = v94
                                    default:
                                        var v95 uint8 = 70
                                        _ = v95
                                    }
                                    var v96 uint8 = 1
                                    _ = v96
                                }
                                _ = v90
                            } else {
                                switch fi69 {
                                case 9223372036854775807:
                                    v97 := GenS12{F0: 0}
                                    _ = v97
                                }
                                if false {
                                    for fi98 := 0; fi98 < 3; fi98++ {
                                        if (!false) {
                                            v4 = v4
                                        }
                                        v99 := GenS9{F0: 1, F1: -112}
                                        _ = fi98
                                        _ = v99
                                    }
                                } else {
                                    v100 := GenS13{F0: 0, F1: 65535}
                                    _ = v100
                                }
                            }
                            _ = v87
                        }
                    }
                    var v101 int16 = (v1 << ((358) & 63))
                    for fi102 := 0; fi102 < 13; fi102++ {
                        switch "" {
                        case "":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case "hello":
                            switch (-(fi102)) {
                            case 4294967295:
                                v4 = v4
                            default:
                                v1 = -1
                            }
                        }
                        if (((^(fi102)) >= (fi69 / ((1) | 1))) || true) {
                            var v103 bool = (879 >= 639)
                            _ = v103
                        }
                        _ = fi102
                    }
                    _ = fi69
                    _ = v101
                }
            } else {
                for fi104 := 0; fi104 < 11; fi104++ {
                    v3 = -1
                    v105 := GenS10{F0: (v1 - (v1 << ((450) & 63))), F1: 95, F2: 65535}
                    _ = fi104
                    _ = v105
                }
            }
        }
        v106 := GenS7{F0: 32767, F1: (^(v3))}
        _ = v106
    } else {
        v1 -= v1
        var v107 uint = 65535
        switch 9223372036854775807 {
        case -9223372036854775808:
            if false {
                var v108 float64 = 1e10
                fmt.Println(v1)
                if (false || false) {
                    _ = len(v5)
                    var v109 string = v4
                    _ = v109
                } else {
                    v1 = (v1 << ((4294967295) & 63))
                    for fi110 := 0; fi110 < 16; fi110++ {
                        v107 += 887
                        v1 = v1
                        _ = fi110
                    }
                    if true {
                        fmt.Println(v4)
                        if (1 > 17) {
                            delete(v5, (v4 + "hello"))
                        }
                        fmt.Println(v3)
                    } else {
                        _ = len(v5)
                    }
                }
                _ = v108
            } else {
                v107 = v107
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                switch 42 {
                case 128:
                    var v111 float32 = 0.0
                    _ = v111
                case 0:
                    for fi112 := 0; fi112 < 3; fi112++ {
                        fmt.Println(v1)
                        _ = fi112
                    }
                default:
                    if (!false) {
                        v1 = v1
                        switch v4 {
                        case "a":
                            v4 = (((v4 + (("foo bar" + v4) + (v4 + "foo bar"))) + "") + v4)
                        case "foo bar":
                            v5["foo bar"] = genFn0(42, -1, 127)
                        case "\n\t":
                            for fi113 := 0; fi113 < 15; fi113++ {
                                v1 += (-(v1))
                                fmt.Println(v4, v3)
                                for fi114 := 0; fi114 < 7; fi114++ {
                                    switch fi114 {
                                    case 815:
                                        v1 = 1
                                    case 2147483647:
                                        var v115 float32 = ((((3.14159 + (-2.25 * (3.14159 - (3.14159 - (1e10 + 3.14159))))) / -2.25) * (3.14159 / ((1e10 * 3.14159) - 3.14159))) * 1.5)
                                        _ = v115
                                    default:
                                        v107 += (v107 % ((v107) | 1))
                                    }
                                    v116 := GenS14{F0: 670, F1: 0}
                                    v107 = v107
                                    _ = fi114
                                    _ = v116
                                }
                                _ = fi113
                            }
                        }
                        _ = len(v6)
                    }
                }
            }
        }
        _ = v107
    }
    v117 := GenS7{F0: 42, F1: v3}
    _ = len(v5)
    if (!(!(1 != 2147483647))) {
        v4 = v4
    } else {
        var v118 int = -9223372036854775808
        _ = v118
    }
    if len(v6) > 0 {
        v6[0] = -271
    }
    v1 = v1
    switch -9223372036854775808 {
    case 262:
        v119 := GenS15{F0: 42, F1: (v3 ^ v3)}
        _ = v119
    case 128:
        v120 := GenS1{F0: (v1 + 1)}
        _ = v120
    case 32768:
        for fi121 := 0; fi121 < 15; fi121++ {
            for fi122 := 0; fi122 < 16; fi122++ {
                v123 := v5[((("foo bar" + v4) + ((v4 + v4) + "foo bar")) + "foo bar")]
                var v124 int32 = -2147483648
                v5[((v4 + "a") + v4)] = 2147483648
                _ = fi122
                _ = v123
                _ = v124
            }
            v125 := GenS1{F0: (v1 & -32768)}
            v126 := GenS16{F0: 1}
            _ = fi121
            _ = v125
            _ = v126
        }
    }
    delete(v5, ((("" + v4) + (((v4 + v4) + v4) + (v4 + "\n\t"))) + v4))
    for fi127 := 0; fi127 < 13; fi127++ {
        _ = len(v6)
        _ = fi127
    }
    if true {
        v128 := GenS17{F0: 18446744073709551615, F1: 1}
        _ = v128
    }
    v3 -= (-(v3))
    v3 -= (v3 * (-(v3)))
    var v129 int = 696
    if true {
        if false {
            _ = len(v2)
            var v130 int16 = v1
            switch v4 {
            case "hello":
                v131 := GenS1{F0: v130}
                _ = v131
            }
            _ = v130
        }
        var v132 uint = 0
        _ = v132
    } else {
        v133 := v5[((v4 + ("" + v4)) + (v4 + (v4 + "hello")))]
        for fi134 := 0; fi134 < 6; fi134++ {
            if true {
                v135 := GenS4{F0: ("" + v4)}
                if false {
                    v1 = (v1 + 0)
                } else {
                    for fi136 := 0; fi136 < 5; fi136++ {
                        v137 := v5[v4]
                        _ = len(v5)
                        delete(v5, (v4 + (v4 + ("foo bar" + v4))))
                        _ = fi136
                        _ = v137
                    }
                    for fi138 := 0; fi138 < 8; fi138++ {
                        if true {
                            v139 := GenS6{F0: 42}
                            v4 = ((v4 + v4) + ("hello" + v4))
                            switch v129 {
                            case -9223372036854775808:
                                fmt.Println(v3)
                            case 4294967295:
                                var v140 int64 = (v3 * (v3 + v3))
                                _ = v140
                            }
                            _ = v139
                        } else {
                            if (1 <= 18446744073709551615) {
                                if len(v6) > 0 {
                                    v6[0] = fi138
                                }
                                var v141 float64 = ((((-2.25 / 1e10) * (((3.14159 + (0.1 * (0.1 - (-2.25 - 1e10)))) * 0.1) - float64((-2.25 * 0.0)))) - float64(0.0)) / float64(0.1))
                                _ = v141
                            } else {
                                v129 = fi134
                                v142 := GenS18{F0: v4, F1: (fi138 ^ (-(fi138))), F2: 127}
                                _ = v142
                            }
                            if (1.5 > (((3.14159 / (1.5 / -2.25)) / (-2.25 * 1e10)) - 0.1)) {
                                var v143 uint64 = 4294967295
                                _ = v143
                            }
                        }
                        v144 := GenS4{F0: v4}
                        var v145 uint = 310
                        _ = fi138
                        _ = v144
                        _ = v145
                    }
                }
                _ = len(v5)
                _ = v135
            } else {
                if true {
                    v1 += (v1 % (((v1 + 1000)) | 1))
                    _ = len(v5)
                }
                v146 := GenS19{F0: -2147483648, F1: 42}
                _ = v146
            }
            v147 := GenS1{F0: (v1 / (((v1 | 32767)) | 1))}
            v129 = 0
            _ = fi134
            _ = v147
        }
        for fi148 := 0; fi148 < 8; fi148++ {
            if false {
                v4 = v4
                if (v3 == 0) {
                    switch v4 {
                    case "a":
                        var v149 string = "hello"
                        _ = v149
                    case "hello":
                        v150 := GenS19{F0: 1, F1: 42}
                        _ = v150
                    }
                    v129 += (v129 >> ((289) & 63))
                    var v151 uint32 = 0
                    _ = v151
                }
            } else {
                switch ((("a" + "") + ("a" + v4)) + v4) {
                case "\n\t":
                    _ = len(v5)
                case "a":
                    switch v129 {
                    case 4294967295:
                        fmt.Println(v1, v133)
                    case 127:
                        for fi152 := 0; fi152 < 2; fi152++ {
                            fmt.Println(v3, v1)
                            var v153 int16 = v1
                            v4 = ""
                            _ = len(v6)
                            v3 -= (v3 % ((v3) | 1))
                            _ = len(v6)
                            v154 := GenS16{F0: 18446744073709551615}
                            _ = fi152
                            _ = v153
                            _ = v154
                        }
                    case 0:
                        if true {
                            v155 := GenS6{F0: -122}
                            _ = v155
                        }
                    }
                case "foo bar":
                    v129 = (fi148 << ((412) & 63))
                default:
                    v133 = v133
                }
                v156 := GenS19{F0: 2147483647, F1: 1}
                v157 := GenS20{F0: (v133 & 233), F1: v4, F2: 1}
                _ = v156
                _ = v157
            }
            var v158 uint32 = (v133 + (v133 + v133))
            _ = fi148
            _ = v158
        }
        _ = v133
    }
    if (true && (true && (v3 >= (v3 & 161)))) {
        for fi159 := 0; fi159 < 15; fi159++ {
            v160 := GenS14{F0: 2, F1: 0}
            v129 -= (^(fi159))
            if true {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if true {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    v4 = "hello"
                }
            }
            _ = fi159
            _ = v160
        }
        fmt.Println(v3, v1, v1)
        if true {
            v129 -= v129
        }
    } else {
        fmt.Println(v3, v1)
        var v161 bool = ((false && true) && ((((1e10 + 0.0) == (float32((3.14159 / (((1.5 - -2.25) - 1.5) - (float64(-2.25) + 3.14159)))) / float32((1e10 - 0.1)))) && (2147483647 <= 2147483648)) && (v4 > "")))
        _ = v161
    }
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v117
    _ = v129
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
