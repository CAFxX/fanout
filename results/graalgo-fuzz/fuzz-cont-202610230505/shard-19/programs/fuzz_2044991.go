package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 int8; F2 uint8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int8; F1 int }
type GenS12 struct { F0 string; F1 int32 }
type GenS13 struct { F0 int32 }
func genFn0() uint64 {
    var v5 uint32 = 42
    for fi6 := 0; fi6 < 5; fi6++ {
        switch fi6 {
        case 0:
            var v7 int16 = 1000
            _ = v7
        case -321:
            switch fi6 {
            case 32767:
                v8 := GenS1{F0: ("" + "foo bar")}
                _ = v8
            }
        case 32767:
            var v9 string = "foo bar"
            _ = v9
        }
        var v10 bool = true
        v5 = 828
        _ = fi6
        _ = v10
    }
    _ = v5
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{128, 42, 4294967295, -9223372036854775808}
    var v3 bool = (42 > -9223372036854775808)
    var v4 uint32 = 78
    var v11 int16 = 1
    if (128 > 32768) {
        var v12 uint = 42
        v4 = v4
        var v13 int = 343
        _ = v12
        _ = v13
    } else {
        v3 = v3
        for fi14 := 0; fi14 < 8; fi14++ {
            if v3 {
                var v15 float32 = 0.0
                _ = v15
            } else {
                if v3 {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        var v17 string = ("hello" + (("\n\t" + "foo bar") + "a"))
                        _ = fi16
                        _ = v17
                    }
                    switch (fi14 - fi14) {
                    case -772:
                        v4 = (v4 + (v4 ^ v4))
                    case -458:
                        v2 = append(v2, fi14)
                    case 512:
                        if false {
                            v18 := GenS2{F0: 1, F1: 255, F2: 0}
                            v2 = append(v2, 127)
                            var v19 int32 = 2147483647
                            _ = v18
                            _ = v19
                        }
                    }
                } else {
                    switch -555 {
                    case 2147483647:
                        if ((v3 || v3) && ((false || v3) && false)) {
                            switch "foo bar" {
                            case "a":
                                v3 = false
                            default:
                                delete(v1, (fi14 / ((fi14) | 1)))
                            }
                        }
                    default:
                        delete(v1, fi14)
                    }
                }
                var v20 uint32 = (v4 + v4)
                var v21 int16 = 1
                _ = v20
                _ = v21
            }
            var v22 float32 = 3.14159
            if (3.14159 > 1.5) {
                var v23 int64 = 9223372036854775807
                _ = v23
            } else {
                var v24 int16 = v11
                switch "hello" {
                case "\n\t":
                    v24 -= -680
                case "a":
                    var v25 uint32 = (v4 << ((4294967295) & 63))
                    _ = v25
                case "foo bar":
                    var v26 float32 = float32(1e10)
                    _ = v26
                }
                _ = v24
            }
            _ = fi14
            _ = v22
        }
    }
    v27 := GenS3{F0: 42}
    v3 = ((43 <= -93) && v3)
    fmt.Println(v4)
    for fi28 := 0; fi28 < 17; fi28++ {
        for fi29 := 0; fi29 < 13; fi29++ {
            switch fi28 {
            case 4294967295:
                v4 -= v4
            default:
                fmt.Println(v11, fi28, fi29)
            }
            _ = fi29
        }
        var v30 int32 = -446
        var v31 uint16 = 0
        _ = fi28
        _ = v30
        _ = v31
    }
    var v32 float64 = -2.25
    v3 = (32767 != 145)
    if len(v2) > 0 {
        v2[0] = 127
    }
    fmt.Println(v11, v4, v3)
    fmt.Println(v3, v3)
    for fi33 := 0; fi33 < 3; fi33++ {
        var v34 uint = 0
        _ = fi33
        _ = v34
    }
    if false {
        switch 32767 {
        case -1:
            for fi35 := 0; fi35 < 20; fi35++ {
                if (42 <= 4294967295) {
                    _ = len(v2)
                }
                switch ("hello" + ("hello" + ((("\n\t" + "\n\t") + "") + "foo bar"))) {
                case "hello":
                    for fi36 := 0; fi36 < 15; fi36++ {
                        var v37 int8 = 0
                        _ = fi36
                        _ = v37
                    }
                default:
                    if v3 {
                        for fi38 := 0; fi38 < 4; fi38++ {
                            v39 := GenS4{F0: 0}
                            _ = fi38
                            _ = v39
                        }
                        for fi40 := 0; fi40 < 8; fi40++ {
                            var v41 uint64 = 1
                            v41 -= (v41 / ((v41) | 1))
                            var v42 float64 = v32
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        _ = len(v1)
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            v43 := GenS3{F0: 4294967295}
                            _ = v43
                        case "foo bar":
                            v4 = (v4 | 370)
                        case "hello":
                            var v44 string = ((("\n\t" + "") + "") + "hello")
                            _ = v44
                        default:
                            fmt.Println(v11)
                        }
                    }
                }
                _ = fi35
            }
        case 2147483647:
            var v45 string = "\n\t"
            _ = v45
        case 117:
            switch "" {
            case "\n\t":
                v46 := GenS5{F0: ("" + ("\n\t" + ""))}
                _ = v46
            case "foo bar":
                for fi47 := 0; fi47 < 14; fi47++ {
                    v11 = (v11 & -1)
                    for fi48 := 0; fi48 < 5; fi48++ {
                        v11 = (v11 ^ (^(v11)))
                        _ = fi48
                    }
                    _ = fi47
                }
            case "hello":
                v49 := GenS1{F0: "hello"}
                _ = v49
            }
        default:
            for fi50 := 0; fi50 < 13; fi50++ {
                v51 := GenS6{F0: 337, F1: 1, F2: v11}
                switch (fi50 & 32767) {
                case -18:
                    switch (-(fi50)) {
                    case 128:
                        v11 -= (v11 + v11)
                    case -1:
                        if false {
                            var v52 uint8 = 2
                            _ = v52
                        } else {
                            v3 = true
                            v4 -= (v4 << ((2) & 63))
                        }
                    }
                case -2147483648:
                    v53 := GenS7{F0: 42}
                    _ = v53
                case -225:
                    v32 = 1.5
                default:
                    v54 := GenS8{F0: 556}
                    _ = v54
                }
                if ((("a" != ("" + "hello")) || (3.14159 <= (1.5 * ((float32((v32 + v32)) + 3.14159) / (-2.25 / 1e10))))) && false) {
                    v55 := GenS8{F0: 255}
                    for fi56 := 0; fi56 < 10; fi56++ {
                        var v57 float64 = (v32 / v32)
                        if (!(!(2 <= uint(v4)))) {
                            var v58 uint32 = v4
                            switch "\n\t" {
                            case "foo bar":
                                var v59 int8 = 92
                                _ = v59
                            case "":
                                if true {
                                    var v60 int = fi56
                                    switch 32767 {
                                    case 128:
                                        v61 := GenS6{F0: 42, F1: uint16(v58), F2: v11}
                                        _ = v61
                                    }
                                    _ = v60
                                }
                            case "hello":
                                if false {
                                    fmt.Println(v3, v4, fi50)
                                    v62 := GenS9{F0: "foo bar", F1: 1, F2: 255}
                                    _ = v62
                                } else {
                                    if v3 {
                                        v57 -= (((v32 * v57) / v57) / v32)
                                    }
                                    delete(v1, fi56)
                                }
                            }
                            v63 := GenS3{F0: 864}
                            _ = v58
                            _ = v63
                        }
                        var v64 bool = true
                        _ = fi56
                        _ = v57
                        _ = v64
                    }
                    v3 = (((3.14159 <= (3.14159 * 1e10)) && v3) && (544 == 1))
                    _ = v55
                } else {
                    v65 := GenS10{F0: 255}
                    v66 := GenS4{F0: 42}
                    if (0 != 2) {
                        for fi67 := 0; fi67 < 1; fi67++ {
                            fmt.Println(v3, v11)
                            _ = fi67
                        }
                        v11 += (v11 << ((255) & 63))
                        var v68 int8 = 1
                        _ = v68
                    }
                    _ = v65
                    _ = v66
                }
                _ = fi50
                _ = v51
            }
        }
        if false {
            v69 := GenS10{F0: 42}
            var v70 int64 = -1
            v4 += v4
            _ = v69
            _ = v70
        }
        _ = len(v2)
    } else {
        if ((212 < 2) || v3) {
            var v71 uint = 65535
            _ = v71
        } else {
            fmt.Println(v32, v3)
        }
        fmt.Println(v11, v11, v4)
    }
    var v72 uint = 2
    for fi73 := 0; fi73 < 15; fi73++ {
        v11 -= 32767
        _ = fi73
    }
    switch "\n\t" {
    case "a":
        var v74 int32 = -10
        _ = v74
    }
    _ = len(v2)
    switch 32767 {
    case -9223372036854775808:
        for fi75 := 0; fi75 < 1; fi75++ {
            v72 += v72
            v3 = (!(!(!v3)))
            if (2 <= 904) {
                for fi76 := 0; fi76 < 9; fi76++ {
                    v32 = -2.25
                    _ = fi76
                }
                fmt.Println(fi75, fi75, fi75)
                v3 = v3
            }
            _ = fi75
        }
    default:
        if false {
            for fi77 := 0; fi77 < 4; fi77++ {
                for fi78 := 0; fi78 < 7; fi78++ {
                    fmt.Println(v72, v32, v4)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v2) > 0 {
                        v2[0] = (fi77 / ((-1) | 1))
                    }
                    var v79 uint32 = v4
                    _ = fi78
                    _ = v79
                }
                _ = fi77
            }
            if ("foo bar" != "hello") {
                for fi80 := 0; fi80 < 15; fi80++ {
                    if (!(-106 != 127)) {
                        if (v3 || v3) {
                            if v3 {
                                for fi81 := 0; fi81 < 5; fi81++ {
                                    fmt.Println(v4, v3, v4)
                                    _ = fi81
                                }
                            } else {
                                var v82 float64 = v32
                                if true {
                                    v4 += v4
                                    var v83 float32 = 3.14159
                                    if true {
                                        if (true || v3) {
                                            if (v3 || false) {
                                                if len(v2) > 0 {
                                                    v2[0] = fi80
                                                }
                                            }
                                            switch -2147483648 {
                                            case 42:
                                                if (42 >= 1) {
                                                    var v84 uint = uint(v4)
                                                    _ = v84
                                                }
                                            case 128:
                                                v32 = (float64(-2.25) / ((-2.25 / v32) + (3.14159 * 1.5)))
                                            }
                                        } else {
                                            if (((fi80 >> ((267) & 63)) != fi80) || false) {
                                                if v3 {
                                                    switch fi80 {
                                                    case 128:
                                                        var v85 uint8 = 205
                                                        _ = v85
                                                    default:
                                                        var v86 int16 = 1000
                                                        _ = v86
                                                    }
                                                    switch int(v11) {
                                                    case 128:
                                                        var v87 int64 = -1
                                                        _ = v87
                                                    case -2147483648:
                                                        v3 = (24 != -85)
                                                    }
                                                    if (!v3) {
                                                        if v3 {
                                                            v88 := GenS3{F0: 103}
                                                            _ = v88
                                                        }
                                                    } else {
                                                        for fi89 := 0; fi89 < 18; fi89++ {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                if v3 {
                                                                    v82 = (float64(v83) * 1e10)
                                                                    var v90 float64 = v32
                                                                    var v91 uint = 222
                                                                    _ = v90
                                                                    _ = v91
                                                                }
                                                            default:
                                                                v32 = v32
                                                            }
                                                            v1[(fi89 & 172)] = true
                                                            var v92 uint16 = 255
                                                            _ = fi89
                                                            _ = v92
                                                        }
                                                        var v93 float32 = v83
                                                        if false {
                                                            fmt.Println(v11)
                                                            switch (fi80 * (fi80 << ((v72) & 63))) {
                                                            case 127:
                                                                var v94 uint16 = 42
                                                                _ = v94
                                                            case 42:
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            case 975:
                                                                v93 += 1.5
                                                            }
                                                            switch fi80 {
                                                            case 977:
                                                                v95 := GenS11{F0: -81, F1: fi80}
                                                                _ = v95
                                                            case -781:
                                                                for fi96 := 0; fi96 < 18; fi96++ {
                                                                    var v97 int16 = v11
                                                                    v82 = v32
                                                                    _ = fi96
                                                                    _ = v97
                                                                }
                                                            default:
                                                                if v3 {
                                                                    fmt.Println(v11, v72)
                                                                } else {
                                                                    var v98 uint8 = 42
                                                                    if v3 {
                                                                        var v99 uint16 = 34
                                                                        for fi100 := 0; fi100 < 4; fi100++ {
                                                                            v2 = append(v2, fi80)
                                                                            _ = fi100
                                                                        }
                                                                        var v101 float64 = 0.1
                                                                        _ = v101
                                                                        _ = v99
                                                                    }
                                                                    _ = v98
                                                                }
                                                            }
                                                        }
                                                        _ = v93
                                                    }
                                                }
                                            } else {
                                                v82 -= (v82 - float64(v83))
                                                v102 := GenS6{F0: 809, F1: 2, F2: 0}
                                                _ = v102
                                            }
                                        }
                                        if (1000 > -32768) {
                                            var v103 int32 = 0
                                            var v104 bool = false
                                            _ = v103
                                            _ = v104
                                        } else {
                                            v105 := GenS1{F0: "foo bar"}
                                            v72 = v72
                                            _ = v105
                                        }
                                    } else {
                                        switch "" {
                                        case "a":
                                            v106 := GenS12{F0: "", F1: 2147483647}
                                            _ = v106
                                        }
                                    }
                                    _ = v83
                                }
                                var v107 int32 = 0
                                _ = v107
                                _ = v82
                            }
                        } else {
                            var v108 float32 = 1.5
                            var v109 bool = (fi80 > (fi80 & (fi80 * 1)))
                            _ = v108
                            _ = v109
                        }
                        fmt.Println(v32, v11, v32)
                        if (194 < 4294967295) {
                            var v110 int64 = -1
                            _ = v110
                        } else {
                            if (0 < v72) {
                                for fi111 := 0; fi111 < 20; fi111++ {
                                    if (2147483647 >= (v4 ^ 131)) {
                                        v112 := GenS13{F0: -2147483648}
                                        var v113 string = ("" + (("\n\t" + (("a" + "a") + "\n\t")) + ("a" + "")))
                                        _ = v112
                                        _ = v113
                                    } else {
                                        switch (fi111 >> (((v72 << ((v72) & 63))) & 63)) {
                                        case 42:
                                            fmt.Println(v32, v72, fi80)
                                        case 32768:
                                            fmt.Println(v3, v11, fi111)
                                        case 2147483647:
                                            var v114 int64 = -9223372036854775808
                                            _ = v114
                                        }
                                    }
                                    var v115 int16 = (v11 ^ (v11 | v11))
                                    _ = fi111
                                    _ = v115
                                }
                                var v116 uint8 = 42
                                _ = v116
                            } else {
                                var v117 bool = false
                                for fi118 := 0; fi118 < 12; fi118++ {
                                    var v119 int32 = -760
                                    _ = fi118
                                    _ = v119
                                }
                                _ = v117
                            }
                        }
                    }
                    switch fi80 {
                    case -541:
                        v2 = append(v2, (fi80 - (fi80 << (((v72 ^ v72)) & 63))))
                    case 128:
                        if (128 == 128) {
                            var v120 float32 = 0.1
                            for fi121 := 0; fi121 < 4; fi121++ {
                                v3 = v3
                                _ = fi121
                            }
                            _ = v120
                        } else {
                            var v122 uint = (v72 - uint(v4))
                            for fi123 := 0; fi123 < 4; fi123++ {
                                var v124 uint = v122
                                var v125 float64 = 3.14159
                                if true {
                                    v4 = (^(v4))
                                    if (0 == -2147483648) {
                                        fmt.Println(fi80, v122)
                                        switch fi123 {
                                        case -582:
                                            v124 = (v122 % ((967) | 1))
                                        case -172:
                                            fmt.Println(v3, fi123)
                                        case -2147483648:
                                            v124 = (v72 | (v122 % (((v124 & v124)) | 1)))
                                        default:
                                            v126 := v1[int(v4)]
                                            _ = v126
                                        }
                                        v122 += (^(v122))
                                    }
                                }
                                _ = fi123
                                _ = v124
                                _ = v125
                            }
                            v122 = 18446744073709551615
                            _ = v122
                        }
                    case 2147483647:
                        var v127 int32 = 1
                        _ = v127
                    default:
                        v1[(fi80 * (fi80 | 0))] = (!(!true))
                    }
                    if ((("\n\t" + (("a" + ("foo bar" + "\n\t")) + "a")) + ("\n\t" + "a")) == ("foo bar" + "foo bar")) {
                        fmt.Println(v32, v32)
                    }
                    _ = fi80
                }
            } else {
                v3 = true
                switch 32768 {
                case 4294967295:
                    if (!v3) {
                        for fi128 := 0; fi128 < 11; fi128++ {
                            v4 += 555
                            v3 = (false || v3)
                            _ = fi128
                        }
                        for fi129 := 0; fi129 < 3; fi129++ {
                            v2 = append(v2, (^(fi129)))
                            _ = fi129
                        }
                        v11 = 578
                    } else {
                        v1[-1] = (-529 == int32(v72))
                        var v130 float32 = (float32(v32) - float32(v32))
                        _ = v130
                    }
                case 0:
                    var v131 string = ("foo bar" + "\n\t")
                    _ = v131
                }
            }
            var v132 uint8 = 129
            _ = v132
        } else {
            var v133 uint = 0
            var v134 bool = (2147483647 <= 2147483647)
            _ = v133
            _ = v134
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v72
}
