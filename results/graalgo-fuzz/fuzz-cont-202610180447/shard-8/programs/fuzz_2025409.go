package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }
type GenS3 struct { F0 uint16; F1 int8; F2 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64; F1 int32; F2 int }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint64; F1 int8 }
type GenS11 struct { F0 uint; F1 int16 }
type GenS12 struct { F0 int8; F1 int32; F2 int8 }
type GenS13 struct { F0 uint8; F1 int16 }
type GenS14 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS15 struct { F0 int; F1 int; F2 bool }
func genFn0(p0 uint8) uint {
    var v7 float32 = 0.0
    switch "foo bar" {
    case "a":
        v8 := GenS1{F0: 0, F1: p0}
        _ = v8
    case "foo bar":
        var v9 int64 = 0
        _ = v9
    }
    switch -352 {
    case 128:
        v10 := GenS2{F0: (p0 + p0), F1: 42, F2: (("hello" + "hello") + "")}
        _ = v10
    case -744:
        v11 := GenS3{F0: 255, F1: -78, F2: 1}
        _ = v11
    default:
        v12 := []int{-733, -252, 2147483647}
        _ = v12
    }
    var v13 uint64 = 18446744073709551615
    v14 := GenS3{F0: 42, F1: -85, F2: 0}
    _ = p0
    _ = v13
    _ = v14
    _ = v7
    return 401
}
func genFn1(p0 int, p1 string, p2 int) uint {
    fmt.Println(p1)
    var v15 int = 2147483647
    switch "foo bar" {
    case "a":
        if true {
            p2 = v15
            v16 := GenS1{F0: 2, F1: 128}
            _ = v16
        } else {
            var v17 int = 4294967295
            if (976 >= 42) {
                v18 := []int{}
                v19 := GenS4{F0: 0}
                fmt.Println(p0, v15)
                _ = v18
                _ = v19
            }
            for fi20 := 0; fi20 < 10; fi20++ {
                v21 := []int{}
                v21 = append(v21, (fi20 % (((v15 - (p0 >> ((569) & 63)))) | 1)))
                var v22 uint = genFn0(127)
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v17
        }
    case "":
        p1 = (p1 + ("\n\t" + p1))
    case "\n\t":
        p0 = (p0 >> ((757) & 63))
    default:
        var v23 uint8 = 42
        _ = v23
    }
    p2 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 631
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{4294967295, 65535}
    var v3 uint = 255
    var v4 int8 = -128
    var v5 uint64 = 0
    v6 := make(map[uint32]uint32)
    for fi24 := 0; fi24 < 20; fi24++ {
        v25 := GenS3{F0: 936, F1: v4, F2: 752}
        var v26 bool = (993 <= 255)
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 bool = true
    for fi28 := 0; fi28 < 9; fi28++ {
        v29 := v1[fi28]
        v30 := GenS5{F0: 535}
        _ = fi28
        _ = v29
        _ = v30
    }
    if (true || (1 < -1)) {
        _ = len(v2)
        var v31 float32 = 1e10
        _ = v31
    }
    v3 = (^(v3))
    switch (((("foo bar" + "hello") + "a") + "a") + "hello") {
    case "hello":
        if v27 {
            fmt.Println(v4, v5)
        }
    }
    if true {
        if v27 {
            if (v27 || ((v5 >> ((v3) & 63)) < v5)) {
                if v27 {
                    switch ("foo bar" + "a") {
                    case "hello":
                        v27 = true
                    case "\n\t":
                        if false {
                            v4 = (v4 | (v4 ^ (v4 & v4)))
                            v3 = v3
                            for fi32 := 0; fi32 < 17; fi32++ {
                                for fi33 := 0; fi33 < 19; fi33++ {
                                    switch ("foo bar" + "") {
                                    case "foo bar":
                                        v34 := GenS6{F0: v3}
                                        _ = v34
                                    case "hello":
                                        if false {
                                            fmt.Println(v4)
                                        } else {
                                            v4 -= v4
                                            v3 = (^(v3))
                                            if (-32768 != 338) {
                                                var v35 int64 = 1
                                                v27 = true
                                                v5 -= 323
                                                _ = v35
                                            } else {
                                                var v36 float64 = (1.5 / ((0.1 + float64(1e10)) - 3.14159))
                                                switch 2147483647 {
                                                case 836:
                                                    v37 := GenS4{F0: 2}
                                                    _ = v37
                                                }
                                                delete(v6, 42)
                                                _ = v36
                                            }
                                            var v38 int32 = -871
                                            _ = v38
                                        }
                                        if ("a" == ("hello" + "a")) {
                                            var v39 float32 = 1.5
                                            v40 := GenS7{F0: -1}
                                            _ = v39
                                            _ = v40
                                        } else {
                                            var v41 uint32 = 319
                                            _ = v41
                                        }
                                        var v42 string = ("\n\t" + ("foo bar" + "hello"))
                                        _ = v42
                                    case "\n\t":
                                        fmt.Println(v4)
                                    }
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(v3, fi32)
                                    case "a":
                                        if v27 {
                                            if ((0 < 9223372036854775807) && (4294967295 < 0)) {
                                                if (v27 && (!((false || true) && v27))) {
                                                    var v43 int32 = 47
                                                    var v44 int8 = v4
                                                    _ = len(v1)
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    v27 = true
                                                    var v45 uint32 = 0
                                                    _ = v45
                                                }
                                            }
                                            v5 = (v5 / ((1) | 1))
                                            var v46 uint64 = v5
                                            if v27 {
                                                switch 372 {
                                                case 0:
                                                    v47 := GenS4{F0: 2147483647}
                                                    _ = v47
                                                case 4294967295:
                                                    var v48 uint32 = 805
                                                    _ = v48
                                                case 1:
                                                    v46 -= (v46 / (((v46 * 2)) | 1))
                                                }
                                            } else {
                                                v49 := GenS4{F0: 2}
                                                switch fi33 {
                                                case 1:
                                                    if v27 {
                                                        if v27 {
                                                            v27 = (196 != 2)
                                                            v2 = append(v2, 0)
                                                            fmt.Println(v4, v5)
                                                        } else {
                                                            v27 = true
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(fi32, fi33, v5)
                                                }
                                                var v50 float64 = 0.1
                                                _ = v49
                                                _ = v50
                                            }
                                            v4 += -128
                                            v3 -= genFn1(fi33, ("a" + (("\n\t" + "a") + "\n\t")), fi32)
                                            _ = v46
                                        }
                                    case "foo bar":
                                        if (255 <= 128) {
                                            var v51 int32 = -2147483648
                                            if false {
                                                var v52 float32 = (0.1 / float32((0.1 + 0.0)))
                                                _ = v52
                                            }
                                            _ = v51
                                        } else {
                                            var v53 float32 = 0.0
                                            v2 = append(v2, 0)
                                            v27 = v27
                                            _ = v53
                                        }
                                    }
                                    var v54 int8 = v4
                                    _ = fi33
                                    _ = v54
                                }
                                v27 = true
                                v4 -= 127
                                _ = fi32
                            }
                        } else {
                            v4 += -1
                            var v55 uint8 = 2
                            v55 = (v55 << ((v3) & 63))
                            _ = v55
                        }
                    case "a":
                        var v56 float64 = ((0.0 + ((-2.25 * 1.5) - 3.14159)) - (-2.25 / 1.0))
                        _ = v56
                    default:
                        if false {
                            var v57 float32 = float32(1.5)
                            _ = v57
                        }
                    }
                    switch -354 {
                    case 746:
                        fmt.Println(v27, v27)
                    default:
                        var v58 int8 = (v4 * v4)
                        _ = v58
                    }
                } else {
                    var v59 uint = v3
                    if v27 {
                        fmt.Println(v27, v4, v5)
                        if ((!v27) || (!v27)) {
                            for fi60 := 0; fi60 < 17; fi60++ {
                                v61 := GenS8{F0: 0, F1: 1, F2: (^(fi60))}
                                fmt.Println(v5, v27)
                                _ = fi60
                                _ = v61
                            }
                            var v62 int8 = v4
                            v59 = (v59 | v59)
                            _ = v62
                        } else {
                            v63 := GenS5{F0: 308}
                            v64 := GenS6{F0: 65535}
                            v27 = (v5 != (v5 * (v5 >> (((v59 * 65535)) & 63))))
                            _ = v63
                            _ = v64
                        }
                        var v65 int = -565
                        _ = v65
                    } else {
                        switch ("" + "\n\t") {
                        case "\n\t":
                            if v27 {
                                if (1 == 0) {
                                    switch (("" + "") + "hello") {
                                    case "hello":
                                        v66 := v6[2]
                                        _ = v66
                                    }
                                    for fi67 := 0; fi67 < 1; fi67++ {
                                        v68 := GenS6{F0: v59}
                                        v3 -= (v59 - genFn1((-(fi67)), (("\n\t" + "\n\t") + "\n\t"), 128))
                                        fmt.Println(fi67, v27, v27)
                                        _ = fi67
                                        _ = v68
                                    }
                                    var v69 int16 = -1
                                    _ = v69
                                }
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                for fi70 := 0; fi70 < 2; fi70++ {
                                    for fi71 := 0; fi71 < 19; fi71++ {
                                        var v72 bool = (0 > 78)
                                        var v73 uint16 = 885
                                        _ = fi71
                                        _ = v72
                                        _ = v73
                                    }
                                    _ = fi70
                                }
                            }
                        case "a":
                            for fi74 := 0; fi74 < 2; fi74++ {
                                switch (fi74 << (((v3 / (((v59 / ((4294967295) | 1))) | 1))) & 63)) {
                                case 9223372036854775807:
                                    for fi75 := 0; fi75 < 19; fi75++ {
                                        switch "\n\t" {
                                        case "":
                                            v76 := v6[65535]
                                            v77 := GenS9{F0: (v4 / (((v4 / ((v4) | 1))) | 1))}
                                            var v78 int32 = -2147483648
                                            _ = v76
                                            _ = v77
                                            _ = v78
                                        case "a":
                                            v2 = append(v2, 42)
                                        case "\n\t":
                                            v1[840] = -1
                                            v79 := GenS10{F0: (v5 | 2), F1: -128}
                                            v59 += uint(v5)
                                            _ = v79
                                        }
                                        fmt.Println(fi75)
                                        v4 -= (v4 % ((127) | 1))
                                        _ = fi75
                                    }
                                case 128:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case 797:
                                    switch fi74 {
                                    case -9223372036854775808:
                                        var v80 bool = (9223372036854775807 != 525)
                                        _ = v80
                                    default:
                                        for fi81 := 0; fi81 < 11; fi81++ {
                                            var v82 int = fi74
                                            switch (-(fi74)) {
                                            case 32768:
                                                var v83 float32 = 0.0
                                                _ = v83
                                            case 422:
                                                switch (^(v82)) {
                                                case -2147483648:
                                                    v84 := GenS11{F0: v59, F1: -1}
                                                    if ((1e10 <= 0.1) || (!v27)) {
                                                        _ = len(v6)
                                                        switch 2147483647 {
                                                        case 9223372036854775807:
                                                            fmt.Println(fi74, v27)
                                                        default:
                                                            v3 = v59
                                                        }
                                                        v85 := GenS12{F0: v4, F1: 2147483647, F2: v4}
                                                        v86 := v6[787]
                                                        _ = v85
                                                        _ = v86
                                                    } else {
                                                        v87 := GenS11{F0: (v3 - (v3 >> ((genFn1(fi81, "", 2147483647)) & 63))), F1: 0}
                                                        var v88 float64 = (0.1 * ((0.1 / 3.14159) + (3.14159 - float64(1.5))))
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                    _ = v84
                                                case 32767:
                                                    var v89 bool = true
                                                    _ = v89
                                                case 4294967295:
                                                    v5 += (^(v5))
                                                default:
                                                    fmt.Println(v4, v59, v27)
                                                }
                                            case 4294967295:
                                                fmt.Println(v5)
                                            }
                                            _ = fi81
                                            _ = v82
                                        }
                                    }
                                default:
                                    v3 = (^(v3))
                                }
                                var v90 int = fi74
                                _ = fi74
                                _ = v90
                            }
                        default:
                            if v27 {
                                fmt.Println(v27, v5, v4)
                                for fi91 := 0; fi91 < 7; fi91++ {
                                    if true {
                                        for fi92 := 0; fi92 < 1; fi92++ {
                                            switch fi92 {
                                            case 127:
                                                delete(v6, 132)
                                            case 128:
                                                var v93 int16 = 42
                                                _ = v93
                                            case 550:
                                                v94 := GenS13{F0: 42, F1: -32768}
                                                _ = v94
                                            default:
                                                var v95 int16 = int16(v4)
                                                _ = v95
                                            }
                                            _ = fi92
                                        }
                                        var v96 int32 = -250
                                        _ = v96
                                    }
                                    if (127 == 91) {
                                        var v97 uint8 = 0
                                        _ = v97
                                    }
                                    for fi98 := 0; fi98 < 15; fi98++ {
                                        if (0 > (fi98 ^ (fi98 + (fi98 / ((fi91) | 1))))) {
                                            fmt.Println(v3, v4, v4)
                                        }
                                        v4 = (v4 >> (((v59 << (((^(v59))) & 63))) & 63))
                                        _ = fi98
                                    }
                                    _ = fi91
                                }
                                for fi99 := 0; fi99 < 1; fi99++ {
                                    if (v59 < 65535) {
                                        var v100 uint = uint(v5)
                                        _ = v100
                                    }
                                    var v101 uint = 0
                                    var v102 int8 = (v4 >> ((genFn0(75)) & 63))
                                    _ = fi99
                                    _ = v101
                                    _ = v102
                                }
                            } else {
                                var v103 bool = false
                                v4 = v4
                                _ = v103
                            }
                        }
                        var v104 int = -147
                        _ = v104
                    }
                    _ = v59
                }
                v27 = false
            }
        } else {
            v105 := GenS14{F0: 128, F1: -1, F2: 152}
            switch 2147483647 {
            case 740:
                v2 = append(v2, 2147483648)
            default:
                var v106 float64 = (float64((((3.14159 - (0.1 + float32(1e10))) * 0.0) - 0.0)) * 0.1)
                _ = v106
            }
            _ = v105
        }
    } else {
        var v107 bool = (321 < -2147483648)
        v5 = (v5 ^ v5)
        for fi108 := 0; fi108 < 14; fi108++ {
            switch 1 {
            case 127:
                if (v27 || true) {
                    if true {
                        fmt.Println(v5, v27, v4)
                        var v109 int8 = v4
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v109
                    } else {
                        v1[fi108] = (fi108 ^ (-(fi108)))
                    }
                }
            case 128:
                var v110 bool = true
                _ = v110
            case -776:
                var v111 uint32 = 607
                _ = v111
            default:
                fmt.Println(fi108)
            }
            v5 = (^(v5))
            _ = fi108
        }
        _ = v107
    }
    if false {
        v27 = v27
        switch ((("\n\t" + ("hello" + ("hello" + "\n\t"))) + ("\n\t" + "\n\t")) + (("hello" + "a") + "\n\t")) {
        case "a":
            var v112 float32 = 1e10
            _ = v112
        case "\n\t":
            for fi113 := 0; fi113 < 14; fi113++ {
                var v114 int8 = -128
                for fi115 := 0; fi115 < 9; fi115++ {
                    v116 := GenS8{F0: 574, F1: -1, F2: (fi113 | -1)}
                    _ = fi115
                    _ = v116
                }
                _ = fi113
                _ = v114
            }
        case "foo bar":
            for fi117 := 0; fi117 < 3; fi117++ {
                v118 := GenS1{F0: v5, F1: 2}
                v4 -= -54
                _ = fi117
                _ = v118
            }
        }
        fmt.Println(v5, v27, v3)
    } else {
        fmt.Println(v3, v5)
        if true {
            switch "a" {
            case "\n\t":
                var v119 float64 = (0.0 - -2.25)
                _ = v119
            case "hello":
                v4 = (-(v4))
            case "":
                fmt.Println(v5, v3, v5)
            }
        } else {
            v120 := GenS15{F0: 32767, F1: -1, F2: ((v27 || (1 < 212)) && ((-32768 <= 32767) || false))}
            _ = v120
        }
    }
    if v27 {
        var v121 uint8 = 140
        fmt.Println(v3, v121)
        _ = v121
    } else {
        for fi122 := 0; fi122 < 8; fi122++ {
            delete(v6, 565)
            _ = fi122
        }
        for fi123 := 0; fi123 < 5; fi123++ {
            var v124 int8 = v4
            _ = fi123
            _ = v124
        }
        fmt.Println(v5, v4, v5)
    }
    v3 = (v3 % (((v3 * (^(v3)))) | 1))
    v125 := GenS9{F0: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v125
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
