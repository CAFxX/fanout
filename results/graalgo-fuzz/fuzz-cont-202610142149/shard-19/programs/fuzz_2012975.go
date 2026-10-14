package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint64; F1 uint; F2 int16 }
type GenS7 struct { F0 bool; F1 uint64; F2 int64 }
type GenS8 struct { F0 string; F1 uint; F2 int16 }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 int8; F1 int64 }
type GenS11 struct { F0 string; F1 uint16 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 string; F1 uint32 }
type GenS14 struct { F0 int16; F1 int8; F2 uint64 }
type GenS15 struct { F0 uint8; F1 string }
type GenS16 struct { F0 uint8; F1 int16; F2 string }
type GenS17 struct { F0 int8 }
func genFn0(p0 uint16, p1 uint8) uint64 {
    v6 := GenS1{F0: "foo bar", F1: (("a" + "\n\t") + "a")}
    var v7 string = ("\n\t" + ("" + "foo bar"))
    if (!((1.5 != -2.25) || (42 <= 562))) {
        var v8 float64 = 1e10
        _ = v8
    }
    fmt.Println(p1)
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int8 = 100
        var v11 int = fi9
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return uint64(p0)
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]string)
    v3 := []uint32{4294967295, 2147483647}
    var v4 int16 = 0
    v5 := []string{"", "hello"}
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(v1)
        for fi13 := 0; fi13 < 3; fi13++ {
            v1 += 0
            fmt.Println(fi12, fi12, v4)
            _ = fi13
        }
        _ = fi12
    }
    v14 := GenS2{F0: -1, F1: (v1 % ((111) | 1))}
    v4 += v4
    v15 := GenS3{F0: 650}
    v1 = v1
    switch "" {
    case "hello":
        var v16 int8 = 1
        _ = v16
    }
    fmt.Println(v1, v1)
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 int64 = -887
        _ = fi17
        _ = v18
    }
    v4 -= 129
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            v21 := GenS4{F0: -327}
            v1 -= 59
            var v22 string = ("" + "foo bar")
            if false {
                if false {
                    v4 = v4
                    fmt.Println(fi19, v1, v22)
                    v22 = v22
                }
                var v23 uint64 = uint64(v1)
                _ = v23
            } else {
                var v24 int = -640
                switch (v22 + "foo bar") {
                case "\n\t":
                    v25 := GenS1{F0: (v22 + (v22 + "")), F1: ""}
                    _ = v25
                case "a":
                    var v26 string = (v22 + "foo bar")
                    v27 := GenS2{F0: -1, F1: v1}
                    if false {
                        fmt.Println(v24, v26)
                        var v28 int = fi20
                        _ = v28
                    }
                    v29 := GenS4{F0: 2147483647}
                    if (!(!(0.1 < (0.0 / (3.14159 - (3.14159 + (1e10 - (float64((float32(1.5) / (0.1 - (1.5 - float32(float64(float32(1.5))))))) / 1e10)))))))) {
                        if false {
                            if true {
                                var v30 uint16 = 2
                                v1 = (^(v1))
                                _ = v30
                            }
                            var v31 float32 = 0.1
                            v22 = (v22 + "\n\t")
                            _ = v31
                        }
                        if ((!(4294967295 == 0)) || (!(!true))) {
                            var v32 uint = 289
                            var v33 uint64 = 2
                            var v34 int16 = -804
                            var v35 float32 = (0.0 + 1e10)
                            _ = v32
                            _ = v33
                            _ = v34
                            _ = v35
                        } else {
                            switch (-(fi19)) {
                            case 127:
                                v36 := GenS4{F0: 558}
                                _ = v36
                            default:
                                var v37 uint64 = 2
                                _ = v37
                            }
                            if false {
                                v38 := GenS5{F0: 566, F1: 2, F2: v4}
                                var v39 uint64 = 0
                                switch ("a" + "foo bar") {
                                case "a":
                                    v22 = (v26 + "hello")
                                }
                                _ = v38
                                _ = v39
                            } else {
                                v24 = fi19
                                if (!false) {
                                    fmt.Println(v26)
                                    switch (v22 + v22) {
                                    case "":
                                        if true {
                                            if false {
                                                if true {
                                                    switch 2147483647 {
                                                    case 4294967295:
                                                        v40 := GenS6{F0: 42, F1: 42, F2: v4}
                                                        _ = v40
                                                    case 32768:
                                                        _ = len(v2)
                                                    case 0:
                                                        if false {
                                                            v41 := GenS5{F0: 0, F1: 0, F2: v4}
                                                            switch v26 {
                                                            case "foo bar":
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            default:
                                                                var v42 string = "\n\t"
                                                                _ = v42
                                                            }
                                                            v43 := GenS2{F0: 100, F1: 2}
                                                            _ = v41
                                                            _ = v43
                                                        }
                                                    }
                                                    v4 += v4
                                                    var v44 float32 = (0.0 * 1.5)
                                                    _ = v44
                                                } else {
                                                    var v45 int32 = -2147483648
                                                    var v46 float32 = float32(float64((1.5 / ((0.0 + 0.1) + ((((3.14159 * 0.0) / 1.0) - 1.5) - float32(-2.25))))))
                                                    fmt.Println(v1, v26)
                                                    _ = v45
                                                    _ = v46
                                                }
                                                var v47 int64 = -1
                                                fmt.Println(v24, v1, v4)
                                                _ = v47
                                            }
                                            switch (fi19 / ((fi19) | 1)) {
                                            case 1:
                                                if ((0.1 / -2.25) <= ((3.14159 - 3.14159) - 1.5)) {
                                                    v48 := GenS7{F0: true, F1: genFn0(0, v1), F2: -550}
                                                    v26 = v26
                                                    _ = v48
                                                }
                                            case 32767:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            switch v22 {
                                            case "foo bar":
                                                v4 = v4
                                            default:
                                                v4 = v4
                                            }
                                            if false {
                                                var v49 uint32 = 2147483648
                                                if (2 > 4294967295) {
                                                    delete(v2, (fi19 % ((fi19) | 1)))
                                                } else {
                                                    switch (("foo bar" + ("hello" + (("hello" + v26) + v26))) + v22) {
                                                    case "foo bar":
                                                        v50 := GenS8{F0: "hello", F1: 0, F2: (v4 ^ (v4 + -1))}
                                                        _ = v50
                                                    case "":
                                                        v24 = (fi20 ^ (fi19 ^ 32768))
                                                    default:
                                                        if true {
                                                            fmt.Println(v24, fi20)
                                                            v49 = v49
                                                            v51 := GenS1{F0: (v22 + v26), F1: (("a" + ((v26 + v26) + "a")) + v22)}
                                                            var v52 float32 = 1e10
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                    }
                                                    _ = len(v2)
                                                    v53 := GenS8{F0: ("hello" + (v22 + v26)), F1: 0, F2: (v4 & v4)}
                                                    switch fi20 {
                                                    case -9223372036854775808:
                                                        v49 = (v49 % ((v49) | 1))
                                                    case 2147483647:
                                                        if false {
                                                            v24 = -1
                                                            if false {
                                                                v1 = v1
                                                                v26 = v22
                                                                v54 := GenS7{F0: (!(fi20 > (fi20 + fi19))), F1: 18446744073709551615, F2: -677}
                                                                v55 := GenS9{F0: -1, F1: 1}
                                                                _ = v54
                                                                _ = v55
                                                            } else {
                                                                var v56 float32 = 0.0
                                                                _ = v56
                                                            }
                                                        } else {
                                                            delete(v2, 128)
                                                        }
                                                    default:
                                                        v57 := GenS10{F0: 42, F1: 9223372036854775807}
                                                        _ = v57
                                                    }
                                                    _ = v53
                                                }
                                                var v58 string = "foo bar"
                                                _ = v49
                                                _ = v58
                                            } else {
                                                var v59 int8 = 1
                                                var v60 uint16 = 324
                                                _ = v59
                                                _ = v60
                                            }
                                        } else {
                                            delete(v2, 32767)
                                            v61 := v2[(v24 % ((v24) | 1))]
                                            var v62 int32 = -1
                                            _ = v61
                                            _ = v62
                                        }
                                    case "\n\t":
                                        var v63 uint32 = 2147483647
                                        _ = v63
                                    }
                                }
                                var v64 int = (fi19 ^ v24)
                                _ = v64
                            }
                            v4 -= 1
                        }
                    } else {
                        v22 = "a"
                        v4 += v4
                    }
                    _ = v26
                    _ = v27
                    _ = v29
                }
                _ = v24
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        if (!(42 == 42)) {
            if (372 != -877) {
                fmt.Println(v4, v4)
                switch 42 {
                case 2147483647:
                    for fi65 := 0; fi65 < 17; fi65++ {
                        v66 := GenS3{F0: 42}
                        if (!(v1 == (v1 + (v1 - v1)))) {
                            v4 = (v4 % ((v4) | 1))
                            var v67 int8 = -106
                            if false {
                                var v68 int = (-(fi19))
                                _ = v68
                            }
                            v1 += (^(v1))
                            v4 += 449
                            _ = v67
                        }
                        _ = fi65
                        _ = v66
                    }
                default:
                    var v69 uint32 = 2147483648
                    _ = v69
                }
            } else {
                v1 = (v1 + 255)
            }
            var v70 uint = 682
            switch "" {
            case "":
                if false {
                    var v71 uint8 = (v1 & v1)
                    v71 = v71
                    _ = v71
                } else {
                    v72 := GenS11{F0: ("a" + "hello"), F1: 0}
                    fmt.Println(v70, fi19)
                    _ = v72
                }
            case "foo bar":
                if false {
                    var v73 float32 = float32(0.1)
                    _ = v73
                } else {
                    if false {
                        fmt.Println(fi19, fi19, v1)
                        if (!("" >= ("" + ("hello" + "hello")))) {
                            var v74 uint8 = v1
                            fmt.Println(v70)
                            fmt.Println(fi19)
                            _ = v74
                        } else {
                            v4 += v4
                            var v75 string = ""
                            for fi76 := 0; fi76 < 6; fi76++ {
                                var v77 uint16 = uint16(v1)
                                _ = len(v2)
                                var v78 float64 = 1e10
                                _ = fi76
                                _ = v77
                                _ = v78
                            }
                            _ = v75
                        }
                        fmt.Println(v1)
                    }
                    for fi79 := 0; fi79 < 18; fi79++ {
                        switch (("" + ("" + "\n\t")) + "hello") {
                        case "a":
                            v1 = uint8(v70)
                        case "foo bar":
                            switch 32767 {
                            case -1:
                                var v80 uint64 = genFn0(65535, (v1 - v1))
                                _ = v80
                            case 127:
                                v81 := GenS8{F0: ("\n\t" + ""), F1: 4294967295, F2: (^(v4))}
                                _ = v81
                            default:
                                v70 = v70
                            }
                        }
                        var v82 uint8 = v1
                        var v83 float64 = (1.5 * 0.1)
                        _ = fi79
                        _ = v82
                        _ = v83
                    }
                    for fi84 := 0; fi84 < 7; fi84++ {
                        v70 = 65535
                        _ = fi84
                    }
                }
            }
            _ = v70
        }
        _ = fi19
    }
    if (0.0 > 1e10) {
        if (-9223372036854775808 > 489) {
            var v85 uint = 490
            v86 := GenS12{F0: v85}
            if true {
                v87 := GenS13{F0: ("a" + ("a" + "")), F1: 2}
                if false {
                    delete(v2, 128)
                    for fi88 := 0; fi88 < 16; fi88++ {
                        delete(v2, (fi88 - 9223372036854775807))
                        v4 = int16(v1)
                        _ = fi88
                    }
                } else {
                    fmt.Println(v85, v1, v4)
                }
                fmt.Println(v1)
                _ = v87
            } else {
                if (false || (32767 <= 0)) {
                    for fi89 := 0; fi89 < 13; fi89++ {
                        if false {
                            switch 32768 {
                            case 32768:
                                v85 = v85
                            }
                            for fi90 := 0; fi90 < 9; fi90++ {
                                var v91 uint16 = 0
                                _ = fi90
                                _ = v91
                            }
                            v1 -= v1
                        }
                        v85 -= v85
                        for fi92 := 0; fi92 < 1; fi92++ {
                            if (uint32(v85) != 1) {
                                var v93 float64 = float64(0.0)
                                v94 := GenS5{F0: uint64(v85), F1: 451, F2: 42}
                                _ = len(v3)
                                _ = v93
                                _ = v94
                            } else {
                                _ = len(v2)
                                v4 = (v4 ^ v4)
                                v1 = (v1 ^ (v1 % ((v1) | 1)))
                            }
                            fmt.Println(v4)
                            fmt.Println(fi92, v1)
                            _ = fi92
                        }
                        _ = fi89
                    }
                    var v95 float32 = (3.14159 * ((float32(0.1) + (1.5 / 0.1)) - -2.25))
                    if false {
                        for fi96 := 0; fi96 < 1; fi96++ {
                            var v97 float32 = v95
                            var v98 uint = 58
                            fmt.Println(fi96, fi96)
                            _ = fi96
                            _ = v97
                            _ = v98
                        }
                        var v99 int32 = -2147483648
                        v100 := v2[32767]
                        _ = v100
                        _ = v99
                    } else {
                        v101 := GenS1{F0: "a", F1: "a"}
                        fmt.Println(v95, v4, v95)
                        var v102 uint8 = v1
                        _ = v101
                        _ = v102
                    }
                    _ = v95
                } else {
                    v2[int(v4)] = ""
                    _ = len(v2)
                    v103 := v2[42]
                    _ = v103
                }
                _ = len(v3)
            }
            _ = v85
            _ = v86
        } else {
            for fi104 := 0; fi104 < 4; fi104++ {
                switch ("" + "") {
                case "hello":
                    fmt.Println(v1, fi104)
                case "":
                    for fi105 := 0; fi105 < 8; fi105++ {
                        var v106 uint64 = 18446744073709551615
                        v107 := GenS14{F0: v4, F1: -128, F2: (^(v106))}
                        var v108 int8 = 1
                        v2[9223372036854775807] = "hello"
                        switch (fi104 % ((fi105) | 1)) {
                        case 4294967295:
                            v1 = (v1 << ((65535) & 63))
                        case 2147483647:
                            v108 += 100
                        case -2147483648:
                            v2[fi105] = "hello"
                        default:
                            v1 -= v1
                        }
                        _ = fi105
                        _ = v106
                        _ = v107
                        _ = v108
                    }
                default:
                    var v109 string = (("a" + (("a" + "foo bar") + "a")) + "\n\t")
                    _ = v109
                }
                v1 = v1
                _ = fi104
            }
        }
        _ = len(v3)
        var v110 int = 0
        _ = v110
    } else {
        for fi111 := 0; fi111 < 3; fi111++ {
            var v112 uint8 = v1
            for fi113 := 0; fi113 < 14; fi113++ {
                if (1 == fi111) {
                    if true {
                        var v114 int = (-(fi111))
                        _ = len(v2)
                        _ = v114
                    } else {
                        var v115 int64 = 298
                        v112 = (v112 ^ v112)
                        v4 += (v4 >> ((4294967295) & 63))
                        _ = v115
                    }
                    v4 += (v4 | 0)
                    var v116 int32 = 0
                    _ = v116
                } else {
                    fmt.Println(fi111, fi111)
                    v112 = (v1 + v112)
                    fmt.Println(v4, fi113)
                }
                if true {
                    v112 += (v1 >> ((845) & 63))
                    fmt.Println(fi111, v112)
                }
                _ = fi113
            }
            for fi117 := 0; fi117 < 8; fi117++ {
                v1 -= v112
                _ = fi117
            }
            _ = fi111
            _ = v112
        }
        fmt.Println(v4)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v118 int = 42
    switch 32768 {
    case 0:
        switch "" {
        case "\n\t":
            v119 := GenS3{F0: 4294967295}
            _ = v119
        case "":
            if (2 > 575) {
                var v120 int16 = v4
                switch 127 {
                case 9223372036854775807:
                    v4 += (v4 - 1000)
                case -336:
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v118, v1)
                    case "":
                        v121 := GenS15{F0: (^(v1)), F1: ("hello" + "a")}
                        _ = v121
                    case "foo bar":
                        for fi122 := 0; fi122 < 8; fi122++ {
                            if false {
                                v123 := GenS16{F0: 127, F1: 1000, F2: "hello"}
                                _ = v123
                            }
                            switch 4294967295 {
                            case 0:
                                _ = len(v2)
                            }
                            var v124 bool = true
                            _ = fi122
                            _ = v124
                        }
                    }
                }
                _ = v120
            } else {
                fmt.Println(v1, v118, v4)
                if true {
                    if (!false) {
                        var v125 float32 = ((1e10 * (-2.25 - 3.14159)) / (1e10 / 3.14159))
                        if true {
                            v126 := GenS17{F0: 42}
                            switch v118 {
                            case -9223372036854775808:
                                v127 := GenS15{F0: (^(v1)), F1: "a"}
                                _ = v127
                            case 32767:
                                v125 += v125
                            case 2147483647:
                                switch v118 {
                                case 0:
                                    fmt.Println(v4, v4)
                                }
                            default:
                                v128 := GenS15{F0: v1, F1: ""}
                                _ = v128
                            }
                            _ = len(v2)
                            _ = v126
                        } else {
                            if (("a" + ((("a" + ("" + "foo bar")) + "a") + "hello")) > "a") {
                                for fi129 := 0; fi129 < 3; fi129++ {
                                    var v130 uint64 = 595
                                    _ = fi129
                                    _ = v130
                                }
                            } else {
                                var v131 float64 = ((0.1 - 0.1) + ((((1e10 / 1.0) / (-2.25 - 3.14159)) + 3.14159) + 1.5))
                                _ = v131
                            }
                            v1 = 128
                            v132 := v2[(v118 / ((v118) | 1))]
                            _ = v132
                        }
                        _ = v125
                    }
                } else {
                    v1 = v1
                }
                delete(v2, -2147483648)
            }
        }
    case 128:
        if len(v3) > 0 {
            v3[0] = 625
        }
    default:
        v1 = v1
    }
    v118 += (v118 ^ v118)
    fmt.Println("v1:", v1)
    fmt.Println("v118:", v118)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v118
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
