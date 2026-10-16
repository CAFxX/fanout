package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint32; F2 int }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint; F1 string; F2 uint64 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 int32; F2 bool }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 bool; F1 string }

func main() {
    var v1 int8 = 42
    var v2 int8 = (v1 << ((738) & 63))
    v3 := make(map[uint32]bool)
    var v4 int8 = (v1 / (((v1 >> ((0) & 63))) | 1))
    v5 := make(map[uint32]string)
    var v6 float32 = 3.14159
    v7 := []int{32767, 1}
    v4 = v2
    v5[751] = ("\n\t" + (("\n\t" + "\n\t") + ("a" + "")))
    var v8 string = ""
    if false {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint16 = 0
            if true {
                v11 := GenS1{F0: 420, F1: 815}
                if (-32768 <= 0) {
                    v12 := GenS2{F0: (v10 / ((1) | 1)), F1: 9223372036854775807, F2: 2}
                    for fi13 := 0; fi13 < 6; fi13++ {
                        v2 = (v2 % ((v1) | 1))
                        _ = fi13
                    }
                    _ = v12
                } else {
                    v4 = 0
                }
                v14 := v3[42]
                _ = v11
                _ = v14
            } else {
                var v15 float64 = float64(v6)
                switch (^(fi9)) {
                case 621:
                    if false {
                        var v16 uint32 = 2147483647
                        for fi17 := 0; fi17 < 15; fi17++ {
                            fmt.Println(v1)
                            v18 := GenS3{F0: false, F1: v16, F2: 9223372036854775807}
                            _ = fi17
                            _ = v18
                        }
                        _ = v16
                    } else {
                        v19 := GenS2{F0: v10, F1: 733, F2: 42}
                        v20 := GenS2{F0: v10, F1: -9223372036854775808, F2: 0}
                        if (2147483647 > 600) {
                            for fi21 := 0; fi21 < 8; fi21++ {
                                v1 -= -98
                                v15 = 0.1
                                switch fi9 {
                                case -2147483648:
                                    fmt.Println(v6)
                                }
                                _ = fi21
                            }
                            if (-1 > -1) {
                                v22 := GenS3{F0: (!true), F1: 82, F2: 806}
                                _ = v22
                            }
                            var v23 uint32 = 2147483648
                            _ = v23
                        }
                        _ = v19
                        _ = v20
                    }
                case -313:
                    var v24 int32 = -1
                    _ = v24
                }
                fmt.Println(v15, v6)
                _ = v15
            }
            for fi25 := 0; fi25 < 2; fi25++ {
                v26 := GenS1{F0: 65535, F1: 42}
                v10 -= v10
                if true {
                    v27 := GenS2{F0: (v10 - v10), F1: 42, F2: uint32(v10)}
                    if true {
                        v28 := GenS2{F0: v10, F1: 9223372036854775807, F2: 4294967295}
                        if false {
                            if (1000 == 669) {
                                if (false || (false || false)) {
                                    var v29 float32 = ((float32(0.1) * (v6 * v6)) / float32(3.14159))
                                    v30 := GenS4{F0: 577, F1: 4294967295}
                                    _ = v29
                                    _ = v30
                                }
                            } else {
                                if false {
                                    v10 = (v10 + v10)
                                    v31 := GenS1{F0: 586, F1: 0}
                                    _ = v31
                                } else {
                                    v1 = -39
                                }
                                var v32 int64 = 42
                                v7 = append(v7, fi25)
                                _ = v32
                            }
                        } else {
                            delete(v3, 2147483647)
                            var v33 int32 = 42
                            v4 = (^(v4))
                            _ = v33
                        }
                        _ = v28
                    }
                    _ = v27
                } else {
                    switch fi25 {
                    case -588:
                        switch (v8 + "\n\t") {
                        case "":
                            if len(v7) > 0 {
                                v7[0] = fi9
                            }
                        case "\n\t":
                            switch ((v8 + v8) + v8) {
                            case "\n\t":
                                var v34 int64 = 9223372036854775807
                                _ = v34
                            }
                        case "a":
                            _ = len(v7)
                        }
                    case 42:
                        var v35 float64 = 1e10
                        _ = v35
                    }
                    v6 = 0.1
                }
                fmt.Println(v1, v6)
                v8 = v8
                _ = fi25
                _ = v26
            }
            _ = fi9
            _ = v10
        }
        if false {
            if false {
                fmt.Println(v8)
                var v36 int32 = -462
                var v37 uint64 = 433
                _ = v36
                _ = v37
            }
        }
    }
    fmt.Println(v8, v6)
    if true {
        var v38 uint = 86
        _ = v38
    } else {
        if (!false) {
            v1 += v2
            v6 += (3.14159 - (v6 / v6))
        }
        switch v8 {
        case "foo bar":
            v5[1] = (v8 + ("" + v8))
        case "\n\t":
            if true {
                switch 183 {
                case 9223372036854775807:
                    var v39 uint64 = 1
                    _ = v39
                case -689:
                    if (!(2 <= 465)) {
                        var v40 uint8 = 1
                        _ = v40
                    }
                default:
                    fmt.Println(v8, v6, v6)
                }
            } else {
                if (0 > 961) {
                    var v41 int8 = (-(v1))
                    var v42 float32 = v6
                    switch (v8 + "hello") {
                    case "\n\t":
                        v8 = (("foo bar" + v8) + (v8 + "hello"))
                    case "a":
                        if (true && true) {
                            v42 += float32((1e10 - -2.25))
                            v43 := GenS5{F0: 255, F1: v8, F2: 18446744073709551615}
                            for fi44 := 0; fi44 < 8; fi44++ {
                                switch v8 {
                                case "foo bar":
                                    var v45 float64 = 0.1
                                    _ = v45
                                case "hello":
                                    fmt.Println(v42, v1, v41)
                                default:
                                    for fi46 := 0; fi46 < 15; fi46++ {
                                        var v47 int = (fi44 >> ((255) & 63))
                                        _ = fi46
                                        _ = v47
                                    }
                                }
                                _ = fi44
                            }
                            _ = v43
                        } else {
                            v42 = 1e10
                        }
                    case "":
                        if true {
                            for fi48 := 0; fi48 < 16; fi48++ {
                                v1 = -1
                                switch "hello" {
                                case "":
                                    if len(v7) > 0 {
                                        v7[0] = (fi48 % (((fi48 - (fi48 << ((0) & 63)))) | 1))
                                    }
                                case "hello":
                                    v41 = v2
                                default:
                                    _ = len(v7)
                                }
                                v49 := GenS6{F0: false, F1: 0}
                                _ = fi48
                                _ = v49
                            }
                        } else {
                            var v50 bool = (-1 >= 128)
                            var v51 int64 = 9223372036854775807
                            if (!v50) {
                                v8 = v8
                                var v52 uint32 = 4294967295
                                _ = v52
                            }
                            _ = v50
                            _ = v51
                        }
                    }
                    _ = v41
                    _ = v42
                } else {
                    if (!(734 > 819)) {
                        for fi53 := 0; fi53 < 15; fi53++ {
                            for fi54 := 0; fi54 < 1; fi54++ {
                                var v55 int = 763
                                if ((1e10 / float32(-2.25)) < float32(1.5)) {
                                    v5[110] = "foo bar"
                                }
                                switch fi54 {
                                case 0:
                                    v4 = (v2 / (((v1 ^ (v1 ^ 100))) | 1))
                                    var v56 bool = false
                                    var v57 int = 127
                                    v58 := GenS2{F0: 1, F1: 9223372036854775807, F2: 123}
                                    if (1 < -548) {
                                        var v59 uint32 = 4294967295
                                        _ = v59
                                    }
                                    var v60 uint8 = 0
                                    v57 = fi54
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                    _ = v60
                                case 4294967295:
                                    if false {
                                        v1 -= v1
                                        if (v1 <= v1) {
                                            v61 := GenS4{F0: 65535, F1: 266}
                                            if false {
                                                switch 4294967295 {
                                                case -2147483648:
                                                    var v62 int32 = -1
                                                    _ = v62
                                                case 1:
                                                    v5[4294967295] = ("hello" + "hello")
                                                    var v63 uint64 = 876
                                                    _ = v63
                                                default:
                                                    v55 += int(v2)
                                                }
                                                v6 += v6
                                                fmt.Println(fi54, v6, v2)
                                                v4 -= (v1 / ((100) | 1))
                                                v64 := GenS6{F0: (float64(float32(float64((3.14159 * float32(3.14159))))) != 1e10), F1: 440}
                                                v1 += (^(v4))
                                                v65 := GenS7{F0: 42}
                                                delete(v3, 1)
                                                var v66 int = (fi54 & v55)
                                                if len(v7) > 0 {
                                                    v7[0] = v66
                                                }
                                                var v67 int32 = -257
                                                var v68 float32 = 1.5
                                                if false {
                                                    var v69 bool = false
                                                    _ = v69
                                                }
                                                fmt.Println(v4)
                                                _ = len(v7)
                                                fmt.Println(v2)
                                                fmt.Println(v2)
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                            } else {
                                                v1 += 127
                                            }
                                            if (18446744073709551615 >= 173) {
                                                var v70 int8 = (v2 + v1)
                                                var v71 int = fi53
                                                _ = v70
                                                _ = v71
                                            }
                                            _ = v61
                                        } else {
                                            v72 := v5[400]
                                            if (-1 < 1) {
                                                fmt.Println(v4, fi54, fi54)
                                                if false {
                                                    var v73 int32 = 1
                                                    if true {
                                                        if true {
                                                            var v74 int16 = -32768
                                                            _ = v74
                                                        } else {
                                                            v6 += v6
                                                        }
                                                        v1 = int8(fi53)
                                                    }
                                                    v2 -= 42
                                                    fmt.Println(v8, fi54)
                                                    _ = v73
                                                }
                                                switch 32767 {
                                                case 783:
                                                    v8 = "hello"
                                                default:
                                                    switch "foo bar" {
                                                    case "":
                                                        _ = len(v7)
                                                    default:
                                                        fmt.Println(v2, fi53, fi54)
                                                    }
                                                }
                                            }
                                            v2 -= v2
                                            _ = v72
                                        }
                                        v55 -= fi54
                                    }
                                case 42:
                                    switch v55 {
                                    case -2147483648:
                                        if false {
                                            switch "a" {
                                            case "\n\t":
                                                if (true && true) {
                                                    switch fi53 {
                                                    case -103:
                                                        v1 += (v2 ^ 1)
                                                    case -782:
                                                        v7 = append(v7, fi53)
                                                    case 32767:
                                                        v75 := GenS5{F0: 65535, F1: v8, F2: 4294967295}
                                                        _ = v75
                                                    }
                                                    var v76 int64 = 9223372036854775807
                                                    _ = v76
                                                } else {
                                                    switch (fi54 * fi54) {
                                                    case 42:
                                                        if false {
                                                            if (true || true) {
                                                                var v77 uint8 = 0
                                                                switch (fi54 & fi53) {
                                                                case -2147483648:
                                                                    v78 := GenS8{F0: 967, F1: 0, F2: true}
                                                                    _ = v78
                                                                case 0:
                                                                    if len(v7) > 0 {
                                                                        v7[0] = fi54
                                                                    }
                                                                case -1:
                                                                    switch (v8 + (v8 + v8)) {
                                                                    case "":
                                                                        if len(v7) >= 2 {
                                                                            v7 = v7[1:len(v7)]
                                                                        }
                                                                    }
                                                                default:
                                                                    var v79 uint64 = 4294967295
                                                                    if true {
                                                                        var v80 int64 = 22
                                                                        v81 := GenS9{F0: (fi53 >> ((18446744073709551615) & 63))}
                                                                        switch (fi54 - 88) {
                                                                        case 128:
                                                                            v82 := GenS10{F0: 284}
                                                                            _ = v82
                                                                        case 32767:
                                                                            switch v55 {
                                                                            case 4294967295:
                                                                                var v83 string = v8
                                                                                _ = v83
                                                                            case 0:
                                                                                switch (v55 | int(v79)) {
                                                                                case -870:
                                                                                    if (0.0 <= (((1e10 * -2.25) / 1.0) * 1e10)) {
                                                                                        v84 := GenS10{F0: 2}
                                                                                        _ = v84
                                                                                    } else {
                                                                                        v79 = 18446744073709551615
                                                                                        var v85 int64 = (v80 ^ v80)
                                                                                        var v86 int32 = 191
                                                                                        fmt.Println(v79, v1)
                                                                                        if (!false) {
                                                                                            v87 := GenS11{F0: false, F1: "a"}
                                                                                            _ = len(v3)
                                                                                            _ = v87
                                                                                        } else {
                                                                                            var v88 float32 = (v6 + v6)
                                                                                            var v89 int = (^(v55))
                                                                                            _ = v88
                                                                                            _ = v89
                                                                                        }
                                                                                        _ = v85
                                                                                        _ = v86
                                                                                    }
                                                                                case 4294967295:
                                                                                    fmt.Println(v77, v1)
                                                                                }
                                                                            }
                                                                        }
                                                                        var v90 float32 = v6
                                                                        _ = v80
                                                                        _ = v81
                                                                        _ = v90
                                                                    } else {
                                                                        v4 += (v2 & (^(v2)))
                                                                    }
                                                                    if true {
                                                                        if (!false) {
                                                                            v2 = (v1 & v2)
                                                                        }
                                                                        var v91 string = ""
                                                                        switch (fi54 & 1) {
                                                                        case 127:
                                                                            fmt.Println(v55, v77)
                                                                        case 2147483647:
                                                                            var v92 uint8 = v77
                                                                            _ = v92
                                                                        case -9223372036854775808:
                                                                            if (v4 != (v4 & (v4 - 42))) {
                                                                                v93 := GenS2{F0: 240, F1: 9223372036854775807, F2: 2147483647}
                                                                                fmt.Println(v8, v2, v8)
                                                                                _ = v93
                                                                            }
                                                                            _ = len(v7)
                                                                            _ = len(v7)
                                                                        }
                                                                        _ = v91
                                                                    } else {
                                                                        v94 := GenS2{F0: 255, F1: 0, F2: 2}
                                                                        _ = v94
                                                                    }
                                                                    _ = v79
                                                                }
                                                                v95 := v3[1]
                                                                v1 = v1
                                                                _ = v77
                                                                _ = v95
                                                            }
                                                        } else {
                                                            v2 += 94
                                                            if false {
                                                                v96 := GenS4{F0: 249, F1: 24}
                                                                v4 = (v1 | v4)
                                                                _ = v96
                                                            }
                                                            switch v8 {
                                                            case "":
                                                                if (!(!(v6 > (v6 - ((1.5 + 1e10) - float32(1.5)))))) {
                                                                    v97 := v5[2147483648]
                                                                    v6 = float32(float64(v6))
                                                                    if len(v7) >= 2 {
                                                                        v7 = v7[1:len(v7)]
                                                                    }
                                                                    var v98 int64 = 0
                                                                    _ = v97
                                                                    _ = v98
                                                                }
                                                            }
                                                        }
                                                        var v99 float32 = (v6 * (((0.0 + v6) - (v6 - v6)) * v6))
                                                        if ((!false) || true) {
                                                            if (!false) {
                                                                v100 := GenS11{F0: (712 == 848), F1: v8}
                                                                v101 := v5[7]
                                                                var v102 int8 = v2
                                                                var v103 string = ("a" + "a")
                                                                v1 = v102
                                                                fmt.Println(v55, fi53, v2)
                                                                var v104 float32 = 3.14159
                                                                _ = v100
                                                                _ = v101
                                                                _ = v102
                                                                _ = v103
                                                                _ = v104
                                                            } else {
                                                                var v105 uint32 = 42
                                                                fmt.Println(fi54)
                                                                v106 := GenS1{F0: 18446744073709551615, F1: -1}
                                                                _ = v105
                                                                _ = v106
                                                            }
                                                            switch (("" + "\n\t") + v8) {
                                                            case "foo bar":
                                                                var v107 float32 = 0.1
                                                                _ = v107
                                                            case "":
                                                                fmt.Println(v4, v2)
                                                            }
                                                        }
                                                        _ = v99
                                                    default:
                                                        var v108 int8 = (v4 + (v2 & (v1 - -77)))
                                                        _ = v108
                                                    }
                                                }
                                            }
                                            fmt.Println(v8)
                                        }
                                    }
                                }
                                var v109 bool = true
                                _ = fi54
                                _ = v109
                                _ = v55
                            }
                            var v110 float64 = 0.0
                            _ = fi53
                            _ = v110
                        }
                        fmt.Println(v2, v6)
                        fmt.Println(v8, v4)
                    } else {
                        if (!(42 != 0)) {
                            switch (v8 + "") {
                            case "":
                                var v111 int64 = 0
                                _ = v111
                            case "\n\t":
                                if len(v7) > 0 {
                                    v7[0] = 898
                                }
                            case "a":
                                v112 := v5[2147483647]
                                _ = v112
                            }
                        }
                        for fi113 := 0; fi113 < 9; fi113++ {
                            v6 = v6
                            _ = fi113
                        }
                    }
                }
                v1 += -51
            }
        case "a":
            for fi114 := 0; fi114 < 17; fi114++ {
                var v115 int = (fi114 / ((-1) | 1))
                switch ("hello" + (((((v8 + ("foo bar" + ("foo bar" + v8))) + "a") + "") + "") + v8)) {
                case "\n\t":
                    fmt.Println(v6, v6)
                case "hello":
                    var v116 int16 = 486
                    _ = v116
                default:
                    for fi117 := 0; fi117 < 1; fi117++ {
                        var v118 float64 = (3.14159 + 3.14159)
                        v6 -= (v6 * v6)
                        _ = fi117
                        _ = v118
                    }
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi114
                _ = v115
            }
        default:
            v6 += v6
        }
        switch 2147483647 {
        case 37:
            v119 := GenS5{F0: 1, F1: (((v8 + (v8 + v8)) + v8) + "a"), F2: 618}
            _ = v119
        }
    }
    var v120 float64 = (1e10 * ((1.5 - -2.25) / 1.0))
    if ((684 < 255) && false) {
        for fi121 := 0; fi121 < 9; fi121++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi121
        }
    }
    if (true && (1 >= 18446744073709551615)) {
        if (1000 >= 1) {
            v120 = (v120 + v120)
        } else {
            for fi122 := 0; fi122 < 20; fi122++ {
                v6 = (v6 + 0.1)
                _ = fi122
            }
            var v123 uint32 = 4294967295
            _ = v123
        }
        for fi124 := 0; fi124 < 19; fi124++ {
            _ = len(v7)
            v125 := GenS5{F0: 65535, F1: (v8 + ""), F2: 18446744073709551615}
            v126 := GenS8{F0: 1, F1: 308, F2: true}
            _ = fi124
            _ = v125
            _ = v126
        }
        var v127 string = (v8 + v8)
        _ = v127
    }
    var v128 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v120
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
