package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 uint16; F2 int }
type GenS8 struct { F0 int32; F1 uint64; F2 string }
type GenS9 struct { F0 int16; F1 uint64 }
type GenS10 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int64; F1 uint64 }
type GenS13 struct { F0 uint16; F1 int8 }
type GenS14 struct { F0 uint16; F1 int8; F2 uint }
type GenS15 struct { F0 uint16; F1 uint64 }
type GenS16 struct { F0 uint32; F1 uint32 }
type GenS17 struct { F0 int32; F1 uint32 }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 uint16 }
func genFn0() uint64 {
    switch "" {
    case "":
        var v7 int32 = 1
        _ = v7
    case "a":
        var v8 uint8 = 204
        _ = v8
    case "foo bar":
        var v9 bool = false
        _ = v9
    default:
        var v10 uint64 = 18446744073709551615
        _ = v10
    }
    v11 := []int{-9223372036854775808}
    _ = v11
    return 145
}
func genFn1(p0 uint32, p1 uint16, p2 int) uint8 {
    fmt.Println(p2, p0)
    v12 := GenS1{F0: "foo bar", F1: (p1 | (^(p1)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return uint8(p1)
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    var v3 int32 = -2147483648
    var v4 int = -9223372036854775808
    v5 := []string{"\n\t", ""}
    var v6 int8 = 33
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS1{F0: ("\n\t" + "a"), F1: 0}
        _ = fi13
        _ = v14
    }
    v4 += 4294967295
    switch ("" + "\n\t") {
    case "hello":
        v4 = v4
    }
    var v15 uint32 = 1
    var v16 uint8 = 255
    var v17 int16 = -32768
    if (v15 != uint32(v16)) {
        if true {
            if true {
                if (!true) {
                    v18 := GenS2{F0: v15, F1: 255}
                    for fi19 := 0; fi19 < 10; fi19++ {
                        var v20 uint64 = 296
                        var v21 int8 = 102
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v17 -= (v17 / ((v17) | 1))
                    _ = v18
                }
                var v22 int32 = v3
                var v23 int16 = (^(v17))
                _ = v22
                _ = v23
            } else {
                for fi24 := 0; fi24 < 9; fi24++ {
                    switch int(v16) {
                    case 1:
                        v25 := GenS3{F0: 2, F1: (-1 == int64(v3))}
                        _ = v25
                    case 42:
                        var v26 bool = (true || true)
                        _ = v26
                    case 32768:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    delete(v2, "\n\t")
                    var v27 int32 = 2147483647
                    _ = fi24
                    _ = v27
                }
            }
            _ = len(v5)
            var v28 int = (v4 ^ v4)
            _ = v28
        } else {
            v5 = append(v5, "foo bar")
        }
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 string = ""
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    if (true && (42 == 716)) {
        for fi32 := 0; fi32 < 1; fi32++ {
            v1[(v15 | v15)] = true
            _ = fi32
        }
        v33 := GenS2{F0: 2147483647, F1: 42}
        _ = v33
    } else {
        v34 := v2[""]
        v35 := GenS2{F0: v15, F1: 32767}
        _ = v34
        _ = v35
    }
    v36 := GenS3{F0: 4294967295, F1: false}
    if true {
        for fi37 := 0; fi37 < 18; fi37++ {
            _ = len(v1)
            _ = fi37
        }
        fmt.Println(v6)
        v3 = v3
    } else {
        for fi38 := 0; fi38 < 11; fi38++ {
            v4 = (v4 | v4)
            var v39 int8 = v6
            switch "\n\t" {
            case "\n\t":
                for fi40 := 0; fi40 < 12; fi40++ {
                    v41 := GenS3{F0: 42, F1: (!(false || (1.5 <= 0.1)))}
                    delete(v2, "")
                    _ = fi40
                    _ = v41
                }
            case "":
                if (42 > genFn0()) {
                    v17 += 0
                    switch "foo bar" {
                    case "hello":
                        switch "foo bar" {
                        case "\n\t":
                            v6 = v39
                        case "":
                            switch 32767 {
                            case 650:
                                v39 -= v6
                            case 9223372036854775807:
                                for fi42 := 0; fi42 < 2; fi42++ {
                                    fmt.Println(v16, v17, v17)
                                    if ((!(!(!true))) || (float32(1e10) > float32(0.1))) {
                                        var v43 uint8 = v16
                                        v44 := GenS2{F0: v15, F1: 101}
                                        var v45 int64 = -9223372036854775808
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                    _ = fi42
                                }
                            default:
                                v46 := GenS4{F0: false, F1: (v6 / ((0) | 1))}
                                _ = v46
                            }
                        }
                    case "a":
                        var v47 uint32 = 0
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 4; fi48++ {
                        v4 -= (v4 + fi38)
                        v49 := GenS5{F0: 0}
                        v6 = 0
                        v15 -= 854
                        switch 4294967295 {
                        case 4294967295:
                            if true {
                                if true {
                                    var v50 float32 = -2.25
                                    _ = v50
                                }
                            } else {
                                if (-2.25 == float64(((0.1 / 1e10) / (3.14159 - (0.0 + 0.1))))) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v51 float64 = (-2.25 / 1.0)
                                    v6 = v6
                                    _ = v51
                                } else {
                                    v52 := GenS6{F0: 255}
                                    if ((^(v16)) > (v16 ^ (^(v16)))) {
                                        _ = len(v5)
                                        var v53 uint16 = 42
                                        _ = v53
                                    }
                                    var v54 uint32 = 458
                                    _ = v52
                                    _ = v54
                                }
                                switch (fi38 % ((-9223372036854775808) | 1)) {
                                case 127:
                                    v55 := GenS7{F0: (v16 / ((v16) | 1)), F1: 491, F2: (fi48 / ((-9223372036854775808) | 1))}
                                    _ = v55
                                default:
                                    var v56 int64 = -9223372036854775808
                                    _ = v56
                                }
                            }
                        case 32767:
                            if false {
                                var v57 int = v4
                                fmt.Println(v3)
                                _ = v57
                            }
                        case -9223372036854775808:
                            switch "\n\t" {
                            case "hello":
                                var v58 int8 = v39
                                v59 := GenS8{F0: v3, F1: genFn0(), F2: ""}
                                _ = v58
                                _ = v59
                            case "":
                                if (148 != v15) {
                                    switch "hello" {
                                    case "a":
                                        if true {
                                            var v60 string = ""
                                            _ = v60
                                        } else {
                                            if (v17 == v17) {
                                                if (((3.14159 / (3.14159 + (1e10 + (0.1 * -2.25)))) <= 0.0) || (!(true && (false || ((float64(float32((0.1 - ((-2.25 * -2.25) + 1.5)))) >= 3.14159) || true))))) {
                                                    v17 = (^(v17))
                                                    if ((0 >= 0) || true) {
                                                        var v61 float32 = float32(3.14159)
                                                        fmt.Println(v39, v3)
                                                        if len(v5) > 0 {
                                                            v5[0] = "hello"
                                                        }
                                                        _ = v61
                                                    }
                                                    if (!true) {
                                                        fmt.Println(v17, v3)
                                                        v3 += (-(v3))
                                                        if false {
                                                            if false {
                                                                v62 := GenS4{F0: false, F1: 100}
                                                                v63 := GenS9{F0: 42, F1: uint64(v16)}
                                                                _ = v62
                                                                _ = v63
                                                            }
                                                            v5 = append(v5, "")
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        }
                                                    } else {
                                                        var v64 int8 = v6
                                                        var v65 int = (v4 * (-(v4)))
                                                        var v66 float64 = 0.1
                                                        _ = v64
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                } else {
                                                    v67 := GenS3{F0: genFn0(), F1: false}
                                                    _ = v67
                                                }
                                                v68 := v2["hello"]
                                                v69 := GenS2{F0: v68, F1: uint16(v16)}
                                                _ = v68
                                                _ = v69
                                            }
                                            _ = len(v5)
                                            v15 = (v15 % ((v15) | 1))
                                        }
                                    case "foo bar":
                                        if false {
                                            var v70 uint = 517
                                            _ = v70
                                        } else {
                                            switch "foo bar" {
                                            case "hello":
                                                if false {
                                                    var v71 float32 = float32(0.0)
                                                    switch (v4 >> ((42) & 63)) {
                                                    case -522:
                                                        v72 := GenS10{F0: int64(v3), F1: (v15 / ((457) | 1)), F2: 2}
                                                        _ = v72
                                                    case -9223372036854775808:
                                                        var v73 uint = 65535
                                                        _ = v73
                                                    }
                                                    var v74 int8 = v39
                                                    v75 := GenS3{F0: 0, F1: false}
                                                    _ = v71
                                                    _ = v74
                                                    _ = v75
                                                }
                                            case "a":
                                                if false {
                                                    if (!(1e10 != (1e10 / 1.5))) {
                                                        v76 := GenS7{F0: genFn1(321, 0, (v4 + (v4 | fi48))), F1: 0, F2: (fi48 / (((fi38 ^ (-(fi48)))) | 1))}
                                                        _ = v76
                                                    }
                                                    if false {
                                                        var v77 int = v4
                                                        _ = v77
                                                    } else {
                                                        fmt.Println(v6, fi38)
                                                    }
                                                } else {
                                                    switch v4 {
                                                    case 123:
                                                        switch "\n\t" {
                                                        case "a":
                                                            if (4294967295 <= 920) {
                                                                fmt.Println(v17, v6)
                                                                var v78 bool = (false || true)
                                                                _ = v78
                                                            }
                                                        case "":
                                                            if true {
                                                                var v79 string = ("" + "foo bar")
                                                                _ = v79
                                                            }
                                                        }
                                                    case 127:
                                                        var v80 int = 42
                                                        _ = v80
                                                    case -2147483648:
                                                        v15 = v15
                                                        if (!true) {
                                                            fmt.Println(v3)
                                                            v15 -= v15
                                                            v81 := GenS11{F0: genFn1(v15, 82, 9223372036854775807)}
                                                            v15 = 65535
                                                            _ = v81
                                                        }
                                                    }
                                                    if true {
                                                        v82 := GenS12{F0: 42, F1: 0}
                                                        if (int64(v39) > -9223372036854775808) {
                                                            if true {
                                                                v39 += v39
                                                                var v83 int64 = -676
                                                                switch "hello" {
                                                                case "a":
                                                                    switch "a" {
                                                                    case "":
                                                                        if true {
                                                                            v17 += 1000
                                                                        } else {
                                                                            var v84 int64 = -494
                                                                            fmt.Println(v17)
                                                                            _ = v84
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    var v85 int = 2147483647
                                                                    _ = v85
                                                                }
                                                                var v86 int8 = 58
                                                                _ = v83
                                                                _ = v86
                                                            }
                                                            delete(v2, "foo bar")
                                                        } else {
                                                            if true {
                                                                fmt.Println(v39)
                                                            } else {
                                                                fmt.Println(v15, v39, fi48)
                                                            }
                                                            switch "foo bar" {
                                                            case "a":
                                                                fmt.Println(fi38, v3)
                                                            }
                                                        }
                                                        v6 = (v6 + 1)
                                                        _ = v82
                                                    } else {
                                                        var v87 int64 = 42
                                                        _ = v87
                                                    }
                                                    fmt.Println(v16, v39, v39)
                                                    v1[v15] = true
                                                }
                                            default:
                                                switch "" {
                                                case "hello":
                                                    fmt.Println(fi38, fi38, v16)
                                                    if (!(!(!((float64(0.0) - float64(-2.25)) != 0.0)))) {
                                                        v17 = v17
                                                        v17 += 1
                                                        v5 = append(v5, "hello")
                                                    }
                                                case "\n\t":
                                                    var v88 int = (v4 * 128)
                                                    _ = v88
                                                case "":
                                                    delete(v1, (^(v15)))
                                                default:
                                                    var v89 uint32 = (v15 - 493)
                                                    _ = v89
                                                }
                                            }
                                            var v90 uint8 = 1
                                            v90 = v90
                                            var v91 int16 = (-(v17))
                                            v4 -= 32767
                                            _ = v90
                                            _ = v91
                                        }
                                    case "\n\t":
                                        var v92 uint8 = v16
                                        _ = v92
                                    default:
                                        var v93 string = "a"
                                        _ = v93
                                    }
                                    if false {
                                        var v94 int16 = (v17 * (^(v17)))
                                        var v95 uint32 = v15
                                        v95 -= (v15 % ((v95) | 1))
                                        _ = v94
                                        _ = v95
                                    }
                                }
                            case "a":
                                var v96 uint32 = 4294967295
                                _ = v96
                            default:
                                var v97 uint8 = genFn1(v15, 760, fi48)
                                _ = v97
                            }
                        default:
                            switch "" {
                            case "hello":
                                v98 := GenS13{F0: 46, F1: int8(v17)}
                                _ = v98
                            case "a":
                                var v99 int64 = 9223372036854775807
                                _ = v99
                            case "\n\t":
                                v100 := GenS14{F0: 65535, F1: v6, F2: 18446744073709551615}
                                _ = v100
                            }
                        }
                        _ = fi48
                        _ = v49
                    }
                } else {
                    v1[(v15 ^ 182)] = ((((-1 == 0) && false) || (!((v3 << ((684) & 63)) == (^(v3))))) && false)
                }
            }
            _ = fi38
            _ = v39
        }
    }
    v6 -= v6
    if (true && true) {
        if false {
            v17 = 42
        }
    } else {
        fmt.Println(v3, v4, v4)
        if true {
            for fi101 := 0; fi101 < 6; fi101++ {
                var v102 uint16 = 255
                var v103 uint32 = (v15 - v15)
                _ = fi101
                _ = v102
                _ = v103
            }
            v16 -= 132
            v104 := GenS15{F0: 590, F1: 778}
            _ = v104
        } else {
            var v105 uint8 = v16
            v4 += v4
            switch v4 {
            case 1:
                var v106 int = -573
                _ = v106
            case -9223372036854775808:
                if false {
                    for fi107 := 0; fi107 < 5; fi107++ {
                        v105 -= genFn1((v15 % (((v15 * (v15 - (v15 | v15)))) | 1)), 255, (v4 << ((65535) & 63)))
                        _ = len(v5)
                        if true {
                            v17 = v17
                            v108 := GenS1{F0: "\n\t", F1: 0}
                            fmt.Println(v105, v4, fi107)
                            _ = v108
                        }
                        _ = fi107
                    }
                    for fi109 := 0; fi109 < 15; fi109++ {
                        v4 -= fi109
                        _ = fi109
                    }
                    switch 9223372036854775807 {
                    case 153:
                        v110 := GenS16{F0: (^(v15)), F1: 146}
                        _ = v110
                    case 992:
                        v111 := GenS7{F0: (v105 >> ((255) & 63)), F1: 179, F2: (v4 + v4)}
                        _ = v111
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
            case 2147483647:
                if ("foo bar" <= ("" + "a")) {
                    for fi112 := 0; fi112 < 7; fi112++ {
                        if true {
                            if (v16 < (v105 & (v105 << ((2) & 63)))) {
                                if false {
                                    var v113 string = ("\n\t" + "a")
                                    v16 = (v105 ^ (^(v105)))
                                    _ = v113
                                } else {
                                    switch "hello" {
                                    case "\n\t":
                                        var v114 uint32 = (^(v15))
                                        _ = v114
                                    }
                                    v3 = v3
                                    if (42 > 18446744073709551615) {
                                        fmt.Println(v17)
                                        if false {
                                            v5 = append(v5, (("hello" + "") + ("" + "")))
                                            var v115 uint8 = (v16 + v105)
                                            if len(v5) > 0 {
                                                v5[0] = (("\n\t" + (("\n\t" + ("hello" + ("hello" + "\n\t"))) + "a")) + ("\n\t" + "\n\t"))
                                            }
                                            _ = v115
                                        } else {
                                            fmt.Println(v6, v105)
                                            var v116 uint8 = 255
                                            if (!true) {
                                                var v117 int16 = v17
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                var v118 uint64 = 127
                                                _ = v117
                                                _ = v118
                                            } else {
                                                v105 += (^(v116))
                                            }
                                            _ = v116
                                        }
                                        _ = len(v1)
                                    } else {
                                        if true {
                                            v119 := GenS17{F0: (^(v3)), F1: (^(v15))}
                                            _ = v119
                                        } else {
                                            v15 -= (v15 & (v15 & (^(v15))))
                                            if (((("" + ("foo bar" + ("foo bar" + ("a" + ("\n\t" + "\n\t"))))) >= ("\n\t" + "")) && (-1 >= 195)) || false) {
                                                v6 = (v6 >> ((255) & 63))
                                                v3 += (-(v3))
                                                var v120 int16 = (v17 << ((65535) & 63))
                                                _ = v120
                                            }
                                        }
                                    }
                                }
                                if false {
                                    v121 := GenS1{F0: "hello", F1: uint16(v15)}
                                    var v122 int64 = int64(v6)
                                    v17 = (v17 + v17)
                                    _ = v121
                                    _ = v122
                                }
                            }
                            var v123 float64 = 0.0
                            _ = v123
                        } else {
                            if true {
                                var v124 int64 = 9223372036854775807
                                _ = v124
                            } else {
                                v6 = (-(v6))
                            }
                        }
                        var v125 uint16 = 2
                        v126 := GenS18{F0: 715}
                        _ = fi112
                        _ = v125
                        _ = v126
                    }
                } else {
                    if ((true || true) || (v6 != (v6 ^ v6))) {
                        v16 += (v16 ^ (v105 - (^(v16))))
                    }
                    var v127 int8 = v6
                    var v128 float32 = 0.0
                    _ = v127
                    _ = v128
                }
            }
            _ = v105
        }
    }
    if (42 >= (v3 ^ v3)) {
        v16 -= (v16 - (v16 - 127))
    }
    if true {
        switch ((("\n\t" + "foo bar") + "\n\t") + (("" + ("foo bar" + "")) + "hello")) {
        case "":
            switch 32767 {
            case 283:
                v129 := GenS11{F0: (v16 ^ 1)}
                _ = v129
            case 4294967295:
                fmt.Println(v4)
            }
        }
    }
    switch v4 {
    case 128:
        if (!true) {
            switch v4 {
            case 32768:
                for fi130 := 0; fi130 < 9; fi130++ {
                    if ((false && (false || false)) && true) {
                        delete(v2, "foo bar")
                        v131 := GenS19{F0: 42}
                        _ = v131
                    } else {
                        var v132 int = 416
                        _ = v132
                    }
                    v3 = (v3 - (v3 * (v3 | v3)))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi130
                }
            case -2147483648:
                v133 := GenS5{F0: (v3 ^ (-(v3)))}
                _ = v133
            case 609:
                fmt.Println(v3, v17, v17)
            default:
                switch "hello" {
                case "\n\t":
                    var v134 uint16 = 255
                    _ = v134
                case "foo bar":
                    for fi135 := 0; fi135 < 5; fi135++ {
                        if (false && true) {
                            v136 := GenS13{F0: 42, F1: -128}
                            _ = v136
                        }
                        _ = fi135
                    }
                case "a":
                    v4 += (v4 ^ v4)
                }
            }
            var v137 float32 = 3.14159
            var v138 int = (v4 / (((^(v4))) | 1))
            _ = v137
            _ = v138
        } else {
            v139 := GenS6{F0: v16}
            _ = v139
        }
    case 32768:
        v140 := GenS3{F0: 0, F1: true}
        _ = v140
    case 32767:
        for fi141 := 0; fi141 < 1; fi141++ {
            fmt.Println(v3, v16, v3)
            var v142 string = ""
            _ = fi141
            _ = v142
        }
    default:
        fmt.Println(v4, v4)
    }
    for fi143 := 0; fi143 < 4; fi143++ {
        v16 -= (^(v16))
        v16 -= (v16 + v16)
        _ = fi143
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
