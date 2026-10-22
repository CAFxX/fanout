package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 int32; F2 int16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS12 struct { F0 uint; F1 uint16; F2 bool }
type GenS13 struct { F0 int8; F1 int32; F2 uint }
type GenS14 struct { F0 int; F1 int8 }
type GenS15 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint8, p1 string, p2 int32) uint64 {
    p0 = 255
    if (false && true) {
        for fi4 := 0; fi4 < 11; fi4++ {
            p0 += p0
            _ = fi4
        }
        v5 := GenS1{F0: 100, F1: 42}
        for fi6 := 0; fi6 < 4; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                switch fi6 {
                case 827:
                    v8 := []int{1, (fi7 + fi6)}
                    _ = v8
                case 128:
                    var v9 float64 = 0.1
                    _ = v9
                default:
                    fmt.Println(p0, fi6, p1)
                }
                _ = fi7
            }
            _ = fi6
        }
        _ = v5
    } else {
        p2 += 885
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 uint, p1 int32, p2 bool) int8 {
    var v10 int8 = 127
    v11 := GenS1{F0: v10, F1: 2}
    v12 := GenS1{F0: 112, F1: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return v10
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 42
    v3 := []uint32{729, 2, 2147483648}
    if ((0.0 - ((((0.1 / 1.0) / 1.5) / 1e10) - 0.1)) < -2.25) {
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS1{F0: (v2 << ((575) & 63)), F1: 98}
            for fi15 := 0; fi15 < 18; fi15++ {
                v16 := GenS2{F0: ""}
                v2 += (v2 * (v2 >> ((4294967295) & 63)))
                v2 = 42
                fmt.Println(v2, fi15, fi15)
                _ = fi15
                _ = v16
            }
            var v17 int32 = -1
            _ = fi13
            _ = v14
            _ = v17
        }
        v2 = v2
        v18 := GenS2{F0: "\n\t"}
        _ = v18
    } else {
        v2 += v2
        v2 -= v2
    }
    switch -9223372036854775808 {
    case 57:
        v19 := v1[42]
        _ = v19
    case 32768:
        fmt.Println(v2)
    case -2147483648:
        var v20 int = -1
        _ = v20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            if true {
                if true {
                    if false {
                        var v23 uint = 628
                        switch fi22 {
                        case 128:
                            fmt.Println(fi21)
                        case 32768:
                            delete(v1, 4294967295)
                        }
                        _ = v23
                    } else {
                        v2 += (^(v2))
                        if (false || (true || true)) {
                            var v24 int64 = -984
                            _ = v24
                        } else {
                            fmt.Println(fi22, fi21)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v25 := v1[4294967295]
                            v26 := GenS3{F0: false, F1: v25}
                            _ = v25
                            _ = v26
                        }
                    }
                    v2 = 127
                    fmt.Println(fi22, v2)
                }
                if (-2147483648 != -1) {
                    var v27 string = "hello"
                    if false {
                        var v28 int = (fi22 - -2147483648)
                        switch fi21 {
                        case 2147483647:
                            switch "a" {
                            case "":
                                v29 := GenS2{F0: (v27 + (v27 + ((("\n\t" + v27) + ("foo bar" + v27)) + v27)))}
                                _ = v29
                            case "\n\t":
                                v30 := GenS1{F0: (^(v2)), F1: 89}
                                _ = v30
                            case "a":
                                fmt.Println(fi22, v28, v2)
                            }
                        case -9223372036854775808:
                            fmt.Println(fi22)
                        case 622:
                            var v31 int64 = 925
                            _ = v31
                        }
                        _ = v28
                    } else {
                        var v32 string = v27
                        v33 := GenS1{F0: v2, F1: 32767}
                        var v34 float32 = 3.14159
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    switch v27 {
                    case "a":
                        v35 := GenS2{F0: v27}
                        _ = v35
                    case "foo bar":
                        if (false && false) {
                            if true {
                                var v36 float32 = 0.1
                                _ = v36
                            }
                            v37 := GenS4{F0: 4294967295}
                            _ = v37
                        } else {
                            if ((v2 != v2) && (1 <= -9223372036854775808)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            if false {
                                fmt.Println(fi22, fi22, fi21)
                                switch fi21 {
                                case 32768:
                                    v38 := GenS4{F0: 313}
                                    v2 = 100
                                    _ = v38
                                case 878:
                                    var v39 int8 = v2
                                    _ = v39
                                }
                            } else {
                                delete(v1, 770)
                                v40 := GenS4{F0: 678}
                                _ = v40
                            }
                        }
                    default:
                        var v41 uint = 996
                        _ = v41
                    }
                    _ = v27
                } else {
                    v42 := GenS5{F0: 0, F1: 552, F2: 1}
                    _ = v42
                }
            } else {
                var v43 int16 = 1
                v2 += v2
                _ = v43
            }
            var v44 uint32 = 0
            v45 := v1[v44]
            _ = fi22
            _ = v44
            _ = v45
        }
        _ = fi21
    }
    _ = len(v3)
    switch -424 {
    case 335:
        for fi46 := 0; fi46 < 4; fi46++ {
            switch (("\n\t" + "") + (("hello" + "hello") + ("hello" + "hello"))) {
            case "":
                switch 0 {
                case 621:
                    v2 = v2
                case 9223372036854775807:
                    if false {
                        v2 -= v2
                    } else {
                        for fi47 := 0; fi47 < 4; fi47++ {
                            var v48 float64 = 3.14159
                            fmt.Println(v2)
                            var v49 float32 = 3.14159
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                        if (!true) {
                            v50 := GenS6{F0: 1}
                            var v51 float64 = (1e10 * 0.1)
                            v51 -= -2.25
                            _ = v50
                            _ = v51
                        }
                        v2 = v2
                    }
                default:
                    v2 -= (v2 >> ((727) & 63))
                }
            case "\n\t":
                for fi52 := 0; fi52 < 17; fi52++ {
                    switch (fi46 << ((219) & 63)) {
                    case 2147483647:
                        v2 -= v2
                    default:
                        v53 := GenS7{F0: fi52, F1: 1, F2: 222}
                        _ = v53
                    }
                    if (false || false) {
                        switch -483 {
                        case -84:
                            v2 = (v2 + 116)
                        }
                    }
                    _ = fi52
                }
            case "a":
                v54 := GenS7{F0: fi46, F1: 42, F2: -1}
                _ = v54
            default:
                delete(v1, 42)
            }
            v2 = genFn1(686, -1, true)
            var v55 int = 2147483647
            _ = fi46
            _ = v55
        }
    case 4294967295:
        var v56 int8 = (v2 >> ((47) & 63))
        _ = v56
    }
    if (false || (791 == -101)) {
        var v57 int64 = -1
        _ = v57
    } else {
        var v58 float32 = -2.25
        _ = v58
    }
    if true {
        var v59 bool = ((float64(1.5) + 0.1) != ((0.0 / float64((((0.1 * 1e10) - float32(float64(0.0))) + -2.25))) * float64(((1e10 / ((1e10 / (-2.25 * 3.14159)) * float32(-2.25))) + (((1e10 - 0.1) + (((3.14159 / -2.25) + 1.5) - 1.5)) - (float32(3.14159) / 1e10))))))
        v60 := GenS3{F0: true, F1: 1}
        _ = v59
        _ = v60
    } else {
        switch 2147483647 {
        case 32768:
            if false {
                if true {
                    v2 += (v2 * v2)
                    for fi61 := 0; fi61 < 16; fi61++ {
                        switch "hello" {
                        case "":
                            switch 577 {
                            case 0:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case 32768:
                                var v62 float64 = -2.25
                                _ = v62
                            }
                        case "foo bar":
                            _ = len(v3)
                        case "\n\t":
                            v2 = 42
                        default:
                            for fi63 := 0; fi63 < 2; fi63++ {
                                if false {
                                    if (-2147483648 < 1) {
                                        v2 = int8(fi63)
                                        v2 += v2
                                        if (1 > 32767) {
                                            if false {
                                                v64 := v1[65535]
                                                v64 += (v64 << ((18446744073709551615) & 63))
                                                _ = v64
                                            }
                                            var v65 uint8 = 128
                                            _ = v65
                                        }
                                        switch (fi61 | (fi63 << ((468) & 63))) {
                                        case -2147483648:
                                            switch ("a" + ("hello" + "a")) {
                                            case "foo bar":
                                                if (false && false) {
                                                    v2 -= 0
                                                    v2 = int8(fi61)
                                                }
                                            case "a":
                                                if true {
                                                    var v66 uint64 = 11
                                                    delete(v1, 698)
                                                    v67 := GenS5{F0: int32(v66), F1: 271, F2: 2}
                                                    _ = v66
                                                    _ = v67
                                                } else {
                                                    if false {
                                                        v2 -= genFn1(18446744073709551615, 0, true)
                                                    }
                                                    if (42 <= 0) {
                                                        var v68 int8 = v2
                                                        _ = v68
                                                    }
                                                    if false {
                                                        if false {
                                                            if (255 > 255) {
                                                                v69 := GenS8{F0: (((("" + ("foo bar" + ("\n\t" + "a"))) + "foo bar") + ("hello" + "\n\t")) + "hello")}
                                                                var v70 uint8 = 238
                                                                var v71 int16 = -32768
                                                                _ = v69
                                                                _ = v70
                                                                _ = v71
                                                            } else {
                                                                var v72 uint = 209
                                                                v2 = (v2 | (v2 % (((v2 * (v2 ^ v2))) | 1)))
                                                                v73 := GenS2{F0: "a"}
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                        }
                                                        var v74 uint = 18446744073709551615
                                                        if ("\n\t" >= (("" + "a") + ("a" + "a"))) {
                                                            v75 := GenS1{F0: (v2 >> (((v74 | 0)) & 63)), F1: 32767}
                                                            var v76 int32 = -723
                                                            v2 = v2
                                                            if (!true) {
                                                                if false {
                                                                    var v77 uint8 = 1
                                                                    var v78 int64 = 1
                                                                    _ = len(v1)
                                                                    _ = v77
                                                                    _ = v78
                                                                }
                                                                if (128 != 160) {
                                                                    v76 -= (v76 * (v76 % ((0) | 1)))
                                                                    v74 -= 42
                                                                    v2 += v2
                                                                }
                                                            }
                                                            v76 = -1
                                                            v79 := v1[2147483647]
                                                            var v80 int64 = -764
                                                            _ = v75
                                                            _ = v76
                                                            _ = v79
                                                            _ = v80
                                                        } else {
                                                            v2 = v2
                                                            if len(v3) > 0 {
                                                                v3[0] = 4294967295
                                                            }
                                                            v81 := GenS9{F0: "hello"}
                                                            _ = v81
                                                        }
                                                        _ = v74
                                                    } else {
                                                        v2 = (v2 >> ((65535) & 63))
                                                        v82 := GenS10{F0: 66}
                                                        _ = v82
                                                    }
                                                }
                                            case "\n\t":
                                                var v83 float64 = 0.1
                                                v2 -= 1
                                                var v84 int32 = int32(fi63)
                                                if true {
                                                    delete(v1, 0)
                                                } else {
                                                    fmt.Println(v2)
                                                }
                                                _ = v83
                                                _ = v84
                                            default:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v2 = 0
                                                default:
                                                    v85 := GenS8{F0: "foo bar"}
                                                    _ = v85
                                                }
                                            }
                                        case 537:
                                            v2 -= 127
                                        case 32767:
                                            switch "hello" {
                                            case "\n\t":
                                                if len(v3) > 0 {
                                                    v3[0] = 42
                                                }
                                            case "":
                                                fmt.Println(v2, v2)
                                            default:
                                                var v86 int16 = -534
                                                _ = v86
                                            }
                                        }
                                    } else {
                                        var v87 float64 = (3.14159 * 0.1)
                                        switch (fi63 >> ((192) & 63)) {
                                        case 4294967295:
                                            if (!true) {
                                                v88 := GenS11{F0: (v2 + v2), F1: 0, F2: 2}
                                                fmt.Println(fi63, v87)
                                                if true {
                                                    if false {
                                                        var v89 float32 = -2.25
                                                        _ = v89
                                                    } else {
                                                        if false {
                                                            v2 = (v2 / ((v2) | 1))
                                                        }
                                                        v87 += -2.25
                                                    }
                                                    v90 := GenS12{F0: 40, F1: 679, F2: true}
                                                    v2 = (v2 ^ genFn1(2, 2147483647, false))
                                                    _ = v90
                                                }
                                                _ = v88
                                            }
                                        case 32767:
                                            switch -9223372036854775808 {
                                            case 0:
                                                var v91 float32 = 1e10
                                                _ = v91
                                            case 32767:
                                                _ = len(v3)
                                            case 800:
                                                if true {
                                                    if (!false) {
                                                        v2 += 1
                                                    }
                                                    v2 = (v2 + -100)
                                                }
                                            default:
                                                v92 := GenS13{F0: v2, F1: 0, F2: 65535}
                                                _ = v92
                                            }
                                        }
                                        v93 := v1[42]
                                        _ = v87
                                        _ = v93
                                    }
                                }
                                fmt.Println(fi63)
                                v94 := GenS7{F0: fi63, F1: -1, F2: 0}
                                v95 := GenS14{F0: fi61, F1: 127}
                                _ = fi63
                                _ = v94
                                _ = v95
                            }
                        }
                        v96 := GenS8{F0: ""}
                        v97 := GenS14{F0: 4294967295, F1: v2}
                        _ = fi61
                        _ = v96
                        _ = v97
                    }
                } else {
                    v2 = v2
                    var v98 uint32 = 1
                    _ = v98
                }
            }
        case -1:
            var v99 bool = false
            _ = v99
        default:
            var v100 int32 = 0
            _ = v100
        }
        switch -2147483648 {
        case -2147483648:
            if (!true) {
                switch 9223372036854775807 {
                case 2147483647:
                    for fi101 := 0; fi101 < 1; fi101++ {
                        switch "foo bar" {
                        case "":
                            if (668 != 42) {
                                v2 += 127
                                var v102 uint = 18446744073709551615
                                _ = v102
                            } else {
                                v2 -= (v2 - (^(v2)))
                                _ = len(v3)
                                for fi103 := 0; fi103 < 6; fi103++ {
                                    if (4294967295 <= 747) {
                                        var v104 uint8 = 2
                                        _ = len(v1)
                                        v105 := GenS10{F0: 0}
                                        _ = v104
                                        _ = v105
                                    } else {
                                        if (!(65535 != 255)) {
                                            v2 = v2
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v106 uint8 = 42
                                            _ = v106
                                        default:
                                            var v107 uint16 = 0
                                            _ = v107
                                        }
                                        _ = len(v1)
                                        var v108 int16 = -249
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v108
                                    }
                                    _ = fi103
                                }
                            }
                        }
                        if true {
                            switch "a" {
                            case "":
                                var v109 int64 = 0
                                _ = v109
                            case "hello":
                                v2 = genFn1(0, 486, (50 != 2))
                            default:
                                if true {
                                    var v110 bool = false
                                    if false {
                                        v2 -= v2
                                        if len(v3) > 0 {
                                            v3[0] = 2
                                        }
                                        v2 += v2
                                    }
                                    _ = v110
                                }
                            }
                            var v111 uint8 = 42
                            v112 := GenS7{F0: fi101, F1: 1, F2: -1}
                            _ = v111
                            _ = v112
                        }
                        _ = fi101
                    }
                default:
                    var v113 uint8 = 1
                    _ = v113
                }
            } else {
                switch 947 {
                case 32767:
                    fmt.Println(v2)
                case 9223372036854775807:
                    v2 -= (-(v2))
                }
                if (160 >= 61) {
                    for fi114 := 0; fi114 < 11; fi114++ {
                        v2 = (v2 / (((v2 ^ (v2 | genFn1(18446744073709551615, -1, true)))) | 1))
                        fmt.Println(fi114, v2)
                        _ = fi114
                    }
                    var v115 int8 = genFn1(65535, -1, true)
                    var v116 int64 = 381
                    _ = v115
                    _ = v116
                } else {
                    for fi117 := 0; fi117 < 8; fi117++ {
                        v2 += genFn1(1, 297, (2147483648 <= 42))
                        v118 := GenS15{F0: 242, F1: 131}
                        _ = fi117
                        _ = v118
                    }
                    fmt.Println(v2)
                }
                var v119 string = "\n\t"
                _ = v119
            }
        }
        for fi120 := 0; fi120 < 20; fi120++ {
            delete(v1, 865)
            v2 -= (v2 * v2)
            _ = fi120
        }
    }
    v2 = 100
    for fi121 := 0; fi121 < 7; fi121++ {
        v2 += (v2 - 42)
        var v122 float32 = -2.25
        _ = fi121
        _ = v122
    }
    var v123 float32 = (0.1 + (3.14159 + 0.0))
    switch 4294967295 {
    case 42:
        v124 := GenS7{F0: 128, F1: -1, F2: -704}
        _ = v124
    }
    if (!(true || false)) {
        var v125 uint8 = 1
        fmt.Println(v2, v2)
        _ = v125
    }
    fmt.Println(v2)
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v123
    _ = v2
    _ = v3
}
