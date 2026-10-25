package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint; F2 int }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint32; F1 int16; F2 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int }
func genFn0(p0 int16) uint8 {
    if true {
        v5 := []int{4294967295, 9223372036854775807}
        for fi6 := 0; fi6 < 7; fi6++ {
            p0 = 0
            switch fi6 {
            case 9223372036854775807:
                if len(v5) > 0 {
                    v5[0] = (fi6 * (fi6 ^ (^(fi6))))
                }
            case 128:
                switch fi6 {
                case 0:
                    for fi7 := 0; fi7 < 3; fi7++ {
                        if len(v5) > 0 {
                            v5[0] = fi7
                        }
                        var v8 int64 = 875
                        _ = fi7
                        _ = v8
                    }
                default:
                    if true {
                        for fi9 := 0; fi9 < 13; fi9++ {
                            p0 = p0
                            p0 = p0
                            v5 = append(v5, 127)
                            _ = fi9
                        }
                        p0 = 1000
                    } else {
                        p0 = (p0 / ((1) | 1))
                        for fi10 := 0; fi10 < 11; fi10++ {
                            var v11 float32 = ((((((0.1 / 1.5) - 0.1) * 0.0) / 1e10) / -2.25) + ((3.14159 * (1.5 - ((1.5 * (0.0 + 0.1)) / 1e10))) - (3.14159 - 0.1)))
                            _ = fi10
                            _ = v11
                        }
                        fmt.Println(fi6, p0)
                    }
                }
            case 709:
                for fi12 := 0; fi12 < 1; fi12++ {
                    switch ((("hello" + "hello") + "\n\t") + "") {
                    case "foo bar":
                        var v13 int16 = -289
                        _ = v13
                    case "\n\t":
                        if (3.14159 <= (1e10 / float64(((float32(1e10) - 0.1) * (((0.0 - (float32(0.0) + 1.5)) * 3.14159) + (float32((0.0 - float64((1e10 + ((-2.25 * -2.25) + 3.14159))))) - float32(1.5))))))) {
                            p0 += 68
                            p0 += (p0 + p0)
                            fmt.Println(fi6)
                        }
                    }
                    var v14 uint = 1
                    var v15 float32 = -2.25
                    _ = fi12
                    _ = v14
                    _ = v15
                }
            default:
                switch "" {
                case "foo bar":
                    if ((("foo bar" <= "a") && false) && false) {
                        if false {
                            p0 = p0
                            v16 := GenS1{F0: 1, F1: true, F2: 969}
                            for fi17 := 0; fi17 < 13; fi17++ {
                                switch ("foo bar" + ("" + "hello")) {
                                case "":
                                    var v18 int = fi17
                                    _ = v18
                                }
                                _ = len(v5)
                                if len(v5) > 0 {
                                    v5[0] = fi6
                                }
                                _ = fi17
                            }
                            _ = v16
                        } else {
                            if len(v5) > 0 {
                                v5[0] = 32768
                            }
                            v19 := GenS1{F0: 42, F1: ((3.14159 == (0.1 + 3.14159)) || (-453 > -543)), F2: 312}
                            var v20 int16 = (p0 | (p0 & (p0 ^ 747)))
                            _ = v19
                            _ = v20
                        }
                        if false {
                            var v21 uint8 = 8
                            if true {
                                v21 = (v21 % ((0) | 1))
                                var v22 int64 = 42
                                fmt.Println(v21)
                                _ = v22
                            }
                            var v23 int8 = 0
                            _ = v21
                            _ = v23
                        }
                    }
                default:
                    v5 = append(v5, fi6)
                }
            }
            _ = fi6
        }
        if (!false) {
            var v24 int64 = 42
            if false {
                fmt.Println(v24, p0)
            }
            _ = v24
        }
        _ = v5
    } else {
        v25 := []int{671}
        fmt.Println(p0)
        _ = v25
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(p0)
    case "foo bar":
        var v26 int8 = 1
        _ = v26
    case "":
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            var v27 uint64 = 1
            _ = v27
        case "hello":
            v28 := []int{128, 2147483647}
            _ = v28
        default:
            v29 := GenS2{F0: 65535, F1: 18446744073709551615, F2: -1}
            _ = v29
        }
    default:
        if (42 > 2) {
            var v30 int8 = 1
            switch ("\n\t" + "") {
            case "\n\t":
                v31 := []int{4294967295, 42}
                _ = v31
            case "hello":
                switch ((("" + ("\n\t" + "foo bar")) + "hello") + ("\n\t" + "")) {
                case "foo bar":
                    for fi32 := 0; fi32 < 17; fi32++ {
                        v33 := []int{fi32}
                        _ = fi32
                        _ = v33
                    }
                case "a":
                    var v34 uint32 = 703
                    _ = v34
                case "\n\t":
                    v30 -= v30
                default:
                    v35 := []int{9223372036854775807, 9223372036854775807, -914}
                    _ = v35
                }
            }
            _ = v30
        }
    }
    v36 := []int{}
    _ = p0
    _ = v36
    return 2
}
func genFn1(p0 int64, p1 uint64) int {
    switch 32767 {
    case -643:
        v37 := []int{-800, -368}
        _ = v37
    case 127:
        var v38 int64 = 1
        _ = v38
    case 128:
        fmt.Println(p0, p1)
    }
    var v39 uint32 = 392
    switch 127 {
    case 360:
        v39 += v39
    case 1:
        var v40 int64 = 42
        _ = v40
    case -9223372036854775808:
        v41 := GenS2{F0: (v39 ^ v39), F1: 1, F2: 32768}
        _ = v41
    default:
        if true {
            for fi42 := 0; fi42 < 20; fi42++ {
                fmt.Println(v39, p1, v39)
                if (!(-32768 != 0)) {
                    v43 := GenS1{F0: 32767, F1: true, F2: (^(p1))}
                    var v44 bool = (true && false)
                    if (false && (v44 || v44)) {
                        for fi45 := 0; fi45 < 12; fi45++ {
                            switch 42 {
                            case 674:
                                if (3.14159 >= -2.25) {
                                    v46 := GenS3{F0: "\n\t", F1: 1}
                                    v47 := GenS4{F0: 42, F1: ("foo bar" + "foo bar")}
                                    _ = v46
                                    _ = v47
                                } else {
                                    if true {
                                        p0 -= (p0 + (p0 + -647))
                                    }
                                }
                            default:
                                if false {
                                    if v44 {
                                        fmt.Println(v39, p0)
                                        v48 := []int{fi42, (fi42 & (fi42 << ((670) & 63)))}
                                        _ = v48
                                    }
                                    v44 = v44
                                } else {
                                    if v44 {
                                        v49 := GenS3{F0: "\n\t", F1: 127}
                                        var v50 float64 = (1e10 * 0.1)
                                        fmt.Println(fi45, p1)
                                        _ = v49
                                        _ = v50
                                    }
                                    v51 := []int{fi45, fi45, fi45}
                                    _ = v51
                                }
                            }
                            if (70 != 100) {
                                fmt.Println(fi45)
                            } else {
                                v52 := []int{591, (-(fi42)), fi45, fi42}
                                p0 = (p0 % (((p0 + -9223372036854775808)) | 1))
                                _ = v52
                            }
                            _ = fi45
                        }
                    }
                    _ = v43
                    _ = v44
                } else {
                    p1 += 165
                    v53 := GenS5{F0: 0, F1: -32768, F2: fi42}
                    _ = v53
                }
                v54 := []int{fi42, int(v39), (-(fi42))}
                _ = fi42
                _ = v54
            }
            var v55 uint16 = 1
            _ = v55
        } else {
            v39 = (v39 ^ 1)
            if (!false) {
                var v56 int64 = p0
                var v57 int64 = int64(p1)
                if (float32(3.14159) == -2.25) {
                    if (true || true) {
                        if ((-36 <= int8(p1)) || (genFn0(-88) > 115)) {
                            if (((-2.25 - 3.14159) + (3.14159 * (1.5 + 1.5))) > (0.0 + (0.0 * 1.5))) {
                                v58 := GenS6{F0: (p1 + p1)}
                                _ = v58
                            }
                            var v59 int64 = v56
                            _ = v59
                        }
                        p0 += 0
                        p1 = 18446744073709551615
                    }
                    var v60 uint16 = 42
                    _ = v60
                } else {
                    p1 += 0
                }
                _ = v56
                _ = v57
            }
        }
    }
    var v61 bool = (1.5 <= -2.25)
    _ = p0
    _ = p1
    _ = v39
    _ = v61
    return 128
}

func main() {
    var v1 uint = 266
    var v2 bool = false
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    for fi62 := 0; fi62 < 10; fi62++ {
        fmt.Println(v2, fi62, v1)
        v63 := []int{fi62, fi62, (-(fi62))}
        _ = fi62
        _ = v63
    }
    if (4294967295 != 664) {
        if (0.0 != ((1e10 - (-2.25 * float32(3.14159))) - 0.0)) {
            v64 := GenS1{F0: 2, F1: v2, F2: 114}
            if v2 {
                var v65 uint32 = 499
                switch 850 {
                case 9223372036854775807:
                    var v66 int32 = 462
                    _ = v66
                case 4294967295:
                    v65 = (^(v65))
                case 0:
                    _ = len(v4)
                }
                v65 = (v65 & v65)
                _ = v65
            }
            v2 = false
            _ = v64
        } else {
            var v67 float64 = (0.1 * 3.14159)
            _ = v67
        }
        v1 = v1
    } else {
        for fi68 := 0; fi68 < 17; fi68++ {
            delete(v3, 4294967295)
            v69 := []int{}
            _ = fi68
            _ = v69
        }
    }
    v2 = v2
    for fi70 := 0; fi70 < 5; fi70++ {
        if v2 {
            v71 := []int{274, (fi70 % ((fi70) | 1)), -971, genFn1(42, 2)}
            _ = v71
        } else {
            var v72 uint64 = uint64(v1)
            for fi73 := 0; fi73 < 9; fi73++ {
                switch 4294967295 {
                case 2147483647:
                    var v74 string = "foo bar"
                    _ = v74
                }
                switch fi73 {
                case -1:
                    v75 := v4[755]
                    _ = v75
                }
                fmt.Println(v2, v72)
                var v76 string = ("hello" + "")
                switch "hello" {
                case "":
                    var v77 float64 = 0.0
                    _ = v77
                case "foo bar":
                    fmt.Println(v1, v2)
                case "hello":
                    v78 := GenS1{F0: 65535, F1: (-32768 <= 1000), F2: 992}
                    _ = v78
                }
                var v79 string = ""
                _ = fi73
                _ = v76
                _ = v79
            }
            _ = v72
        }
        v1 -= v1
        var v80 bool = v2
        _ = fi70
        _ = v80
    }
    switch -729 {
    case -9223372036854775808:
        switch 42 {
        case 32768:
            v2 = (-2147483648 < 1)
        case 0:
            if v2 {
                for fi81 := 0; fi81 < 20; fi81++ {
                    v82 := []int{fi81, genFn1(211, 42)}
                    v2 = v2
                    _ = fi81
                    _ = v82
                }
                var v83 string = "\n\t"
                v2 = v2
                _ = v83
            }
        case 256:
            for fi84 := 0; fi84 < 3; fi84++ {
                var v85 int32 = -780
                var v86 int64 = 680
                v87 := GenS3{F0: "foo bar", F1: 0}
                _ = fi84
                _ = v85
                _ = v86
                _ = v87
            }
        default:
            v88 := v3[1]
            _ = v88
        }
    case -1:
        if v2 {
            if (0.1 > 1e10) {
                v1 += 1
                for fi89 := 0; fi89 < 5; fi89++ {
                    v1 -= 1
                    fmt.Println(v1)
                    if (!((((0 == 408) && v2) || ((2 != 1) || (true || (42 >= 1)))) && (-493 < 32767))) {
                        switch (fi89 / ((612) | 1)) {
                        case -747:
                            var v90 uint8 = 42
                            _ = v90
                        case 1:
                            for fi91 := 0; fi91 < 6; fi91++ {
                                if true {
                                    v92 := []int{genFn1(623, 401), fi91, (-(fi91))}
                                    fmt.Println(v2, v1, v1)
                                    _ = v92
                                } else {
                                    v2 = ((-31 != 34) || (121 > 32767))
                                }
                                switch ("foo bar" + "foo bar") {
                                case "a":
                                    v93 := GenS4{F0: int32(v1), F1: (("foo bar" + "hello") + "\n\t")}
                                    _ = v93
                                case "foo bar":
                                    v94 := GenS5{F0: 2147483648, F1: 185, F2: fi91}
                                    _ = v94
                                default:
                                    if v2 {
                                        v2 = (208 != 1)
                                    }
                                }
                                v95 := v4[uint32(v1)]
                                _ = fi91
                                _ = v95
                            }
                        case -616:
                            var v96 uint = (v1 + v1)
                            _ = v96
                        }
                        _ = len(v3)
                    }
                    _ = fi89
                }
            } else {
                v97 := []int{}
                var v98 uint64 = 2
                v99 := GenS3{F0: "a", F1: 40}
                _ = v97
                _ = v98
                _ = v99
            }
        } else {
            switch -61 {
            case 32767:
                v1 = 42
            default:
                v2 = ((255 >= 48) || (4294967295 < 1))
            }
        }
    default:
        var v100 uint = 808
        _ = v100
    }
    var v101 string = ""
    for fi102 := 0; fi102 < 14; fi102++ {
        v3[1] = (v2 && (true || v2))
        fmt.Println(v2, v101, v2)
        var v103 float32 = 0.1
        _ = fi102
        _ = v103
    }
    var v104 float32 = 0.1
    v1 = (v1 + 4294967295)
    for fi105 := 0; fi105 < 6; fi105++ {
        if false {
            if (0 != 42) {
                fmt.Println(v101, v101, v1)
            }
        } else {
            switch 128 {
            case -2147483648:
                _ = len(v3)
            case -1:
                switch v101 {
                case "foo bar":
                    for fi106 := 0; fi106 < 4; fi106++ {
                        if (32 <= -1) {
                            var v107 string = (((("a" + (("" + v101) + v101)) + "hello") + "") + (v101 + v101))
                            v108 := GenS5{F0: 684, F1: 0, F2: (fi105 % (((fi105 % (((fi106 + fi105)) | 1))) | 1))}
                            _ = v107
                            _ = v108
                        }
                        v109 := []int{}
                        _ = len(v4)
                        v101 = (((v101 + (v101 + ("hello" + "hello"))) + ((v101 + "hello") + v101)) + "")
                        switch fi105 {
                        case -654:
                            if len(v109) >= 2 {
                                v109 = v109[1:len(v109)]
                            }
                        case 0:
                            if true {
                                var v110 float64 = (1e10 + 1.5)
                                v111 := GenS7{F0: 4294967295, F1: v101}
                                fmt.Println(v104)
                                _ = v110
                                _ = v111
                            }
                        case 32768:
                            if true {
                                var v112 uint8 = 128
                                _ = v112
                            }
                        default:
                            v101 = v101
                        }
                        fmt.Println(v101, fi106)
                        _ = fi106
                        _ = v109
                    }
                case "":
                    var v113 uint = v1
                    _ = v113
                default:
                    v114 := v3[0]
                    _ = v114
                }
            }
            var v115 bool = (!(v2 || (2 > 358)))
            for fi116 := 0; fi116 < 1; fi116++ {
                v1 += (^(v1))
                switch (fi105 * fi105) {
                case 2147483647:
                    if (true || (!(1 == 2))) {
                        v117 := GenS7{F0: 2, F1: v101}
                        _ = v117
                    }
                case 4294967295:
                    v1 += 0
                case 18:
                    v118 := GenS8{F0: 2147483647}
                    _ = v118
                default:
                    fmt.Println(v104, fi105)
                }
                v2 = false
                var v119 int32 = 0
                v1 -= (v1 | (v1 / ((252) | 1)))
                _ = fi116
                _ = v119
            }
            _ = v115
        }
        var v120 string = v101
        v101 = "foo bar"
        _ = fi105
        _ = v120
    }
    v104 -= v104
    v121 := []int{-9223372036854775808, 1}
    var v122 uint32 = 2147483647
    for fi123 := 0; fi123 < 20; fi123++ {
        if (80 != 111) {
            var v124 float32 = (float32(((float64(float32(0.1)) + ((1e10 - (1.5 * (3.14159 / 0.1))) - ((1e10 * (0.1 + (((1e10 * float64(float32(3.14159))) / float64(float32(0.1))) * (1.5 + ((1.5 / -2.25) / 1.0))))) + 3.14159))) / 0.1)) * v104)
            var v125 uint16 = uint16(v1)
            v126 := GenS9{F0: 1}
            _ = v124
            _ = v125
            _ = v126
        } else {
            _ = len(v4)
            v122 = v122
            for fi127 := 0; fi127 < 13; fi127++ {
                var v128 int = 128
                var v129 int = (v128 % (((v128 % ((fi123) | 1))) | 1))
                v121 = append(v121, 32768)
                _ = fi127
                _ = v128
                _ = v129
            }
        }
        _ = fi123
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v104:", v104)
    fmt.Println("slicelen v121:", len(v121), v121)
    fmt.Println("v122:", v122)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v101
    _ = v104
    _ = v121
    _ = v122
    _ = v2
    _ = v3
    _ = v4
}
