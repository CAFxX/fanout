package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 bool }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS11 struct { F0 int64 }
func genFn0(p0 uint64, p1 int16) int8 {
    for fi4 := 0; fi4 < 2; fi4++ {
        p1 = p1
        var v5 bool = ("foo bar" > "\n\t")
        if v5 {
            if v5 {
                var v6 uint16 = 255
                _ = v6
            } else {
                switch ("" + "\n\t") {
                case "hello":
                    if (2147483647 <= -294) {
                        fmt.Println(p0)
                        p1 = (p1 >> ((255) & 63))
                    }
                default:
                    var v7 uint16 = 32767
                    _ = v7
                }
            }
            var v8 float32 = -2.25
            _ = v8
        }
        _ = fi4
        _ = v5
    }
    if (!true) {
        var v9 int64 = 9223372036854775807
        v10 := GenS1{F0: 42, F1: (p0 - p0), F2: 127}
        _ = v10
        _ = v9
    }
    if (-85 <= 42) {
        v11 := GenS1{F0: 0, F1: (^(p0)), F2: 128}
        var v12 bool = (65535 < 255)
        v13 := []int{-9223372036854775808, -9223372036854775808, 42}
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    return -87
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 382
    v3 := []string{(("hello" + "hello") + "")}
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS1{F0: 32767, F1: 411, F2: 42}
            v17 := GenS1{F0: 255, F1: 2, F2: 0}
            if (42 > 922) {
                v18 := GenS1{F0: 32767, F1: 97, F2: 128}
                var v19 int16 = -32768
                fmt.Println(v2, fi15)
                _ = v18
                _ = v19
            } else {
                var v20 string = (("\n\t" + ((("" + "foo bar") + "hello") + "\n\t")) + "hello")
                var v21 float64 = float64(1.5)
                _ = v20
                _ = v21
            }
            v22 := GenS1{F0: 2, F1: 55, F2: 0}
            v2 = v2
            _ = fi15
            _ = v16
            _ = v17
            _ = v22
        }
        _ = fi14
    }
    var v23 string = ("" + "foo bar")
    if true {
        v1[(v2 + (v2 >> ((42) & 63)))] = 2147483648
    } else {
        var v24 uint = 18446744073709551615
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            if (v23 < v23) {
                var v27 int8 = -10
                fmt.Println(fi26)
                _ = v27
            } else {
                v28 := GenS2{F0: 255, F1: 127, F2: false}
                v23 = (("" + (v23 + "hello")) + "hello")
                v2 = 42
                _ = v28
            }
            switch (v23 + v23) {
            case "":
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    var v29 bool = true
                    v30 := v1[(v2 + 9223372036854775807)]
                    switch -1 {
                    case 32768:
                        v1[v2] = (v30 + 42)
                    case -1:
                        v30 += (^(v30))
                    case -241:
                        var v31 string = v23
                        _ = v31
                    }
                    v2 += -2147483648
                    _ = v29
                    _ = v30
                case "":
                    fmt.Println(v2, fi26, v23)
                    v2 += (fi25 - v2)
                    _ = len(v1)
                    v32 := GenS2{F0: 419, F1: 127, F2: (65535 <= 596)}
                    var v33 uint16 = 727
                    _ = v32
                    _ = v33
                case "a":
                    if true {
                        fmt.Println(fi26, v23)
                        switch fi26 {
                        case 390:
                            v2 -= (-(fi26))
                        default:
                            v34 := GenS3{F0: 402, F1: 1}
                            _ = v34
                        }
                    }
                    if true {
                        v35 := GenS3{F0: 32768, F1: 127}
                        v2 = (v2 * v2)
                        _ = v35
                    } else {
                        v23 = "hello"
                        switch "" {
                        case "foo bar":
                            v36 := GenS2{F0: 65535, F1: 0, F2: (32767 <= 485)}
                            _ = v36
                        case "a":
                            v37 := GenS1{F0: 2, F1: 4294967295, F2: 194}
                            _ = v37
                        case "":
                            v23 = ""
                        }
                        if false {
                            switch v23 {
                            case "\n\t":
                                fmt.Println(fi25, v23, fi26)
                            case "":
                                if (-480 != -613) {
                                    fmt.Println(fi25)
                                    v38 := GenS4{F0: 42}
                                    _ = v38
                                }
                                fmt.Println(v2)
                                var v39 int = v2
                                _ = v39
                            default:
                                v1[fi25] = 65535
                            }
                            var v40 string = "foo bar"
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (true || (!true)) {
                                var v41 string = "foo bar"
                                var v42 int32 = -1
                                _ = v41
                                _ = v42
                            }
                            _ = v40
                        } else {
                            if (42 < 880) {
                                switch v23 {
                                case "a":
                                    v2 += (fi26 >> ((1) & 63))
                                case "\n\t":
                                    v43 := GenS4{F0: 2147483648}
                                    _ = v43
                                case "hello":
                                    if true {
                                        var v44 float32 = (-2.25 * 1.5)
                                        var v45 uint64 = 2
                                        fmt.Println(v44, v45, v45)
                                        _ = v44
                                        _ = v45
                                    } else {
                                        var v46 bool = true
                                        _ = v46
                                    }
                                }
                                fmt.Println(fi25, fi26, v23)
                            } else {
                                if (768 != 1) {
                                    var v47 float32 = -2.25
                                    switch (-(fi26)) {
                                    case 127:
                                        fmt.Println(v23, v47)
                                    case 32768:
                                        switch "foo bar" {
                                        case "\n\t":
                                            v48 := GenS5{F0: 684, F1: 4294967295, F2: 4294967295}
                                            _ = v48
                                        case "hello":
                                            v2 += 128
                                            v2 = (fi25 / ((fi25) | 1))
                                            v49 := GenS1{F0: 2, F1: 0, F2: 0}
                                            _ = v49
                                        case "a":
                                            v1[v2] = 4294967295
                                        }
                                    case -9223372036854775808:
                                        switch "a" {
                                        case "hello":
                                            var v50 int = fi25
                                            _ = v50
                                        }
                                    }
                                    v51 := GenS3{F0: fi25, F1: 1}
                                    v23 = "\n\t"
                                    _ = v47
                                    _ = v51
                                } else {
                                    v52 := GenS5{F0: 309, F1: 389, F2: 65535}
                                    v53 := GenS4{F0: 0}
                                    _ = v52
                                    _ = v53
                                }
                                var v54 int64 = -1
                                _ = v54
                            }
                        }
                    }
                default:
                    v2 += (v2 >> ((411) & 63))
                }
            case "foo bar":
                var v55 bool = false
                _ = v55
            default:
                var v56 int = (fi25 % ((fi26) | 1))
                _ = v56
            }
            v23 = v23
            _ = fi26
        }
        var v57 int32 = 1
        switch (fi25 / ((v2) | 1)) {
        case 1:
            if (!(genFn0(0, 1000) >= 47)) {
                v58 := GenS6{F0: 42, F1: 851}
                _ = v58
            }
        }
        _ = fi25
        _ = v57
    }
    if false {
        var v59 uint = 42
        v2 = (v2 >> (((^(v59))) & 63))
        if ((42 > -2147483648) || ((3.14159 >= 1.5) && (128 == 42))) {
            if ((v2 - (v2 / (((v2 ^ (v2 << ((v59) & 63)))) | 1))) >= v2) {
                if (0 > -848) {
                    v60 := GenS7{F0: -847, F1: 2147483647}
                    switch v2 {
                    case -9223372036854775808:
                        v23 = "foo bar"
                    case -1:
                        v61 := GenS3{F0: 56, F1: 2}
                        _ = v61
                    default:
                        switch (v2 >> (((v59 - (v59 << (((^(v59))) & 63)))) & 63)) {
                        case 128:
                            var v62 int32 = 42
                            _ = v62
                        default:
                            var v63 uint = v59
                            _ = v63
                        }
                    }
                    _ = v60
                }
                var v64 bool = true
                _ = v64
            }
            for fi65 := 0; fi65 < 2; fi65++ {
                var v66 float64 = 3.14159
                if false {
                    var v67 int64 = 0
                    _ = v67
                } else {
                    v68 := GenS7{F0: 281, F1: 355}
                    var v69 uint32 = 4294967295
                    _ = v68
                    _ = v69
                }
                _ = fi65
                _ = v66
            }
            v70 := GenS8{F0: v23}
            _ = v70
        }
        _ = v59
    } else {
        var v71 uint32 = 2147483648
        if true {
            for fi72 := 0; fi72 < 16; fi72++ {
                switch (v23 + "a") {
                case "\n\t":
                    if true {
                        fmt.Println(v2)
                        for fi73 := 0; fi73 < 12; fi73++ {
                            if (65535 < 457) {
                                if true {
                                    var v74 uint = 65535
                                    v75 := GenS7{F0: -365, F1: (v71 - 2147483647)}
                                    v23 = v23
                                    _ = v74
                                    _ = v75
                                } else {
                                    var v76 bool = ((0.0 / float32(0.1)) >= float32(-2.25))
                                    v23 = "\n\t"
                                    delete(v1, (-(fi73)))
                                    _ = v76
                                }
                                v77 := v1[fi72]
                                var v78 int = v2
                                _ = v77
                                _ = v78
                            } else {
                                var v79 uint = 301
                                var v80 uint = uint(v71)
                                _ = v79
                                _ = v80
                            }
                            _ = fi73
                        }
                        if (true && false) {
                            var v81 int16 = 32767
                            _ = v81
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = (v23 + "\n\t")
                        }
                    }
                case "":
                    switch 4294967295 {
                    case 42:
                        for fi82 := 0; fi82 < 9; fi82++ {
                            v83 := GenS3{F0: (fi82 - 42), F1: 127}
                            v84 := GenS7{F0: 42, F1: v71}
                            _ = fi82
                            _ = v83
                            _ = v84
                        }
                    case 9223372036854775807:
                        switch "" {
                        case "hello":
                            v3 = append(v3, (v23 + v23))
                        default:
                            v23 = v23
                        }
                    case 32768:
                        for fi85 := 0; fi85 < 16; fi85++ {
                            v71 -= v71
                            _ = fi85
                        }
                    default:
                        var v86 int32 = 0
                        _ = v86
                    }
                case "hello":
                    switch v2 {
                    case 9223372036854775807:
                        for fi87 := 0; fi87 < 4; fi87++ {
                            var v88 int8 = genFn0(4294967295, 42)
                            v23 = v23
                            _ = fi87
                            _ = v88
                        }
                    case 4294967295:
                        if ((fi72 * v2) <= v2) {
                            v89 := GenS8{F0: (("hello" + v23) + ("hello" + v23))}
                            switch "hello" {
                            case "hello":
                                switch 0 {
                                case -81:
                                    var v90 float32 = (3.14159 + 1e10)
                                    _ = v90
                                case 4294967295:
                                    _ = len(v3)
                                }
                            }
                            _ = v89
                        } else {
                            v2 = (^(fi72))
                            fmt.Println(v71, v23, fi72)
                        }
                    }
                default:
                    switch v2 {
                    case -569:
                        if ((0.1 / float32((0.1 + 0.0))) == -2.25) {
                            for fi91 := 0; fi91 < 1; fi91++ {
                                if len(v3) > 0 {
                                    v3[0] = v23
                                }
                                if true {
                                    v23 = (v23 + v23)
                                    switch (v2 | 4294967295) {
                                    case -455:
                                        var v92 int = fi91
                                        _ = v92
                                    }
                                } else {
                                    var v93 float64 = float64(3.14159)
                                    if false {
                                        v94 := GenS9{F0: 1}
                                        if true {
                                            var v95 uint = 146
                                            var v96 uint32 = (v71 * 42)
                                            v23 = v23
                                            _ = v95
                                            _ = v96
                                        }
                                        _ = v94
                                    } else {
                                        if false {
                                            if (!false) {
                                                v23 = (v23 + v23)
                                            } else {
                                                switch 127 {
                                                case -2147483648:
                                                    _ = len(v3)
                                                case 667:
                                                    var v97 int = (fi72 - 375)
                                                    _ = v97
                                                case -1:
                                                    _ = len(v3)
                                                default:
                                                    switch v23 {
                                                    case "\n\t":
                                                        var v98 uint8 = 1
                                                        _ = v98
                                                    default:
                                                        v2 += fi91
                                                    }
                                                }
                                                v23 = v23
                                            }
                                        }
                                    }
                                    var v99 int = v2
                                    _ = v93
                                    _ = v99
                                }
                                _ = fi91
                            }
                            var v100 bool = false
                            _ = v100
                        } else {
                            fmt.Println(v2, v23, v2)
                            if false {
                                fmt.Println(fi72, v71)
                                fmt.Println(v71, fi72, v71)
                            } else {
                                if false {
                                    var v101 int16 = -1
                                    var v102 uint = 1
                                    _ = v101
                                    _ = v102
                                }
                            }
                            v71 -= (v71 << ((65535) & 63))
                        }
                    }
                }
                _ = fi72
            }
        }
        v2 += v2
        _ = v71
    }
    var v103 uint16 = 346
    for fi104 := 0; fi104 < 14; fi104++ {
        if ("\n\t" < v23) {
            var v105 bool = (!(0.1 <= (((1.5 * 0.0) - 1.5) * (-2.25 / (1.5 * -2.25)))))
            switch ((v23 + "foo bar") + v23) {
            case "hello":
                for fi106 := 0; fi106 < 4; fi106++ {
                    v2 = fi104
                    _ = len(v1)
                    if v105 {
                        v23 = "foo bar"
                    }
                    _ = fi106
                }
            }
            _ = v105
        }
        v103 = 42
        _ = fi104
    }
    v107 := GenS10{F0: v103, F1: 1, F2: 42}
    v2 = (v2 / ((v2) | 1))
    v23 = (v23 + v23)
    var v108 bool = ((false && true) || true)
    for fi109 := 0; fi109 < 8; fi109++ {
        for fi110 := 0; fi110 < 11; fi110++ {
            if (!false) {
                fmt.Println(v103)
            }
            v111 := GenS3{F0: (^(v2)), F1: 49}
            v103 -= (v103 & (^(v103)))
            var v112 int8 = 100
            v113 := GenS4{F0: 2147483647}
            v114 := GenS10{F0: v103, F1: 42, F2: 2}
            var v115 int8 = (v112 + (v112 + 20))
            v3 = append(v3, (("\n\t" + v23) + "hello"))
            var v116 bool = v108
            _ = fi110
            _ = v111
            _ = v112
            _ = v113
            _ = v114
            _ = v115
            _ = v116
        }
        if v108 {
            v2 = (fi109 + 0)
            var v117 int64 = 1
            fmt.Println(v108, v103)
            _ = v117
        } else {
            switch (fi109 | 209) {
            case 127:
                v118 := GenS6{F0: 1, F1: 42}
                _ = v118
            default:
                fmt.Println(fi109, v108)
            }
            for fi119 := 0; fi119 < 11; fi119++ {
                _ = len(v3)
                fmt.Println(v23, fi109)
                var v120 int16 = int16(fi119)
                v120 += (v120 ^ (^(v120)))
                _ = fi119
                _ = v120
            }
            for fi121 := 0; fi121 < 11; fi121++ {
                v1[fi109] = 0
                v103 = 42
                fmt.Println(v103)
                _ = fi121
            }
        }
        if v108 {
            _ = len(v3)
        } else {
            var v122 uint8 = 0
            _ = len(v3)
            v23 = (v23 + "a")
            _ = v122
        }
        _ = fi109
    }
    v123 := GenS8{F0: v23}
    if (-2 > 0) {
        var v124 bool = false
        var v125 string = v23
        var v126 float32 = 1.5
        _ = v124
        _ = v125
        _ = v126
    }
    var v127 float64 = (3.14159 - 1e10)
    if (false || v108) {
        if v108 {
            v128 := GenS3{F0: (v2 >> ((4294967295) & 63)), F1: 201}
            _ = v128
        } else {
            var v129 int8 = 1
            v130 := GenS11{F0: 0}
            _ = v129
            _ = v130
        }
        var v131 uint32 = 42
        _ = v131
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("v108:", v108)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v103
    _ = v107
    _ = v108
    _ = v123
    _ = v127
    _ = v2
    _ = v23
    _ = v3
}
