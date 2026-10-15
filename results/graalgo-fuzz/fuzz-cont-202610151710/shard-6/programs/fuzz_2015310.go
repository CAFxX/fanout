package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 string; F2 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool; F1 int16; F2 uint8 }
type GenS10 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS11 struct { F0 bool; F1 int8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int16; F1 string }
type GenS14 struct { F0 int64; F1 uint32; F2 int16 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS17 struct { F0 string; F1 int; F2 uint8 }
type GenS18 struct { F0 int16; F1 uint; F2 int }
type GenS19 struct { F0 string }
type GenS20 struct { F0 string; F1 int16 }

func main() {
    var v1 int = 2147483647
    var v2 uint16 = 299
    var v3 float32 = (1e10 - 3.14159)
    var v4 float64 = (0.1 + (0.0 - (1.5 * ((1e10 / ((((((((3.14159 + (3.14159 - 1e10)) - 0.0) * 3.14159) - ((((0.1 * (0.1 / 0.1)) * (float64(v3) - 0.1)) / (1.5 * 3.14159)) * -2.25)) / 1e10) - (float64(1e10) * (1e10 + (3.14159 - 0.1)))) * 1.5) / -2.25)) * 0.0))))
    v5 := []uint32{0, 2147483647, 661, 825}
    var v6 int8 = 7
    var v7 uint64 = 0
    var v8 float64 = float64(1.5)
    switch v1 {
    case 646:
        switch "hello" {
        case "":
            fmt.Println(v4, v3, v7)
        }
    default:
        if true {
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 int16 = 0
                if true {
                    fmt.Println(v10, v6, v4)
                    v2 -= (v2 ^ 32767)
                }
                _ = fi9
                _ = v10
            }
        } else {
            if (!(9223372036854775807 > 223)) {
                var v11 int16 = -32768
                _ = v11
            } else {
                fmt.Println(v8, v2)
                var v12 int16 = 1000
                v12 = (v12 & -1)
                _ = v12
            }
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS1{F0: 431}
                switch "hello" {
                case "\n\t":
                    if (-1 >= int64(v6)) {
                        v1 = (fi13 >> ((1) & 63))
                    } else {
                        fmt.Println(v8, fi13)
                        var v15 float64 = v4
                        var v16 uint16 = (v2 << ((2) & 63))
                        _ = v15
                        _ = v16
                    }
                default:
                    if true {
                        v1 = -675
                        fmt.Println(v8)
                        var v17 float64 = v8
                        _ = v17
                    } else {
                        v4 = ((v4 + v8) * 0.1)
                    }
                }
                _ = fi13
                _ = v14
            }
            var v18 bool = (353 != 42)
            _ = v18
        }
    }
    v19 := GenS2{F0: 212}
    switch (v1 * 0) {
    case 9223372036854775807:
        switch (v1 % ((v1) | 1)) {
        case -234:
            fmt.Println(v8, v2, v4)
        case -80:
            fmt.Println(v4, v2, v7)
        case 2147483647:
            var v20 string = ""
            _ = v20
        }
    case -9223372036854775808:
        var v21 uint = 1
        _ = v21
    }
    v3 = 3.14159
    v6 += (v6 >> ((1) & 63))
    fmt.Println(v3)
    if true {
        var v22 int = (-(v1))
        for fi23 := 0; fi23 < 16; fi23++ {
            v6 = v6
            _ = fi23
        }
        _ = v22
    } else {
        switch v1 {
        case -2147483648:
            switch ("a" + "") {
            case "\n\t":
                v24 := GenS3{F0: (v2 ^ (v2 / ((956) | 1))), F1: "hello", F2: v1}
                _ = v24
            case "foo bar":
                if len(v5) > 0 {
                    v5[0] = 65535
                }
            case "a":
                switch "a" {
                case "a":
                    var v25 int16 = -350
                    _ = v25
                case "hello":
                    v6 = (^(v6))
                default:
                    if (-1 < -32768) {
                        if (false && (!(true || true))) {
                            fmt.Println(v2, v8, v2)
                            fmt.Println(v8, v2, v6)
                        } else {
                            switch "a" {
                            case "":
                                v3 += 3.14159
                            default:
                                var v26 string = ""
                                _ = v26
                            }
                            var v27 uint32 = 65535
                            v1 = (^(v1))
                            _ = v27
                        }
                    } else {
                        for fi28 := 0; fi28 < 18; fi28++ {
                            if (v3 >= v3) {
                                var v29 string = "hello"
                                _ = v29
                            } else {
                                for fi30 := 0; fi30 < 10; fi30++ {
                                    v3 += v3
                                    v31 := GenS3{F0: v2, F1: "\n\t", F2: fi30}
                                    _ = fi30
                                    _ = v31
                                }
                                _ = len(v5)
                                for fi32 := 0; fi32 < 7; fi32++ {
                                    var v33 int64 = 9223372036854775807
                                    _ = fi32
                                    _ = v33
                                }
                            }
                            fmt.Println(v1)
                            _ = fi28
                        }
                        var v34 bool = false
                        _ = v34
                    }
                }
            default:
                var v35 float32 = (v3 * 0.0)
                _ = v35
            }
        case -9223372036854775808:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case -338:
            var v36 int = v1
            _ = v36
        default:
            for fi37 := 0; fi37 < 4; fi37++ {
                if len(v5) > 0 {
                    v5[0] = 426
                }
                fmt.Println(v8)
                for fi38 := 0; fi38 < 8; fi38++ {
                    v39 := GenS4{F0: -295}
                    _ = fi38
                    _ = v39
                }
                _ = fi37
            }
        }
    }
    for fi40 := 0; fi40 < 16; fi40++ {
        var v41 int8 = v6
        var v42 bool = (1 >= 2)
        _ = fi40
        _ = v41
        _ = v42
    }
    v43 := GenS3{F0: (v2 ^ v2), F1: "hello", F2: int(v6)}
    v1 -= (v1 << ((397) & 63))
    for fi44 := 0; fi44 < 3; fi44++ {
        v1 -= 42
        for fi45 := 0; fi45 < 11; fi45++ {
            fmt.Println(v8, fi44, fi44)
            if (-9223372036854775808 < 1) {
                var v46 int64 = -9223372036854775808
                _ = v46
            } else {
                if (!(!true)) {
                    _ = len(v5)
                    _ = len(v5)
                    if false {
                        v7 = v7
                        switch ("\n\t" + (("hello" + "\n\t") + "")) {
                        case "\n\t":
                            v8 += (v4 - float64(v3))
                        }
                    } else {
                        var v47 float32 = 3.14159
                        v5 = append(v5, 4294967295)
                        v48 := GenS5{F0: v6}
                        v49 := GenS6{F0: 255}
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                }
            }
            v3 = v3
            _ = fi45
        }
        _ = fi44
    }
    var v50 int16 = 1
    for fi51 := 0; fi51 < 16; fi51++ {
        var v52 uint16 = v2
        switch 32768 {
        case 612:
            if (!true) {
                v50 -= (-(v50))
                v52 = (^(v2))
                v3 += 3.14159
            }
        case 301:
            if true {
                var v53 bool = true
                _ = v53
            } else {
                v54 := GenS7{F0: 1, F1: 959}
                var v55 uint32 = 113
                if false {
                    for fi56 := 0; fi56 < 2; fi56++ {
                        v57 := GenS8{F0: 65535}
                        var v58 uint32 = 2
                        v59 := GenS9{F0: true, F1: -32768, F2: 42}
                        _ = fi56
                        _ = v57
                        _ = v58
                        _ = v59
                    }
                    v60 := GenS8{F0: 65535}
                    _ = v60
                } else {
                    v55 -= 898
                    if (v1 > fi51) {
                        v1 -= 127
                        if false {
                            var v61 uint8 = 127
                            v6 -= (v6 | v6)
                            _ = v61
                        } else {
                            if (!(2147483647 < -1)) {
                                if false {
                                    switch v1 {
                                    case 32768:
                                        v2 = v2
                                    case 2147483647:
                                        for fi62 := 0; fi62 < 3; fi62++ {
                                            v1 -= v1
                                            v63 := GenS2{F0: 36}
                                            fmt.Println(v3, v6)
                                            _ = fi62
                                            _ = v63
                                        }
                                    case -2147483648:
                                        v3 -= v3
                                    }
                                    fmt.Println(fi51, v4, v52)
                                }
                            } else {
                                for fi64 := 0; fi64 < 10; fi64++ {
                                    fmt.Println(v6)
                                    var v65 uint8 = 143
                                    v66 := GenS2{F0: (v65 ^ (v65 >> ((4294967295) & 63)))}
                                    _ = fi64
                                    _ = v65
                                    _ = v66
                                }
                                v50 -= v50
                            }
                        }
                        fmt.Println(v3, v4, v2)
                    } else {
                        var v67 bool = (v4 >= (3.14159 / (v4 + 3.14159)))
                        _ = v67
                    }
                    if false {
                        fmt.Println(v52, v1)
                        switch "a" {
                        case "a":
                            v55 -= (v55 * 604)
                        case "\n\t":
                            v55 = 65535
                        case "foo bar":
                            v7 = 2
                        default:
                            v5 = append(v5, 0)
                        }
                        v2 -= v52
                    } else {
                        fmt.Println(v7)
                        v68 := GenS7{F0: 255, F1: 2147483647}
                        for fi69 := 0; fi69 < 8; fi69++ {
                            v70 := GenS1{F0: (v50 + (^(v50)))}
                            if ((!false) || true) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            } else {
                                var v71 string = "hello"
                                _ = v71
                            }
                            fmt.Println(v50)
                            var v72 float32 = v3
                            _ = fi69
                            _ = v70
                            _ = v72
                        }
                        _ = v68
                    }
                }
                _ = v54
                _ = v55
            }
        case 475:
            var v73 uint16 = 2
            _ = v73
        }
        _ = fi51
        _ = v52
    }
    v7 = (v7 << ((115) & 63))
    switch ("a" + (("a" + "\n\t") + "\n\t")) {
    case "foo bar":
        for fi74 := 0; fi74 < 17; fi74++ {
            switch "hello" {
            case "a":
                var v75 uint = 2
                _ = v75
            case "hello":
                switch fi74 {
                case 240:
                    fmt.Println(v3, v50)
                case -2147483648:
                    v1 = int(v6)
                case -1:
                    for fi76 := 0; fi76 < 7; fi76++ {
                        v2 = 65535
                        v2 = v2
                        var v77 bool = (v50 < (v50 | -1))
                        if true {
                            if ((v2 == (v2 | v2)) || v77) {
                                v77 = ((float32(v4) - v3) == v3)
                                if (v7 != v7) {
                                    v78 := GenS4{F0: int(v6)}
                                    if (42 < 18446744073709551615) {
                                        v5 = append(v5, uint32(v7))
                                        if false {
                                            if true {
                                                if true {
                                                    if false {
                                                        v79 := GenS10{F0: v50, F1: 1, F2: v7}
                                                        _ = v79
                                                    }
                                                    if v77 {
                                                        v2 = v2
                                                        if (835 < 957) {
                                                            var v80 uint8 = 0
                                                            var v81 uint64 = v7
                                                            if (!((false && true) && ((^(v81)) != (^(v81))))) {
                                                                if (((^(v6)) > v6) && false) {
                                                                    v6 = (v6 % ((v6) | 1))
                                                                }
                                                            }
                                                            _ = v80
                                                            _ = v81
                                                        } else {
                                                            _ = len(v5)
                                                            v2 -= v2
                                                        }
                                                        _ = len(v5)
                                                    }
                                                } else {
                                                    switch (("\n\t" + "hello") + ("a" + "\n\t")) {
                                                    case "foo bar":
                                                        var v82 uint32 = 42
                                                        _ = v82
                                                    case "\n\t":
                                                        switch "hello" {
                                                        case "\n\t":
                                                            var v83 int8 = 1
                                                            _ = v83
                                                        default:
                                                            v4 -= 1e10
                                                        }
                                                    case "hello":
                                                        v3 = 1e10
                                                    }
                                                }
                                                fmt.Println(v4, v6)
                                                switch -9223372036854775808 {
                                                case 0:
                                                    v7 -= (v7 % ((v7) | 1))
                                                default:
                                                    v6 -= -1
                                                }
                                                var v84 float32 = 0.0
                                                _ = v84
                                            } else {
                                                var v85 uint = 1
                                                switch (v1 * (^(fi76))) {
                                                case 0:
                                                    if ((!(2 < v7)) && v77) {
                                                        fmt.Println(v4)
                                                        v77 = v77
                                                    } else {
                                                        _ = len(v5)
                                                    }
                                                case -1:
                                                    v86 := GenS11{F0: v77, F1: int8(v50)}
                                                    _ = v86
                                                case 42:
                                                    var v87 uint = v85
                                                    _ = v87
                                                }
                                                fmt.Println(fi74)
                                                _ = v85
                                            }
                                        } else {
                                            v88 := GenS11{F0: true, F1: v6}
                                            var v89 uint = 923
                                            _ = v88
                                            _ = v89
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = 2147483648
                                        }
                                        var v90 bool = true
                                        _ = v90
                                    } else {
                                        switch -559 {
                                        case -384:
                                            switch ("\n\t" + "foo bar") {
                                            case "a":
                                                fmt.Println(v7, fi74, v2)
                                            case "\n\t":
                                                v91 := GenS12{F0: 149}
                                                _ = v91
                                            default:
                                                var v92 uint32 = 285
                                                if true {
                                                    v6 += (^(v6))
                                                    v5 = append(v5, (v92 - (v92 << ((1) & 63))))
                                                    v7 = uint64(v2)
                                                } else {
                                                    fmt.Println(v2)
                                                    if true {
                                                        var v93 float32 = ((v3 - v3) * v3)
                                                        _ = v93
                                                    } else {
                                                        _ = len(v5)
                                                        var v94 float64 = (float64(v3) + v8)
                                                        var v95 int8 = int8(v7)
                                                        v96 := GenS9{F0: true, F1: int16(v7), F2: 2}
                                                        v2 += v2
                                                        v5 = append(v5, v92)
                                                        _ = v94
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                    if false {
                                                        if (-9223372036854775808 >= 1) {
                                                            fmt.Println(v4)
                                                            v92 = (v92 >> ((2) & 63))
                                                            fmt.Println(v4)
                                                        }
                                                    } else {
                                                        if v77 {
                                                            _ = len(v5)
                                                            switch v1 {
                                                            case 0:
                                                                if (("\n\t" >= ("\n\t" + "")) || false) {
                                                                    v97 := GenS4{F0: fi76}
                                                                    var v98 int64 = int64(v7)
                                                                    v50 = v50
                                                                    _ = v97
                                                                    _ = v98
                                                                } else {
                                                                    if v77 {
                                                                        var v99 int32 = -2147483648
                                                                        if false {
                                                                            if true {
                                                                                v100 := GenS7{F0: uint8(v2), F1: (v99 >> ((910) & 63))}
                                                                                _ = v100
                                                                            } else {
                                                                                if false {
                                                                                    v92 -= v92
                                                                                    var v101 int32 = (v99 << ((uint(v92)) & 63))
                                                                                    _ = v101
                                                                                }
                                                                            }
                                                                            if (0.0 > v3) {
                                                                                var v102 uint8 = 1
                                                                                var v103 int32 = (v99 - 737)
                                                                                fmt.Println(v99, fi74)
                                                                                v5 = append(v5, v92)
                                                                                switch "" {
                                                                                case "a":
                                                                                    v1 -= -1
                                                                                }
                                                                                _ = v102
                                                                                _ = v103
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v77, v2)
                                                                            var v104 float64 = (v8 * v8)
                                                                            v3 = v3
                                                                            v105 := GenS10{F0: (v50 >> ((0) & 63)), F1: 2, F2: v7}
                                                                            v2 = (^(v2))
                                                                            _ = v104
                                                                            _ = v105
                                                                        }
                                                                        _ = v99
                                                                    }
                                                                    fmt.Println(v4, v2)
                                                                    _ = len(v5)
                                                                    var v106 float64 = v8
                                                                    var v107 float64 = float64(1.5)
                                                                    switch v1 {
                                                                    case 998:
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            v106 -= (v8 + v106)
                                                                        case "":
                                                                            var v108 uint16 = (v2 ^ (v2 + v2))
                                                                            _ = v108
                                                                        default:
                                                                            v109 := GenS7{F0: 218, F1: 42}
                                                                            _ = v109
                                                                        }
                                                                    case -2147483648:
                                                                        var v110 float64 = 3.14159
                                                                        _ = v110
                                                                    }
                                                                    _ = v106
                                                                    _ = v107
                                                                }
                                                            case 1:
                                                                _ = len(v5)
                                                            default:
                                                                if ((v77 && true) || (100 < v6)) {
                                                                    var v111 float32 = ((0.1 - v3) + 1.5)
                                                                    var v112 int64 = 9223372036854775807
                                                                    switch "" {
                                                                    case "a":
                                                                        var v113 uint = 1
                                                                        _ = v113
                                                                    }
                                                                    _ = v111
                                                                    _ = v112
                                                                }
                                                            }
                                                            v114 := GenS13{F0: v50, F1: "foo bar"}
                                                            _ = v114
                                                        }
                                                        switch fi76 {
                                                        case 32767:
                                                            v3 = (float32(float64(float32(3.14159))) + (float32((v4 / ((3.14159 + v8) * ((v8 + 0.1) + float64((1e10 + ((v3 * v3) / ((-2.25 / float32(1e10)) - v3)))))))) * v3))
                                                        case -1:
                                                            var v115 uint8 = 127
                                                            switch ("a" + "\n\t") {
                                                            case "":
                                                                v2 = 955
                                                            default:
                                                                fmt.Println(v6, v1)
                                                            }
                                                            _ = v115
                                                        }
                                                        if ((v6 != v6) && true) {
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            var v116 string = ""
                                                            _ = v116
                                                        }
                                                    }
                                                }
                                                _ = len(v5)
                                                var v117 bool = false
                                                _ = v117
                                                _ = v92
                                            }
                                        }
                                        var v118 uint64 = v7
                                        switch 1 {
                                        case 32767:
                                            var v119 float32 = (1.5 * v3)
                                            v8 += (3.14159 / (1e10 / float64(1e10)))
                                            v118 += 2
                                            _ = v119
                                        case 181:
                                            v120 := GenS14{F0: -396, F1: 2147483647, F2: (-(v50))}
                                            _ = v120
                                        }
                                        _ = v118
                                    }
                                    v121 := GenS15{F0: 42}
                                    _ = v121
                                    _ = v78
                                } else {
                                    var v122 string = "foo bar"
                                    v6 = v6
                                    var v123 float64 = v4
                                    var v124 int8 = v6
                                    _ = v122
                                    _ = v123
                                    _ = v124
                                }
                            } else {
                                if false {
                                    fmt.Println(v3, v1)
                                    var v125 int32 = 42
                                    _ = v125
                                }
                                v2 -= v2
                                var v126 int16 = -537
                                _ = v126
                            }
                            if v77 {
                                if ((v6 / ((v6) | 1)) >= (v6 | (v6 >> ((65535) & 63)))) {
                                    _ = len(v5)
                                } else {
                                    v6 = 0
                                    v1 += -2147483648
                                }
                                v127 := GenS16{F0: (v6 * v6), F1: (v7 << ((2) & 63)), F2: 929}
                                _ = v127
                            } else {
                                if true {
                                    switch ("foo bar" + ("a" + ("\n\t" + ""))) {
                                    case "foo bar":
                                        v128 := GenS7{F0: 255, F1: -176}
                                        _ = v128
                                    }
                                    v6 += (v6 + 127)
                                    _ = len(v5)
                                }
                            }
                        }
                        _ = fi76
                        _ = v77
                    }
                default:
                    for fi129 := 0; fi129 < 5; fi129++ {
                        v5 = append(v5, uint32(v2))
                        _ = fi129
                    }
                }
            }
            var v130 int = (v1 >> ((834) & 63))
            if false {
                switch int(v7) {
                case 2147483647:
                    v131 := GenS17{F0: "", F1: fi74, F2: 255}
                    _ = v131
                case 32767:
                    _ = len(v5)
                }
                var v132 uint8 = 127
                _ = v132
            } else {
                if false {
                    v1 -= 42
                    fmt.Println(v3, v1)
                }
                v130 += (fi74 + 149)
            }
            _ = fi74
            _ = v130
        }
    case "hello":
        if (v6 == v6) {
            if false {
                for fi133 := 0; fi133 < 17; fi133++ {
                    if true {
                        if true {
                            var v134 uint64 = v7
                            v8 -= float64(float32(0.0))
                            for fi135 := 0; fi135 < 9; fi135++ {
                                var v136 string = ("" + "")
                                _ = fi135
                                _ = v136
                            }
                            _ = v134
                        } else {
                            var v137 string = ("a" + "\n\t")
                            _ = len(v5)
                            _ = v137
                        }
                        var v138 int8 = (^(v6))
                        v139 := GenS18{F0: -823, F1: 0, F2: (v1 >> ((1) & 63))}
                        _ = v138
                        _ = v139
                    } else {
                        v6 -= (v6 >> ((42) & 63))
                        var v140 int8 = (v6 & 127)
                        var v141 float64 = ((float64(v3) * 1.5) / v4)
                        _ = v140
                        _ = v141
                    }
                    v6 -= -128
                    _ = fi133
                }
                v3 += v3
            }
            for fi142 := 0; fi142 < 15; fi142++ {
                v4 = v4
                v143 := GenS7{F0: 255, F1: 594}
                _ = len(v5)
                _ = fi142
                _ = v143
            }
            if false {
                switch -1 {
                case 4294967295:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case 2147483647:
                    fmt.Println(v2)
                }
                for fi144 := 0; fi144 < 16; fi144++ {
                    v8 -= (v8 / v8)
                    _ = fi144
                }
            }
        }
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = 2
        }
    default:
        v145 := GenS19{F0: "a"}
        _ = v145
    }
    switch "hello" {
    case "foo bar":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        if (-104 == 0) {
            v146 := GenS20{F0: "", F1: -472}
            var v147 uint16 = 42
            _ = len(v5)
            _ = v146
            _ = v147
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v50
    _ = v6
    _ = v7
    _ = v8
}
