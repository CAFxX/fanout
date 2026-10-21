package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint64; F1 int64; F2 int16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 string; F1 int32; F2 int16 }
type GenS12 struct { F0 string; F1 int8; F2 string }
type GenS13 struct { F0 uint32; F1 int }
type GenS14 struct { F0 int32; F1 uint8; F2 string }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int64; F1 uint64 }
type GenS17 struct { F0 int; F1 bool }
type GenS18 struct { F0 int32; F1 int }
type GenS19 struct { F0 int64; F1 int16 }
type GenS20 struct { F0 int32; F1 uint8 }
type GenS21 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 42
    var v3 bool = false
    v4 := make(map[string]string)
    var v5 float32 = ((-2.25 / (1.5 / (3.14159 - float32(-2.25)))) / 1.0)
    v6 := []int{1, 51, 4294967295}
    if (196 != 1) {
        if (1e10 != float64(1.5)) {
            if v3 {
                switch (("\n\t" + "hello") + "") {
                case "\n\t":
                    var v7 int8 = -128
                    _ = v7
                case "hello":
                    v8 := GenS1{F0: 1, F1: 42, F2: 1}
                    _ = v8
                }
                if len(v6) > 0 {
                    v6[0] = 2147483647
                }
            } else {
                switch -1 {
                case -2147483648:
                    if v3 {
                        v5 += v5
                    }
                case 294:
                    for fi9 := 0; fi9 < 5; fi9++ {
                        v10 := GenS2{F0: 2, F1: (v2 | v2)}
                        fmt.Println(v3, fi9)
                        var v11 int8 = 42
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                case 1:
                    v12 := GenS3{F0: 32767, F1: v2, F2: 0}
                    _ = v12
                default:
                    fmt.Println(v5, v3)
                }
            }
            v5 = v5
        }
    } else {
        fmt.Println(v2)
        fmt.Println(v3, v2, v3)
    }
    var v13 float32 = v5
    v14 := GenS2{F0: 4294967295, F1: (v2 | v2)}
    var v15 int = -617
    v2 = (v2 / (((v2 << ((42) & 63))) | 1))
    if (!(false && v3)) {
        var v16 bool = v3
        if v16 {
            if ((-156 < (v15 | (v15 >> ((65535) & 63)))) || (1 >= v15)) {
                v2 += 0
                v17 := GenS4{F0: int8(v2)}
                _ = v17
            }
        }
        _ = v16
    } else {
        for fi18 := 0; fi18 < 14; fi18++ {
            for fi19 := 0; fi19 < 18; fi19++ {
                v3 = v3
                _ = fi19
            }
            switch "hello" {
            case "":
                for fi20 := 0; fi20 < 15; fi20++ {
                    if (42 == v2) {
                        v21 := GenS2{F0: 4294967295, F1: 2}
                        fmt.Println(v3, v13)
                        if v3 {
                            v3 = (0 >= 32767)
                        } else {
                            v22 := GenS5{F0: "foo bar", F1: fi18}
                            if false {
                                v23 := GenS1{F0: 42, F1: 128, F2: 65535}
                                _ = len(v4)
                                v24 := GenS3{F0: 1, F1: v2, F2: uint32(v2)}
                                _ = v23
                                _ = v24
                            } else {
                                var v25 bool = (55 == 101)
                                if v25 {
                                    if (!(!true)) {
                                        var v26 int64 = 9223372036854775807
                                        _ = v26
                                    }
                                    switch "" {
                                    case "foo bar":
                                        if v3 {
                                            v27 := GenS6{F0: 1, F1: -1, F2: -742}
                                            var v28 int8 = -66
                                            fmt.Println(fi18)
                                            _ = v27
                                            _ = v28
                                        } else {
                                            v3 = false
                                            fmt.Println(v25)
                                            _ = len(v6)
                                        }
                                    case "\n\t":
                                        if false {
                                            var v29 bool = (100 != 2)
                                            _ = v29
                                        }
                                    case "":
                                        v3 = false
                                    default:
                                        v4["hello"] = "hello"
                                    }
                                    v5 -= -2.25
                                }
                                if v25 {
                                    if (0.1 == float64((float32(-2.25) / (3.14159 * (v5 * (v5 - (float32(1e10) - (v13 + (-2.25 * float32(1.5)))))))))) {
                                        v30 := GenS7{F0: v2}
                                        switch "foo bar" {
                                        case "foo bar":
                                            switch fi18 {
                                            case -111:
                                                v15 -= v15
                                                switch "hello" {
                                                case "a":
                                                    var v31 int = (v15 * fi20)
                                                    _ = v31
                                                case "":
                                                    v32 := GenS8{F0: 1, F1: -1}
                                                    _ = v32
                                                default:
                                                    v15 -= (v15 + v15)
                                                }
                                            case 127:
                                                if v3 {
                                                    if ("a" <= "") {
                                                        var v33 int16 = -32768
                                                        var v34 int32 = 2147483647
                                                        if (164 > 127) {
                                                            v34 -= 1
                                                            var v35 float64 = 1e10
                                                            fmt.Println(v25)
                                                            _ = v35
                                                        } else {
                                                            var v36 uint16 = 0
                                                            switch (v15 & -2147483648) {
                                                            case 696:
                                                                fmt.Println(v33, v15, fi18)
                                                            }
                                                            _ = v36
                                                        }
                                                        v37 := GenS8{F0: (v2 | 583), F1: 42}
                                                        _ = v33
                                                        _ = v34
                                                        _ = v37
                                                    }
                                                } else {
                                                    if ("hello" <= "a") {
                                                        v38 := GenS9{F0: (!v3)}
                                                        if (!true) {
                                                            if (!false) {
                                                                v39 := GenS1{F0: 0, F1: 235, F2: 2}
                                                                var v40 int16 = -1
                                                                _ = v39
                                                                _ = v40
                                                            }
                                                            if (-128 != 42) {
                                                                var v41 int = v15
                                                                var v42 float32 = v13
                                                                _ = v41
                                                                _ = v42
                                                            } else {
                                                                if v3 {
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        if v25 {
                                                                            if ((v2 >> ((42) & 63)) >= (v2 & 255)) {
                                                                                v43 := GenS10{F0: 1}
                                                                                switch (("a" + ((("foo bar" + "foo bar") + ("" + ("hello" + "\n\t"))) + "hello")) + ("" + ("a" + "foo bar"))) {
                                                                                case "foo bar":
                                                                                    v5 += float32((0.1 - -2.25))
                                                                                case "hello":
                                                                                    var v44 uint32 = 877
                                                                                    _ = v44
                                                                                default:
                                                                                    if v3 {
                                                                                        switch "foo bar" {
                                                                                        case "foo bar":
                                                                                            v15 -= fi18
                                                                                        case "\n\t":
                                                                                            var v45 int16 = 1
                                                                                            _ = v45
                                                                                        }
                                                                                        switch ("" + "hello") {
                                                                                        case "foo bar":
                                                                                            var v46 uint8 = 0
                                                                                            _ = v46
                                                                                        default:
                                                                                            v47 := v1[391]
                                                                                            _ = v47
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if (204 == 0) {
                                                                                    if v3 {
                                                                                        var v48 int64 = -1
                                                                                        fmt.Println(v13, v25)
                                                                                        _ = v48
                                                                                    }
                                                                                    if v25 {
                                                                                        var v49 int32 = 570
                                                                                        var v50 uint16 = 255
                                                                                        var v51 uint8 = 128
                                                                                        _ = v49
                                                                                        _ = v50
                                                                                        _ = v51
                                                                                    } else {
                                                                                        if v25 {
                                                                                            var v52 int8 = 5
                                                                                            v4["hello"] = "foo bar"
                                                                                            _ = v52
                                                                                        }
                                                                                        v5 = float32(((1e10 + 1e10) + 3.14159))
                                                                                        if v3 {
                                                                                            fmt.Println(v2)
                                                                                            var v53 string = "foo bar"
                                                                                            _ = v53
                                                                                        }
                                                                                        switch ("" + "a") {
                                                                                        case "":
                                                                                            fmt.Println(fi20, v2, v15)
                                                                                            v3 = v25
                                                                                        default:
                                                                                            v2 = (v2 * v2)
                                                                                        }
                                                                                    }
                                                                                    if len(v6) > 0 {
                                                                                        v6[0] = (fi20 % ((4294967295) | 1))
                                                                                    }
                                                                                }
                                                                                _ = v43
                                                                            } else {
                                                                                fmt.Println(v15, v5)
                                                                                var v54 float64 = (1.5 + 0.0)
                                                                                v15 += fi18
                                                                                _ = v54
                                                                            }
                                                                            var v55 int16 = 1000
                                                                            v13 -= 0.0
                                                                            v3 = ((127 != 100) || true)
                                                                            switch (("hello" + "a") + ("" + "")) {
                                                                            case "a":
                                                                                v5 = -2.25
                                                                            case "\n\t":
                                                                                v56 := v1[0]
                                                                                _ = v56
                                                                            }
                                                                            v13 += (float32(1e10) - float32(1e10))
                                                                            _ = v55
                                                                        }
                                                                    default:
                                                                        v5 -= (((1e10 - ((v13 * 1e10) + float32((0.1 + 1e10)))) - float32(((0.0 / 1.0) / 1e10))) + (v5 / v5))
                                                                    }
                                                                    v25 = (v2 < 46)
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v57 := GenS11{F0: "a", F1: -1, F2: 1}
                                                                        _ = v57
                                                                    default:
                                                                        if (1.5 <= ((float32(0.1) * v13) * 0.0)) {
                                                                            if v25 {
                                                                                v58 := GenS12{F0: ((("a" + (("\n\t" + ((("hello" + "") + "foo bar") + "hello")) + "")) + "") + "foo bar"), F1: -32, F2: "foo bar"}
                                                                                _ = v58
                                                                            }
                                                                        } else {
                                                                            v3 = v25
                                                                            fmt.Println(v2, v5, v3)
                                                                        }
                                                                    }
                                                                    fmt.Println(v5, v13)
                                                                    v59 := GenS13{F0: 677, F1: v15}
                                                                    _ = v59
                                                                }
                                                                if len(v6) > 0 {
                                                                    v6[0] = -9223372036854775808
                                                                }
                                                                var v60 int64 = 42
                                                                _ = v60
                                                            }
                                                        }
                                                        fmt.Println(fi20)
                                                        _ = v38
                                                    } else {
                                                        v3 = (!(!v3))
                                                        switch (v15 % (((v15 & (fi18 - (fi18 - fi18)))) | 1)) {
                                                        case 4294967295:
                                                            v61 := v1[0]
                                                            var v62 uint = 2
                                                            if len(v6) > 0 {
                                                                v6[0] = 1
                                                            }
                                                            _ = v61
                                                            _ = v62
                                                        case 118:
                                                            v6 = append(v6, (v15 | (fi20 % ((-716) | 1))))
                                                        case 0:
                                                            v4[((("a" + "") + "\n\t") + "foo bar")] = "\n\t"
                                                        default:
                                                            v63 := GenS14{F0: 465, F1: 128, F2: "foo bar"}
                                                            if false {
                                                                var v64 int64 = 42
                                                                var v65 int = (-(fi20))
                                                                v66 := GenS15{F0: -32768}
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                            } else {
                                                                v25 = true
                                                                v67 := GenS15{F0: 0}
                                                                _ = v67
                                                            }
                                                            _ = v63
                                                        }
                                                        var v68 uint32 = 1
                                                        switch ("a" + "hello") {
                                                        case "a":
                                                            v69 := GenS1{F0: 194, F1: 2, F2: 255}
                                                            _ = v69
                                                        case "\n\t":
                                                            var v70 uint8 = 255
                                                            _ = v70
                                                        case "hello":
                                                            v68 += v68
                                                        default:
                                                            v71 := GenS13{F0: (v68 >> ((319) & 63)), F1: (v15 + -224)}
                                                            _ = v71
                                                        }
                                                        _ = v68
                                                    }
                                                    v15 -= 2147483647
                                                    if ((42 >= 388) && true) {
                                                        if (42 >= 18446744073709551615) {
                                                            var v72 uint16 = v2
                                                            fmt.Println(v72)
                                                            v72 -= (^(v72))
                                                            _ = v72
                                                        }
                                                        var v73 int32 = 42
                                                        var v74 string = ("hello" + "foo bar")
                                                        _ = v73
                                                        _ = v74
                                                    } else {
                                                        switch "" {
                                                        case "\n\t":
                                                            var v75 uint8 = 128
                                                            _ = v75
                                                        case "hello":
                                                            v76 := GenS5{F0: "hello", F1: 9223372036854775807}
                                                            _ = v76
                                                        default:
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        }
                                                    }
                                                }
                                            default:
                                                v13 = v13
                                            }
                                        }
                                        _ = len(v6)
                                        if (v25 && (!(fi20 != (fi20 | (fi20 | v15))))) {
                                            var v77 uint32 = 65535
                                            _ = v77
                                        }
                                        if (false && (!v3)) {
                                            var v78 int = v15
                                            _ = v78
                                        } else {
                                            _ = len(v6)
                                            v13 = v5
                                            var v79 uint8 = 0
                                            _ = v79
                                        }
                                        _ = v30
                                    } else {
                                        v13 = 1e10
                                        v25 = ((fi18 < 9223372036854775807) && (9223372036854775807 != 42))
                                        var v80 int = fi18
                                        _ = v80
                                    }
                                    switch (fi20 + v15) {
                                    case 452:
                                        v13 -= 3.14159
                                    case 127:
                                        if true {
                                            v25 = (0 == 42)
                                            var v81 uint = 42
                                            fmt.Println(fi18)
                                            _ = v81
                                        }
                                    case 90:
                                        delete(v4, "")
                                    default:
                                        delete(v1, 1)
                                    }
                                    v25 = ("foo bar" > "")
                                }
                                _ = v25
                            }
                            _ = v22
                        }
                        _ = v21
                    }
                    switch "" {
                    case "foo bar":
                        fmt.Println(v3, v15)
                    case "hello":
                        v5 = float32(-2.25)
                    default:
                        v2 -= (^(v2))
                    }
                    _ = fi20
                }
            case "a":
                var v82 uint8 = 127
                _ = v82
            case "hello":
                v6 = append(v6, (v15 ^ fi18))
            default:
                switch (fi18 | fi18) {
                case -804:
                    if true {
                        switch fi18 {
                        case 42:
                            for fi83 := 0; fi83 < 2; fi83++ {
                                var v84 int = 427
                                var v85 uint16 = v2
                                v86 := GenS5{F0: (("foo bar" + "hello") + ""), F1: -120}
                                _ = fi83
                                _ = v84
                                _ = v85
                                _ = v86
                            }
                        }
                    } else {
                        for fi87 := 0; fi87 < 10; fi87++ {
                            var v88 int16 = -970
                            v89 := GenS16{F0: 511, F1: 4294967295}
                            v90 := GenS5{F0: "hello", F1: (v15 & -362)}
                            if (!v3) {
                                var v91 int8 = -42
                                v88 += (v88 - 42)
                                var v92 int8 = v91
                                _ = v91
                                _ = v92
                            }
                            var v93 uint64 = 2
                            fmt.Println(fi87)
                            v6 = append(v6, (^(v15)))
                            v15 = 153
                            if false {
                                delete(v1, 4294967295)
                                v3 = false
                                v94 := GenS17{F0: v15, F1: (1e10 > 0.1)}
                                _ = v94
                            }
                            _ = len(v6)
                            v2 -= (^(v2))
                            _ = fi87
                            _ = v88
                            _ = v89
                            _ = v90
                            _ = v93
                        }
                    }
                default:
                    for fi95 := 0; fi95 < 6; fi95++ {
                        switch "foo bar" {
                        case "hello":
                            switch (("a" + ("foo bar" + "hello")) + "\n\t") {
                            case "a":
                                if v3 {
                                    var v96 int32 = -816
                                    var v97 int8 = -1
                                    v15 -= (fi95 & (^(fi95)))
                                    switch v15 {
                                    case 663:
                                        _ = len(v1)
                                    }
                                    v98 := GenS15{F0: -1}
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    if v3 {
                                        var v99 int16 = 928
                                        switch "a" {
                                        case "hello":
                                            fmt.Println(fi18, fi18)
                                        case "a":
                                            switch "foo bar" {
                                            case "hello":
                                                if false {
                                                    var v100 int16 = -1
                                                    var v101 uint8 = 1
                                                    _ = v100
                                                    _ = v101
                                                } else {
                                                    var v102 int32 = (v96 ^ v96)
                                                    v1[4294967295] = ((false && v3) && ((v99 ^ v99) > v99))
                                                    _ = v102
                                                }
                                            case "foo bar":
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                        default:
                                            v103 := GenS18{F0: -2147483648, F1: int(v96)}
                                            _ = v103
                                        }
                                        _ = v99
                                    } else {
                                        if v3 {
                                            v97 = 0
                                        } else {
                                            v104 := GenS15{F0: 42}
                                            v2 -= v2
                                            _ = v104
                                        }
                                        var v105 float32 = 0.0
                                        if (((((3.14159 + 1.5) * (1e10 + -2.25)) * 1e10) < -2.25) && (!false)) {
                                            switch fi95 {
                                            case 32767:
                                                var v106 uint64 = 18446744073709551615
                                                _ = v106
                                            case 0:
                                                v15 = int(v2)
                                            default:
                                                if (42 == v97) {
                                                    v107 := GenS19{F0: 9223372036854775807, F1: 1000}
                                                    v2 = v2
                                                    _ = v107
                                                }
                                            }
                                            var v108 uint16 = (v2 % (((v2 % ((704) | 1))) | 1))
                                            switch 32768 {
                                            case 4294967295:
                                                v13 += ((v13 / 1.0) * v5)
                                            case -701:
                                                var v109 int32 = v96
                                                _ = v109
                                            case 128:
                                                fmt.Println(v13, v96)
                                            default:
                                                _ = len(v6)
                                            }
                                            _ = v108
                                        }
                                        _ = v105
                                    }
                                    v3 = (v96 > (v96 ^ (v96 % ((-1) | 1))))
                                    _ = v96
                                    _ = v97
                                    _ = v98
                                } else {
                                    v3 = v3
                                }
                            default:
                                v110 := GenS5{F0: "hello", F1: (fi95 ^ -219)}
                                _ = v110
                            }
                        default:
                            delete(v4, "")
                            v15 += fi95
                        }
                        var v111 bool = ((255 != 2) && (1 == -25))
                        v2 -= (v2 | v2)
                        _ = fi95
                        _ = v111
                    }
                }
            }
            _ = fi18
        }
        if v3 {
            fmt.Println(v2, v2, v15)
            var v112 int8 = 42
            _ = v112
        } else {
            if v3 {
                v13 -= 1e10
            } else {
                v6 = append(v6, v15)
                if (10 <= -128) {
                    v3 = (42 >= (v2 % ((42) | 1)))
                    if true {
                        v113 := GenS17{F0: 4294967295, F1: v3}
                        if v3 {
                            var v114 int16 = 42
                            if (128 == 33) {
                                var v115 uint8 = 127
                                _ = v115
                            }
                            _ = v114
                        }
                        var v116 uint8 = 1
                        _ = v113
                        _ = v116
                    }
                } else {
                    for fi117 := 0; fi117 < 12; fi117++ {
                        v5 -= 0.0
                        _ = fi117
                    }
                    _ = len(v6)
                    v5 -= (((((float32(1e10) - ((((0.1 / 1.0) * -2.25) / (float32(0.0) + v13)) - v5)) - v13) * float32(((0.1 - 1e10) * 3.14159))) / (1e10 + v13)) - float32(((-2.25 - (float64(v5) * ((float64(3.14159) * 0.1) + 1.5))) + (-2.25 / 1.5))))
                }
                if true {
                    var v118 float64 = ((float64(-2.25) * float64(v5)) * ((((0.1 + float64(0.0)) + -2.25) / (3.14159 + 0.1)) + float64((v5 + v13))))
                    if true {
                        v119 := GenS20{F0: 260, F1: 70}
                        _ = v119
                    } else {
                        v15 -= (v15 >> ((0) & 63))
                    }
                    switch v15 {
                    case 149:
                        v120 := GenS21{F0: -1, F1: 283}
                        _ = v120
                    }
                    _ = v118
                }
            }
            for fi121 := 0; fi121 < 1; fi121++ {
                var v122 string = "a"
                _ = fi121
                _ = v122
            }
            switch "" {
            case "\n\t":
                fmt.Println(v13, v3)
            case "hello":
                v123 := v1[2147483648]
                _ = v123
            }
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = len(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
