package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 int8; F1 uint; F2 bool }
type GenS9 struct { F0 uint32; F1 uint8; F2 string }
type GenS10 struct { F0 uint8; F1 int8; F2 uint }
type GenS11 struct { F0 int16; F1 bool; F2 int8 }
type GenS12 struct { F0 int16; F1 int8 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int32; F1 uint32 }
type GenS17 struct { F0 uint; F1 uint16; F2 int }
type GenS18 struct { F0 uint8; F1 uint32 }
func genFn0() bool {
    v6 := GenS1{F0: true}
    var v7 bool = false
    _ = v6
    _ = v7
    return ((float64(-2.25) + 0.0) <= (1e10 * float64(1e10)))
}
func genFn1(p0 int8, p1 uint16, p2 bool) int8 {
    switch 42 {
    case 2147483647:
        p2 = (2147483647 >= -112)
    case -2147483648:
        var v8 uint = 741
        _ = v8
    }
    switch 723 {
    case 42:
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint8 = 1
            _ = fi9
            _ = v10
        }
    case 127:
        p0 = -27
    default:
        if p2 {
            v11 := GenS2{F0: p0}
            v12 := []int{-1, 9223372036854775807, 42}
            switch 731 {
            case -288:
                var v13 uint32 = 4294967295
                _ = v13
            case -1:
                p1 = p1
            case 127:
                fmt.Println(p2, p1)
            default:
                if (0 >= 753) {
                    fmt.Println(p1, p0)
                } else {
                    var v14 int = 259
                    var v15 string = "hello"
                    _ = v14
                    _ = v15
                }
            }
            _ = v11
            _ = v12
        } else {
            for fi16 := 0; fi16 < 19; fi16++ {
                fmt.Println(p0)
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 string = "a"
                    v19 := GenS3{F0: 2}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v20 := GenS4{F0: fi16, F1: -9223372036854775808}
                _ = fi16
                _ = v20
            }
            var v21 int32 = -1
            _ = v21
        }
    }
    if p2 {
        switch 2147483647 {
        case -270:
            var v22 uint64 = 2
            _ = v22
        case 964:
            var v23 string = ("a" + "foo bar")
            _ = v23
        default:
            if p2 {
                p1 = p1
            } else {
                v24 := GenS5{F0: 1}
                if ((!false) || false) {
                    for fi25 := 0; fi25 < 12; fi25++ {
                        for fi26 := 0; fi26 < 15; fi26++ {
                            v27 := []int{-9223372036854775808, (^(fi25)), fi26, 329}
                            _ = fi26
                            _ = v27
                        }
                        if (p2 || true) {
                            v28 := GenS3{F0: 824}
                            if (42 != 42) {
                                switch "\n\t" {
                                case "hello":
                                    var v29 int32 = 42
                                    _ = v29
                                }
                                if (p1 < p1) {
                                    v30 := []int{fi25}
                                    switch (("hello" + "") + "") {
                                    case "hello":
                                        for fi31 := 0; fi31 < 14; fi31++ {
                                            v32 := GenS2{F0: p0}
                                            _ = fi31
                                            _ = v32
                                        }
                                    case "foo bar":
                                        v30 = append(v30, fi25)
                                    case "":
                                        fmt.Println(p1)
                                    }
                                    _ = v30
                                }
                            }
                            _ = v28
                        } else {
                            var v33 int32 = -804
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 16; fi34++ {
                            fmt.Println(fi25, p1)
                            var v35 uint8 = 0
                            switch fi25 {
                            case 0:
                                if p2 {
                                    v36 := GenS6{F0: p1, F1: (("\n\t" + ("foo bar" + "foo bar")) + "foo bar")}
                                    p0 -= (^(p0))
                                    switch "\n\t" {
                                    case "hello":
                                        p2 = ((^(v35)) > 0)
                                    case "\n\t":
                                        p1 = p1
                                        fmt.Println(fi25, v35)
                                    case "a":
                                        if (878 != 2147483648) {
                                            v37 := GenS7{F0: "hello", F1: 0}
                                            var v38 string = ("\n\t" + "a")
                                            v35 = (v35 - v35)
                                            v39 := GenS5{F0: 0}
                                            var v40 uint8 = (v35 << ((255) & 63))
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        }
                                    }
                                    _ = v36
                                } else {
                                    var v41 uint16 = p1
                                    var v42 string = "hello"
                                    _ = v41
                                    _ = v42
                                }
                            case 127:
                                var v43 int64 = 0
                                _ = v43
                            case 2147483647:
                                v44 := []int{32768}
                                _ = v44
                            default:
                                var v45 int64 = -1
                                _ = v45
                            }
                            _ = fi34
                            _ = v35
                        }
                        _ = fi25
                    }
                }
                var v46 uint32 = 4294967295
                _ = v24
                _ = v46
            }
        }
        v47 := GenS1{F0: true}
        for fi48 := 0; fi48 < 19; fi48++ {
            for fi49 := 0; fi49 < 9; fi49++ {
                v50 := []int{}
                var v51 int16 = -32768
                var v52 uint32 = 2147483647
                var v53 int32 = 1
                var v54 int64 = 0
                _ = fi49
                _ = v50
                _ = v51
                _ = v52
                _ = v53
                _ = v54
            }
            _ = fi48
        }
        _ = v47
    } else {
        var v55 uint8 = 127
        _ = v55
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 216
    var v3 int = 2147483647
    var v4 int8 = 127
    var v5 uint = 255
    var v56 int16 = -300
    switch "foo bar" {
    case "foo bar":
        v57 := v1[700]
        _ = v57
    default:
        if false {
            fmt.Println(v5)
            v2 = v56
        } else {
            v56 -= v56
        }
    }
    delete(v1, 465)
    v58 := GenS5{F0: 42}
    v5 -= 42
    v4 = genFn1((v4 + (v4 % ((0) | 1))), 255, false)
    switch (v3 | v3) {
    case -9223372036854775808:
        var v59 int16 = v2
        _ = v59
    case 88:
        v5 += (v5 ^ (v5 % ((v5) | 1)))
    default:
        v60 := GenS5{F0: 42}
        _ = v60
    }
    v61 := []int{}
    if (255 >= 0) {
        if false {
            if true {
                for fi62 := 0; fi62 < 15; fi62++ {
                    fmt.Println(v2, v2, v4)
                    fmt.Println(v2, v2, v3)
                    _ = fi62
                }
                switch (^(v3)) {
                case 32768:
                    var v63 int8 = v4
                    _ = v63
                case 2147483647:
                    fmt.Println(v2, v4, v4)
                case 127:
                    var v64 int = v3
                    _ = v64
                default:
                    if (-2147483648 == 0) {
                        var v65 int32 = -56
                        _ = v65
                    }
                }
            } else {
                switch (v3 + v3) {
                case 4294967295:
                    for fi66 := 0; fi66 < 13; fi66++ {
                        var v67 string = "hello"
                        var v68 bool = (true && (0 < -129))
                        for fi69 := 0; fi69 < 6; fi69++ {
                            if (v68 && v68) {
                                if v68 {
                                    var v70 float64 = (3.14159 / -2.25)
                                    var v71 uint = v5
                                    v1[2147483647] = ((v4 ^ (v4 - v4)) >= genFn1((v4 | 42), 1, true))
                                    _ = v70
                                    _ = v71
                                }
                                fmt.Println(fi66, v2)
                            } else {
                                if v68 {
                                    if (61 < 134) {
                                        _ = len(v61)
                                        v56 -= (v2 & v2)
                                        var v72 float32 = -2.25
                                        switch v67 {
                                        case "hello":
                                            v73 := GenS8{F0: 100, F1: v5, F2: true}
                                            _ = v73
                                        }
                                        var v74 string = v67
                                        _ = v72
                                        _ = v74
                                    }
                                } else {
                                    v75 := GenS9{F0: 0, F1: 182, F2: v67}
                                    if false {
                                        switch v3 {
                                        case -935:
                                            v4 = (v4 | (v4 * -47))
                                        case -2147483648:
                                            if (!true) {
                                                var v76 bool = (!(!v68))
                                                _ = v76
                                            }
                                        case 0:
                                            switch (v67 + (v67 + v67)) {
                                            case "hello":
                                                if (!true) {
                                                    fmt.Println(v67)
                                                } else {
                                                    var v77 uint16 = 481
                                                    _ = v77
                                                }
                                            case "\n\t":
                                                fmt.Println(fi69, v4, v2)
                                            default:
                                                fmt.Println(v5, v4, v4)
                                                v67 = v67
                                                v78 := GenS10{F0: 1, F1: -128, F2: v5}
                                                _ = v78
                                            }
                                        }
                                        v61 = append(v61, fi66)
                                        if v68 {
                                            v56 -= v2
                                            v2 -= 0
                                            switch "a" {
                                            case "hello":
                                                _ = len(v1)
                                                fmt.Println(fi66)
                                            case "":
                                                v5 = (v5 << (((v5 % ((v5) | 1))) & 63))
                                            }
                                        }
                                    } else {
                                        v79 := GenS6{F0: 719, F1: v67}
                                        var v80 float32 = 0.1
                                        _ = v79
                                        _ = v80
                                    }
                                    _ = v75
                                }
                                if (65535 > 0) {
                                    var v81 uint = v5
                                    var v82 uint = 42
                                    _ = v81
                                    _ = v82
                                }
                            }
                            if true {
                                v67 = "a"
                            } else {
                                if false {
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(v4, v56)
                                    case "foo bar":
                                        switch v3 {
                                        case 0:
                                            fmt.Println(fi66)
                                        case 9223372036854775807:
                                            v61 = append(v61, (fi69 / (((v3 * fi66)) | 1)))
                                        }
                                    case "a":
                                        v3 += (v3 + (^(fi69)))
                                    }
                                    var v83 uint8 = 131
                                    var v84 uint16 = 342
                                    _ = v83
                                    _ = v84
                                }
                            }
                            _ = fi69
                        }
                        _ = fi66
                        _ = v67
                        _ = v68
                    }
                default:
                    if false {
                        var v85 uint16 = 97
                        _ = v85
                    } else {
                        for fi86 := 0; fi86 < 15; fi86++ {
                            fmt.Println(v5, v3, v2)
                            v3 = 127
                            var v87 uint32 = 1
                            _ = fi86
                            _ = v87
                        }
                        v5 += (^(v5))
                        if false {
                            v3 = v3
                            v88 := GenS3{F0: v5}
                            v89 := GenS9{F0: 0, F1: 42, F2: ("foo bar" + "a")}
                            _ = v88
                            _ = v89
                        } else {
                            if true {
                                v90 := GenS7{F0: "hello", F1: 18446744073709551615}
                                for fi91 := 0; fi91 < 2; fi91++ {
                                    for fi92 := 0; fi92 < 7; fi92++ {
                                        if false {
                                            if false {
                                                v93 := GenS11{F0: v2, F1: true, F2: (v4 >> (((v5 * (v5 & v5))) & 63))}
                                                _ = v93
                                            }
                                            if false {
                                                _ = len(v1)
                                                if true {
                                                    if false {
                                                        var v94 string = (("\n\t" + "") + "hello")
                                                        v5 -= v5
                                                        v2 = v56
                                                        switch fi92 {
                                                        case 127:
                                                            var v95 string = v94
                                                            _ = v95
                                                        default:
                                                            if false {
                                                                var v96 int = (v3 >> ((v5) & 63))
                                                                fmt.Println(v5, v56)
                                                                _ = v96
                                                            } else {
                                                                var v97 uint16 = 255
                                                                v4 += (-(v4))
                                                                v5 = (^(v5))
                                                                _ = v97
                                                            }
                                                        }
                                                        var v98 int16 = -32768
                                                        fmt.Println(v98, fi91, v94)
                                                        _ = v94
                                                        _ = v98
                                                    } else {
                                                        v99 := GenS12{F0: v56, F1: (v4 % ((v4) | 1))}
                                                        _ = v99
                                                    }
                                                    switch ("" + ("" + "hello")) {
                                                    case "hello":
                                                        _ = len(v1)
                                                    case "":
                                                        fmt.Println(v5, v56)
                                                    default:
                                                        var v100 int8 = (^(v4))
                                                        _ = v100
                                                    }
                                                    v56 += (v56 >> ((97) & 63))
                                                    v101 := GenS3{F0: v5}
                                                    _ = v101
                                                } else {
                                                    var v102 uint16 = 97
                                                    if len(v61) > 0 {
                                                        v61[0] = 127
                                                    }
                                                    v4 = 42
                                                    _ = v102
                                                }
                                            } else {
                                                _ = len(v61)
                                            }
                                            v103 := GenS13{F0: 1}
                                            var v104 bool = (!false)
                                            var v105 int = 128
                                            _ = v103
                                            _ = v104
                                            _ = v105
                                        } else {
                                            fmt.Println(v4, v4, v2)
                                        }
                                        v106 := GenS14{F0: 128}
                                        var v107 int8 = int8(fi91)
                                        _ = fi92
                                        _ = v106
                                        _ = v107
                                    }
                                    fmt.Println(v4)
                                    v4 = (v4 ^ v4)
                                    _ = fi91
                                }
                                v56 += v2
                                _ = v90
                            } else {
                                v4 -= (v4 * (v4 * (v4 / ((v4) | 1))))
                                v1[2147483647] = (v3 <= (v3 + v3))
                            }
                            if (false || true) {
                                v2 = 265
                                v108 := GenS4{F0: v3, F1: v3}
                                var v109 uint = (v5 / ((2) | 1))
                                _ = v108
                                _ = v109
                            }
                            var v110 int8 = -108
                            _ = v110
                        }
                    }
                }
                v61 = append(v61, (v3 / ((v3) | 1)))
                switch "\n\t" {
                case "foo bar":
                    switch "" {
                    case "hello":
                        for fi111 := 0; fi111 < 15; fi111++ {
                            var v112 float64 = 1e10
                            _ = len(v1)
                            for fi113 := 0; fi113 < 15; fi113++ {
                                var v114 int16 = (v56 % ((v2) | 1))
                                switch "foo bar" {
                                case "foo bar":
                                    if len(v61) > 0 {
                                        v61[0] = int(v4)
                                    }
                                case "hello":
                                    if ((v114 ^ v114) == (v2 | (v114 + (v56 - v2)))) {
                                        if (-481 == 9223372036854775807) {
                                            if ("" <= "") {
                                                if ((0 <= 0) && false) {
                                                    if ((v4 != v4) || ((float64(0.0) - (v112 + (float64(float32((v112 - (v112 / 1.0)))) - (float64(1e10) - v112)))) == v112)) {
                                                        if true {
                                                            if (true || true) {
                                                                fmt.Println(v4)
                                                                if true {
                                                                    v5 += v5
                                                                    v112 += v112
                                                                    var v115 int32 = -1
                                                                    var v116 int8 = v4
                                                                    if (0.1 > 1e10) {
                                                                        if (fi113 < 42) {
                                                                            v117 := v1[935]
                                                                            fmt.Println(fi111, v4, v2)
                                                                            _ = len(v1)
                                                                            _ = v117
                                                                        }
                                                                        v118 := GenS15{F0: (fi111 | (fi111 * (fi113 / ((265) | 1))))}
                                                                        var v119 uint16 = 862
                                                                        _ = v118
                                                                        _ = v119
                                                                    }
                                                                    _ = v115
                                                                    _ = v116
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v56 = 1000
                                                } else {
                                                    if (4294967295 <= 0) {
                                                        var v120 int = (fi111 * 848)
                                                        if (-128 != (v4 % ((genFn1(v4, 0, false)) | 1))) {
                                                            var v121 int64 = -1
                                                            _ = len(v1)
                                                            v122 := GenS3{F0: 177}
                                                            _ = v121
                                                            _ = v122
                                                        }
                                                        switch (v3 >> ((65535) & 63)) {
                                                        case 1:
                                                            delete(v1, 0)
                                                        case -359:
                                                            var v123 int8 = -128
                                                            _ = v123
                                                        }
                                                        v120 = fi113
                                                        var v124 uint32 = 0
                                                        var v125 uint16 = 1
                                                        switch fi111 {
                                                        case 128:
                                                            v126 := GenS1{F0: true}
                                                            _ = v126
                                                        default:
                                                            v56 += (v114 - (v114 << ((42) & 63)))
                                                        }
                                                        if true {
                                                            fmt.Println(v114, v2)
                                                            fmt.Println(v125)
                                                            switch (-(v120)) {
                                                            case 32768:
                                                                delete(v1, 4294967295)
                                                            default:
                                                                v114 = v2
                                                            }
                                                        }
                                                        _ = v120
                                                        _ = v124
                                                        _ = v125
                                                    }
                                                    var v127 int16 = (v114 & v56)
                                                    _ = v127
                                                }
                                                var v128 int = (v3 * fi113)
                                                v3 = -2147483648
                                                _ = v128
                                            }
                                        } else {
                                            v1[0] = false
                                        }
                                        v3 = -125
                                    }
                                default:
                                    if (128 == int(v56)) {
                                        fmt.Println(v2, v3)
                                    } else {
                                        switch 127 {
                                        case 0:
                                            var v129 int64 = int64(v4)
                                            _ = v129
                                        case 127:
                                            v130 := GenS16{F0: -879, F1: 439}
                                            _ = v130
                                        default:
                                            var v131 string = "\n\t"
                                            _ = v131
                                        }
                                        v132 := GenS17{F0: 914, F1: 737, F2: 42}
                                        _ = v132
                                    }
                                }
                                var v133 int32 = 0
                                fmt.Println(v2, v133, v4)
                                _ = fi113
                                _ = v114
                                _ = v133
                            }
                            _ = fi111
                            _ = v112
                        }
                    case "":
                        var v134 int16 = (v2 >> (((^(v5))) & 63))
                        _ = v134
                    default:
                        delete(v1, 0)
                    }
                default:
                    v56 -= v56
                }
            }
        } else {
            v5 = (v5 ^ 2)
            fmt.Println(v3, v4, v4)
        }
        v135 := GenS18{F0: 2, F1: 42}
        _ = v135
    }
    if len(v61) > 0 {
        v61[0] = (-(v3))
    }
    v2 = (v2 ^ (v2 + v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v58
    _ = v61
}
