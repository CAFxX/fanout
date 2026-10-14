package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int64; F1 uint64 }
type GenS12 struct { F0 int16; F1 uint; F2 uint8 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint32; F1 uint8 }
type GenS15 struct { F0 uint; F1 int16; F2 uint16 }
func genFn0(p0 string, p1 int32) uint16 {
    p0 = "foo bar"
    v7 := GenS1{F0: -128}
    switch p0 {
    case "hello":
        v8 := GenS2{F0: false}
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 16; fi9++ {
            if true {
                var v10 uint8 = 42
                if true {
                    switch "hello" {
                    case "\n\t":
                        v11 := []int{(^(fi9)), 0, fi9, -565}
                        _ = v11
                    case "hello":
                        p0 = p0
                    case "":
                        var v12 uint64 = uint64(v10)
                        _ = v12
                    default:
                        p1 = 1
                    }
                    p1 = (p1 ^ 314)
                    var v13 float64 = -2.25
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 7; fi14++ {
                        v15 := []int{fi9, int(v10), fi14, fi14}
                        v16 := GenS1{F0: 100}
                        v10 = v10
                        switch (fi9 ^ (fi14 * (fi14 | (fi9 ^ fi9)))) {
                        case 4294967295:
                            p0 = (("hello" + ("a" + (p0 + ((("hello" + (p0 + "hello")) + (p0 + p0)) + "\n\t")))) + "foo bar")
                        }
                        var v17 int32 = p1
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                }
                v10 -= (v10 & v10)
                _ = v10
            } else {
                var v18 float32 = ((1.5 * ((1e10 * (float32(1e10) + 1e10)) * 1.5)) * 0.1)
                v19 := GenS1{F0: -77}
                if false {
                    v18 = ((v18 - float32(-2.25)) / v18)
                }
                _ = v18
                _ = v19
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := []int{fi9}
                _ = fi20
                _ = v21
            }
            _ = fi9
        }
    default:
        p0 = "a"
    }
    _ = p0
    _ = p1
    _ = v7
    return 2
}

func main() {
    var v1 uint16 = 377
    v2 := []uint32{65535, 2147483648}
    v3 := []string{"", "a", ("\n\t" + "foo bar"), ("hello" + "hello")}
    var v4 float64 = ((1.5 * (((float64(float32((0.0 - float64(0.1)))) + float64((0.0 - 3.14159))) / (1e10 + 3.14159)) / 1e10)) * 1e10)
    var v5 string = ""
    var v6 uint64 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 -= genFn0(((("a" + "") + v5) + v5), 42)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v22 uint32 = 636
    v23 := GenS3{F0: 0, F1: 4294967295}
    var v24 int16 = 32767
    _ = len(v3)
    var v25 uint64 = 2
    if (true && true) {
        for fi26 := 0; fi26 < 5; fi26++ {
            switch "\n\t" {
            case "\n\t":
                v27 := GenS1{F0: -96}
                _ = v27
            case "":
                v6 = v6
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            switch (("hello" + "hello") + (((v5 + (v5 + v5)) + (v5 + ((v5 + v5) + "foo bar"))) + "\n\t")) {
            case "a":
                v28 := GenS4{F0: (v1 & v1), F1: v24}
                _ = v28
            case "foo bar":
                var v29 bool = (!(94 < 1))
                _ = v29
            case "":
                var v30 float64 = float64(float32(((1e10 - float64(-2.25)) * v4)))
                _ = v30
            default:
                var v31 uint64 = v6
                _ = v31
            }
            switch v5 {
            case "\n\t":
                var v32 int32 = -1
                _ = v32
            }
            _ = fi26
        }
        switch (v5 + v5) {
        case "foo bar":
            if true {
                switch -2147483648 {
                case 32768:
                    switch "foo bar" {
                    case "":
                        v33 := GenS5{F0: -522, F1: v25}
                        _ = v33
                    case "a":
                        v1 = (v1 & v1)
                    default:
                        var v34 uint32 = v22
                        _ = v34
                    }
                case 4294967295:
                    switch 570 {
                    case 128:
                        fmt.Println(v6, v25, v24)
                    case 786:
                        if false {
                            v35 := GenS6{F0: 9223372036854775807, F1: 100}
                            if false {
                                var v36 int32 = -1
                                v5 = v5
                                _ = v36
                            } else {
                                var v37 uint32 = 2147483648
                                v2 = append(v2, v22)
                                if false {
                                    for fi38 := 0; fi38 < 3; fi38++ {
                                        v2 = append(v2, (v37 ^ v37))
                                        if true {
                                            v39 := GenS1{F0: 127}
                                            _ = v39
                                        }
                                        _ = fi38
                                    }
                                    switch 843 {
                                    case 370:
                                        for fi40 := 0; fi40 < 11; fi40++ {
                                            fmt.Println(v24, v22, v37)
                                            _ = fi40
                                        }
                                    case -1:
                                        for fi41 := 0; fi41 < 17; fi41++ {
                                            fmt.Println(v24, v1, v6)
                                            var v42 bool = ((v24 / (((-(v24))) | 1)) >= v24)
                                            _ = fi41
                                            _ = v42
                                        }
                                    case 127:
                                        switch v5 {
                                        case "\n\t":
                                            v22 += uint32(v1)
                                        case "":
                                            v4 -= 0.0
                                        case "foo bar":
                                            switch (((v5 + (v5 + ("foo bar" + v5))) + "hello") + v5) {
                                            case "hello":
                                                switch "foo bar" {
                                                case "":
                                                    if true {
                                                        var v43 int = 127
                                                        for fi44 := 0; fi44 < 19; fi44++ {
                                                            v5 = v5
                                                            _ = fi44
                                                        }
                                                        v6 -= v25
                                                        _ = v43
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v37, v22, v24)
                                                case "a":
                                                    v45 := GenS6{F0: -9223372036854775808, F1: -1}
                                                    _ = v45
                                                }
                                            }
                                        }
                                    default:
                                        _ = len(v3)
                                    }
                                    var v46 uint16 = v1
                                    _ = v46
                                } else {
                                    v25 = 512
                                }
                                _ = v37
                            }
                            v47 := GenS7{F0: 1}
                            _ = v35
                            _ = v47
                        }
                    case 9223372036854775807:
                        fmt.Println(v22)
                    }
                default:
                    for fi48 := 0; fi48 < 17; fi48++ {
                        var v49 float32 = 1.5
                        v50 := GenS1{F0: 127}
                        _ = len(v2)
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                }
                switch 128 {
                case 42:
                    for fi51 := 0; fi51 < 19; fi51++ {
                        v52 := GenS8{F0: 647}
                        if ((v1 + (v1 | v1)) > v1) {
                            switch "foo bar" {
                            case "\n\t":
                                var v53 float32 = 1.5
                                _ = v53
                            case "a":
                                for fi54 := 0; fi54 < 15; fi54++ {
                                    v55 := GenS2{F0: (654 == genFn0(((v5 + v5) + (("" + "a") + "a")), -2147483648))}
                                    _ = fi54
                                    _ = v55
                                }
                            default:
                                v56 := GenS9{F0: 127, F1: (v24 - v24)}
                                _ = v56
                            }
                            v25 = (v25 | (v25 + 4294967295))
                            switch fi51 {
                            case 2147483647:
                                for fi57 := 0; fi57 < 6; fi57++ {
                                    _ = len(v3)
                                    fmt.Println(v5)
                                    v58 := GenS3{F0: 0, F1: v6}
                                    _ = len(v2)
                                    _ = fi57
                                    _ = v58
                                }
                            case -878:
                                _ = len(v2)
                            default:
                                var v59 uint = 818
                                _ = v59
                            }
                        } else {
                            v25 -= 18446744073709551615
                            v1 = v1
                            v1 = (v1 | v1)
                        }
                        _ = fi51
                        _ = v52
                    }
                case -2147483648:
                    if false {
                        v60 := GenS10{F0: -1}
                        _ = v60
                    }
                }
                var v61 uint32 = (v22 - (v22 - v22))
                _ = v61
            } else {
                switch -74 {
                case 946:
                    if (1 < -614) {
                        switch 707 {
                        case -354:
                            switch "a" {
                            case "hello":
                                v25 = v6
                            case "\n\t":
                                if false {
                                    v62 := GenS11{F0: -9223372036854775808, F1: (^(v25))}
                                    v63 := GenS2{F0: ((^(v24)) == -516)}
                                    _ = v62
                                    _ = v63
                                }
                            }
                        case -186:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            switch v5 {
                            case "a":
                                if len(v2) > 0 {
                                    v2[0] = 318
                                }
                            case "hello":
                                if true {
                                    switch -2147483648 {
                                    case 32768:
                                        v64 := GenS6{F0: -9223372036854775808, F1: -1}
                                        _ = v64
                                    case 32767:
                                        if len(v2) > 0 {
                                            v2[0] = (v22 >> ((1) & 63))
                                        }
                                    }
                                } else {
                                    v65 := GenS2{F0: ((1e10 / (v4 * 0.0)) >= v4)}
                                    var v66 bool = (32768 >= -9223372036854775808)
                                    for fi67 := 0; fi67 < 2; fi67++ {
                                        fmt.Println(v25, v1, fi67)
                                        fmt.Println(fi67)
                                        _ = fi67
                                    }
                                    _ = v65
                                    _ = v66
                                }
                            case "foo bar":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                for fi68 := 0; fi68 < 7; fi68++ {
                                    v1 = genFn0("a", -1)
                                    v24 -= 0
                                    _ = fi68
                                }
                            }
                        }
                    } else {
                        v69 := GenS11{F0: 1, F1: (^(v25))}
                        _ = v69
                    }
                default:
                    if (v24 == v24) {
                        v1 -= (v1 / ((genFn0(v5, 42)) | 1))
                        fmt.Println(v5)
                    }
                }
                if len(v3) > 0 {
                    v3[0] = (v5 + (v5 + (("" + v5) + v5)))
                }
                var v70 uint8 = 102
                _ = v70
            }
        default:
            for fi71 := 0; fi71 < 17; fi71++ {
                var v72 string = v5
                if len(v2) > 0 {
                    v2[0] = v22
                }
                var v73 int = fi71
                _ = fi71
                _ = v72
                _ = v73
            }
        }
        if (813 > 400) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v74 uint8 = 42
            _ = v74
        }
    }
    if false {
        v75 := GenS12{F0: (v24 * v24), F1: 4294967295, F2: 128}
        v5 = v5
        _ = v75
    } else {
        v22 = (v22 << ((2) & 63))
        for fi76 := 0; fi76 < 19; fi76++ {
            if (v6 != v6) {
                for fi77 := 0; fi77 < 1; fi77++ {
                    var v78 uint = 42
                    var v79 int16 = 42
                    if (int(v24) != 32768) {
                        if (genFn0(v5, 42) < genFn0(v5, 0)) {
                            v78 = (^(v78))
                            var v80 float32 = -2.25
                            var v81 string = v5
                            _ = v80
                            _ = v81
                        } else {
                            var v82 uint = v78
                            if false {
                                v6 = 888
                            }
                            if ((v1 & uint16(v25)) < (v1 * 136)) {
                                switch (v5 + v5) {
                                case "a":
                                    if ((!true) || true) {
                                        v83 := GenS10{F0: 0}
                                        _ = v83
                                    } else {
                                        var v84 int = fi77
                                        v5 = v5
                                        fmt.Println(fi76, v78, v25)
                                        v78 = 273
                                        if false {
                                            _ = len(v2)
                                            if true {
                                                v84 += fi76
                                            } else {
                                                fmt.Println(v1, v22, v25)
                                                fmt.Println(v4, v5)
                                                v85 := GenS9{F0: 0, F1: 1000}
                                                _ = v85
                                            }
                                            v24 -= (v79 >> ((v82) & 63))
                                        }
                                        v22 = (v22 << ((v82) & 63))
                                        _ = v84
                                    }
                                case "foo bar":
                                    v86 := GenS2{F0: true}
                                    switch (fi77 << (((v82 >> ((uint(v6)) & 63))) & 63)) {
                                    case 30:
                                        fmt.Println(v22, v82, v78)
                                    case 32767:
                                        var v87 int8 = 42
                                        _ = v87
                                    }
                                    v88 := GenS1{F0: -83}
                                    _ = v86
                                    _ = v88
                                case "":
                                    var v89 uint = uint(v6)
                                    _ = v89
                                }
                            }
                            if ((v4 * 1.5) != (-2.25 / -2.25)) {
                                v79 -= v24
                            }
                            var v90 int = fi77
                            _ = v82
                            _ = v90
                        }
                        v79 = v79
                        if false {
                            if true {
                                fmt.Println(v79)
                            } else {
                                var v91 int16 = (^(v24))
                                var v92 string = "foo bar"
                                _ = v91
                                _ = v92
                            }
                        } else {
                            fmt.Println(v25, v4)
                        }
                        v25 = (v25 << ((1) & 63))
                        v93 := GenS9{F0: -74, F1: v79}
                        _ = v93
                    } else {
                        var v94 uint = (v78 ^ (v78 + 1))
                        fmt.Println(v79)
                        _ = v94
                    }
                    _ = fi77
                    _ = v78
                    _ = v79
                }
                var v95 string = "a"
                _ = v95
            } else {
                for fi96 := 0; fi96 < 6; fi96++ {
                    var v97 int8 = 29
                    if false {
                        if false {
                            if true {
                                v4 -= v4
                                v98 := GenS6{F0: -9223372036854775808, F1: v97}
                                var v99 uint16 = v1
                                _ = v98
                                _ = v99
                            }
                        } else {
                            fmt.Println(v22, v97)
                        }
                        if len(v2) > 0 {
                            v2[0] = v22
                        }
                        v25 -= v6
                        v22 -= v22
                    } else {
                        v100 := GenS13{F0: v97}
                        v4 -= (((v4 + v4) + (((0.1 - v4) * ((1.5 - v4) * (0.0 / v4))) + (float64(-2.25) * v4))) + 1.5)
                        _ = v100
                    }
                    fmt.Println(v25, v4, v24)
                    _ = fi96
                    _ = v97
                }
                switch (("a" + (v5 + "a")) + v5) {
                case "a":
                    v101 := GenS12{F0: v24, F1: 42, F2: 136}
                    _ = v101
                case "hello":
                    if false {
                        var v102 int32 = -2147483648
                        _ = v102
                    } else {
                        var v103 uint = 65535
                        v25 -= (v25 % (((v25 / (((v25 * v6)) | 1))) | 1))
                        _ = v103
                    }
                default:
                    for fi104 := 0; fi104 < 3; fi104++ {
                        var v105 string = (v5 + v5)
                        var v106 int8 = 13
                        v25 = (^(v25))
                        _ = fi104
                        _ = v105
                        _ = v106
                    }
                }
                _ = len(v3)
            }
            fmt.Println(v1)
            if false {
                v107 := GenS14{F0: (^(v22)), F1: 1}
                var v108 int = (fi76 % ((fi76) | 1))
                for fi109 := 0; fi109 < 15; fi109++ {
                    v22 -= 65535
                    if len(v3) > 0 {
                        v3[0] = v5
                    }
                    _ = fi109
                }
                _ = v107
                _ = v108
            } else {
                var v110 int64 = 9223372036854775807
                _ = v110
            }
            _ = fi76
        }
    }
    var v111 int8 = 42
    fmt.Println(v25, v111)
    v112 := GenS15{F0: 18446744073709551615, F1: v24, F2: 0}
    fmt.Println(v22, v5, v111)
    v1 -= v1
    v25 += (^(v25))
    if true {
        if false {
            for fi113 := 0; fi113 < 4; fi113++ {
                var v114 uint = 80
                for fi115 := 0; fi115 < 8; fi115++ {
                    v1 -= v1
                    v25 = (v25 >> ((uint(v25)) & 63))
                    if (v5 > ("foo bar" + (v5 + v5))) {
                        var v116 int = fi115
                        v6 -= (v25 | (v25 & (v6 / (((v6 >> (((v114 | 625)) & 63))) | 1))))
                        _ = v116
                    }
                    v117 := GenS10{F0: -769}
                    var v118 int8 = v111
                    v118 += v118
                    _ = fi115
                    _ = v117
                    _ = v118
                }
                if (false || (-9223372036854775808 < 1)) {
                    var v119 float64 = v4
                    _ = v119
                }
                _ = fi113
                _ = v114
            }
        }
        for fi120 := 0; fi120 < 8; fi120++ {
            var v121 uint8 = uint8(v22)
            for fi122 := 0; fi122 < 15; fi122++ {
                if true {
                    var v123 uint16 = (v1 ^ uint16(v22))
                    _ = v123
                }
                v2 = append(v2, uint32(v1))
                var v124 int = 4294967295
                v24 = (v24 + (v24 >> ((255) & 63)))
                if false {
                    if ((-1 < v111) || (!(-1 <= 196))) {
                        var v125 int16 = 42
                        _ = v125
                    }
                    switch v5 {
                    case "a":
                        var v126 int16 = v24
                        _ = v126
                    }
                    v124 = int(v1)
                }
                _ = fi122
                _ = v124
            }
            _ = fi120
            _ = v121
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v111
    _ = v112
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
