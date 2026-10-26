package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint; F1 int8; F2 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32; F1 int16 }
type GenS10 struct { F0 bool; F1 int16 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 string; F1 uint64 }
type GenS14 struct { F0 int32; F1 uint16 }
func genFn0(p0 int32) uint {
    fmt.Println(p0)
    var v7 int16 = 141
    fmt.Println(v7, p0)
    v8 := GenS1{F0: 671, F1: 255, F2: 0}
    v7 += (v7 - 32767)
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    v3 := []int{127, 83, 32768}
    v4 := make(map[int]string)
    v5 := []int{740}
    v6 := []int{}
    var v9 uint8 = 0
    if false {
        switch 32767 {
        case -487:
            v9 += v9
        default:
            var v10 uint32 = 2147483648
            _ = v10
        }
        var v11 int8 = -114
        v11 = (v11 & v11)
        _ = v11
    }
    delete(v1, "foo bar")
    if false {
        var v12 float64 = 1e10
        _ = v12
    }
    v9 += v9
    var v13 float32 = (0.1 * 1.5)
    v14 := GenS1{F0: 1, F1: v9, F2: 97}
    if true {
        for fi15 := 0; fi15 < 12; fi15++ {
            if (2147483648 == 407) {
                switch (fi15 % (((fi15 ^ (fi15 - int(v9)))) | 1)) {
                case 0:
                    var v16 uint64 = 309
                    _ = v16
                case 238:
                    for fi17 := 0; fi17 < 16; fi17++ {
                        if true {
                            var v18 uint8 = (v9 * v9)
                            if (true && ((v13 < ((1e10 / (0.0 / v13)) * (float32((((3.14159 + 0.0) / 1.0) / -2.25)) + (-2.25 - (float32(float64(float32(0.1))) - (1e10 - v13)))))) || (!true))) {
                                var v19 int32 = 2147483647
                                if false {
                                    v20 := GenS1{F0: fi15, F1: v18, F2: 42}
                                    switch (fi15 ^ (fi15 & (fi17 * fi17))) {
                                    case -2147483648:
                                        v9 = 128
                                    case 799:
                                        var v21 uint16 = 32767
                                        _ = v21
                                    case 32768:
                                        fmt.Println(v19, v13, fi15)
                                    default:
                                        fmt.Println(fi17)
                                    }
                                    _ = v20
                                } else {
                                    if false {
                                        v22 := GenS2{F0: 67, F1: 42, F2: 65535}
                                        v9 = (v18 - v18)
                                        _ = v22
                                    } else {
                                        delete(v2, 801)
                                        v2[-2147483648] = ("a" + "a")
                                    }
                                }
                                _ = v19
                            } else {
                                var v23 uint32 = 42
                                v23 += (v23 % ((v23) | 1))
                                switch "\n\t" {
                                case "hello":
                                    if (!(-2147483648 > 664)) {
                                        v9 += v9
                                        v24 := GenS2{F0: 65535, F1: 42, F2: 803}
                                        _ = v24
                                    } else {
                                        v25 := GenS3{F0: 4294967295, F1: uint64(v9), F2: 649}
                                        _ = v25
                                    }
                                }
                                _ = v23
                            }
                            _ = v18
                        } else {
                            var v26 float64 = -2.25
                            var v27 int = int(v9)
                            var v28 bool = (v26 == float64((-2.25 - v13)))
                            switch (v27 << ((607) & 63)) {
                            case 228:
                                v29 := GenS1{F0: (fi17 - fi15), F1: 128, F2: -7}
                                _ = v29
                            case 127:
                                var v30 uint = 212
                                _ = v30
                            default:
                                var v31 bool = v28
                                switch "" {
                                case "hello":
                                    v32 := GenS2{F0: 580, F1: 127, F2: 566}
                                    _ = v32
                                case "\n\t":
                                    var v33 float32 = float32(v26)
                                    v34 := GenS3{F0: 803, F1: 0, F2: 0}
                                    _ = v33
                                    _ = v34
                                default:
                                    var v35 uint = 42
                                    fmt.Println(v9, v28)
                                    v31 = v31
                                    if true {
                                        _ = len(v5)
                                        v36 := GenS3{F0: v35, F1: 1, F2: 71}
                                        _ = v36
                                    }
                                    var v37 int16 = 32767
                                    var v38 uint16 = 32767
                                    _ = v35
                                    _ = v37
                                    _ = v38
                                }
                                _ = v31
                            }
                            v28 = (!v28)
                            fmt.Println(fi15, v28, v9)
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        if ((v9 / ((v9) | 1)) < v9) {
                            if (42 < 42) {
                                switch 361 {
                                case 0:
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v39 string = ""
                                        _ = v39
                                    case "hello":
                                        switch fi15 {
                                        case 4294967295:
                                            v9 = v9
                                        case 127:
                                            v13 = v13
                                        case -2147483648:
                                            switch "hello" {
                                            case "hello":
                                                v9 = 255
                                                v9 = (v9 << ((65535) & 63))
                                            case "\n\t":
                                                switch fi15 {
                                                case -1:
                                                    v40 := GenS4{F0: 80}
                                                    _ = v40
                                                case 182:
                                                    if false {
                                                        if (false && (("" >= ("a" + "foo bar")) && (fi15 > 2147483647))) {
                                                            fmt.Println(fi15, v13)
                                                        } else {
                                                            var v41 uint = 255
                                                            _ = v41
                                                        }
                                                        var v42 uint = genFn0(-329)
                                                        _ = v42
                                                    }
                                                case -2147483648:
                                                    if (false || (v9 != v9)) {
                                                        fmt.Println(v13, fi15, fi15)
                                                        if (4294967295 < 4294967295) {
                                                            var v43 int8 = 100
                                                            _ = v43
                                                        }
                                                        v13 += v13
                                                        v44 := GenS5{F0: 0}
                                                        _ = v44
                                                    }
                                                }
                                            }
                                        }
                                    case "\n\t":
                                        v13 += v13
                                    }
                                case 355:
                                    v13 -= (0.1 / v13)
                                default:
                                    v13 = 0.1
                                }
                                v45 := GenS2{F0: 413, F1: 0, F2: 1}
                                if (2 < 65535) {
                                    switch (fi17 * fi15) {
                                    case 32768:
                                        var v46 uint64 = 42
                                        _ = v46
                                    default:
                                        var v47 bool = true
                                        _ = v47
                                    }
                                    v48 := v2[(fi15 << ((4294967295) & 63))]
                                    _ = v48
                                } else {
                                    var v49 int64 = 9223372036854775807
                                    if (42 <= 1) {
                                        var v50 bool = ("hello" < "")
                                        v13 = 0.0
                                        v9 = 208
                                        _ = v50
                                    } else {
                                        v49 += 738
                                    }
                                    _ = v49
                                }
                                _ = v45
                            } else {
                                v51 := GenS6{F0: -2147483648, F1: false}
                                v52 := GenS5{F0: 205}
                                var v53 uint32 = 2147483648
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                        } else {
                            var v54 int16 = -516
                            v9 = 37
                            if true {
                                _ = len(v1)
                            }
                            _ = v54
                        }
                        _ = fi17
                    }
                default:
                    var v55 uint8 = 19
                    _ = v55
                }
                var v56 int = fi15
                v56 -= (v56 << ((255) & 63))
                _ = v56
            }
            switch "\n\t" {
            case "hello":
                v9 += v9
            case "foo bar":
                var v57 uint16 = 65535
                _ = v57
            case "":
                var v58 bool = ((!true) && true)
                _ = v58
            }
            var v59 bool = (false && (1 != genFn0(42)))
            _ = fi15
            _ = v59
        }
        var v60 int8 = -1
        if false {
            fmt.Println(v13, v9, v60)
        } else {
            var v61 uint = genFn0(-2147483648)
            _ = v61
        }
        _ = v60
    } else {
        var v62 int16 = 32767
        for fi63 := 0; fi63 < 8; fi63++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (2 == genFn0(-1)) {
                v64 := GenS6{F0: 1, F1: (!((0.1 != -2.25) || true))}
                _ = v64
            }
            _ = fi63
        }
        _ = v62
    }
    switch 553 {
    case 1:
        for fi65 := 0; fi65 < 16; fi65++ {
            var v66 bool = ((!(!false)) && (65535 > 908))
            _ = fi65
            _ = v66
        }
    case 42:
        v13 = v13
    case 0:
        for fi67 := 0; fi67 < 10; fi67++ {
            if (v9 == (^(v9))) {
                v9 = v9
                v9 = v9
            }
            v68 := GenS7{F0: v9, F1: 0, F2: 1}
            var v69 float64 = 1.5
            _ = fi67
            _ = v68
            _ = v69
        }
    default:
        var v70 uint32 = 1
        _ = v70
    }
    for fi71 := 0; fi71 < 19; fi71++ {
        for fi72 := 0; fi72 < 17; fi72++ {
            if true {
                switch (^(fi71)) {
                case 71:
                    v73 := GenS8{F0: fi71}
                    _ = v73
                case -185:
                    if true {
                        if ("\n\t" != "a") {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            fmt.Println(v13, fi72)
                        } else {
                            var v74 uint = 18446744073709551615
                            if true {
                                var v75 uint16 = 65535
                                _ = v75
                            }
                            var v76 uint32 = 0
                            _ = v74
                            _ = v76
                        }
                        var v77 uint8 = v9
                        if false {
                            switch 128 {
                            case -2147483648:
                                var v78 int32 = 484
                                _ = v78
                            }
                            if true {
                                if (v9 != v9) {
                                    if (true || ((true || true) && false)) {
                                        v79 := GenS9{F0: -2147483648, F1: 3}
                                        var v80 int = (^(fi72))
                                        switch 9223372036854775807 {
                                        case 32768:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case 239:
                                            switch fi72 {
                                            case 42:
                                                v81 := GenS8{F0: 0}
                                                _ = v81
                                            case -380:
                                                if true {
                                                    v77 = 236
                                                    switch fi72 {
                                                    case 2147483647:
                                                        var v82 bool = (!(0.0 != (-2.25 + -2.25)))
                                                        _ = v82
                                                    case -9223372036854775808:
                                                        switch v80 {
                                                        case 4294967295:
                                                            var v83 int8 = -1
                                                            _ = v83
                                                        case 42:
                                                            delete(v2, 127)
                                                        case 1:
                                                            var v84 uint64 = uint64(v9)
                                                            _ = v84
                                                        }
                                                    }
                                                }
                                            }
                                        case 208:
                                            fmt.Println(fi71, v77)
                                        }
                                        _ = v79
                                        _ = v80
                                    }
                                    v85 := GenS10{F0: true, F1: 32767}
                                    _ = v85
                                } else {
                                    v9 -= 114
                                    v13 = (((v13 * (1.5 * 0.0)) / v13) * v13)
                                }
                            } else {
                                fmt.Println(v77, fi71)
                                v13 -= ((v13 + v13) + v13)
                                var v86 bool = (-2147483648 < (-(fi71)))
                                _ = v86
                            }
                            if true {
                                switch ("\n\t" + ("" + "hello")) {
                                case "\n\t":
                                    _ = len(v1)
                                case "":
                                    if (-2.25 <= -2.25) {
                                        fmt.Println(v77)
                                        if false {
                                            if true {
                                                v87 := GenS10{F0: true, F1: int16(fi72)}
                                                _ = len(v6)
                                                v77 += (v9 - (v77 >> ((1) & 63)))
                                                _ = v87
                                            } else {
                                                _ = len(v2)
                                                v88 := GenS11{F0: (-61 <= 1)}
                                                if ((1 != 2147483648) && (!false)) {
                                                    fmt.Println(fi72, v13, v9)
                                                }
                                                _ = v88
                                            }
                                            if (!(530 > 613)) {
                                                v9 = v9
                                                v89 := GenS12{F0: 814}
                                                if len(v6) > 0 {
                                                    v6[0] = 1
                                                }
                                                _ = v89
                                            }
                                        }
                                        if ((!(128 != v77)) && (false && true)) {
                                            if false {
                                                if len(v5) > 0 {
                                                    v5[0] = fi71
                                                }
                                                delete(v1, "a")
                                                switch "foo bar" {
                                                case "hello":
                                                    v2[(fi71 - -654)] = ("hello" + "foo bar")
                                                    var v90 uint = 2
                                                    var v91 int64 = 234
                                                    _ = v90
                                                    _ = v91
                                                case "foo bar":
                                                    if (("\n\t" + "hello") > ((("a" + ("\n\t" + ("a" + "a"))) + "\n\t") + "")) {
                                                        switch fi72 {
                                                        case 502:
                                                            v77 = v77
                                                        case 127:
                                                            var v92 bool = (!false)
                                                            _ = v92
                                                        case 32768:
                                                            v9 = (v77 | 0)
                                                        default:
                                                            v9 = 164
                                                        }
                                                        if len(v6) > 0 {
                                                            v6[0] = (fi71 << ((42) & 63))
                                                        }
                                                    } else {
                                                        var v93 uint32 = 465
                                                        v93 = (v93 ^ (v93 - v93))
                                                        _ = v93
                                                    }
                                                case "":
                                                    var v94 int32 = 2147483647
                                                    _ = v94
                                                default:
                                                    v95 := v2[(fi72 >> ((genFn0(0)) & 63))]
                                                    _ = v95
                                                }
                                                v96 := v2[fi72]
                                                if (((1e10 / float64(v13)) / 1.0) < (1.5 + (-2.25 + (-2.25 / 1.5)))) {
                                                    fmt.Println(v77, v77)
                                                } else {
                                                    var v97 int64 = -1
                                                    if false {
                                                        v98 := GenS13{F0: ((("a" + v96) + v96) + ("" + v96)), F1: 42}
                                                        _ = v98
                                                    }
                                                    _ = v97
                                                }
                                                delete(v1, ("" + "a"))
                                                _ = v96
                                            }
                                            switch (fi72 * fi72) {
                                            case 0:
                                                v99 := GenS6{F0: 0, F1: true}
                                                _ = v99
                                            case 42:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case -2147483648:
                                                v9 += v77
                                            default:
                                                if true {
                                                    switch (fi72 << ((65535) & 63)) {
                                                    case 127:
                                                        v9 += 127
                                                    default:
                                                        var v100 uint16 = 49
                                                        _ = v100
                                                    }
                                                    var v101 float64 = 0.0
                                                    _ = v101
                                                } else {
                                                    switch (^(fi71)) {
                                                    case 0:
                                                        var v102 string = ("a" + "foo bar")
                                                        _ = v102
                                                    default:
                                                        switch "\n\t" {
                                                        case "a":
                                                            v103 := GenS14{F0: -561, F1: 237}
                                                            _ = v103
                                                        case "\n\t":
                                                            v13 = (v13 + v13)
                                                        case "":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        default:
                                                            var v104 int32 = 0
                                                            _ = v104
                                                        }
                                                    }
                                                    if (!(!true)) {
                                                        var v105 float32 = v13
                                                        v13 += 1.5
                                                        if true {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v106 := GenS13{F0: ("a" + ""), F1: 1}
                                                                _ = v106
                                                            case "a":
                                                                var v107 float32 = v13
                                                                _ = v107
                                                            case "":
                                                                var v108 uint32 = 90
                                                                v109 := v2[fi71]
                                                                _ = v108
                                                                _ = v109
                                                            }
                                                            var v110 int32 = -2147483648
                                                            _ = v110
                                                        } else {
                                                            v111 := GenS8{F0: (fi71 | fi71)}
                                                            fmt.Println(v105, v13)
                                                            _ = v111
                                                        }
                                                        _ = v105
                                                    } else {
                                                        if false {
                                                            var v112 int64 = 578
                                                            var v113 string = "hello"
                                                            v114 := GenS1{F0: (fi71 * (fi71 - fi71)), F1: (v9 / (((^(v9))) | 1)), F2: 100}
                                                            _ = v112
                                                            _ = v113
                                                            _ = v114
                                                        }
                                                        v77 -= 1
                                                        v13 = ((3.14159 + v13) / (3.14159 / (((((v13 * ((v13 * v13) + v13)) / float32(1.5)) * v13) + -2.25) - v13)))
                                                        v115 := GenS12{F0: 42}
                                                        _ = v115
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                if true {
                                    var v116 float64 = float64((float32(0.1) / 1e10))
                                    var v117 uint64 = 0
                                    fmt.Println(v117, v116, v117)
                                    _ = v116
                                    _ = v117
                                }
                            } else {
                                fmt.Println(fi71)
                                var v118 string = ((("a" + (("a" + "hello") + "\n\t")) + (("hello" + "\n\t") + "\n\t")) + "hello")
                                v119 := GenS14{F0: 783, F1: 65535}
                                v3 = append(v3, (fi71 / ((fi71) | 1)))
                                _ = v118
                                _ = v119
                            }
                        }
                        delete(v2, 134)
                        _ = v77
                    }
                case 4294967295:
                    switch fi72 {
                    case 42:
                        v13 -= float32(1.5)
                    }
                    _ = len(v3)
                    v120 := v4[(fi72 >> ((42) & 63))]
                    _ = v120
                }
                if true {
                    _ = len(v2)
                }
            } else {
                var v121 uint8 = v9
                v122 := v4[fi72]
                _ = v121
                _ = v122
            }
            switch "" {
            case "foo bar":
                v13 = v13
            case "hello":
                v5 = append(v5, fi72)
            case "a":
                v123 := GenS8{F0: (fi71 | fi72)}
                _ = v123
            }
            var v124 float64 = (((0.1 / 0.1) / 0.1) * (0.1 + 1.5))
            _ = fi72
            _ = v124
        }
        v2[4294967295] = ("\n\t" + ((("hello" + "foo bar") + "hello") + "\n\t"))
        _ = fi71
    }
    for fi125 := 0; fi125 < 4; fi125++ {
        var v126 string = "hello"
        v127 := GenS12{F0: 977}
        _ = fi125
        _ = v126
        _ = v127
    }
    v9 = v9
    var v128 int8 = 112
    switch -2147483648 {
    case 42:
        fmt.Println(v13, v9)
    case 127:
        delete(v4, 127)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v128:", v128)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v128
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
