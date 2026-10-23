package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 bool; F2 int64 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint; F1 string; F2 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 uint16; F1 uint; F2 string }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 int; F1 bool }
type GenS11 struct { F0 bool; F1 int64 }

func main() {
    var v1 int8 = 1
    v2 := make(map[int]string)
    var v3 int = 2147483647
    var v4 int = -1
    var v5 uint8 = 127
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int = 277
        _ = fi6
        _ = v7
    }
    v3 -= (-(v3))
    var v8 uint32 = 2147483647
    var v9 float32 = 0.0
    v1 += 0
    var v10 uint16 = 42
    for fi11 := 0; fi11 < 7; fi11++ {
        v3 -= (fi11 / ((v4) | 1))
        _ = fi11
    }
    if false {
        var v12 int64 = -796
        v9 += v9
        _ = v12
    } else {
        v8 -= (^(v8))
        v13 := GenS1{F0: 65535, F1: (v8 & (v8 - v8)), F2: v8}
        var v14 int32 = 42
        _ = v13
        _ = v14
    }
    if true {
        var v15 int64 = 0
        _ = v15
    }
    v16 := GenS2{F0: (v5 & v5)}
    var v17 uint16 = v10
    switch v3 {
    case 209:
        if true {
            switch v3 {
            case 4294967295:
                if ((^(v8)) != 2147483647) {
                    v18 := GenS3{F0: 1, F1: false, F2: 982}
                    _ = v18
                }
            }
            if false {
                var v19 int64 = -1
                for fi20 := 0; fi20 < 6; fi20++ {
                    var v21 uint32 = (v8 | 272)
                    _ = len(v2)
                    _ = fi20
                    _ = v21
                }
                for fi22 := 0; fi22 < 14; fi22++ {
                    if ((v8 % (((v8 >> ((0) & 63))) | 1)) == (v8 >> ((0) & 63))) {
                        if true {
                            delete(v2, 32768)
                            var v23 float64 = float64(float32(float64(v9)))
                            var v24 uint = 18446744073709551615
                            _ = v23
                            _ = v24
                        } else {
                            v10 = v17
                        }
                        if true {
                            var v25 float32 = ((float32(-2.25) * v9) / (1e10 + 0.0))
                            _ = v25
                        } else {
                            var v26 float64 = (float64((v9 / ((3.14159 / 3.14159) * 3.14159))) / (0.1 - float64(1.5)))
                            switch "foo bar" {
                            case "a":
                                var v27 float32 = 1.5
                                _ = v27
                            case "foo bar":
                                for fi28 := 0; fi28 < 2; fi28++ {
                                    v26 = (1e10 * (v26 - (float64((v9 + v9)) * v26)))
                                    var v29 int = int(v17)
                                    fmt.Println(v17)
                                    v1 -= v1
                                    if false {
                                        var v30 int64 = int64(v10)
                                        switch "\n\t" {
                                        case "\n\t":
                                            if true {
                                                var v31 float64 = v26
                                                fmt.Println(v3)
                                                _ = v31
                                            } else {
                                                if true {
                                                    if (!(((v3 >= (fi28 * (v3 - (fi28 & (v3 ^ (v3 << ((42) & 63))))))) && false) || true)) {
                                                        var v32 int = int(v5)
                                                        switch fi28 {
                                                        case -624:
                                                            v19 += 0
                                                        case 1:
                                                            fmt.Println(v9, v9, v10)
                                                        case -9223372036854775808:
                                                            var v33 uint16 = (v17 >> ((18446744073709551615) & 63))
                                                            _ = v33
                                                        }
                                                        var v34 int = v29
                                                        _ = v32
                                                        _ = v34
                                                    } else {
                                                        var v35 float64 = v26
                                                        _ = v35
                                                    }
                                                    var v36 int32 = 42
                                                    v19 = (-(v30))
                                                    _ = v36
                                                }
                                                if true {
                                                    v37 := []int{-821, -446, v3, fi28}
                                                    var v38 int = (v4 ^ (v3 >> ((42) & 63)))
                                                    _ = v37
                                                    _ = v38
                                                }
                                            }
                                        case "foo bar":
                                            if true {
                                                fmt.Println(v29, v3, v19)
                                                var v39 uint16 = 255
                                                _ = v39
                                            }
                                        case "a":
                                            fmt.Println(v30, v19, v29)
                                        default:
                                            var v40 uint8 = (^(v5))
                                            if true {
                                                switch "a" {
                                                case "foo bar":
                                                    var v41 uint64 = 4294967295
                                                    _ = v41
                                                case "":
                                                    v42 := GenS4{F0: 42, F1: v8}
                                                    _ = v42
                                                case "\n\t":
                                                    switch "" {
                                                    case "hello":
                                                        v1 -= (v1 * 59)
                                                    case "":
                                                        v3 += v3
                                                    case "a":
                                                        fmt.Println(v40)
                                                        delete(v2, 402)
                                                        v43 := GenS3{F0: (fi22 & 32767), F1: true, F2: int64(v40)}
                                                        var v44 bool = (v26 > -2.25)
                                                        _ = v43
                                                        _ = v44
                                                    default:
                                                        switch int(v1) {
                                                        case 32768:
                                                            fmt.Println(v4)
                                                        case -769:
                                                            var v45 int16 = 42
                                                            _ = v45
                                                        case 2147483647:
                                                            switch "foo bar" {
                                                            case "hello":
                                                                if false {
                                                                    if false {
                                                                        delete(v2, v29)
                                                                        v19 = v19
                                                                    } else {
                                                                        var v46 bool = (1 < 512)
                                                                        var v47 float64 = 1.5
                                                                        _ = v46
                                                                        _ = v47
                                                                    }
                                                                    if false {
                                                                        var v48 int64 = (v30 | (v30 + v30))
                                                                        var v49 uint64 = 374
                                                                        switch int(v30) {
                                                                        case 4294967295:
                                                                            if true {
                                                                                v50 := GenS2{F0: 59}
                                                                                var v51 uint8 = (v5 | v5)
                                                                                _ = v50
                                                                                _ = v51
                                                                            }
                                                                        case -1:
                                                                            if true {
                                                                                fmt.Println(v30)
                                                                            } else {
                                                                                var v52 int16 = 0
                                                                                _ = v52
                                                                            }
                                                                        }
                                                                        fmt.Println(v26, v17)
                                                                        _ = v48
                                                                        _ = v49
                                                                    } else {
                                                                        if ((v26 <= v26) || ((v17 ^ (^(v10))) != (v17 + (^(v10))))) {
                                                                            delete(v2, (v4 - v3))
                                                                            v10 += v10
                                                                            switch (fi22 % (((v29 % ((fi28) | 1))) | 1)) {
                                                                            case 4294967295:
                                                                                if (false || (32767 < 1000)) {
                                                                                    if (((v9 + -2.25) / float32((0.1 - v26))) != (v9 + 3.14159)) {
                                                                                        var v53 int16 = -110
                                                                                        v8 -= v8
                                                                                        _ = v53
                                                                                    }
                                                                                    switch v29 {
                                                                                    case 2147483647:
                                                                                        var v54 uint = 255
                                                                                        _ = v54
                                                                                    case 4294967295:
                                                                                        var v55 float64 = float64(1.5)
                                                                                        _ = v55
                                                                                    case -1:
                                                                                        fmt.Println(v1)
                                                                                    default:
                                                                                        v56 := GenS5{F0: 1, F1: "\n\t", F2: v1}
                                                                                        _ = v56
                                                                                    }
                                                                                    var v57 int = fi22
                                                                                    _ = v57
                                                                                }
                                                                            case -2147483648:
                                                                                v19 -= -328
                                                                            }
                                                                        }
                                                                        switch fi22 {
                                                                        case -613:
                                                                            v58 := []int{}
                                                                            _ = v58
                                                                        case 445:
                                                                            switch fi28 {
                                                                            case 128:
                                                                                if ((42 != 391) || false) {
                                                                                    if true {
                                                                                        if false {
                                                                                            v17 = v17
                                                                                            switch ("\n\t" + "foo bar") {
                                                                                            case "hello":
                                                                                                v17 -= (v17 + (v17 & (^(v10))))
                                                                                            case "":
                                                                                                var v59 uint8 = (v5 / ((v40) | 1))
                                                                                                _ = v59
                                                                                            default:
                                                                                                switch "a" {
                                                                                                case "hello":
                                                                                                    var v60 bool = true
                                                                                                    _ = v60
                                                                                                case "a":
                                                                                                    var v61 float32 = v9
                                                                                                    _ = v61
                                                                                                default:
                                                                                                    if (false || false) {
                                                                                                        var v62 uint8 = v40
                                                                                                        _ = v62
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            if true {
                                                                                                if (350 >= 32767) {
                                                                                                    var v63 string = "\n\t"
                                                                                                    _ = v63
                                                                                                }
                                                                                                if false {
                                                                                                    if false {
                                                                                                        var v64 uint16 = 692
                                                                                                        v65 := v2[(v3 * (fi22 / ((42) | 1)))]
                                                                                                        v3 -= v29
                                                                                                        v30 = (v19 / ((766) | 1))
                                                                                                        switch fi28 {
                                                                                                        case 128:
                                                                                                            v66 := []int{v29}
                                                                                                            _ = v66
                                                                                                        case 32767:
                                                                                                            switch v65 {
                                                                                                            case "":
                                                                                                                v67 := []int{(v4 | 128)}
                                                                                                                _ = v67
                                                                                                            case "a":
                                                                                                                if (42 == 65535) {
                                                                                                                    if true {
                                                                                                                        v2[191] = v65
                                                                                                                        if true {
                                                                                                                            switch "a" {
                                                                                                                            case "a":
                                                                                                                                if true {
                                                                                                                                    fmt.Println(v30, v29, v64)
                                                                                                                                    v19 = (v19 / ((v19) | 1))
                                                                                                                                    var v68 uint32 = v8
                                                                                                                                    _ = v68
                                                                                                                                } else {
                                                                                                                                    v69 := GenS2{F0: 2}
                                                                                                                                    _ = v69
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                switch (v29 & v3) {
                                                                                                                                case -2147483648:
                                                                                                                                    v70 := GenS2{F0: (v5 ^ (v5 >> ((255) & 63)))}
                                                                                                                                    _ = v70
                                                                                                                                case 2147483647:
                                                                                                                                    v19 += v30
                                                                                                                                default:
                                                                                                                                    var v71 int16 = -32768
                                                                                                                                    _ = v71
                                                                                                                                }
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            var v72 int = (-(v3))
                                                                                                                            v3 -= v29
                                                                                                                            _ = v72
                                                                                                                        }
                                                                                                                        var v73 float32 = v9
                                                                                                                        if false {
                                                                                                                            var v74 int16 = 42
                                                                                                                            _ = v74
                                                                                                                        }
                                                                                                                        _ = v73
                                                                                                                    } else {
                                                                                                                        v75 := []int{v4, (v4 / (((fi28 >> ((1) & 63))) | 1))}
                                                                                                                        _ = v75
                                                                                                                    }
                                                                                                                    v9 += (3.14159 / float32(0.1))
                                                                                                                    fmt.Println(v64, v64)
                                                                                                                } else {
                                                                                                                    var v76 float64 = v26
                                                                                                                    _ = v76
                                                                                                                }
                                                                                                            }
                                                                                                        case 32768:
                                                                                                            var v77 float32 = ((-2.25 / v9) * 0.1)
                                                                                                            _ = v77
                                                                                                        default:
                                                                                                            if false {
                                                                                                                v78 := []int{536, (v29 - 4294967295)}
                                                                                                                _ = v78
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v64
                                                                                                        _ = v65
                                                                                                    }
                                                                                                    v3 += fi28
                                                                                                    var v79 uint8 = v40
                                                                                                    _ = v79
                                                                                                }
                                                                                                var v80 int16 = 48
                                                                                                _ = v80
                                                                                            }
                                                                                        } else {
                                                                                            v81 := GenS6{F0: 65535}
                                                                                            _ = v81
                                                                                        }
                                                                                        if false {
                                                                                            _ = len(v2)
                                                                                            if false {
                                                                                                var v82 int16 = 0
                                                                                                fmt.Println(v82, v5)
                                                                                                fmt.Println(fi22)
                                                                                                _ = v82
                                                                                            }
                                                                                            v83 := []int{v4, fi22}
                                                                                            _ = v83
                                                                                        }
                                                                                        fmt.Println(v30, fi22, v8)
                                                                                    } else {
                                                                                        var v84 uint32 = v8
                                                                                        v30 = 9223372036854775807
                                                                                        v85 := GenS7{F0: -128, F1: 42}
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                    }
                                                                                    var v86 string = "foo bar"
                                                                                    v1 -= (v1 ^ (v1 - (-(v1))))
                                                                                    _ = v86
                                                                                } else {
                                                                                    v10 += 1
                                                                                }
                                                                            }
                                                                        case 1:
                                                                            var v87 int16 = 596
                                                                            _ = v87
                                                                        default:
                                                                            v1 = (^(v1))
                                                                        }
                                                                        switch (v4 ^ v4) {
                                                                        case 128:
                                                                            var v88 int32 = -481
                                                                            _ = v88
                                                                        default:
                                                                            fmt.Println(v17, v19)
                                                                        }
                                                                        v8 = v8
                                                                        if false {
                                                                            var v89 int = v29
                                                                            v8 = v8
                                                                            switch "hello" {
                                                                            case "":
                                                                                v1 -= 0
                                                                            default:
                                                                                if ((1000 < 32767) || (42 >= 2)) {
                                                                                    var v90 int64 = -9223372036854775808
                                                                                    if false {
                                                                                        v91 := GenS7{F0: v1, F1: 2}
                                                                                        var v92 uint = 42
                                                                                        delete(v2, (v4 / (((v89 / ((v29) | 1))) | 1)))
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                    }
                                                                                    switch (^(v4)) {
                                                                                    case 9223372036854775807:
                                                                                        v93 := []int{}
                                                                                        _ = v93
                                                                                    case -1:
                                                                                        if ((v10 + (v17 - v10)) > v17) {
                                                                                            var v94 bool = true
                                                                                            var v95 int8 = 13
                                                                                            fmt.Println(v89, v9)
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                        } else {
                                                                                            v96 := []int{}
                                                                                            var v97 bool = false
                                                                                            _ = v96
                                                                                            _ = v97
                                                                                        }
                                                                                    default:
                                                                                        v40 = 234
                                                                                    }
                                                                                    _ = v90
                                                                                } else {
                                                                                    var v98 uint = 675
                                                                                    if true {
                                                                                        switch fi22 {
                                                                                        case 127:
                                                                                            var v99 int = (fi22 ^ (fi22 ^ v89))
                                                                                            v10 = 42
                                                                                            _ = v99
                                                                                        case 32768:
                                                                                            switch "a" {
                                                                                            case "":
                                                                                                var v100 int32 = -650
                                                                                                _ = v100
                                                                                            case "\n\t":
                                                                                                var v101 uint64 = 501
                                                                                                switch "foo bar" {
                                                                                                case "":
                                                                                                    var v102 int16 = 42
                                                                                                    _ = v102
                                                                                                }
                                                                                                _ = v101
                                                                                            case "hello":
                                                                                                v103 := GenS8{F0: (v17 | 42), F1: uint(v8), F2: "\n\t"}
                                                                                                _ = v103
                                                                                            default:
                                                                                                v104 := []int{}
                                                                                                _ = v104
                                                                                            }
                                                                                        case 9223372036854775807:
                                                                                            if ("hello" > "a") {
                                                                                                var v105 uint16 = v17
                                                                                                var v106 uint64 = 1
                                                                                                _ = v105
                                                                                                _ = v106
                                                                                            }
                                                                                        }
                                                                                        if false {
                                                                                            if false {
                                                                                                if false {
                                                                                                    v107 := GenS7{F0: v1, F1: uint64(v40)}
                                                                                                    _ = v107
                                                                                                }
                                                                                            } else {
                                                                                                delete(v2, (fi22 | v29))
                                                                                                if (!((v10 / ((v17) | 1)) == (v17 - 250))) {
                                                                                                    v108 := GenS2{F0: v40}
                                                                                                    switch 42 {
                                                                                                    case -765:
                                                                                                        if (!(v98 <= (v98 ^ 0))) {
                                                                                                            var v109 string = ""
                                                                                                            _ = v109
                                                                                                        }
                                                                                                    case 922:
                                                                                                        if (!false) {
                                                                                                            var v110 int32 = -113
                                                                                                            if (false || true) {
                                                                                                                v111 := GenS4{F0: -351, F1: (v8 % ((v8) | 1))}
                                                                                                                var v112 float64 = ((v26 - v26) - v26)
                                                                                                                _ = v111
                                                                                                                _ = v112
                                                                                                            }
                                                                                                            fmt.Println(v98, v8)
                                                                                                            _ = v110
                                                                                                        } else {
                                                                                                            v113 := []int{-136, v4, 2147483647}
                                                                                                            _ = v113
                                                                                                        }
                                                                                                    case -1:
                                                                                                        var v114 float32 = float32(float64(-2.25))
                                                                                                        _ = v114
                                                                                                    default:
                                                                                                        v115 := v2[(-(v4))]
                                                                                                        _ = v115
                                                                                                    }
                                                                                                    v89 += (v89 + (fi28 >> ((v98) & 63)))
                                                                                                    _ = v108
                                                                                                }
                                                                                                var v116 int32 = -2147483648
                                                                                                var v117 string = "hello"
                                                                                                _ = v116
                                                                                                _ = v117
                                                                                            }
                                                                                        } else {
                                                                                            v118 := GenS7{F0: (v1 % ((v1) | 1)), F1: 2}
                                                                                            v1 += (v1 + (^(v1)))
                                                                                            var v119 float64 = 0.1
                                                                                            v2[(fi28 % ((v3) | 1))] = "a"
                                                                                            _ = v118
                                                                                            _ = v119
                                                                                        }
                                                                                    }
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        fmt.Println(v9)
                                                                                    case "foo bar":
                                                                                        var v120 int32 = -637
                                                                                        _ = v120
                                                                                    case "hello":
                                                                                        v26 += ((1e10 / (v26 - (float64(float32(v26)) / ((v26 * v26) / (v26 + v26))))) * v26)
                                                                                    default:
                                                                                        if true {
                                                                                            delete(v2, int(v98))
                                                                                            v30 = 645
                                                                                            var v121 int64 = -476
                                                                                            switch "\n\t" {
                                                                                            case "":
                                                                                                var v122 uint8 = v5
                                                                                                _ = v122
                                                                                            default:
                                                                                                v10 = (v10 >> ((788) & 63))
                                                                                            }
                                                                                            v123 := GenS9{F0: 346, F1: true}
                                                                                            v98 = v98
                                                                                            _ = v121
                                                                                            _ = v123
                                                                                        } else {
                                                                                            fmt.Println(v29, v8, fi22)
                                                                                            var v124 uint16 = 65535
                                                                                            v17 -= uint16(v5)
                                                                                            var v125 string = "foo bar"
                                                                                            var v126 int16 = -1
                                                                                            v30 -= v30
                                                                                            _ = v124
                                                                                            _ = v125
                                                                                            _ = v126
                                                                                        }
                                                                                    }
                                                                                    _ = v98
                                                                                }
                                                                            }
                                                                            _ = v89
                                                                        }
                                                                    }
                                                                    var v127 bool = false
                                                                    if true {
                                                                        var v128 int8 = (v1 ^ (v1 ^ (v1 + (v1 + (v1 - v1)))))
                                                                        v129 := []int{v3, v3, int(v10)}
                                                                        v1 = (v128 + v128)
                                                                        _ = v128
                                                                        _ = v129
                                                                    }
                                                                    if (!(127 <= -9223372036854775808)) {
                                                                        if (127 <= (-(v1))) {
                                                                            var v130 string = "hello"
                                                                            var v131 int16 = 32767
                                                                            _ = v130
                                                                            _ = v131
                                                                        }
                                                                        v4 += (^(fi28))
                                                                        var v132 float32 = (v9 / 1e10)
                                                                        var v133 int8 = (v1 | -51)
                                                                        var v134 uint = 42
                                                                        _ = v132
                                                                        _ = v133
                                                                        _ = v134
                                                                    } else {
                                                                        v135 := GenS4{F0: int16(v8), F1: 65535}
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            var v136 uint64 = 0
                                                                            _ = v136
                                                                        default:
                                                                            fmt.Println(v19, v8)
                                                                        }
                                                                        fmt.Println(v5)
                                                                        _ = v135
                                                                    }
                                                                    fmt.Println(v40, v9)
                                                                    _ = v127
                                                                } else {
                                                                    if true {
                                                                        v137 := []int{(fi22 % ((-9223372036854775808) | 1)), fi22, (v4 / ((1) | 1)), 964}
                                                                        v29 = v4
                                                                        var v138 string = ""
                                                                        _ = v137
                                                                        _ = v138
                                                                    } else {
                                                                        fmt.Println(v9)
                                                                        v26 += (v26 - v26)
                                                                        switch (fi22 >> ((65535) & 63)) {
                                                                        case 449:
                                                                            v139 := []int{(v29 << ((4294967295) & 63))}
                                                                            _ = v139
                                                                        }
                                                                    }
                                                                    v140 := GenS4{F0: 32767, F1: (v8 * (v8 << ((42) & 63)))}
                                                                    _ = v140
                                                                }
                                                                if (!false) {
                                                                    if false {
                                                                        v141 := []int{(v4 >> ((4294967295) & 63)), (v29 << ((370) & 63))}
                                                                        v142 := GenS7{F0: (v1 + 0), F1: 1}
                                                                        _ = v141
                                                                        _ = v142
                                                                    }
                                                                    var v143 uint32 = (v8 << ((255) & 63))
                                                                    v144 := v2[(-(fi22))]
                                                                    _ = v143
                                                                    _ = v144
                                                                }
                                                                fmt.Println(v5, fi28)
                                                            case "\n\t":
                                                                var v145 int64 = v30
                                                                _ = v145
                                                            default:
                                                                fmt.Println(v8, v9, v10)
                                                            }
                                                        }
                                                    }
                                                }
                                                v1 -= (v1 / (((^(v1))) | 1))
                                            }
                                            switch (fi28 | (v3 * v29)) {
                                            case 0:
                                                var v146 uint = 590
                                                v4 += fi28
                                                var v147 float32 = float32(v26)
                                                _ = v146
                                                _ = v147
                                            default:
                                                v148 := GenS10{F0: fi22, F1: (!((v30 - v19) == (v19 >> ((68) & 63))))}
                                                _ = v148
                                            }
                                            _ = v40
                                        }
                                        _ = v30
                                    }
                                    _ = fi28
                                    _ = v29
                                }
                            default:
                                fmt.Println(fi22)
                            }
                            for fi149 := 0; fi149 < 4; fi149++ {
                                v1 = 1
                                v2[v3] = ("hello" + "")
                                if (("" + "hello") >= "") {
                                    v150 := []int{}
                                    _ = v150
                                }
                                v3 += (fi149 * (fi149 << ((4294967295) & 63)))
                                var v151 int16 = 158
                                _ = fi149
                                _ = v151
                            }
                            _ = v26
                        }
                        var v152 int = (v3 & (-(fi22)))
                        _ = v152
                    } else {
                        if false {
                            var v153 int = 4294967295
                            if (-32768 == -124) {
                                var v154 int8 = (v1 >> ((2) & 63))
                                fmt.Println(v5)
                                _ = v154
                            } else {
                                var v155 float32 = float32(0.1)
                                var v156 int8 = v1
                                _ = v155
                                _ = v156
                            }
                            v19 = v19
                            _ = v153
                        } else {
                            switch ("a" + "") {
                            case "":
                                var v157 int8 = v1
                                _ = v157
                            case "\n\t":
                                var v158 float64 = 1.5
                                _ = v158
                            case "foo bar":
                                v19 = 42
                            }
                            var v159 int = (v4 << ((1) & 63))
                            _ = v159
                        }
                        v1 = -128
                        v160 := []int{(v4 % (((v3 - v4)) | 1)), v4, 950, (fi22 & int(v5))}
                        _ = v160
                    }
                    _ = fi22
                }
                _ = v19
            }
        }
    default:
        fmt.Println(v1, v1)
    }
    for fi161 := 0; fi161 < 14; fi161++ {
        v3 += v4
        fmt.Println(fi161, v3, v4)
        var v162 uint = 573
        _ = fi161
        _ = v162
    }
    if ((^(v5)) >= v5) {
        v9 = 0.1
    } else {
        for fi163 := 0; fi163 < 10; fi163++ {
            switch "\n\t" {
            case "foo bar":
                if ("hello" > "foo bar") {
                    switch "foo bar" {
                    case "\n\t":
                        for fi164 := 0; fi164 < 18; fi164++ {
                            v3 = int(v10)
                            v1 = 42
                            var v165 uint8 = (v5 & uint8(v8))
                            if true {
                                v166 := GenS1{F0: 89, F1: v8, F2: 2}
                                fmt.Println(v10, v10)
                                var v167 uint32 = uint32(v165)
                                _ = v166
                                _ = v167
                            }
                            v168 := GenS11{F0: (true || true), F1: -9223372036854775808}
                            _ = fi164
                            _ = v165
                            _ = v168
                        }
                    case "foo bar":
                        v3 = v3
                    case "hello":
                        var v169 uint64 = 2
                        _ = v169
                    default:
                        var v170 uint16 = v17
                        _ = v170
                    }
                    var v171 float32 = ((1e10 / v9) / 0.1)
                    _ = v171
                } else {
                    v5 += (v5 % (((v5 & v5)) | 1))
                    var v172 string = "\n\t"
                    _ = v172
                }
            }
            _ = fi163
        }
    }
    if (-9223372036854775808 < 0) {
        v173 := []int{(v4 & (v4 + v4)), v3}
        v174 := GenS10{F0: v4, F1: false}
        _ = v173
        _ = v174
    }
    for fi175 := 0; fi175 < 1; fi175++ {
        var v176 uint64 = 2
        _ = fi175
        _ = v176
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
