package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 int16; F2 int64 }
type GenS6 struct { F0 bool; F1 string; F2 uint64 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 int8; F1 int32; F2 string }
type GenS9 struct { F0 int; F1 uint32 }
func genFn0() int32 {
    v5 := GenS1{F0: 132}
    var v6 int16 = 32767
    var v7 uint8 = 171
    v8 := []int{-70, -533, 4294967295, 128}
    switch 9223372036854775807 {
    case 2147483647:
        for fi9 := 0; fi9 < 8; fi9++ {
            switch 32768 {
            case 467:
                if true {
                    var v10 uint64 = 0
                    _ = v10
                }
            case -278:
                fmt.Println(v7, v6, v7)
            case 32767:
                v7 = v7
            }
            _ = fi9
        }
    case 4294967295:
        for fi11 := 0; fi11 < 4; fi11++ {
            v12 := GenS1{F0: 42}
            _ = fi11
            _ = v12
        }
    case -722:
        _ = len(v8)
    default:
        var v13 float64 = 1e10
        _ = v13
    }
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int, p1 string) uint32 {
    v14 := GenS2{F0: 0, F1: 0}
    v15 := []int{p0, p0}
    var v16 int16 = 1000
    p0 -= p0
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "\n\t"
    var v3 uint = 242
    v4 := make(map[int]bool)
    if (833 == 32768) {
        v17 := GenS2{F0: v3, F1: 255}
        for fi18 := 0; fi18 < 10; fi18++ {
            delete(v1, 65535)
            switch (v2 + ((v2 + "a") + "")) {
            case "\n\t":
                var v19 bool = false
                _ = v19
            default:
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 int8 = 42
                    _ = fi20
                    _ = v21
                }
            }
            _ = fi18
        }
        _ = v17
    } else {
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                v24 := []int{}
                _ = fi23
                _ = v24
            }
            fmt.Println(fi22, fi22, v3)
            _ = fi22
        }
        var v25 int16 = 1
        var v26 int32 = 0
        _ = v25
        _ = v26
    }
    switch 4294967295 {
    case 0:
        var v27 int = 1
        _ = v27
    }
    switch 357 {
    case 2147483647:
        v28 := GenS1{F0: 277}
        _ = v28
    case 0:
        var v29 bool = true
        _ = v29
    default:
        for fi30 := 0; fi30 < 4; fi30++ {
            v31 := []int{(fi30 ^ (fi30 - -1)), (fi30 * (fi30 >> ((v3) & 63))), (fi30 * -1)}
            _ = fi30
            _ = v31
        }
    }
    var v32 bool = false
    v33 := GenS2{F0: (v3 ^ v3), F1: 409}
    if v32 {
        v3 -= (v3 / ((v3) | 1))
        switch ((v2 + v2) + (v2 + "")) {
        case "\n\t":
            var v34 uint32 = 365
            _ = v34
        case "a":
            if (((v2 == "") && (!false)) || v32) {
                var v35 bool = (0 != 255)
                _ = v35
            }
        case "foo bar":
            v36 := []int{-755, -257, 42, -9223372036854775808}
            _ = v36
        default:
            fmt.Println(v2)
        }
        fmt.Println(v32, v32)
    } else {
        var v37 string = v2
        _ = v37
    }
    v32 = false
    fmt.Println(v32, v32, v32)
    for fi38 := 0; fi38 < 7; fi38++ {
        switch fi38 {
        case -1:
            var v39 int32 = 87
            _ = v39
        case 32767:
            var v40 float32 = ((-2.25 - (((((-2.25 / float32(float64((0.1 - ((1e10 - 0.1) / -2.25))))) * ((1e10 / float32((3.14159 - -2.25))) / 1.0)) * 1.5) - float32((float64(0.1) / float64((1e10 + -2.25))))) + 1.5)) / 1.0)
            _ = v40
        }
        switch (fi38 * (^(fi38))) {
        case -414:
            if (!(v32 || false)) {
                var v41 uint = v3
                v42 := GenS2{F0: 1, F1: 0}
                _ = v41
                _ = v42
            }
        case -324:
            var v43 int32 = 2147483647
            _ = v43
        }
        _ = fi38
    }
    v44 := GenS2{F0: v3, F1: uint16(v3)}
    if (32767 != -15) {
        fmt.Println(v3)
        for fi45 := 0; fi45 < 14; fi45++ {
            if true {
                switch ((v2 + ((v2 + "\n\t") + v2)) + v2) {
                case "\n\t":
                    if false {
                        _ = len(v1)
                        switch ((v2 + ((((((v2 + v2) + v2) + "\n\t") + "hello") + ("\n\t" + "")) + v2)) + "foo bar") {
                        case "\n\t":
                            if false {
                                if (0.1 >= 1e10) {
                                    switch v2 {
                                    case "":
                                        v3 -= (v3 & v3)
                                    default:
                                        v46 := []int{(fi45 * (fi45 << ((v3) & 63))), fi45, 128, fi45}
                                        _ = v46
                                    }
                                    v47 := []int{1}
                                    if v32 {
                                        if false {
                                            for fi48 := 0; fi48 < 8; fi48++ {
                                                var v49 uint16 = 255
                                                _ = fi48
                                                _ = v49
                                            }
                                            var v50 uint8 = 128
                                            v47 = append(v47, fi45)
                                            _ = v50
                                        } else {
                                            switch (v2 + ((v2 + v2) + v2)) {
                                            case "\n\t":
                                                fmt.Println(fi45, v3)
                                            case "":
                                                for fi51 := 0; fi51 < 3; fi51++ {
                                                    if false {
                                                        if v32 {
                                                            switch ((v2 + "\n\t") + v2) {
                                                            case "":
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    v52 := GenS3{F0: 65535}
                                                                    _ = v52
                                                                case "a":
                                                                    switch fi45 {
                                                                    case -486:
                                                                        v3 = (v3 | v3)
                                                                    case 954:
                                                                        if false {
                                                                            if (545 >= 42) {
                                                                                fmt.Println(v3)
                                                                                if (!(42 != 2147483647)) {
                                                                                    switch ("\n\t" + v2) {
                                                                                    case "foo bar":
                                                                                        v53 := GenS1{F0: 458}
                                                                                        _ = v53
                                                                                    case "":
                                                                                        fmt.Println(fi51, v2)
                                                                                    case "\n\t":
                                                                                        v54 := GenS1{F0: 1}
                                                                                        _ = v54
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi45)
                                                                            } else {
                                                                                v32 = ((705 <= 538) || true)
                                                                            }
                                                                            if true {
                                                                                switch (fi51 << (((v3 | v3)) & 63)) {
                                                                                case 9:
                                                                                    var v55 float32 = ((0.0 / (float32(1.5) / 1.5)) + float32(0.0))
                                                                                    _ = v55
                                                                                default:
                                                                                    var v56 string = "\n\t"
                                                                                    _ = v56
                                                                                }
                                                                                if v32 {
                                                                                    v32 = (1.5 <= 0.1)
                                                                                } else {
                                                                                    if len(v47) > 0 {
                                                                                        v47[0] = fi51
                                                                                    }
                                                                                }
                                                                                delete(v4, fi51)
                                                                                v32 = false
                                                                                switch fi45 {
                                                                                case 32767:
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        if len(v47) > 0 {
                                                                                            v47[0] = (fi45 / ((fi45) | 1))
                                                                                        }
                                                                                    case "\n\t":
                                                                                        fmt.Println(fi51)
                                                                                    case "foo bar":
                                                                                        v3 = 65535
                                                                                    default:
                                                                                        v3 = v3
                                                                                    }
                                                                                case 42:
                                                                                    if v32 {
                                                                                        v57 := GenS3{F0: 255}
                                                                                        _ = v57
                                                                                    } else {
                                                                                        v58 := GenS1{F0: 2}
                                                                                        v59 := GenS4{F0: 1}
                                                                                        _ = v58
                                                                                        _ = v59
                                                                                    }
                                                                                case 128:
                                                                                    v60 := GenS5{F0: (((1e10 / 1.0) + float32(1e10)) == float32(1.5)), F1: 32767, F2: -887}
                                                                                    _ = v60
                                                                                default:
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        v47 = append(v47, (fi51 - 128))
                                                                                    }
                                                                                }
                                                                                var v61 uint32 = 42
                                                                                v62 := GenS5{F0: (-1 >= 339), F1: 42, F2: 42}
                                                                                _ = v61
                                                                                _ = v62
                                                                            } else {
                                                                                if len(v47) >= 2 {
                                                                                    v47 = v47[1:len(v47)]
                                                                                }
                                                                                v32 = ((!true) || (!v32))
                                                                                var v63 string = v2
                                                                                delete(v1, 2)
                                                                                _ = v63
                                                                            }
                                                                            var v64 int32 = -1
                                                                            _ = v64
                                                                        } else {
                                                                            if v32 {
                                                                                v2 = "hello"
                                                                                if len(v47) >= 2 {
                                                                                    v47 = v47[1:len(v47)]
                                                                                }
                                                                                var v65 bool = v32
                                                                                _ = v65
                                                                            } else {
                                                                                v66 := v1[4294967295]
                                                                                _ = v66
                                                                            }
                                                                            if v32 {
                                                                                var v67 uint16 = 255
                                                                                v68 := GenS1{F0: v67}
                                                                                fmt.Println(v32, v2, v3)
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    if v32 {
                                                                        var v69 float32 = 1.5
                                                                        v70 := GenS5{F0: v32, F1: 32767, F2: -340}
                                                                        _ = v69
                                                                        _ = v70
                                                                    } else {
                                                                        if true {
                                                                            v32 = (!false)
                                                                        } else {
                                                                            v32 = ("hello" > (v2 + v2))
                                                                            var v71 uint64 = 364
                                                                            if (3.14159 != 1.5) {
                                                                                var v72 float32 = 0.0
                                                                                var v73 float32 = v72
                                                                                if ((((((fi45 ^ (fi45 + (fi45 ^ -2147483648))) >= fi45) && v32) || ((v71 & (v71 << ((v3) & 63))) >= v71)) && v32) && true) {
                                                                                    v74 := GenS2{F0: 1, F1: 1}
                                                                                    _ = v74
                                                                                }
                                                                                _ = v72
                                                                                _ = v73
                                                                            }
                                                                            _ = v71
                                                                        }
                                                                        switch (v2 + v2) {
                                                                        case "a":
                                                                            delete(v4, fi51)
                                                                        case "foo bar":
                                                                            v3 = 1
                                                                        case "\n\t":
                                                                            if len(v47) > 0 {
                                                                                v47[0] = 32768
                                                                            }
                                                                        default:
                                                                            var v75 int64 = -1
                                                                            _ = v75
                                                                        }
                                                                        v76 := GenS4{F0: 255}
                                                                        _ = v76
                                                                    }
                                                                default:
                                                                    v32 = (92 >= 12)
                                                                }
                                                            case "\n\t":
                                                                switch fi51 {
                                                                case 32767:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        switch 1 {
                                                                        case -1:
                                                                            v77 := GenS6{F0: v32, F1: v2, F2: 142}
                                                                            var v78 int64 = 275
                                                                            _ = v77
                                                                            _ = v78
                                                                        case -9223372036854775808:
                                                                            var v79 int64 = 9223372036854775807
                                                                            _ = v79
                                                                        }
                                                                    case "hello":
                                                                        var v80 int8 = 0
                                                                        _ = v80
                                                                    case "a":
                                                                        v81 := GenS5{F0: false, F1: 42, F2: 1}
                                                                        _ = v81
                                                                    }
                                                                case -9223372036854775808:
                                                                    v32 = true
                                                                default:
                                                                    if v32 {
                                                                        fmt.Println(v32)
                                                                        var v82 int8 = 1
                                                                        var v83 uint16 = 255
                                                                        _ = v82
                                                                        _ = v83
                                                                    } else {
                                                                        var v84 bool = (2 >= 128)
                                                                        _ = v84
                                                                    }
                                                                }
                                                            default:
                                                                v32 = (!(!(v2 < v2)))
                                                            }
                                                            switch (fi51 | 951) {
                                                            case 127:
                                                                var v85 uint = 1
                                                                if v32 {
                                                                    if (255 != (v85 % ((v85) | 1))) {
                                                                        var v86 int16 = 42
                                                                        var v87 int = (fi45 % ((fi45) | 1))
                                                                        _ = v86
                                                                        _ = v87
                                                                    }
                                                                    fmt.Println(fi51)
                                                                    if false {
                                                                        if true {
                                                                            v47 = append(v47, fi45)
                                                                            var v88 float64 = 1e10
                                                                            _ = v88
                                                                        } else {
                                                                            var v89 int16 = -829
                                                                            v90 := v1[65535]
                                                                            _ = v89
                                                                            _ = v90
                                                                        }
                                                                    } else {
                                                                        v91 := v4[fi45]
                                                                        v3 = 0
                                                                        var v92 uint32 = 65535
                                                                        _ = v91
                                                                        _ = v92
                                                                    }
                                                                } else {
                                                                    v85 += (^(v3))
                                                                }
                                                                v85 -= v3
                                                                _ = v85
                                                            }
                                                        } else {
                                                            if v32 {
                                                                if (v3 < (v3 - (v3 - (v3 * v3)))) {
                                                                    var v93 int = (fi45 / (((fi45 << ((0) & 63))) | 1))
                                                                    v2 = (v2 + v2)
                                                                    if ((((((((1e10 - -2.25) - (1e10 * 1e10)) + 1.5) * (0.1 + (float32(((float64(0.1) + (1.5 - ((0.1 - -2.25) * 1.5))) * 1.5)) + (1.5 / 1e10)))) < 1e10) || (((float32((float64((1e10 / float32(((float64(1.5) / 1.5) * 3.14159)))) + -2.25)) - -2.25) - float32(1e10)) <= (-2.25 * 3.14159))) || false) || (false && v32)) {
                                                                        fmt.Println(v93, fi51, v2)
                                                                        fmt.Println(v93, v93, v32)
                                                                    }
                                                                    _ = v93
                                                                }
                                                                v32 = (42 != 304)
                                                            }
                                                            var v94 float32 = (1e10 + 1e10)
                                                            _ = v94
                                                        }
                                                    }
                                                    var v95 float64 = (1e10 - 0.1)
                                                    v47 = append(v47, fi51)
                                                    _ = fi51
                                                    _ = v95
                                                }
                                            case "hello":
                                                var v96 string = v2
                                                _ = v96
                                            }
                                            v3 += (^(v3))
                                            var v97 uint = v3
                                            _ = v97
                                        }
                                        for fi98 := 0; fi98 < 20; fi98++ {
                                            var v99 uint = 837
                                            _ = fi98
                                            _ = v99
                                        }
                                        v3 = v3
                                    }
                                    _ = v47
                                } else {
                                    v3 += (v3 ^ (v3 >> (((v3 << (((^(v3))) & 63))) & 63)))
                                    if (!v32) {
                                        if true {
                                            v3 -= 42
                                        } else {
                                            var v100 string = (v2 + ("" + v2))
                                            v101 := GenS3{F0: 255}
                                            if (100 != -128) {
                                                var v102 int32 = -2147483648
                                                for fi103 := 0; fi103 < 8; fi103++ {
                                                    v104 := GenS7{F0: ("foo bar" + "a"), F1: 1}
                                                    v2 = (((v2 + ("hello" + ("hello" + ""))) + "\n\t") + "hello")
                                                    _ = fi103
                                                    _ = v104
                                                }
                                                _ = v102
                                            }
                                            _ = v100
                                            _ = v101
                                        }
                                        for fi105 := 0; fi105 < 7; fi105++ {
                                            if v32 {
                                                var v106 int32 = 0
                                                v107 := []int{(fi45 & 158)}
                                                _ = v106
                                                _ = v107
                                            } else {
                                                var v108 bool = (9223372036854775807 <= 1)
                                                var v109 int = fi45
                                                v110 := GenS6{F0: v108, F1: "", F2: 0}
                                                _ = v108
                                                _ = v109
                                                _ = v110
                                            }
                                            var v111 uint32 = 65535
                                            _ = fi105
                                            _ = v111
                                        }
                                    } else {
                                        var v112 int16 = 202
                                        v32 = false
                                        _ = v112
                                    }
                                }
                            }
                        default:
                            v4[fi45] = v32
                        }
                        var v113 uint16 = 65535
                        _ = v113
                    }
                }
                for fi114 := 0; fi114 < 15; fi114++ {
                    v32 = v32
                    v3 = (v3 & v3)
                    v32 = (-58 == -9223372036854775808)
                    _ = fi114
                }
                for fi115 := 0; fi115 < 16; fi115++ {
                    var v116 float64 = 1.5
                    v2 = (((v2 + v2) + "") + ((v2 + (v2 + ("foo bar" + (v2 + (v2 + ""))))) + v2))
                    _ = fi115
                    _ = v116
                }
            }
            _ = fi45
        }
        fmt.Println(v3, v32)
    }
    for fi117 := 0; fi117 < 16; fi117++ {
        var v118 string = (((("hello" + (v2 + v2)) + (v2 + v2)) + "hello") + v2)
        if (892 < 42) {
            for fi119 := 0; fi119 < 3; fi119++ {
                var v120 int32 = genFn0()
                v121 := []int{(fi119 - fi119)}
                _ = len(v1)
                v122 := GenS8{F0: 127, F1: genFn0(), F2: "foo bar"}
                if (255 <= 785) {
                    v123 := GenS9{F0: fi117, F1: 2147483647}
                    v124 := GenS9{F0: fi117, F1: 2}
                    v125 := GenS9{F0: fi117, F1: 703}
                    _ = v123
                    _ = v124
                    _ = v125
                }
                if (1 > 0) {
                    fmt.Println(v120)
                } else {
                    var v126 uint32 = 2147483648
                    _ = v126
                }
                _ = fi119
                _ = v120
                _ = v121
                _ = v122
            }
        }
        var v127 bool = ((1.5 / 0.1) == 3.14159)
        _ = fi117
        _ = v118
        _ = v127
    }
    v128 := []int{}
    v2 = "a"
    var v129 uint32 = 4294967295
    v2 = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v128:", len(v128), v128)
    fmt.Println("v129:", v129)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v128
    _ = v129
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v44
}
