package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint16; F1 uint }
type GenS11 struct { F0 bool; F1 int8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int16; F1 string; F2 string }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint; F1 int64 }
type GenS16 struct { F0 int; F1 int16 }
type GenS17 struct { F0 uint8; F1 int64; F2 int64 }
type GenS18 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS19 struct { F0 uint16; F1 uint16 }
type GenS20 struct { F0 uint64; F1 uint16 }
type GenS21 struct { F0 int16 }
func genFn0(p0 int64, p1 string) int16 {
    v7 := []int{1, 32768, 2147483647, 0}
    var v8 bool = false
    v7 = append(v7, 4294967295)
    var v9 uint = 1
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            p0 = (p0 >> ((v9) & 63))
            _ = fi11
        }
        if ((0.0 * 3.14159) != 3.14159) {
            _ = len(v7)
            v12 := GenS1{F0: v9, F1: 42}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 12; fi13++ {
                if v8 {
                    switch "hello" {
                    case "a":
                        if (float32((float64((1.5 + -2.25)) - ((float64(1e10) + (((((1e10 + (1e10 + 0.1)) - 1.5) / 1.0) + (0.1 + 3.14159)) * (0.1 - (float64(float32(3.14159)) * (-2.25 / ((float64(-2.25) / (float64(3.14159) + (1.5 + (0.1 * 3.14159)))) * (((0.0 + (0.1 * (1e10 - 3.14159))) - 3.14159) * ((1.5 / 1.0) / 0.1)))))))) / -2.25))) < 3.14159) {
                            v14 := GenS1{F0: v9, F1: 0}
                            _ = v14
                        }
                    case "":
                        switch fi13 {
                        case 671:
                            v7 = append(v7, 4294967295)
                        case 981:
                            p0 = 1
                        case 127:
                            switch fi13 {
                            case -726:
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            default:
                                fmt.Println(p0, p1, fi13)
                            }
                        }
                    default:
                        switch "foo bar" {
                        case "a":
                            var v15 int32 = -289
                            _ = v15
                        case "\n\t":
                            v16 := GenS2{F0: -749}
                            _ = v16
                        }
                    }
                    switch (fi13 * (-(fi10))) {
                    case -120:
                        _ = len(v7)
                    case 4294967295:
                        var v17 bool = (0 <= -2147483648)
                        _ = v17
                    case 264:
                        var v18 bool = (65535 > 1)
                        _ = v18
                    }
                    switch (fi10 & fi10) {
                    case 2147483647:
                        var v19 int16 = 1000
                        _ = v19
                    case -9223372036854775808:
                        p0 -= -1
                    case 32767:
                        p1 = "a"
                    default:
                        fmt.Println(v9, fi10, fi13)
                        switch p1 {
                        case "":
                            v20 := GenS1{F0: 1, F1: 100}
                            _ = v20
                        case "foo bar":
                            if (-1 >= -128) {
                                switch (^(fi13)) {
                                case -2147483648:
                                    fmt.Println(p1, v9)
                                case 165:
                                    if (42 < 0) {
                                        v21 := GenS1{F0: 18446744073709551615, F1: 42}
                                        var v22 int64 = -9223372036854775808
                                        _ = v21
                                        _ = v22
                                    }
                                }
                            } else {
                                if (800 >= 0) {
                                    v8 = (255 < 65535)
                                    v23 := GenS1{F0: v9, F1: 1}
                                    _ = v23
                                } else {
                                    p1 = p1
                                    p1 = p1
                                    var v24 int8 = 42
                                    _ = v24
                                }
                                var v25 int32 = -262
                                if (359 <= 42) {
                                    if (!v8) {
                                        if false {
                                            var v26 bool = v8
                                            p0 -= (^(p0))
                                            fmt.Println(p0)
                                            _ = v26
                                        } else {
                                            if (65535 < 236) {
                                                if (4294967295 == 865) {
                                                    v25 += v25
                                                    fmt.Println(v9, v9, p1)
                                                } else {
                                                    p0 = p0
                                                    if v8 {
                                                        switch fi10 {
                                                        case 885:
                                                            if v8 {
                                                                var v27 int16 = -144
                                                                _ = v27
                                                            } else {
                                                                v8 = ((!v8) || v8)
                                                                fmt.Println(p1, p1, p0)
                                                            }
                                                        default:
                                                            var v28 int64 = 42
                                                            _ = v28
                                                        }
                                                        v25 = int32(v9)
                                                    }
                                                    v29 := GenS3{F0: 1000}
                                                    _ = v29
                                                }
                                                if (-91 <= 32767) {
                                                    if v8 {
                                                        p1 = (p1 + p1)
                                                        v30 := GenS4{F0: 1, F1: int64(fi10)}
                                                        switch p1 {
                                                        case "hello":
                                                            var v31 int16 = -795
                                                            var v32 uint8 = 126
                                                            _ = v31
                                                            _ = v32
                                                        case "a":
                                                            var v33 bool = v8
                                                            _ = v33
                                                        }
                                                        var v34 string = "foo bar"
                                                        if len(v7) > 0 {
                                                            v7[0] = (fi13 | (fi10 | fi13))
                                                        }
                                                        _ = v30
                                                        _ = v34
                                                    }
                                                }
                                                v35 := GenS3{F0: 1000}
                                                v36 := GenS4{F0: 18446744073709551615, F1: 42}
                                                switch p1 {
                                                case "hello":
                                                    fmt.Println(p1)
                                                case "\n\t":
                                                    if (!(1 <= 2)) {
                                                        v7 = append(v7, (fi13 >> ((v9) & 63)))
                                                        if false {
                                                            v25 -= 1
                                                        } else {
                                                            var v37 int64 = 9223372036854775807
                                                            v37 = 9223372036854775807
                                                            v9 += 942
                                                            _ = v37
                                                        }
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                    }
                                                default:
                                                    var v38 int16 = int16(fi13)
                                                    _ = v38
                                                }
                                                fmt.Println(p1, p0, p1)
                                                if (v9 > 18446744073709551615) {
                                                    var v39 string = p1
                                                    var v40 int64 = p0
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                    switch (fi10 >> (((^(v9))) & 63)) {
                                                    case 964:
                                                        var v41 uint32 = 923
                                                        _ = v41
                                                    case 128:
                                                        p1 = v39
                                                    case 9223372036854775807:
                                                        var v42 int8 = -79
                                                        _ = v42
                                                    default:
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                        v43 := GenS4{F0: 464, F1: (v40 & (v40 >> ((42) & 63)))}
                                                        if v8 {
                                                            if v8 {
                                                                if (!(559 <= 2)) {
                                                                    v44 := GenS5{F0: 0}
                                                                    switch 0 {
                                                                    case -833:
                                                                        var v45 float32 = 0.0
                                                                        _ = v45
                                                                    default:
                                                                        fmt.Println(p0, v25)
                                                                    }
                                                                    _ = v44
                                                                }
                                                                switch (fi13 << ((v9) & 63)) {
                                                                case 536:
                                                                    var v46 float64 = 1.5
                                                                    _ = v46
                                                                }
                                                            }
                                                            if len(v7) > 0 {
                                                                v7[0] = (fi13 | fi10)
                                                            }
                                                            var v47 int = (fi13 % (((fi13 - (fi10 / ((0) | 1)))) | 1))
                                                            _ = v47
                                                        }
                                                        _ = v43
                                                    }
                                                    _ = v39
                                                    _ = v40
                                                }
                                                fmt.Println(v25)
                                                _ = v35
                                                _ = v36
                                            } else {
                                                v25 = int32(v9)
                                                switch (p1 + (p1 + p1)) {
                                                case "a":
                                                    var v48 float64 = -2.25
                                                    _ = v48
                                                case "\n\t":
                                                    if v8 {
                                                        switch p1 {
                                                        case "a":
                                                            var v49 int16 = -969
                                                            _ = v49
                                                        default:
                                                            v7 = append(v7, 4294967295)
                                                        }
                                                    } else {
                                                        v50 := GenS3{F0: int16(fi10)}
                                                        switch fi10 {
                                                        case -9223372036854775808:
                                                            var v51 int8 = 32
                                                            _ = v51
                                                        default:
                                                            p0 = (p0 * p0)
                                                        }
                                                        var v52 uint16 = 2
                                                        _ = v50
                                                        _ = v52
                                                    }
                                                case "foo bar":
                                                    _ = len(v7)
                                                }
                                            }
                                        }
                                        _ = len(v7)
                                        var v53 int = fi13
                                        _ = v53
                                    } else {
                                        if (1e10 >= (0.0 / 1e10)) {
                                            v54 := GenS4{F0: 51, F1: p0}
                                            p1 = "foo bar"
                                            v9 = 0
                                            v8 = (0.0 <= (1e10 / 1.0))
                                            _ = v54
                                        } else {
                                            v7 = append(v7, fi10)
                                        }
                                        switch fi10 {
                                        case 1:
                                            p1 = (p1 + ((p1 + "a") + p1))
                                        case 32768:
                                            v8 = v8
                                        case 127:
                                            if (true && false) {
                                                v25 = (^(v25))
                                                _ = len(v7)
                                                v55 := GenS3{F0: -32768}
                                                v25 = (-(v25))
                                                _ = len(v7)
                                                _ = v55
                                            }
                                        }
                                    }
                                    v56 := GenS6{F0: 65535}
                                    v57 := GenS2{F0: 486}
                                    var v58 float64 = 1.5
                                    v7 = append(v7, (fi10 + (fi13 / ((fi10) | 1))))
                                    v25 -= v25
                                    if ((v8 && ((v25 <= int32(p0)) && (240 != 144))) || v8) {
                                        var v59 uint32 = 2147483647
                                        _ = v59
                                    } else {
                                        v60 := GenS7{F0: (^(fi13)), F1: p1}
                                        v61 := GenS8{F0: (p0 - (p0 >> (((v9 - (v9 << (((v9 ^ v9)) & 63)))) & 63)))}
                                        _ = v60
                                        _ = v61
                                    }
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                } else {
                                    var v62 bool = (v8 || v8)
                                    _ = v62
                                }
                                v63 := GenS2{F0: -702}
                                p0 -= -558
                                v64 := GenS8{F0: p0}
                                _ = v25
                                _ = v63
                                _ = v64
                            }
                            if (!v8) {
                                v9 = 2
                                var v65 string = p1
                                _ = v65
                            }
                        case "a":
                            var v66 uint = v9
                            _ = v66
                        }
                    }
                } else {
                    p0 -= p0
                }
                _ = fi13
            }
            v8 = true
        }
        v67 := GenS9{F0: v9}
        _ = fi10
        _ = v67
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int32, p1 uint64, p2 int16) string {
    p2 -= p2
    var v68 string = "\n\t"
    for fi69 := 0; fi69 < 4; fi69++ {
        v70 := []int{}
        v71 := GenS6{F0: 2}
        _ = fi69
        _ = v70
        _ = v71
    }
    var v72 int32 = (p0 & (p0 ^ p0))
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v68
    _ = v72
    return ((("" + "a") + ("" + "hello")) + v68)
}

func main() {
    var v1 float32 = 0.1
    var v2 int16 = -32768
    var v3 float64 = 0.0
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    v6 := []int{4294967295, 32768, 42}
    fmt.Println(v3, v3, v1)
    for fi73 := 0; fi73 < 7; fi73++ {
        v6 = append(v6, fi73)
        v74 := GenS4{F0: 169, F1: 0}
        v3 += 3.14159
        _ = fi73
        _ = v74
    }
    v6 = append(v6, 32768)
    fmt.Println(v2, v1)
    v1 += ((v1 / v1) + (v1 - (v1 * (v1 - (((((v1 * float32(3.14159)) - v1) - v1) + v1) - float32(((0.1 + v3) - (float64(v1) - float64(v1)))))))))
    if false {
        switch -285 {
        case -565:
            v75 := GenS6{F0: 65535}
            _ = v75
        default:
            v6 = append(v6, -9223372036854775808)
        }
        for fi76 := 0; fi76 < 16; fi76++ {
            var v77 uint = 18446744073709551615
            if (true || (false || ((1 > 1) || false))) {
                if ("" != "\n\t") {
                    if true {
                        var v78 uint32 = 65535
                        switch (fi76 ^ (fi76 - (fi76 - (fi76 + -1)))) {
                        case 4294967295:
                            var v79 float32 = ((v1 - (float32(-2.25) * float32((v3 + v3)))) * float32(float64(-2.25)))
                            _ = v79
                        case -9223372036854775808:
                            switch (fi76 >> (((^(v77))) & 63)) {
                            case -2147483648:
                                for fi80 := 0; fi80 < 14; fi80++ {
                                    if (!true) {
                                        if false {
                                            v81 := GenS2{F0: (v2 >> ((v77) & 63))}
                                            v3 = (v3 + (v3 - v3))
                                            _ = v81
                                        } else {
                                            v3 = (((((1e10 / v3) + float64(v1)) * v3) + (1.5 * v3)) * v3)
                                            var v82 uint = 18446744073709551615
                                            _ = v82
                                        }
                                        v83 := GenS7{F0: fi76, F1: ("" + "foo bar")}
                                        var v84 float64 = v3
                                        var v85 uint8 = 6
                                        _ = len(v4)
                                        var v86 float64 = 0.1
                                        v86 = v3
                                        var v87 int8 = 127
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                    } else {
                                        var v88 float32 = 1e10
                                        v78 += (v78 >> ((0) & 63))
                                        var v89 uint64 = 2
                                        switch "foo bar" {
                                        case "":
                                            var v90 int16 = (v2 / (((v2 / (((v2 + (v2 ^ 0))) | 1))) | 1))
                                            _ = v90
                                        case "\n\t":
                                            v91 := v5[fi80]
                                            _ = v91
                                        case "a":
                                            v3 = v3
                                        default:
                                            v92 := GenS10{F0: 1, F1: 351}
                                            var v93 int32 = -2147483648
                                            _ = v92
                                            _ = v93
                                        }
                                        var v94 uint16 = 310
                                        _ = v88
                                        _ = v89
                                        _ = v94
                                    }
                                    if true {
                                        fmt.Println(v2, v78)
                                        v78 += 0
                                    }
                                    _ = fi80
                                }
                            case 32768:
                                v95 := GenS11{F0: (127 < 1), F1: -128}
                                _ = v95
                            case 4294967295:
                                switch fi76 {
                                case 2147483647:
                                    v96 := GenS12{F0: genFn0(0, "hello")}
                                    _ = v96
                                }
                            default:
                                for fi97 := 0; fi97 < 17; fi97++ {
                                    var v98 uint32 = (v78 ^ 918)
                                    fmt.Println(v78, v98)
                                    if (!false) {
                                        switch -777 {
                                        case -2147483648:
                                            if (v77 <= (v77 << (((^(v77))) & 63))) {
                                                if (true && true) {
                                                    switch -307 {
                                                    case -1:
                                                        var v99 uint16 = 765
                                                        _ = v99
                                                    case 0:
                                                        var v100 float64 = float64(0.1)
                                                        _ = v100
                                                    default:
                                                        v6 = append(v6, fi97)
                                                    }
                                                    v2 = 0
                                                }
                                            } else {
                                                var v101 uint8 = 255
                                                v78 -= v98
                                                var v102 float64 = v3
                                                v102 = v3
                                                _ = v101
                                                _ = v102
                                            }
                                        }
                                        _ = len(v6)
                                        fmt.Println(v1, v77)
                                    } else {
                                        if true {
                                            if false {
                                                v4["foo bar"] = (-(fi97))
                                                v3 -= (1e10 - 1e10)
                                                fmt.Println(v98, v77)
                                            }
                                            if false {
                                                var v103 bool = (2 >= 127)
                                                _ = v103
                                            }
                                        }
                                        delete(v4, "hello")
                                    }
                                    _ = fi97
                                    _ = v98
                                }
                            }
                        default:
                            for fi104 := 0; fi104 < 16; fi104++ {
                                switch "\n\t" {
                                case "":
                                    switch (fi76 >> (((v77 | 4294967295)) & 63)) {
                                    case 42:
                                        v105 := GenS13{F0: (v2 & genFn0(-1, "a")), F1: (("hello" + ("foo bar" + ("hello" + ((("hello" + "hello") + "") + "foo bar")))) + "hello"), F2: ((("" + (((("hello" + "foo bar") + "a") + "\n\t") + "a")) + (("" + "hello") + "a")) + "")}
                                        _ = v105
                                    case 32767:
                                        var v106 uint16 = 825
                                        _ = v106
                                    case 32768:
                                        if true {
                                            v107 := GenS14{F0: 255}
                                            _ = len(v5)
                                            if (!((!(v78 <= (^(v78)))) && true)) {
                                                var v108 string = ""
                                                _ = v108
                                            }
                                            _ = v107
                                        }
                                    }
                                case "a":
                                    var v109 string = "a"
                                    _ = v109
                                }
                                fmt.Println(v2, v3, v77)
                                var v110 uint = (v77 ^ 528)
                                var v111 int = 4294967295
                                v111 = (fi104 & (v111 >> ((v110) & 63)))
                                switch fi104 {
                                case 42:
                                    v4[("" + (("" + ("a" + "a")) + ""))] = 32767
                                case 32767:
                                    var v112 uint = 2
                                    _ = v112
                                default:
                                    v113 := v5[fi76]
                                    _ = v113
                                }
                                _ = fi104
                                _ = v110
                                _ = v111
                            }
                        }
                        _ = v78
                    } else {
                        for fi114 := 0; fi114 < 14; fi114++ {
                            if true {
                                v115 := v4[("hello" + "\n\t")]
                                if len(v6) > 0 {
                                    v6[0] = (fi76 ^ (fi76 % ((v115) | 1)))
                                }
                                _ = v115
                            } else {
                                v3 += (3.14159 + v3)
                                v116 := v5[-1]
                                var v117 uint16 = 75
                                _ = v116
                                _ = v117
                            }
                            _ = len(v4)
                            _ = fi114
                        }
                    }
                    var v118 bool = ((^(fi76)) <= (fi76 % ((127) | 1)))
                    _ = v118
                } else {
                    v3 += 3.14159
                    v77 -= v77
                }
                var v119 int64 = -386
                if (uint8(v77) == 255) {
                    switch "foo bar" {
                    case "hello":
                        var v120 int = fi76
                        _ = v120
                    }
                    v121 := GenS15{F0: 258, F1: v119}
                    var v122 int8 = 0
                    _ = v121
                    _ = v122
                } else {
                    var v123 uint32 = 1
                    if (!true) {
                        for fi124 := 0; fi124 < 11; fi124++ {
                            if (true || true) {
                                var v125 string = "hello"
                                v126 := GenS16{F0: (fi124 * 32768), F1: genFn0((v119 + 0), (v125 + "\n\t"))}
                                fmt.Println(v119, v119, fi76)
                                _ = v125
                                _ = v126
                            } else {
                                v3 += 1e10
                                var v127 uint32 = 233
                                if false {
                                    var v128 uint64 = 18446744073709551615
                                    _ = v128
                                }
                                v123 = v127
                                _ = v127
                            }
                            v119 += v119
                            fmt.Println(v1, v3)
                            _ = fi124
                        }
                        v129 := GenS16{F0: 4294967295, F1: (v2 | genFn0((v119 & 0), "\n\t"))}
                        _ = v129
                    }
                    _ = v123
                }
                _ = v119
            } else {
                var v130 float32 = ((0.1 * 3.14159) - 3.14159)
                switch "a" {
                case "":
                    v4[""] = (^(fi76))
                case "foo bar":
                    fmt.Println(v77, v3, v77)
                case "\n\t":
                    if true {
                        var v131 uint8 = 1
                        _ = v131
                    }
                default:
                    if true {
                        v2 -= genFn0(-640, ((("hello" + ("" + ("\n\t" + "a"))) + "foo bar") + "a"))
                        var v132 int64 = 9223372036854775807
                        if true {
                            v130 -= ((v130 - 0.0) * v1)
                            v133 := GenS10{F0: 192, F1: 261}
                            v2 = v2
                            _ = v133
                        }
                        _ = v132
                    }
                }
                _ = v130
            }
            if (623 != 2147483648) {
                var v134 uint8 = 189
                if (false || false) {
                    fmt.Println(fi76, fi76, v1)
                    delete(v4, ("hello" + "foo bar"))
                }
                _ = v134
            } else {
                switch fi76 {
                case -1:
                    fmt.Println(v77, v2, v77)
                case 1:
                    v2 += v2
                default:
                    switch "" {
                    case "a":
                        v77 = (^(v77))
                    case "foo bar":
                        v135 := GenS17{F0: 42, F1: 0, F2: -128}
                        _ = v135
                    default:
                        if (false || (0.1 != 1e10)) {
                            var v136 string = "foo bar"
                            _ = v136
                        }
                    }
                }
                var v137 float32 = v1
                _ = v137
            }
            _ = fi76
            _ = v77
        }
    }
    fmt.Println(v3)
    var v138 int64 = 0
    var v139 uint = 65535
    delete(v4, "a")
    switch (("\n\t" + "") + "\n\t") {
    case "\n\t":
        for fi140 := 0; fi140 < 14; fi140++ {
            if (42 <= (fi140 << (((v139 / (((^(v139))) | 1))) & 63))) {
                v2 = (v2 ^ v2)
                switch fi140 {
                case 422:
                    v1 = ((0.0 / float32((v3 * (v3 / (((float64(1.5) / float64(((v1 * v1) * ((3.14159 - v1) + 1.5)))) * (1e10 - 3.14159)) - 0.0))))) - 0.0)
                case 0:
                    if (4294967295 > 24) {
                        v141 := GenS18{F0: 149, F1: 133, F2: 1}
                        var v142 int32 = 653
                        _ = v141
                        _ = v142
                    } else {
                        var v143 int16 = int16(v139)
                        switch (("a" + (("hello" + "foo bar") + ("" + (("hello" + (("" + "\n\t") + "hello")) + ((("" + ("\n\t" + "\n\t")) + "foo bar") + "hello"))))) + "hello") {
                        case "\n\t":
                            if false {
                                switch "" {
                                case "foo bar":
                                    v144 := GenS12{F0: 989}
                                    _ = v144
                                case "a":
                                    var v145 float64 = v3
                                    _ = v145
                                case "\n\t":
                                    if len(v6) > 0 {
                                        v6[0] = fi140
                                    }
                                }
                                v146 := GenS19{F0: 32767, F1: 114}
                                _ = v146
                            }
                        case "hello":
                            if true {
                                fmt.Println(v2)
                                v2 = v2
                                if false {
                                    fmt.Println(fi140, v138, v139)
                                    if true {
                                        for fi147 := 0; fi147 < 7; fi147++ {
                                            if false {
                                                v143 = 1000
                                                v138 -= -492
                                            }
                                            v148 := GenS18{F0: 1, F1: 231, F2: 640}
                                            if false {
                                                if (false || false) {
                                                    v138 -= (^(v138))
                                                    fmt.Println(v1, fi147)
                                                    fmt.Println(v138)
                                                }
                                                if false {
                                                    v149 := GenS17{F0: 42, F1: v138, F2: v138}
                                                    v6 = append(v6, (fi147 / ((fi140) | 1)))
                                                    if true {
                                                        v143 += (v143 % ((-1) | 1))
                                                        switch (fi147 % (((fi140 | 510)) | 1)) {
                                                        case 42:
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v150 uint = v139
                                                                _ = v150
                                                            }
                                                        case 2147483647:
                                                            if (((!false) && (v3 <= float64(0.0))) && false) {
                                                                v143 -= (v143 >> ((42) & 63))
                                                            }
                                                        }
                                                    } else {
                                                        _ = len(v4)
                                                        fmt.Println(v138, fi147)
                                                        var v151 float32 = (((3.14159 * 3.14159) - 0.0) - v1)
                                                        _ = v151
                                                    }
                                                    if true {
                                                        var v152 bool = true
                                                        _ = v152
                                                    }
                                                    _ = v149
                                                }
                                            } else {
                                                v138 -= (v138 >> ((v139) & 63))
                                                var v153 uint32 = 2147483647
                                                _ = v153
                                            }
                                            fmt.Println(v1, v2, v143)
                                            fmt.Println(v3, v138, fi140)
                                            _ = fi147
                                            _ = v148
                                        }
                                        v154 := GenS3{F0: (-(v2))}
                                        _ = v154
                                    }
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                            } else {
                                var v155 bool = false
                                v156 := GenS20{F0: 4294967295, F1: 1}
                                _ = v155
                                _ = v156
                            }
                        case "foo bar":
                            if (!false) {
                                if true {
                                    var v157 float64 = (3.14159 * (((v3 - 1.5) + ((float64(v1) / -2.25) * (((v3 * v3) - v3) / 1.0))) + v3))
                                    for fi158 := 0; fi158 < 10; fi158++ {
                                        fmt.Println(v138, v138, v2)
                                        _ = fi158
                                    }
                                    _ = v157
                                }
                                fmt.Println(v138)
                            } else {
                                v159 := GenS4{F0: 196, F1: v138}
                                _ = v159
                            }
                        }
                        v160 := GenS13{F0: genFn0(v138, "a"), F1: "", F2: "\n\t"}
                        _ = v143
                        _ = v160
                    }
                case 728:
                    fmt.Println(fi140, v2, fi140)
                }
            }
            _ = fi140
        }
    case "hello":
        switch "foo bar" {
        case "a":
            fmt.Println(v1)
        }
    case "foo bar":
        for fi161 := 0; fi161 < 6; fi161++ {
            v162 := GenS21{F0: (v2 >> (((^(v139))) & 63))}
            _ = fi161
            _ = v162
        }
    default:
        v138 -= 9223372036854775807
    }
    var v163 int64 = 42
    var v164 string = (("foo bar" + (("\n\t" + ("a" + "\n\t")) + ("hello" + ("" + "hello")))) + "a")
    v165 := GenS17{F0: 255, F1: v163, F2: (v138 ^ v138)}
    var v166 int = int(v139)
    var v167 string = (v164 + v164)
    fmt.Println("v1:", v1)
    fmt.Println("v138:", v138)
    fmt.Println("v139:", v139)
    fmt.Println("v163:", v163)
    fmt.Println("v164:", v164)
    fmt.Println("v166:", v166)
    fmt.Println("v167:", v167)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v138
    _ = v139
    _ = v163
    _ = v164
    _ = v165
    _ = v166
    _ = v167
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
