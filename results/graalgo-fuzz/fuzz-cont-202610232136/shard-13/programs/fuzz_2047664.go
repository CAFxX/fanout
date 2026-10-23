package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int; F1 uint; F2 int }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 uint8; F1 int8; F2 int64 }
type GenS7 struct { F0 int64; F1 uint8 }
type GenS8 struct { F0 uint8; F1 int32 }
func genFn0() int16 {
    var v6 uint8 = 1
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int32 = 42
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        switch "\n\t" {
        case "hello":
            v6 = v6
        default:
            switch ("foo bar" + "\n\t") {
            case "a":
                v6 -= v6
            case "":
                var v10 float32 = 3.14159
                _ = v10
            case "hello":
                v11 := []int{-588, (fi9 % ((fi9) | 1))}
                _ = v11
            default:
                switch (((("a" + "hello") + "foo bar") + "") + "a") {
                case "\n\t":
                    var v12 int16 = -32768
                    _ = v12
                case "":
                    v6 -= (v6 | v6)
                }
            }
        }
        var v13 int16 = 0
        var v14 uint16 = 42
        _ = fi9
        _ = v13
        _ = v14
    }
    v6 -= (v6 + v6)
    _ = v6
    return 42
}
func genFn1(p0 int64) int8 {
    p0 -= (^(p0))
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return -1
}

func main() {
    v1 := []uint32{0}
    v2 := make(map[uint32]uint32)
    var v3 string = ("" + "\n\t")
    var v4 uint64 = 1
    var v5 int8 = 42
    var v15 uint8 = 0
    switch v3 {
    case "a":
        v4 -= uint64(v15)
    case "\n\t":
        v1 = append(v1, 553)
    case "foo bar":
        switch 4294967295 {
        case 32768:
            var v16 uint64 = v4
            _ = v16
        case 1:
            fmt.Println(v3, v15, v4)
        case -2147483648:
            if true {
                if (false || false) {
                    v15 = 255
                    for fi17 := 0; fi17 < 15; fi17++ {
                        fmt.Println(v4, v4)
                        var v18 uint64 = v4
                        _ = fi17
                        _ = v18
                    }
                    var v19 bool = true
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        _ = len(v2)
                        switch fi20 {
                        case 32767:
                            v21 := GenS1{F0: 1}
                            _ = v21
                        case -9223372036854775808:
                            var v22 bool = false
                            _ = v22
                        case -816:
                            var v23 float64 = (0.0 / 1.0)
                            _ = v23
                        default:
                            fmt.Println(v3, v5, v15)
                        }
                        _ = fi20
                    }
                    if ((v15 | v15) == 42) {
                        switch ((("a" + (v3 + "hello")) + ("hello" + (v3 + "hello"))) + v3) {
                        case "hello":
                            if true {
                                var v24 float64 = float64((((3.14159 - float32(1.5)) - ((0.0 - (((3.14159 / 1.0) * 1.5) - 3.14159)) + -2.25)) / (((3.14159 - 0.0) / ((1e10 + (0.1 + 0.0)) * float32(float64(1e10)))) / 0.1)))
                                v25 := GenS1{F0: -1}
                                v26 := GenS1{F0: 346}
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                        case "":
                            var v27 int64 = 1
                            _ = v27
                        case "a":
                            if true {
                                v28 := v2[65535]
                                if (!(-1 > 0)) {
                                    if true {
                                        if len(v1) > 0 {
                                            v1[0] = (v28 - uint32(v4))
                                        }
                                    }
                                    var v29 int16 = 0
                                    _ = v29
                                } else {
                                    var v30 int = -2147483648
                                    v31 := GenS1{F0: 0}
                                    var v32 int = v30
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                                _ = v28
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                for fi33 := 0; fi33 < 17; fi33++ {
                                    for fi34 := 0; fi34 < 16; fi34++ {
                                        v35 := v2[4294967295]
                                        v4 -= v4
                                        _ = fi34
                                        _ = v35
                                    }
                                    for fi36 := 0; fi36 < 11; fi36++ {
                                        var v37 bool = false
                                        if (0 != 65535) {
                                            fmt.Println(v3, v4)
                                        } else {
                                            var v38 uint16 = 384
                                            if v37 {
                                                if (v38 <= 32767) {
                                                    v5 -= 127
                                                }
                                            }
                                            v3 = ""
                                            var v39 uint32 = 65535
                                            v39 = 1
                                            _ = v38
                                            _ = v39
                                        }
                                        _ = fi36
                                        _ = v37
                                    }
                                    for fi40 := 0; fi40 < 15; fi40++ {
                                        v41 := GenS2{F0: (42 >= 32767), F1: 735}
                                        var v42 uint = 358
                                        _ = fi40
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = fi33
                                }
                            }
                        default:
                            v5 += (v5 & v5)
                        }
                        v3 = "hello"
                        if true {
                            if false {
                                v4 += (v4 + uint64(v15))
                            } else {
                                var v43 int32 = -625
                                var v44 int16 = 795
                                _ = v43
                                _ = v44
                            }
                        } else {
                            switch v3 {
                            case "foo bar":
                                var v45 int8 = v5
                                _ = v45
                            case "":
                                var v46 uint8 = (v15 / ((v15) | 1))
                                _ = v46
                            case "hello":
                                switch v3 {
                                case "\n\t":
                                    var v47 int8 = (v5 << ((18446744073709551615) & 63))
                                    _ = v47
                                case "foo bar":
                                    v3 = v3
                                case "a":
                                    if true {
                                        v5 = v5
                                    } else {
                                        delete(v2, 65535)
                                    }
                                default:
                                    var v48 float32 = float32(0.1)
                                    _ = v48
                                }
                            default:
                                fmt.Println(v4, v3, v3)
                            }
                            v49 := GenS2{F0: false, F1: 590}
                            _ = v49
                        }
                    }
                }
                if false {
                    switch v3 {
                    case "foo bar":
                        v50 := GenS2{F0: false, F1: 255}
                        _ = v50
                    case "hello":
                        delete(v2, 0)
                    case "":
                        for fi51 := 0; fi51 < 6; fi51++ {
                            for fi52 := 0; fi52 < 14; fi52++ {
                                var v53 uint8 = (v15 << ((0) & 63))
                                _ = fi52
                                _ = v53
                            }
                            _ = fi51
                        }
                    }
                    if (!false) {
                        var v54 float64 = 0.0
                        fmt.Println(v3, v4, v5)
                        var v55 uint8 = v15
                        _ = v54
                        _ = v55
                    } else {
                        fmt.Println(v3)
                        for fi56 := 0; fi56 < 7; fi56++ {
                            fmt.Println(v4, v15)
                            var v57 uint16 = 242
                            _ = fi56
                            _ = v57
                        }
                    }
                } else {
                    switch ((v3 + "foo bar") + v3) {
                    case "foo bar":
                        switch 32768 {
                        case -2147483648:
                            fmt.Println(v4, v15)
                        case -564:
                            v58 := GenS3{F0: v5, F1: 32767}
                            _ = v58
                        case -9223372036854775808:
                            var v59 uint32 = 42
                            _ = v59
                        default:
                            switch 2147483647 {
                            case -1:
                                var v60 float64 = 1.5
                                _ = v60
                            }
                        }
                    default:
                        v15 = v15
                    }
                }
                if false {
                    for fi61 := 0; fi61 < 2; fi61++ {
                        v4 += 589
                        var v62 int64 = -1
                        if (363 != 18446744073709551615) {
                            for fi63 := 0; fi63 < 20; fi63++ {
                                var v64 uint = 18446744073709551615
                                switch fi61 {
                                case 2147483647:
                                    if true {
                                        v65 := GenS1{F0: 175}
                                        _ = v65
                                    }
                                case 715:
                                    fmt.Println(v64, v62)
                                case 128:
                                    var v66 int = fi63
                                    _ = v66
                                default:
                                    switch v3 {
                                    case "\n\t":
                                        v2[4294967295] = 2147483648
                                        var v67 float32 = float32(((0.1 - 3.14159) * ((float64(3.14159) + ((1.5 / 3.14159) + float64(-2.25))) * (0.0 + 3.14159))))
                                        fmt.Println(v15)
                                        v68 := GenS4{F0: fi61, F1: v64, F2: fi61}
                                        _ = v67
                                        _ = v68
                                    case "":
                                        switch (^(fi63)) {
                                        case 728:
                                            v69 := v2[2147483648]
                                            _ = v69
                                        case -501:
                                            var v70 uint16 = 274
                                            _ = v70
                                        default:
                                            if false {
                                                v71 := GenS1{F0: 0}
                                                _ = v71
                                            } else {
                                                var v72 string = (v3 + v3)
                                                switch 1 {
                                                case 865:
                                                    if false {
                                                        var v73 uint8 = v15
                                                        _ = v73
                                                    }
                                                }
                                                _ = v72
                                            }
                                        }
                                    case "a":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        var v74 int32 = 1
                                        _ = v74
                                    }
                                }
                                _ = fi63
                                _ = v64
                            }
                            v1 = append(v1, 0)
                            if true {
                                fmt.Println(v3, v4)
                                if ((("a" + "foo bar") == v3) || true) {
                                    if (!(!(!(!true)))) {
                                        var v75 int32 = -15
                                        _ = v75
                                    } else {
                                        var v76 int = fi61
                                        _ = v76
                                    }
                                }
                            } else {
                                var v77 int64 = v62
                                _ = v77
                            }
                        }
                        _ = fi61
                        _ = v62
                    }
                }
            } else {
                var v78 uint16 = 2
                v79 := v2[2]
                v80 := v2[(v79 + v79)]
                _ = v78
                _ = v79
                _ = v80
            }
        }
    default:
        if true {
            if (((596 == 4294967295) && true) || (4294967295 == -2147483648)) {
                delete(v2, 746)
                if ((1e10 + ((3.14159 / 1.0) + (((1.5 / (0.0 + 3.14159)) / 1.5) / 0.1))) == 1e10) {
                    var v81 uint64 = v4
                    _ = v81
                } else {
                    switch 2147483647 {
                    case 4294967295:
                        for fi82 := 0; fi82 < 17; fi82++ {
                            fmt.Println(v4)
                            _ = fi82
                        }
                    }
                }
            }
            var v83 string = v3
            for fi84 := 0; fi84 < 8; fi84++ {
                var v85 uint64 = v4
                _ = fi84
                _ = v85
            }
            _ = v83
        }
    }
    if false {
        fmt.Println(v5)
    }
    switch 629 {
    case -1:
        if (float32(0.0) <= 0.1) {
            for fi86 := 0; fi86 < 17; fi86++ {
                switch "foo bar" {
                case "foo bar":
                    v3 = v3
                case "\n\t":
                    for fi87 := 0; fi87 < 6; fi87++ {
                        v5 -= genFn1(-1)
                        v2[1] = 4294967295
                        _ = fi87
                    }
                case "hello":
                    v5 = 0
                default:
                    switch v3 {
                    case "a":
                        switch (fi86 & 128) {
                        case -2147483648:
                            for fi88 := 0; fi88 < 18; fi88++ {
                                var v89 uint8 = v15
                                _ = fi88
                                _ = v89
                            }
                        case 127:
                            for fi90 := 0; fi90 < 18; fi90++ {
                                var v91 int64 = 0
                                delete(v2, 2147483648)
                                _ = fi90
                                _ = v91
                            }
                        case 772:
                            fmt.Println(fi86, v15)
                        default:
                            for fi92 := 0; fi92 < 5; fi92++ {
                                v5 = (v5 - -128)
                                var v93 int32 = 0
                                _ = fi92
                                _ = v93
                            }
                        }
                    }
                }
                _ = fi86
            }
            v5 = v5
            v4 += v4
        }
    }
    v4 = (v4 * v4)
    v94 := v2[2147483647]
    if false {
        var v95 uint32 = v94
        v96 := GenS5{F0: -578, F1: (v4 / ((288) | 1))}
        _ = len(v2)
        _ = v95
        _ = v96
    } else {
        v15 = 2
        for fi97 := 0; fi97 < 18; fi97++ {
            for fi98 := 0; fi98 < 2; fi98++ {
                var v99 uint = 65535
                _ = fi98
                _ = v99
            }
            if (-603 <= 1) {
                for fi100 := 0; fi100 < 14; fi100++ {
                    v101 := GenS1{F0: -2147483648}
                    delete(v2, 0)
                    _ = fi100
                    _ = v101
                }
            } else {
                v94 = (v94 << ((2) & 63))
                v102 := v2[v94]
                _ = v102
            }
            v3 = (v3 + (v3 + ""))
            _ = fi97
        }
        for fi103 := 0; fi103 < 5; fi103++ {
            fmt.Println(fi103, fi103, fi103)
            switch (-(fi103)) {
            case 42:
                switch fi103 {
                case 32768:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case -9223372036854775808:
                    if true {
                        var v104 int8 = (^(v5))
                        _ = v104
                    } else {
                        if false {
                            _ = len(v1)
                            for fi105 := 0; fi105 < 6; fi105++ {
                                fmt.Println(v3, v3, v15)
                                v15 = (v15 >> ((42) & 63))
                                v5 -= (v5 >> ((515) & 63))
                                _ = fi105
                            }
                            var v106 uint8 = 128
                            _ = v106
                        }
                        var v107 uint32 = 0
                        v3 = ""
                        _ = v107
                    }
                default:
                    v108 := GenS6{F0: (v15 | (v15 - v15)), F1: genFn1(0), F2: 147}
                    _ = v108
                }
            case 9223372036854775807:
                switch v3 {
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "":
                    v2[v94] = (v94 % (((v94 ^ (v94 & 0))) | 1))
                case "a":
                    for fi109 := 0; fi109 < 10; fi109++ {
                        if (0 != -1) {
                            fmt.Println(fi109, v94)
                        } else {
                            var v110 int8 = 100
                            if false {
                                if ((3.14159 / (1.5 + (1e10 / 0.1))) != 3.14159) {
                                    if (1 < 0) {
                                        var v111 int16 = 1
                                        _ = v111
                                    } else {
                                        v94 -= (^(v94))
                                        if (42 <= 1000) {
                                            v2[v94] = (v94 * v94)
                                            if (false || ("" == ("foo bar" + ""))) {
                                                v3 = ((v3 + "") + v3)
                                                if true {
                                                    v1 = append(v1, v94)
                                                    var v112 int64 = -548
                                                    _ = v112
                                                }
                                                var v113 int32 = 190
                                                _ = v113
                                            } else {
                                                v114 := GenS7{F0: 54, F1: 149}
                                                _ = v114
                                            }
                                        }
                                    }
                                    switch fi103 {
                                    case 0:
                                        fmt.Println(v110, v5)
                                    case 4294967295:
                                        switch v3 {
                                        case "a":
                                            v3 = "a"
                                        }
                                        v115 := GenS8{F0: (v15 >> ((4294967295) & 63)), F1: 42}
                                        _ = v115
                                    case 32768:
                                        v4 = (v4 ^ (^(v4)))
                                    default:
                                        delete(v2, v94)
                                    }
                                    v116 := GenS3{F0: -1, F1: 255}
                                    _ = v116
                                } else {
                                    if false {
                                        v3 = "hello"
                                        v3 = v3
                                    }
                                }
                            } else {
                                if false {
                                    var v117 string = v3
                                    _ = v117
                                }
                                v4 -= (v4 ^ v4)
                                var v118 uint = 18446744073709551615
                                _ = v118
                            }
                            _ = v110
                        }
                        _ = fi109
                    }
                }
            case 2147483647:
                v15 -= v15
            }
            fmt.Println(v3, v94)
            _ = fi103
        }
    }
    v94 += v94
    if (((0.0 * ((3.14159 + 1e10) * (1e10 / 1e10))) > 3.14159) || false) {
        var v119 int = -9223372036854775808
        v94 -= v94
        _ = v119
    } else {
        if (!(!true)) {
            var v120 int64 = 394
            switch v3 {
            case "":
                var v121 uint16 = 32767
                _ = v121
            case "hello":
                v122 := GenS3{F0: v5, F1: 42}
                _ = v122
            case "a":
                var v123 bool = true
                _ = v123
            }
            fmt.Println(v15, v15)
            _ = v120
        }
    }
    v15 += v15
    var v124 float64 = -2.25
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v124:", v124)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v124
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v94
}
