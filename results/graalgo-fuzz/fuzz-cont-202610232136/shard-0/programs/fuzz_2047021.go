package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 uint; F2 uint64 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16; F1 string }
type GenS9 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int16 }

func main() {
    var v1 uint16 = 0
    v2 := []uint32{65535}
    var v3 uint64 = uint64(v1)
    v3 = (v3 % ((136) | 1))
    v1 -= 898
    var v4 bool = false
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int = (fi5 % (((fi5 >> ((1) & 63))) | 1))
            var v8 float64 = ((0.1 / 1.0) - 3.14159)
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    switch "hello" {
    case "hello":
        if v4 {
            v1 -= 1
            var v9 uint32 = 406
            if v4 {
                if len(v2) > 0 {
                    v2[0] = v9
                }
                _ = len(v2)
            } else {
                if (!(2 == 1)) {
                    var v10 float64 = (float64(1e10) / 1.0)
                    _ = v10
                } else {
                    if true {
                        var v11 int32 = 0
                        _ = v11
                    }
                    if (((v3 >> ((uint(v9)) & 63)) <= v3) && v4) {
                        fmt.Println(v1)
                        _ = len(v2)
                        var v12 bool = false
                        _ = v12
                    }
                    if v4 {
                        for fi13 := 0; fi13 < 15; fi13++ {
                            for fi14 := 0; fi14 < 10; fi14++ {
                                if (!v4) {
                                    switch (fi14 & fi14) {
                                    case -1:
                                        var v15 int8 = 127
                                        if (!(1 == 255)) {
                                            var v16 string = "\n\t"
                                            var v17 int64 = 1
                                            _ = v16
                                            _ = v17
                                        }
                                        var v18 int = 128
                                        if true {
                                            if ("\n\t" > "hello") {
                                                if true {
                                                    var v19 uint = 333
                                                    _ = v19
                                                } else {
                                                    var v20 int16 = 577
                                                    v20 = (v20 << ((678) & 63))
                                                    _ = v20
                                                }
                                            }
                                            var v21 bool = (128 < 1)
                                            var v22 uint64 = v3
                                            _ = v21
                                            _ = v22
                                        } else {
                                            var v23 uint64 = v3
                                            fmt.Println(v4, v15)
                                            _ = v23
                                        }
                                        var v24 string = "a"
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if (((0.0 - 0.0) / 1e10) > 0.0) {
                                            v25 := GenS1{F0: v15, F1: v3}
                                            if (349 > (v9 - v9)) {
                                                var v26 int16 = -929
                                                if (255 != 0) {
                                                    v2 = append(v2, 459)
                                                } else {
                                                    var v27 uint64 = (v3 + v3)
                                                    var v28 uint32 = 1
                                                    _ = v27
                                                    _ = v28
                                                }
                                                _ = v26
                                            } else {
                                                var v29 uint64 = 18446744073709551615
                                                if (v24 >= (v24 + ("\n\t" + "hello"))) {
                                                    switch 255 {
                                                    case 9223372036854775807:
                                                        if true {
                                                            if v4 {
                                                                if (v15 > v15) {
                                                                    v30 := GenS1{F0: (v15 - (v15 & (v15 >> ((2) & 63)))), F1: v3}
                                                                    v24 = "a"
                                                                    var v31 string = (v24 + ("a" + (v24 + ("a" + (v24 + "a")))))
                                                                    _ = v30
                                                                    _ = v31
                                                                }
                                                                if ((!v4) || false) {
                                                                    switch fi14 {
                                                                    case -9223372036854775808:
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            if len(v2) > 0 {
                                                                                v2[0] = v9
                                                                            }
                                                                        case "hello":
                                                                            var v32 uint64 = 42
                                                                            _ = v32
                                                                        default:
                                                                            _ = len(v2)
                                                                        }
                                                                    case -505:
                                                                        if (!(1 <= -1)) {
                                                                            var v33 float64 = 1e10
                                                                            var v34 uint8 = 0
                                                                            _ = v33
                                                                            _ = v34
                                                                        } else {
                                                                            v35 := GenS1{F0: 0, F1: v29}
                                                                            var v36 uint16 = v1
                                                                            var v37 uint16 = v1
                                                                            _ = v35
                                                                            _ = v36
                                                                            _ = v37
                                                                        }
                                                                    case 32767:
                                                                        if (!v4) {
                                                                            switch v24 {
                                                                            case "foo bar":
                                                                                _ = len(v2)
                                                                            case "hello":
                                                                                if ((!(v24 > (("hello" + "hello") + v24))) || (v4 || (!v4))) {
                                                                                    fmt.Println(v9, fi13, v24)
                                                                                }
                                                                            default:
                                                                                v15 += -42
                                                                            }
                                                                            var v38 uint16 = (v1 | v1)
                                                                            fmt.Println(v1, v4)
                                                                            var v39 float64 = 0.0
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            _ = v38
                                                                            _ = v39
                                                                        }
                                                                        _ = len(v2)
                                                                        v1 -= 65535
                                                                    }
                                                                    _ = len(v2)
                                                                } else {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                            }
                                                            v2 = append(v2, v9)
                                                            var v40 uint8 = 128
                                                            _ = v40
                                                        }
                                                    default:
                                                        if v4 {
                                                            var v41 int16 = int16(v15)
                                                            _ = v41
                                                        } else {
                                                            v29 -= uint64(v9)
                                                        }
                                                        if (255 >= 146) {
                                                            var v42 int8 = (v15 * (v15 * (^(v15))))
                                                            if v4 {
                                                                var v43 uint = 1
                                                                _ = v43
                                                            } else {
                                                                v9 += uint32(v1)
                                                            }
                                                            switch ("" + "\n\t") {
                                                            case "":
                                                                var v44 int64 = 9223372036854775807
                                                                _ = v44
                                                            }
                                                            v18 = (-(fi13))
                                                            _ = v42
                                                        }
                                                    }
                                                    v45 := GenS1{F0: (-(v15)), F1: (v29 << ((255) & 63))}
                                                    _ = v45
                                                } else {
                                                    var v46 int8 = (v15 % ((1) | 1))
                                                    v3 -= (^(v3))
                                                    if false {
                                                        var v47 int16 = 42
                                                        var v48 int16 = (v47 ^ 417)
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                    _ = v46
                                                }
                                                var v49 string = "foo bar"
                                                _ = v29
                                                _ = v49
                                            }
                                            switch (((v24 + "") + "foo bar") + ("hello" + v24)) {
                                            case "foo bar":
                                                var v50 uint32 = v9
                                                _ = v50
                                            case "\n\t":
                                                var v51 uint64 = v3
                                                _ = v51
                                            default:
                                                if ((v1 >> ((2) & 63)) >= v1) {
                                                    var v52 int32 = -346
                                                    if (!((!(v4 && true)) || (42 >= 207))) {
                                                        v4 = (0.0 == (3.14159 + 1e10))
                                                    }
                                                    if ((v3 | v3) <= 770) {
                                                        if len(v2) > 0 {
                                                            v2[0] = v9
                                                        }
                                                    }
                                                    _ = v52
                                                } else {
                                                    v53 := GenS2{F0: 14, F1: v1}
                                                    if ((v15 + (v15 % ((1) | 1))) == v15) {
                                                        var v54 int16 = -615
                                                        v2 = append(v2, 141)
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v54
                                                    } else {
                                                        v55 := GenS1{F0: v15, F1: 42}
                                                        _ = v55
                                                    }
                                                    v18 -= (-(v18))
                                                    _ = v53
                                                }
                                            }
                                            _ = v25
                                        }
                                        v9 += (v9 | v9)
                                        _ = v15
                                        _ = v18
                                        _ = v24
                                    case -920:
                                        var v56 bool = v4
                                        _ = v56
                                    case -706:
                                        v3 -= v3
                                    default:
                                        v4 = (v3 >= (v3 - v3))
                                        var v57 uint8 = 128
                                        if (!v4) {
                                            switch fi14 {
                                            case -854:
                                                switch "" {
                                                case "":
                                                    v58 := GenS1{F0: -100, F1: v3}
                                                    _ = v58
                                                case "hello":
                                                    v9 -= v9
                                                }
                                            default:
                                                _ = len(v2)
                                            }
                                            v59 := GenS3{F0: fi13, F1: "foo bar"}
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v60 float64 = -2.25
                                            var v61 string = "hello"
                                            _ = v59
                                            _ = v60
                                            _ = v61
                                        } else {
                                            var v62 int8 = 42
                                            var v63 uint8 = v57
                                            if ((v63 / ((v57) | 1)) == v63) {
                                                v62 = v62
                                            }
                                            _ = v62
                                            _ = v63
                                        }
                                        _ = v57
                                    }
                                    if v4 {
                                        v1 = (v1 / (((^(v1))) | 1))
                                    } else {
                                        fmt.Println(v3, v9)
                                        v3 = v3
                                    }
                                }
                                var v64 uint64 = uint64(v1)
                                _ = fi14
                                _ = v64
                            }
                            v65 := GenS4{F0: (1 <= -580)}
                            _ = fi13
                            _ = v65
                        }
                    } else {
                        if v4 {
                            v2 = append(v2, v9)
                            var v66 string = "a"
                            _ = v66
                        }
                        fmt.Println(v1, v4, v1)
                    }
                }
                var v67 uint = 1
                _ = v67
            }
            _ = v9
        } else {
            for fi68 := 0; fi68 < 5; fi68++ {
                if false {
                    v2 = append(v2, 4294967295)
                    for fi69 := 0; fi69 < 4; fi69++ {
                        v70 := GenS5{F0: -551, F1: 4294967295, F2: v3}
                        _ = fi69
                        _ = v70
                    }
                    for fi71 := 0; fi71 < 8; fi71++ {
                        v72 := GenS3{F0: fi71, F1: ""}
                        switch "hello" {
                        case "":
                            var v73 int16 = -32768
                            _ = v73
                        }
                        _ = fi71
                        _ = v72
                    }
                } else {
                    for fi74 := 0; fi74 < 16; fi74++ {
                        var v75 string = "foo bar"
                        v75 = ((("\n\t" + "") + "") + v75)
                        _ = fi74
                        _ = v75
                    }
                    for fi76 := 0; fi76 < 2; fi76++ {
                        v77 := GenS3{F0: (fi76 >> ((1) & 63)), F1: ("foo bar" + "\n\t")}
                        switch fi68 {
                        case 32768:
                            v1 = (v1 - (v1 >> ((1) & 63)))
                        case 2147483647:
                            var v78 int64 = 0
                            _ = v78
                        default:
                            var v79 uint32 = 65535
                            _ = v79
                        }
                        _ = fi76
                        _ = v77
                    }
                }
                for fi80 := 0; fi80 < 3; fi80++ {
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi80
                }
                if v4 {
                    for fi81 := 0; fi81 < 19; fi81++ {
                        v82 := GenS1{F0: 1, F1: v3}
                        switch fi68 {
                        case 32767:
                            _ = len(v2)
                        case 0:
                            fmt.Println(v4, fi68)
                        case -1:
                            var v83 uint8 = 255
                            _ = v83
                        default:
                            if true {
                                v4 = v4
                                var v84 uint32 = 0
                                _ = v84
                            }
                        }
                        v2 = append(v2, 2147483647)
                        _ = fi81
                        _ = v82
                    }
                    switch "foo bar" {
                    case "hello":
                        var v85 uint = 4294967295
                        _ = v85
                    }
                    if true {
                        fmt.Println(v1, v4, v4)
                        var v86 int16 = 32767
                        _ = v86
                    }
                } else {
                    switch fi68 {
                    case 127:
                        _ = len(v2)
                    case -9223372036854775808:
                        var v87 uint = 42
                        _ = v87
                    case 358:
                        switch ("foo bar" + (("a" + ("hello" + "a")) + ("foo bar" + ""))) {
                        case "hello":
                            fmt.Println(v3, fi68, v1)
                        }
                    }
                    if (96 > 45) {
                        _ = len(v2)
                    }
                }
                _ = fi68
            }
            fmt.Println(v1, v4, v1)
        }
    }
    if false {
        v88 := GenS6{F0: v4, F1: v1}
        if len(v2) > 0 {
            v2[0] = 439
        }
        v3 -= v3
        _ = v88
    }
    fmt.Println(v3)
    if true {
        if ((2 <= 607) && (false || v4)) {
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
        }
    } else {
        for fi89 := 0; fi89 < 2; fi89++ {
            var v90 uint32 = 0
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi89
            _ = v90
        }
        if len(v2) > 0 {
            v2[0] = 65535
        }
    }
    fmt.Println(v1, v4)
    for fi91 := 0; fi91 < 13; fi91++ {
        if len(v2) > 0 {
            v2[0] = 0
        }
        if len(v2) > 0 {
            v2[0] = 599
        }
        if true {
            v4 = v4
            for fi92 := 0; fi92 < 6; fi92++ {
                switch ("a" + ("foo bar" + ((("\n\t" + ("foo bar" + ("hello" + ("" + "foo bar")))) + "a") + ("foo bar" + "foo bar")))) {
                case "hello":
                    v93 := GenS1{F0: 0, F1: v3}
                    _ = v93
                case "":
                    if v4 {
                        fmt.Println(fi91, v1, v4)
                        v3 -= 1
                        var v94 uint8 = 42
                        switch fi92 {
                        case -878:
                            if len(v2) > 0 {
                                v2[0] = 65535
                            }
                            switch "\n\t" {
                            case "\n\t":
                                var v95 float32 = 0.1
                                _ = v95
                            }
                            switch (fi91 >> ((42) & 63)) {
                            case -9223372036854775808:
                                var v96 bool = (v1 > (v1 & uint16(v94)))
                                _ = v96
                            default:
                                switch fi91 {
                                case 32768:
                                    v97 := GenS3{F0: fi91, F1: "a"}
                                    _ = v97
                                case 1:
                                    v4 = (-596 > 1000)
                                    v98 := GenS5{F0: 327, F1: 255, F2: v3}
                                    var v99 bool = ((100 > 127) || (false && true))
                                    _ = v98
                                    _ = v99
                                }
                                var v100 int8 = -27
                                var v101 int32 = 42
                                _ = v100
                                _ = v101
                            }
                        case -9223372036854775808:
                            v102 := GenS1{F0: 1, F1: 4294967295}
                            _ = v102
                        case 32767:
                            v2 = append(v2, 2)
                        default:
                            _ = len(v2)
                        }
                        if (1 >= -1) {
                            var v103 int8 = -88
                            _ = v103
                        } else {
                            if ((127 == (v94 - 0)) || v4) {
                                var v104 uint = 18446744073709551615
                                if v4 {
                                    switch fi92 {
                                    case -9223372036854775808:
                                        var v105 float32 = ((1e10 * (3.14159 * 1.5)) * 1e10)
                                        _ = v105
                                    }
                                    if v4 {
                                        fmt.Println(v3)
                                        var v106 float32 = (-2.25 - float32(3.14159))
                                        fmt.Println(v4, v1)
                                        _ = v106
                                    }
                                    var v107 int8 = 127
                                    if v4 {
                                        v108 := GenS6{F0: v4, F1: v1}
                                        if (v4 || (false || v4)) {
                                            if ((1 > (v94 | 0)) && ((65535 > v1) || true)) {
                                                switch -2147483648 {
                                                case 9223372036854775807:
                                                    if v4 {
                                                        v109 := GenS7{F0: 2147483648}
                                                        _ = v109
                                                    }
                                                default:
                                                    v104 -= v104
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = 65535
                                                }
                                            } else {
                                                var v110 uint = v104
                                                fmt.Println(v4)
                                                _ = v110
                                            }
                                            v107 = 25
                                            var v111 uint64 = v3
                                            v1 = v1
                                            v112 := GenS8{F0: v1, F1: ("a" + ("" + "hello"))}
                                            _ = v111
                                            _ = v112
                                        } else {
                                            fmt.Println(v4, v107, v104)
                                            var v113 float64 = (3.14159 + 0.1)
                                            var v114 bool = v4
                                            _ = v113
                                            _ = v114
                                        }
                                        v115 := GenS4{F0: (v1 < v1)}
                                        var v116 int8 = v107
                                        _ = v108
                                        _ = v115
                                        _ = v116
                                    }
                                    var v117 string = ("a" + (("a" + "\n\t") + "foo bar"))
                                    _ = v107
                                    _ = v117
                                }
                                v1 -= 32767
                                _ = v104
                            }
                            v1 += v1
                        }
                        v4 = v4
                        switch "foo bar" {
                        case "a":
                            v4 = (!false)
                        }
                        _ = v94
                    }
                default:
                    if (!false) {
                        if (true || true) {
                            if (255 <= 42) {
                                v1 = (v1 & (v1 ^ v1))
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = 2
                                }
                            }
                            var v118 int32 = 42
                            _ = v118
                        } else {
                            var v119 int64 = 42
                            if len(v2) > 0 {
                                v2[0] = 476
                            }
                            if len(v2) > 0 {
                                v2[0] = 65535
                            }
                            _ = v119
                        }
                        v1 = (v1 >> ((uint(v3)) & 63))
                        v4 = v4
                        _ = len(v2)
                        fmt.Println(v1, fi92, fi91)
                        if (v4 || v4) {
                            v120 := GenS5{F0: -1, F1: 674, F2: (^(v3))}
                            v4 = ((("\n\t" <= "") || (!v4)) && v4)
                            _ = v120
                        } else {
                            v121 := GenS6{F0: (128 == uint8(v1)), F1: 42}
                            v122 := GenS9{F0: 1, F1: 0, F2: 838}
                            _ = v121
                            _ = v122
                        }
                        v1 -= (v1 + v1)
                    }
                }
                v4 = v4
                _ = fi92
            }
            v2 = append(v2, 1)
        }
        _ = fi91
    }
    v3 -= 4294967295
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 -= (v1 & v1)
    var v123 int = -420
    if v4 {
        fmt.Println(v3, v3)
        switch v123 {
        case 127:
            for fi124 := 0; fi124 < 9; fi124++ {
                var v125 uint64 = v3
                v4 = ("foo bar" >= ("foo bar" + "a"))
                v1 = (v1 / (((v1 % ((v1) | 1))) | 1))
                _ = fi124
                _ = v125
            }
        }
        if (65535 == 65535) {
            fmt.Println(v123, v3)
            if (-2147483648 != v123) {
                var v126 float32 = 0.0
                if len(v2) > 0 {
                    v2[0] = 762
                }
                v3 = (v3 ^ (v3 ^ 1))
                _ = v126
            }
        } else {
            v127 := GenS10{F0: 13}
            v128 := GenS4{F0: true}
            if (127 <= (-(v123))) {
                var v129 int64 = 42
                for fi130 := 0; fi130 < 7; fi130++ {
                    if (v4 || true) {
                        var v131 float32 = -2.25
                        if ((fi130 | -9223372036854775808) == 32768) {
                            v123 = (fi130 + (v123 - 9223372036854775807))
                            v129 = (v129 ^ (v129 % ((v129) | 1)))
                            v123 += (^(v123))
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v123 -= 638
                        }
                        _ = v131
                    }
                    _ = len(v2)
                    _ = fi130
                }
                switch v123 {
                case 645:
                    for fi132 := 0; fi132 < 8; fi132++ {
                        v133 := GenS3{F0: (-(fi132)), F1: ("foo bar" + "a")}
                        v134 := GenS11{F0: -32768}
                        var v135 uint16 = (v1 ^ (^(v1)))
                        _ = fi132
                        _ = v133
                        _ = v134
                        _ = v135
                    }
                case 4294967295:
                    v1 = (v1 ^ 32767)
                }
                _ = v129
            } else {
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
                var v136 uint16 = v1
                if false {
                    v3 = v3
                } else {
                    v123 = v123
                    for fi137 := 0; fi137 < 4; fi137++ {
                        v2 = append(v2, 42)
                        v138 := GenS5{F0: -32768, F1: 255, F2: v3}
                        var v139 int = fi137
                        _ = fi137
                        _ = v138
                        _ = v139
                    }
                    v4 = true
                }
                _ = v136
            }
            _ = v127
            _ = v128
        }
    } else {
        var v140 string = (("\n\t" + "") + (("\n\t" + ("a" + ("\n\t" + ""))) + "foo bar"))
        if (2 < 2) {
            var v141 float64 = float64(3.14159)
            v142 := GenS4{F0: v4}
            var v143 string = ((((((("\n\t" + "") + "") + (v140 + (v140 + (v140 + (v140 + v140))))) + "a") + "") + ((v140 + (v140 + (("foo bar" + "hello") + (v140 + "a")))) + "hello")) + (((((v140 + ("\n\t" + "hello")) + v140) + v140) + v140) + (v140 + v140)))
            _ = v141
            _ = v142
            _ = v143
        } else {
            v1 += 32767
        }
        _ = v140
    }
    var v144 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v144:", v144)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v123
    _ = v144
    _ = v2
    _ = v3
    _ = v4
}
